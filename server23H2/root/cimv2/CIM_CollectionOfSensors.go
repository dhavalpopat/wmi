// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_CollectionOfSensors struct
type CIM_CollectionOfSensors struct {
	*CIM_Component
}

func NewCIM_CollectionOfSensorsEx1(instance *cim.WmiInstance) (newInstance *CIM_CollectionOfSensors, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_CollectionOfSensors{
		CIM_Component: tmp,
	}
	return
}

func NewCIM_CollectionOfSensorsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_CollectionOfSensors, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_CollectionOfSensors{
		CIM_Component: tmp,
	}
	return
}
