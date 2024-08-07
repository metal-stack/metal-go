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

// V1FirewallEgressRule v1 firewall egress rule
//
// swagger:model v1.FirewallEgressRule
type V1FirewallEgressRule struct {

	// an optional comment describing what this rule is used for
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// the ports affected by this rule
	// Required: true
	Ports []int32 `json:"ports" yaml:"ports"`

	// the protocol for the rule, defaults to tcp
	// Enum: ["tcp","udp"]
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// the cidrs affected by this rule
	// Required: true
	To []string `json:"to" yaml:"to"`
}

// Validate validates this v1 firewall egress rule
func (m *V1FirewallEgressRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FirewallEgressRule) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	return nil
}

var v1FirewallEgressRuleTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1FirewallEgressRuleTypeProtocolPropEnum = append(v1FirewallEgressRuleTypeProtocolPropEnum, v)
	}
}

const (

	// V1FirewallEgressRuleProtocolTCP captures enum value "tcp"
	V1FirewallEgressRuleProtocolTCP string = "tcp"

	// V1FirewallEgressRuleProtocolUDP captures enum value "udp"
	V1FirewallEgressRuleProtocolUDP string = "udp"
)

// prop value enum
func (m *V1FirewallEgressRule) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1FirewallEgressRuleTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1FirewallEgressRule) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *V1FirewallEgressRule) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 firewall egress rule based on context it is used
func (m *V1FirewallEgressRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FirewallEgressRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FirewallEgressRule) UnmarshalBinary(b []byte) error {
	var res V1FirewallEgressRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
