// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PlayKubePod PlayKubePod represents a single pod and associated containers created by play kube
//
// swagger:model PlayKubePod
type PlayKubePod struct {

	// Containers - the IDs of the containers running in the created pod.
	Containers []string `json:"Containers"`

	// ID - ID of the pod created as a result of play kube.
	ID string `json:"ID,omitempty"`

	// Logs - non-fatal errors and log messages while processing.
	Logs []string `json:"Logs"`
}

// Validate validates this play kube pod
func (m *PlayKubePod) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this play kube pod based on context it is used
func (m *PlayKubePod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayKubePod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayKubePod) UnmarshalBinary(b []byte) error {
	var res PlayKubePod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
