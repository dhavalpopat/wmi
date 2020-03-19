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

// Win32_PerfFormattedData_NETFramework_NETCLRJit struct
type Win32_PerfFormattedData_NETFramework_NETCLRJit struct {
	*Win32_PerfFormattedData

	//
	ILBytesJittedPersec uint32

	//
	NumberofILBytesJitted uint32

	//
	NumberofMethodsJitted uint32

	//
	PercentTimeinJit uint32

	//
	StandardJitFailures uint32

	//
	TotalNumberofILBytesJitted uint32
}

func NewWin32_PerfFormattedData_NETFramework_NETCLRJitEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_NETFramework_NETCLRJit, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_NETFramework_NETCLRJit{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_NETFramework_NETCLRJitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_NETFramework_NETCLRJit, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_NETFramework_NETCLRJit{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetILBytesJittedPersec sets the value of ILBytesJittedPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) SetPropertyILBytesJittedPersec(value uint32) (err error) {
	return instance.SetProperty("ILBytesJittedPersec", value)
}

// GetILBytesJittedPersec gets the value of ILBytesJittedPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) GetPropertyILBytesJittedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ILBytesJittedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofILBytesJitted sets the value of NumberofILBytesJitted for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) SetPropertyNumberofILBytesJitted(value uint32) (err error) {
	return instance.SetProperty("NumberofILBytesJitted", value)
}

// GetNumberofILBytesJitted gets the value of NumberofILBytesJitted for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) GetPropertyNumberofILBytesJitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofILBytesJitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofMethodsJitted sets the value of NumberofMethodsJitted for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) SetPropertyNumberofMethodsJitted(value uint32) (err error) {
	return instance.SetProperty("NumberofMethodsJitted", value)
}

// GetNumberofMethodsJitted gets the value of NumberofMethodsJitted for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) GetPropertyNumberofMethodsJitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofMethodsJitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeinJit sets the value of PercentTimeinJit for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) SetPropertyPercentTimeinJit(value uint32) (err error) {
	return instance.SetProperty("PercentTimeinJit", value)
}

// GetPercentTimeinJit gets the value of PercentTimeinJit for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) GetPropertyPercentTimeinJit() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentTimeinJit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardJitFailures sets the value of StandardJitFailures for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) SetPropertyStandardJitFailures(value uint32) (err error) {
	return instance.SetProperty("StandardJitFailures", value)
}

// GetStandardJitFailures gets the value of StandardJitFailures for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) GetPropertyStandardJitFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardJitFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofILBytesJitted sets the value of TotalNumberofILBytesJitted for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) SetPropertyTotalNumberofILBytesJitted(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofILBytesJitted", value)
}

// GetTotalNumberofILBytesJitted gets the value of TotalNumberofILBytesJitted for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRJit) GetPropertyTotalNumberofILBytesJitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofILBytesJitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
