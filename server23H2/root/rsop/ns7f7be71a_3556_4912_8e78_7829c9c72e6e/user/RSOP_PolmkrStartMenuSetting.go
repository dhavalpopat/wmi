// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS7F7BE71A_3556_4912_8E78_7829C9C72E6E.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrStartMenuSetting struct
type RSOP_PolmkrStartMenuSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrStartMenuSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrStartMenuSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrStartMenuSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrStartMenuSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrStartMenuSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrStartMenuSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
