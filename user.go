package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/user"
	"github.com/metal-stack/metal-go/api/models"
)

// UserGetResponse is the response of a UserGet action
type UserGetResponse struct {
	User *models.SecurityUser
}

// UserGet return a User from given token
func (d *Driver) UserGet(token string) (*UserGetResponse, error) {
	response := &UserGetResponse{}
	params := user.NewGetUserParams()
	params.Token = token

	resp, err := d.client.User.GetUser(params, nil)
	if err != nil {
		return response, err
	}
	response.User = resp.Payload
	return response, nil
}
