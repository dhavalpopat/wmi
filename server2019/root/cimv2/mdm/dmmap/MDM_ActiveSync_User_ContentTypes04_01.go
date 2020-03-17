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

// MDM_ActiveSync_User_ContentTypes04_01 struct
type MDM_ActiveSync_User_ContentTypes04_01 struct {
	cim.WmiInstance

	//
	Enabled string

	//
	InstanceID string

	//
	Name string

	//
	ParentID string
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) SetPropertyEnabled(value string) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) GetPropertyEnabled() (value string, err error) {
	retValue, err := instance.GetProperty("Enabled")
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
func (instance *MDM_ActiveSync_User_ContentTypes04_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) GetPropertyInstanceID() (value string, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) GetPropertyName() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ActiveSync_User_ContentTypes04_01) GetPropertyParentID() (value string, err error) {
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
