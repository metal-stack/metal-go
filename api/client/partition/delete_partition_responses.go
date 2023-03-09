// Code generated by go-swagger; DO NOT EDIT.

package partition

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

// DeletePartitionReader is a Reader for the DeletePartition structure.
type DeletePartitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePartitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePartitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePartitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePartitionOK creates a DeletePartitionOK with default headers values
func NewDeletePartitionOK() *DeletePartitionOK {
	return &DeletePartitionOK{}
}

/*
DeletePartitionOK describes a response with status code 200, with default header values.

OK
*/
type DeletePartitionOK struct {
	Payload *models.V1PartitionResponse
}

// IsSuccess returns true when this delete partition o k response has a 2xx status code
func (o *DeletePartitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete partition o k response has a 3xx status code
func (o *DeletePartitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete partition o k response has a 4xx status code
func (o *DeletePartitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete partition o k response has a 5xx status code
func (o *DeletePartitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete partition o k response a status code equal to that given
func (o *DeletePartitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete partition o k response
func (o *DeletePartitionOK) Code() int {
	return 200
}

func (o *DeletePartitionOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/partition/{id}][%d] deletePartitionOK  %+v", 200, o.Payload)
}

func (o *DeletePartitionOK) String() string {
	return fmt.Sprintf("[DELETE /v1/partition/{id}][%d] deletePartitionOK  %+v", 200, o.Payload)
}

func (o *DeletePartitionOK) GetPayload() *models.V1PartitionResponse {
	return o.Payload
}

func (o *DeletePartitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PartitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePartitionDefault creates a DeletePartitionDefault with default headers values
func NewDeletePartitionDefault(code int) *DeletePartitionDefault {
	return &DeletePartitionDefault{
		_statusCode: code,
	}
}

/*
DeletePartitionDefault describes a response with status code -1, with default header values.

Error
*/
type DeletePartitionDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this delete partition default response has a 2xx status code
func (o *DeletePartitionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete partition default response has a 3xx status code
func (o *DeletePartitionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete partition default response has a 4xx status code
func (o *DeletePartitionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete partition default response has a 5xx status code
func (o *DeletePartitionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete partition default response a status code equal to that given
func (o *DeletePartitionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete partition default response
func (o *DeletePartitionDefault) Code() int {
	return o._statusCode
}

func (o *DeletePartitionDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/partition/{id}][%d] deletePartition default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePartitionDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/partition/{id}][%d] deletePartition default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePartitionDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeletePartitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
