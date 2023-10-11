// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// FindTenantsReader is a Reader for the FindTenants structure.
type FindTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindTenantsOK creates a FindTenantsOK with default headers values
func NewFindTenantsOK() *FindTenantsOK {
	return &FindTenantsOK{}
}

/*
FindTenantsOK describes a response with status code 200, with default header values.

OK
*/
type FindTenantsOK struct {
	Payload []*models.V1TenantResponse
}

// IsSuccess returns true when this find tenants o k response has a 2xx status code
func (o *FindTenantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find tenants o k response has a 3xx status code
func (o *FindTenantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find tenants o k response has a 4xx status code
func (o *FindTenantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find tenants o k response has a 5xx status code
func (o *FindTenantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find tenants o k response a status code equal to that given
func (o *FindTenantsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find tenants o k response
func (o *FindTenantsOK) Code() int {
	return 200
}

func (o *FindTenantsOK) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/find][%d] findTenantsOK  %+v", 200, o.Payload)
}

func (o *FindTenantsOK) String() string {
	return fmt.Sprintf("[POST /v1/tenant/find][%d] findTenantsOK  %+v", 200, o.Payload)
}

func (o *FindTenantsOK) GetPayload() []*models.V1TenantResponse {
	return o.Payload
}

func (o *FindTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindTenantsDefault creates a FindTenantsDefault with default headers values
func NewFindTenantsDefault(code int) *FindTenantsDefault {
	return &FindTenantsDefault{
		_statusCode: code,
	}
}

/*
FindTenantsDefault describes a response with status code -1, with default header values.

Error
*/
type FindTenantsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this find tenants default response has a 2xx status code
func (o *FindTenantsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find tenants default response has a 3xx status code
func (o *FindTenantsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find tenants default response has a 4xx status code
func (o *FindTenantsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find tenants default response has a 5xx status code
func (o *FindTenantsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find tenants default response a status code equal to that given
func (o *FindTenantsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the find tenants default response
func (o *FindTenantsDefault) Code() int {
	return o._statusCode
}

func (o *FindTenantsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/find][%d] findTenants default  %+v", o._statusCode, o.Payload)
}

func (o *FindTenantsDefault) String() string {
	return fmt.Sprintf("[POST /v1/tenant/find][%d] findTenants default  %+v", o._statusCode, o.Payload)
}

func (o *FindTenantsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}