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

// Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU struct
type Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU struct {
	*Win32_PerfRawData

	//
	TCPcurrentconnections uint32
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPUEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPUEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetTCPcurrentconnections sets the value of TCPcurrentconnections for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU) SetPropertyTCPcurrentconnections(value uint32) (err error) {
	return instance.SetProperty("TCPcurrentconnections", (value))
}

// GetTCPcurrentconnections gets the value of TCPcurrentconnections for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU) GetPropertyTCPcurrentconnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPcurrentconnections")
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
