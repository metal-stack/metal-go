// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineState v1 machine state
//
// swagger:model v1.MachineState
type V1MachineState struct {

	// a description why this machine is in the given state
	// Required: true
	Description *string `json:"description" yaml:"description"`

	// the user that changed the state
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// the version of metal hammer which put the machine in waiting state
	// Required: true
	MetalHammerVersion *string `json:"metal_hammer_version" yaml:"metal_hammer_version"`

	// the state of this machine. empty means available for all
	// Required: true
	// Enum: ["","LOCKED","RESERVED"]
	Value *string `json:"value" yaml:"value"`
}

// Validate validates this v1 machine state
func (m *V1MachineState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetalHammerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineState) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineState) validateMetalHammerVersion(formats strfmt.Registry) error {

	if err := validate.Required("metal_hammer_version", "body", m.MetalHammerVersion); err != nil {
		return err
	}

	return nil
}

var v1MachineStateTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","LOCKED","RESERVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1MachineStateTypeValuePropEnum = append(v1MachineStateTypeValuePropEnum, v)
	}
}

const (

	// V1MachineStateValueEmpty captures enum value ""
	V1MachineStateValueEmpty string = ""

	// V1MachineStateValueLOCKED captures enum value "LOCKED"
	V1MachineStateValueLOCKED string = "LOCKED"

	// V1MachineStateValueRESERVED captures enum value "RESERVED"
	V1MachineStateValueRESERVED string = "RESERVED"
)

// prop value enum
func (m *V1MachineState) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1MachineStateTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1MachineState) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine state based on context it is used
func (m *V1MachineState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineState) UnmarshalBinary(b []byte) error {
	var res V1MachineState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
