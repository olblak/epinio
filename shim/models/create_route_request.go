// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateRouteRequest create route request
//
// swagger:model createRouteRequest
type CreateRouteRequest struct {

	// The guid of the associated domain
	DomainGUID string `json:"domain_guid,omitempty"`

	// The guid of the route.
	GUID string `json:"guid,omitempty"`

	// The host portion of the route
	Host GenericObject `json:"host,omitempty"`

	// The guid of the associated space
	SpaceGUID string `json:"space_guid,omitempty"`
}

// Validate validates this create route request
func (m *CreateRouteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRouteRequest) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if err := m.Host.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("host")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateRouteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRouteRequest) UnmarshalBinary(b []byte) error {
	var res CreateRouteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}