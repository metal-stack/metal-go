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
)

// NewFindIPMIMachineParams creates a new FindIPMIMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindIPMIMachineParams() *FindIPMIMachineParams {
	return &FindIPMIMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindIPMIMachineParamsWithTimeout creates a new FindIPMIMachineParams object
// with the ability to set a timeout on a request.
func NewFindIPMIMachineParamsWithTimeout(timeout time.Duration) *FindIPMIMachineParams {
	return &FindIPMIMachineParams{
		timeout: timeout,
	}
}

// NewFindIPMIMachineParamsWithContext creates a new FindIPMIMachineParams object
// with the ability to set a context for a request.
func NewFindIPMIMachineParamsWithContext(ctx context.Context) *FindIPMIMachineParams {
	return &FindIPMIMachineParams{
		Context: ctx,
	}
}

// NewFindIPMIMachineParamsWithHTTPClient creates a new FindIPMIMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindIPMIMachineParamsWithHTTPClient(client *http.Client) *FindIPMIMachineParams {
	return &FindIPMIMachineParams{
		HTTPClient: client,
	}
}

/*
FindIPMIMachineParams contains all the parameters to send to the API endpoint

	for the find IP m i machine operation.

	Typically these are written to a http.Request.
*/
type FindIPMIMachineParams struct {

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find IP m i machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindIPMIMachineParams) WithDefaults() *FindIPMIMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find IP m i machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindIPMIMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find IP m i machine params
func (o *FindIPMIMachineParams) WithTimeout(timeout time.Duration) *FindIPMIMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find IP m i machine params
func (o *FindIPMIMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find IP m i machine params
func (o *FindIPMIMachineParams) WithContext(ctx context.Context) *FindIPMIMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find IP m i machine params
func (o *FindIPMIMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find IP m i machine params
func (o *FindIPMIMachineParams) WithHTTPClient(client *http.Client) *FindIPMIMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find IP m i machine params
func (o *FindIPMIMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find IP m i machine params
func (o *FindIPMIMachineParams) WithID(id string) *FindIPMIMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find IP m i machine params
func (o *FindIPMIMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindIPMIMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
