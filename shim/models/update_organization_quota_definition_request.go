// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrganizationQuotaDefinitionRequest update organization quota definition request
//
// swagger:model updateOrganizationQuotaDefinitionRequest
type UpdateOrganizationQuotaDefinitionRequest struct {

	// The maximum amount of memory in megabyte an application instance can have. (-1 represents an unlimited amount)
	InstanceMemoryLimit string `json:"instance_memory_limit,omitempty"`

	// How much memory in megabyte an organization can have.
	MemoryLimit string `json:"memory_limit,omitempty"`

	// The name for the Organization Quota Definition.
	Name string `json:"name,omitempty"`

	// If an organization can have non basic services
	NonBasicServicesAllowed bool `json:"non_basic_services_allowed,omitempty"`

	// How many routes an organization can have.
	TotalRoutes string `json:"total_routes,omitempty"`

	// How many services an organization can have.
	TotalServices string `json:"total_services,omitempty"`

	// If an organization can have a trial db.
	TrialDbAllowed GenericObject `json:"trial_db_allowed,omitempty"`
}

// Validate validates this update organization quota definition request
func (m *UpdateOrganizationQuotaDefinitionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrialDbAllowed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOrganizationQuotaDefinitionRequest) validateTrialDbAllowed(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialDbAllowed) { // not required
		return nil
	}

	if err := m.TrialDbAllowed.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trial_db_allowed")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrganizationQuotaDefinitionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrganizationQuotaDefinitionRequest) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationQuotaDefinitionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}