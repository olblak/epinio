// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllMatchingResourcesRequest list all matching resources request
//
// swagger:model listAllMatchingResourcesRequest
type ListAllMatchingResourcesRequest struct {

	// The sha1
	Sha1 string `json:"sha1,omitempty"`

	// The size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this list all matching resources request
func (m *ListAllMatchingResourcesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllMatchingResourcesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllMatchingResourcesRequest) UnmarshalBinary(b []byte) error {
	var res ListAllMatchingResourcesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}