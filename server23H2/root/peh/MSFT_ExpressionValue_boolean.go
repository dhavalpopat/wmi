// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ExpressionValue_boolean struct
type MSFT_ExpressionValue_boolean struct {
	*MSFT_ExpressionValue

	//
	value bool
}

func NewMSFT_ExpressionValue_booleanEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionValue_boolean, err error) {
	tmp, err := NewMSFT_ExpressionValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue_boolean{
		MSFT_ExpressionValue: tmp,
	}
	return
}

func NewMSFT_ExpressionValue_booleanEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionValue_boolean, err error) {
	tmp, err := NewMSFT_ExpressionValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue_boolean{
		MSFT_ExpressionValue: tmp,
	}
	return
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_boolean) SetPropertyvalue(value bool) (err error) {
	return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_boolean) GetPropertyvalue() (value bool, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
