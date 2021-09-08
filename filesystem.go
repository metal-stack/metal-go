package metalgo

import (
	"github.com/metal-stack/metal-go/client/operations"
	"github.com/metal-stack/metal-go/models"
)

// FilesystemLayoutList return all machine filesystemlayouts
func (d *Driver) FilesystemLayoutList() ([]*models.V1FilesystemLayoutResponse, error) {
	listFilesystemLayouts := operations.NewListFilesystemLayoutsParams()
	resp, err := d.Client.ListFilesystemLayouts(listFilesystemLayouts, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutGet return a filesystemlayout
func (d *Driver) FilesystemLayoutGet(filesystemlayoutID string) (*models.V1FilesystemLayoutResponse, error) {
	request := operations.NewGetFilesystemLayoutParams()
	request.ID = filesystemlayoutID
	resp, err := d.Client.GetFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutCreate create a filesystemlayout
func (d *Driver) FilesystemLayoutCreate(fcr models.V1FilesystemLayoutCreateRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := operations.NewCreateFilesystemLayoutParams()
	request.SetBody(&fcr)
	resp, err := d.Client.CreateFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutUpdate create a filesystemlayout
func (d *Driver) FilesystemLayoutUpdate(fur models.V1FilesystemLayoutUpdateRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := operations.NewUpdateFilesystemLayoutParams()
	request.SetBody(&fur)
	resp, err := d.Client.UpdateFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutDelete return a filesystemlayout
func (d *Driver) FilesystemLayoutDelete(filesystemlayoutID string) (*models.V1FilesystemLayoutResponse, error) {
	request := operations.NewDeleteFilesystemLayoutParams()
	request.ID = filesystemlayoutID
	resp, err := d.Client.DeleteFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutTry with size and image
func (d *Driver) FilesystemLayoutTry(try models.V1FilesystemLayoutTryRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := operations.NewTryFilesystemLayoutParams()
	request.SetBody(&try)
	resp, err := d.Client.TryFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutMatch with machine and filesystemlayout
func (d *Driver) FilesystemLayoutMatch(match models.V1FilesystemLayoutMatchRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := operations.NewMatchFilesystemLayoutParams()
	request.SetBody(&match)
	resp, err := d.Client.MatchFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
