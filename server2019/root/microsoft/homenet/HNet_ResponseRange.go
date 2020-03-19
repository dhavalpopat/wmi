// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_ResponseRange struct
type HNet_ResponseRange struct {
	*cim.WmiInstance

	//
	EndPort uint16

	//
	IPProtocol uint8

	//
	StartPort uint16
}

func NewHNet_ResponseRangeEx1(instance *cim.WmiInstance) (newInstance *HNet_ResponseRange, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_ResponseRange{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_ResponseRangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_ResponseRange, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_ResponseRange{
		WmiInstance: tmp,
	}
	return
}

// SetEndPort sets the value of EndPort for the instance
func (instance *HNet_ResponseRange) SetPropertyEndPort(value uint16) (err error) {
	return instance.SetProperty("EndPort", value)
}

// GetEndPort gets the value of EndPort for the instance
func (instance *HNet_ResponseRange) GetPropertyEndPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("EndPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPProtocol sets the value of IPProtocol for the instance
func (instance *HNet_ResponseRange) SetPropertyIPProtocol(value uint8) (err error) {
	return instance.SetProperty("IPProtocol", value)
}

// GetIPProtocol gets the value of IPProtocol for the instance
func (instance *HNet_ResponseRange) GetPropertyIPProtocol() (value uint8, err error) {
	retValue, err := instance.GetProperty("IPProtocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartPort sets the value of StartPort for the instance
func (instance *HNet_ResponseRange) SetPropertyStartPort(value uint16) (err error) {
	return instance.SetProperty("StartPort", value)
}

// GetStartPort gets the value of StartPort for the instance
func (instance *HNet_ResponseRange) GetPropertyStartPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("StartPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}