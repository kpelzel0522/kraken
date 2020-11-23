// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewLibpodListNetworkParams creates a new LibpodListNetworkParams object
// with the default values initialized.
func NewLibpodListNetworkParams() *LibpodListNetworkParams {
	var ()
	return &LibpodListNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodListNetworkParamsWithTimeout creates a new LibpodListNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodListNetworkParamsWithTimeout(timeout time.Duration) *LibpodListNetworkParams {
	var ()
	return &LibpodListNetworkParams{

		timeout: timeout,
	}
}

// NewLibpodListNetworkParamsWithContext creates a new LibpodListNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodListNetworkParamsWithContext(ctx context.Context) *LibpodListNetworkParams {
	var ()
	return &LibpodListNetworkParams{

		Context: ctx,
	}
}

// NewLibpodListNetworkParamsWithHTTPClient creates a new LibpodListNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodListNetworkParamsWithHTTPClient(client *http.Client) *LibpodListNetworkParams {
	var ()
	return &LibpodListNetworkParams{
		HTTPClient: client,
	}
}

/*LibpodListNetworkParams contains all the parameters to send to the API endpoint
for the libpod list network operation typically these are written to a http.Request
*/
type LibpodListNetworkParams struct {

	/*Filter
	  Provide filter values (e.g. 'name=podman')

	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod list network params
func (o *LibpodListNetworkParams) WithTimeout(timeout time.Duration) *LibpodListNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod list network params
func (o *LibpodListNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod list network params
func (o *LibpodListNetworkParams) WithContext(ctx context.Context) *LibpodListNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod list network params
func (o *LibpodListNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod list network params
func (o *LibpodListNetworkParams) WithHTTPClient(client *http.Client) *LibpodListNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod list network params
func (o *LibpodListNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the libpod list network params
func (o *LibpodListNetworkParams) WithFilter(filter *string) *LibpodListNetworkParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the libpod list network params
func (o *LibpodListNetworkParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodListNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
