package metalgo

import (
	"os"
	"time"

	"github.com/go-openapi/runtime"

	"github.com/metal-stack/metal-go/api/client/firmware"
	"github.com/metal-stack/metal-go/api/models"
)

// FirmwaresResponse contains all firmwares matching the requested parameters
type FirmwaresResponse struct {
	Firmwares *models.V1FirmwaresResponse
}

// UploadFirmware uploads the given firmware for the given vendor and board, which is tagged as specified revision
func (d *Driver) UploadFirmware(kind FirmwareKind, vendor, board, revision, file string) (*firmware.UploadFirmwareOK, error) {
	uploadFirmware := firmware.NewUploadFirmwareParams().WithTimeout(5 * time.Minute)
	uploadFirmware.Kind = string(kind)
	uploadFirmware.Vendor = vendor
	uploadFirmware.Board = board
	uploadFirmware.Revision = revision
	reader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	uploadFirmware.File = runtime.NamedReader(revision, reader)

	return d.Firmware().UploadFirmware(uploadFirmware, nil)
}

// RemoveFirmware removes the given firmware revision of the given vendor and board
func (d *Driver) RemoveFirmware(kind FirmwareKind, vendor, board, revision string) (*firmware.RemoveFirmwareOK, error) {
	removeFirmware := firmware.NewRemoveFirmwareParams().WithTimeout(5 * time.Minute)
	removeFirmware.Kind = string(kind)
	removeFirmware.Vendor = vendor
	removeFirmware.Board = board
	removeFirmware.Revision = revision

	return d.Firmware().RemoveFirmware(removeFirmware, nil)
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
	availableFirmwares := firmware.NewListFirmwaresParams()
	k := string(kind)
	availableFirmwares.Kind = &k
	availableFirmwares.Vendor = &vendor
	availableFirmwares.Board = &board
	availableFirmwares.MachineID = machineID

	response := new(FirmwaresResponse)
	resp, err := d.Firmware().ListFirmwares(availableFirmwares, nil)
	if err != nil {
		return response, err
	}
	response.Firmwares = resp.Payload
	return response, nil
}
