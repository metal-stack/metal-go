// Code generated by go-swagger; DO NOT EDIT.

package image

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

// NewDeleteImageParams creates a new DeleteImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteImageParams() *DeleteImageParams {
	return &DeleteImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImageParamsWithTimeout creates a new DeleteImageParams object
// with the ability to set a timeout on a request.
func NewDeleteImageParamsWithTimeout(timeout time.Duration) *DeleteImageParams {
	return &DeleteImageParams{
		timeout: timeout,
	}
}

// NewDeleteImageParamsWithContext creates a new DeleteImageParams object
// with the ability to set a context for a request.
func NewDeleteImageParamsWithContext(ctx context.Context) *DeleteImageParams {
	return &DeleteImageParams{
		Context: ctx,
	}
}

// NewDeleteImageParamsWithHTTPClient creates a new DeleteImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteImageParamsWithHTTPClient(client *http.Client) *DeleteImageParams {
	return &DeleteImageParams{
		HTTPClient: client,
	}
}

/*
DeleteImageParams contains all the parameters to send to the API endpoint

	for the delete image operation.

	Typically these are written to a http.Request.
*/
type DeleteImageParams struct {

	/* ID.

	   identifier of the image
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteImageParams) WithDefaults() *DeleteImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete image params
func (o *DeleteImageParams) WithTimeout(timeout time.Duration) *DeleteImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete image params
func (o *DeleteImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete image params
func (o *DeleteImageParams) WithContext(ctx context.Context) *DeleteImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete image params
func (o *DeleteImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete image params
func (o *DeleteImageParams) WithHTTPClient(client *http.Client) *DeleteImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete image params
func (o *DeleteImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete image params
func (o *DeleteImageParams) WithID(id string) *DeleteImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete image params
func (o *DeleteImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
