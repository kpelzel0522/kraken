// Code generated by go-swagger; DO NOT EDIT.

package manifests

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
	"github.com/go-openapi/swag"
)

// NewCreateParams creates a new CreateParams object
// with the default values initialized.
func NewCreateParams() *CreateParams {
	var ()
	return &CreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateParamsWithTimeout creates a new CreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateParamsWithTimeout(timeout time.Duration) *CreateParams {
	var ()
	return &CreateParams{

		timeout: timeout,
	}
}

// NewCreateParamsWithContext creates a new CreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateParamsWithContext(ctx context.Context) *CreateParams {
	var ()
	return &CreateParams{

		Context: ctx,
	}
}

// NewCreateParamsWithHTTPClient creates a new CreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateParamsWithHTTPClient(client *http.Client) *CreateParams {
	var ()
	return &CreateParams{
		HTTPClient: client,
	}
}

/*CreateParams contains all the parameters to send to the API endpoint
for the create operation typically these are written to a http.Request
*/
type CreateParams struct {

	/*All
	  add all contents if given list

	*/
	All *bool
	/*Image
	  name of the image

	*/
	Image *string
	/*Name
	  manifest list name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create params
func (o *CreateParams) WithTimeout(timeout time.Duration) *CreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create params
func (o *CreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create params
func (o *CreateParams) WithContext(ctx context.Context) *CreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create params
func (o *CreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) WithHTTPClient(client *http.Client) *CreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the create params
func (o *CreateParams) WithAll(all *bool) *CreateParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the create params
func (o *CreateParams) SetAll(all *bool) {
	o.All = all
}

// WithImage adds the image to the create params
func (o *CreateParams) WithImage(image *string) *CreateParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the create params
func (o *CreateParams) SetImage(image *string) {
	o.Image = image
}

// WithName adds the name to the create params
func (o *CreateParams) WithName(name string) *CreateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create params
func (o *CreateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool
		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {
			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}

	}

	if o.Image != nil {

		// query param image
		var qrImage string
		if o.Image != nil {
			qrImage = *o.Image
		}
		qImage := qrImage
		if qImage != "" {
			if err := r.SetQueryParam("image", qImage); err != nil {
				return err
			}
		}

	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
