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

// UserConfig UserConfig configures the user namespace for the container
//
// swagger:model UserConfig
type UserConfig struct {

	// group add
	GroupAdd []string `json:"GroupAdd"`

	// ID mappings
	IDMappings *IDMappingOptions `json:"IDMappings,omitempty"`

	// user
	User string `json:"User,omitempty"`

	// userns mode
	UsernsMode UsernsMode `json:"UsernsMode,omitempty"`
}

// Validate validates this user config
func (m *UserConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIDMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsernsMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserConfig) validateIDMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.IDMappings) { // not required
		return nil
	}

	if m.IDMappings != nil {
		if err := m.IDMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IDMappings")
			}
			return err
		}
	}

	return nil
}

func (m *UserConfig) validateUsernsMode(formats strfmt.Registry) error {
	if swag.IsZero(m.UsernsMode) { // not required
		return nil
	}

	if err := m.UsernsMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("UsernsMode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this user config based on the context it is used
func (m *UserConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIDMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsernsMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserConfig) contextValidateIDMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.IDMappings != nil {
		if err := m.IDMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IDMappings")
			}
			return err
		}
	}

	return nil
}

func (m *UserConfig) contextValidateUsernsMode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UsernsMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("UsernsMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserConfig) UnmarshalBinary(b []byte) error {
	var res UserConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
