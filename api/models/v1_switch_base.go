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

// V1SwitchBase A switch that can register at the api.
//
// swagger:model v1.SwitchBase
type V1SwitchBase struct {

	// command to access the console of the switch
	ConsoleCommand string `json:"console_command,omitempty" yaml:"console_command,omitempty"`

	// the ip address of the management interface of the switch
	ManagementIP string `json:"management_ip,omitempty" yaml:"management_ip,omitempty"`

	// the user to connect to the switch
	ManagementUser string `json:"management_user,omitempty" yaml:"management_user,omitempty"`

	// the mode the switch currently has
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// the operating system the switch currently has
	Os *V1SwitchOS `json:"os,omitempty" yaml:"os,omitempty"`

	// the id of the rack in which this switch is located
	// Required: true
	RackID *string `json:"rack_id" yaml:"rack_id"`
}

// Validate validates this v1 switch base
func (m *V1SwitchBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRackID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchBase) validateOs(formats strfmt.Registry) error {
	if swag.IsZero(m.Os) { // not required
		return nil
	}

	if m.Os != nil {
		if err := m.Os.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os")
			}
			return err
		}
	}

	return nil
}

func (m *V1SwitchBase) validateRackID(formats strfmt.Registry) error {

	if err := validate.Required("rack_id", "body", m.RackID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 switch base based on the context it is used
func (m *V1SwitchBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchBase) contextValidateOs(ctx context.Context, formats strfmt.Registry) error {

	if m.Os != nil {

		if swag.IsZero(m.Os) { // not required
			return nil
		}

		if err := m.Os.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SwitchBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SwitchBase) UnmarshalBinary(b []byte) error {
	var res V1SwitchBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
