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

// V1Common v1 common
//
// swagger:model v1.Common
type V1Common struct {

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this v1 common
func (m *V1Common) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Common) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 common based on context it is used
func (m *V1Common) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Common) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Common) UnmarshalBinary(b []byte) error {
	var res V1Common
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
