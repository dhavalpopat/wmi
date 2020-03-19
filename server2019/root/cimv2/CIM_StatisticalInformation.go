// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_StatisticalInformation struct
type CIM_StatisticalInformation struct {
	*cim.WmiInstance

	//
	Caption string

	//
	Description string

	//
	Name string
}

func NewCIM_StatisticalInformationEx1(instance *cim.WmiInstance) (newInstance *CIM_StatisticalInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_StatisticalInformation{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_StatisticalInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_StatisticalInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_StatisticalInformation{
		WmiInstance: tmp,
	}
	return
}

// SetCaption sets the value of Caption for the instance
func (instance *CIM_StatisticalInformation) SetPropertyCaption(value string) (err error) {
	return instance.SetProperty("Caption", value)
}

// GetCaption gets the value of Caption for the instance
func (instance *CIM_StatisticalInformation) GetPropertyCaption() (value string, err error) {
	retValue, err := instance.GetProperty("Caption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *CIM_StatisticalInformation) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_StatisticalInformation) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *CIM_StatisticalInformation) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *CIM_StatisticalInformation) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
