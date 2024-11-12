// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// MigrateSwitchReader is a Reader for the MigrateSwitch structure.
type MigrateSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMigrateSwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMigrateSwitchOK creates a MigrateSwitchOK with default headers values
func NewMigrateSwitchOK() *MigrateSwitchOK {
	return &MigrateSwitchOK{}
}

/*
MigrateSwitchOK describes a response with status code 200, with default header values.

OK
*/
type MigrateSwitchOK struct {
	Payload *models.V1SwitchResponse
}

// IsSuccess returns true when this migrate switch o k response has a 2xx status code
func (o *MigrateSwitchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this migrate switch o k response has a 3xx status code
func (o *MigrateSwitchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate switch o k response has a 4xx status code
func (o *MigrateSwitchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this migrate switch o k response has a 5xx status code
func (o *MigrateSwitchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate switch o k response a status code equal to that given
func (o *MigrateSwitchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the migrate switch o k response
func (o *MigrateSwitchOK) Code() int {
	return 200
}

func (o *MigrateSwitchOK) Error() string {
	return fmt.Sprintf("[POST /v1/switch/migrate][%d] migrateSwitchOK  %+v", 200, o.Payload)
}

func (o *MigrateSwitchOK) String() string {
	return fmt.Sprintf("[POST /v1/switch/migrate][%d] migrateSwitchOK  %+v", 200, o.Payload)
}

func (o *MigrateSwitchOK) GetPayload() *models.V1SwitchResponse {
	return o.Payload
}

func (o *MigrateSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateSwitchDefault creates a MigrateSwitchDefault with default headers values
func NewMigrateSwitchDefault(code int) *MigrateSwitchDefault {
	return &MigrateSwitchDefault{
		_statusCode: code,
	}
}

/*
MigrateSwitchDefault describes a response with status code -1, with default header values.

Error
*/
type MigrateSwitchDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this migrate switch default response has a 2xx status code
func (o *MigrateSwitchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this migrate switch default response has a 3xx status code
func (o *MigrateSwitchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this migrate switch default response has a 4xx status code
func (o *MigrateSwitchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this migrate switch default response has a 5xx status code
func (o *MigrateSwitchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this migrate switch default response a status code equal to that given
func (o *MigrateSwitchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the migrate switch default response
func (o *MigrateSwitchDefault) Code() int {
	return o._statusCode
}

func (o *MigrateSwitchDefault) Error() string {
	return fmt.Sprintf("[POST /v1/switch/migrate][%d] migrateSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *MigrateSwitchDefault) String() string {
	return fmt.Sprintf("[POST /v1/switch/migrate][%d] migrateSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *MigrateSwitchDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *MigrateSwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
