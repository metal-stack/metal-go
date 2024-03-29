// Code generated by go-swagger; DO NOT EDIT.

package partition

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

// NewDeletePartitionParams creates a new DeletePartitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePartitionParams() *DeletePartitionParams {
	return &DeletePartitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePartitionParamsWithTimeout creates a new DeletePartitionParams object
// with the ability to set a timeout on a request.
func NewDeletePartitionParamsWithTimeout(timeout time.Duration) *DeletePartitionParams {
	return &DeletePartitionParams{
		timeout: timeout,
	}
}

// NewDeletePartitionParamsWithContext creates a new DeletePartitionParams object
// with the ability to set a context for a request.
func NewDeletePartitionParamsWithContext(ctx context.Context) *DeletePartitionParams {
	return &DeletePartitionParams{
		Context: ctx,
	}
}

// NewDeletePartitionParamsWithHTTPClient creates a new DeletePartitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePartitionParamsWithHTTPClient(client *http.Client) *DeletePartitionParams {
	return &DeletePartitionParams{
		HTTPClient: client,
	}
}

/*
DeletePartitionParams contains all the parameters to send to the API endpoint

	for the delete partition operation.

	Typically these are written to a http.Request.
*/
type DeletePartitionParams struct {

	/* ID.

	   identifier of the Partition
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete partition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePartitionParams) WithDefaults() *DeletePartitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete partition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePartitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete partition params
func (o *DeletePartitionParams) WithTimeout(timeout time.Duration) *DeletePartitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete partition params
func (o *DeletePartitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete partition params
func (o *DeletePartitionParams) WithContext(ctx context.Context) *DeletePartitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete partition params
func (o *DeletePartitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete partition params
func (o *DeletePartitionParams) WithHTTPClient(client *http.Client) *DeletePartitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete partition params
func (o *DeletePartitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete partition params
func (o *DeletePartitionParams) WithID(id string) *DeletePartitionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete partition params
func (o *DeletePartitionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePartitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
