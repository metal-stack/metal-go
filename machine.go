package metalgo

import (
	"github.com/metal-stack/metal-go/client/operations"
	"github.com/metal-stack/metal-go/models"
)

// MachineCreateRequest contains data for a machine creation
type MachineCreateRequest struct {
	Description      string
	Hostname         string
	Name             string
	UserData         string
	Size             string
	FilesystemLayout string
	Project          string
	Partition        string
	Image            string
	Tags             []string
	SSHPublicKeys    []string
	UUID             string
	Networks         []MachineAllocationNetwork
	IPs              []string
}

// MachineCreateRequest contains data for a machine update
type MachineUpdateRequest struct {
	ID          string
	Description *string
	Tags        []string
}

// MachineFindRequest contains criteria for a machine listing
type MachineFindRequest struct {
	ID          *string
	Name        *string
	PartitionID *string
	SizeID      *string
	RackID      *string
	Tags        []string

	// allocation
	AllocationName      *string
	AllocationProject   *string
	AllocationImageID   *string
	AllocationHostname  *string
	AllocationSucceeded *bool
	AllocationRole      *string

	// network
	NetworkIDs                 []string
	NetworkPrefixes            []string
	NetworkIPs                 []string
	NetworkDestinationPrefixes []string
	NetworkVrfs                []int64
	NetworkPrivate             *bool
	NetworkASNs                []int64
	NetworkNat                 *bool
	NetworkUnderlay            *bool

	// hardware
	HardwareMemory   *int64
	HardwareCPUCores *int64

	// nics
	NicsMacAddresses         []string
	NicsNames                []string
	NicsVrfs                 []string
	NicsNeighborMacAddresses []string
	NicsNeighborNames        []string
	NicsNeighborVrfs         []string

	// disks
	DiskNames []string
	DiskSizes []int64

	// state
	StateValue *string

	// ipmi
	IpmiAddress    *string
	IpmiMacAddress *string
	IpmiUser       *string
	IpmiInterface  *string

	// fru
	FruChassisPartNumber   *string
	FruChassisPartSerial   *string
	FruBoardMfg            *string
	FruBoardMfgSerial      *string
	FruBoardPartNumber     *string
	FruProductManufacturer *string
	FruProductPartNumber   *string
	FruProductSerial       *string
}

// MachineAllocationNetwork contains configuration for machine networks
type MachineAllocationNetwork struct {
	Autoacquire bool
	NetworkID   string
}

func (n *MachineCreateRequest) translateNetworks() []*models.V1MachineAllocationNetwork {
	var nets []*models.V1MachineAllocationNetwork
	for i := range n.Networks {
		net := models.V1MachineAllocationNetwork{
			Networkid:   &n.Networks[i].NetworkID,
			Autoacquire: &n.Networks[i].Autoacquire,
		}
		nets = append(nets, &net)
	}
	return nets
}

// MachineCreateResponse is returned when a machine was created
type MachineCreateResponse struct {
	Machine *models.V1MachineResponse
}

// MachineUpdateResponse is returned when a machine was updated
type MachineUpdateResponse struct {
	Machine *models.V1MachineResponse
}

// MachineListResponse contains the machine list result
type MachineListResponse struct {
	Machines []*models.V1MachineResponse
}

// MachineGetResponse contains the machine get result
type MachineGetResponse struct {
	Machine *models.V1MachineResponse
}

// MachineIPMIGetResponse contains the machine ipmi get result
type MachineIPMIGetResponse struct {
	Machine *models.V1MachineIPMIResponse
}

// MachineIPMIListResponse contains the machine ipmi list result
type MachineIPMIListResponse struct {
	Machines []*models.V1MachineIPMIResponse
}

// MachineIPMIReports contains the machine ipmi report
type MachineIPMIReports struct {
	Reports *models.V1MachineIpmiReports
}

// MachineIPMIReportResponse contains the machine ipmi report result
type MachineIPMIReportResponse struct {
	Response *models.V1MachineIpmiReportResponse
}

// MachineDeleteResponse contains the machine delete result
type MachineDeleteResponse struct {
	Machine *models.V1MachineResponse
}

// MachinePowerResponse contains the machine power result
type MachinePowerResponse struct {
	Machine *models.V1MachineResponse
}

// ChassisIdentifyLEDPowerResponse contains the machine LED power result
type ChassisIdentifyLEDPowerResponse struct {
	Machine *models.V1MachineResponse
}

