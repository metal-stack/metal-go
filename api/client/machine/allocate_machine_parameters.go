// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
)

// NewAllocateMachineParams creates a new AllocateMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllocateMachineParams() *AllocateMachineParams {
	return &AllocateMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllocateMachineParamsWithTimeout creates a new AllocateMachineParams object
// with the ability to set a timeout on a request.
func NewAllocateMachineParamsWithTimeout(timeout time.Duration) *AllocateMachineParams {
	return &AllocateMachineParams{
		timeout: timeout,
	}
}

// NewAllocateMachineParamsWithContext creates a new AllocateMachineParams object
// with the ability to set a context for a request.
func NewAllocateMachineParamsWithContext(ctx context.Context) *AllocateMachineParams {
	return &AllocateMachineParams{
		Context: ctx,
	}
}

// NewAllocateMachineParamsWithHTTPClient creates a new AllocateMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllocateMachineParamsWithHTTPClient(client *http.Client) *AllocateMachineParams {
	return &AllocateMachineParams{
		HTTPClient: client,
	}
}

/*
AllocateMachineParams contains all the parameters to send to the API endpoint

	for the allocate machine operation.

	Typically these are written to a http.Request.
*/
type AllocateMachineParams struct {

	// Body.
	Body *models.V1MachineAllocateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allocate machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocateMachineParams) WithDefaults() *AllocateMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allocate machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocateMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the allocate machine params
func (o *AllocateMachineParams) WithTimeout(timeout time.Duration) *AllocateMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocate machine params
func (o *AllocateMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocate machine params
func (o *AllocateMachineParams) WithContext(ctx context.Context) *AllocateMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocate machine params
func (o *AllocateMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocate machine params
func (o *AllocateMachineParams) WithHTTPClient(client *http.Client) *AllocateMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocate machine params
func (o *AllocateMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the allocate machine params
func (o *AllocateMachineParams) WithBody(body *models.V1MachineAllocateRequest) *AllocateMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the allocate machine params
func (o *AllocateMachineParams) SetBody(body *models.V1MachineAllocateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AllocateMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
