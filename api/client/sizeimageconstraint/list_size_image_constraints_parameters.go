// Code generated by go-swagger; DO NOT EDIT.

package sizeimageconstraint

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

// NewListSizeImageConstraintsParams creates a new ListSizeImageConstraintsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSizeImageConstraintsParams() *ListSizeImageConstraintsParams {
	return &ListSizeImageConstraintsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSizeImageConstraintsParamsWithTimeout creates a new ListSizeImageConstraintsParams object
// with the ability to set a timeout on a request.
func NewListSizeImageConstraintsParamsWithTimeout(timeout time.Duration) *ListSizeImageConstraintsParams {
	return &ListSizeImageConstraintsParams{
		timeout: timeout,
	}
}

// NewListSizeImageConstraintsParamsWithContext creates a new ListSizeImageConstraintsParams object
// with the ability to set a context for a request.
func NewListSizeImageConstraintsParamsWithContext(ctx context.Context) *ListSizeImageConstraintsParams {
	return &ListSizeImageConstraintsParams{
		Context: ctx,
	}
}

// NewListSizeImageConstraintsParamsWithHTTPClient creates a new ListSizeImageConstraintsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSizeImageConstraintsParamsWithHTTPClient(client *http.Client) *ListSizeImageConstraintsParams {
	return &ListSizeImageConstraintsParams{
		HTTPClient: client,
	}
}

/* ListSizeImageConstraintsParams contains all the parameters to send to the API endpoint
   for the list size image constraints operation.

   Typically these are written to a http.Request.
*/
type ListSizeImageConstraintsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list size image constraints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSizeImageConstraintsParams) WithDefaults() *ListSizeImageConstraintsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list size image constraints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSizeImageConstraintsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list size image constraints params
func (o *ListSizeImageConstraintsParams) WithTimeout(timeout time.Duration) *ListSizeImageConstraintsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list size image constraints params
func (o *ListSizeImageConstraintsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list size image constraints params
func (o *ListSizeImageConstraintsParams) WithContext(ctx context.Context) *ListSizeImageConstraintsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list size image constraints params
func (o *ListSizeImageConstraintsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list size image constraints params
func (o *ListSizeImageConstraintsParams) WithHTTPClient(client *http.Client) *ListSizeImageConstraintsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list size image constraints params
func (o *ListSizeImageConstraintsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSizeImageConstraintsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}