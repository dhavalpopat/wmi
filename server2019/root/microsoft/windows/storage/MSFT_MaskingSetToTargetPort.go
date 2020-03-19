// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MaskingSetToTargetPort struct
type MSFT_MaskingSetToTargetPort struct {
	*cim.WmiInstance

	//
	MaskingSet MSFT_MaskingSet

	//
	TargetPort MSFT_TargetPort
}

func NewMSFT_MaskingSetToTargetPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_MaskingSetToTargetPort, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MaskingSetToTargetPort{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MaskingSetToTargetPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MaskingSetToTargetPort, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MaskingSetToTargetPort{
		WmiInstance: tmp,
	}
	return
}

// SetMaskingSet sets the value of MaskingSet for the instance
func (instance *MSFT_MaskingSetToTargetPort) SetPropertyMaskingSet(value MSFT_MaskingSet) (err error) {
	return instance.SetProperty("MaskingSet", value)
}

// GetMaskingSet gets the value of MaskingSet for the instance
func (instance *MSFT_MaskingSetToTargetPort) GetPropertyMaskingSet() (value MSFT_MaskingSet, err error) {
	retValue, err := instance.GetProperty("MaskingSet")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_MaskingSet)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetPort sets the value of TargetPort for the instance
func (instance *MSFT_MaskingSetToTargetPort) SetPropertyTargetPort(value MSFT_TargetPort) (err error) {
	return instance.SetProperty("TargetPort", value)
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *MSFT_MaskingSetToTargetPort) GetPropertyTargetPort() (value MSFT_TargetPort, err error) {
	retValue, err := instance.GetProperty("TargetPort")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TargetPort)
	if !ok {
		// TODO: Set an error
	}
	return
}
