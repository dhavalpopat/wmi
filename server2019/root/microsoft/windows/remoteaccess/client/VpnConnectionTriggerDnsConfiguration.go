// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// VpnConnectionTriggerDnsConfiguration struct
type VpnConnectionTriggerDnsConfiguration struct {
	cim.WmiInstance

	//
	ConnectionName string

	//
	DnsIPAddress []string

	//
	DnsSuffix string

	//
	DnsSuffixSearchList []string
}

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", value)
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyConnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsIPAddress sets the value of DnsIPAddress for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyDnsIPAddress(value []string) (err error) {
	return instance.SetProperty("DnsIPAddress", value)
}

// GetDnsIPAddress gets the value of DnsIPAddress for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyDnsIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsIPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", value)
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffixSearchList sets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) SetPropertyDnsSuffixSearchList(value []string) (err error) {
	return instance.SetProperty("DnsSuffixSearchList", value)
}

// GetDnsSuffixSearchList gets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTriggerDnsConfiguration) GetPropertyDnsSuffixSearchList() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsSuffixSearchList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
