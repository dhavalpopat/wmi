// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceStatusSuccess struct
type MSFT_NetServiceStatusSuccess struct {
	MSFT_SCMEventLogEvent

	//
	Control string

	//
	Service string
}

// SetControl sets the value of Control for the instance
func (instance *MSFT_NetServiceStatusSuccess) SetPropertyControl(value string) (err error) {
	return instance.SetProperty("Control", value)
}

// GetControl gets the value of Control for the instance
func (instance *MSFT_NetServiceStatusSuccess) GetPropertyControl() (value string, err error) {
	retValue, err := instance.GetProperty("Control")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceStatusSuccess) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceStatusSuccess) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
