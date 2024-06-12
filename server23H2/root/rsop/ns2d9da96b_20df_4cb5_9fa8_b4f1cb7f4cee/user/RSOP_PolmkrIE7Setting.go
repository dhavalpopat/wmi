// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS2D9DA96B_20DF_4CB5_9FA8_B4F1CB7F4CEE.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrIE7Setting struct
type RSOP_PolmkrIE7Setting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrIE7SettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrIE7Setting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrIE7Setting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrIE7SettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrIE7Setting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrIE7Setting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
