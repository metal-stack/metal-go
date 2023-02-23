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

// NotifySwitchReader is a Reader for the NotifySwitch structure.
type NotifySwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotifySwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotifySwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNotifySwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNotifySwitchOK creates a NotifySwitchOK with default headers values
func NewNotifySwitchOK() *NotifySwitchOK {
	return &NotifySwitchOK{}
}

/*
NotifySwitchOK describes a response with status code 200, with default header values.

OK
*/
type NotifySwitchOK struct {
	Payload *models.V1SwitchResponse
}

// IsSuccess returns true when this notify switch o k response has a 2xx status code
func (o *NotifySwitchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notify switch o k response has a 3xx status code
func (o *NotifySwitchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify switch o k response has a 4xx status code
func (o *NotifySwitchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notify switch o k response has a 5xx status code
func (o *NotifySwitchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notify switch o k response a status code equal to that given
func (o *NotifySwitchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notify switch o k response
func (o *NotifySwitchOK) Code() int {
	return 200
}

func (o *NotifySwitchOK) Error() string {
	return fmt.Sprintf("[POST /v1/switch/{id}/notify][%d] notifySwitchOK  %+v", 200, o.Payload)
}

func (o *NotifySwitchOK) String() string {
	return fmt.Sprintf("[POST /v1/switch/{id}/notify][%d] notifySwitchOK  %+v", 200, o.Payload)
}

func (o *NotifySwitchOK) GetPayload() *models.V1SwitchResponse {
	return o.Payload
}

func (o *NotifySwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifySwitchDefault creates a NotifySwitchDefault with default headers values
func NewNotifySwitchDefault(code int) *NotifySwitchDefault {
	return &NotifySwitchDefault{
		_statusCode: code,
	}
}

/*
NotifySwitchDefault describes a response with status code -1, with default header values.

Error
*/
type NotifySwitchDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this notify switch default response has a 2xx status code
func (o *NotifySwitchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this notify switch default response has a 3xx status code
func (o *NotifySwitchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this notify switch default response has a 4xx status code
func (o *NotifySwitchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this notify switch default response has a 5xx status code
func (o *NotifySwitchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this notify switch default response a status code equal to that given
func (o *NotifySwitchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the notify switch default response
func (o *NotifySwitchDefault) Code() int {
	return o._statusCode
}

func (o *NotifySwitchDefault) Error() string {
	return fmt.Sprintf("[POST /v1/switch/{id}/notify][%d] notifySwitch default  %+v", o._statusCode, o.Payload)
}

func (o *NotifySwitchDefault) String() string {
	return fmt.Sprintf("[POST /v1/switch/{id}/notify][%d] notifySwitch default  %+v", o._statusCode, o.Payload)
}

func (o *NotifySwitchDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *NotifySwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
