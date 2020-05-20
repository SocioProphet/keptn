// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/mongodb-datastore/models"
)

// GetOpenEventsOKCode is the HTTP code returned for type GetOpenEventsOK
const GetOpenEventsOKCode int = 200

/*GetOpenEventsOK ok

swagger:response getOpenEventsOK
*/
type GetOpenEventsOK struct {

	/*
	  In: Body
	*/
	Payload *GetOpenEventsOKBody `json:"body,omitempty"`
}

// NewGetOpenEventsOK creates GetOpenEventsOK with default headers values
func NewGetOpenEventsOK() *GetOpenEventsOK {

	return &GetOpenEventsOK{}
}

// WithPayload adds the payload to the get open events o k response
func (o *GetOpenEventsOK) WithPayload(payload *GetOpenEventsOKBody) *GetOpenEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get open events o k response
func (o *GetOpenEventsOK) SetPayload(payload *GetOpenEventsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOpenEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetOpenEventsDefault error

swagger:response getOpenEventsDefault
*/
type GetOpenEventsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOpenEventsDefault creates GetOpenEventsDefault with default headers values
func NewGetOpenEventsDefault(code int) *GetOpenEventsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetOpenEventsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get open events default response
func (o *GetOpenEventsDefault) WithStatusCode(code int) *GetOpenEventsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get open events default response
func (o *GetOpenEventsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get open events default response
func (o *GetOpenEventsDefault) WithPayload(payload *models.Error) *GetOpenEventsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get open events default response
func (o *GetOpenEventsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOpenEventsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}