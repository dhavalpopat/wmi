// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WFAPIGP struct
type WFAPIGP struct {
	*EventTrace

	// Enable Flags
	Flags WFAPIGP_Flags
}

func NewWFAPIGPEx1(instance *cim.WmiInstance) (newInstance *WFAPIGP, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WFAPIGP{
		EventTrace: tmp,
	}
	return
}

func NewWFAPIGPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WFAPIGP, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WFAPIGP{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *WFAPIGP) SetPropertyFlags(value WFAPIGP_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *WFAPIGP) GetPropertyFlags() (value WFAPIGP_Flags, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WFAPIGP_Flags(valuetmp)

	return
}
