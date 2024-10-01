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

// V1PowerSupplyStatus v1 power supply status
//
// swagger:model v1.PowerSupplyStatus
type V1PowerSupplyStatus struct {

	// health
	// Required: true
	Health *string `json:"health" yaml:"health"`

	// state
	// Required: true
	State *string `json:"state" yaml:"state"`
}

// Validate validates this v1 power supply status
func (m *V1PowerSupplyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PowerSupplyStatus) validateHealth(formats strfmt.Registry) error {

	if err := validate.Required("health", "body", m.Health); err != nil {
		return err
	}

	return nil
}

func (m *V1PowerSupplyStatus) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 power supply status based on context it is used
func (m *V1PowerSupplyStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PowerSupplyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PowerSupplyStatus) UnmarshalBinary(b []byte) error {
	var res V1PowerSupplyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
