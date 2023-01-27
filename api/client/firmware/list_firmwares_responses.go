// Code generated by go-swagger; DO NOT EDIT.

package firmware

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

// ListFirmwaresReader is a Reader for the ListFirmwares structure.
type ListFirmwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFirmwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFirmwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListFirmwaresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFirmwaresOK creates a ListFirmwaresOK with default headers values
func NewListFirmwaresOK() *ListFirmwaresOK {
	return &ListFirmwaresOK{}
}

/*
ListFirmwaresOK describes a response with status code 200, with default header values.

OK
*/
type ListFirmwaresOK struct {
	Payload *models.V1FirmwaresResponse
}

// IsSuccess returns true when this list firmwares o k response has a 2xx status code
func (o *ListFirmwaresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list firmwares o k response has a 3xx status code
func (o *ListFirmwaresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list firmwares o k response has a 4xx status code
func (o *ListFirmwaresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list firmwares o k response has a 5xx status code
func (o *ListFirmwaresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list firmwares o k response a status code equal to that given
func (o *ListFirmwaresOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list firmwares o k response
func (o *ListFirmwaresOK) Code() int {
	return 200
}

func (o *ListFirmwaresOK) Error() string {
	return fmt.Sprintf("[GET /v1/firmware][%d] listFirmwaresOK  %+v", 200, o.Payload)
}

func (o *ListFirmwaresOK) String() string {
	return fmt.Sprintf("[GET /v1/firmware][%d] listFirmwaresOK  %+v", 200, o.Payload)
}

func (o *ListFirmwaresOK) GetPayload() *models.V1FirmwaresResponse {
	return o.Payload
}

func (o *ListFirmwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1FirmwaresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFirmwaresDefault creates a ListFirmwaresDefault with default headers values
func NewListFirmwaresDefault(code int) *ListFirmwaresDefault {
	return &ListFirmwaresDefault{
		_statusCode: code,
	}
}

/*
ListFirmwaresDefault describes a response with status code -1, with default header values.

Error
*/
type ListFirmwaresDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this list firmwares default response has a 2xx status code
func (o *ListFirmwaresDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list firmwares default response has a 3xx status code
func (o *ListFirmwaresDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list firmwares default response has a 4xx status code
func (o *ListFirmwaresDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list firmwares default response has a 5xx status code
func (o *ListFirmwaresDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list firmwares default response a status code equal to that given
func (o *ListFirmwaresDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list firmwares default response
func (o *ListFirmwaresDefault) Code() int {
	return o._statusCode
}

func (o *ListFirmwaresDefault) Error() string {
	return fmt.Sprintf("[GET /v1/firmware][%d] listFirmwares default  %+v", o._statusCode, o.Payload)
}

func (o *ListFirmwaresDefault) String() string {
	return fmt.Sprintf("[GET /v1/firmware][%d] listFirmwares default  %+v", o._statusCode, o.Payload)
}

func (o *ListFirmwaresDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListFirmwaresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
