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

// V1SizeImageConstraintBase v1 size image constraint base
//
// swagger:model v1.SizeImageConstraintBase
type V1SizeImageConstraintBase struct {

	// a list of images for this constraints apply
	// Required: true
	Images map[string]string `json:"images" yaml:"images"`
}

// Validate validates this v1 size image constraint base
func (m *V1SizeImageConstraintBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeImageConstraintBase) validateImages(formats strfmt.Registry) error {

	if err := validate.Required("images", "body", m.Images); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 size image constraint base based on context it is used
func (m *V1SizeImageConstraintBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeImageConstraintBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeImageConstraintBase) UnmarshalBinary(b []byte) error {
	var res V1SizeImageConstraintBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
