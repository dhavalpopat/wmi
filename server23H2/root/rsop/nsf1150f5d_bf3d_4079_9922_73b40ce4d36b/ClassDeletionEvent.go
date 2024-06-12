// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSF1150F5D_BF3D_4079_9922_73B40CE4D36B
//////////////////////////////////////////////
package nsf1150f5d_bf3d_4079_9922_73b40ce4d36b

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ClassDeletionEvent struct
type __ClassDeletionEvent struct {
	*__ClassOperationEvent
}

func New__ClassDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__ClassDeletionEvent, err error) {
	tmp, err := New__ClassOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassDeletionEvent{
		__ClassOperationEvent: tmp,
	}
	return
}

func New__ClassDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassDeletionEvent, err error) {
	tmp, err := New__ClassOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassDeletionEvent{
		__ClassOperationEvent: tmp,
	}
	return
}
