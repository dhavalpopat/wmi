// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS77DB142C_6A49_489B_8538_882E7276E951.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrSchedTaskSetting struct
type RSOP_PolmkrSchedTaskSetting struct {
	*RSOP_PolmkrTaskSetting
}

func NewRSOP_PolmkrSchedTaskSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrSchedTaskSetting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrSchedTaskSetting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrSchedTaskSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrSchedTaskSetting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrSchedTaskSetting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}
