package metalgo

import (
	"github.com/go-openapi/runtime"
	"github.com/metal-stack/metal-go/api/client/machine"
	"os"
	"time"
)

// UploadFirmware uploads the given firmware for the given machine
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
