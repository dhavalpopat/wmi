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

// CIM_ProductSoftwareFeatures struct
type CIM_ProductSoftwareFeatures struct {
	*cim.WmiInstance

	//
	Component CIM_SoftwareFeature

	//
	Product CIM_Product
}

func NewCIM_ProductSoftwareFeaturesEx1(instance *cim.WmiInstance) (newInstance *CIM_ProductSoftwareFeatures, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ProductSoftwareFeatures{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ProductSoftwareFeaturesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProductSoftwareFeatures, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProductSoftwareFeatures{
		WmiInstance: tmp,
	}
	return
}

// SetComponent sets the value of Component for the instance
func (instance *CIM_ProductSoftwareFeatures) SetPropertyComponent(value CIM_SoftwareFeature) (err error) {
	return instance.SetProperty("Component", value)
}

// GetComponent gets the value of Component for the instance
func (instance *CIM_ProductSoftwareFeatures) GetPropertyComponent() (value CIM_SoftwareFeature, err error) {
	retValue, err := instance.GetProperty("Component")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SoftwareFeature)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProduct sets the value of Product for the instance
func (instance *CIM_ProductSoftwareFeatures) SetPropertyProduct(value CIM_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_ProductSoftwareFeatures) GetPropertyProduct() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("Product")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}