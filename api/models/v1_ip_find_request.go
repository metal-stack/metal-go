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

// V1IPFindRequest v1 IP find request
//
// swagger:model v1.IPFindRequest
type V1IPFindRequest struct {

	// addressfamily
	// Enum: ["IPv4","IPv6"]
	Addressfamily string `json:"addressfamily,omitempty" yaml:"addressfamily,omitempty"`

	// a unique identifier for this ip address allocation, can be used to distinguish between ip address allocation over time.
	Allocationuuid string `json:"allocationuuid,omitempty" yaml:"allocationuuid,omitempty"`

	// the address (ipv4 or ipv6) of this ip
	Ipaddress string `json:"ipaddress,omitempty" yaml:"ipaddress,omitempty"`

	// the machine an ip address is associated to
	Machineid string `json:"machineid,omitempty" yaml:"machineid,omitempty"`

	// the name of the ip address
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// the network this ip allocate request address belongs to
	Networkid string `json:"networkid,omitempty" yaml:"networkid,omitempty"`

	// the prefix of the network this ip address belongs to
	Networkprefix string `json:"networkprefix,omitempty" yaml:"networkprefix,omitempty"`

	// the project this ip address belongs to, empty if not strong coupled
	Projectid string `json:"projectid,omitempty" yaml:"projectid,omitempty"`

	// the tags that are assigned to this ip address
	Tags []string `json:"tags" yaml:"tags"`

	// the type of the ip address, ephemeral or static
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// Validate validates this v1 IP find request
func (m *V1IPFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressfamily(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1IpFindRequestTypeAddressfamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1IpFindRequestTypeAddressfamilyPropEnum = append(v1IpFindRequestTypeAddressfamilyPropEnum, v)
	}
}

const (

	// V1IPFindRequestAddressfamilyIPV4 captures enum value "IPv4"
	V1IPFindRequestAddressfamilyIPV4 string = "IPv4"

	// V1IPFindRequestAddressfamilyIPV6 captures enum value "IPv6"
	V1IPFindRequestAddressfamilyIPV6 string = "IPv6"
)

// prop value enum
func (m *V1IPFindRequest) validateAddressfamilyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1IpFindRequestTypeAddressfamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1IPFindRequest) validateAddressfamily(formats strfmt.Registry) error {
	if swag.IsZero(m.Addressfamily) { // not required
		return nil
	}

	// value enum
	if err := m.validateAddressfamilyEnum("addressfamily", "body", m.Addressfamily); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 IP find request based on context it is used
func (m *V1IPFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPFindRequest) UnmarshalBinary(b []byte) error {
	var res V1IPFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
