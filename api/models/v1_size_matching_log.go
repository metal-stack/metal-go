// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SizeMatchingLog v1 size matching log
//
// swagger:model v1.SizeMatchingLog
type V1SizeMatchingLog struct {

	// constraints
	// Required: true
	Constraints []*V1SizeConstraintMatchingLog `json:"constraints" yaml:"constraints"`

	// log
	// Required: true
	Log *string `json:"log" yaml:"log"`

	// match
	// Required: true
	Match *bool `json:"match" yaml:"match"`

	// name
	// Required: true
	Name *string `json:"name" yaml:"name"`
}

// Validate validates this v1 size matching log
func (m *V1SizeMatchingLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
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

func (m *V1SizeMatchingLog) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("constraints", "body", m.Constraints); err != nil {
		return err
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SizeMatchingLog) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeMatchingLog) validateMatch(formats strfmt.Registry) error {

	if err := validate.Required("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeMatchingLog) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 size matching log based on the context it is used
func (m *V1SizeMatchingLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeMatchingLog) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Constraints); i++ {

		if m.Constraints[i] != nil {

			if swag.IsZero(m.Constraints[i]) { // not required
				return nil
			}

			if err := m.Constraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeMatchingLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeMatchingLog) UnmarshalBinary(b []byte) error {
	var res V1SizeMatchingLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
