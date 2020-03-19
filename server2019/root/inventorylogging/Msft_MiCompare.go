// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_MiCompare struct
type Msft_MiCompare struct {
	*Msft_MiStream

	//
	Input Msft_MiStream

	//
	OnlyUpdateSnapshot bool

	//
	SuppressionHint Msft_MiCompareSuppression
}

func NewMsft_MiCompareEx1(instance *cim.WmiInstance) (newInstance *Msft_MiCompare, err error) {
	tmp, err := NewMsft_MiStreamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_MiCompare{
		Msft_MiStream: tmp,
	}
	return
}

func NewMsft_MiCompareEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiCompare, err error) {
	tmp, err := NewMsft_MiStreamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiCompare{
		Msft_MiStream: tmp,
	}
	return
}

// SetInput sets the value of Input for the instance
func (instance *Msft_MiCompare) SetPropertyInput(value Msft_MiStream) (err error) {
	return instance.SetProperty("Input", value)
}

// GetInput gets the value of Input for the instance
func (instance *Msft_MiCompare) GetPropertyInput() (value Msft_MiStream, err error) {
	retValue, err := instance.GetProperty("Input")
	if err != nil {
		return
	}
	value, ok := retValue.(Msft_MiStream)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOnlyUpdateSnapshot sets the value of OnlyUpdateSnapshot for the instance
func (instance *Msft_MiCompare) SetPropertyOnlyUpdateSnapshot(value bool) (err error) {
	return instance.SetProperty("OnlyUpdateSnapshot", value)
}

// GetOnlyUpdateSnapshot gets the value of OnlyUpdateSnapshot for the instance
func (instance *Msft_MiCompare) GetPropertyOnlyUpdateSnapshot() (value bool, err error) {
	retValue, err := instance.GetProperty("OnlyUpdateSnapshot")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuppressionHint sets the value of SuppressionHint for the instance
func (instance *Msft_MiCompare) SetPropertySuppressionHint(value Msft_MiCompareSuppression) (err error) {
	return instance.SetProperty("SuppressionHint", value)
}

// GetSuppressionHint gets the value of SuppressionHint for the instance
func (instance *Msft_MiCompare) GetPropertySuppressionHint() (value Msft_MiCompareSuppression, err error) {
	retValue, err := instance.GetProperty("SuppressionHint")
	if err != nil {
		return
	}
	value, ok := retValue.(Msft_MiCompareSuppression)
	if !ok {
		// TODO: Set an error
	}
	return
}