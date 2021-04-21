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

// V1DiskPartition v1 disk partition
//
// swagger:model v1.DiskPartition
type V1DiskPartition struct {

	// g p t type
	// Required: true
	GPTType *string `json:"GPTType"`

	// label
	// Required: true
	Label *string `json:"Label"`

	// number
	// Required: true
	Number *int64 `json:"Number"`

	// size
	// Required: true
	Size *int64 `json:"Size"`
}

// Validate validates this v1 disk partition
func (m *V1DiskPartition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGPTType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DiskPartition) validateGPTType(formats strfmt.Registry) error {

	if err := validate.Required("GPTType", "body", m.GPTType); err != nil {
		return err
	}

	return nil
}

func (m *V1DiskPartition) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("Label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *V1DiskPartition) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("Number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *V1DiskPartition) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 disk partition based on context it is used
func (m *V1DiskPartition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DiskPartition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DiskPartition) UnmarshalBinary(b []byte) error {
	var res V1DiskPartition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
