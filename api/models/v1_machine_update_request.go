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

// V1MachineUpdateRequest v1 machine update request
//
// swagger:model v1.MachineUpdateRequest
type V1MachineUpdateRequest struct {

	// a description for this machine
	// Required: true
	Description *string `json:"description" yaml:"description"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// the public ssh keys to access the machine with
	SSHPubKeys []string `json:"ssh_pub_keys" yaml:"ssh_pub_keys"`

	// tags for this machine.
	Tags []string `json:"tags" yaml:"tags"`
}

// Validate validates this v1 machine update request
func (m *V1MachineUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
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

func (m *V1MachineUpdateRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine update request based on context it is used
func (m *V1MachineUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineUpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
