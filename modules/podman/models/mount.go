// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Mount Mount specifies a mount for a container.
//
// swagger:model Mount
type Mount struct {

	// Destination is the absolute path where the mount will be placed in the container.
	Destination string `json:"destination,omitempty"`

	// Options are fstab style mount options.
	Options []string `json:"options"`

	// Source specifies the source path of the mount.
	Source string `json:"source,omitempty"`

	// Type specifies the mount kind.
	Type string `json:"type,omitempty"`
}

// Validate validates this mount
func (m *Mount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mount based on context it is used
func (m *Mount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Mount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mount) UnmarshalBinary(b []byte) error {
	var res Mount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
