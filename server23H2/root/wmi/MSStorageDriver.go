// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSStorageDriver struct
type MSStorageDriver struct {
	*cim.WmiInstance
}

func NewMSStorageDriverEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver{
		WmiInstance: tmp,
	}
	return
}

func NewMSStorageDriverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSStorageDriver, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver{
		WmiInstance: tmp,
	}
	return
}
