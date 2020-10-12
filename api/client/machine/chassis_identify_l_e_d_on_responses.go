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
)

// ChassisIdentifyLEDOnReader is a Reader for the ChassisIdentifyLEDOn structure.
type ChassisIdentifyLEDOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChassisIdentifyLEDOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChassisIdentifyLEDOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChassisIdentifyLEDOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChassisIdentifyLEDOnOK creates a ChassisIdentifyLEDOnOK with default headers values
func NewChassisIdentifyLEDOnOK() *ChassisIdentifyLEDOnOK {
	return &ChassisIdentifyLEDOnOK{}
}

/*ChassisIdentifyLEDOnOK handles this case with default header values.

OK
*/
type ChassisIdentifyLEDOnOK struct {
	Payload *models.V1MachineResponse
}

func (o *ChassisIdentifyLEDOnOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/chassis-identify-led-on][%d] chassisIdentifyLEDOnOK  %+v", 200, o.Payload)
}

func (o *ChassisIdentifyLEDOnOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *ChassisIdentifyLEDOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChassisIdentifyLEDOnDefault creates a ChassisIdentifyLEDOnDefault with default headers values
func NewChassisIdentifyLEDOnDefault(code int) *ChassisIdentifyLEDOnDefault {
	return &ChassisIdentifyLEDOnDefault{
		_statusCode: code,
	}
}

/*ChassisIdentifyLEDOnDefault handles this case with default header values.

Error
*/
type ChassisIdentifyLEDOnDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the chassis identify l e d on default response
func (o *ChassisIdentifyLEDOnDefault) Code() int {
	return o._statusCode
}

func (o *ChassisIdentifyLEDOnDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/chassis-identify-led-on][%d] chassisIdentifyLEDOn default  %+v", o._statusCode, o.Payload)
}

func (o *ChassisIdentifyLEDOnDefault) GetPayload() *models.HttperrorsHTTPErrorResponse {
	return o.Payload
}

func (o *ChassisIdentifyLEDOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
