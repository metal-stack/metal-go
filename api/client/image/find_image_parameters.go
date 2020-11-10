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

// NewFindImageParams creates a new FindImageParams object
// with the default values initialized.
func NewFindImageParams() *FindImageParams {
	var ()
	return &FindImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindImageParamsWithTimeout creates a new FindImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindImageParamsWithTimeout(timeout time.Duration) *FindImageParams {
	var ()
	return &FindImageParams{

		timeout: timeout,
	}
}

// NewFindImageParamsWithContext creates a new FindImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindImageParamsWithContext(ctx context.Context) *FindImageParams {
	var ()
	return &FindImageParams{

		Context: ctx,
	}
}

// NewFindImageParamsWithHTTPClient creates a new FindImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindImageParamsWithHTTPClient(client *http.Client) *FindImageParams {
	var ()
	return &FindImageParams{
		HTTPClient: client,
	}
}

/*FindImageParams contains all the parameters to send to the API endpoint
for the find image operation typically these are written to a http.Request
*/
type FindImageParams struct {

	/*ID
	  identifier of the image

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find image params
func (o *FindImageParams) WithTimeout(timeout time.Duration) *FindImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find image params
func (o *FindImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find image params
func (o *FindImageParams) WithContext(ctx context.Context) *FindImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find image params
func (o *FindImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find image params
func (o *FindImageParams) WithHTTPClient(client *http.Client) *FindImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find image params
func (o *FindImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find image params
func (o *FindImageParams) WithID(id string) *FindImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find image params
func (o *FindImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
