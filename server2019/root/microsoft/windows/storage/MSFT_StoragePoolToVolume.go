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

// MSFT_StoragePoolToVolume struct
type MSFT_StoragePoolToVolume struct {
	cim.WmiInstance

	//
	StoragePool MSFT_StoragePool

	//
	Volume MSFT_Volume
}

// SetStoragePool sets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToVolume) SetPropertyStoragePool(value MSFT_StoragePool) (err error) {
	return instance.SetProperty("StoragePool", value)
}

// GetStoragePool gets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToVolume) GetPropertyStoragePool() (value MSFT_StoragePool, err error) {
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

// SetVolume sets the value of Volume for the instance
func (instance *MSFT_StoragePoolToVolume) SetPropertyVolume(value MSFT_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *MSFT_StoragePoolToVolume) GetPropertyVolume() (value MSFT_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}
