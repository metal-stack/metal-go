package metalgo

import (
	"os"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/metal-stack/metal-go/client/operations"
	"github.com/metal-stack/metal-go/models"
)

// FirmwaresResponse contains all firmwares matching the requested parameters
type FirmwaresResponse struct {
	Firmwares *models.V1FirmwaresResponse
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
func (d *Driver) UploadFirmware(kind FirmwareKind, vendor, board, revision, file string) (*operations.UploadFirmwareOK, error) {
	uploadFirmware := operations.NewUploadFirmwareParams().WithTimeout(5 * time.Minute)
	uploadFirmware.Kind = string(kind)
	uploadFirmware.Vendor = vendor
	uploadFirmware.Board = board
	uploadFirmware.Revision = revision
	reader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	uploadFirmware.File = runtime.NamedReader(revision, reader)

	return d.Client.UploadFirmware(uploadFirmware, nil)
}

// RemoveFirmware removes the given firmware revision of the given vendor and board
func (d *Driver) RemoveFirmware(kind FirmwareKind, vendor, board, revision string) (*operations.RemoveFirmwareOK, error) {
	removeFirmware := operations.NewRemoveFirmwareParams().WithTimeout(5 * time.Minute)
	removeFirmware.Kind = string(kind)
	removeFirmware.Vendor = vendor
	removeFirmware.Board = board
	removeFirmware.Revision = revision

	return d.Client.RemoveFirmware(removeFirmware, nil)
}

// ListFirmwares returns all firmwares of given kind that matches given vendor and board (if not empty).
func (d *Driver) ListFirmwares(kind FirmwareKind, vendor, board string) (*FirmwaresResponse, error) {
	return d.listFirmwares(kind, vendor, board, nil)
}

// MachineListFirmwares returns all firmwares of given kind for given machine
func (d *Driver) MachineListFirmwares(kind FirmwareKind, machineID string) (*FirmwaresResponse, error) {
	return d.listFirmwares(kind, "", "", &machineID)
}

func (d *Driver) listFirmwares(kind FirmwareKind, vendor, board string, machineID *string) (*FirmwaresResponse, error) {
	availableFirmwares := operations.NewListFirmwaresParams()
	k := string(kind)
	availableFirmwares.Kind = &k
	availableFirmwares.Vendor = &vendor
	availableFirmwares.Board = &board
	availableFirmwares.MachineID = machineID

	response := new(FirmwaresResponse)
	resp, err := d.Client.ListFirmwares(availableFirmwares, nil)
	if err != nil {
		return response, err
	}
	response.Firmwares = resp.Payload
	return response, nil
}

// MachineUpdateFirmware updates given firmware of given machine
func (d *Driver) MachineUpdateFirmware(kind FirmwareKind, machineID, revision, description string) (*MachineUpdateFirmwareResponse, error) {
	updateFirmware := operations.NewUpdateFirmwareParams()
	updateFirmware.ID = machineID
	k := string(kind)
	updateFirmware.Body = &models.V1MachineUpdateFirmwareRequest{
		Kind:        &k,
		Revision:    &revision,
		Description: &description,
	}

	response := &MachineUpdateFirmwareResponse{
		Kind: kind,
	}
	resp, err := d.Client.UpdateFirmware(updateFirmware, nil)
	if err != nil {
		return response, err
	}
	response.Machine = resp.Payload
	return response, nil
}
