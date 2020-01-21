// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// ReinstallMachineReader is a Reader for the ReinstallMachine structure.
type ReinstallMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReinstallMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReinstallMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 504:
		result := NewReinstallMachineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReinstallMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReinstallMachineOK creates a ReinstallMachineOK with default headers values
func NewReinstallMachineOK() *ReinstallMachineOK {
	return &ReinstallMachineOK{}
}

/*ReinstallMachineOK handles this case with default header values.

OK
*/
type ReinstallMachineOK struct {
	Payload *models.V1MachineResponse
}

func (o *ReinstallMachineOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/reinstall][%d] reinstallMachineOK  %+v", 200, o.Payload)
}

func (o *ReinstallMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReinstallMachineGatewayTimeout creates a ReinstallMachineGatewayTimeout with default headers values
func NewReinstallMachineGatewayTimeout() *ReinstallMachineGatewayTimeout {
	return &ReinstallMachineGatewayTimeout{}
}

/*ReinstallMachineGatewayTimeout handles this case with default header values.

Timeout
*/
type ReinstallMachineGatewayTimeout struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *ReinstallMachineGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/reinstall][%d] reinstallMachineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ReinstallMachineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReinstallMachineDefault creates a ReinstallMachineDefault with default headers values
func NewReinstallMachineDefault(code int) *ReinstallMachineDefault {
	return &ReinstallMachineDefault{
		_statusCode: code,
	}
}

/*ReinstallMachineDefault handles this case with default header values.

Error
*/
type ReinstallMachineDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the reinstall machine default response
func (o *ReinstallMachineDefault) Code() int {
	return o._statusCode
}

func (o *ReinstallMachineDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/reinstall][%d] reinstallMachine default  %+v", o._statusCode, o.Payload)
}

func (o *ReinstallMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
