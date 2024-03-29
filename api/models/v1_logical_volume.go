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

// V1LogicalVolume v1 logical volume
//
// swagger:model v1.LogicalVolume
type V1LogicalVolume struct {

	// the type of this logical volume can be either linear|striped|raid1
	// Required: true
	Lvmtype *string `json:"lvmtype" yaml:"lvmtype"`

	// the name of the logical volume
	// Required: true
	Name *string `json:"name" yaml:"name"`

	// size in mebibytes (MiB) of this volume
	// Required: true
	Size *int64 `json:"size" yaml:"size"`

	// the name of the volume group where to create the logical volume onto
	// Required: true
	Volumegroup *string `json:"volumegroup" yaml:"volumegroup"`
}

// Validate validates this v1 logical volume
func (m *V1LogicalVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLvmtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumegroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LogicalVolume) validateLvmtype(formats strfmt.Registry) error {

	if err := validate.Required("lvmtype", "body", m.Lvmtype); err != nil {
		return err
	}

	return nil
}

func (m *V1LogicalVolume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1LogicalVolume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *V1LogicalVolume) validateVolumegroup(formats strfmt.Registry) error {

	if err := validate.Required("volumegroup", "body", m.Volumegroup); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 logical volume based on context it is used
func (m *V1LogicalVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1LogicalVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LogicalVolume) UnmarshalBinary(b []byte) error {
	var res V1LogicalVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
