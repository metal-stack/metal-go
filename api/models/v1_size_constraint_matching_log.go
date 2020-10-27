// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SizeConstraintMatchingLog v1 size constraint matching log
//
// swagger:model v1.SizeConstraintMatchingLog
type V1SizeConstraintMatchingLog struct {

	// the size constraint to which this log relates to
	// Required: true
	Constraint *V1SizeConstraint `json:"constraint"`

	// a string represention of the matching condition
	// Required: true
	Log *string `json:"log"`

	// indicates whether the constraint matched or not
	// Required: true
	Match *bool `json:"match"`
}

// Validate validates this v1 size constraint matching log
func (m *V1SizeConstraintMatchingLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeConstraintMatchingLog) validateConstraint(formats strfmt.Registry) error {

	if err := validate.Required("constraint", "body", m.Constraint); err != nil {
		return err
	}

	if m.Constraint != nil {
		if err := m.Constraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraint")
			}
			return err
		}
	}

	return nil
}

func (m *V1SizeConstraintMatchingLog) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeConstraintMatchingLog) validateMatch(formats strfmt.Registry) error {

	if err := validate.Required("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeConstraintMatchingLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeConstraintMatchingLog) UnmarshalBinary(b []byte) error {
	var res V1SizeConstraintMatchingLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
