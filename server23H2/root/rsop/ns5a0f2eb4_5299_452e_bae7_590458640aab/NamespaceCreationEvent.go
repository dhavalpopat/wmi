// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS5A0F2EB4_5299_452E_BAE7_590458640AAB
//////////////////////////////////////////////
package ns5a0f2eb4_5299_452e_bae7_590458640aab

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceCreationEvent struct
type __NamespaceCreationEvent struct {
	*__NamespaceOperationEvent
}

func New__NamespaceCreationEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceCreationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceCreationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceCreationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceCreationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}
