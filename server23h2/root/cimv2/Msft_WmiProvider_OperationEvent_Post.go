// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_WmiProvider_OperationEvent_Post struct
type Msft_WmiProvider_OperationEvent_Post struct {
	*Msft_WmiProvider_OperationEvent
}

func NewMsft_WmiProvider_OperationEvent_PostEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_OperationEvent_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_OperationEvent_Post{
		Msft_WmiProvider_OperationEvent: tmp,
	}
	return
}

func NewMsft_WmiProvider_OperationEvent_PostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_OperationEvent_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_OperationEvent_Post{
		Msft_WmiProvider_OperationEvent: tmp,
	}
	return
}
