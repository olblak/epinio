// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveOrganizationOKCode is the HTTP code returned for type RetrieveOrganizationOK
const RetrieveOrganizationOKCode int = 200

/*RetrieveOrganizationOK successful response

swagger:response retrieveOrganizationOK
*/
type RetrieveOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveOrganizationResponseResource `json:"body,omitempty"`
}

// NewRetrieveOrganizationOK creates RetrieveOrganizationOK with default headers values
func NewRetrieveOrganizationOK() *RetrieveOrganizationOK {

	return &RetrieveOrganizationOK{}
}

// WithPayload adds the payload to the retrieve organization o k response
func (o *RetrieveOrganizationOK) WithPayload(payload *models.RetrieveOrganizationResponseResource) *RetrieveOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve organization o k response
func (o *RetrieveOrganizationOK) SetPayload(payload *models.RetrieveOrganizationResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}