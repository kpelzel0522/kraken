// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerNode ContainerNode stores information about the node that a container
// is running on.  It's only used by the Docker Swarm standalone API
//
// swagger:model ContainerNode
type ContainerNode struct {

	// addr
	Addr string `json:"Addr,omitempty"`

	// cpus
	Cpus int64 `json:"Cpus,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// IP address
	IPAddress string `json:"IP,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// memory
	Memory int64 `json:"Memory,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this container node
func (m *ContainerNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container node based on context it is used
func (m *ContainerNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerNode) UnmarshalBinary(b []byte) error {
	var res ContainerNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
