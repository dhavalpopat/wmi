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

// WSP_StorageSubsystemToReplicationCapabilities struct
type WSP_StorageSubsystemToReplicationCapabilities struct {
	*MSFT_StorageSubSystemToReplicationCapabilities
}

func NewWSP_StorageSubsystemToReplicationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageSubsystemToReplicationCapabilities, err error) {
	tmp, err := NewMSFT_StorageSubSystemToReplicationCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToReplicationCapabilities{
		MSFT_StorageSubSystemToReplicationCapabilities: tmp,
	}
	return
}

func NewWSP_StorageSubsystemToReplicationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageSubsystemToReplicationCapabilities, err error) {
	tmp, err := NewMSFT_StorageSubSystemToReplicationCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToReplicationCapabilities{
		MSFT_StorageSubSystemToReplicationCapabilities: tmp,
	}
	return
}
