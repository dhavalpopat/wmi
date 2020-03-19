// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_NTLogEventUser struct
type Win32_NTLogEventUser struct {
	*cim.WmiInstance

	//
	Record Win32_NTLogEvent

	//
	User Win32_UserAccount
}

func NewWin32_NTLogEventUserEx1(instance *cim.WmiInstance) (newInstance *Win32_NTLogEventUser, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_NTLogEventUser{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_NTLogEventUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_NTLogEventUser, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_NTLogEventUser{
		WmiInstance: tmp,
	}
	return
}

// SetRecord sets the value of Record for the instance
func (instance *Win32_NTLogEventUser) SetPropertyRecord(value Win32_NTLogEvent) (err error) {
	return instance.SetProperty("Record", value)
}

// GetRecord gets the value of Record for the instance
func (instance *Win32_NTLogEventUser) GetPropertyRecord() (value Win32_NTLogEvent, err error) {
	retValue, err := instance.GetProperty("Record")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_NTLogEvent)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUser sets the value of User for the instance
func (instance *Win32_NTLogEventUser) SetPropertyUser(value Win32_UserAccount) (err error) {
	return instance.SetProperty("User", value)
}

// GetUser gets the value of User for the instance
func (instance *Win32_NTLogEventUser) GetPropertyUser() (value Win32_UserAccount, err error) {
	retValue, err := instance.GetProperty("User")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_UserAccount)
	if !ok {
		// TODO: Set an error
	}
	return
}
