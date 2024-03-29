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

// NewDeleteFilesystemLayoutParams creates a new DeleteFilesystemLayoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFilesystemLayoutParams() *DeleteFilesystemLayoutParams {
	return &DeleteFilesystemLayoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFilesystemLayoutParamsWithTimeout creates a new DeleteFilesystemLayoutParams object
// with the ability to set a timeout on a request.
func NewDeleteFilesystemLayoutParamsWithTimeout(timeout time.Duration) *DeleteFilesystemLayoutParams {
	return &DeleteFilesystemLayoutParams{
		timeout: timeout,
	}
}

// NewDeleteFilesystemLayoutParamsWithContext creates a new DeleteFilesystemLayoutParams object
// with the ability to set a context for a request.
func NewDeleteFilesystemLayoutParamsWithContext(ctx context.Context) *DeleteFilesystemLayoutParams {
	return &DeleteFilesystemLayoutParams{
		Context: ctx,
	}
}

// NewDeleteFilesystemLayoutParamsWithHTTPClient creates a new DeleteFilesystemLayoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFilesystemLayoutParamsWithHTTPClient(client *http.Client) *DeleteFilesystemLayoutParams {
	return &DeleteFilesystemLayoutParams{
		HTTPClient: client,
	}
}

/*
DeleteFilesystemLayoutParams contains all the parameters to send to the API endpoint

	for the delete filesystem layout operation.

	Typically these are written to a http.Request.
*/
type DeleteFilesystemLayoutParams struct {

	/* ID.

	   identifier of the filesystemlayout
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete filesystem layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFilesystemLayoutParams) WithDefaults() *DeleteFilesystemLayoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete filesystem layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFilesystemLayoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) WithTimeout(timeout time.Duration) *DeleteFilesystemLayoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) WithContext(ctx context.Context) *DeleteFilesystemLayoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) WithHTTPClient(client *http.Client) *DeleteFilesystemLayoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) WithID(id string) *DeleteFilesystemLayoutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete filesystem layout params
func (o *DeleteFilesystemLayoutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFilesystemLayoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
