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

// Win32_SystemDriver struct
type Win32_SystemDriver struct {
	*Win32_BaseService
}

func NewWin32_SystemDriverEx1(instance *cim.WmiInstance) (newInstance *Win32_SystemDriver, err error) {
	tmp, err := NewWin32_BaseServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_SystemDriver{
		Win32_BaseService: tmp,
	}
	return
}

func NewWin32_SystemDriverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SystemDriver, err error) {
	tmp, err := NewWin32_BaseServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SystemDriver{
		Win32_BaseService: tmp,
	}
	return
}
