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

// PodNetworkConfig PodNetworkConfig contains networking configuration for a pod.
//
// swagger:model PodNetworkConfig
type PodNetworkConfig struct {

	// CNINetworks is a list of CNI networks that the infra container will
	// join. As, by default, containers share their network with the infra
	// container, these networks will effectively be joined by the
	// entire pod.
	// Only available when NetNS is set to Bridge, the default for root.
	// Optional.
	CNINetworks []string `json:"cni_networks"`

	// DNSOption is a set of DNS options that will be used in the infra
	// container's resolv.conf, which will, by default, be shared with all
	// containers in the pod.
	// Conflicts with NoInfra=true.
	// Optional.
	DNSOption []string `json:"dns_option"`

	// DNSSearch is a set of DNS search domains that will be used in the
	// infra container's resolv.conf, which will, by default, be shared with
	// all containers in the pod.
	// If not provided, DNS search domains from the host's resolv.conf will
	// be used.
	// Conflicts with NoInfra=true.
	// Optional.
	DNSSearch []string `json:"dns_search"`

	// DNSServer is a set of DNS servers that will be used in the infra
	// container's resolv.conf, which will, by default, be shared with all
	// containers in the pod.
	// If not provided, the host's DNS servers will be used, unless the only
	// server set is a localhost address. As the container cannot connect to
	// the host's localhost, a default server will instead be set.
	// Conflicts with NoInfra=true.
	// Optional.
	DNSServer []IP `json:"dns_server"`

	// HostAdd is a set of hosts that will be added to the infra container's
	// etc/hosts that will, by default, be shared with all containers in
	// the pod.
	// Conflicts with NoInfra=true and NoManageHosts.
	// Optional.
	HostAdd []string `json:"hostadd"`

	// NetworkOptions are additional options for each network
	// Optional.
	NetworkOptions map[string][]string `json:"network_options,omitempty"`

	// NoManageHosts indicates that /etc/hosts should not be managed by the
	// pod. Instead, each container will create a separate /etc/hosts as
	// they would if not in a pod.
	// Conflicts with HostAdd.
	NoManageHosts bool `json:"no_manage_hosts,omitempty"`

	// NoManageResolvConf indicates that /etc/resolv.conf should not be
	// managed by the pod. Instead, each container will create and manage a
	// separate resolv.conf as if they had not joined a pod.
	// Conflicts with NoInfra=true and DNSServer, DNSSearch, DNSOption.
	// Optional.
	NoManageResolvConf bool `json:"no_manage_resolv_conf,omitempty"`

	// PortMappings is a set of ports to map into the infra container.
	// As, by default, containers share their network with the infra
	// container, this will forward the ports to the entire pod.
	// Only available if NetNS is set to Bridge or Slirp.
	// Optional.
	PortMappings []*PortMapping `json:"portmappings"`

	// netns
	Netns *Namespace `json:"netns,omitempty"`

	// static ip
	StaticIP IP `json:"static_ip,omitempty"`

	// static mac
	StaticMac HardwareAddr `json:"static_mac,omitempty"`
}

// Validate validates this pod network config
func (m *PodNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticMac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodNetworkConfig) validateDNSServer(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSServer) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSServer); i++ {

		if err := m.DNSServer[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_server" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PodNetworkConfig) validatePortMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.PortMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.PortMappings); i++ {
		if swag.IsZero(m.PortMappings[i]) { // not required
			continue
		}

		if m.PortMappings[i] != nil {
			if err := m.PortMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portmappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodNetworkConfig) validateNetns(formats strfmt.Registry) error {
	if swag.IsZero(m.Netns) { // not required
		return nil
	}

	if m.Netns != nil {
		if err := m.Netns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netns")
			}
			return err
		}
	}

	return nil
}

func (m *PodNetworkConfig) validateStaticIP(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticIP) { // not required
		return nil
	}

	if err := m.StaticIP.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ip")
		}
		return err
	}

	return nil
}

func (m *PodNetworkConfig) validateStaticMac(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticMac) { // not required
		return nil
	}

	if err := m.StaticMac.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_mac")
		}
		return err
	}

	return nil
}

// ContextValidate validate this pod network config based on the context it is used
func (m *PodNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticMac(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodNetworkConfig) contextValidateDNSServer(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServer); i++ {

		if err := m.DNSServer[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_server" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PodNetworkConfig) contextValidatePortMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PortMappings); i++ {

		if m.PortMappings[i] != nil {
			if err := m.PortMappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portmappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodNetworkConfig) contextValidateNetns(ctx context.Context, formats strfmt.Registry) error {

	if m.Netns != nil {
		if err := m.Netns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netns")
			}
			return err
		}
	}

	return nil
}

func (m *PodNetworkConfig) contextValidateStaticIP(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticIP.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ip")
		}
		return err
	}

	return nil
}

func (m *PodNetworkConfig) contextValidateStaticMac(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticMac.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_mac")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodNetworkConfig) UnmarshalBinary(b []byte) error {
	var res PodNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
