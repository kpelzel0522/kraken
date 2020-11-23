// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IDMap IDMap contains a single entry for user namespace range remapping. An array
// of IDMap entries represents the structure that will be provided to the Linux
// kernel for creating a user namespace.
//
// swagger:model IDMap
type IDMap struct {

	// container ID
	ContainerID int64 `json:"container_id,omitempty"`

	// host ID
	HostID int64 `json:"host_id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this ID map
func (m *IDMap) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ID map based on context it is used
func (m *IDMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IDMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDMap) UnmarshalBinary(b []byte) error {
	var res IDMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
