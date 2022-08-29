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

// NewDeleteSizeImageConstraintParams creates a new DeleteSizeImageConstraintParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSizeImageConstraintParams() *DeleteSizeImageConstraintParams {
	return &DeleteSizeImageConstraintParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSizeImageConstraintParamsWithTimeout creates a new DeleteSizeImageConstraintParams object
// with the ability to set a timeout on a request.
func NewDeleteSizeImageConstraintParamsWithTimeout(timeout time.Duration) *DeleteSizeImageConstraintParams {
	return &DeleteSizeImageConstraintParams{
		timeout: timeout,
	}
}

// NewDeleteSizeImageConstraintParamsWithContext creates a new DeleteSizeImageConstraintParams object
// with the ability to set a context for a request.
func NewDeleteSizeImageConstraintParamsWithContext(ctx context.Context) *DeleteSizeImageConstraintParams {
	return &DeleteSizeImageConstraintParams{
		Context: ctx,
	}
}

// NewDeleteSizeImageConstraintParamsWithHTTPClient creates a new DeleteSizeImageConstraintParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSizeImageConstraintParamsWithHTTPClient(client *http.Client) *DeleteSizeImageConstraintParams {
	return &DeleteSizeImageConstraintParams{
		HTTPClient: client,
	}
}

/*
DeleteSizeImageConstraintParams contains all the parameters to send to the API endpoint

	for the delete size image constraint operation.

	Typically these are written to a http.Request.
*/
type DeleteSizeImageConstraintParams struct {

	/* ID.

	   identifier of the size
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete size image constraint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSizeImageConstraintParams) WithDefaults() *DeleteSizeImageConstraintParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete size image constraint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSizeImageConstraintParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) WithTimeout(timeout time.Duration) *DeleteSizeImageConstraintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) WithContext(ctx context.Context) *DeleteSizeImageConstraintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) WithHTTPClient(client *http.Client) *DeleteSizeImageConstraintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) WithID(id string) *DeleteSizeImageConstraintParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete size image constraint params
func (o *DeleteSizeImageConstraintParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSizeImageConstraintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
