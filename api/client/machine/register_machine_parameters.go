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

// NewRegisterMachineParams creates a new RegisterMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterMachineParams() *RegisterMachineParams {
	return &RegisterMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterMachineParamsWithTimeout creates a new RegisterMachineParams object
// with the ability to set a timeout on a request.
func NewRegisterMachineParamsWithTimeout(timeout time.Duration) *RegisterMachineParams {
	return &RegisterMachineParams{
		timeout: timeout,
	}
}

// NewRegisterMachineParamsWithContext creates a new RegisterMachineParams object
// with the ability to set a context for a request.
func NewRegisterMachineParamsWithContext(ctx context.Context) *RegisterMachineParams {
	return &RegisterMachineParams{
		Context: ctx,
	}
}

// NewRegisterMachineParamsWithHTTPClient creates a new RegisterMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterMachineParamsWithHTTPClient(client *http.Client) *RegisterMachineParams {
	return &RegisterMachineParams{
		HTTPClient: client,
	}
}

/*
RegisterMachineParams contains all the parameters to send to the API endpoint

	for the register machine operation.

	Typically these are written to a http.Request.
*/
type RegisterMachineParams struct {

	// Body.
	Body *models.V1MachineRegisterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterMachineParams) WithDefaults() *RegisterMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register machine params
func (o *RegisterMachineParams) WithTimeout(timeout time.Duration) *RegisterMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register machine params
func (o *RegisterMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register machine params
func (o *RegisterMachineParams) WithContext(ctx context.Context) *RegisterMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register machine params
func (o *RegisterMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register machine params
func (o *RegisterMachineParams) WithHTTPClient(client *http.Client) *RegisterMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register machine params
func (o *RegisterMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register machine params
func (o *RegisterMachineParams) WithBody(body *models.V1MachineRegisterRequest) *RegisterMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register machine params
func (o *RegisterMachineParams) SetBody(body *models.V1MachineRegisterRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
