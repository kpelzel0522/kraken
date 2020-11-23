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

// NewLibpodImagesPullParams creates a new LibpodImagesPullParams object
// with the default values initialized.
func NewLibpodImagesPullParams() *LibpodImagesPullParams {
	var (
		tLSVerifyDefault = bool(true)
	)
	return &LibpodImagesPullParams{
		TLSVerify: &tLSVerifyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodImagesPullParamsWithTimeout creates a new LibpodImagesPullParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodImagesPullParamsWithTimeout(timeout time.Duration) *LibpodImagesPullParams {
	var (
		tLSVerifyDefault = bool(true)
	)
	return &LibpodImagesPullParams{
		TLSVerify: &tLSVerifyDefault,

		timeout: timeout,
	}
}

// NewLibpodImagesPullParamsWithContext creates a new LibpodImagesPullParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodImagesPullParamsWithContext(ctx context.Context) *LibpodImagesPullParams {
	var (
		tlsVerifyDefault = bool(true)
	)
	return &LibpodImagesPullParams{
		TLSVerify: &tlsVerifyDefault,

		Context: ctx,
	}
}

// NewLibpodImagesPullParamsWithHTTPClient creates a new LibpodImagesPullParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodImagesPullParamsWithHTTPClient(client *http.Client) *LibpodImagesPullParams {
	var (
		tlsVerifyDefault = bool(true)
	)
	return &LibpodImagesPullParams{
		TLSVerify:  &tlsVerifyDefault,
		HTTPClient: client,
	}
}

