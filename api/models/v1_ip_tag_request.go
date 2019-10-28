// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IPTagRequest v1 IP tag request
// swagger:model v1.IPTagRequest
type V1IPTagRequest struct {

	// clusterid
	Clusterid string `json:"clusterid,omitempty"`

	// the address (ipv4 or ipv6) of this ip
	// Required: true
	// Unique: true
	Ipaddress *string `json:"ipaddress"`

	// machineid
	Machineid string `json:"machineid,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this v1 IP tag request
func (m *V1IPTagRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIpaddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IPTagRequest) validateIpaddress(formats strfmt.Registry) error {

	if err := validate.Required("ipaddress", "body", m.Ipaddress); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IPTagRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPTagRequest) UnmarshalBinary(b []byte) error {
	var res V1IPTagRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}