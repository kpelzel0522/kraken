// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hpc/kraken/modules/podman/models"
)

// CompatInspectNetworkReader is a Reader for the CompatInspectNetwork structure.
type CompatInspectNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompatInspectNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompatInspectNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCompatInspectNetworkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompatInspectNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompatInspectNetworkOK creates a CompatInspectNetworkOK with default headers values
func NewCompatInspectNetworkOK() *CompatInspectNetworkOK {
	return &CompatInspectNetworkOK{}
}

/*CompatInspectNetworkOK handles this case with default header values.

Network inspect
*/
type CompatInspectNetworkOK struct {
	Payload *models.NetworkResource
}

func (o *CompatInspectNetworkOK) Error() string {
	return fmt.Sprintf("[GET /networks/{name}][%d] compatInspectNetworkOK  %+v", 200, o.Payload)
}

func (o *CompatInspectNetworkOK) GetPayload() *models.NetworkResource {
	return o.Payload
}

func (o *CompatInspectNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompatInspectNetworkNotFound creates a CompatInspectNetworkNotFound with default headers values
func NewCompatInspectNetworkNotFound() *CompatInspectNetworkNotFound {
	return &CompatInspectNetworkNotFound{}
}

/*CompatInspectNetworkNotFound handles this case with default header values.

No such network
*/
type CompatInspectNetworkNotFound struct {
	Payload *CompatInspectNetworkNotFoundBody
}

func (o *CompatInspectNetworkNotFound) Error() string {
	return fmt.Sprintf("[GET /networks/{name}][%d] compatInspectNetworkNotFound  %+v", 404, o.Payload)
}

func (o *CompatInspectNetworkNotFound) GetPayload() *CompatInspectNetworkNotFoundBody {
	return o.Payload
}

func (o *CompatInspectNetworkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CompatInspectNetworkNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompatInspectNetworkInternalServerError creates a CompatInspectNetworkInternalServerError with default headers values
func NewCompatInspectNetworkInternalServerError() *CompatInspectNetworkInternalServerError {
	return &CompatInspectNetworkInternalServerError{}
}

/*CompatInspectNetworkInternalServerError handles this case with default header values.

Internal server error
*/
type CompatInspectNetworkInternalServerError struct {
	Payload *CompatInspectNetworkInternalServerErrorBody
}

func (o *CompatInspectNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[GET /networks/{name}][%d] compatInspectNetworkInternalServerError  %+v", 500, o.Payload)
}

func (o *CompatInspectNetworkInternalServerError) GetPayload() *CompatInspectNetworkInternalServerErrorBody {
	return o.Payload
}

func (o *CompatInspectNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CompatInspectNetworkInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CompatInspectNetworkInternalServerErrorBody compat inspect network internal server error body
swagger:model CompatInspectNetworkInternalServerErrorBody
*/
type CompatInspectNetworkInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this compat inspect network internal server error body
func (o *CompatInspectNetworkInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this compat inspect network internal server error body based on context it is used
func (o *CompatInspectNetworkInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CompatInspectNetworkInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompatInspectNetworkInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CompatInspectNetworkInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CompatInspectNetworkNotFoundBody compat inspect network not found body
swagger:model CompatInspectNetworkNotFoundBody
*/
type CompatInspectNetworkNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this compat inspect network not found body
func (o *CompatInspectNetworkNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this compat inspect network not found body based on context it is used
func (o *CompatInspectNetworkNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CompatInspectNetworkNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompatInspectNetworkNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CompatInspectNetworkNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
