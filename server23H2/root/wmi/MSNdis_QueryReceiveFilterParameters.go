// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_QueryReceiveFilterParameters struct
type MSNdis_QueryReceiveFilterParameters struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string
}

func NewMSNdis_QueryReceiveFilterParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_QueryReceiveFilterParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_QueryReceiveFilterParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_QueryReceiveFilterParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_QueryReceiveFilterParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_QueryReceiveFilterParameters{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_QueryReceiveFilterParameters) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_QueryReceiveFilterParameters) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_QueryReceiveFilterParameters) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_QueryReceiveFilterParameters) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

//

// <param name="Header" type="MSNdis_WmiMethodHeader "></param>
// <param name="ReceiveFilterParameters" type="MSNdis_ReceiveFilterParameters "></param>

// <param name="ReceiveFilterParameters" type="MSNdis_ReceiveFilterParameters "></param>
func (instance *MSNdis_QueryReceiveFilterParameters) WmiQueryReceiveFilterParameters( /* IN */ Header MSNdis_WmiMethodHeader,
	/* IN/OUT */ ReceiveFilterParameters MSNdis_ReceiveFilterParameters) (err error) {
	_, err = instance.InvokeMethod("WmiQueryReceiveFilterParameters", Header)
	if err != nil {
		return
	}
	return

}
