// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_DeviceMemoryAddress struct
type Win32_DeviceMemoryAddress struct {
	*Win32_SystemMemoryResource

	//
	MemoryType string
}

func NewWin32_DeviceMemoryAddressEx1(instance *cim.WmiInstance) (newInstance *Win32_DeviceMemoryAddress, err error) {
	tmp, err := NewWin32_SystemMemoryResourceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_DeviceMemoryAddress{
		Win32_SystemMemoryResource: tmp,
	}
	return
}

func NewWin32_DeviceMemoryAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_DeviceMemoryAddress, err error) {
	tmp, err := NewWin32_SystemMemoryResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_DeviceMemoryAddress{
		Win32_SystemMemoryResource: tmp,
	}
	return
}

// SetMemoryType sets the value of MemoryType for the instance
func (instance *Win32_DeviceMemoryAddress) SetPropertyMemoryType(value string) (err error) {
	return instance.SetProperty("MemoryType", value)
}

// GetMemoryType gets the value of MemoryType for the instance
func (instance *Win32_DeviceMemoryAddress) GetPropertyMemoryType() (value string, err error) {
	retValue, err := instance.GetProperty("MemoryType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
