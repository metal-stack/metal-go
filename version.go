package metalgo

import (
	"github.com/metal-stack/metal-go/api/models"
)

// VersionGetResponse is the response of a VersionGet action
type VersionGetResponse struct {
	Version *models.RestVersion
}

// VersionGet return a Version
func (d *Driver) VersionGet() (*VersionGetResponse, error) {
	response := &VersionGetResponse{}
	resp, err := d.Client.Version.Info(nil, nil)
	if err != nil {
		return response, err
	}
	response.Version = resp.Payload
	return response, nil
}
