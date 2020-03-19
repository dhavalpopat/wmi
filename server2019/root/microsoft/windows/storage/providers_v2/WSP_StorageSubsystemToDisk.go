// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_StorageSubsystemToDisk struct
type WSP_StorageSubsystemToDisk struct {
	*MSFT_StorageSubSystemToDisk
}

func NewWSP_StorageSubsystemToDiskEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageSubsystemToDisk, err error) {
	tmp, err := NewMSFT_StorageSubSystemToDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToDisk{
		MSFT_StorageSubSystemToDisk: tmp,
	}
	return
}

func NewWSP_StorageSubsystemToDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageSubsystemToDisk, err error) {
	tmp, err := NewMSFT_StorageSubSystemToDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToDisk{
		MSFT_StorageSubSystemToDisk: tmp,
	}
	return
}
