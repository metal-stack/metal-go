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

// V1ImageResponse v1 image response
//
// swagger:model v1.ImageResponse
type V1ImageResponse struct {

	// the last changed timestamp of this entity
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed,omitempty" yaml:"changed,omitempty"`

	// classification of this image
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// the creation time of this entity
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty" yaml:"created,omitempty"`

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

// Validate validates this v1 image response
func (m *V1ImageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1ImageResponse) validateChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.Changed) { // not required
		return nil
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageResponse) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageResponse) validateExpirationDate(formats strfmt.Registry) error {

	if err := validate.Required("expirationDate", "body", m.ExpirationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 image response based on the context it is used
func (m *V1ImageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChanged(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageResponse) contextValidateChanged(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "changed", "body", strfmt.DateTime(m.Changed)); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageResponse) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageResponse) UnmarshalBinary(b []byte) error {
	var res V1ImageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
