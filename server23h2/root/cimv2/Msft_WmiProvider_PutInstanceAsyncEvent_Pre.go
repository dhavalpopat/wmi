// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msft_WmiProvider_PutInstanceAsyncEvent_Pre struct
type Msft_WmiProvider_PutInstanceAsyncEvent_Pre struct {
	*Msft_WmiProvider_OperationEvent_Pre

	//
	Flags uint32

	//
	InstanceObject interface{}
}

func NewMsft_WmiProvider_PutInstanceAsyncEvent_PreEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PreEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_PutInstanceAsyncEvent_Pre{
		Msft_WmiProvider_OperationEvent_Pre: tmp,
	}
	return
}

func NewMsft_WmiProvider_PutInstanceAsyncEvent_PreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PreEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_PutInstanceAsyncEvent_Pre{
		Msft_WmiProvider_OperationEvent_Pre: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInstanceObject sets the value of InstanceObject for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) SetPropertyInstanceObject(value interface{}) (err error) {
	return instance.SetProperty("InstanceObject", (value))
}

// GetInstanceObject gets the value of InstanceObject for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) GetPropertyInstanceObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("InstanceObject")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
