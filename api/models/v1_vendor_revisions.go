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

// V1VendorRevisions v1 vendor revisions
//
// swagger:model v1.VendorRevisions
type V1VendorRevisions struct {

	// vendor revisions
	// Required: true
	VendorRevisions map[string]V1BoardRevisions `json:"VendorRevisions" yaml:"VendorRevisions"`
}

// Validate validates this v1 vendor revisions
func (m *V1VendorRevisions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVendorRevisions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VendorRevisions) validateVendorRevisions(formats strfmt.Registry) error {

	if err := validate.Required("VendorRevisions", "body", m.VendorRevisions); err != nil {
		return err
	}

	for k := range m.VendorRevisions {

		if err := validate.Required("VendorRevisions"+"."+k, "body", m.VendorRevisions[k]); err != nil {
			return err
		}
		if val, ok := m.VendorRevisions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("VendorRevisions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("VendorRevisions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 vendor revisions based on the context it is used
func (m *V1VendorRevisions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVendorRevisions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VendorRevisions) contextValidateVendorRevisions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("VendorRevisions", "body", m.VendorRevisions); err != nil {
		return err
	}

	for k := range m.VendorRevisions {

		if val, ok := m.VendorRevisions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VendorRevisions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VendorRevisions) UnmarshalBinary(b []byte) error {
	var res V1VendorRevisions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
