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

// V1Raid v1 raid
//
// swagger:model v1.Raid
type V1Raid struct {

	// the name of the resulting array device
	// Required: true
	Arrayname *string `json:"arrayname" yaml:"arrayname"`

	// the options to use to create the raid array
	Createoptions []string `json:"createoptions" yaml:"createoptions"`

	// list of devices to form the raid array from
	Devices []string `json:"devices" yaml:"devices"`

	// raid level to create, should be 0 or 1
	// Required: true
	Level *string `json:"level" yaml:"level"`

	// number of spares for the raid array
	// Required: true
	Spares *int32 `json:"spares" yaml:"spares"`
}

// Validate validates this v1 raid
func (m *V1Raid) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrayname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpares(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Raid) validateArrayname(formats strfmt.Registry) error {

	if err := validate.Required("arrayname", "body", m.Arrayname); err != nil {
		return err
	}

	return nil
}

func (m *V1Raid) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *V1Raid) validateSpares(formats strfmt.Registry) error {

	if err := validate.Required("spares", "body", m.Spares); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 raid based on context it is used
func (m *V1Raid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Raid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Raid) UnmarshalBinary(b []byte) error {
	var res V1Raid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
