package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

// HTTP code for type PostDhcpLeaseCreated
const PostDhcpLeaseCreatedCode int = 201

/*PostDhcpLeaseCreated post dhcp lease created

swagger:response postDhcpLeaseCreated
*/
type PostDhcpLeaseCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DhcpLeaseInput `json:"body,omitempty"`
}

// NewPostDhcpLeaseCreated creates PostDhcpLeaseCreated with default headers values
func NewPostDhcpLeaseCreated() *PostDhcpLeaseCreated {
	return &PostDhcpLeaseCreated{}
}

// WithPayload adds the payload to the post dhcp lease created response
func (o *PostDhcpLeaseCreated) WithPayload(payload *models.DhcpLeaseInput) *PostDhcpLeaseCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dhcp lease created response
func (o *PostDhcpLeaseCreated) SetPayload(payload *models.DhcpLeaseInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDhcpLeaseCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PostDhcpLeaseUnauthorized
const PostDhcpLeaseUnauthorizedCode int = 401

/*PostDhcpLeaseUnauthorized post dhcp lease unauthorized

swagger:response postDhcpLeaseUnauthorized
*/
type PostDhcpLeaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDhcpLeaseUnauthorized creates PostDhcpLeaseUnauthorized with default headers values
func NewPostDhcpLeaseUnauthorized() *PostDhcpLeaseUnauthorized {
	return &PostDhcpLeaseUnauthorized{}
}

// WithPayload adds the payload to the post dhcp lease unauthorized response
func (o *PostDhcpLeaseUnauthorized) WithPayload(payload *models.Error) *PostDhcpLeaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dhcp lease unauthorized response
func (o *PostDhcpLeaseUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDhcpLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PostDhcpLeaseConflict
const PostDhcpLeaseConflictCode int = 409

/*PostDhcpLeaseConflict post dhcp lease conflict

swagger:response postDhcpLeaseConflict
*/
type PostDhcpLeaseConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDhcpLeaseConflict creates PostDhcpLeaseConflict with default headers values
func NewPostDhcpLeaseConflict() *PostDhcpLeaseConflict {
	return &PostDhcpLeaseConflict{}
}

// WithPayload adds the payload to the post dhcp lease conflict response
func (o *PostDhcpLeaseConflict) WithPayload(payload *models.Error) *PostDhcpLeaseConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dhcp lease conflict response
func (o *PostDhcpLeaseConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDhcpLeaseConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PostDhcpLeaseInternalServerError
const PostDhcpLeaseInternalServerErrorCode int = 500

/*PostDhcpLeaseInternalServerError post dhcp lease internal server error

swagger:response postDhcpLeaseInternalServerError
*/
type PostDhcpLeaseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDhcpLeaseInternalServerError creates PostDhcpLeaseInternalServerError with default headers values
func NewPostDhcpLeaseInternalServerError() *PostDhcpLeaseInternalServerError {
	return &PostDhcpLeaseInternalServerError{}
}

// WithPayload adds the payload to the post dhcp lease internal server error response
func (o *PostDhcpLeaseInternalServerError) WithPayload(payload *models.Error) *PostDhcpLeaseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dhcp lease internal server error response
func (o *PostDhcpLeaseInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDhcpLeaseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}