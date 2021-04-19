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

// V1Filesystem v1 filesystem
//
// swagger:model v1.Filesystem
type V1Filesystem struct {

	// device
	// Required: true
	Device *string `json:"Device"`

	// format
	// Required: true
	Format *string `json:"Format"`

	// label
	// Required: true
	Label *string `json:"Label"`

	// mount options
	// Required: true
	MountOptions []string `json:"MountOptions"`

	// options
	// Required: true
	Options []string `json:"Options"`

	// path
	// Required: true
	Path *string `json:"Path"`
}

// Validate validates this v1 filesystem
func (m *V1Filesystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Filesystem) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("Device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *V1Filesystem) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("Format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *V1Filesystem) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("Label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *V1Filesystem) validateMountOptions(formats strfmt.Registry) error {

	if err := validate.Required("MountOptions", "body", m.MountOptions); err != nil {
		return err
	}

	return nil
}

func (m *V1Filesystem) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("Options", "body", m.Options); err != nil {
		return err
	}

	return nil
}

func (m *V1Filesystem) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("Path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 filesystem based on context it is used
func (m *V1Filesystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Filesystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Filesystem) UnmarshalBinary(b []byte) error {
	var res V1Filesystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
