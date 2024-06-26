// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VideoBIOSFeature struct
type CIM_VideoBIOSFeature struct {
	*CIM_SoftwareFeature

	//
	CharacteristicDescriptions []string

	//
	Characteristics []uint16
}

func NewCIM_VideoBIOSFeatureEx1(instance *cim.WmiInstance) (newInstance *CIM_VideoBIOSFeature, err error) {
	tmp, err := NewCIM_SoftwareFeatureEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VideoBIOSFeature{
		CIM_SoftwareFeature: tmp,
	}
	return
}

func NewCIM_VideoBIOSFeatureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VideoBIOSFeature, err error) {
	tmp, err := NewCIM_SoftwareFeatureEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VideoBIOSFeature{
		CIM_SoftwareFeature: tmp,
	}
	return
}

// SetCharacteristicDescriptions sets the value of CharacteristicDescriptions for the instance
func (instance *CIM_VideoBIOSFeature) SetPropertyCharacteristicDescriptions(value []string) (err error) {
	return instance.SetProperty("CharacteristicDescriptions", (value))
}

// GetCharacteristicDescriptions gets the value of CharacteristicDescriptions for the instance
func (instance *CIM_VideoBIOSFeature) GetPropertyCharacteristicDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CharacteristicDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *CIM_VideoBIOSFeature) SetPropertyCharacteristics(value []uint16) (err error) {
	return instance.SetProperty("Characteristics", (value))
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *CIM_VideoBIOSFeature) GetPropertyCharacteristics() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Characteristics")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}
