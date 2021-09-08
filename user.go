package metalgo

import (
	"github.com/metal-stack/metal-go/client/operations"
	"github.com/metal-stack/metal-go/models"
)

// Me return the connecting User
func (d *Driver) Me() (*models.V1User, error) {
	params := operations.NewGetMeParams()
	resp, err := d.Client.GetMe(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
