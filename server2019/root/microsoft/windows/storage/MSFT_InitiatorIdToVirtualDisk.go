// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_InitiatorIdToVirtualDisk struct
type MSFT_InitiatorIdToVirtualDisk struct {
	*cim.WmiInstance

	//
	InitiatorId MSFT_InitiatorId

	//
	VirtualDisk MSFT_VirtualDisk
}

func NewMSFT_InitiatorIdToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_InitiatorIdToVirtualDisk, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_InitiatorIdToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_InitiatorIdToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_InitiatorIdToVirtualDisk, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_InitiatorIdToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

// SetInitiatorId sets the value of InitiatorId for the instance
func (instance *MSFT_InitiatorIdToVirtualDisk) SetPropertyInitiatorId(value MSFT_InitiatorId) (err error) {
	return instance.SetProperty("InitiatorId", value)
}

// GetInitiatorId gets the value of InitiatorId for the instance
func (instance *MSFT_InitiatorIdToVirtualDisk) GetPropertyInitiatorId() (value MSFT_InitiatorId, err error) {
	retValue, err := instance.GetProperty("InitiatorId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_InitiatorId)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_InitiatorIdToVirtualDisk) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_InitiatorIdToVirtualDisk) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("VirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}