// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1ProjectResponse v1 project response
// swagger:model v1.ProjectResponse
type V1ProjectResponse struct {

	// description
	Description string `json:"description,omitempty"`

	// meta
	Meta *V1Meta `json:"meta,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// quotas
	Quotas *V1QuotaSet `json:"quotas,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this v1 project response
func (m *V1ProjectResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectResponse) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectResponse) validateQuotas(formats strfmt.Registry) error {

	if swag.IsZero(m.Quotas) { // not required
		return nil
	}

	if m.Quotas != nil {
		if err := m.Quotas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quotas")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectResponse) UnmarshalBinary(b []byte) error {
	var res V1ProjectResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
