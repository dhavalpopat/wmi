// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ExpressionIdentifier struct
type MSFT_ExpressionIdentifier struct {
	*MSFT_Expression

	//
	name string
}

func NewMSFT_ExpressionIdentifierEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionIdentifier, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionIdentifier{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionIdentifierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionIdentifier, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionIdentifier{
		MSFT_Expression: tmp,
	}
	return
}

// Setname sets the value of name for the instance
func (instance *MSFT_ExpressionIdentifier) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", value)
}

// Getname gets the value of name for the instance
func (instance *MSFT_ExpressionIdentifier) GetPropertyname() (value string, err error) {
	retValue, err := instance.GetProperty("name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}