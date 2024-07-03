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

// V1FirewallFindRequest v1 firewall find request
//
// swagger:model v1.FirewallFindRequest
type V1FirewallFindRequest struct {

	// allocation hostname
	AllocationHostname string `json:"allocation_hostname,omitempty" yaml:"allocation_hostname,omitempty"`

	// allocation image id
	AllocationImageID string `json:"allocation_image_id,omitempty" yaml:"allocation_image_id,omitempty"`

	// allocation name
	AllocationName string `json:"allocation_name,omitempty" yaml:"allocation_name,omitempty"`

	// allocation project
	AllocationProject string `json:"allocation_project,omitempty" yaml:"allocation_project,omitempty"`

	// allocation role
	AllocationRole string `json:"allocation_role,omitempty" yaml:"allocation_role,omitempty"`

	// allocation succeeded
	AllocationSucceeded bool `json:"allocation_succeeded,omitempty" yaml:"allocation_succeeded,omitempty"`

	// disk names
	DiskNames []string `json:"disk_names" yaml:"disk_names"`

	// disk sizes
	DiskSizes []int64 `json:"disk_sizes" yaml:"disk_sizes"`

	// fru board mfg
	FruBoardMfg string `json:"fru_board_mfg,omitempty" yaml:"fru_board_mfg,omitempty"`

	// fru board mfg serial
	FruBoardMfgSerial string `json:"fru_board_mfg_serial,omitempty" yaml:"fru_board_mfg_serial,omitempty"`

	// fru board part number
	FruBoardPartNumber string `json:"fru_board_part_number,omitempty" yaml:"fru_board_part_number,omitempty"`

	// fru chassis part number
	FruChassisPartNumber string `json:"fru_chassis_part_number,omitempty" yaml:"fru_chassis_part_number,omitempty"`

	// fru chassis part serial
	FruChassisPartSerial string `json:"fru_chassis_part_serial,omitempty" yaml:"fru_chassis_part_serial,omitempty"`

	// fru product manufacturer
	FruProductManufacturer string `json:"fru_product_manufacturer,omitempty" yaml:"fru_product_manufacturer,omitempty"`

	// fru product part number
	FruProductPartNumber string `json:"fru_product_part_number,omitempty" yaml:"fru_product_part_number,omitempty"`

	// fru product serial
	FruProductSerial string `json:"fru_product_serial,omitempty" yaml:"fru_product_serial,omitempty"`

	// hardware memory
	HardwareMemory int64 `json:"hardware_memory,omitempty" yaml:"hardware_memory,omitempty"`

	// id
	ID string `json:"id,omitempty" yaml:"id,omitempty"`

	// ipmi address
	IpmiAddress string `json:"ipmi_address,omitempty" yaml:"ipmi_address,omitempty"`

	// ipmi interface
	IpmiInterface string `json:"ipmi_interface,omitempty" yaml:"ipmi_interface,omitempty"`

	// ipmi mac address
	IpmiMacAddress string `json:"ipmi_mac_address,omitempty" yaml:"ipmi_mac_address,omitempty"`

	// ipmi user
	IpmiUser string `json:"ipmi_user,omitempty" yaml:"ipmi_user,omitempty"`

	// name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// network asns
	NetworkAsns []int64 `json:"network_asns" yaml:"network_asns"`

	// network destination prefixes
	NetworkDestinationPrefixes []string `json:"network_destination_prefixes" yaml:"network_destination_prefixes"`

	// network ids
	NetworkIds []string `json:"network_ids" yaml:"network_ids"`

	// network ips
	NetworkIps []string `json:"network_ips" yaml:"network_ips"`

	// network prefixes
	NetworkPrefixes []string `json:"network_prefixes" yaml:"network_prefixes"`

	// network vrfs
	NetworkVrfs []int64 `json:"network_vrfs" yaml:"network_vrfs"`

	// nics mac addresses
	NicsMacAddresses []string `json:"nics_mac_addresses" yaml:"nics_mac_addresses"`

	// nics names
	NicsNames []string `json:"nics_names" yaml:"nics_names"`

	// nics neighbor mac addresses
	NicsNeighborMacAddresses []string `json:"nics_neighbor_mac_addresses" yaml:"nics_neighbor_mac_addresses"`

	// nics neighbor names
	NicsNeighborNames []string `json:"nics_neighbor_names" yaml:"nics_neighbor_names"`

	// nics neighbor vrfs
	NicsNeighborVrfs []string `json:"nics_neighbor_vrfs" yaml:"nics_neighbor_vrfs"`

	// nics vrfs
	NicsVrfs []string `json:"nics_vrfs" yaml:"nics_vrfs"`

	// partition id
	PartitionID string `json:"partition_id,omitempty" yaml:"partition_id,omitempty"`

	// rackid
	Rackid string `json:"rackid,omitempty" yaml:"rackid,omitempty"`

	// sizeid
	Sizeid string `json:"sizeid,omitempty" yaml:"sizeid,omitempty"`

	// state value
	// Enum: ["","LOCKED","RESERVED"]
	StateValue string `json:"state_value,omitempty" yaml:"state_value,omitempty"`

	// tags
	Tags []string `json:"tags" yaml:"tags"`
}

// Validate validates this v1 firewall find request
func (m *V1FirewallFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1FirewallFindRequestTypeStateValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","LOCKED","RESERVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1FirewallFindRequestTypeStateValuePropEnum = append(v1FirewallFindRequestTypeStateValuePropEnum, v)
	}
}

const (

	// V1FirewallFindRequestStateValueEmpty captures enum value ""
	V1FirewallFindRequestStateValueEmpty string = ""

	// V1FirewallFindRequestStateValueLOCKED captures enum value "LOCKED"
	V1FirewallFindRequestStateValueLOCKED string = "LOCKED"

	// V1FirewallFindRequestStateValueRESERVED captures enum value "RESERVED"
	V1FirewallFindRequestStateValueRESERVED string = "RESERVED"
)

// prop value enum
func (m *V1FirewallFindRequest) validateStateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1FirewallFindRequestTypeStateValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1FirewallFindRequest) validateStateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.StateValue) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateValueEnum("state_value", "body", m.StateValue); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 firewall find request based on context it is used
func (m *V1FirewallFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FirewallFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FirewallFindRequest) UnmarshalBinary(b []byte) error {
	var res V1FirewallFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
