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

// NewListPartitionsParams creates a new ListPartitionsParams object
// with the default values initialized.
func NewListPartitionsParams() *ListPartitionsParams {

	return &ListPartitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPartitionsParamsWithTimeout creates a new ListPartitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPartitionsParamsWithTimeout(timeout time.Duration) *ListPartitionsParams {

	return &ListPartitionsParams{

		timeout: timeout,
	}
}

// NewListPartitionsParamsWithContext creates a new ListPartitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPartitionsParamsWithContext(ctx context.Context) *ListPartitionsParams {

	return &ListPartitionsParams{

		Context: ctx,
	}
}

// NewListPartitionsParamsWithHTTPClient creates a new ListPartitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPartitionsParamsWithHTTPClient(client *http.Client) *ListPartitionsParams {

	return &ListPartitionsParams{
		HTTPClient: client,
	}
}

/*ListPartitionsParams contains all the parameters to send to the API endpoint
for the list partitions operation typically these are written to a http.Request
*/
type ListPartitionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list partitions params
func (o *ListPartitionsParams) WithTimeout(timeout time.Duration) *ListPartitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list partitions params
func (o *ListPartitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list partitions params
func (o *ListPartitionsParams) WithContext(ctx context.Context) *ListPartitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list partitions params
func (o *ListPartitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list partitions params
func (o *ListPartitionsParams) WithHTTPClient(client *http.Client) *ListPartitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list partitions params
func (o *ListPartitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListPartitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
