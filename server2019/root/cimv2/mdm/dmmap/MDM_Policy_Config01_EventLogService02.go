// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Config01_EventLogService02 struct
type MDM_Policy_Config01_EventLogService02 struct {
	cim.WmiInstance

	//
	ControlEventLogBehavior string

	//
	InstanceID string

	//
	ParentID string

	//
	SpecifyMaximumFileSizeApplicationLog string

	//
	SpecifyMaximumFileSizeSecurityLog string

	//
	SpecifyMaximumFileSizeSystemLog string
}

// SetControlEventLogBehavior sets the value of ControlEventLogBehavior for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertyControlEventLogBehavior(value string) (err error) {
	return instance.SetProperty("ControlEventLogBehavior", value)
}

// GetControlEventLogBehavior gets the value of ControlEventLogBehavior for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertyControlEventLogBehavior() (value string, err error) {
	retValue, err := instance.GetProperty("ControlEventLogBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyMaximumFileSizeApplicationLog sets the value of SpecifyMaximumFileSizeApplicationLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertySpecifyMaximumFileSizeApplicationLog(value string) (err error) {
	return instance.SetProperty("SpecifyMaximumFileSizeApplicationLog", value)
}

// GetSpecifyMaximumFileSizeApplicationLog gets the value of SpecifyMaximumFileSizeApplicationLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertySpecifyMaximumFileSizeApplicationLog() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaximumFileSizeApplicationLog")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyMaximumFileSizeSecurityLog sets the value of SpecifyMaximumFileSizeSecurityLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertySpecifyMaximumFileSizeSecurityLog(value string) (err error) {
	return instance.SetProperty("SpecifyMaximumFileSizeSecurityLog", value)
}

// GetSpecifyMaximumFileSizeSecurityLog gets the value of SpecifyMaximumFileSizeSecurityLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertySpecifyMaximumFileSizeSecurityLog() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaximumFileSizeSecurityLog")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyMaximumFileSizeSystemLog sets the value of SpecifyMaximumFileSizeSystemLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertySpecifyMaximumFileSizeSystemLog(value string) (err error) {
	return instance.SetProperty("SpecifyMaximumFileSizeSystemLog", value)
}

// GetSpecifyMaximumFileSizeSystemLog gets the value of SpecifyMaximumFileSizeSystemLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertySpecifyMaximumFileSizeSystemLog() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaximumFileSizeSystemLog")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
