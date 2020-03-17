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

// MDM_Policy_Config01_Experience02 struct
type MDM_Policy_Config01_Experience02 struct {
	cim.WmiInstance

	//
	AllowClipboardHistory int32

	//
	AllowCortana int32

	//
	AllowDeviceDiscovery int32

	//
	AllowFindMyDevice int32

	//
	AllowManualMDMUnenrollment int32

	//
	AllowSaveAsOfOfficeFiles int32

	//
	AllowScreenCapture int32

	//
	AllowSharingOfOfficeFiles int32

	//
	AllowSIMErrorDialogPromptWhenNoSIM int32

	//
	AllowSyncMySettings int32

	//
	AllowWindowsConsumerFeatures int32

	//
	AllowWindowsTips int32

	//
	DoNotShowFeedbackNotifications int32

	//
	DoNotSyncBrowserSettings int32

	//
	InstanceID string

	//
	ParentID string

	//
	PreventUsersFromTurningOnBrowserSyncing int32

	//
	ShowLockOnUserTile int32
}

// SetAllowClipboardHistory sets the value of AllowClipboardHistory for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowClipboardHistory(value int32) (err error) {
	return instance.SetProperty("AllowClipboardHistory", value)
}

// GetAllowClipboardHistory gets the value of AllowClipboardHistory for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowClipboardHistory() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowClipboardHistory")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowCortana sets the value of AllowCortana for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowCortana(value int32) (err error) {
	return instance.SetProperty("AllowCortana", value)
}

// GetAllowCortana gets the value of AllowCortana for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowCortana() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCortana")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowDeviceDiscovery sets the value of AllowDeviceDiscovery for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowDeviceDiscovery(value int32) (err error) {
	return instance.SetProperty("AllowDeviceDiscovery", value)
}

// GetAllowDeviceDiscovery gets the value of AllowDeviceDiscovery for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowDeviceDiscovery() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDeviceDiscovery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowFindMyDevice sets the value of AllowFindMyDevice for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowFindMyDevice(value int32) (err error) {
	return instance.SetProperty("AllowFindMyDevice", value)
}

// GetAllowFindMyDevice gets the value of AllowFindMyDevice for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowFindMyDevice() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFindMyDevice")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowManualMDMUnenrollment sets the value of AllowManualMDMUnenrollment for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowManualMDMUnenrollment(value int32) (err error) {
	return instance.SetProperty("AllowManualMDMUnenrollment", value)
}

// GetAllowManualMDMUnenrollment gets the value of AllowManualMDMUnenrollment for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowManualMDMUnenrollment() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowManualMDMUnenrollment")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSaveAsOfOfficeFiles sets the value of AllowSaveAsOfOfficeFiles for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowSaveAsOfOfficeFiles(value int32) (err error) {
	return instance.SetProperty("AllowSaveAsOfOfficeFiles", value)
}

// GetAllowSaveAsOfOfficeFiles gets the value of AllowSaveAsOfOfficeFiles for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowSaveAsOfOfficeFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSaveAsOfOfficeFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowScreenCapture sets the value of AllowScreenCapture for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowScreenCapture(value int32) (err error) {
	return instance.SetProperty("AllowScreenCapture", value)
}

// GetAllowScreenCapture gets the value of AllowScreenCapture for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowScreenCapture() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowScreenCapture")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSharingOfOfficeFiles sets the value of AllowSharingOfOfficeFiles for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowSharingOfOfficeFiles(value int32) (err error) {
	return instance.SetProperty("AllowSharingOfOfficeFiles", value)
}

// GetAllowSharingOfOfficeFiles gets the value of AllowSharingOfOfficeFiles for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowSharingOfOfficeFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSharingOfOfficeFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSIMErrorDialogPromptWhenNoSIM sets the value of AllowSIMErrorDialogPromptWhenNoSIM for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowSIMErrorDialogPromptWhenNoSIM(value int32) (err error) {
	return instance.SetProperty("AllowSIMErrorDialogPromptWhenNoSIM", value)
}

// GetAllowSIMErrorDialogPromptWhenNoSIM gets the value of AllowSIMErrorDialogPromptWhenNoSIM for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowSIMErrorDialogPromptWhenNoSIM() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSIMErrorDialogPromptWhenNoSIM")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSyncMySettings sets the value of AllowSyncMySettings for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowSyncMySettings(value int32) (err error) {
	return instance.SetProperty("AllowSyncMySettings", value)
}

// GetAllowSyncMySettings gets the value of AllowSyncMySettings for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowSyncMySettings() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSyncMySettings")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsConsumerFeatures sets the value of AllowWindowsConsumerFeatures for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowWindowsConsumerFeatures(value int32) (err error) {
	return instance.SetProperty("AllowWindowsConsumerFeatures", value)
}

// GetAllowWindowsConsumerFeatures gets the value of AllowWindowsConsumerFeatures for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowWindowsConsumerFeatures() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsConsumerFeatures")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsTips sets the value of AllowWindowsTips for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyAllowWindowsTips(value int32) (err error) {
	return instance.SetProperty("AllowWindowsTips", value)
}

// GetAllowWindowsTips gets the value of AllowWindowsTips for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyAllowWindowsTips() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsTips")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotShowFeedbackNotifications sets the value of DoNotShowFeedbackNotifications for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyDoNotShowFeedbackNotifications(value int32) (err error) {
	return instance.SetProperty("DoNotShowFeedbackNotifications", value)
}

// GetDoNotShowFeedbackNotifications gets the value of DoNotShowFeedbackNotifications for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyDoNotShowFeedbackNotifications() (value int32, err error) {
	retValue, err := instance.GetProperty("DoNotShowFeedbackNotifications")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotSyncBrowserSettings sets the value of DoNotSyncBrowserSettings for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyDoNotSyncBrowserSettings(value int32) (err error) {
	return instance.SetProperty("DoNotSyncBrowserSettings", value)
}

// GetDoNotSyncBrowserSettings gets the value of DoNotSyncBrowserSettings for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyDoNotSyncBrowserSettings() (value int32, err error) {
	retValue, err := instance.GetProperty("DoNotSyncBrowserSettings")
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
func (instance *MDM_Policy_Config01_Experience02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_Experience02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyParentID() (value string, err error) {
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

// SetPreventUsersFromTurningOnBrowserSyncing sets the value of PreventUsersFromTurningOnBrowserSyncing for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyPreventUsersFromTurningOnBrowserSyncing(value int32) (err error) {
	return instance.SetProperty("PreventUsersFromTurningOnBrowserSyncing", value)
}

// GetPreventUsersFromTurningOnBrowserSyncing gets the value of PreventUsersFromTurningOnBrowserSyncing for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyPreventUsersFromTurningOnBrowserSyncing() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventUsersFromTurningOnBrowserSyncing")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShowLockOnUserTile sets the value of ShowLockOnUserTile for the instance
func (instance *MDM_Policy_Config01_Experience02) SetPropertyShowLockOnUserTile(value int32) (err error) {
	return instance.SetProperty("ShowLockOnUserTile", value)
}

// GetShowLockOnUserTile gets the value of ShowLockOnUserTile for the instance
func (instance *MDM_Policy_Config01_Experience02) GetPropertyShowLockOnUserTile() (value int32, err error) {
	retValue, err := instance.GetProperty("ShowLockOnUserTile")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
