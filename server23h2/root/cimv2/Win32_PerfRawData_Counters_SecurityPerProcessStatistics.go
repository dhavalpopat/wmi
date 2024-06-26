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

// Win32_PerfRawData_Counters_SecurityPerProcessStatistics struct
type Win32_PerfRawData_Counters_SecurityPerProcessStatistics struct {
	*Win32_PerfRawData

	//
	ContextHandles uint32

	//
	CredentialHandles uint32
}

func NewWin32_PerfRawData_Counters_SecurityPerProcessStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_SecurityPerProcessStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_SecurityPerProcessStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_SecurityPerProcessStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_SecurityPerProcessStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_SecurityPerProcessStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetContextHandles sets the value of ContextHandles for the instance
func (instance *Win32_PerfRawData_Counters_SecurityPerProcessStatistics) SetPropertyContextHandles(value uint32) (err error) {
	return instance.SetProperty("ContextHandles", (value))
}

// GetContextHandles gets the value of ContextHandles for the instance
func (instance *Win32_PerfRawData_Counters_SecurityPerProcessStatistics) GetPropertyContextHandles() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContextHandles")
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

// SetCredentialHandles sets the value of CredentialHandles for the instance
func (instance *Win32_PerfRawData_Counters_SecurityPerProcessStatistics) SetPropertyCredentialHandles(value uint32) (err error) {
	return instance.SetProperty("CredentialHandles", (value))
}

// GetCredentialHandles gets the value of CredentialHandles for the instance
func (instance *Win32_PerfRawData_Counters_SecurityPerProcessStatistics) GetPropertyCredentialHandles() (value uint32, err error) {
	retValue, err := instance.GetProperty("CredentialHandles")
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