type FirmwareKind string

const (
	Bios FirmwareKind = "bios"
	Bmc  FirmwareKind = "bmc"
)

// MachineDiskResponse contains the machine Disk result
type MachineDiskResponse struct {
	Machine *models.V1MachineResponse
}

// MachinePxeResponse contains the machine Pxe result
type MachinePxeResponse struct {
	Machine *models.V1MachineResponse
}

// MachineStateResponse contains the machine bios result
type MachineStateResponse struct {
	Machine *models.V1MachineResponse
}

// MachineCreate creates a single metal machine
func (d *Driver) MachineCreate(mcr *MachineCreateRequest) (*MachineCreateResponse, error) {
	response := &MachineCreateResponse{}

	allocateRequest := &models.V1MachineAllocateRequest{
		Description:        mcr.Description,
		Partitionid:        &mcr.Partition,
		Hostname:           mcr.Hostname,
		Imageid:            &mcr.Image,
		Name:               mcr.Name,
		UUID:               mcr.UUID,
		Projectid:          &mcr.Project,
		Sizeid:             &mcr.Size,
		Filesystemlayoutid: mcr.FilesystemLayout,
		SSHPubKeys:         mcr.SSHPublicKeys,
		UserData:           mcr.UserData,
		Tags:               mcr.Tags,
		Networks:           mcr.translateNetworks(),
		Ips:                mcr.IPs,
	}
	allocMachine := operations.NewAllocateMachineParams()
	allocMachine.SetBody(allocateRequest)

	resp, err := d.Client.AllocateMachine(allocMachine, nil)
	if err != nil {
		return response, err
	}

	response.Machine = resp.Payload

	return response, nil
}

