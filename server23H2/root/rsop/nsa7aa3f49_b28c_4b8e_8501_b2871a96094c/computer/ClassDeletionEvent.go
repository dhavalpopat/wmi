// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA7AA3F49_B28C_4B8E_8501_B2871A96094C.Computer
//////////////////////////////////////////////
package computer

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