/*LibpodImagesPullParams contains all the parameters to send to the API endpoint
for the libpod images pull operation typically these are written to a http.Request
*/
type LibpodImagesPullParams struct {

	/*AllTags
	  Pull all tagged images in the repository.

	*/
	AllTags *bool
	/*Credentials
	  username:password for the registry

	*/
	Credentials *string
	/*OverrideArch
	  Pull image for the specified architecture.

	*/
	OverrideArch *string
	/*OverrideOS
	  Pull image for the specified operating system.

	*/
	OverrideOS *string
	/*OverrideVariant
	  Pull image for the specified variant.

	*/
	OverrideVariant *string
	/*Reference
	  Mandatory reference to the image (e.g., quay.io/image/name:tag)

	*/
	Reference *string
	/*TLSVerify
	  Require TLS verification.

	*/
	TLSVerify *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod images pull params
func (o *LibpodImagesPullParams) WithTimeout(timeout time.Duration) *LibpodImagesPullParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod images pull params
func (o *LibpodImagesPullParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod images pull params
func (o *LibpodImagesPullParams) WithContext(ctx context.Context) *LibpodImagesPullParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod images pull params
func (o *LibpodImagesPullParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod images pull params
func (o *LibpodImagesPullParams) WithHTTPClient(client *http.Client) *LibpodImagesPullParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod images pull params
func (o *LibpodImagesPullParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllTags adds the allTags to the libpod images pull params
func (o *LibpodImagesPullParams) WithAllTags(allTags *bool) *LibpodImagesPullParams {
	o.SetAllTags(allTags)
	return o
}

// SetAllTags adds the allTags to the libpod images pull params
func (o *LibpodImagesPullParams) SetAllTags(allTags *bool) {
	o.AllTags = allTags
}

// WithCredentials adds the credentials to the libpod images pull params
func (o *LibpodImagesPullParams) WithCredentials(credentials *string) *LibpodImagesPullParams {
	o.SetCredentials(credentials)
	return o
}

// SetCredentials adds the credentials to the libpod images pull params
func (o *LibpodImagesPullParams) SetCredentials(credentials *string) {
	o.Credentials = credentials
}

// WithOverrideArch adds the overrideArch to the libpod images pull params
func (o *LibpodImagesPullParams) WithOverrideArch(overrideArch *string) *LibpodImagesPullParams {
	o.SetOverrideArch(overrideArch)
	return o
}

// SetOverrideArch adds the overrideArch to the libpod images pull params
func (o *LibpodImagesPullParams) SetOverrideArch(overrideArch *string) {
	o.OverrideArch = overrideArch
}

// WithOverrideOS adds the overrideOS to the libpod images pull params
func (o *LibpodImagesPullParams) WithOverrideOS(overrideOS *string) *LibpodImagesPullParams {
	o.SetOverrideOS(overrideOS)
	return o
}

// SetOverrideOS adds the overrideOS to the libpod images pull params
func (o *LibpodImagesPullParams) SetOverrideOS(overrideOS *string) {
	o.OverrideOS = overrideOS
}

// WithOverrideVariant adds the overrideVariant to the libpod images pull params
func (o *LibpodImagesPullParams) WithOverrideVariant(overrideVariant *string) *LibpodImagesPullParams {
	o.SetOverrideVariant(overrideVariant)
	return o
}

// SetOverrideVariant adds the overrideVariant to the libpod images pull params
func (o *LibpodImagesPullParams) SetOverrideVariant(overrideVariant *string) {
	o.OverrideVariant = overrideVariant
}

// WithReference adds the reference to the libpod images pull params
func (o *LibpodImagesPullParams) WithReference(reference *string) *LibpodImagesPullParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the libpod images pull params
func (o *LibpodImagesPullParams) SetReference(reference *string) {
	o.Reference = reference
}

// WithTLSVerify adds the tLSVerify to the libpod images pull params
func (o *LibpodImagesPullParams) WithTLSVerify(tLSVerify *bool) *LibpodImagesPullParams {
	o.SetTLSVerify(tLSVerify)
	return o
}

// SetTLSVerify adds the tlsVerify to the libpod images pull params
func (o *LibpodImagesPullParams) SetTLSVerify(tLSVerify *bool) {
	o.TLSVerify = tLSVerify
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodImagesPullParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllTags != nil {

		// query param allTags
		var qrAllTags bool
		if o.AllTags != nil {
			qrAllTags = *o.AllTags
		}
		qAllTags := swag.FormatBool(qrAllTags)
		if qAllTags != "" {
			if err := r.SetQueryParam("allTags", qAllTags); err != nil {
				return err
			}
		}

	}

	if o.Credentials != nil {

		// query param credentials
		var qrCredentials string
		if o.Credentials != nil {
			qrCredentials = *o.Credentials
		}
		qCredentials := qrCredentials
		if qCredentials != "" {
			if err := r.SetQueryParam("credentials", qCredentials); err != nil {
				return err
			}
		}

	}

	if o.OverrideArch != nil {

		// query param overrideArch
		var qrOverrideArch string
		if o.OverrideArch != nil {
			qrOverrideArch = *o.OverrideArch
		}
		qOverrideArch := qrOverrideArch
		if qOverrideArch != "" {
			if err := r.SetQueryParam("overrideArch", qOverrideArch); err != nil {
				return err
			}
		}

	}

	if o.OverrideOS != nil {

		// query param overrideOS
		var qrOverrideOS string
		if o.OverrideOS != nil {
			qrOverrideOS = *o.OverrideOS
		}
		qOverrideOS := qrOverrideOS
		if qOverrideOS != "" {
			if err := r.SetQueryParam("overrideOS", qOverrideOS); err != nil {
				return err
			}
		}

	}

	if o.OverrideVariant != nil {

		// query param overrideVariant
		var qrOverrideVariant string
		if o.OverrideVariant != nil {
			qrOverrideVariant = *o.OverrideVariant
		}
		qOverrideVariant := qrOverrideVariant
		if qOverrideVariant != "" {
			if err := r.SetQueryParam("overrideVariant", qOverrideVariant); err != nil {
				return err
			}
		}

	}

	if o.Reference != nil {

		// query param reference
		var qrReference string
		if o.Reference != nil {
			qrReference = *o.Reference
		}
		qReference := qrReference
		if qReference != "" {
			if err := r.SetQueryParam("reference", qReference); err != nil {
				return err
			}
		}

	}

	if o.TLSVerify != nil {

		// query param tlsVerify
		var qrTLSVerify bool
		if o.TLSVerify != nil {
			qrTLSVerify = *o.TLSVerify
		}
		qTLSVerify := swag.FormatBool(qrTLSVerify)
		if qTLSVerify != "" {
			if err := r.SetQueryParam("tlsVerify", qTLSVerify); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
