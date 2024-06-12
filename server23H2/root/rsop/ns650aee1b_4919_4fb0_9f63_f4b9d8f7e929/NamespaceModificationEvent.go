// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS650AEE1B_4919_4FB0_9F63_F4B9D8F7E929
//////////////////////////////////////////////
package ns650aee1b_4919_4fb0_9f63_f4b9d8f7e929

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __NamespaceModificationEvent struct
type __NamespaceModificationEvent struct {
	*__NamespaceOperationEvent

	//
	PreviousNamespace __Namespace
}

func New__NamespaceModificationEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceModificationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceModificationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceModificationEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceModificationEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

// SetPreviousNamespace sets the value of PreviousNamespace for the instance
func (instance *__NamespaceModificationEvent) SetPropertyPreviousNamespace(value __Namespace) (err error) {
	return instance.SetProperty("PreviousNamespace", (value))
}

// GetPreviousNamespace gets the value of PreviousNamespace for the instance
func (instance *__NamespaceModificationEvent) GetPropertyPreviousNamespace() (value __Namespace, err error) {
	retValue, err := instance.GetProperty("PreviousNamespace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(__Namespace)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " __Namespace is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = __Namespace(valuetmp)

	return
}
