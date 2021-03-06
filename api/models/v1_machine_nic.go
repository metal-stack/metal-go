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

// V1MachineNic v1 machine nic
//
// swagger:model v1.MachineNic
type V1MachineNic struct {

	// the mac address of this network interface
	// Required: true
	Mac *string `json:"mac"`

	// the name of this network interface
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this v1 machine nic
func (m *V1MachineNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineNic) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("mac", "body", m.Mac); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine nic based on context it is used
func (m *V1MachineNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineNic) UnmarshalBinary(b []byte) error {
	var res V1MachineNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
