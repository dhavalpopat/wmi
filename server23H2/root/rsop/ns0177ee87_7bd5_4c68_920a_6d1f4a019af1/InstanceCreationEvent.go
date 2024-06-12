// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS0177EE87_7BD5_4C68_920A_6D1F4A019AF1
//////////////////////////////////////////////
package ns0177ee87_7bd5_4c68_920a_6d1f4a019af1

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceCreationEvent struct
type __InstanceCreationEvent struct {
	*__InstanceOperationEvent
}

func New__InstanceCreationEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceCreationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceCreationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

func New__InstanceCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceCreationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceCreationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}
