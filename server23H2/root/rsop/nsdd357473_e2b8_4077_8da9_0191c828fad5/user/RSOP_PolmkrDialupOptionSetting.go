// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSDD357473_E2B8_4077_8DA9_0191C828FAD5.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrDialupOptionSetting struct
type RSOP_PolmkrDialupOptionSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrDialupOptionSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrDialupOptionSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrDialupOptionSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrDialupOptionSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrDialupOptionSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrDialupOptionSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
