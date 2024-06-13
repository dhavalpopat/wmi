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

// CIM_MemoryOnCard struct
type CIM_MemoryOnCard struct {
	*CIM_PackagedComponent
}

func NewCIM_MemoryOnCardEx1(instance *cim.WmiInstance) (newInstance *CIM_MemoryOnCard, err error) {
	tmp, err := NewCIM_PackagedComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MemoryOnCard{
		CIM_PackagedComponent: tmp,
	}
	return
}

func NewCIM_MemoryOnCardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MemoryOnCard, err error) {
	tmp, err := NewCIM_PackagedComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MemoryOnCard{
		CIM_PackagedComponent: tmp,
	}
	return
}
