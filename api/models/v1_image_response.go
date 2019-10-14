// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ImageResponse v1 image response
// swagger:model v1.ImageResponse
type V1ImageResponse struct {

	// the last changed timestamp of this entity
	// Required: true
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed"`

	// the creation time of this entity
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// features of this image
	Features []string `json:"features"`

	// the unique ID of this entity
	// Required: true
	// Unique: true
	ID *string `json:"id"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// the url of this image
	URL string `json:"url,omitempty"`

	// date to which it is allowed to allocate machines from
	// Required: true
	// Format: date-time
	Validto *strfmt.DateTime `json:"validto"`
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

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidto(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageResponse) validateChanged(formats strfmt.Registry) error {

	if err := validate.Required("changed", "body", strfmt.DateTime(m.Changed)); err != nil {
		return err
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
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

func (m *V1ImageResponse) validateValidto(formats strfmt.Registry) error {

	if err := validate.Required("validto", "body", m.Validto); err != nil {
		return err
	}

	if err := validate.FormatOf("validto", "body", "date-time", m.Validto.String(), formats); err != nil {
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
