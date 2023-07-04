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

// V1SizeImageConstraintCreateRequest v1 size image constraint create request
//
// swagger:model v1.SizeImageConstraintCreateRequest
type V1SizeImageConstraintCreateRequest struct {

	// a list of constraints that for this size
	Constraints *V1SizeImageConstraintBase `json:"constraints,omitempty" yaml:"constraints,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this v1 size image constraint create request
func (m *V1SizeImageConstraintCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeImageConstraintCreateRequest) validateConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	if m.Constraints != nil {
		if err := m.Constraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constraints")
			}
			return err
		}
	}

	return nil
}

func (m *V1SizeImageConstraintCreateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 size image constraint create request based on the context it is used
func (m *V1SizeImageConstraintCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeImageConstraintCreateRequest) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.Constraints != nil {

		if swag.IsZero(m.Constraints) { // not required
			return nil
		}

		if err := m.Constraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constraints")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeImageConstraintCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeImageConstraintCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1SizeImageConstraintCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
