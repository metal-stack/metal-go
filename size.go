package metalgo

import (
	"net/http"

	"github.com/metal-stack/metal-go/client/operations"
	"github.com/metal-stack/metal-go/models"
)

// SizeListResponse is the response of a SizeList action
type SizeListResponse struct {
	Size []*models.V1SizeResponse
}

// SizeGetResponse is the response of a SizeGet action
type SizeGetResponse struct {
	Size *models.V1SizeResponse
}

// SizeCreateRequest is the request to create a new Size
type SizeCreateRequest struct {
	ID          string
	Name        string
	Description string
	Constraints []*models.V1SizeConstraint
}

// SizeCreateResponse is the response of a SizeList action
type SizeCreateResponse struct {
	Size *models.V1SizeResponse
}

// SizeTryResponse is the response of a SizeTry action
type SizeTryResponse struct {
	Logs []*models.V1SizeMatchingLog
}

// SizeList return all machine sizes
func (d *Driver) SizeList() (*SizeListResponse, error) {
	response := &SizeListResponse{}
	listSizes := operations.NewListSizesParams()
	resp, err := d.Client.ListSizes(listSizes, nil)
	if err != nil {
		return response, err
	}
	response.Size = resp.Payload
	return response, nil
}

// SizeGet return a size
func (d *Driver) SizeGet(sizeID string) (*SizeGetResponse, error) {
	response := &SizeGetResponse{}
	request := operations.NewFindSizeParams()
	request.ID = sizeID
	resp, err := d.Client.FindSize(request, nil)
	if err != nil {
		return response, err
	}
	response.Size = resp.Payload
	return response, nil
}

// SizeTry will return the chosen size with given Hardware specs.
func (d *Driver) SizeTry(cores int32, memory, storage uint64) (*SizeTryResponse, error) {
	response := &SizeTryResponse{}

	m := int64(memory)
	s := int64(storage)
	diskName := "/dev/trydisk"

	hardware := &models.V1MachineHardwareExtended{
		CPUCores: &cores,
		Memory:   &m,
		Disks: []*models.V1MachineBlockDevice{{
			Name: &diskName,
			Size: &s,
		}},
	}

	trySize := operations.NewFromHardwareParams()
	trySize.Body = hardware

	resp, err := d.Client.FromHardware(trySize, nil)
	if err == nil {
		response.Logs = []*models.V1SizeMatchingLog{resp.Payload}
	} else {
		//nolint:errorlint
		if e, ok := err.(*operations.FromHardwareDefault); ok {
			if e.Code() == http.StatusNotFound {
				response.Logs = []*models.V1SizeMatchingLog{}
				err = nil
			}
		}
	}

	return response, err
}

// SizeCreate create a size
func (d *Driver) SizeCreate(pcr SizeCreateRequest) (*SizeCreateResponse, error) {
	response := &SizeCreateResponse{}

	createSize := &models.V1SizeCreateRequest{
		ID:          &pcr.ID,
		Name:        pcr.Name,
		Description: pcr.Description,
		Constraints: pcr.Constraints,
	}
	request := operations.NewCreateSizeParams()
	request.SetBody(createSize)
	resp, err := d.Client.CreateSize(request, nil)
	if err != nil {
		return response, err
	}
	response.Size = resp.Payload
	return response, nil
}

// SizeUpdate create a size
func (d *Driver) SizeUpdate(pcr SizeCreateRequest) (*SizeCreateResponse, error) {
	response := &SizeCreateResponse{}

	updateSize := &models.V1SizeUpdateRequest{
		ID:          &pcr.ID,
		Name:        pcr.Name,
		Description: pcr.Description,
		Constraints: pcr.Constraints,
	}
	request := operations.NewUpdateSizeParams()
	request.SetBody(updateSize)
	resp, err := d.Client.UpdateSize(request, nil)
	if err != nil {
		return response, err
	}
	response.Size = resp.Payload
	return response, nil
}

// SizeDelete return a size
func (d *Driver) SizeDelete(sizeID string) (*SizeGetResponse, error) {
	response := &SizeGetResponse{}
	request := operations.NewDeleteSizeParams()
	request.ID = sizeID
	resp, err := d.Client.DeleteSize(request, nil)
	if err != nil {
		return response, err
	}
	response.Size = resp.Payload
	return response, nil
}
