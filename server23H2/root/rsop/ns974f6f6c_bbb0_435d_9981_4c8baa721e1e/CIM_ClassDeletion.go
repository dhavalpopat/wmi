// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS974F6F6C_BBB0_435D_9981_4C8BAA721E1E
//////////////////////////////////////////////
package ns974f6f6c_bbb0_435d_9981_4c8baa721e1e

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
