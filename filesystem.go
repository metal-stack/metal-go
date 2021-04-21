package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/filesystemlayout"
	"github.com/metal-stack/metal-go/api/models"
)

// FilesystemLayoutList return all machine filesystemlayouts
func (d *Driver) FilesystemLayoutList() ([]*models.V1FilesystemLayoutResponse, error) {
	listFilesystemLayouts := filesystemlayout.NewListFilesystemLayoutsParams()
	resp, err := d.filesystemlayout.ListFilesystemLayouts(listFilesystemLayouts, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutGet return a filesystemlayout
func (d *Driver) FilesystemLayoutGet(filesystemlayoutID string) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewGetFilesystemLayoutParams()
	request.ID = filesystemlayoutID
	resp, err := d.filesystemlayout.GetFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutCreate create a filesystemlayout
func (d *Driver) FilesystemLayoutCreate(fcr models.V1FilesystemLayoutCreateRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewCreateFilesystemLayoutParams()
	request.SetBody(&fcr)
	resp, err := d.filesystemlayout.CreateFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutUpdate create a filesystemlayout
func (d *Driver) FilesystemLayoutUpdate(fur models.V1FilesystemLayoutUpdateRequest) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewUpdateFilesystemLayoutParams()
	request.SetBody(&fur)
	resp, err := d.filesystemlayout.UpdateFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// FilesystemLayoutDelete return a filesystemlayout
func (d *Driver) FilesystemLayoutDelete(filesystemlayoutID string) (*models.V1FilesystemLayoutResponse, error) {
	request := filesystemlayout.NewDeleteFilesystemLayoutParams()
	request.ID = filesystemlayoutID
	resp, err := d.filesystemlayout.DeleteFilesystemLayout(request, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
