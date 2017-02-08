package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*ListMachinesOK list machines o k

swagger:response listMachinesOK
*/
type ListMachinesOK struct {

	/*
	  In: Body
	*/
	Payload ListMachinesOKBody `json:"body,omitempty"`
}

// NewListMachinesOK creates ListMachinesOK with default headers values
func NewListMachinesOK() *ListMachinesOK {
	return &ListMachinesOK{}
}

// WithPayload adds the payload to the list machines o k response
func (o *ListMachinesOK) WithPayload(payload ListMachinesOKBody) *ListMachinesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list machines o k response
func (o *ListMachinesOK) SetPayload(payload ListMachinesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMachinesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ListMachinesUnauthorized list machines unauthorized

swagger:response listMachinesUnauthorized
*/
type ListMachinesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewListMachinesUnauthorized creates ListMachinesUnauthorized with default headers values
func NewListMachinesUnauthorized() *ListMachinesUnauthorized {
	return &ListMachinesUnauthorized{}
}

// WithPayload adds the payload to the list machines unauthorized response
func (o *ListMachinesUnauthorized) WithPayload(payload *models.Result) *ListMachinesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list machines unauthorized response
func (o *ListMachinesUnauthorized) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMachinesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListMachinesNotFound list machines not found

swagger:response listMachinesNotFound
*/
type ListMachinesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewListMachinesNotFound creates ListMachinesNotFound with default headers values
func NewListMachinesNotFound() *ListMachinesNotFound {
	return &ListMachinesNotFound{}
}

// WithPayload adds the payload to the list machines not found response
func (o *ListMachinesNotFound) WithPayload(payload *models.Result) *ListMachinesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list machines not found response
func (o *ListMachinesNotFound) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMachinesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListMachinesInternalServerError list machines internal server error

swagger:response listMachinesInternalServerError
*/
type ListMachinesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewListMachinesInternalServerError creates ListMachinesInternalServerError with default headers values
func NewListMachinesInternalServerError() *ListMachinesInternalServerError {
	return &ListMachinesInternalServerError{}
}

// WithPayload adds the payload to the list machines internal server error response
func (o *ListMachinesInternalServerError) WithPayload(payload *models.Result) *ListMachinesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list machines internal server error response
func (o *ListMachinesInternalServerError) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMachinesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}