// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS814BDA2A_0BB7_4B36_B7B7_0A0141C4C544
//////////////////////////////////////////////
package ns814bda2a_0bb7_4b36_b7b7_0a0141c4c544

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __NamespaceOperationEvent struct
type __NamespaceOperationEvent struct {
	*__Event

	//
	TargetNamespace __Namespace
}

func New__NamespaceOperationEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceOperationEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceOperationEvent{
		__Event: tmp,
	}
	return
}

func New__NamespaceOperationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceOperationEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceOperationEvent{
		__Event: tmp,
	}
	return
}

// SetTargetNamespace sets the value of TargetNamespace for the instance
func (instance *__NamespaceOperationEvent) SetPropertyTargetNamespace(value __Namespace) (err error) {
	return instance.SetProperty("TargetNamespace", (value))
}

// GetTargetNamespace gets the value of TargetNamespace for the instance
func (instance *__NamespaceOperationEvent) GetPropertyTargetNamespace() (value __Namespace, err error) {
	retValue, err := instance.GetProperty("TargetNamespace")
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
