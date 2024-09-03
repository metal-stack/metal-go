// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SizeReservationListRequest v1 size reservation list request
//
// swagger:model v1.SizeReservationListRequest
type V1SizeReservationListRequest struct {

	// the partition id of this size reservation
	Partitionid string `json:"partitionid,omitempty" yaml:"partitionid,omitempty"`

	// the project id of this size reservation
	Projectid string `json:"projectid,omitempty" yaml:"projectid,omitempty"`

	// the size id of this size reservation
	Sizeid string `json:"sizeid,omitempty" yaml:"sizeid,omitempty"`

	// the tenant of this size reservation
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`
}

// Validate validates this v1 size reservation list request
func (m *V1SizeReservationListRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 size reservation list request based on context it is used
func (m *V1SizeReservationListRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeReservationListRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeReservationListRequest) UnmarshalBinary(b []byte) error {
	var res V1SizeReservationListRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}