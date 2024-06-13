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

// MSNdis_CoTransmitPdusOk struct
type MSNdis_CoTransmitPdusOk struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	NdisCoTransmitPdusOk uint64
}

func NewMSNdis_CoTransmitPdusOkEx1(instance *cim.WmiInstance) (newInstance *MSNdis_CoTransmitPdusOk, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_CoTransmitPdusOk{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_CoTransmitPdusOkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_CoTransmitPdusOk, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_CoTransmitPdusOk{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_CoTransmitPdusOk) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_CoTransmitPdusOk) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_CoTransmitPdusOk) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_CoTransmitPdusOk) GetPropertyInstanceName() (value string, err error) {
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

// SetNdisCoTransmitPdusOk sets the value of NdisCoTransmitPdusOk for the instance
func (instance *MSNdis_CoTransmitPdusOk) SetPropertyNdisCoTransmitPdusOk(value uint64) (err error) {
	return instance.SetProperty("NdisCoTransmitPdusOk", (value))
}

// GetNdisCoTransmitPdusOk gets the value of NdisCoTransmitPdusOk for the instance
func (instance *MSNdis_CoTransmitPdusOk) GetPropertyNdisCoTransmitPdusOk() (value uint64, err error) {
	retValue, err := instance.GetProperty("NdisCoTransmitPdusOk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
