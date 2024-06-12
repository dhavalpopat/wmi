// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSFC0B587C_1879_46D8_AB82_13F61D22CD8F
//////////////////////////////////////////////
package nsfc0b587c_1879_46d8_ab82_13f61d22cd8f

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("ConsumerClassNames", (value))
}

// GetConsumerClassNames gets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) GetPropertyConsumerClassNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ConsumerClassNames")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
