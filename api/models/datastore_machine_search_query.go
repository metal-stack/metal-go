// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatastoreMachineSearchQuery datastore machine search query
//
// swagger:model datastore.MachineSearchQuery
type DatastoreMachineSearchQuery struct {

	// allocation hostname
	AllocationHostname string `json:"allocation_hostname,omitempty"`

	// allocation image id
	AllocationImageID string `json:"allocation_image_id,omitempty"`

	// allocation name
	AllocationName string `json:"allocation_name,omitempty"`

	// allocation project
	AllocationProject string `json:"allocation_project,omitempty"`

	// allocation succeeded
	AllocationSucceeded bool `json:"allocation_succeeded,omitempty"`

	// disk names
	DiskNames []string `json:"disk_names"`

	// disk sizes
	DiskSizes []int64 `json:"disk_sizes"`

	// fru board mfg
	FruBoardMfg string `json:"fru_board_mfg,omitempty"`

	// fru board mfg serial
	FruBoardMfgSerial string `json:"fru_board_mfg_serial,omitempty"`

	// fru board part number
	FruBoardPartNumber string `json:"fru_board_part_number,omitempty"`

	// fru chassis part number
	FruChassisPartNumber string `json:"fru_chassis_part_number,omitempty"`

	// fru chassis part serial
	FruChassisPartSerial string `json:"fru_chassis_part_serial,omitempty"`

	// fru product manufacturer
	FruProductManufacturer string `json:"fru_product_manufacturer,omitempty"`

	// fru product part number
	FruProductPartNumber string `json:"fru_product_part_number,omitempty"`

	// fru product serial
	FruProductSerial string `json:"fru_product_serial,omitempty"`

	// hardware cpu cores
	HardwareCPUCores int64 `json:"hardware_cpu_cores,omitempty"`

	// hardware memory
	HardwareMemory int64 `json:"hardware_memory,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ipmi address
	IpmiAddress string `json:"ipmi_address,omitempty"`

	// ipmi interface
	IpmiInterface string `json:"ipmi_interface,omitempty"`

	// ipmi mac address
	IpmiMacAddress string `json:"ipmi_mac_address,omitempty"`

	// ipmi user
	IpmiUser string `json:"ipmi_user,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// network asns
	NetworkAsns []int64 `json:"network_asns"`

	// network destination prefixes
	NetworkDestinationPrefixes []string `json:"network_destination_prefixes"`

	// network ids
	NetworkIds []string `json:"network_ids"`

	// network ips
	NetworkIps []string `json:"network_ips"`

	// network nat
	NetworkNat bool `json:"network_nat,omitempty"`

	// network prefixes
	NetworkPrefixes []string `json:"network_prefixes"`

	// network private
	NetworkPrivate bool `json:"network_private,omitempty"`

	// network underlay
	NetworkUnderlay bool `json:"network_underlay,omitempty"`

	// network vrfs
	NetworkVrfs []int64 `json:"network_vrfs"`

	// nics mac addresses
	NicsMacAddresses []string `json:"nics_mac_addresses"`

	// nics names
	NicsNames []string `json:"nics_names"`

	// nics neighbor mac addresses
	NicsNeighborMacAddresses []string `json:"nics_neighbor_mac_addresses"`

	// nics neighbor names
	NicsNeighborNames []string `json:"nics_neighbor_names"`

	// nics neighbor vrfs
	NicsNeighborVrfs []string `json:"nics_neighbor_vrfs"`

	// nics vrfs
	NicsVrfs []string `json:"nics_vrfs"`

	// partition id
	PartitionID string `json:"partition_id,omitempty"`

	// rackid
	Rackid string `json:"rackid,omitempty"`

	// sizeid
	Sizeid string `json:"sizeid,omitempty"`

	// state value
	StateValue string `json:"state_value,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this datastore machine search query
func (m *DatastoreMachineSearchQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datastore machine search query based on context it is used
func (m *DatastoreMachineSearchQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreMachineSearchQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreMachineSearchQuery) UnmarshalBinary(b []byte) error {
	var res DatastoreMachineSearchQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}