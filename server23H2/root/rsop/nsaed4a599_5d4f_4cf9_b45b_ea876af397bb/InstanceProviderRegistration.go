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
