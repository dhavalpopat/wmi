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

// Win32_ProcessStopTrace struct
type Win32_ProcessStopTrace struct {
	*Win32_ProcessTrace

	//
	ExitStatus uint32
}

func NewWin32_ProcessStopTraceEx1(instance *cim.WmiInstance) (newInstance *Win32_ProcessStopTrace, err error) {
	tmp, err := NewWin32_ProcessTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ProcessStopTrace{
		Win32_ProcessTrace: tmp,
	}
	return
}

func NewWin32_ProcessStopTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ProcessStopTrace, err error) {
	tmp, err := NewWin32_ProcessTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ProcessStopTrace{
		Win32_ProcessTrace: tmp,
	}
	return
}

// SetExitStatus sets the value of ExitStatus for the instance
func (instance *Win32_ProcessStopTrace) SetPropertyExitStatus(value uint32) (err error) {
	return instance.SetProperty("ExitStatus", (value))
}

// GetExitStatus gets the value of ExitStatus for the instance
func (instance *Win32_ProcessStopTrace) GetPropertyExitStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExitStatus")
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
