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

// Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting struct
type Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting struct {
	*Win32_PerfRawData

	//
	ConnectedClients uint32

	//
	UpdatedPixelsPersec uint32
}

func NewWin32_PerfRawData_RemotePerfProvider_HyperVVMRemotingEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_RemotePerfProvider_HyperVVMRemotingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetConnectedClients sets the value of ConnectedClients for the instance
func (instance *Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting) SetPropertyConnectedClients(value uint32) (err error) {
	return instance.SetProperty("ConnectedClients", (value))
}

// GetConnectedClients gets the value of ConnectedClients for the instance
func (instance *Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting) GetPropertyConnectedClients() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectedClients")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetUpdatedPixelsPersec sets the value of UpdatedPixelsPersec for the instance
func (instance *Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting) SetPropertyUpdatedPixelsPersec(value uint32) (err error) {
	return instance.SetProperty("UpdatedPixelsPersec", (value))
}

// GetUpdatedPixelsPersec gets the value of UpdatedPixelsPersec for the instance
func (instance *Win32_PerfRawData_RemotePerfProvider_HyperVVMRemoting) GetPropertyUpdatedPixelsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdatedPixelsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
