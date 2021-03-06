// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Magicking/Biffin/models"
)

// PostMapOKCode is the HTTP code returned for type PostMapOK
const PostMapOKCode int = 200

/*PostMapOK The ID of the new map

swagger:response postMapOK
*/
type PostMapOK struct {

	/*
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewPostMapOK creates PostMapOK with default headers values
func NewPostMapOK() *PostMapOK {

	return &PostMapOK{}
}

// WithPayload adds the payload to the post map o k response
func (o *PostMapOK) WithPayload(payload int64) *PostMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post map o k response
func (o *PostMapOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PostMapDefault Unexpected error

swagger:response postMapDefault
*/
type PostMapDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMapDefault creates PostMapDefault with default headers values
func NewPostMapDefault(code int) *PostMapDefault {
	if code <= 0 {
		code = 500
	}

	return &PostMapDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post map default response
func (o *PostMapDefault) WithStatusCode(code int) *PostMapDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post map default response
func (o *PostMapDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post map default response
func (o *PostMapDefault) WithPayload(payload *models.Error) *PostMapDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post map default response
func (o *PostMapDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMapDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
