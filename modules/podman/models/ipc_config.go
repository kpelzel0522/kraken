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

// IpcConfig IpcConfig configures the ipc namespace for the container
//
// swagger:model IpcConfig
type IpcConfig struct {

	// ipc mode
	IpcMode IpcMode `json:"IpcMode,omitempty"`
}

// Validate validates this ipc config
func (m *IpcConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIpcMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpcConfig) validateIpcMode(formats strfmt.Registry) error {
	if swag.IsZero(m.IpcMode) { // not required
		return nil
	}

	if err := m.IpcMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("IpcMode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this ipc config based on the context it is used
func (m *IpcConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIpcMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpcConfig) contextValidateIpcMode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IpcMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("IpcMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpcConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpcConfig) UnmarshalBinary(b []byte) error {
	var res IpcConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
