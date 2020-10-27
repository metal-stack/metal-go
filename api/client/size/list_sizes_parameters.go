// Code generated by go-swagger; DO NOT EDIT.

package size

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

// NewListSizesParams creates a new ListSizesParams object
// with the default values initialized.
func NewListSizesParams() *ListSizesParams {

	return &ListSizesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSizesParamsWithTimeout creates a new ListSizesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSizesParamsWithTimeout(timeout time.Duration) *ListSizesParams {

	return &ListSizesParams{

		timeout: timeout,
	}
}

// NewListSizesParamsWithContext creates a new ListSizesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSizesParamsWithContext(ctx context.Context) *ListSizesParams {

	return &ListSizesParams{

		Context: ctx,
	}
}

// NewListSizesParamsWithHTTPClient creates a new ListSizesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSizesParamsWithHTTPClient(client *http.Client) *ListSizesParams {

	return &ListSizesParams{
		HTTPClient: client,
	}
}

/*ListSizesParams contains all the parameters to send to the API endpoint
for the list sizes operation typically these are written to a http.Request
*/
type ListSizesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sizes params
func (o *ListSizesParams) WithTimeout(timeout time.Duration) *ListSizesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sizes params
func (o *ListSizesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sizes params
func (o *ListSizesParams) WithContext(ctx context.Context) *ListSizesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sizes params
func (o *ListSizesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sizes params
func (o *ListSizesParams) WithHTTPClient(client *http.Client) *ListSizesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sizes params
func (o *ListSizesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSizesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
