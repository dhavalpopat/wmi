// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS4C15617D_7EC1_44F0_9878_0B4CCDB5F690
//////////////////////////////////////////////
package ns4c15617d_7ec1_44f0_9878_0b4ccdb5f690

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
