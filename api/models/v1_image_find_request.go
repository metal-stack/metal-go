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

// V1ImageFindRequest v1 image find request
//
// swagger:model v1.ImageFindRequest
type V1ImageFindRequest struct {

	// classification
	// Enum: ["deprecated","preview","supported"]
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// features
	Features []string `json:"features" yaml:"features"`

	// id
	ID string `json:"id,omitempty" yaml:"id,omitempty"`

	// name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// os
	Os string `json:"os,omitempty" yaml:"os,omitempty"`

	// version
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

// Validate validates this v1 image find request
func (m *V1ImageFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1ImageFindRequestTypeClassificationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deprecated","preview","supported"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ImageFindRequestTypeClassificationPropEnum = append(v1ImageFindRequestTypeClassificationPropEnum, v)
	}
}

const (

	// V1ImageFindRequestClassificationDeprecated captures enum value "deprecated"
	V1ImageFindRequestClassificationDeprecated string = "deprecated"

	// V1ImageFindRequestClassificationPreview captures enum value "preview"
	V1ImageFindRequestClassificationPreview string = "preview"

	// V1ImageFindRequestClassificationSupported captures enum value "supported"
	V1ImageFindRequestClassificationSupported string = "supported"
)

// prop value enum
func (m *V1ImageFindRequest) validateClassificationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ImageFindRequestTypeClassificationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ImageFindRequest) validateClassification(formats strfmt.Registry) error {
	if swag.IsZero(m.Classification) { // not required
		return nil
	}

	// value enum
	if err := m.validateClassificationEnum("classification", "body", m.Classification); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 image find request based on context it is used
func (m *V1ImageFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageFindRequest) UnmarshalBinary(b []byte) error {
	var res V1ImageFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
