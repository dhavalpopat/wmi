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

// CIM_TemperatureSensor struct
type CIM_TemperatureSensor struct {
	*CIM_NumericSensor
}

func NewCIM_TemperatureSensorEx1(instance *cim.WmiInstance) (newInstance *CIM_TemperatureSensor, err error) {
	tmp, err := NewCIM_NumericSensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_TemperatureSensor{
		CIM_NumericSensor: tmp,
	}
	return
}

func NewCIM_TemperatureSensorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_TemperatureSensor, err error) {
	tmp, err := NewCIM_NumericSensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_TemperatureSensor{
		CIM_NumericSensor: tmp,
	}
	return
}
