// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS2F6DB41A_F363_4809_AAEE_1C61C7F5A2EB
//////////////////////////////////////////////
package ns2f6db41a_f363_4809_aaee_1c61c7f5a2eb

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ClassDeletion struct
type CIM_ClassDeletion struct {
	*CIM_ClassIndication
}

func NewCIM_ClassDeletionEx1(instance *cim.WmiInstance) (newInstance *CIM_ClassDeletion, err error) {
	tmp, err := NewCIM_ClassIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassDeletion{
		CIM_ClassIndication: tmp,
	}
	return
}

func NewCIM_ClassDeletionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ClassDeletion, err error) {
	tmp, err := NewCIM_ClassIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassDeletion{
		CIM_ClassIndication: tmp,
	}
	return
}
