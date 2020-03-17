// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSoP_PolicySettingStatus struct
type RSoP_PolicySettingStatus struct {
	cim.WmiInstance

	//
	errorCode uint32

	//
	eventID uint32

	//
	eventLogName string

	//
	eventSource string

	//
	eventTime string

	//
	id string

	//
	status int32
}

// SeterrorCode sets the value of errorCode for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyerrorCode(value uint32) (err error) {
	return instance.SetProperty("errorCode", value)
}

// GeterrorCode gets the value of errorCode for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyerrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("errorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SeteventID sets the value of eventID for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventID(value uint32) (err error) {
	return instance.SetProperty("eventID", value)
}

// GeteventID gets the value of eventID for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventID() (value uint32, err error) {
	retValue, err := instance.GetProperty("eventID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SeteventLogName sets the value of eventLogName for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventLogName(value string) (err error) {
	return instance.SetProperty("eventLogName", value)
}

// GeteventLogName gets the value of eventLogName for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventLogName() (value string, err error) {
	retValue, err := instance.GetProperty("eventLogName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SeteventSource sets the value of eventSource for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventSource(value string) (err error) {
	return instance.SetProperty("eventSource", value)
}

// GeteventSource gets the value of eventSource for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventSource() (value string, err error) {
	retValue, err := instance.GetProperty("eventSource")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SeteventTime sets the value of eventTime for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventTime(value string) (err error) {
	return instance.SetProperty("eventTime", value)
}

// GeteventTime gets the value of eventTime for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventTime() (value string, err error) {
	retValue, err := instance.GetProperty("eventTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setid sets the value of id for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", value)
}

// Getid gets the value of id for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setstatus sets the value of status for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertystatus(value int32) (err error) {
	return instance.SetProperty("status", value)
}

// Getstatus gets the value of status for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertystatus() (value int32, err error) {
	retValue, err := instance.GetProperty("status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
