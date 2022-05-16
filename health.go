package metalgo

import (
	"github.com/metal-stack/metal-go/models"
)

// HealthGetResponse is the response of a HealthGet action
type HealthGetResponse struct {
	Health *models.RestHealthResponse
}

// HealthGet returns the health status
func (d *Driver) HealthGet() (*HealthGetResponse, error) {
	response := &HealthGetResponse{}
	resp, err := d.Client.Health(nil, nil)
	if err != nil {
		return response, err
	}
	response.Health = resp.Payload
	return response, nil
}
