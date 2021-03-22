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

// V1FirmwaresResponse v1 firmwares response
//
// swagger:model v1.FirmwaresResponse
type V1FirmwaresResponse struct {

	// list of firmwares per board per vendor per kind
	// Required: true
	Revisions map[string]V1FirmwaresResponseRevisions `json:"revisions"`
}

// Validate validates this v1 firmwares response
func (m *V1FirmwaresResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevisions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FirmwaresResponse) validateRevisions(formats strfmt.Registry) error {

	if err := validate.Required("revisions", "body", m.Revisions); err != nil {
		return err
	}

	for k := range m.Revisions {

		if err := validate.Required("revisions"+"."+k, "body", m.Revisions[k]); err != nil {
			return err
		}

		if err := validate.Required("revisions"+"."+k, "body", m.Revisions); err != nil {
			return err
		}

		if val, ok := m.Revisions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 firmwares response based on the context it is used
func (m *V1FirmwaresResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRevisions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FirmwaresResponse) contextValidateRevisions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("revisions", "body", m.Revisions); err != nil {
		return err
	}

	for k := range m.Revisions {

		if err := validate.Required("revisions"+"."+k, "body", m.Revisions); err != nil {
			return err
		}

		if val, ok := m.Revisions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1FirmwaresResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FirmwaresResponse) UnmarshalBinary(b []byte) error {
	var res V1FirmwaresResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
