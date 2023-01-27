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

// ListImagesReader is a Reader for the ListImages structure.
type ListImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListImagesOK creates a ListImagesOK with default headers values
func NewListImagesOK() *ListImagesOK {
	return &ListImagesOK{}
}

/*
ListImagesOK describes a response with status code 200, with default header values.

OK
*/
type ListImagesOK struct {
	Payload []*models.V1ImageResponse
}

// IsSuccess returns true when this list images o k response has a 2xx status code
func (o *ListImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list images o k response has a 3xx status code
func (o *ListImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list images o k response has a 4xx status code
func (o *ListImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list images o k response has a 5xx status code
func (o *ListImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list images o k response a status code equal to that given
func (o *ListImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list images o k response
func (o *ListImagesOK) Code() int {
	return 200
}

func (o *ListImagesOK) Error() string {
	return fmt.Sprintf("[GET /v1/image][%d] listImagesOK  %+v", 200, o.Payload)
}

func (o *ListImagesOK) String() string {
	return fmt.Sprintf("[GET /v1/image][%d] listImagesOK  %+v", 200, o.Payload)
}

func (o *ListImagesOK) GetPayload() []*models.V1ImageResponse {
	return o.Payload
}

func (o *ListImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesDefault creates a ListImagesDefault with default headers values
func NewListImagesDefault(code int) *ListImagesDefault {
	return &ListImagesDefault{
		_statusCode: code,
	}
}

/*
ListImagesDefault describes a response with status code -1, with default header values.

Error
*/
type ListImagesDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this list images default response has a 2xx status code
func (o *ListImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list images default response has a 3xx status code
func (o *ListImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list images default response has a 4xx status code
func (o *ListImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list images default response has a 5xx status code
func (o *ListImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list images default response a status code equal to that given
func (o *ListImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list images default response
func (o *ListImagesDefault) Code() int {
	return o._statusCode
}

func (o *ListImagesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/image][%d] listImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListImagesDefault) String() string {
	return fmt.Sprintf("[GET /v1/image][%d] listImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListImagesDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
