// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSF8393E37_A916_40B8_B064_5A700BB71C87.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_InstCreation struct
type CIM_InstCreation struct {
	*CIM_InstIndication
}

func NewCIM_InstCreationEx1(instance *cim.WmiInstance) (newInstance *CIM_InstCreation, err error) {
	tmp, err := NewCIM_InstIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_InstCreation{
		CIM_InstIndication: tmp,
	}
	return
}

func NewCIM_InstCreationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_InstCreation, err error) {
	tmp, err := NewCIM_InstIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_InstCreation{
		CIM_InstIndication: tmp,
	}
	return
}
