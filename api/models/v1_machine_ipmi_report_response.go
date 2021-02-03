// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineIpmiReportResponse v1 machine ipmi report response
//
// swagger:model v1.MachineIpmiReportResponse
type V1MachineIpmiReportResponse struct {

	// the machine uuids that triggered a creation of a machine entity
	// Required: true
	Created []string `json:"created"`

	// the machine uuids that triggered an update of ipmi data
	// Required: true
	Updated []string `json:"updated"`
}

// Validate validates this v1 machine ipmi report response
func (m *V1MachineIpmiReportResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineIpmiReportResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIpmiReportResponse) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine ipmi report response based on context it is used
func (m *V1MachineIpmiReportResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineIpmiReportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineIpmiReportResponse) UnmarshalBinary(b []byte) error {
	var res V1MachineIpmiReportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
