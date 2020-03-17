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

// MDM_Policy_User_Result01_Display02 struct
type MDM_Policy_User_Result01_Display02 struct {
	cim.WmiInstance

	//
	EnablePerProcessDpi int32

	//
	InstanceID string

	//
	ParentID string
}

// SetEnablePerProcessDpi sets the value of EnablePerProcessDpi for the instance
func (instance *MDM_Policy_User_Result01_Display02) SetPropertyEnablePerProcessDpi(value int32) (err error) {
	return instance.SetProperty("EnablePerProcessDpi", value)
}

// GetEnablePerProcessDpi gets the value of EnablePerProcessDpi for the instance
func (instance *MDM_Policy_User_Result01_Display02) GetPropertyEnablePerProcessDpi() (value int32, err error) {
	retValue, err := instance.GetProperty("EnablePerProcessDpi")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Display02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Display02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_User_Result01_Display02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_Display02) GetPropertyParentID() (value string, err error) {
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
