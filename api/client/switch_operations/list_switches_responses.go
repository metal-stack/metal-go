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

// ListSwitchesReader is a Reader for the ListSwitches structure.
type ListSwitchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSwitchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSwitchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSwitchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSwitchesOK creates a ListSwitchesOK with default headers values
func NewListSwitchesOK() *ListSwitchesOK {
	return &ListSwitchesOK{}
}

/*
ListSwitchesOK describes a response with status code 200, with default header values.

OK
*/
type ListSwitchesOK struct {
	Payload []*models.V1SwitchResponse
}

// IsSuccess returns true when this list switches o k response has a 2xx status code
func (o *ListSwitchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list switches o k response has a 3xx status code
func (o *ListSwitchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list switches o k response has a 4xx status code
func (o *ListSwitchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list switches o k response has a 5xx status code
func (o *ListSwitchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list switches o k response a status code equal to that given
func (o *ListSwitchesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListSwitchesOK) Error() string {
	return fmt.Sprintf("[GET /v1/switch][%d] listSwitchesOK  %+v", 200, o.Payload)
}

func (o *ListSwitchesOK) String() string {
	return fmt.Sprintf("[GET /v1/switch][%d] listSwitchesOK  %+v", 200, o.Payload)
}

func (o *ListSwitchesOK) GetPayload() []*models.V1SwitchResponse {
	return o.Payload
}

func (o *ListSwitchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSwitchesDefault creates a ListSwitchesDefault with default headers values
func NewListSwitchesDefault(code int) *ListSwitchesDefault {
	return &ListSwitchesDefault{
		_statusCode: code,
	}
}

/*
ListSwitchesDefault describes a response with status code -1, with default header values.

Error
*/
type ListSwitchesDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the list switches default response
func (o *ListSwitchesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list switches default response has a 2xx status code
func (o *ListSwitchesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list switches default response has a 3xx status code
func (o *ListSwitchesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list switches default response has a 4xx status code
func (o *ListSwitchesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list switches default response has a 5xx status code
func (o *ListSwitchesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list switches default response a status code equal to that given
func (o *ListSwitchesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListSwitchesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/switch][%d] listSwitches default  %+v", o._statusCode, o.Payload)
}

func (o *ListSwitchesDefault) String() string {
	return fmt.Sprintf("[GET /v1/switch][%d] listSwitches default  %+v", o._statusCode, o.Payload)
}

func (o *ListSwitchesDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListSwitchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
