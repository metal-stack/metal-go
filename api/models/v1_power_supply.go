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

// V1PowerSupply v1 power supply
//
// swagger:model v1.PowerSupply
type V1PowerSupply struct {

	// status
	// Required: true
	Status *V1PowerSupplyStatus `json:"status" yaml:"status"`
}

// Validate validates this v1 power supply
func (m *V1PowerSupply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PowerSupply) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 power supply based on the context it is used
func (m *V1PowerSupply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PowerSupply) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PowerSupply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PowerSupply) UnmarshalBinary(b []byte) error {
	var res V1PowerSupply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}