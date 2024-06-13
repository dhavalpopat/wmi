// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PageFault_V2 struct
type PageFault_V2 struct {
	*MSNT_SystemTrace
}

func NewPageFault_V2Ex1(instance *cim.WmiInstance) (newInstance *PageFault_V2, err error) {
	tmp, err := NewMSNT_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_V2{
		MSNT_SystemTrace: tmp,
	}
	return
}

func NewPageFault_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_V2, err error) {
	tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_V2{
		MSNT_SystemTrace: tmp,
	}
	return
}
