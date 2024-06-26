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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ODBCDriverAttribute struct
type Win32_ODBCDriverAttribute struct {
	*Win32_SettingCheck
}

func NewWin32_ODBCDriverAttributeEx1(instance *cim.WmiInstance) (newInstance *Win32_ODBCDriverAttribute, err error) {
	tmp, err := NewWin32_SettingCheckEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ODBCDriverAttribute{
		Win32_SettingCheck: tmp,
	}
	return
}

func NewWin32_ODBCDriverAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ODBCDriverAttribute, err error) {
	tmp, err := NewWin32_SettingCheckEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ODBCDriverAttribute{
		Win32_SettingCheck: tmp,
	}
	return
}
