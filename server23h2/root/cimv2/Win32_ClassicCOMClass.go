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

// Win32_ClassicCOMClass struct
type Win32_ClassicCOMClass struct {
	*Win32_COMClass

	//
	ComponentId string
}

func NewWin32_ClassicCOMClassEx1(instance *cim.WmiInstance) (newInstance *Win32_ClassicCOMClass, err error) {
	tmp, err := NewWin32_COMClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ClassicCOMClass{
		Win32_COMClass: tmp,
	}
	return
}

func NewWin32_ClassicCOMClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ClassicCOMClass, err error) {
	tmp, err := NewWin32_COMClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ClassicCOMClass{
		Win32_COMClass: tmp,
	}
	return
}

// SetComponentId sets the value of ComponentId for the instance
func (instance *Win32_ClassicCOMClass) SetPropertyComponentId(value string) (err error) {
	return instance.SetProperty("ComponentId", (value))
}

// GetComponentId gets the value of ComponentId for the instance
func (instance *Win32_ClassicCOMClass) GetPropertyComponentId() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentId")
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
