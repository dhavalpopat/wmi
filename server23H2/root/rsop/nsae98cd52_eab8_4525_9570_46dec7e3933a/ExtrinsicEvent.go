// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSAE98CD52_EAB8_4525_9570_46DEC7E3933A
//////////////////////////////////////////////
package nsae98cd52_eab8_4525_9570_46dec7e3933a

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ExtrinsicEvent struct
type __ExtrinsicEvent struct {
	*__Event
}

func New__ExtrinsicEventEx1(instance *cim.WmiInstance) (newInstance *__ExtrinsicEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ExtrinsicEvent{
		__Event: tmp,
	}
	return
}

func New__ExtrinsicEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ExtrinsicEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ExtrinsicEvent{
		__Event: tmp,
	}
	return
}
