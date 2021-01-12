// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveJobThatWasSuccessfulOKCode is the HTTP code returned for type RetrieveJobThatWasSuccessfulOK
const RetrieveJobThatWasSuccessfulOKCode int = 200

/*RetrieveJobThatWasSuccessfulOK successful response

swagger:response retrieveJobThatWasSuccessfulOK
*/
type RetrieveJobThatWasSuccessfulOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveJobThatWasSuccessfulResponseResource `json:"body,omitempty"`
}

// NewRetrieveJobThatWasSuccessfulOK creates RetrieveJobThatWasSuccessfulOK with default headers values
func NewRetrieveJobThatWasSuccessfulOK() *RetrieveJobThatWasSuccessfulOK {

	return &RetrieveJobThatWasSuccessfulOK{}
}

// WithPayload adds the payload to the retrieve job that was successful o k response
func (o *RetrieveJobThatWasSuccessfulOK) WithPayload(payload *models.RetrieveJobThatWasSuccessfulResponseResource) *RetrieveJobThatWasSuccessfulOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve job that was successful o k response
func (o *RetrieveJobThatWasSuccessfulOK) SetPayload(payload *models.RetrieveJobThatWasSuccessfulResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveJobThatWasSuccessfulOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}