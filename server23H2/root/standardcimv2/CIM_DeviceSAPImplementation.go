// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_DeviceSAPImplementation struct
type CIM_DeviceSAPImplementation struct {
	*CIM_Dependency
}

func NewCIM_DeviceSAPImplementationEx1(instance *cim.WmiInstance) (newInstance *CIM_DeviceSAPImplementation, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DeviceSAPImplementation{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_DeviceSAPImplementationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DeviceSAPImplementation, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DeviceSAPImplementation{
		CIM_Dependency: tmp,
	}
	return
}
