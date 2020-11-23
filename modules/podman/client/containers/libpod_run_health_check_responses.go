// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hpc/kraken/modules/podman/models"
)

// LibpodRunHealthCheckReader is a Reader for the LibpodRunHealthCheck structure.
type LibpodRunHealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodRunHealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodRunHealthCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodRunHealthCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewLibpodRunHealthCheckConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodRunHealthCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodRunHealthCheckOK creates a LibpodRunHealthCheckOK with default headers values
func NewLibpodRunHealthCheckOK() *LibpodRunHealthCheckOK {
	return &LibpodRunHealthCheckOK{}
}

/*LibpodRunHealthCheckOK handles this case with default header values.

Healthcheck
*/
type LibpodRunHealthCheckOK struct {
	Payload *LibpodRunHealthCheckOKBody
}

func (o *LibpodRunHealthCheckOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name:.*}/healthcheck][%d] libpodRunHealthCheckOK  %+v", 200, o.Payload)
}

func (o *LibpodRunHealthCheckOK) GetPayload() *LibpodRunHealthCheckOKBody {
	return o.Payload
}

func (o *LibpodRunHealthCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodRunHealthCheckOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodRunHealthCheckNotFound creates a LibpodRunHealthCheckNotFound with default headers values
func NewLibpodRunHealthCheckNotFound() *LibpodRunHealthCheckNotFound {
	return &LibpodRunHealthCheckNotFound{}
}

/*LibpodRunHealthCheckNotFound handles this case with default header values.

No such container
*/
type LibpodRunHealthCheckNotFound struct {
	Payload *LibpodRunHealthCheckNotFoundBody
}

func (o *LibpodRunHealthCheckNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name:.*}/healthcheck][%d] libpodRunHealthCheckNotFound  %+v", 404, o.Payload)
}

func (o *LibpodRunHealthCheckNotFound) GetPayload() *LibpodRunHealthCheckNotFoundBody {
	return o.Payload
}

func (o *LibpodRunHealthCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodRunHealthCheckNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodRunHealthCheckConflict creates a LibpodRunHealthCheckConflict with default headers values
func NewLibpodRunHealthCheckConflict() *LibpodRunHealthCheckConflict {
	return &LibpodRunHealthCheckConflict{}
}

/*LibpodRunHealthCheckConflict handles this case with default header values.

container has no healthcheck or is not running
*/
type LibpodRunHealthCheckConflict struct {
}

func (o *LibpodRunHealthCheckConflict) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name:.*}/healthcheck][%d] libpodRunHealthCheckConflict ", 409)
}

func (o *LibpodRunHealthCheckConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodRunHealthCheckInternalServerError creates a LibpodRunHealthCheckInternalServerError with default headers values
func NewLibpodRunHealthCheckInternalServerError() *LibpodRunHealthCheckInternalServerError {
	return &LibpodRunHealthCheckInternalServerError{}
}

/*LibpodRunHealthCheckInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodRunHealthCheckInternalServerError struct {
	Payload *LibpodRunHealthCheckInternalServerErrorBody
}

func (o *LibpodRunHealthCheckInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name:.*}/healthcheck][%d] libpodRunHealthCheckInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodRunHealthCheckInternalServerError) GetPayload() *LibpodRunHealthCheckInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodRunHealthCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodRunHealthCheckInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodRunHealthCheckInternalServerErrorBody libpod run health check internal server error body
swagger:model LibpodRunHealthCheckInternalServerErrorBody
*/
type LibpodRunHealthCheckInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod run health check internal server error body
func (o *LibpodRunHealthCheckInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod run health check internal server error body based on context it is used
func (o *LibpodRunHealthCheckInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodRunHealthCheckInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodRunHealthCheckInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodRunHealthCheckInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodRunHealthCheckNotFoundBody libpod run health check not found body
swagger:model LibpodRunHealthCheckNotFoundBody
*/
type LibpodRunHealthCheckNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod run health check not found body
func (o *LibpodRunHealthCheckNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod run health check not found body based on context it is used
func (o *LibpodRunHealthCheckNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodRunHealthCheckNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodRunHealthCheckNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodRunHealthCheckNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodRunHealthCheckOKBody libpod run health check o k body
swagger:model LibpodRunHealthCheckOKBody
*/
type LibpodRunHealthCheckOKBody struct {

	// FailingStreak is the number of consecutive failed healthchecks
	FailingStreak int64 `json:"FailingStreak,omitempty"`

	// Log describes healthcheck attempts and results
	Log []*models.HealthCheckLog `json:"Log"`

	// Status healthy or unhealthy
	Status string `json:"Status,omitempty"`
}

// Validate validates this libpod run health check o k body
func (o *LibpodRunHealthCheckOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodRunHealthCheckOKBody) validateLog(formats strfmt.Registry) error {
	if swag.IsZero(o.Log) { // not required
		return nil
	}

	for i := 0; i < len(o.Log); i++ {
		if swag.IsZero(o.Log[i]) { // not required
			continue
		}

		if o.Log[i] != nil {
			if err := o.Log[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libpodRunHealthCheckOK" + "." + "Log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this libpod run health check o k body based on the context it is used
func (o *LibpodRunHealthCheckOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodRunHealthCheckOKBody) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Log); i++ {

		if o.Log[i] != nil {
			if err := o.Log[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libpodRunHealthCheckOK" + "." + "Log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *LibpodRunHealthCheckOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodRunHealthCheckOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodRunHealthCheckOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
