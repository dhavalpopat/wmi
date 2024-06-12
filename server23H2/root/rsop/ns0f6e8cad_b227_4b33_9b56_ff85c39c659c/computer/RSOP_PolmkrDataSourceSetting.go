// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS0F6E8CAD_B227_4B33_9B56_FF85C39C659C.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrDataSourceSetting struct
type RSOP_PolmkrDataSourceSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrDataSourceSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrDataSourceSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrDataSourceSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrDataSourceSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrDataSourceSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrDataSourceSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
