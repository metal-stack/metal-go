// Code generated by go-swagger; DO NOT EDIT.

package firewall

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

// AllocateFirewallReader is a Reader for the AllocateFirewall structure.
type AllocateFirewallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocateFirewallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllocateFirewallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAllocateFirewallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllocateFirewallOK creates a AllocateFirewallOK with default headers values
func NewAllocateFirewallOK() *AllocateFirewallOK {
	return &AllocateFirewallOK{}
}

/*
AllocateFirewallOK describes a response with status code 200, with default header values.

OK
*/
type AllocateFirewallOK struct {
	Payload *models.V1FirewallResponse
}

// IsSuccess returns true when this allocate firewall o k response has a 2xx status code
func (o *AllocateFirewallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this allocate firewall o k response has a 3xx status code
func (o *AllocateFirewallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allocate firewall o k response has a 4xx status code
func (o *AllocateFirewallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this allocate firewall o k response has a 5xx status code
func (o *AllocateFirewallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this allocate firewall o k response a status code equal to that given
func (o *AllocateFirewallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the allocate firewall o k response
func (o *AllocateFirewallOK) Code() int {
	return 200
}

func (o *AllocateFirewallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/firewall/allocate][%d] allocateFirewallOK %s", 200, payload)
}

func (o *AllocateFirewallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/firewall/allocate][%d] allocateFirewallOK %s", 200, payload)
}

func (o *AllocateFirewallOK) GetPayload() *models.V1FirewallResponse {
	return o.Payload
}

func (o *AllocateFirewallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1FirewallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateFirewallDefault creates a AllocateFirewallDefault with default headers values
func NewAllocateFirewallDefault(code int) *AllocateFirewallDefault {
	return &AllocateFirewallDefault{
		_statusCode: code,
	}
}

/*
AllocateFirewallDefault describes a response with status code -1, with default header values.

Error
*/
type AllocateFirewallDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this allocate firewall default response has a 2xx status code
func (o *AllocateFirewallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this allocate firewall default response has a 3xx status code
func (o *AllocateFirewallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this allocate firewall default response has a 4xx status code
func (o *AllocateFirewallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this allocate firewall default response has a 5xx status code
func (o *AllocateFirewallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this allocate firewall default response a status code equal to that given
func (o *AllocateFirewallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the allocate firewall default response
func (o *AllocateFirewallDefault) Code() int {
	return o._statusCode
}

func (o *AllocateFirewallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/firewall/allocate][%d] allocateFirewall default %s", o._statusCode, payload)
}

func (o *AllocateFirewallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/firewall/allocate][%d] allocateFirewall default %s", o._statusCode, payload)
}

func (o *AllocateFirewallDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AllocateFirewallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
