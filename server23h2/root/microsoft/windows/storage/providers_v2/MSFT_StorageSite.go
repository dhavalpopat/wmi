// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageSite struct
type MSFT_StorageSite struct {
	*MSFT_StorageFaultDomain
}

func NewMSFT_StorageSiteEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageSite, err error) {
	tmp, err := NewMSFT_StorageFaultDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSite{
		MSFT_StorageFaultDomain: tmp,
	}
	return
}

func NewMSFT_StorageSiteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageSite, err error) {
	tmp, err := NewMSFT_StorageFaultDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSite{
		MSFT_StorageFaultDomain: tmp,
	}
	return
}
