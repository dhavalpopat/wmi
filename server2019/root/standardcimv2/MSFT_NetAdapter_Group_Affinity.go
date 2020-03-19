// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_Group_Affinity struct
type MSFT_NetAdapter_Group_Affinity struct {
	*cim.WmiInstance

	//
	ProcessorAffinityMask uint64

	//
	ProcessorGroup uint16
}

func NewMSFT_NetAdapter_Group_AffinityEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_Group_Affinity, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_Group_Affinity{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_Group_AffinityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_Group_Affinity, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_Group_Affinity{
		WmiInstance: tmp,
	}
	return
}

// SetProcessorAffinityMask sets the value of ProcessorAffinityMask for the instance
func (instance *MSFT_NetAdapter_Group_Affinity) SetPropertyProcessorAffinityMask(value uint64) (err error) {
	return instance.SetProperty("ProcessorAffinityMask", value)
}

// GetProcessorAffinityMask gets the value of ProcessorAffinityMask for the instance
func (instance *MSFT_NetAdapter_Group_Affinity) GetPropertyProcessorAffinityMask() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProcessorAffinityMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorGroup sets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_Group_Affinity) SetPropertyProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("ProcessorGroup", value)
}

// GetProcessorGroup gets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_Group_Affinity) GetPropertyProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
