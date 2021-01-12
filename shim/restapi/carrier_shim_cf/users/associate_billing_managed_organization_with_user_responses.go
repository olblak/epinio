// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// AssociateBillingManagedOrganizationWithUserCreatedCode is the HTTP code returned for type AssociateBillingManagedOrganizationWithUserCreated
const AssociateBillingManagedOrganizationWithUserCreatedCode int = 201

/*AssociateBillingManagedOrganizationWithUserCreated successful response

swagger:response associateBillingManagedOrganizationWithUserCreated
*/
type AssociateBillingManagedOrganizationWithUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AssociateBillingManagedOrganizationWithUserResponseResource `json:"body,omitempty"`
}

// NewAssociateBillingManagedOrganizationWithUserCreated creates AssociateBillingManagedOrganizationWithUserCreated with default headers values
func NewAssociateBillingManagedOrganizationWithUserCreated() *AssociateBillingManagedOrganizationWithUserCreated {

	return &AssociateBillingManagedOrganizationWithUserCreated{}
}

// WithPayload adds the payload to the associate billing managed organization with user created response
func (o *AssociateBillingManagedOrganizationWithUserCreated) WithPayload(payload *models.AssociateBillingManagedOrganizationWithUserResponseResource) *AssociateBillingManagedOrganizationWithUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate billing managed organization with user created response
func (o *AssociateBillingManagedOrganizationWithUserCreated) SetPayload(payload *models.AssociateBillingManagedOrganizationWithUserResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateBillingManagedOrganizationWithUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}