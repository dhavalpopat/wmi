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

// Win32_DependentService struct
type Win32_DependentService struct {
	*CIM_ServiceServiceDependency
}

func NewWin32_DependentServiceEx1(instance *cim.WmiInstance) (newInstance *Win32_DependentService, err error) {
	tmp, err := NewCIM_ServiceServiceDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_DependentService{
		CIM_ServiceServiceDependency: tmp,
	}
	return
}

func NewWin32_DependentServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_DependentService, err error) {
	tmp, err := NewCIM_ServiceServiceDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_DependentService{
		CIM_ServiceServiceDependency: tmp,
	}
	return
}
