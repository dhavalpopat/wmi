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

// Win32_PnPDevicePropertyBooleanArray struct
type Win32_PnPDevicePropertyBooleanArray struct {
	*Win32_PnPDeviceProperty

	//
	Data []bool
}

func NewWin32_PnPDevicePropertyBooleanArrayEx1(instance *cim.WmiInstance) (newInstance *Win32_PnPDevicePropertyBooleanArray, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyBooleanArray{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

func NewWin32_PnPDevicePropertyBooleanArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PnPDevicePropertyBooleanArray, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyBooleanArray{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyBooleanArray) SetPropertyData(value []bool) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyBooleanArray) GetPropertyData() (value []bool, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}
