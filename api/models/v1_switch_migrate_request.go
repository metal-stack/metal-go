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

// V1SwitchMigrateRequest v1 switch migrate request
//
// swagger:model v1.SwitchMigrateRequest
type V1SwitchMigrateRequest struct {

	// the id of the new switch to migrate to
	// Required: true
	NewSwitchID *string `json:"new_switch_id" yaml:"new_switch_id"`

	// the id of the switch that should be migrated away from
	// Required: true
	OldSwitchID *string `json:"old_switch_id" yaml:"old_switch_id"`
}

// Validate validates this v1 switch migrate request
func (m *V1SwitchMigrateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewSwitchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldSwitchID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchMigrateRequest) validateNewSwitchID(formats strfmt.Registry) error {

	if err := validate.Required("new_switch_id", "body", m.NewSwitchID); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchMigrateRequest) validateOldSwitchID(formats strfmt.Registry) error {

	if err := validate.Required("old_switch_id", "body", m.OldSwitchID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 switch migrate request based on context it is used
func (m *V1SwitchMigrateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SwitchMigrateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SwitchMigrateRequest) UnmarshalBinary(b []byte) error {
	var res V1SwitchMigrateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
