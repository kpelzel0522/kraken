// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OverlayVolume OverlayVolume holds information about a overlay volume that will be mounted into
// the container.
//
// swagger:model OverlayVolume
type OverlayVolume struct {

	// Destination is the absolute path where the mount will be placed in the container.
	Destination string `json:"destination,omitempty"`

	// Source specifies the source path of the mount.
	Source string `json:"source,omitempty"`
}

// Validate validates this overlay volume
func (m *OverlayVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this overlay volume based on context it is used
func (m *OverlayVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OverlayVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OverlayVolume) UnmarshalBinary(b []byte) error {
	var res OverlayVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
