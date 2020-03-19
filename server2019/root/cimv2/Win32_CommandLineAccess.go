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

// Win32_CommandLineAccess struct
type Win32_CommandLineAccess struct {
	*CIM_ServiceAccessPoint

	//
	CommandLine string
}

func NewWin32_CommandLineAccessEx1(instance *cim.WmiInstance) (newInstance *Win32_CommandLineAccess, err error) {
	tmp, err := NewCIM_ServiceAccessPointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_CommandLineAccess{
		CIM_ServiceAccessPoint: tmp,
	}
	return
}

func NewWin32_CommandLineAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_CommandLineAccess, err error) {
	tmp, err := NewCIM_ServiceAccessPointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_CommandLineAccess{
		CIM_ServiceAccessPoint: tmp,
	}
	return
}

// SetCommandLine sets the value of CommandLine for the instance
func (instance *Win32_CommandLineAccess) SetPropertyCommandLine(value string) (err error) {
	return instance.SetProperty("CommandLine", value)
}

// GetCommandLine gets the value of CommandLine for the instance
func (instance *Win32_CommandLineAccess) GetPropertyCommandLine() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLine")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}