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

// Win32_LogicalShareSecuritySetting struct
type Win32_LogicalShareSecuritySetting struct {
	*Win32_SecuritySetting

	//
	Name string
}

func NewWin32_LogicalShareSecuritySettingEx1(instance *cim.WmiInstance) (newInstance *Win32_LogicalShareSecuritySetting, err error) {
	tmp, err := NewWin32_SecuritySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_LogicalShareSecuritySetting{
		Win32_SecuritySetting: tmp,
	}
	return
}

func NewWin32_LogicalShareSecuritySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_LogicalShareSecuritySetting, err error) {
	tmp, err := NewWin32_SecuritySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_LogicalShareSecuritySetting{
		Win32_SecuritySetting: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_LogicalShareSecuritySetting) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Win32_LogicalShareSecuritySetting) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
