// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewFreeNetworkParams creates a new FreeNetworkParams object
// with the default values initialized.
func NewFreeNetworkParams() *FreeNetworkParams {
	var ()
	return &FreeNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFreeNetworkParamsWithTimeout creates a new FreeNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFreeNetworkParamsWithTimeout(timeout time.Duration) *FreeNetworkParams {
	var ()
	return &FreeNetworkParams{

		timeout: timeout,
	}
}

// NewFreeNetworkParamsWithContext creates a new FreeNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewFreeNetworkParamsWithContext(ctx context.Context) *FreeNetworkParams {
	var ()
	return &FreeNetworkParams{

		Context: ctx,
	}
}

// NewFreeNetworkParamsWithHTTPClient creates a new FreeNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFreeNetworkParamsWithHTTPClient(client *http.Client) *FreeNetworkParams {
	var ()
	return &FreeNetworkParams{
		HTTPClient: client,
	}
}

/*FreeNetworkParams contains all the parameters to send to the API endpoint
for the free network operation typically these are written to a http.Request
*/
type FreeNetworkParams struct {

	/*ID
	  identifier of the network

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the free network params
func (o *FreeNetworkParams) WithTimeout(timeout time.Duration) *FreeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the free network params
func (o *FreeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the free network params
func (o *FreeNetworkParams) WithContext(ctx context.Context) *FreeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the free network params
func (o *FreeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the free network params
func (o *FreeNetworkParams) WithHTTPClient(client *http.Client) *FreeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the free network params
func (o *FreeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the free network params
func (o *FreeNetworkParams) WithID(id string) *FreeNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the free network params
func (o *FreeNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FreeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
