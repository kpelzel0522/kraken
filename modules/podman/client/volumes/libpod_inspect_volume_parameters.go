// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewLibpodInspectVolumeParams creates a new LibpodInspectVolumeParams object
// with the default values initialized.
func NewLibpodInspectVolumeParams() *LibpodInspectVolumeParams {
	var ()
	return &LibpodInspectVolumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodInspectVolumeParamsWithTimeout creates a new LibpodInspectVolumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodInspectVolumeParamsWithTimeout(timeout time.Duration) *LibpodInspectVolumeParams {
	var ()
	return &LibpodInspectVolumeParams{

		timeout: timeout,
	}
}

// NewLibpodInspectVolumeParamsWithContext creates a new LibpodInspectVolumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodInspectVolumeParamsWithContext(ctx context.Context) *LibpodInspectVolumeParams {
	var ()
	return &LibpodInspectVolumeParams{

		Context: ctx,
	}
}

// NewLibpodInspectVolumeParamsWithHTTPClient creates a new LibpodInspectVolumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodInspectVolumeParamsWithHTTPClient(client *http.Client) *LibpodInspectVolumeParams {
	var ()
	return &LibpodInspectVolumeParams{
		HTTPClient: client,
	}
}

/*LibpodInspectVolumeParams contains all the parameters to send to the API endpoint
for the libpod inspect volume operation typically these are written to a http.Request
*/
type LibpodInspectVolumeParams struct {

	/*Name
	  the name or ID of the volume

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) WithTimeout(timeout time.Duration) *LibpodInspectVolumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) WithContext(ctx context.Context) *LibpodInspectVolumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) WithHTTPClient(client *http.Client) *LibpodInspectVolumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) WithName(name string) *LibpodInspectVolumeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod inspect volume params
func (o *LibpodInspectVolumeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodInspectVolumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
