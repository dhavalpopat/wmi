// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __EventConsumerProviderRegistration struct
type __EventConsumerProviderRegistration struct {
	*__ProviderRegistration

	//
	ConsumerClassNames []string
}

func New__EventConsumerProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__EventConsumerProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventConsumerProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__EventConsumerProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventConsumerProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventConsumerProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

// SetConsumerClassNames sets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) SetPropertyConsumerClassNames(value []string) (err error) {
	return instance.SetProperty("ConsumerClassNames", value)
}

// GetConsumerClassNames gets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) GetPropertyConsumerClassNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ConsumerClassNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
