// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NSE4751BA1_9751_4F55_96B3_16813F2826DE
//////////////////////////////////////////////
package nse4751ba1_9751_4f55_96b3_16813f2826de

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
