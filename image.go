package metalgo

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/metal-stack/metal-go/client/image"
	"github.com/metal-stack/metal-go/models"
)

// ImageListResponse is the response of a ImageList action
type ImageListResponse struct {
	Image []*models.V1ImageResponse
}

// ImageGetResponse is the response of a ImageList action
type ImageGetResponse struct {
	Image *models.V1ImageResponse
}

// ImageCreateRequest is the response of a ImageList action
type ImageCreateRequest struct {
	ID             string
	Name           string
	Description    string
	URL            string
	Features       []string
	ExpirationDate *time.Time
	Classification *string
}

// ImageCreateResponse is the response of a ImageList action
type ImageCreateResponse struct {
	Image *models.V1ImageResponse
}

// ImageList return all machine images
func (d *Driver) ImageList() (*ImageListResponse, error) {
	response := &ImageListResponse{}
	listImages := image.NewListImagesParams()
	resp, err := d.Image().ListImages(listImages, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageListWithUsage return all machine images with usage info
func (d *Driver) ImageListWithUsage() (*ImageListResponse, error) {
	response := &ImageListResponse{}
	listImages := image.NewListImagesParams()
	showUsage := true
	listImages.WithShowUsage(&showUsage)
	resp, err := d.Image().ListImages(listImages, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageQueryByID return all machine images
func (d *Driver) ImageQueryByID(imageID string) (*ImageListResponse, error) {
	response := &ImageListResponse{}
	request := image.NewQueryImagesByIDParams()
	request.ID = imageID
	resp, err := d.Image().QueryImagesByID(request, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageGet return a image
func (d *Driver) ImageGet(imageID string) (*ImageGetResponse, error) {
	response := &ImageGetResponse{}
	request := image.NewFindImageParams()
	request.ID = imageID
	resp, err := d.Image().FindImage(request, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageGetLatest returns the latest image to a given imageID if no patch version was specified.
func (d *Driver) ImageGetLatest(imageID string) (*ImageGetResponse, error) {
	response := &ImageGetResponse{}
	request := image.NewFindLatestImageParams()
	request.ID = imageID
	resp, err := d.Image().FindLatestImage(request, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageCreate create a image
func (d *Driver) ImageCreate(icr ImageCreateRequest) (*ImageCreateResponse, error) {
	response := &ImageCreateResponse{}

	createImage := &models.V1ImageCreateRequest{
		Description: icr.Description,
		Features:    icr.Features,
		ID:          &icr.ID,
		Name:        icr.Name,
		URL:         &icr.URL,
	}
	request := image.NewCreateImageParams()
	request.SetBody(createImage)
	resp, err := d.Image().CreateImage(request, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageUpdate create a image
func (d *Driver) ImageUpdate(icr ImageCreateRequest) (*ImageCreateResponse, error) {
	response := &ImageCreateResponse{}

	updateImage := &models.V1ImageUpdateRequest{
		Description: icr.Description,
		Features:    icr.Features,
		ID:          &icr.ID,
		Name:        icr.Name,
		URL:         icr.URL,
	}
	if icr.ExpirationDate != nil {
		expt := strfmt.DateTime(*icr.ExpirationDate)
		updateImage.ExpirationDate = &expt
	}
	if icr.Classification != nil {
		updateImage.Classification = *icr.Classification
	}

	request := image.NewUpdateImageParams()
	request.SetBody(updateImage)
	resp, err := d.Image().UpdateImage(request, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}

// ImageDelete return a image
func (d *Driver) ImageDelete(imageID string) (*ImageGetResponse, error) {
	response := &ImageGetResponse{}
	request := image.NewDeleteImageParams()
	request.ID = imageID
	resp, err := d.Image().DeleteImage(request, nil)
	if err != nil {
		return response, err
	}
	response.Image = resp.Payload
	return response, nil
}
