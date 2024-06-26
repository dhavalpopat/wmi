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

// Win32_VoltageProbe struct
type Win32_VoltageProbe struct {
	*CIM_VoltageSensor
}

func NewWin32_VoltageProbeEx1(instance *cim.WmiInstance) (newInstance *Win32_VoltageProbe, err error) {
	tmp, err := NewCIM_VoltageSensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_VoltageProbe{
		CIM_VoltageSensor: tmp,
	}
	return
}

func NewWin32_VoltageProbeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_VoltageProbe, err error) {
	tmp, err := NewCIM_VoltageSensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_VoltageProbe{
		CIM_VoltageSensor: tmp,
	}
	return
}
