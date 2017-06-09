package backend

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func randString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("Failed to read random\n")
		return "ARGH!"
	}
	base64 := base64.URLEncoding.EncodeToString(b)
	return base64[:n]
}

// Config configures a Manager.
type JwtConfig struct {
	// digital signing method, defaults to jwt.SigningMethodHS256 (SHA256)
	Method jwt.SigningMethod
}

// Manager is a JSON Web Token (JWT) Provider which create or retrieves tokens
// with a particular signing key and options.
type JwtManager struct {
	key    []byte
	method jwt.SigningMethod
}

// New creates a new Manager which provides JWTs using the given signing key.
// Defaults to signing with SHA256 HMAC (jwt.SigningMethodHS256)
func NewJwtManager(key []byte, configs ...JwtConfig) *JwtManager {
	var c JwtConfig
	if len(configs) == 0 {
		c = JwtConfig{}
	} else {
		c = configs[0]
	}
	m := &JwtManager{
		key:    key,
		method: c.Method,
	}
	m.setDefaults()
	return m
}

func (m *JwtManager) setDefaults() {
	if m.method == nil {
		m.method = jwt.SigningMethodHS256
	}
}

// getKey accepts an unverified JWT and returns the signing/verification key.
// Also ensures tha the token's algorithm matches the signing method expected
// by the manager.
func (m *JwtManager) getKey(unverified *jwt.Token) (interface{}, error) {
	// require token alg to match the set signing method, do not allow none
	if meth := unverified.Method; meth == nil || meth.Alg() != m.method.Alg() {
		return nil, jwt.ErrHashUnavailable
	}
	return m.key, nil
}

func encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))

	// iv =  initialization vector
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func decrypt(key []byte, b64ciphertext string) (string, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(b64ciphertext)

	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// Sign digitally signs a *jwt.Token using the token's method and the manager's
// signing key to return a string
func (m *JwtManager) sign(token *jwt.Token) (string, error) {
	jwtString, err := token.SignedString(m.key)
	if err != nil {
		return "", err
	}

	return encrypt(m.key, jwtString)
}

// Claim is an individial specifier for something we are allowed access to.
type Claim struct {
	Scope    string `json:"scope"`
	Action   string `json:"action"`
	Specific string `json:"specific"`
}

// Match tests to see if this claim allows access for the specified
// scope, action, and specific item.
//
// If the Claim has `*` for any field, it matches all possible values
// for that field.
func (c *Claim) Match(scope, action, specific string) bool {
	return (c.Scope == scope || c.Scope == "*") &&
		(c.Action == action || c.Action == "*") &&
		(c.Specific == specific || c.Specific == "*")
}

// DrpCustomClaims is a JWT token that contains a list of all the
// things this token allows access to.
type DrpCustomClaims struct {
	DrpClaims []Claim `json:"drp_claims"`
	jwt.StandardClaims
}

// Match tests all the claims in this Token to find one that matches.
func (d *DrpCustomClaims) Match(scope, action, specific string) bool {
	for _, claim := range d.DrpClaims {
		if claim.Match(scope, action, specific) {
			return true
		}
	}
	return false
}

// NewClaim creates a new, unsigned Token that doesn't allow access to anything.
// You must call Seal() to turn this into a signed JWT token.
func NewClaim(user string, ttl int) *DrpCustomClaims {
	d := time.Duration(ttl) * time.Second
	res := &DrpCustomClaims{DrpClaims: []Claim{}}
	res.IssuedAt = time.Now().Unix()
	res.ExpiresAt = time.Now().Add(d).Unix()
	res.Issuer = "digitalrebar provision"
	res.Id = user
	return res
}

// Add adds a discrete Claim to our custom Token class.
func (d *DrpCustomClaims) Add(scope, action, specific string) *DrpCustomClaims {
	d.DrpClaims = append(d.DrpClaims, Claim{scope, action, specific})
	return d
}

// Seal turns our custom Token class into a signed JWT Token.
func (d *DrpCustomClaims) Seal(m *JwtManager) (string, error) {
	return m.sign(jwt.NewWithClaims(m.method, d))
}

// Get gets the signed JWT from the Authorization header. If the token is
// missing, expired, or the signature does not validate, returns an error.
func (m *JwtManager) get(encTokenString string) (*DrpCustomClaims, error) {
	tokenString, err := decrypt(m.key, encTokenString)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &DrpCustomClaims{}, m.getKey)
	if err != nil {
		return nil, err
	}

	if drpCustomClaim, ok := token.Claims.(*DrpCustomClaims); !ok {
		return nil, errors.New(fmt.Sprintf("Missing claim structure: %v\n", token))
	} else {
		return drpCustomClaim, nil
	}
}