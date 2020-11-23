// Code generated by go-swagger; DO NOT EDIT.

package images

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

// LibpodChangesImagesReader is a Reader for the LibpodChangesImages structure.
type LibpodChangesImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodChangesImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodChangesImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodChangesImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodChangesImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodChangesImagesOK creates a LibpodChangesImagesOK with default headers values
func NewLibpodChangesImagesOK() *LibpodChangesImagesOK {
	return &LibpodChangesImagesOK{}
}

/*LibpodChangesImagesOK handles this case with default header values.

Array of Changes
*/
type LibpodChangesImagesOK struct {
}

func (o *LibpodChangesImagesOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/changes][%d] libpodChangesImagesOK ", 200)
}

func (o *LibpodChangesImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodChangesImagesNotFound creates a LibpodChangesImagesNotFound with default headers values
func NewLibpodChangesImagesNotFound() *LibpodChangesImagesNotFound {
	return &LibpodChangesImagesNotFound{}
}

/*LibpodChangesImagesNotFound handles this case with default header values.

No such container
*/
type LibpodChangesImagesNotFound struct {
	Payload *LibpodChangesImagesNotFoundBody
}

func (o *LibpodChangesImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/changes][%d] libpodChangesImagesNotFound  %+v", 404, o.Payload)
}

func (o *LibpodChangesImagesNotFound) GetPayload() *LibpodChangesImagesNotFoundBody {
	return o.Payload
}

func (o *LibpodChangesImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodChangesImagesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodChangesImagesInternalServerError creates a LibpodChangesImagesInternalServerError with default headers values
func NewLibpodChangesImagesInternalServerError() *LibpodChangesImagesInternalServerError {
	return &LibpodChangesImagesInternalServerError{}
}

/*LibpodChangesImagesInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodChangesImagesInternalServerError struct {
	Payload *LibpodChangesImagesInternalServerErrorBody
}

func (o *LibpodChangesImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/changes][%d] libpodChangesImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodChangesImagesInternalServerError) GetPayload() *LibpodChangesImagesInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodChangesImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodChangesImagesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodChangesImagesInternalServerErrorBody libpod changes images internal server error body
swagger:model LibpodChangesImagesInternalServerErrorBody
*/
type LibpodChangesImagesInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod changes images internal server error body
func (o *LibpodChangesImagesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod changes images internal server error body based on context it is used
func (o *LibpodChangesImagesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodChangesImagesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodChangesImagesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodChangesImagesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodChangesImagesNotFoundBody libpod changes images not found body
swagger:model LibpodChangesImagesNotFoundBody
*/
type LibpodChangesImagesNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod changes images not found body
func (o *LibpodChangesImagesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod changes images not found body based on context it is used
func (o *LibpodChangesImagesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodChangesImagesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodChangesImagesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodChangesImagesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
