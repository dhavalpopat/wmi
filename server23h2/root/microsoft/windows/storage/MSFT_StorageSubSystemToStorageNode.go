// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageSubSystemToStorageNode struct
type MSFT_StorageSubSystemToStorageNode struct {
	*cim.WmiInstance

	//
	StorageNode MSFT_StorageNode

	//
	StorageSubSystem MSFT_StorageSubSystem
}

func NewMSFT_StorageSubSystemToStorageNodeEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageSubSystemToStorageNode, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystemToStorageNode{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageSubSystemToStorageNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageSubSystemToStorageNode, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystemToStorageNode{
		WmiInstance: tmp,
	}
	return
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", (value))
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
	retValue, err := instance.GetProperty("StorageNode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_StorageNode)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_StorageNode is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_StorageNode(valuetmp)

	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", (value))
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
	retValue, err := instance.GetProperty("StorageSubSystem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_StorageSubSystem)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_StorageSubSystem is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_StorageSubSystem(valuetmp)

	return
}
