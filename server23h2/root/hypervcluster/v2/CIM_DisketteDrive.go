// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_DisketteDrive struct
type CIM_DisketteDrive struct {
	*CIM_MediaAccessDevice
}

func NewCIM_DisketteDriveEx1(instance *cim.WmiInstance) (newInstance *CIM_DisketteDrive, err error) {
	tmp, err := NewCIM_MediaAccessDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DisketteDrive{
		CIM_MediaAccessDevice: tmp,
	}
	return
}

func NewCIM_DisketteDriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DisketteDrive, err error) {
	tmp, err := NewCIM_MediaAccessDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DisketteDrive{
		CIM_MediaAccessDevice: tmp,
	}
	return
}
