// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// UpdateFirmwareReader is a Reader for the UpdateFirmware structure.
type UpdateFirmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFirmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFirmwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateFirmwareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFirmwareOK creates a UpdateFirmwareOK with default headers values
func NewUpdateFirmwareOK() *UpdateFirmwareOK {
	return &UpdateFirmwareOK{}
}

/*
UpdateFirmwareOK describes a response with status code 200, with default header values.

OK
*/
type UpdateFirmwareOK struct {
	Payload *models.V1MachineResponse
}

// IsSuccess returns true when this update firmware o k response has a 2xx status code
func (o *UpdateFirmwareOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update firmware o k response has a 3xx status code
func (o *UpdateFirmwareOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update firmware o k response has a 4xx status code
func (o *UpdateFirmwareOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update firmware o k response has a 5xx status code
func (o *UpdateFirmwareOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update firmware o k response a status code equal to that given
func (o *UpdateFirmwareOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update firmware o k response
func (o *UpdateFirmwareOK) Code() int {
	return 200
}

func (o *UpdateFirmwareOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/update-firmware/{id}][%d] updateFirmwareOK %s", 200, payload)
}

func (o *UpdateFirmwareOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/update-firmware/{id}][%d] updateFirmwareOK %s", 200, payload)
}

func (o *UpdateFirmwareOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *UpdateFirmwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirmwareDefault creates a UpdateFirmwareDefault with default headers values
func NewUpdateFirmwareDefault(code int) *UpdateFirmwareDefault {
	return &UpdateFirmwareDefault{
		_statusCode: code,
	}
}

/*
UpdateFirmwareDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateFirmwareDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update firmware default response has a 2xx status code
func (o *UpdateFirmwareDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update firmware default response has a 3xx status code
func (o *UpdateFirmwareDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update firmware default response has a 4xx status code
func (o *UpdateFirmwareDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update firmware default response has a 5xx status code
func (o *UpdateFirmwareDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update firmware default response a status code equal to that given
func (o *UpdateFirmwareDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update firmware default response
func (o *UpdateFirmwareDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFirmwareDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/update-firmware/{id}][%d] updateFirmware default %s", o._statusCode, payload)
}

func (o *UpdateFirmwareDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/update-firmware/{id}][%d] updateFirmware default %s", o._statusCode, payload)
}

func (o *UpdateFirmwareDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateFirmwareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
