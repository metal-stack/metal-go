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

// GetMachineConsolePasswordReader is a Reader for the GetMachineConsolePassword structure.
type GetMachineConsolePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineConsolePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineConsolePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMachineConsolePasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachineConsolePasswordOK creates a GetMachineConsolePasswordOK with default headers values
func NewGetMachineConsolePasswordOK() *GetMachineConsolePasswordOK {
	return &GetMachineConsolePasswordOK{}
}

/*
GetMachineConsolePasswordOK describes a response with status code 200, with default header values.

OK
*/
type GetMachineConsolePasswordOK struct {
	Payload *models.V1MachineConsolePasswordResponse
}

// IsSuccess returns true when this get machine console password o k response has a 2xx status code
func (o *GetMachineConsolePasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get machine console password o k response has a 3xx status code
func (o *GetMachineConsolePasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine console password o k response has a 4xx status code
func (o *GetMachineConsolePasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get machine console password o k response has a 5xx status code
func (o *GetMachineConsolePasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine console password o k response a status code equal to that given
func (o *GetMachineConsolePasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get machine console password o k response
func (o *GetMachineConsolePasswordOK) Code() int {
	return 200
}

func (o *GetMachineConsolePasswordOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/machine/consolepassword][%d] getMachineConsolePasswordOK %s", 200, payload)
}

func (o *GetMachineConsolePasswordOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/machine/consolepassword][%d] getMachineConsolePasswordOK %s", 200, payload)
}

func (o *GetMachineConsolePasswordOK) GetPayload() *models.V1MachineConsolePasswordResponse {
	return o.Payload
}

func (o *GetMachineConsolePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineConsolePasswordResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineConsolePasswordDefault creates a GetMachineConsolePasswordDefault with default headers values
func NewGetMachineConsolePasswordDefault(code int) *GetMachineConsolePasswordDefault {
	return &GetMachineConsolePasswordDefault{
		_statusCode: code,
	}
}

/*
GetMachineConsolePasswordDefault describes a response with status code -1, with default header values.

Error
*/
type GetMachineConsolePasswordDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get machine console password default response has a 2xx status code
func (o *GetMachineConsolePasswordDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get machine console password default response has a 3xx status code
func (o *GetMachineConsolePasswordDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get machine console password default response has a 4xx status code
func (o *GetMachineConsolePasswordDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get machine console password default response has a 5xx status code
func (o *GetMachineConsolePasswordDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get machine console password default response a status code equal to that given
func (o *GetMachineConsolePasswordDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get machine console password default response
func (o *GetMachineConsolePasswordDefault) Code() int {
	return o._statusCode
}

func (o *GetMachineConsolePasswordDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/machine/consolepassword][%d] getMachineConsolePassword default %s", o._statusCode, payload)
}

func (o *GetMachineConsolePasswordDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/machine/consolepassword][%d] getMachineConsolePassword default %s", o._statusCode, payload)
}

func (o *GetMachineConsolePasswordDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetMachineConsolePasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
