// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// UdpIp_V1 struct
type UdpIp_V1 struct {
	*MSNT_SystemTrace
}

func NewUdpIp_V1Ex1(instance *cim.WmiInstance) (newInstance *UdpIp_V1, err error) {
	tmp, err := NewMSNT_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UdpIp_V1{
		MSNT_SystemTrace: tmp,
	}
	return
}

func NewUdpIp_V1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UdpIp_V1, err error) {
	tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UdpIp_V1{
		MSNT_SystemTrace: tmp,
	}
	return
}
