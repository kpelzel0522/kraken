// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// PauseContainerReader is a Reader for the PauseContainer structure.
type PauseContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPauseContainerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPauseContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPauseContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPauseContainerNoContent creates a PauseContainerNoContent with default headers values
func NewPauseContainerNoContent() *PauseContainerNoContent {
	return &PauseContainerNoContent{}
}

/*PauseContainerNoContent handles this case with default header values.

no error
*/
type PauseContainerNoContent struct {
}

func (o *PauseContainerNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] pauseContainerNoContent ", 204)
}

func (o *PauseContainerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPauseContainerNotFound creates a PauseContainerNotFound with default headers values
func NewPauseContainerNotFound() *PauseContainerNotFound {
	return &PauseContainerNotFound{}
}

/*PauseContainerNotFound handles this case with default header values.

No such container
*/
type PauseContainerNotFound struct {
	Payload *PauseContainerNotFoundBody
}

func (o *PauseContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] pauseContainerNotFound  %+v", 404, o.Payload)
}

func (o *PauseContainerNotFound) GetPayload() *PauseContainerNotFoundBody {
	return o.Payload
}

func (o *PauseContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PauseContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseContainerInternalServerError creates a PauseContainerInternalServerError with default headers values
func NewPauseContainerInternalServerError() *PauseContainerInternalServerError {
	return &PauseContainerInternalServerError{}
}

/*PauseContainerInternalServerError handles this case with default header values.

Internal server error
*/
type PauseContainerInternalServerError struct {
	Payload *PauseContainerInternalServerErrorBody
}

func (o *PauseContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] pauseContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *PauseContainerInternalServerError) GetPayload() *PauseContainerInternalServerErrorBody {
	return o.Payload
}

func (o *PauseContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PauseContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PauseContainerInternalServerErrorBody pause container internal server error body
swagger:model PauseContainerInternalServerErrorBody
*/
type PauseContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pause container internal server error body
func (o *PauseContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pause container internal server error body based on context it is used
func (o *PauseContainerInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PauseContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PauseContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PauseContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PauseContainerNotFoundBody pause container not found body
swagger:model PauseContainerNotFoundBody
*/
type PauseContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pause container not found body
func (o *PauseContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pause container not found body based on context it is used
func (o *PauseContainerNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PauseContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PauseContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PauseContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
