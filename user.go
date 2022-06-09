package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/user"
	"github.com/metal-stack/metal-go/api/models"
)

// Me return the connecting User
func (d *Driver) Me() (*models.V1User, error) {
	params := user.NewGetMeParams()
	resp, err := d.User().GetMe(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
