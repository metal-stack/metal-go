package metalgo

import (
	sw "github.com/metal-stack/metal-go/api/client/switch_operations"
	"github.com/metal-stack/metal-go/api/models"
)

// SwitchListResponse is the response of a SwitchList action
type SwitchListResponse struct {
	Switch []*models.V1SwitchResponse
}

// SwitchGetResponse is the response of a SwitchGet action
type SwitchGetResponse struct {
	Switch *models.V1SwitchResponse
}

// SwitchUpdateRequest contains properties to update a switch
type SwitchUpdateRequest struct {
	ID          string
	Name        string
	Description string
	RackID      string
	Mode        string
}

// SwitchList return all switches
func (d *Driver) SwitchList() (*SwitchListResponse, error) {
	response := &SwitchListResponse{}
	listSwitchs := sw.NewListSwitchesParams()
	resp, err := d.sw.ListSwitches(listSwitchs, nil)
	if err != nil {
		return response, err
	}
	response.Switch = resp.Payload
	return response, nil
}

// SwitchGet return a switch
func (d *Driver) SwitchGet(switchID string) (*SwitchGetResponse, error) {
	response := &SwitchGetResponse{}
	get := sw.NewFindSwitchParams()
	get.ID = switchID
	resp, err := d.sw.FindSwitch(get, nil)
	if err != nil {
		return response, err
	}
	response.Switch = resp.Payload
	return response, nil
}

// SwitchUpdate updates a switch
func (d *Driver) SwitchUpdate(sur SwitchUpdateRequest) (*SwitchGetResponse, error) {
	response := &SwitchGetResponse{}
	updateSwitch := &models.V1SwitchUpdateRequest{
		ID:          &sur.ID,
		Name:        sur.Name,
		Description: sur.Description,
		RackID:      &sur.RackID,
		Mode:        sur.Mode,
	}
	request := sw.NewUpdateSwitchParams()
	request.SetBody(updateSwitch)
	resp, err := d.sw.UpdateSwitch(request, nil)
	if err != nil {
		return response, err
	}
	response.Switch = resp.Payload
	return response, nil
}
