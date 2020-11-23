// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodStatsContainersReader is a Reader for the LibpodStatsContainers structure.
type LibpodStatsContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodStatsContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodStatsContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodStatsContainersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodStatsContainersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodStatsContainersOK creates a LibpodStatsContainersOK with default headers values
func NewLibpodStatsContainersOK() *LibpodStatsContainersOK {
	return &LibpodStatsContainersOK{}
}

/*LibpodStatsContainersOK handles this case with default header values.

no error
*/
type LibpodStatsContainersOK struct {
}

func (o *LibpodStatsContainersOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/stats][%d] libpodStatsContainersOK ", 200)
}

func (o *LibpodStatsContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodStatsContainersNotFound creates a LibpodStatsContainersNotFound with default headers values
func NewLibpodStatsContainersNotFound() *LibpodStatsContainersNotFound {
	return &LibpodStatsContainersNotFound{}
}

/*LibpodStatsContainersNotFound handles this case with default header values.

No such container
*/
type LibpodStatsContainersNotFound struct {
	Payload *LibpodStatsContainersNotFoundBody
}

func (o *LibpodStatsContainersNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/stats][%d] libpodStatsContainersNotFound  %+v", 404, o.Payload)
}

func (o *LibpodStatsContainersNotFound) GetPayload() *LibpodStatsContainersNotFoundBody {
	return o.Payload
}

func (o *LibpodStatsContainersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStatsContainersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodStatsContainersInternalServerError creates a LibpodStatsContainersInternalServerError with default headers values
func NewLibpodStatsContainersInternalServerError() *LibpodStatsContainersInternalServerError {
	return &LibpodStatsContainersInternalServerError{}
}

/*LibpodStatsContainersInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodStatsContainersInternalServerError struct {
	Payload *LibpodStatsContainersInternalServerErrorBody
}

func (o *LibpodStatsContainersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/stats][%d] libpodStatsContainersInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodStatsContainersInternalServerError) GetPayload() *LibpodStatsContainersInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodStatsContainersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStatsContainersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodStatsContainersInternalServerErrorBody libpod stats containers internal server error body
swagger:model LibpodStatsContainersInternalServerErrorBody
*/
type LibpodStatsContainersInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stats containers internal server error body
func (o *LibpodStatsContainersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod stats containers internal server error body based on context it is used
func (o *LibpodStatsContainersInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStatsContainersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStatsContainersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodStatsContainersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodStatsContainersNotFoundBody libpod stats containers not found body
swagger:model LibpodStatsContainersNotFoundBody
*/
type LibpodStatsContainersNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stats containers not found body
func (o *LibpodStatsContainersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod stats containers not found body based on context it is used
func (o *LibpodStatsContainersNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStatsContainersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStatsContainersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodStatsContainersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
