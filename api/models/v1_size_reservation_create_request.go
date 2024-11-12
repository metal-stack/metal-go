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

// V1SizeReservationCreateRequest v1 size reservation create request
//
// swagger:model v1.SizeReservationCreateRequest
type V1SizeReservationCreateRequest struct {

	// the amount of reservations of this size reservation
	// Required: true
	Amount *int32 `json:"amount" yaml:"amount"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// free labels associated with this size reservation.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// the partition id of this size reservation
	// Required: true
	Partitionids []string `json:"partitionids" yaml:"partitionids"`

	// the project id of this size reservation
	// Required: true
	Projectid *string `json:"projectid" yaml:"projectid"`

	// the size id of this size reservation
	// Required: true
	Sizeid *string `json:"sizeid" yaml:"sizeid"`
}

// Validate validates this v1 size reservation create request
func (m *V1SizeReservationCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeReservationCreateRequest) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeReservationCreateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeReservationCreateRequest) validatePartitionids(formats strfmt.Registry) error {

	if err := validate.Required("partitionids", "body", m.Partitionids); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeReservationCreateRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeReservationCreateRequest) validateSizeid(formats strfmt.Registry) error {

	if err := validate.Required("sizeid", "body", m.Sizeid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 size reservation create request based on context it is used
func (m *V1SizeReservationCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeReservationCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeReservationCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1SizeReservationCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}