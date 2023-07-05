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

// V1SwitchNic v1 switch nic
//
// swagger:model v1.SwitchNic
type V1SwitchNic struct {

	// configures the bgp filter applied at the switch port
	Filter *V1BGPFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// the identifier of this network interface
	// Required: true
	Identifier *string `json:"identifier" yaml:"identifier"`

	// the mac address of this network interface
	// Required: true
	Mac *string `json:"mac" yaml:"mac"`

	// the name of this network interface
	// Required: true
	Name *string `json:"name" yaml:"name"`

	// the vrf this network interface is part of
	Vrf string `json:"vrf,omitempty" yaml:"vrf,omitempty"`
}

// Validate validates this v1 switch nic
func (m *V1SwitchNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1SwitchNic) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *V1SwitchNic) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchNic) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("mac", "body", m.Mac); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchNic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 switch nic based on the context it is used
func (m *V1SwitchNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchNic) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.Filter != nil {

		if swag.IsZero(m.Filter) { // not required
			return nil
		}

		if err := m.Filter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SwitchNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SwitchNic) UnmarshalBinary(b []byte) error {
	var res V1SwitchNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
