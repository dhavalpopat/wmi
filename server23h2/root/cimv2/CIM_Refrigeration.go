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

// CIM_Refrigeration struct
type CIM_Refrigeration struct {
	*CIM_CoolingDevice
}

func NewCIM_RefrigerationEx1(instance *cim.WmiInstance) (newInstance *CIM_Refrigeration, err error) {
	tmp, err := NewCIM_CoolingDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Refrigeration{
		CIM_CoolingDevice: tmp,
	}
	return
}

func NewCIM_RefrigerationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Refrigeration, err error) {
	tmp, err := NewCIM_CoolingDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Refrigeration{
		CIM_CoolingDevice: tmp,
	}
	return
}
