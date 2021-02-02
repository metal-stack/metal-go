// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IPResponse v1 IP response
//
// swagger:model v1.IPResponse
type V1IPResponse struct {

	// a unique identifier for this ip address allocation, can be used to distinguish between ip address allocation over time.
	// Required: true
	Allocationuuid *string `json:"allocationuuid"`

	// the last changed timestamp of this entity
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed,omitempty"`

	// the creation time of this entity
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the address (ipv4 or ipv6) of this ip
	// Required: true
	// Unique: true
	Ipaddress *string `json:"ipaddress"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// the network this ip allocate request address belongs to
	// Required: true
	Networkid *string `json:"networkid"`

	// the project this ip address belongs to
	// Required: true
	Projectid *string `json:"projectid"`

	// free tags that you associate with this ip.
	Tags []string `json:"tags"`

	// the ip type, ephemeral leads to automatic cleanup of the ip address, static will enable re-use of the ip at a later point in time
	// Required: true
	// Enum: [ephemeral static]
	Type *string `json:"type"`
}

// Validate validates this v1 IP response
func (m *V1IPResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationuuid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpaddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IPResponse) validateAllocationuuid(formats strfmt.Registry) error {

	if err := validate.Required("allocationuuid", "body", m.Allocationuuid); err != nil {
		return err
	}

	return nil
}

func (m *V1IPResponse) validateChanged(formats strfmt.Registry) error {

	if swag.IsZero(m.Changed) { // not required
		return nil
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IPResponse) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IPResponse) validateIpaddress(formats strfmt.Registry) error {

	if err := validate.Required("ipaddress", "body", m.Ipaddress); err != nil {
		return err
	}

	return nil
}

func (m *V1IPResponse) validateNetworkid(formats strfmt.Registry) error {

	if err := validate.Required("networkid", "body", m.Networkid); err != nil {
		return err
	}

	return nil
}

func (m *V1IPResponse) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

var v1IpResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ephemeral","static"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1IpResponseTypeTypePropEnum = append(v1IpResponseTypeTypePropEnum, v)
	}
}

const (

	// V1IPResponseTypeEphemeral captures enum value "ephemeral"
	V1IPResponseTypeEphemeral string = "ephemeral"

	// V1IPResponseTypeStatic captures enum value "static"
	V1IPResponseTypeStatic string = "static"
)

// prop value enum
func (m *V1IPResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1IpResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1IPResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IPResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPResponse) UnmarshalBinary(b []byte) error {
	var res V1IPResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
