// Code generated by go-swagger; DO NOT EDIT.

package firewall

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

// FindFirewallReader is a Reader for the FindFirewall structure.
type FindFirewallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindFirewallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindFirewallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindFirewallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindFirewallOK creates a FindFirewallOK with default headers values
func NewFindFirewallOK() *FindFirewallOK {
	return &FindFirewallOK{}
}

/*
FindFirewallOK describes a response with status code 200, with default header values.

OK
*/
type FindFirewallOK struct {
	Payload *models.V1FirewallResponse
}

// IsSuccess returns true when this find firewall o k response has a 2xx status code
func (o *FindFirewallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find firewall o k response has a 3xx status code
func (o *FindFirewallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find firewall o k response has a 4xx status code
func (o *FindFirewallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find firewall o k response has a 5xx status code
func (o *FindFirewallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find firewall o k response a status code equal to that given
func (o *FindFirewallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find firewall o k response
func (o *FindFirewallOK) Code() int {
	return 200
}

func (o *FindFirewallOK) Error() string {
	return fmt.Sprintf("[GET /v1/firewall/{id}][%d] findFirewallOK  %+v", 200, o.Payload)
}

func (o *FindFirewallOK) String() string {
	return fmt.Sprintf("[GET /v1/firewall/{id}][%d] findFirewallOK  %+v", 200, o.Payload)
}

func (o *FindFirewallOK) GetPayload() *models.V1FirewallResponse {
	return o.Payload
}

func (o *FindFirewallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1FirewallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindFirewallDefault creates a FindFirewallDefault with default headers values
func NewFindFirewallDefault(code int) *FindFirewallDefault {
	return &FindFirewallDefault{
		_statusCode: code,
	}
}

/*
FindFirewallDefault describes a response with status code -1, with default header values.

Error
*/
type FindFirewallDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this find firewall default response has a 2xx status code
func (o *FindFirewallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find firewall default response has a 3xx status code
func (o *FindFirewallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find firewall default response has a 4xx status code
func (o *FindFirewallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find firewall default response has a 5xx status code
func (o *FindFirewallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find firewall default response a status code equal to that given
func (o *FindFirewallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the find firewall default response
func (o *FindFirewallDefault) Code() int {
	return o._statusCode
}

func (o *FindFirewallDefault) Error() string {
	return fmt.Sprintf("[GET /v1/firewall/{id}][%d] findFirewall default  %+v", o._statusCode, o.Payload)
}

func (o *FindFirewallDefault) String() string {
	return fmt.Sprintf("[GET /v1/firewall/{id}][%d] findFirewall default  %+v", o._statusCode, o.Payload)
}

func (o *FindFirewallDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindFirewallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
