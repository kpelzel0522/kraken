// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewLibpodExportImagesParams creates a new LibpodExportImagesParams object
// with the default values initialized.
func NewLibpodExportImagesParams() *LibpodExportImagesParams {
	var ()
	return &LibpodExportImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodExportImagesParamsWithTimeout creates a new LibpodExportImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodExportImagesParamsWithTimeout(timeout time.Duration) *LibpodExportImagesParams {
	var ()
	return &LibpodExportImagesParams{

		timeout: timeout,
	}
}

// NewLibpodExportImagesParamsWithContext creates a new LibpodExportImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodExportImagesParamsWithContext(ctx context.Context) *LibpodExportImagesParams {
	var ()
	return &LibpodExportImagesParams{

		Context: ctx,
	}
}

// NewLibpodExportImagesParamsWithHTTPClient creates a new LibpodExportImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodExportImagesParamsWithHTTPClient(client *http.Client) *LibpodExportImagesParams {
	var ()
	return &LibpodExportImagesParams{
		HTTPClient: client,
	}
}

/*LibpodExportImagesParams contains all the parameters to send to the API endpoint
for the libpod export images operation typically these are written to a http.Request
*/
type LibpodExportImagesParams struct {

	/*Compress
	  use compression on image

	*/
	Compress *bool
	/*Format
	  format for exported image (only docker-archive is supported)

	*/
	Format *string
	/*References
	  references to images to export

	*/
	References []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod export images params
func (o *LibpodExportImagesParams) WithTimeout(timeout time.Duration) *LibpodExportImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod export images params
func (o *LibpodExportImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod export images params
func (o *LibpodExportImagesParams) WithContext(ctx context.Context) *LibpodExportImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod export images params
func (o *LibpodExportImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod export images params
func (o *LibpodExportImagesParams) WithHTTPClient(client *http.Client) *LibpodExportImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod export images params
func (o *LibpodExportImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompress adds the compress to the libpod export images params
func (o *LibpodExportImagesParams) WithCompress(compress *bool) *LibpodExportImagesParams {
	o.SetCompress(compress)
	return o
}

// SetCompress adds the compress to the libpod export images params
func (o *LibpodExportImagesParams) SetCompress(compress *bool) {
	o.Compress = compress
}

// WithFormat adds the format to the libpod export images params
func (o *LibpodExportImagesParams) WithFormat(format *string) *LibpodExportImagesParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the libpod export images params
func (o *LibpodExportImagesParams) SetFormat(format *string) {
	o.Format = format
}

// WithReferences adds the references to the libpod export images params
func (o *LibpodExportImagesParams) WithReferences(references []string) *LibpodExportImagesParams {
	o.SetReferences(references)
	return o
}

// SetReferences adds the references to the libpod export images params
func (o *LibpodExportImagesParams) SetReferences(references []string) {
	o.References = references
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodExportImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Compress != nil {

		// query param compress
		var qrCompress bool
		if o.Compress != nil {
			qrCompress = *o.Compress
		}
		qCompress := swag.FormatBool(qrCompress)
		if qCompress != "" {
			if err := r.SetQueryParam("compress", qCompress); err != nil {
				return err
			}
		}

	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	valuesReferences := o.References

	joinedReferences := swag.JoinByFormat(valuesReferences, "")
	// query array param references
	if err := r.SetQueryParam("references", joinedReferences...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
