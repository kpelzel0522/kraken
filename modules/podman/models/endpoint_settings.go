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

// EndpointSettings EndpointSettings stores the network endpoint details
//
// swagger:model EndpointSettings
type EndpointSettings struct {

	// aliases
	Aliases []string `json:"Aliases"`

	// driver opts
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// endpoint ID
	EndpointID string `json:"EndpointID,omitempty"`

	// gateway
	Gateway string `json:"Gateway,omitempty"`

	// global IPv6 address
	GlobalIPV6Address string `json:"GlobalIPv6Address,omitempty"`

	// global IPv6 prefix len
	GlobalIPV6PrefixLen int64 `json:"GlobalIPv6PrefixLen,omitempty"`

	// IP a m config
	IPAMConfig *EndpointIPAMConfig `json:"IPAMConfig,omitempty"`

	// IP address
	IPAddress string `json:"IPAddress,omitempty"`

	// IP prefix len
	IPPrefixLen int64 `json:"IPPrefixLen,omitempty"`

	// IPv6 gateway
	IPV6Gateway string `json:"IPv6Gateway,omitempty"`

	// links
	Links []string `json:"Links"`

	// mac address
	MacAddress string `json:"MacAddress,omitempty"`

	// Operational data
	NetworkID string `json:"NetworkID,omitempty"`
}

// Validate validates this endpoint settings
func (m *EndpointSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAMConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointSettings) validateIPAMConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAMConfig) { // not required
		return nil
	}

	if m.IPAMConfig != nil {
		if err := m.IPAMConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAMConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoint settings based on the context it is used
func (m *EndpointSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAMConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointSettings) contextValidateIPAMConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.IPAMConfig != nil {
		if err := m.IPAMConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAMConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointSettings) UnmarshalBinary(b []byte) error {
	var res EndpointSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
