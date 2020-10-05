package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/firewall"
	"github.com/metal-stack/metal-go/api/models"
)

// FirewallCreateRequest contains data for a machine creation
type FirewallCreateRequest struct {
	MachineCreateRequest
}

// FirewallCreateResponse is returned when a machine was created
type FirewallCreateResponse struct {
	Firewall *models.V1FirewallResponse
}

// FirewallFindRequest contains criteria for a machine listing
type FirewallFindRequest struct {
	MachineFindRequest
}

// FirewallListResponse contains the machine list result
type FirewallListResponse struct {
	Firewalls []*models.V1FirewallResponse
}

// FirewallGetResponse contains the machine get result
type FirewallGetResponse struct {
	Firewall *models.V1FirewallResponse
}

// FirewallCreate will create a single metal machine
func (d *Driver) FirewallCreate(fcr *FirewallCreateRequest) (*FirewallCreateResponse, error) {
	response := &FirewallCreateResponse{}

	allocateRequest := &models.V1FirewallCreateRequest{
		Description: fcr.Description,
		Partitionid: &fcr.Partition,
		Hostname:    fcr.Hostname,
		Imageid:     &fcr.Image,
		Name:        fcr.Name,
		UUID:        fcr.UUID,
		Projectid:   &fcr.Project,
		Sizeid:      &fcr.Size,
		SSHPubKeys:  fcr.SSHPublicKeys,
		UserData:    fcr.UserData,
		Tags:        fcr.Tags,
		Networks:    fcr.translateNetworks(),
		Ips:         fcr.IPs,
	}

	allocFirewall := firewall.NewAllocateFirewallParams()
	allocFirewall.SetBody(allocateRequest)

	resp, err := d.firewall.AllocateFirewall(allocFirewall, nil)
	if err != nil {
		return response, err
	}
	response.Firewall = resp.Payload

	return response, nil
}

// FirewallList will list all machines
func (d *Driver) FirewallList() (*FirewallListResponse, error) {
	response := &FirewallListResponse{}

	listFirewall := firewall.NewListFirewallsParams()
	resp, err := d.firewall.ListFirewalls(listFirewall, nil)
	if err != nil {
		return response, err
	}
	response.Firewalls = resp.Payload
	return response, nil
}

// FirewallFind will search for firewalls for given criteria
func (d *Driver) FirewallFind(ffr *FirewallFindRequest) (*FirewallListResponse, error) {
	if ffr == nil {
		return d.FirewallList()
	}

	response := &FirewallListResponse{}
	var err error
	var resp *firewall.FindFirewallsOK

	req := &models.V1FirewallFindRequest{
		ID:                         StrDeref(ffr.ID),
		Name:                       StrDeref(ffr.Name),
		PartitionID:                StrDeref(ffr.PartitionID),
		Sizeid:                     StrDeref(ffr.SizeID),
		Rackid:                     StrDeref(ffr.RackID),
		Liveliness:                 StrDeref(ffr.Liveliness),
		Tags:                       ffr.Tags,
		AllocationName:             StrDeref(ffr.AllocationName),
		AllocationProject:          StrDeref(ffr.AllocationProject),
		AllocationImageID:          StrDeref(ffr.AllocationImageID),
		AllocationHostname:         StrDeref(ffr.AllocationHostname),
		AllocationSucceeded:        BoolDeref(ffr.AllocationSucceeded),
		NetworkIds:                 ffr.NetworkIDs,
		NetworkPrefixes:            ffr.NetworkPrefixes,
		NetworkIps:                 ffr.NetworkIPs,
		NetworkDestinationPrefixes: ffr.NetworkDestinationPrefixes,
		NetworkVrfs:                ffr.NetworkVrfs,
		NetworkPrivate:             BoolDeref(ffr.NetworkPrivate),
		NetworkAsns:                ffr.NetworkASNs,
		NetworkNat:                 BoolDeref(ffr.NetworkNat),
		NetworkUnderlay:            BoolDeref(ffr.NetworkUnderlay),
		HardwareMemory:             Int64Deref(ffr.HardwareMemory),
		HardwareCPUCores:           Int64Deref(ffr.HardwareCPUCores),
		NicsMacAddresses:           ffr.NicsMacAddresses,
		NicsNames:                  ffr.NicsNames,
		NicsVrfs:                   ffr.NicsVrfs,
		NicsNeighborMacAddresses:   ffr.NicsNeighborMacAddresses,
		NicsNeighborNames:          ffr.NicsNeighborNames,
		NicsNeighborVrfs:           ffr.NicsNeighborVrfs,
		DiskNames:                  ffr.DiskNames,
		DiskSizes:                  ffr.DiskSizes,
		StateValue:                 StrDeref(ffr.StateValue),
		IPMIAddress:                StrDeref(ffr.IpmiAddress),
		IPMIMacAddress:             StrDeref(ffr.IpmiMacAddress),
		IPMIUser:                   StrDeref(ffr.IpmiUser),
		IPMIInterface:              StrDeref(ffr.IpmiInterface),
		FruChassisPartNumber:       StrDeref(ffr.FruChassisPartNumber),
		FruChassisPartSerial:       StrDeref(ffr.FruChassisPartSerial),
		FruBoardMfg:                StrDeref(ffr.FruBoardMfg),
		FruBoardMfgSerial:          StrDeref(ffr.FruBoardMfgSerial),
		FruBoardPartNumber:         StrDeref(ffr.FruChassisPartNumber),
		FruProductManufacturer:     StrDeref(ffr.FruProductManufacturer),
		FruProductPartNumber:       StrDeref(ffr.FruProductPartNumber),
		FruProductSerial:           StrDeref(ffr.FruProductSerial),
	}
	findFirewalls := firewall.NewFindFirewallsParams()
	findFirewalls.SetBody(req)

	resp, err = d.firewall.FindFirewalls(findFirewalls, nil)
	if err != nil {
		return response, err
	}
	response.Firewalls = resp.Payload
	return response, nil
}

// FirewallGet will only return one machine
func (d *Driver) FirewallGet(machineID string) (*FirewallGetResponse, error) {
	findFirewall := firewall.NewFindFirewallParams()
	findFirewall.ID = machineID

	response := &FirewallGetResponse{}
	resp, err := d.firewall.FindFirewall(findFirewall, nil)
	if err != nil {
		return response, err
	}
	response.Firewall = resp.Payload

	return response, nil
}
