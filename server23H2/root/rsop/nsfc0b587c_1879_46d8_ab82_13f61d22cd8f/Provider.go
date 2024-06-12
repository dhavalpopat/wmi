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

// __Provider struct
type __Provider struct {
	*__SystemClass

	//
	Name string
}

func New__ProviderEx1(instance *cim.WmiInstance) (newInstance *__Provider, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__Provider{
		__SystemClass: tmp,
	}
	return
}

func New__ProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__Provider, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__Provider{
		__SystemClass: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *__Provider) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *__Provider) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
