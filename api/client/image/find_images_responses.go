// Code generated by go-swagger; DO NOT EDIT.

package image

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

// FindImagesReader is a Reader for the FindImages structure.
type FindImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindImagesOK creates a FindImagesOK with default headers values
func NewFindImagesOK() *FindImagesOK {
	return &FindImagesOK{}
}

/*
FindImagesOK describes a response with status code 200, with default header values.

OK
*/
type FindImagesOK struct {
	Payload []*models.V1ImageResponse
}

// IsSuccess returns true when this find images o k response has a 2xx status code
func (o *FindImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find images o k response has a 3xx status code
func (o *FindImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find images o k response has a 4xx status code
func (o *FindImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find images o k response has a 5xx status code
func (o *FindImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find images o k response a status code equal to that given
func (o *FindImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find images o k response
func (o *FindImagesOK) Code() int {
	return 200
}

func (o *FindImagesOK) Error() string {
	return fmt.Sprintf("[POST /v1/image/find][%d] findImagesOK  %+v", 200, o.Payload)
}

func (o *FindImagesOK) String() string {
	return fmt.Sprintf("[POST /v1/image/find][%d] findImagesOK  %+v", 200, o.Payload)
}

func (o *FindImagesOK) GetPayload() []*models.V1ImageResponse {
	return o.Payload
}

func (o *FindImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindImagesDefault creates a FindImagesDefault with default headers values
func NewFindImagesDefault(code int) *FindImagesDefault {
	return &FindImagesDefault{
		_statusCode: code,
	}
}

/*
FindImagesDefault describes a response with status code -1, with default header values.

Error
*/
type FindImagesDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this find images default response has a 2xx status code
func (o *FindImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find images default response has a 3xx status code
func (o *FindImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find images default response has a 4xx status code
func (o *FindImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find images default response has a 5xx status code
func (o *FindImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find images default response a status code equal to that given
func (o *FindImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the find images default response
func (o *FindImagesDefault) Code() int {
	return o._statusCode
}

func (o *FindImagesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/image/find][%d] findImages default  %+v", o._statusCode, o.Payload)
}

func (o *FindImagesDefault) String() string {
	return fmt.Sprintf("[POST /v1/image/find][%d] findImages default  %+v", o._statusCode, o.Payload)
}

func (o *FindImagesDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
