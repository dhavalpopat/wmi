// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NSE31A6F0B_2C62_4FC1_A177_9F88DF963F17
//////////////////////////////////////////////
package nse31a6f0b_2c62_4fc1_a177_9f88df963f17

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceProviderRegistration struct
type __InstanceProviderRegistration struct {
	*__ObjectProviderRegistration
}

func New__InstanceProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__InstanceProviderRegistration, err error) {
	tmp, err := New__ObjectProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceProviderRegistration{
		__ObjectProviderRegistration: tmp,
	}
	return
}

func New__InstanceProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceProviderRegistration, err error) {
	tmp, err := New__ObjectProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceProviderRegistration{
		__ObjectProviderRegistration: tmp,
	}
	return
}
