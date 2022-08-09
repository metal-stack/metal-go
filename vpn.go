package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/vpn"
	"github.com/metal-stack/metal-go/api/models"
)

type VPNGetAuthKeyResponse struct {
	VPN *models.V1VPNResponse
}

func (d *Driver) GetVPNAuthKey(projectID string) (*VPNGetAuthKeyResponse, error) {
	response := &VPNGetAuthKeyResponse{}
	getVPNAuthKeyParams := vpn.NewGetVPNAuthKeyParams()
	getVPNAuthKeyParams.Pid = projectID
	resp, err := d.VPN().GetVPNAuthKey(getVPNAuthKeyParams, nil)
	if err != nil {
		return nil, err
	}
	response.VPN = resp.Payload
	return response, nil
}
