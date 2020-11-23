// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// PortSet PortSet is a collection of structs indexed by Port
//
// swagger:model PortSet
type PortSet map[string]interface{}

// Validate validates this port set
func (m PortSet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port set based on context it is used
func (m PortSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
