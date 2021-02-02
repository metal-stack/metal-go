// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// AvailableBIOSUpdatesReader is a Reader for the AvailableBIOSUpdates structure.
type AvailableBIOSUpdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvailableBIOSUpdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAvailableBIOSUpdatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAvailableBIOSUpdatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAvailableBIOSUpdatesOK creates a AvailableBIOSUpdatesOK with default headers values
func NewAvailableBIOSUpdatesOK() *AvailableBIOSUpdatesOK {
	return &AvailableBIOSUpdatesOK{}
}

/*AvailableBIOSUpdatesOK handles this case with default header values.

OK
*/
type AvailableBIOSUpdatesOK struct {
	Payload *models.V1MachineAvailableUpdates
}

func (o *AvailableBIOSUpdatesOK) Error() string {
	return fmt.Sprintf("[GET /v1/machine/{id}/available-bios-updates][%d] availableBIOSUpdatesOK  %+v", 200, o.Payload)
}

func (o *AvailableBIOSUpdatesOK) GetPayload() *models.V1MachineAvailableUpdates {
	return o.Payload
}

func (o *AvailableBIOSUpdatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineAvailableUpdates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvailableBIOSUpdatesDefault creates a AvailableBIOSUpdatesDefault with default headers values
func NewAvailableBIOSUpdatesDefault(code int) *AvailableBIOSUpdatesDefault {
	return &AvailableBIOSUpdatesDefault{
		_statusCode: code,
	}
}

/*AvailableBIOSUpdatesDefault handles this case with default header values.

Error
*/
type AvailableBIOSUpdatesDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the available b i o s updates default response
func (o *AvailableBIOSUpdatesDefault) Code() int {
	return o._statusCode
}

func (o *AvailableBIOSUpdatesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/machine/{id}/available-bios-updates][%d] availableBIOSUpdates default  %+v", o._statusCode, o.Payload)
}

func (o *AvailableBIOSUpdatesDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AvailableBIOSUpdatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
