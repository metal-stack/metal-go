// Code generated by go-swagger; DO NOT EDIT.

package version

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

// InfoReader is a Reader for the Info structure.
type InfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInfoOK creates a InfoOK with default headers values
func NewInfoOK() *InfoOK {
	return &InfoOK{}
}

/*
InfoOK describes a response with status code 200, with default header values.

OK
*/
type InfoOK struct {
	Payload *models.RestVersion
}

// IsSuccess returns true when this info o k response has a 2xx status code
func (o *InfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this info o k response has a 3xx status code
func (o *InfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this info o k response has a 4xx status code
func (o *InfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this info o k response has a 5xx status code
func (o *InfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this info o k response a status code equal to that given
func (o *InfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *InfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/version][%d] infoOK  %+v", 200, o.Payload)
}

func (o *InfoOK) String() string {
	return fmt.Sprintf("[GET /v1/version][%d] infoOK  %+v", 200, o.Payload)
}

func (o *InfoOK) GetPayload() *models.RestVersion {
	return o.Payload
}

func (o *InfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInfoDefault creates a InfoDefault with default headers values
func NewInfoDefault(code int) *InfoDefault {
	return &InfoDefault{
		_statusCode: code,
	}
}

/*
InfoDefault describes a response with status code -1, with default header values.

Error
*/
type InfoDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the info default response
func (o *InfoDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this info default response has a 2xx status code
func (o *InfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this info default response has a 3xx status code
func (o *InfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this info default response has a 4xx status code
func (o *InfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this info default response has a 5xx status code
func (o *InfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this info default response a status code equal to that given
func (o *InfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *InfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1/version][%d] info default  %+v", o._statusCode, o.Payload)
}

func (o *InfoDefault) String() string {
	return fmt.Sprintf("[GET /v1/version][%d] info default  %+v", o._statusCode, o.Payload)
}

func (o *InfoDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *InfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
