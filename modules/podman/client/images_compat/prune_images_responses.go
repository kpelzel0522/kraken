// Code generated by go-swagger; DO NOT EDIT.

package images_compat

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

// PruneImagesReader is a Reader for the PruneImages structure.
type PruneImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PruneImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPruneImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPruneImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPruneImagesOK creates a PruneImagesOK with default headers values
func NewPruneImagesOK() *PruneImagesOK {
	return &PruneImagesOK{}
}

/*PruneImagesOK handles this case with default header values.

Delete response
*/
type PruneImagesOK struct {
	Payload []*models.ImageDeleteResponse
}

func (o *PruneImagesOK) Error() string {
	return fmt.Sprintf("[POST /images/prune][%d] pruneImagesOK  %+v", 200, o.Payload)
}

func (o *PruneImagesOK) GetPayload() []*models.ImageDeleteResponse {
	return o.Payload
}

func (o *PruneImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPruneImagesInternalServerError creates a PruneImagesInternalServerError with default headers values
func NewPruneImagesInternalServerError() *PruneImagesInternalServerError {
	return &PruneImagesInternalServerError{}
}

/*PruneImagesInternalServerError handles this case with default header values.

Internal server error
*/
type PruneImagesInternalServerError struct {
	Payload *PruneImagesInternalServerErrorBody
}

func (o *PruneImagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/prune][%d] pruneImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PruneImagesInternalServerError) GetPayload() *PruneImagesInternalServerErrorBody {
	return o.Payload
}

func (o *PruneImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PruneImagesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PruneImagesInternalServerErrorBody prune images internal server error body
swagger:model PruneImagesInternalServerErrorBody
*/
type PruneImagesInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this prune images internal server error body
func (o *PruneImagesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prune images internal server error body based on context it is used
func (o *PruneImagesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PruneImagesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PruneImagesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PruneImagesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
