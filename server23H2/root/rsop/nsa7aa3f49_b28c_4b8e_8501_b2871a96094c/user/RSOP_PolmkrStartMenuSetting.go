// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA7AA3F49_B28C_4B8E_8501_B2871A96094C.User
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
