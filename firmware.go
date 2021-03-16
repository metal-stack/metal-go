package metalgo

import (
	"github.com/go-openapi/runtime"
	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/models"
	"os"
	"time"
)

// AvailableFirmwaresResponse contains available firmwares matching the requested parameters
type AvailableFirmwaresResponse struct {
	Kind      FirmwareKind
	Firmwares *models.V1AvailableFirmwares
}

// MachineUpdateFirmwareResponse contains the firmware update result
type MachineUpdateFirmwareResponse struct {
	Kind    FirmwareKind
	Machine *models.V1MachineResponse
}

// MachineFirmwareResponse contains the machine firmware result
type MachineFirmwareResponse struct {
	Kind    FirmwareKind
	Machine *models.V1MachineResponse
}

// UploadFirmware uploads the given firmware for the given vendor and board, which is tagged as specified revision
func (d *Driver) UploadFirmware(kind FirmwareKind, vendor, board, revision, updateFile string) (*machine.UploadFirmwareOK, error) {
	uploadFirmware := machine.NewUploadFirmwareParams().WithTimeout(5 * time.Minute)
	uploadFirmware.Kind = string(kind)
	uploadFirmware.Vendor = vendor
	uploadFirmware.Board = board
	uploadFirmware.Revision = revision
	reader, err := os.Open(updateFile)
	if err != nil {
		return nil, err
	}
	uploadFirmware.File = runtime.NamedReader(revision, reader)

	return d.machine.UploadFirmware(uploadFirmware, nil)
}

// RemoveFirmware removes the given firmware revision of the given vendor and board
func (d *Driver) RemoveFirmware(kind FirmwareKind, vendor, board, revision string) (*machine.RemoveFirmwareOK, error) {
	removeFirmware := machine.NewRemoveFirmwareParams().WithTimeout(5 * time.Minute)
	removeFirmware.Kind = string(kind)
	removeFirmware.Vendor = vendor
	removeFirmware.Board = board
	removeFirmware.Revision = revision

	return d.machine.RemoveFirmware(removeFirmware, nil)
}

// AvailableFirmwares returns all available firmwares of given kind that matches given vendor and board (if not empty).
func (d *Driver) AvailableFirmwares(kind FirmwareKind, vendor, board string) (*AvailableFirmwaresResponse, error) {
	return d.availableFirmwares(kind, vendor, board, nil)
}

// MachineAvailableFirmwares returns all available firmwares of given kind for given machine
func (d *Driver) MachineAvailableFirmwares(kind FirmwareKind, machineID string) (*AvailableFirmwaresResponse, error) {
	return d.availableFirmwares(kind, "", "", &machineID)
}

func (d *Driver) availableFirmwares(kind FirmwareKind, vendor, board string, machineID *string) (*AvailableFirmwaresResponse, error) {
	availableFirmwares := machine.NewAvailableFirmwaresParams()
	k := string(kind)
	availableFirmwares.Kind = &k
	availableFirmwares.Vendor = vendor
	availableFirmwares.Board = board
	availableFirmwares.ID = machineID

	response := &AvailableFirmwaresResponse{
		Kind: kind,
	}
	resp, err := d.machine.AvailableFirmwares(availableFirmwares, nil)
	if err != nil {
		return response, err
	}
	response.Firmwares = resp.Payload
	return response, nil
}

// MachineUpdateFirmware updates given firmware of given machine
func (d *Driver) MachineUpdateFirmware(kind FirmwareKind, machineID, revision, description string) (*MachineUpdateFirmwareResponse, error) {
	updateFirmware := machine.NewUpdateFirmwareParams()
	updateFirmware.ID = machineID
	k := string(kind)
	updateFirmware.Body = &models.V1MachineUpdateFirmware{
		Kind:        &k,
		Revision:    &revision,
		Description: &description,
	}

	response := &MachineUpdateFirmwareResponse{
		Kind: kind,
	}
	resp, err := d.machine.UpdateFirmware(updateFirmware, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}
