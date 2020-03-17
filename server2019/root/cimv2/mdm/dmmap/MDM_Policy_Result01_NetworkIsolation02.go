// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_NetworkIsolation02 struct
type MDM_Policy_Result01_NetworkIsolation02 struct {
	cim.WmiInstance

	//
	EnterpriseCloudResources string

	//
	EnterpriseInternalProxyServers string

	//
	EnterpriseIPRange string

	//
	EnterpriseIPRangesAreAuthoritative int32

	//
	EnterpriseNetworkDomainNames string

	//
	EnterpriseProxyServers string

	//
	EnterpriseProxyServersAreAuthoritative int32

	//
	InstanceID string

	//
	NeutralResources string

	//
	ParentID string
}

// SetEnterpriseCloudResources sets the value of EnterpriseCloudResources for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseCloudResources(value string) (err error) {
	return instance.SetProperty("EnterpriseCloudResources", value)
}

// GetEnterpriseCloudResources gets the value of EnterpriseCloudResources for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseCloudResources() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseCloudResources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseInternalProxyServers sets the value of EnterpriseInternalProxyServers for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseInternalProxyServers(value string) (err error) {
	return instance.SetProperty("EnterpriseInternalProxyServers", value)
}

// GetEnterpriseInternalProxyServers gets the value of EnterpriseInternalProxyServers for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseInternalProxyServers() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseInternalProxyServers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseIPRange sets the value of EnterpriseIPRange for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseIPRange(value string) (err error) {
	return instance.SetProperty("EnterpriseIPRange", value)
}

// GetEnterpriseIPRange gets the value of EnterpriseIPRange for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseIPRange() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseIPRange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseIPRangesAreAuthoritative sets the value of EnterpriseIPRangesAreAuthoritative for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseIPRangesAreAuthoritative(value int32) (err error) {
	return instance.SetProperty("EnterpriseIPRangesAreAuthoritative", value)
}

// GetEnterpriseIPRangesAreAuthoritative gets the value of EnterpriseIPRangesAreAuthoritative for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseIPRangesAreAuthoritative() (value int32, err error) {
	retValue, err := instance.GetProperty("EnterpriseIPRangesAreAuthoritative")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseNetworkDomainNames sets the value of EnterpriseNetworkDomainNames for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseNetworkDomainNames(value string) (err error) {
	return instance.SetProperty("EnterpriseNetworkDomainNames", value)
}

// GetEnterpriseNetworkDomainNames gets the value of EnterpriseNetworkDomainNames for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseNetworkDomainNames() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseNetworkDomainNames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseProxyServers sets the value of EnterpriseProxyServers for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseProxyServers(value string) (err error) {
	return instance.SetProperty("EnterpriseProxyServers", value)
}

// GetEnterpriseProxyServers gets the value of EnterpriseProxyServers for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseProxyServers() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseProxyServers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnterpriseProxyServersAreAuthoritative sets the value of EnterpriseProxyServersAreAuthoritative for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyEnterpriseProxyServersAreAuthoritative(value int32) (err error) {
	return instance.SetProperty("EnterpriseProxyServersAreAuthoritative", value)
}

// GetEnterpriseProxyServersAreAuthoritative gets the value of EnterpriseProxyServersAreAuthoritative for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyEnterpriseProxyServersAreAuthoritative() (value int32, err error) {
	retValue, err := instance.GetProperty("EnterpriseProxyServersAreAuthoritative")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNeutralResources sets the value of NeutralResources for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyNeutralResources(value string) (err error) {
	return instance.SetProperty("NeutralResources", value)
}

// GetNeutralResources gets the value of NeutralResources for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyNeutralResources() (value string, err error) {
	retValue, err := instance.GetProperty("NeutralResources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_NetworkIsolation02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
