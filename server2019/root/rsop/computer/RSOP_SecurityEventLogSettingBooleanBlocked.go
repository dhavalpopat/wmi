// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_SecurityEventLogSettingBooleanBlocked struct
type RSOP_SecurityEventLogSettingBooleanBlocked struct {
	*RSOP_SecuritySettingsBlocked

	//
	KeyName string

	//
	Setting bool

	//
	Type string
}

func NewRSOP_SecurityEventLogSettingBooleanBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_SecurityEventLogSettingBooleanBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecurityEventLogSettingBooleanBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

func NewRSOP_SecurityEventLogSettingBooleanBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_SecurityEventLogSettingBooleanBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecurityEventLogSettingBooleanBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

// SetKeyName sets the value of KeyName for the instance
func (instance *RSOP_SecurityEventLogSettingBooleanBlocked) SetPropertyKeyName(value string) (err error) {
	return instance.SetProperty("KeyName", value)
}

// GetKeyName gets the value of KeyName for the instance
func (instance *RSOP_SecurityEventLogSettingBooleanBlocked) GetPropertyKeyName() (value string, err error) {
	retValue, err := instance.GetProperty("KeyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *RSOP_SecurityEventLogSettingBooleanBlocked) SetPropertySetting(value bool) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *RSOP_SecurityEventLogSettingBooleanBlocked) GetPropertySetting() (value bool, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *RSOP_SecurityEventLogSettingBooleanBlocked) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *RSOP_SecurityEventLogSettingBooleanBlocked) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