// MachineUpdate updated a single metal machine, be cautios that this can remove tags when empty tags are passed
func (d *Driver) MachineUpdate(mur *MachineUpdateRequest) (*MachineUpdateResponse, error) {
	body := &models.V1MachineUpdateRequest{
		ID:   &mur.ID,
		Tags: mur.Tags,
	}

	if mur.Description != nil {
		body.Description = mur.Description
	}

	params := machine.NewUpdateMachineParams().WithBody(body)

	response := &MachineUpdateResponse{}
	resp, err := d.machine.UpdateMachine(params, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineDelete deletes a single metal machine
func (d *Driver) MachineDelete(machineID string) (*MachineDeleteResponse, error) {
	freeMachine := operations.NewFreeMachineParams()
	freeMachine.ID = machineID

	response := &MachineDeleteResponse{}
	resp, err := d.Client.FreeMachine(freeMachine, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineDeleteFromDatabase deletes a single metal machine from the database
func (d *Driver) MachineDeleteFromDatabase(machineID string) (*MachineDeleteResponse, error) {
	p := operations.NewDeleteMachineParams().WithID(machineID)

	response := &MachineDeleteResponse{}
	resp, err := d.Client.DeleteMachine(p, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineGet returns the machine with the given ID
func (d *Driver) MachineGet(id string) (*MachineGetResponse, error) {
	findMachine := operations.NewFindMachineParams()
	findMachine.ID = id

	response := &MachineGetResponse{}
	resp, err := d.Client.FindMachine(findMachine, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload

	return response, nil
}

// MachineConsolePassword returns the consolepassword of the machine
func (d *Driver) MachineConsolePassword(id, reason string) (*models.V1MachineConsolePasswordResponse, error) {
	cp := operations.NewGetMachineConsolePasswordParams()
	cp.Body = &models.V1MachineConsolePasswordRequest{
		ID:     &id,
		Reason: &reason,
	}

	resp, err := d.Client.GetMachineConsolePassword(cp, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// MachineList lists all machines
func (d *Driver) MachineList() (*MachineListResponse, error) {
	listMachines := operations.NewListMachinesParams()
	response := &MachineListResponse{}
	resp, err := d.Client.ListMachines(listMachines, nil)
	if err != nil {
		return response, err
	}
	response.Machines = resp.Payload

	return response, nil
}

// MachineFind lists all machines that match the given properties
func (d *Driver) MachineFind(mfr *MachineFindRequest) (*MachineListResponse, error) {
	if mfr == nil {
		return d.MachineList()
	}

	response := &MachineListResponse{}
	var err error
	var resp *operations.FindMachinesOK

	findMachines := operations.NewFindMachinesParams()
	req := &models.V1MachineFindRequest{
		ID:                         StrDeref(mfr.ID),
		Name:                       StrDeref(mfr.Name),
		PartitionID:                StrDeref(mfr.PartitionID),
		Sizeid:                     StrDeref(mfr.SizeID),
		Rackid:                     StrDeref(mfr.RackID),
		Tags:                       mfr.Tags,
		AllocationName:             StrDeref(mfr.AllocationName),
		AllocationProject:          StrDeref(mfr.AllocationProject),
		AllocationImageID:          StrDeref(mfr.AllocationImageID),
		AllocationHostname:         StrDeref(mfr.AllocationHostname),
		AllocationRole:             StrDeref(mfr.AllocationRole),
		AllocationSucceeded:        BoolDeref(mfr.AllocationSucceeded),
		NetworkIds:                 mfr.NetworkIDs,
		NetworkPrefixes:            mfr.NetworkPrefixes,
		NetworkIps:                 mfr.NetworkIPs,
		NetworkDestinationPrefixes: mfr.NetworkDestinationPrefixes,
		NetworkVrfs:                mfr.NetworkVrfs,
		NetworkPrivate:             BoolDeref(mfr.NetworkPrivate),
		NetworkAsns:                mfr.NetworkASNs,
		NetworkNat:                 BoolDeref(mfr.NetworkNat),
		NetworkUnderlay:            BoolDeref(mfr.NetworkUnderlay),
		HardwareMemory:             Int64Deref(mfr.HardwareMemory),
		HardwareCPUCores:           Int64Deref(mfr.HardwareCPUCores),
		NicsMacAddresses:           mfr.NicsMacAddresses,
		NicsNames:                  mfr.NicsNames,
		NicsVrfs:                   mfr.NicsVrfs,
		NicsNeighborMacAddresses:   mfr.NicsNeighborMacAddresses,
		NicsNeighborNames:          mfr.NicsNeighborNames,
		NicsNeighborVrfs:           mfr.NicsNeighborVrfs,
		DiskNames:                  mfr.DiskNames,
		DiskSizes:                  mfr.DiskSizes,
		StateValue:                 StrDeref(mfr.StateValue),
		IpmiAddress:                StrDeref(mfr.IpmiAddress),
		IpmiMacAddress:             StrDeref(mfr.IpmiMacAddress),
		IpmiUser:                   StrDeref(mfr.IpmiUser),
		IpmiInterface:              StrDeref(mfr.IpmiInterface),
		FruChassisPartNumber:       StrDeref(mfr.FruChassisPartNumber),
		FruChassisPartSerial:       StrDeref(mfr.FruChassisPartSerial),
		FruBoardMfg:                StrDeref(mfr.FruBoardMfg),
		FruBoardMfgSerial:          StrDeref(mfr.FruBoardMfgSerial),
		FruBoardPartNumber:         StrDeref(mfr.FruChassisPartNumber),
		FruProductManufacturer:     StrDeref(mfr.FruProductManufacturer),
		FruProductPartNumber:       StrDeref(mfr.FruProductPartNumber),
		FruProductSerial:           StrDeref(mfr.FruProductSerial),
	}
	findMachines.SetBody(req)

	resp, err = d.Client.FindMachines(findMachines, nil)
	if err != nil {
		return response, err
	}
	response.Machines = resp.Payload

	return response, nil
}

// MachineIPMIGet returns the machine with the given ID including IPMI data
func (d *Driver) MachineIPMIGet(id string) (*MachineIPMIGetResponse, error) {
	findMachine := operations.NewFindIPMIMachineParams().WithID(id)

	response := &MachineIPMIGetResponse{}
	resp, err := d.Client.FindIPMIMachine(findMachine, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload

	return response, nil
}

// MachineIPMIList returns the machine list of the given search query including IPMI data
func (d *Driver) MachineIPMIList(mfr *MachineFindRequest) (*MachineIPMIListResponse, error) {
	response := &MachineIPMIListResponse{}

	findMachines := operations.NewFindIPMIMachinesParams()
	req := &models.V1MachineFindRequest{
		ID:                         StrDeref(mfr.ID),
		Name:                       StrDeref(mfr.Name),
		PartitionID:                StrDeref(mfr.PartitionID),
		Sizeid:                     StrDeref(mfr.SizeID),
		Rackid:                     StrDeref(mfr.RackID),
		Tags:                       mfr.Tags,
		AllocationName:             StrDeref(mfr.AllocationName),
		AllocationProject:          StrDeref(mfr.AllocationProject),
		AllocationImageID:          StrDeref(mfr.AllocationImageID),
		AllocationHostname:         StrDeref(mfr.AllocationHostname),
		AllocationSucceeded:        BoolDeref(mfr.AllocationSucceeded),
		AllocationRole:             StrDeref(mfr.AllocationRole),
		NetworkIds:                 mfr.NetworkIDs,
		NetworkPrefixes:            mfr.NetworkPrefixes,
		NetworkIps:                 mfr.NetworkIPs,
		NetworkDestinationPrefixes: mfr.NetworkDestinationPrefixes,
		NetworkVrfs:                mfr.NetworkVrfs,
		NetworkPrivate:             BoolDeref(mfr.NetworkPrivate),
		NetworkAsns:                mfr.NetworkASNs,
		NetworkNat:                 BoolDeref(mfr.NetworkNat),
		NetworkUnderlay:            BoolDeref(mfr.NetworkUnderlay),
		HardwareMemory:             Int64Deref(mfr.HardwareMemory),
		HardwareCPUCores:           Int64Deref(mfr.HardwareCPUCores),
		NicsMacAddresses:           mfr.NicsMacAddresses,
		NicsNames:                  mfr.NicsNames,
		NicsVrfs:                   mfr.NicsVrfs,
		NicsNeighborMacAddresses:   mfr.NicsNeighborMacAddresses,
		NicsNeighborNames:          mfr.NicsNeighborNames,
		NicsNeighborVrfs:           mfr.NicsNeighborVrfs,
		DiskNames:                  mfr.DiskNames,
		DiskSizes:                  mfr.DiskSizes,
		StateValue:                 StrDeref(mfr.StateValue),
		IpmiAddress:                StrDeref(mfr.IpmiAddress),
		IpmiMacAddress:             StrDeref(mfr.IpmiMacAddress),
		IpmiUser:                   StrDeref(mfr.IpmiUser),
		IpmiInterface:              StrDeref(mfr.IpmiInterface),
		FruChassisPartNumber:       StrDeref(mfr.FruChassisPartNumber),
		FruChassisPartSerial:       StrDeref(mfr.FruChassisPartSerial),
		FruBoardMfg:                StrDeref(mfr.FruBoardMfg),
		FruBoardMfgSerial:          StrDeref(mfr.FruBoardMfgSerial),
		FruBoardPartNumber:         StrDeref(mfr.FruChassisPartNumber),
		FruProductManufacturer:     StrDeref(mfr.FruProductManufacturer),
		FruProductPartNumber:       StrDeref(mfr.FruProductPartNumber),
		FruProductSerial:           StrDeref(mfr.FruProductSerial),
	}
	findMachines.SetBody(req)

	resp, err := d.Client.FindIPMIMachines(findMachines, nil)
	if err != nil {
		return response, err
	}
	response.Machines = resp.Payload

	return response, nil
}

// MachineIPMIReport send leases of this partition to the metal-api
func (d *Driver) MachineIPMIReport(report MachineIPMIReports) (*MachineIPMIReportResponse, error) {
	params := operations.NewIpmiReportParams()
	params.SetBody(report.Reports)
	ok, err := d.Client.IpmiReport(params, nil)
	if err != nil {
		return nil, err
	}
	return &MachineIPMIReportResponse{
		Response: ok.Payload,
	}, nil
}

// MachinePowerOn powers on the given machine
func (d *Driver) MachinePowerOn(machineID string) (*MachinePowerResponse, error) {
	machineOn := operations.NewMachineOnParams()
	machineOn.ID = machineID
	machineOn.Body = []string{}

	response := &MachinePowerResponse{}
	resp, err := d.Client.MachineOn(machineOn, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachinePowerOff powers off the given machine
func (d *Driver) MachinePowerOff(machineID string) (*MachinePowerResponse, error) {
	machineOff := operations.NewMachineOffParams()
	machineOff.ID = machineID
	machineOff.Body = []string{}

	response := &MachinePowerResponse{}
	resp, err := d.Client.MachineOff(machineOff, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachinePowerReset power-resets the given machine
func (d *Driver) MachinePowerReset(machineID string) (*MachinePowerResponse, error) {
	machineReset := operations.NewMachineResetParams()
	machineReset.ID = machineID
	machineReset.Body = []string{}

	response := &MachinePowerResponse{}
	resp, err := d.Client.MachineReset(machineReset, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachinePowerCycle power-cycles the given machine
func (d *Driver) MachinePowerCycle(machineID string) (*MachinePowerResponse, error) {
	machineCycle := operations.NewMachineCycleParams()
	machineCycle.ID = machineID
	machineCycle.Body = []string{}

	response := &MachinePowerResponse{}
	resp, err := d.Client.MachineCycle(machineCycle, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineBootBios boots given machine into BIOS
func (d *Driver) MachineBootBios(machineID string) (*MachineFirmwareResponse, error) {
	machineBios := operations.NewMachineBiosParams()
	machineBios.ID = machineID
	machineBios.Body = []string{}

	response := &MachineFirmwareResponse{}
	resp, err := d.Client.MachineBios(machineBios, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineBootDisk boots given machine from disk
func (d *Driver) MachineBootDisk(machineID string) (*MachineDiskResponse, error) {
	machineDisk := operations.NewMachineDiskParams()
	machineDisk.ID = machineID
	machineDisk.Body = []string{}

	response := &MachineDiskResponse{}
	resp, err := d.Client.MachineDisk(machineDisk, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineBootPxe boots given machine from PXE
func (d *Driver) MachineBootPxe(machineID string) (*MachinePxeResponse, error) {
	machinePxe := operations.NewMachinePxeParams()
	machinePxe.ID = machineID
	machinePxe.Body = []string{}

	response := &MachinePxeResponse{}
	resp, err := d.Client.MachinePxe(machinePxe, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// MachineLock locks a machine to prevent it from being destroyed
func (d *Driver) MachineLock(machineID, description string) (*MachineStateResponse, error) {
	return d.machineState(machineID, "LOCKED", description)
}

// MachineUnLock unlocks a machine
func (d *Driver) MachineUnLock(machineID string) (*MachineStateResponse, error) {
	return d.machineState(machineID, "", "")
}

// MachineReserve reserves a machine for single allocation
func (d *Driver) MachineReserve(machineID, description string) (*MachineStateResponse, error) {
	return d.machineState(machineID, "RESERVED", description)
}

// MachineUnReserve unreserves a machine
func (d *Driver) MachineUnReserve(machineID string) (*MachineStateResponse, error) {
	return d.machineState(machineID, "", "")
}

// MachineReinstall installs given image on already allocated machine
func (d *Driver) MachineReinstall(machineID, imageID, description string) (*MachineGetResponse, error) {
	request := &models.V1MachineReinstallRequest{
		Imageid:     &imageID,
		Description: description,
	}
	machineReinstall := operations.NewReinstallMachineParams()
	machineReinstall.ID = machineID
	machineReinstall.Body = request

	response := &MachineGetResponse{}
	resp, err := d.Client.ReinstallMachine(machineReinstall, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

func (d *Driver) machineState(machineID, state, description string) (*MachineStateResponse, error) {
	machineState := operations.NewSetMachineStateParams()
	machineState.ID = machineID
	machineState.Body = &models.V1MachineState{
		Value:       &state,
		Description: &description,
	}

	response := &MachineStateResponse{}
	resp, err := d.Client.SetMachineState(machineState, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// ChassisIdentifyLEDPowerOn powers on the given machine
func (d *Driver) ChassisIdentifyLEDPowerOn(machineID, description string) (*ChassisIdentifyLEDPowerResponse, error) {
	machineLedOn := operations.NewChassisIdentifyLEDOnParams()
	machineLedOn.ID = machineID
	machineLedOn.Description = &description
	machineLedOn.Body = []string{}

	response := &ChassisIdentifyLEDPowerResponse{}
	resp, err := d.Client.ChassisIdentifyLEDOn(machineLedOn, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}

// ChassisIdentifyLEDPowerOff powers off the given machine
func (d *Driver) ChassisIdentifyLEDPowerOff(machineID, description string) (*ChassisIdentifyLEDPowerResponse, error) {
	machineLedOff := operations.NewChassisIdentifyLEDOffParams()
	machineLedOff.ID = machineID
	machineLedOff.Description = &description
	machineLedOff.Body = []string{}

	response := &ChassisIdentifyLEDPowerResponse{}
	resp, err := d.Client.ChassisIdentifyLEDOff(machineLedOff, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}
