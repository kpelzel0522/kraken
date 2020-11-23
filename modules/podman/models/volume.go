// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Volume Volume volume
//
// swagger:model Volume
type Volume struct {

	// Date/Time the volume was created.
	CreatedAt string `json:"CreatedAt,omitempty"`

	// Name of the volume driver used by the volume.
	// Required: true
	Driver *string `json:"Driver"`

	// User-defined key/value metadata.
	// Required: true
	Labels map[string]string `json:"Labels"`

	// Mount path of the volume on the host.
	// Required: true
	Mountpoint *string `json:"Mountpoint"`

	// Name of the volume.
	// Required: true
	Name *string `json:"Name"`

	// The driver specific options used when creating the volume.
	// Required: true
	Options map[string]string `json:"Options"`

	// The level at which the volume exists. Either `global` for cluster-wide,
	// or `local` for machine level.
	// Required: true
	Scope *string `json:"Scope"`

	// Low-level details about the volume, provided by the volume driver.
	// Details are returned as a map with key/value pairs:
	// `{"key":"value","key2":"value2"}`.
	//
	// The `Status` field is optional, and is omitted if the volume driver
	// does not support this feature.
	Status map[string]interface{} `json:"Status,omitempty"`

	// usage data
	UsageData *VolumeUsageData `json:"UsageData,omitempty"`
}

// Validate validates this volume
func (m *Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("Driver", "body", m.Driver); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("Labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateMountpoint(formats strfmt.Registry) error {

	if err := validate.Required("Mountpoint", "body", m.Mountpoint); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("Options", "body", m.Options); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("Scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateUsageData(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageData) { // not required
		return nil
	}

	if m.UsageData != nil {
		if err := m.UsageData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsageData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume based on the context it is used
func (m *Volume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsageData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) contextValidateUsageData(ctx context.Context, formats strfmt.Registry) error {

	if m.UsageData != nil {
		if err := m.UsageData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsageData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volume) UnmarshalBinary(b []byte) error {
	var res Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
