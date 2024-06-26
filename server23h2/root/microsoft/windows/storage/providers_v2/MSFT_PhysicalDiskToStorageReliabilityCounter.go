// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_PhysicalDiskToStorageReliabilityCounter struct
type MSFT_PhysicalDiskToStorageReliabilityCounter struct {
	*cim.WmiInstance

	//
	PhysicalDisk MSFT_PhysicalDisk

	//
	StorageReliabilityCounter MSFT_StorageReliabilityCounter
}

func NewMSFT_PhysicalDiskToStorageReliabilityCounterEx1(instance *cim.WmiInstance) (newInstance *MSFT_PhysicalDiskToStorageReliabilityCounter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_PhysicalDiskToStorageReliabilityCounter{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_PhysicalDiskToStorageReliabilityCounterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PhysicalDiskToStorageReliabilityCounter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PhysicalDiskToStorageReliabilityCounter{
		WmiInstance: tmp,
	}
	return
}

// SetPhysicalDisk sets the value of PhysicalDisk for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) SetPropertyPhysicalDisk(value MSFT_PhysicalDisk) (err error) {
	return instance.SetProperty("PhysicalDisk", (value))
}

// GetPhysicalDisk gets the value of PhysicalDisk for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) GetPropertyPhysicalDisk() (value MSFT_PhysicalDisk, err error) {
	retValue, err := instance.GetProperty("PhysicalDisk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_PhysicalDisk)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_PhysicalDisk is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_PhysicalDisk(valuetmp)

	return
}

// SetStorageReliabilityCounter sets the value of StorageReliabilityCounter for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) SetPropertyStorageReliabilityCounter(value MSFT_StorageReliabilityCounter) (err error) {
	return instance.SetProperty("StorageReliabilityCounter", (value))
}

// GetStorageReliabilityCounter gets the value of StorageReliabilityCounter for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) GetPropertyStorageReliabilityCounter() (value MSFT_StorageReliabilityCounter, err error) {
	retValue, err := instance.GetProperty("StorageReliabilityCounter")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_StorageReliabilityCounter)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_StorageReliabilityCounter is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_StorageReliabilityCounter(valuetmp)

	return
}
