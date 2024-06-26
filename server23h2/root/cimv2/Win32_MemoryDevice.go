// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_MemoryDevice struct
type Win32_MemoryDevice struct {
	*Win32_SMBIOSMemory

	//
	ErrorGranularity uint16
}

func NewWin32_MemoryDeviceEx1(instance *cim.WmiInstance) (newInstance *Win32_MemoryDevice, err error) {
	tmp, err := NewWin32_SMBIOSMemoryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_MemoryDevice{
		Win32_SMBIOSMemory: tmp,
	}
	return
}

func NewWin32_MemoryDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_MemoryDevice, err error) {
	tmp, err := NewWin32_SMBIOSMemoryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_MemoryDevice{
		Win32_SMBIOSMemory: tmp,
	}
	return
}

// SetErrorGranularity sets the value of ErrorGranularity for the instance
func (instance *Win32_MemoryDevice) SetPropertyErrorGranularity(value uint16) (err error) {
	return instance.SetProperty("ErrorGranularity", (value))
}

// GetErrorGranularity gets the value of ErrorGranularity for the instance
func (instance *Win32_MemoryDevice) GetPropertyErrorGranularity() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorGranularity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
