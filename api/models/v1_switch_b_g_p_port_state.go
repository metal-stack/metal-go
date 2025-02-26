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

// V1SwitchBGPPortState v1 switch b g p port state
//
// swagger:model v1.SwitchBGPPortState
type V1SwitchBGPPortState struct {

	// accepted prefix counter
	// Required: true
	AcceptedPrefixCounter *int64 `json:"AcceptedPrefixCounter" yaml:"AcceptedPrefixCounter"`

	// bgp state
	// Required: true
	BgpState *string `json:"BgpState" yaml:"BgpState"`

	// bgp timer up established
	// Required: true
	BgpTimerUpEstablished *int64 `json:"BgpTimerUpEstablished" yaml:"BgpTimerUpEstablished"`

	// neighbor
	// Required: true
	Neighbor *string `json:"Neighbor" yaml:"Neighbor"`

	// peer group
	// Required: true
	PeerGroup *string `json:"PeerGroup" yaml:"PeerGroup"`

	// sent prefix counter
	// Required: true
	SentPrefixCounter *int64 `json:"SentPrefixCounter" yaml:"SentPrefixCounter"`

	// vrf name
	// Required: true
	VrfName *string `json:"VrfName" yaml:"VrfName"`
}

// Validate validates this v1 switch b g p port state
func (m *V1SwitchBGPPortState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedPrefixCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBgpState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBgpTimerUpEstablished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNeighbor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSentPrefixCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrfName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchBGPPortState) validateAcceptedPrefixCounter(formats strfmt.Registry) error {

	if err := validate.Required("AcceptedPrefixCounter", "body", m.AcceptedPrefixCounter); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchBGPPortState) validateBgpState(formats strfmt.Registry) error {

	if err := validate.Required("BgpState", "body", m.BgpState); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchBGPPortState) validateBgpTimerUpEstablished(formats strfmt.Registry) error {

	if err := validate.Required("BgpTimerUpEstablished", "body", m.BgpTimerUpEstablished); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchBGPPortState) validateNeighbor(formats strfmt.Registry) error {

	if err := validate.Required("Neighbor", "body", m.Neighbor); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchBGPPortState) validatePeerGroup(formats strfmt.Registry) error {

	if err := validate.Required("PeerGroup", "body", m.PeerGroup); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchBGPPortState) validateSentPrefixCounter(formats strfmt.Registry) error {

	if err := validate.Required("SentPrefixCounter", "body", m.SentPrefixCounter); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchBGPPortState) validateVrfName(formats strfmt.Registry) error {

	if err := validate.Required("VrfName", "body", m.VrfName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 switch b g p port state based on context it is used
func (m *V1SwitchBGPPortState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SwitchBGPPortState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SwitchBGPPortState) UnmarshalBinary(b []byte) error {
	var res V1SwitchBGPPortState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
