// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// CreateUserCreatedCode is the HTTP code returned for type CreateUserCreated
const CreateUserCreatedCode int = 201

/*CreateUserCreated successful response

swagger:response createUserCreated
*/
type CreateUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreateUserResponseResource `json:"body,omitempty"`
}

// NewCreateUserCreated creates CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {

	return &CreateUserCreated{}
}

// WithPayload adds the payload to the create user created response
func (o *CreateUserCreated) WithPayload(payload *models.CreateUserResponseResource) *CreateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user created response
func (o *CreateUserCreated) SetPayload(payload *models.CreateUserResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}