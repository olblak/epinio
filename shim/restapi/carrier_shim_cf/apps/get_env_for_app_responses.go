// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// GetEnvForAppOKCode is the HTTP code returned for type GetEnvForAppOK
const GetEnvForAppOKCode int = 200

/*GetEnvForAppOK successful response

swagger:response getEnvForAppOK
*/
type GetEnvForAppOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetEnvForAppResponseResource `json:"body,omitempty"`
}

// NewGetEnvForAppOK creates GetEnvForAppOK with default headers values
func NewGetEnvForAppOK() *GetEnvForAppOK {

	return &GetEnvForAppOK{}
}

// WithPayload adds the payload to the get env for app o k response
func (o *GetEnvForAppOK) WithPayload(payload *models.GetEnvForAppResponseResource) *GetEnvForAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get env for app o k response
func (o *GetEnvForAppOK) SetPayload(payload *models.GetEnvForAppResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEnvForAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}