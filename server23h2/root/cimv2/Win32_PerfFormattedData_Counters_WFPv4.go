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
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Counters_WFPv4 struct
type Win32_PerfFormattedData_Counters_WFPv4 struct {
	*Win32_PerfFormattedData

	//
	ActiveInboundConnections uint32

	//
	ActiveOutboundConnections uint32

	//
	AllowedClassifiesPersec uint32

	//
	BlockedBinds uint32

	//
	InboundConnections uint32

	//
	InboundConnectionsAllowedPersec uint32

	//
	InboundConnectionsBlockedPersec uint32

	//
	InboundPacketsDiscardedPersec uint32

	//
	OutboundConnections uint32

	//
	OutboundConnectionsAllowedPersec uint32

	//
	OutboundConnectionsBlockedPersec uint32

	//
	OutboundPacketsDiscardedPersec uint32

	//
	PacketsDiscardedPersec uint32
}

func NewWin32_PerfFormattedData_Counters_WFPv4Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_WFPv4, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_WFPv4{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_WFPv4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_WFPv4, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_WFPv4{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetActiveInboundConnections sets the value of ActiveInboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyActiveInboundConnections(value uint32) (err error) {
	return instance.SetProperty("ActiveInboundConnections", (value))
}

// GetActiveInboundConnections gets the value of ActiveInboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyActiveInboundConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveInboundConnections")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetActiveOutboundConnections sets the value of ActiveOutboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyActiveOutboundConnections(value uint32) (err error) {
	return instance.SetProperty("ActiveOutboundConnections", (value))
}

// GetActiveOutboundConnections gets the value of ActiveOutboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyActiveOutboundConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveOutboundConnections")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAllowedClassifiesPersec sets the value of AllowedClassifiesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyAllowedClassifiesPersec(value uint32) (err error) {
	return instance.SetProperty("AllowedClassifiesPersec", (value))
}

// GetAllowedClassifiesPersec gets the value of AllowedClassifiesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyAllowedClassifiesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllowedClassifiesPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetBlockedBinds sets the value of BlockedBinds for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyBlockedBinds(value uint32) (err error) {
	return instance.SetProperty("BlockedBinds", (value))
}

// GetBlockedBinds gets the value of BlockedBinds for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyBlockedBinds() (value uint32, err error) {
	retValue, err := instance.GetProperty("BlockedBinds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInboundConnections sets the value of InboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyInboundConnections(value uint32) (err error) {
	return instance.SetProperty("InboundConnections", (value))
}

// GetInboundConnections gets the value of InboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyInboundConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnections")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInboundConnectionsAllowedPersec sets the value of InboundConnectionsAllowedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyInboundConnectionsAllowedPersec(value uint32) (err error) {
	return instance.SetProperty("InboundConnectionsAllowedPersec", (value))
}

// GetInboundConnectionsAllowedPersec gets the value of InboundConnectionsAllowedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyInboundConnectionsAllowedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnectionsAllowedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInboundConnectionsBlockedPersec sets the value of InboundConnectionsBlockedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyInboundConnectionsBlockedPersec(value uint32) (err error) {
	return instance.SetProperty("InboundConnectionsBlockedPersec", (value))
}

// GetInboundConnectionsBlockedPersec gets the value of InboundConnectionsBlockedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyInboundConnectionsBlockedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnectionsBlockedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInboundPacketsDiscardedPersec sets the value of InboundPacketsDiscardedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyInboundPacketsDiscardedPersec(value uint32) (err error) {
	return instance.SetProperty("InboundPacketsDiscardedPersec", (value))
}

// GetInboundPacketsDiscardedPersec gets the value of InboundPacketsDiscardedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyInboundPacketsDiscardedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundPacketsDiscardedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetOutboundConnections sets the value of OutboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyOutboundConnections(value uint32) (err error) {
	return instance.SetProperty("OutboundConnections", (value))
}

// GetOutboundConnections gets the value of OutboundConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyOutboundConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundConnections")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetOutboundConnectionsAllowedPersec sets the value of OutboundConnectionsAllowedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyOutboundConnectionsAllowedPersec(value uint32) (err error) {
	return instance.SetProperty("OutboundConnectionsAllowedPersec", (value))
}

// GetOutboundConnectionsAllowedPersec gets the value of OutboundConnectionsAllowedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyOutboundConnectionsAllowedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundConnectionsAllowedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetOutboundConnectionsBlockedPersec sets the value of OutboundConnectionsBlockedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyOutboundConnectionsBlockedPersec(value uint32) (err error) {
	return instance.SetProperty("OutboundConnectionsBlockedPersec", (value))
}

// GetOutboundConnectionsBlockedPersec gets the value of OutboundConnectionsBlockedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyOutboundConnectionsBlockedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundConnectionsBlockedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetOutboundPacketsDiscardedPersec sets the value of OutboundPacketsDiscardedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyOutboundPacketsDiscardedPersec(value uint32) (err error) {
	return instance.SetProperty("OutboundPacketsDiscardedPersec", (value))
}

// GetOutboundPacketsDiscardedPersec gets the value of OutboundPacketsDiscardedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyOutboundPacketsDiscardedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundPacketsDiscardedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPacketsDiscardedPersec sets the value of PacketsDiscardedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) SetPropertyPacketsDiscardedPersec(value uint32) (err error) {
	return instance.SetProperty("PacketsDiscardedPersec", (value))
}

// GetPacketsDiscardedPersec gets the value of PacketsDiscardedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_WFPv4) GetPropertyPacketsDiscardedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsDiscardedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
