// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllServicesForOrganizationResponse list all services for organization response
//
// swagger:model listAllServicesForOrganizationResponse
type ListAllServicesForOrganizationResponse struct {

	// The active
	Active bool `json:"active,omitempty"`

	// The bindable
	Bindable bool `json:"bindable,omitempty"`

	// The description
	Description string `json:"description,omitempty"`

	// The documentation Url
	DocumentationURL GenericObject `json:"documentation_url,omitempty"`

	// The extra
	Extra GenericObject `json:"extra,omitempty"`

	// The info Url
	InfoURL GenericObject `json:"info_url,omitempty"`

	// The label
	Label string `json:"label,omitempty"`

	// The long Description
	LongDescription GenericObject `json:"long_description,omitempty"`

	// The plan Updateable
	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	// The provider
	Provider string `json:"provider,omitempty"`

	// The requires
	Requires []GenericObject `json:"requires"`

	// The service Broker Guid
	ServiceBrokerGUID string `json:"service_broker_guid,omitempty"`

	// The service Plans Url
	ServicePlansURL string `json:"service_plans_url,omitempty"`

	// The tags
	Tags []GenericObject `json:"tags"`

	// The unique Id
	UniqueID string `json:"unique_id,omitempty"`

	// The url
	URL string `json:"url,omitempty"`

	// The version
	Version string `json:"version,omitempty"`
}

// Validate validates this list all services for organization response
func (m *ListAllServicesForOrganizationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtra(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfoURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequires(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAllServicesForOrganizationResponse) validateDocumentationURL(formats strfmt.Registry) error {

	if swag.IsZero(m.DocumentationURL) { // not required
		return nil
	}

	if err := m.DocumentationURL.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("documentation_url")
		}
		return err
	}

	return nil
}

func (m *ListAllServicesForOrganizationResponse) validateExtra(formats strfmt.Registry) error {

	if swag.IsZero(m.Extra) { // not required
		return nil
	}

	if err := m.Extra.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("extra")
		}
		return err
	}

	return nil
}

func (m *ListAllServicesForOrganizationResponse) validateInfoURL(formats strfmt.Registry) error {

	if swag.IsZero(m.InfoURL) { // not required
		return nil
	}

	if err := m.InfoURL.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("info_url")
		}
		return err
	}

	return nil
}

func (m *ListAllServicesForOrganizationResponse) validateLongDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.LongDescription) { // not required
		return nil
	}

	if err := m.LongDescription.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("long_description")
		}
		return err
	}

	return nil
}

func (m *ListAllServicesForOrganizationResponse) validateRequires(formats strfmt.Registry) error {

	if swag.IsZero(m.Requires) { // not required
		return nil
	}

	for i := 0; i < len(m.Requires); i++ {

		if err := m.Requires[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requires" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ListAllServicesForOrganizationResponse) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := m.Tags[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListAllServicesForOrganizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllServicesForOrganizationResponse) UnmarshalBinary(b []byte) error {
	var res ListAllServicesForOrganizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}