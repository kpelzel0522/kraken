// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RootFS RootFS holds the root fs information of an image
//
// swagger:model RootFS
type RootFS struct {

	// layers
	Layers []Digest `json:"Layers"`

	// type
	Type string `json:"Type,omitempty"`
}

// Validate validates this root f s
func (m *RootFS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RootFS) validateLayers(formats strfmt.Registry) error {
	if swag.IsZero(m.Layers) { // not required
		return nil
	}

	for i := 0; i < len(m.Layers); i++ {

		if err := m.Layers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Layers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this root f s based on the context it is used
func (m *RootFS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLayers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RootFS) contextValidateLayers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Layers); i++ {

		if err := m.Layers[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Layers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RootFS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootFS) UnmarshalBinary(b []byte) error {
	var res RootFS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
