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

// V1ImageUpdateRequest v1 image update request
//
// swagger:model v1.ImageUpdateRequest
type V1ImageUpdateRequest struct {

	// classification of this image
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// expirationDate of this image
	// Required: true
	// Format: date-time
	ExpirationDate *strfmt.DateTime `json:"expirationDate" yaml:"expirationDate"`

	// features of this image
	Features []string `json:"features" yaml:"features"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// the url of this image
	URL string `json:"url,omitempty" yaml:"url,omitempty"`

	// machines where this image is in use
	Usedby []string `json:"usedby" yaml:"usedby"`
}

// Validate validates this v1 image update request
func (m *V1ImageUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationDate(formats); err != nil {
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

func (m *V1ImageUpdateRequest) validateExpirationDate(formats strfmt.Registry) error {

	if err := validate.Required("expirationDate", "body", m.ExpirationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 image update request based on context it is used
func (m *V1ImageUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageUpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1ImageUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
