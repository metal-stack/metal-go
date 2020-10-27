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

	"github.com/metal-stack/metal-go/api/models"
)

// NewCreateSizeParams creates a new CreateSizeParams object
// with the default values initialized.
func NewCreateSizeParams() *CreateSizeParams {
	var ()
	return &CreateSizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSizeParamsWithTimeout creates a new CreateSizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSizeParamsWithTimeout(timeout time.Duration) *CreateSizeParams {
	var ()
	return &CreateSizeParams{

		timeout: timeout,
	}
}

// NewCreateSizeParamsWithContext creates a new CreateSizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSizeParamsWithContext(ctx context.Context) *CreateSizeParams {
	var ()
	return &CreateSizeParams{

		Context: ctx,
	}
}

// NewCreateSizeParamsWithHTTPClient creates a new CreateSizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSizeParamsWithHTTPClient(client *http.Client) *CreateSizeParams {
	var ()
	return &CreateSizeParams{
		HTTPClient: client,
	}
}

/*CreateSizeParams contains all the parameters to send to the API endpoint
for the create size operation typically these are written to a http.Request
*/
type CreateSizeParams struct {

	/*Body*/
	Body *models.V1SizeCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create size params
func (o *CreateSizeParams) WithTimeout(timeout time.Duration) *CreateSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create size params
func (o *CreateSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create size params
func (o *CreateSizeParams) WithContext(ctx context.Context) *CreateSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create size params
func (o *CreateSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create size params
func (o *CreateSizeParams) WithHTTPClient(client *http.Client) *CreateSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create size params
func (o *CreateSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create size params
func (o *CreateSizeParams) WithBody(body *models.V1SizeCreateRequest) *CreateSizeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create size params
func (o *CreateSizeParams) SetBody(body *models.V1SizeCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
