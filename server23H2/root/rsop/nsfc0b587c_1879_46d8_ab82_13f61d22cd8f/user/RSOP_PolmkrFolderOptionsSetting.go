// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSFC0B587C_1879_46D8_AB82_13F61D22CD8F.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrFolderOptionsSetting struct
type RSOP_PolmkrFolderOptionsSetting struct {
	*RSOP_PolmkrFolderOptionSetting
}

func NewRSOP_PolmkrFolderOptionsSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrFolderOptionsSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFolderOptionsSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrFolderOptionsSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrFolderOptionsSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFolderOptionsSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}
