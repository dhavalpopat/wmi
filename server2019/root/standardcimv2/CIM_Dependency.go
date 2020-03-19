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

// CIM_Dependency struct
type CIM_Dependency struct {
	*cim.WmiInstance

	//
	Antecedent CIM_ManagedElement

	//
	Dependent CIM_ManagedElement
}

func NewCIM_DependencyEx1(instance *cim.WmiInstance) (newInstance *CIM_Dependency, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_Dependency{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_DependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Dependency, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Dependency{
		WmiInstance: tmp,
	}
	return
}

// SetAntecedent sets the value of Antecedent for the instance
func (instance *CIM_Dependency) SetPropertyAntecedent(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("Antecedent", value)
}

// GetAntecedent gets the value of Antecedent for the instance
func (instance *CIM_Dependency) GetPropertyAntecedent() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("Antecedent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependent sets the value of Dependent for the instance
func (instance *CIM_Dependency) SetPropertyDependent(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("Dependent", value)
}

// GetDependent gets the value of Dependent for the instance
func (instance *CIM_Dependency) GetPropertyDependent() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("Dependent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
