// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TenantResponse v1 tenant response
//
// swagger:model v1.TenantResponse
type V1TenantResponse struct {

	// default quotas
	DefaultQuotas *V1QuotaSet `json:"default_quotas,omitempty" yaml:"default_quotas,omitempty"`

	// description
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// iam config
	IamConfig *V1IAMConfig `json:"iam_config,omitempty" yaml:"iam_config,omitempty"`

	// meta
	Meta *V1Meta `json:"meta,omitempty" yaml:"meta,omitempty"`

	// name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// quotas
	Quotas *V1QuotaSet `json:"quotas,omitempty" yaml:"quotas,omitempty"`
}

// Validate validates this v1 tenant response
func (m *V1TenantResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultQuotas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIamConfig(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1TenantResponse) validateDefaultQuotas(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultQuotas) { // not required
		return nil
	}

	if m.DefaultQuotas != nil {
		if err := m.DefaultQuotas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_quotas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_quotas")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantResponse) validateIamConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.IamConfig) { // not required
		return nil
	}

	if m.IamConfig != nil {
		if err := m.IamConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iam_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iam_config")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantResponse) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantResponse) validateQuotas(formats strfmt.Registry) error {
	if swag.IsZero(m.Quotas) { // not required
		return nil
	}

	if m.Quotas != nil {
		if err := m.Quotas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quotas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quotas")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 tenant response based on the context it is used
func (m *V1TenantResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultQuotas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIamConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuotas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantResponse) contextValidateDefaultQuotas(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultQuotas != nil {

		if swag.IsZero(m.DefaultQuotas) { // not required
			return nil
		}

		if err := m.DefaultQuotas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_quotas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_quotas")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantResponse) contextValidateIamConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.IamConfig != nil {

		if swag.IsZero(m.IamConfig) { // not required
			return nil
		}

		if err := m.IamConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iam_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iam_config")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantResponse) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {

		if swag.IsZero(m.Meta) { // not required
			return nil
		}

		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantResponse) contextValidateQuotas(ctx context.Context, formats strfmt.Registry) error {

	if m.Quotas != nil {

		if swag.IsZero(m.Quotas) { // not required
			return nil
		}

		if err := m.Quotas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quotas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quotas")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantResponse) UnmarshalBinary(b []byte) error {
	var res V1TenantResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
