// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterVmqElementSetting struct
type MSFT_NetAdapterVmqElementSetting struct {
	*MSFT_NetAdapterElementSettingData
}

func NewMSFT_NetAdapterVmqElementSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterVmqElementSetting, err error) {
	tmp, err := NewMSFT_NetAdapterElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterVmqElementSetting{
		MSFT_NetAdapterElementSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterVmqElementSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterVmqElementSetting, err error) {
	tmp, err := NewMSFT_NetAdapterElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterVmqElementSetting{
		MSFT_NetAdapterElementSettingData: tmp,
	}
	return
}
