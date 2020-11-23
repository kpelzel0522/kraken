// Code generated by go-swagger; DO NOT EDIT.

package images_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hpc/kraken/modules/podman/models"
)

// InspectImageReader is a Reader for the InspectImage structure.
type InspectImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InspectImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInspectImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewInspectImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInspectImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInspectImageOK creates a InspectImageOK with default headers values
func NewInspectImageOK() *InspectImageOK {
	return &InspectImageOK{}
}

/*InspectImageOK handles this case with default header values.

Inspect response
*/
type InspectImageOK struct {
	Payload *InspectImageOKBody
}

func (o *InspectImageOK) Error() string {
	return fmt.Sprintf("[GET /images/{name:.*}/json][%d] inspectImageOK  %+v", 200, o.Payload)
}

func (o *InspectImageOK) GetPayload() *InspectImageOKBody {
	return o.Payload
}

func (o *InspectImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InspectImageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInspectImageNotFound creates a InspectImageNotFound with default headers values
func NewInspectImageNotFound() *InspectImageNotFound {
	return &InspectImageNotFound{}
}

/*InspectImageNotFound handles this case with default header values.

No such image
*/
type InspectImageNotFound struct {
	Payload *InspectImageNotFoundBody
}

func (o *InspectImageNotFound) Error() string {
	return fmt.Sprintf("[GET /images/{name:.*}/json][%d] inspectImageNotFound  %+v", 404, o.Payload)
}

func (o *InspectImageNotFound) GetPayload() *InspectImageNotFoundBody {
	return o.Payload
}

func (o *InspectImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InspectImageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInspectImageInternalServerError creates a InspectImageInternalServerError with default headers values
func NewInspectImageInternalServerError() *InspectImageInternalServerError {
	return &InspectImageInternalServerError{}
}

/*InspectImageInternalServerError handles this case with default header values.

Internal server error
*/
type InspectImageInternalServerError struct {
	Payload *InspectImageInternalServerErrorBody
}

func (o *InspectImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{name:.*}/json][%d] inspectImageInternalServerError  %+v", 500, o.Payload)
}

func (o *InspectImageInternalServerError) GetPayload() *InspectImageInternalServerErrorBody {
	return o.Payload
}

func (o *InspectImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InspectImageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*InspectImageInternalServerErrorBody inspect image internal server error body
swagger:model InspectImageInternalServerErrorBody
*/
type InspectImageInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this inspect image internal server error body
func (o *InspectImageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect image internal server error body based on context it is used
func (o *InspectImageInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InspectImageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InspectImageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res InspectImageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InspectImageNotFoundBody inspect image not found body
swagger:model InspectImageNotFoundBody
*/
type InspectImageNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this inspect image not found body
func (o *InspectImageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect image not found body based on context it is used
func (o *InspectImageNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InspectImageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InspectImageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res InspectImageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InspectImageOKBody inspect image o k body
swagger:model InspectImageOKBody
*/
type InspectImageOKBody struct {

	// architecture
	Architecture string `json:"Architecture,omitempty"`

	// author
	Author string `json:"Author,omitempty"`

	// comment
	Comment string `json:"Comment,omitempty"`

	// config
	Config *models.Config `json:"Config,omitempty"`

	// container
	Container string `json:"Container,omitempty"`

	// container config
	ContainerConfig *models.Config `json:"ContainerConfig,omitempty"`

	// created
	Created string `json:"Created,omitempty"`

	// docker version
	DockerVersion string `json:"DockerVersion,omitempty"`

	// graph driver
	GraphDriver *models.GraphDriverData `json:"GraphDriver,omitempty"`

	// ID
	ID string `json:"Id,omitempty"`

	// metadata
	Metadata *models.ImageMetadata `json:"Metadata,omitempty"`

	// os
	Os string `json:"Os,omitempty"`

	// os version
	OsVersion string `json:"OsVersion,omitempty"`

	// parent
	Parent string `json:"Parent,omitempty"`

	// repo digests
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	RepoTags []string `json:"RepoTags"`

	// root f s
	RootFS *models.RootFS `json:"RootFS,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// variant
	Variant string `json:"Variant,omitempty"`

	// virtual size
	VirtualSize int64 `json:"VirtualSize,omitempty"`
}

// Validate validates this inspect image o k body
func (o *InspectImageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateContainerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGraphDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRootFS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InspectImageOKBody) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.Config) { // not required
		return nil
	}

	if o.Config != nil {
		if err := o.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "Config")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) validateContainerConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.ContainerConfig) { // not required
		return nil
	}

	if o.ContainerConfig != nil {
		if err := o.ContainerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "ContainerConfig")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) validateGraphDriver(formats strfmt.Registry) error {
	if swag.IsZero(o.GraphDriver) { // not required
		return nil
	}

	if o.GraphDriver != nil {
		if err := o.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(o.Metadata) { // not required
		return nil
	}

	if o.Metadata != nil {
		if err := o.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "Metadata")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) validateRootFS(formats strfmt.Registry) error {
	if swag.IsZero(o.RootFS) { // not required
		return nil
	}

	if o.RootFS != nil {
		if err := o.RootFS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "RootFS")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inspect image o k body based on the context it is used
func (o *InspectImageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateContainerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGraphDriver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRootFS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InspectImageOKBody) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.Config != nil {
		if err := o.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "Config")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) contextValidateContainerConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.ContainerConfig != nil {
		if err := o.ContainerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "ContainerConfig")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) contextValidateGraphDriver(ctx context.Context, formats strfmt.Registry) error {

	if o.GraphDriver != nil {
		if err := o.GraphDriver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if o.Metadata != nil {
		if err := o.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "Metadata")
			}
			return err
		}
	}

	return nil
}

func (o *InspectImageOKBody) contextValidateRootFS(ctx context.Context, formats strfmt.Registry) error {

	if o.RootFS != nil {
		if err := o.RootFS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspectImageOK" + "." + "RootFS")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InspectImageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InspectImageOKBody) UnmarshalBinary(b []byte) error {
	var res InspectImageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
