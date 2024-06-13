// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_PhysicalDiskToStorageReliabilityCounter struct
type SPACES_PhysicalDiskToStorageReliabilityCounter struct {
	*MSFT_PhysicalDiskToStorageReliabilityCounter
}

func NewSPACES_PhysicalDiskToStorageReliabilityCounterEx1(instance *cim.WmiInstance) (newInstance *SPACES_PhysicalDiskToStorageReliabilityCounter, err error) {
	tmp, err := NewMSFT_PhysicalDiskToStorageReliabilityCounterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_PhysicalDiskToStorageReliabilityCounter{
		MSFT_PhysicalDiskToStorageReliabilityCounter: tmp,
	}
	return
}

func NewSPACES_PhysicalDiskToStorageReliabilityCounterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_PhysicalDiskToStorageReliabilityCounter, err error) {
	tmp, err := NewMSFT_PhysicalDiskToStorageReliabilityCounterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_PhysicalDiskToStorageReliabilityCounter{
		MSFT_PhysicalDiskToStorageReliabilityCounter: tmp,
	}
	return
}
