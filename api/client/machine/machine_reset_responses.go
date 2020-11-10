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

// MachineResetReader is a Reader for the MachineReset structure.
type MachineResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachineResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMachineResetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMachineResetOK creates a MachineResetOK with default headers values
func NewMachineResetOK() *MachineResetOK {
	return &MachineResetOK{}
}

/*MachineResetOK handles this case with default header values.

OK
*/
type MachineResetOK struct {
	Payload *models.V1MachineResponse
}

func (o *MachineResetOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/reset][%d] machineResetOK  %+v", 200, o.Payload)
}

func (o *MachineResetOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *MachineResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachineResetDefault creates a MachineResetDefault with default headers values
func NewMachineResetDefault(code int) *MachineResetDefault {
	return &MachineResetDefault{
		_statusCode: code,
	}
}

/*MachineResetDefault handles this case with default header values.

Error
*/
type MachineResetDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the machine reset default response
func (o *MachineResetDefault) Code() int {
	return o._statusCode
}

func (o *MachineResetDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/reset][%d] machineReset default  %+v", o._statusCode, o.Payload)
}

func (o *MachineResetDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *MachineResetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
