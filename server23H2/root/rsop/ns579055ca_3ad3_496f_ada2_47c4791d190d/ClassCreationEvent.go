// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS579055CA_3AD3_496F_ADA2_47C4791D190D
//////////////////////////////////////////////
package ns579055ca_3ad3_496f_ada2_47c4791d190d

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
