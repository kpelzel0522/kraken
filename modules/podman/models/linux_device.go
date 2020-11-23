// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinuxDevice LinuxDevice represents the mknod information for a Linux special device file
//
// swagger:model LinuxDevice
type LinuxDevice struct {

	// Gid of the device.
	GID uint32 `json:"gid,omitempty"`

	// Major is the device's major number.
	Major int64 `json:"major,omitempty"`

	// Minor is the device's minor number.
	Minor int64 `json:"minor,omitempty"`

	// Path to the device.
	Path string `json:"path,omitempty"`

	// Device type, block, char, etc.
	Type string `json:"type,omitempty"`

	// UID of the device.
	UID uint32 `json:"uid,omitempty"`

	// file mode
	FileMode FileMode `json:"fileMode,omitempty"`
}

// Validate validates this linux device
func (m *LinuxDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxDevice) validateFileMode(formats strfmt.Registry) error {
	if swag.IsZero(m.FileMode) { // not required
		return nil
	}

	if err := m.FileMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fileMode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this linux device based on the context it is used
func (m *LinuxDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxDevice) contextValidateFileMode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FileMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fileMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinuxDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxDevice) UnmarshalBinary(b []byte) error {
	var res LinuxDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
