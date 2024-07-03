// Code generated by go-swagger; DO NOT EDIT.

package filesystemlayout

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

// DeleteFilesystemLayoutReader is a Reader for the DeleteFilesystemLayout structure.
type DeleteFilesystemLayoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFilesystemLayoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFilesystemLayoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFilesystemLayoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFilesystemLayoutOK creates a DeleteFilesystemLayoutOK with default headers values
func NewDeleteFilesystemLayoutOK() *DeleteFilesystemLayoutOK {
	return &DeleteFilesystemLayoutOK{}
}

/*
DeleteFilesystemLayoutOK describes a response with status code 200, with default header values.

OK
*/
type DeleteFilesystemLayoutOK struct {
	Payload *models.V1FilesystemLayoutResponse
}

// IsSuccess returns true when this delete filesystem layout o k response has a 2xx status code
func (o *DeleteFilesystemLayoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete filesystem layout o k response has a 3xx status code
func (o *DeleteFilesystemLayoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete filesystem layout o k response has a 4xx status code
func (o *DeleteFilesystemLayoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete filesystem layout o k response has a 5xx status code
func (o *DeleteFilesystemLayoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete filesystem layout o k response a status code equal to that given
func (o *DeleteFilesystemLayoutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete filesystem layout o k response
func (o *DeleteFilesystemLayoutOK) Code() int {
	return 200
}

func (o *DeleteFilesystemLayoutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/filesystemlayout/{id}][%d] deleteFilesystemLayoutOK %s", 200, payload)
}

func (o *DeleteFilesystemLayoutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/filesystemlayout/{id}][%d] deleteFilesystemLayoutOK %s", 200, payload)
}

func (o *DeleteFilesystemLayoutOK) GetPayload() *models.V1FilesystemLayoutResponse {
	return o.Payload
}

func (o *DeleteFilesystemLayoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1FilesystemLayoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFilesystemLayoutDefault creates a DeleteFilesystemLayoutDefault with default headers values
func NewDeleteFilesystemLayoutDefault(code int) *DeleteFilesystemLayoutDefault {
	return &DeleteFilesystemLayoutDefault{
		_statusCode: code,
	}
}

/*
DeleteFilesystemLayoutDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteFilesystemLayoutDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this delete filesystem layout default response has a 2xx status code
func (o *DeleteFilesystemLayoutDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete filesystem layout default response has a 3xx status code
func (o *DeleteFilesystemLayoutDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete filesystem layout default response has a 4xx status code
func (o *DeleteFilesystemLayoutDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete filesystem layout default response has a 5xx status code
func (o *DeleteFilesystemLayoutDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete filesystem layout default response a status code equal to that given
func (o *DeleteFilesystemLayoutDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete filesystem layout default response
func (o *DeleteFilesystemLayoutDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFilesystemLayoutDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/filesystemlayout/{id}][%d] deleteFilesystemLayout default %s", o._statusCode, payload)
}

func (o *DeleteFilesystemLayoutDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/filesystemlayout/{id}][%d] deleteFilesystemLayout default %s", o._statusCode, payload)
}

func (o *DeleteFilesystemLayoutDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteFilesystemLayoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
