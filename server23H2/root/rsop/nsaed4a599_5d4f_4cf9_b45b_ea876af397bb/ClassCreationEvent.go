// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSAED4A599_5D4F_4CF9_B45B_EA876AF397BB
//////////////////////////////////////////////
package nsaed4a599_5d4f_4cf9_b45b_ea876af397bb

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ClassCreationEvent struct
type __ClassCreationEvent struct {
	*__ClassOperationEvent
}

func New__ClassCreationEventEx1(instance *cim.WmiInstance) (newInstance *__ClassCreationEvent, err error) {
	tmp, err := New__ClassOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassCreationEvent{
		__ClassOperationEvent: tmp,
	}
	return
}

func New__ClassCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassCreationEvent, err error) {
	tmp, err := New__ClassOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassCreationEvent{
		__ClassOperationEvent: tmp,
	}
	return
}
