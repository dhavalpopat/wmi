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

// Win32_1394Controller struct
type Win32_1394Controller struct {
	*CIM_Controller

	//
	Manufacturer string
}

func NewWin32_1394ControllerEx1(instance *cim.WmiInstance) (newInstance *Win32_1394Controller, err error) {
	tmp, err := NewCIM_ControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_1394Controller{
		CIM_Controller: tmp,
	}
	return
}

func NewWin32_1394ControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_1394Controller, err error) {
	tmp, err := NewCIM_ControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_1394Controller{
		CIM_Controller: tmp,
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *Win32_1394Controller) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *Win32_1394Controller) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
