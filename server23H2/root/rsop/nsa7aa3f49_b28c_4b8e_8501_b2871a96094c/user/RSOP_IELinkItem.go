// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA7AA3F49_B28C_4B8E_8501_B2871A96094C.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IELinkItem struct
type RSOP_IELinkItem struct {
	*RSOP_IEFavoriteOrLinkItem

	//
	rsopID string

	//
	rsopPrecedence uint32
}

func NewRSOP_IELinkItemEx1(instance *cim.WmiInstance) (newInstance *RSOP_IELinkItem, err error) {
	tmp, err := NewRSOP_IEFavoriteOrLinkItemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IELinkItem{
		RSOP_IEFavoriteOrLinkItem: tmp,
	}
	return
}

func NewRSOP_IELinkItemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IELinkItem, err error) {
	tmp, err := NewRSOP_IEFavoriteOrLinkItemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IELinkItem{
		RSOP_IEFavoriteOrLinkItem: tmp,
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IELinkItem) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IELinkItem) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IELinkItem) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IELinkItem) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
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
