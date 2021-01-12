// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveAuditorFromOrganizationCreatedCode is the HTTP code returned for type RemoveAuditorFromOrganizationCreated
const RemoveAuditorFromOrganizationCreatedCode int = 201

/*RemoveAuditorFromOrganizationCreated successful response

swagger:response removeAuditorFromOrganizationCreated
*/
type RemoveAuditorFromOrganizationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveAuditorFromOrganizationResponseResource `json:"body,omitempty"`
}

// NewRemoveAuditorFromOrganizationCreated creates RemoveAuditorFromOrganizationCreated with default headers values
func NewRemoveAuditorFromOrganizationCreated() *RemoveAuditorFromOrganizationCreated {

	return &RemoveAuditorFromOrganizationCreated{}
}

// WithPayload adds the payload to the remove auditor from organization created response
func (o *RemoveAuditorFromOrganizationCreated) WithPayload(payload *models.RemoveAuditorFromOrganizationResponseResource) *RemoveAuditorFromOrganizationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove auditor from organization created response
func (o *RemoveAuditorFromOrganizationCreated) SetPayload(payload *models.RemoveAuditorFromOrganizationResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAuditorFromOrganizationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}