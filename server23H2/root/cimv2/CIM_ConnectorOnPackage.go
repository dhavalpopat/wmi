// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ConnectorOnPackage struct
type CIM_ConnectorOnPackage struct {
	*CIM_Container
}

func NewCIM_ConnectorOnPackageEx1(instance *cim.WmiInstance) (newInstance *CIM_ConnectorOnPackage, err error) {
	tmp, err := NewCIM_ContainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ConnectorOnPackage{
		CIM_Container: tmp,
	}
	return
}

func NewCIM_ConnectorOnPackageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ConnectorOnPackage, err error) {
	tmp, err := NewCIM_ContainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ConnectorOnPackage{
		CIM_Container: tmp,
	}
	return
}
