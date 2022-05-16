package metalgo

import (
	"github.com/metal-stack/metal-go/client/filesystemlayout"
	"github.com/metal-stack/metal-go/models"
)

// FilesystemLayoutList return all machine filesystemlayouts
func (d *Driver) FilesystemLayoutList() ([]*models.V1FilesystemLayoutResponse, error) {
	listFilesystemLayouts := filesystemlayout.NewListFilesystemLayoutsParams()
	resp, err := d.Filesystemlayout.ListFilesystemLayouts(listFilesystemLayouts, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutGet return a filesystemlayout
func (d *Driver) FilesystemLayoutGet(filesystemlayoutID string) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewGetFilesystemLayoutParams()
	request.ID = filesystemlayoutID
	resp, err := d.Filesystemlayout.GetFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutCreate create a filesystemlayout
func (d *Driver) FilesystemLayoutCreate(fcr models.V1FilesystemLayoutCreateRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewCreateFilesystemLayoutParams()
	request.SetBody(&fcr)
	resp, err := d.Filesystemlayout.CreateFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutUpdate create a filesystemlayout
func (d *Driver) FilesystemLayoutUpdate(fur models.V1FilesystemLayoutUpdateRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewUpdateFilesystemLayoutParams()
	request.SetBody(&fur)
	resp, err := d.Filesystemlayout.UpdateFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutDelete return a filesystemlayout
func (d *Driver) FilesystemLayoutDelete(filesystemlayoutID string) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewDeleteFilesystemLayoutParams()
	request.ID = filesystemlayoutID
	resp, err := d.Filesystemlayout.DeleteFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutTry with size and image
func (d *Driver) FilesystemLayoutTry(try models.V1FilesystemLayoutTryRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewTryFilesystemLayoutParams()
	request.SetBody(&try)
	resp, err := d.Filesystemlayout.TryFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutMatch with machine and filesystemlayout
func (d *Driver) FilesystemLayoutMatch(match models.V1FilesystemLayoutMatchRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewMatchFilesystemLayoutParams()
	request.SetBody(&match)
	resp, err := d.Filesystemlayout.MatchFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
