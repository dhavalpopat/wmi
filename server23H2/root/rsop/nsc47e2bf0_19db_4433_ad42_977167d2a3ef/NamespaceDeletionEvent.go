// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSC47E2BF0_19DB_4433_AD42_977167D2A3EF
//////////////////////////////////////////////
package nsc47e2bf0_19db_4433_ad42_977167d2a3ef

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceDeletionEvent struct
type __NamespaceDeletionEvent struct {
	*__NamespaceOperationEvent
}

func New__NamespaceDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceDeletionEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceDeletionEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceDeletionEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceDeletionEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}
