// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_WmiProviderEvent struct
type MSFT_WmiProviderEvent struct {
	*MSFT_WmiEssEvent

	//
	Namespace string

	//
	ProviderName string
}

func NewMSFT_WmiProviderEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_WmiProviderEvent, err error) {
	tmp, err := NewMSFT_WmiEssEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiProviderEvent{
		MSFT_WmiEssEvent: tmp,
	}
	return
}

func NewMSFT_WmiProviderEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WmiProviderEvent, err error) {
	tmp, err := NewMSFT_WmiEssEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiProviderEvent{
		MSFT_WmiEssEvent: tmp,
	}
	return
}

// SetNamespace sets the value of Namespace for the instance
func (instance *MSFT_WmiProviderEvent) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", value)
}

// GetNamespace gets the value of Namespace for the instance
func (instance *MSFT_WmiProviderEvent) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *MSFT_WmiProviderEvent) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", value)
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *MSFT_WmiProviderEvent) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
