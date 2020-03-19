// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_AssignableDeviceDismountSettingData struct
type Msvm_AssignableDeviceDismountSettingData struct {
	*CIM_SettingData

	//
	DeviceInstancePath string

	//
	DeviceLocationPath string

	//
	RequireAcsSupport bool

	//
	RequireDeviceMitigations bool
}

func NewMsvm_AssignableDeviceDismountSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_AssignableDeviceDismountSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AssignableDeviceDismountSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_AssignableDeviceDismountSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AssignableDeviceDismountSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AssignableDeviceDismountSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetDeviceInstancePath sets the value of DeviceInstancePath for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) SetPropertyDeviceInstancePath(value string) (err error) {
	return instance.SetProperty("DeviceInstancePath", value)
}

// GetDeviceInstancePath gets the value of DeviceInstancePath for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) GetPropertyDeviceInstancePath() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceInstancePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceLocationPath sets the value of DeviceLocationPath for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) SetPropertyDeviceLocationPath(value string) (err error) {
	return instance.SetProperty("DeviceLocationPath", value)
}

// GetDeviceLocationPath gets the value of DeviceLocationPath for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) GetPropertyDeviceLocationPath() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceLocationPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireAcsSupport sets the value of RequireAcsSupport for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) SetPropertyRequireAcsSupport(value bool) (err error) {
	return instance.SetProperty("RequireAcsSupport", value)
}

// GetRequireAcsSupport gets the value of RequireAcsSupport for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) GetPropertyRequireAcsSupport() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireAcsSupport")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireDeviceMitigations sets the value of RequireDeviceMitigations for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) SetPropertyRequireDeviceMitigations(value bool) (err error) {
	return instance.SetProperty("RequireDeviceMitigations", value)
}

// GetRequireDeviceMitigations gets the value of RequireDeviceMitigations for the instance
func (instance *Msvm_AssignableDeviceDismountSettingData) GetPropertyRequireDeviceMitigations() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireDeviceMitigations")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
