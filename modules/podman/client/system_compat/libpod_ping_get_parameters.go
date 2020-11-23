// Code generated by go-swagger; DO NOT EDIT.

package system_compat

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

// NewLibpodPingGetParams creates a new LibpodPingGetParams object
// with the default values initialized.
func NewLibpodPingGetParams() *LibpodPingGetParams {

	return &LibpodPingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodPingGetParamsWithTimeout creates a new LibpodPingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodPingGetParamsWithTimeout(timeout time.Duration) *LibpodPingGetParams {

	return &LibpodPingGetParams{

		timeout: timeout,
	}
}

// NewLibpodPingGetParamsWithContext creates a new LibpodPingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodPingGetParamsWithContext(ctx context.Context) *LibpodPingGetParams {

	return &LibpodPingGetParams{

		Context: ctx,
	}
}

// NewLibpodPingGetParamsWithHTTPClient creates a new LibpodPingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodPingGetParamsWithHTTPClient(client *http.Client) *LibpodPingGetParams {

	return &LibpodPingGetParams{
		HTTPClient: client,
	}
}

/*LibpodPingGetParams contains all the parameters to send to the API endpoint
for the libpod ping get operation typically these are written to a http.Request
*/
type LibpodPingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod ping get params
func (o *LibpodPingGetParams) WithTimeout(timeout time.Duration) *LibpodPingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod ping get params
func (o *LibpodPingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod ping get params
func (o *LibpodPingGetParams) WithContext(ctx context.Context) *LibpodPingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod ping get params
func (o *LibpodPingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod ping get params
func (o *LibpodPingGetParams) WithHTTPClient(client *http.Client) *LibpodPingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod ping get params
func (o *LibpodPingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodPingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
