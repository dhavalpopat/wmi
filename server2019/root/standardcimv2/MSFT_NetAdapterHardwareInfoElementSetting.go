// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterHardwareInfoElementSetting struct
type MSFT_NetAdapterHardwareInfoElementSetting struct {
	*MSFT_NetAdapterElementSettingData
}

func NewMSFT_NetAdapterHardwareInfoElementSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterHardwareInfoElementSetting, err error) {
	tmp, err := NewMSFT_NetAdapterElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterHardwareInfoElementSetting{
		MSFT_NetAdapterElementSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterHardwareInfoElementSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterHardwareInfoElementSetting, err error) {
	tmp, err := NewMSFT_NetAdapterElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterHardwareInfoElementSetting{
		MSFT_NetAdapterElementSettingData: tmp,
	}
	return
}
