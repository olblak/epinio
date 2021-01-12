// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveAppFromRouteCreatedCode is the HTTP code returned for type RemoveAppFromRouteCreated
const RemoveAppFromRouteCreatedCode int = 201

/*RemoveAppFromRouteCreated successful response

swagger:response removeAppFromRouteCreated
*/
type RemoveAppFromRouteCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveAppFromRouteResponseResource `json:"body,omitempty"`
}

// NewRemoveAppFromRouteCreated creates RemoveAppFromRouteCreated with default headers values
func NewRemoveAppFromRouteCreated() *RemoveAppFromRouteCreated {

	return &RemoveAppFromRouteCreated{}
}

// WithPayload adds the payload to the remove app from route created response
func (o *RemoveAppFromRouteCreated) WithPayload(payload *models.RemoveAppFromRouteResponseResource) *RemoveAppFromRouteCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove app from route created response
func (o *RemoveAppFromRouteCreated) SetPayload(payload *models.RemoveAppFromRouteResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAppFromRouteCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}