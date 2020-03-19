// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_AffectedJobElement struct
type CIM_AffectedJobElement struct {
	*cim.WmiInstance

	// The ManagedElement affected by the execution of the Job.
	AffectedElement CIM_ManagedElement

	// The Job that is affecting the ManagedElement.
	AffectingElement CIM_Job

	// An enumeration describing the 'effect' on the ManagedElement. This array corresponds to the OtherElementEffectsDescriptions array, where the latter provides details related to the high-level 'effects' enumerated by this property. Additional detail is required if the ElementEffects array contains the value 1, "Other".
	ElementEffects []AffectedJobElement_ElementEffects

	// Provides details for the 'effect' at the corresponding array position in ElementEffects. This information is required whenever ElementEffects contains the value 1 ("Other").
	OtherElementEffectsDescriptions []string
}

func NewCIM_AffectedJobElementEx1(instance *cim.WmiInstance) (newInstance *CIM_AffectedJobElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_AffectedJobElement{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_AffectedJobElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AffectedJobElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AffectedJobElement{
		WmiInstance: tmp,
	}
	return
}

// SetAffectedElement sets the value of AffectedElement for the instance
func (instance *CIM_AffectedJobElement) SetPropertyAffectedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("AffectedElement", value)
}

// GetAffectedElement gets the value of AffectedElement for the instance
func (instance *CIM_AffectedJobElement) GetPropertyAffectedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("AffectedElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAffectingElement sets the value of AffectingElement for the instance
func (instance *CIM_AffectedJobElement) SetPropertyAffectingElement(value CIM_Job) (err error) {
	return instance.SetProperty("AffectingElement", value)
}

// GetAffectingElement gets the value of AffectingElement for the instance
func (instance *CIM_AffectedJobElement) GetPropertyAffectingElement() (value CIM_Job, err error) {
	retValue, err := instance.GetProperty("AffectingElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Job)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElementEffects sets the value of ElementEffects for the instance
func (instance *CIM_AffectedJobElement) SetPropertyElementEffects(value []AffectedJobElement_ElementEffects) (err error) {
	return instance.SetProperty("ElementEffects", value)
}

// GetElementEffects gets the value of ElementEffects for the instance
func (instance *CIM_AffectedJobElement) GetPropertyElementEffects() (value []AffectedJobElement_ElementEffects, err error) {
	retValue, err := instance.GetProperty("ElementEffects")
	if err != nil {
		return
	}
	value, ok := retValue.([]AffectedJobElement_ElementEffects)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherElementEffectsDescriptions sets the value of OtherElementEffectsDescriptions for the instance
func (instance *CIM_AffectedJobElement) SetPropertyOtherElementEffectsDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherElementEffectsDescriptions", value)
}

// GetOtherElementEffectsDescriptions gets the value of OtherElementEffectsDescriptions for the instance
func (instance *CIM_AffectedJobElement) GetPropertyOtherElementEffectsDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherElementEffectsDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
