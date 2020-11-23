// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// LibpodPruneVolumesReader is a Reader for the LibpodPruneVolumes structure.
type LibpodPruneVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodPruneVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodPruneVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodPruneVolumesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodPruneVolumesOK creates a LibpodPruneVolumesOK with default headers values
func NewLibpodPruneVolumesOK() *LibpodPruneVolumesOK {
	return &LibpodPruneVolumesOK{}
}

/*LibpodPruneVolumesOK handles this case with default header values.

Volume prune response
*/
type LibpodPruneVolumesOK struct {
	Payload []*models.VolumePruneReport
}

func (o *LibpodPruneVolumesOK) Error() string {
	return fmt.Sprintf("[POST /libpod/volumes/prune][%d] libpodPruneVolumesOK  %+v", 200, o.Payload)
}

func (o *LibpodPruneVolumesOK) GetPayload() []*models.VolumePruneReport {
	return o.Payload
}

func (o *LibpodPruneVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodPruneVolumesInternalServerError creates a LibpodPruneVolumesInternalServerError with default headers values
func NewLibpodPruneVolumesInternalServerError() *LibpodPruneVolumesInternalServerError {
	return &LibpodPruneVolumesInternalServerError{}
}

/*LibpodPruneVolumesInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodPruneVolumesInternalServerError struct {
	Payload *LibpodPruneVolumesInternalServerErrorBody
}

func (o *LibpodPruneVolumesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/volumes/prune][%d] libpodPruneVolumesInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodPruneVolumesInternalServerError) GetPayload() *LibpodPruneVolumesInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodPruneVolumesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodPruneVolumesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodPruneVolumesInternalServerErrorBody libpod prune volumes internal server error body
swagger:model LibpodPruneVolumesInternalServerErrorBody
*/
type LibpodPruneVolumesInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod prune volumes internal server error body
func (o *LibpodPruneVolumesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod prune volumes internal server error body based on context it is used
func (o *LibpodPruneVolumesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodPruneVolumesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodPruneVolumesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodPruneVolumesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
