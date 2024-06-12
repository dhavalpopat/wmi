// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS5B5202B1_2FC1_4BB0_A3B5_EB3B3EAE1626.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrNetworkOptionSetting struct
type RSOP_PolmkrNetworkOptionSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrNetworkOptionSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrNetworkOptionSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrNetworkOptionSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrNetworkOptionSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrNetworkOptionSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrNetworkOptionSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
