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

/* NotifySwitchOK describes a response with status code 200, with default header values.

OK
*/
type NotifySwitchOK struct {
	Payload *models.V1SwitchResponse
}

func (o *NotifySwitchOK) Error() string {
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

/* NotifySwitchDefault describes a response with status code -1, with default header values.

Error
*/
type NotifySwitchDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the notify switch default response
func (o *NotifySwitchDefault) Code() int {
	return o._statusCode
}

func (o *NotifySwitchDefault) Error() string {
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
