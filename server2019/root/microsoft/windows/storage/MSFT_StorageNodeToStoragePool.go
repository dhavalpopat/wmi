// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageNodeToStoragePool struct
type MSFT_StorageNodeToStoragePool struct {
	cim.WmiInstance

	//
	StorageNode MSFT_StorageNode

	//
	StoragePool MSFT_StoragePool
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToStoragePool) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", value)
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToStoragePool) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
	retValue, err := instance.GetProperty("StorageNode")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageNode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoragePool sets the value of StoragePool for the instance
func (instance *MSFT_StorageNodeToStoragePool) SetPropertyStoragePool(value MSFT_StoragePool) (err error) {
	return instance.SetProperty("StoragePool", value)
}

// GetStoragePool gets the value of StoragePool for the instance
func (instance *MSFT_StorageNodeToStoragePool) GetPropertyStoragePool() (value MSFT_StoragePool, err error) {
	retValue, err := instance.GetProperty("StoragePool")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StoragePool)
	if !ok {
		// TODO: Set an error
	}
	return
}
