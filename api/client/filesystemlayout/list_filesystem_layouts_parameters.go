// Code generated by go-swagger; DO NOT EDIT.

package filesystemlayout

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

// NewListFilesystemLayoutsParams creates a new ListFilesystemLayoutsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListFilesystemLayoutsParams() *ListFilesystemLayoutsParams {
	return &ListFilesystemLayoutsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListFilesystemLayoutsParamsWithTimeout creates a new ListFilesystemLayoutsParams object
// with the ability to set a timeout on a request.
func NewListFilesystemLayoutsParamsWithTimeout(timeout time.Duration) *ListFilesystemLayoutsParams {
	return &ListFilesystemLayoutsParams{
		timeout: timeout,
	}
}

// NewListFilesystemLayoutsParamsWithContext creates a new ListFilesystemLayoutsParams object
// with the ability to set a context for a request.
func NewListFilesystemLayoutsParamsWithContext(ctx context.Context) *ListFilesystemLayoutsParams {
	return &ListFilesystemLayoutsParams{
		Context: ctx,
	}
}

// NewListFilesystemLayoutsParamsWithHTTPClient creates a new ListFilesystemLayoutsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListFilesystemLayoutsParamsWithHTTPClient(client *http.Client) *ListFilesystemLayoutsParams {
	return &ListFilesystemLayoutsParams{
		HTTPClient: client,
	}
}

/*
ListFilesystemLayoutsParams contains all the parameters to send to the API endpoint

	for the list filesystem layouts operation.

	Typically these are written to a http.Request.
*/
type ListFilesystemLayoutsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list filesystem layouts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFilesystemLayoutsParams) WithDefaults() *ListFilesystemLayoutsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list filesystem layouts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFilesystemLayoutsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list filesystem layouts params
func (o *ListFilesystemLayoutsParams) WithTimeout(timeout time.Duration) *ListFilesystemLayoutsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list filesystem layouts params
func (o *ListFilesystemLayoutsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list filesystem layouts params
func (o *ListFilesystemLayoutsParams) WithContext(ctx context.Context) *ListFilesystemLayoutsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list filesystem layouts params
func (o *ListFilesystemLayoutsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list filesystem layouts params
func (o *ListFilesystemLayoutsParams) WithHTTPClient(client *http.Client) *ListFilesystemLayoutsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list filesystem layouts params
func (o *ListFilesystemLayoutsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListFilesystemLayoutsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
