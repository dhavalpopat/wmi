// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ReserveCost struct
type Win32_ReserveCost struct {
	*CIM_Check

	//
	ReserveFolder string

	//
	ReserveKey string

	//
	ReserveLocal uint32

	//
	ReserveSource uint32
}

func NewWin32_ReserveCostEx1(instance *cim.WmiInstance) (newInstance *Win32_ReserveCost, err error) {
	tmp, err := NewCIM_CheckEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ReserveCost{
		CIM_Check: tmp,
	}
	return
}

func NewWin32_ReserveCostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ReserveCost, err error) {
	tmp, err := NewCIM_CheckEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ReserveCost{
		CIM_Check: tmp,
	}
	return
}

// SetReserveFolder sets the value of ReserveFolder for the instance
func (instance *Win32_ReserveCost) SetPropertyReserveFolder(value string) (err error) {
	return instance.SetProperty("ReserveFolder", value)
}

// GetReserveFolder gets the value of ReserveFolder for the instance
func (instance *Win32_ReserveCost) GetPropertyReserveFolder() (value string, err error) {
	retValue, err := instance.GetProperty("ReserveFolder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReserveKey sets the value of ReserveKey for the instance
func (instance *Win32_ReserveCost) SetPropertyReserveKey(value string) (err error) {
	return instance.SetProperty("ReserveKey", value)
}

// GetReserveKey gets the value of ReserveKey for the instance
func (instance *Win32_ReserveCost) GetPropertyReserveKey() (value string, err error) {
	retValue, err := instance.GetProperty("ReserveKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReserveLocal sets the value of ReserveLocal for the instance
func (instance *Win32_ReserveCost) SetPropertyReserveLocal(value uint32) (err error) {
	return instance.SetProperty("ReserveLocal", value)
}

// GetReserveLocal gets the value of ReserveLocal for the instance
func (instance *Win32_ReserveCost) GetPropertyReserveLocal() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReserveLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReserveSource sets the value of ReserveSource for the instance
func (instance *Win32_ReserveCost) SetPropertyReserveSource(value uint32) (err error) {
	return instance.SetProperty("ReserveSource", value)
}

// GetReserveSource gets the value of ReserveSource for the instance
func (instance *Win32_ReserveCost) GetPropertyReserveSource() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReserveSource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
