// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveSpaceFromSecurityGroupCreatedCode is the HTTP code returned for type RemoveSpaceFromSecurityGroupCreated
const RemoveSpaceFromSecurityGroupCreatedCode int = 201

/*RemoveSpaceFromSecurityGroupCreated successful response

swagger:response removeSpaceFromSecurityGroupCreated
*/
type RemoveSpaceFromSecurityGroupCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveSpaceFromSecurityGroupResponseResource `json:"body,omitempty"`
}

// NewRemoveSpaceFromSecurityGroupCreated creates RemoveSpaceFromSecurityGroupCreated with default headers values
func NewRemoveSpaceFromSecurityGroupCreated() *RemoveSpaceFromSecurityGroupCreated {

	return &RemoveSpaceFromSecurityGroupCreated{}
}

// WithPayload adds the payload to the remove space from security group created response
func (o *RemoveSpaceFromSecurityGroupCreated) WithPayload(payload *models.RemoveSpaceFromSecurityGroupResponseResource) *RemoveSpaceFromSecurityGroupCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove space from security group created response
func (o *RemoveSpaceFromSecurityGroupCreated) SetPayload(payload *models.RemoveSpaceFromSecurityGroupResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveSpaceFromSecurityGroupCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}