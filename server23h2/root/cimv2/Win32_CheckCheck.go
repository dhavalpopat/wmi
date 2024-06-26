// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_CheckCheck struct
type Win32_CheckCheck struct {
	*cim.WmiInstance

	//
	Check CIM_Check

	//
	Location CIM_Check
}

func NewWin32_CheckCheckEx1(instance *cim.WmiInstance) (newInstance *Win32_CheckCheck, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_CheckCheck{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_CheckCheckEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_CheckCheck, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_CheckCheck{
		WmiInstance: tmp,
	}
	return
}

// SetCheck sets the value of Check for the instance
func (instance *Win32_CheckCheck) SetPropertyCheck(value CIM_Check) (err error) {
	return instance.SetProperty("Check", (value))
}

// GetCheck gets the value of Check for the instance
func (instance *Win32_CheckCheck) GetPropertyCheck() (value CIM_Check, err error) {
	retValue, err := instance.GetProperty("Check")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_Check)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_Check is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_Check(valuetmp)

	return
}

// SetLocation sets the value of Location for the instance
func (instance *Win32_CheckCheck) SetPropertyLocation(value CIM_Check) (err error) {
	return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *Win32_CheckCheck) GetPropertyLocation() (value CIM_Check, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_Check)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_Check is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_Check(valuetmp)

	return
}
