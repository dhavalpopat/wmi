// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS7344EDEA_49A9_449C_A6C2_37ECE9C9272F.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __MethodProviderRegistration struct
type __MethodProviderRegistration struct {
	*__ProviderRegistration
}

func New__MethodProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__MethodProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__MethodProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__MethodProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__MethodProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__MethodProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}
