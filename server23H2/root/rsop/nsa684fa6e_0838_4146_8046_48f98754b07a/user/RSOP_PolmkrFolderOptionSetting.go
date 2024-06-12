// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA684FA6E_0838_4146_8046_48F98754B07A.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrFolderOptionSetting struct
type RSOP_PolmkrFolderOptionSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrFolderOptionSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrFolderOptionSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFolderOptionSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrFolderOptionSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrFolderOptionSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFolderOptionSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
