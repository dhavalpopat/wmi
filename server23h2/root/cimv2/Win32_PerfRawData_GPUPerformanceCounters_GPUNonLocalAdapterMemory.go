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

// Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory struct
type Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory struct {
	*Win32_PerfRawData

	//
	NonLocalUsage uint64
}

func NewWin32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemoryEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetNonLocalUsage sets the value of NonLocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory) SetPropertyNonLocalUsage(value uint64) (err error) {
	return instance.SetProperty("NonLocalUsage", (value))
}

// GetNonLocalUsage gets the value of NonLocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory) GetPropertyNonLocalUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonLocalUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
