// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// EffectiveWindows8App struct
type EffectiveWindows8App struct {
	*cim.WmiInstance

	//
	DisplayName string

	//
	Enabled bool

	//
	EnabledSource string

	//
	PackageFamilyName string
}

func NewEffectiveWindows8AppEx1(instance *cim.WmiInstance) (newInstance *EffectiveWindows8App, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &EffectiveWindows8App{
		WmiInstance: tmp,
	}
	return
}

func NewEffectiveWindows8AppEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *EffectiveWindows8App, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &EffectiveWindows8App{
		WmiInstance: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *EffectiveWindows8App) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *EffectiveWindows8App) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *EffectiveWindows8App) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *EffectiveWindows8App) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledSource sets the value of EnabledSource for the instance
func (instance *EffectiveWindows8App) SetPropertyEnabledSource(value string) (err error) {
	return instance.SetProperty("EnabledSource", value)
}

// GetEnabledSource gets the value of EnabledSource for the instance
func (instance *EffectiveWindows8App) GetPropertyEnabledSource() (value string, err error) {
	retValue, err := instance.GetProperty("EnabledSource")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageFamilyName sets the value of PackageFamilyName for the instance
func (instance *EffectiveWindows8App) SetPropertyPackageFamilyName(value string) (err error) {
	return instance.SetProperty("PackageFamilyName", value)
}

// GetPackageFamilyName gets the value of PackageFamilyName for the instance
func (instance *EffectiveWindows8App) GetPropertyPackageFamilyName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFamilyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
