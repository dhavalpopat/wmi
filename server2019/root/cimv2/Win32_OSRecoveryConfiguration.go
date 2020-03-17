// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_OSRecoveryConfiguration struct
type Win32_OSRecoveryConfiguration struct {
	CIM_Setting

	//
	AutoReboot bool

	//
	DebugFilePath string

	//
	DebugInfoType uint32

	//
	ExpandedDebugFilePath string

	//
	ExpandedMiniDumpDirectory string

	//
	KernelDumpOnly bool

	//
	MiniDumpDirectory string

	//
	Name string

	//
	OverwriteExistingDebugFile bool

	//
	SendAdminAlert bool

	//
	WriteDebugInfo bool

	//
	WriteToSystemLog bool
}

// SetAutoReboot sets the value of AutoReboot for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyAutoReboot(value bool) (err error) {
	return instance.SetProperty("AutoReboot", value)
}

// GetAutoReboot gets the value of AutoReboot for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyAutoReboot() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoReboot")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDebugFilePath sets the value of DebugFilePath for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyDebugFilePath(value string) (err error) {
	return instance.SetProperty("DebugFilePath", value)
}

// GetDebugFilePath gets the value of DebugFilePath for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyDebugFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("DebugFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDebugInfoType sets the value of DebugInfoType for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyDebugInfoType(value uint32) (err error) {
	return instance.SetProperty("DebugInfoType", value)
}

// GetDebugInfoType gets the value of DebugInfoType for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyDebugInfoType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DebugInfoType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpandedDebugFilePath sets the value of ExpandedDebugFilePath for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyExpandedDebugFilePath(value string) (err error) {
	return instance.SetProperty("ExpandedDebugFilePath", value)
}

// GetExpandedDebugFilePath gets the value of ExpandedDebugFilePath for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyExpandedDebugFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("ExpandedDebugFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpandedMiniDumpDirectory sets the value of ExpandedMiniDumpDirectory for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyExpandedMiniDumpDirectory(value string) (err error) {
	return instance.SetProperty("ExpandedMiniDumpDirectory", value)
}

// GetExpandedMiniDumpDirectory gets the value of ExpandedMiniDumpDirectory for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyExpandedMiniDumpDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("ExpandedMiniDumpDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKernelDumpOnly sets the value of KernelDumpOnly for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyKernelDumpOnly(value bool) (err error) {
	return instance.SetProperty("KernelDumpOnly", value)
}

// GetKernelDumpOnly gets the value of KernelDumpOnly for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyKernelDumpOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("KernelDumpOnly")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMiniDumpDirectory sets the value of MiniDumpDirectory for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyMiniDumpDirectory(value string) (err error) {
	return instance.SetProperty("MiniDumpDirectory", value)
}

// GetMiniDumpDirectory gets the value of MiniDumpDirectory for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyMiniDumpDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("MiniDumpDirectory")
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
func (instance *Win32_OSRecoveryConfiguration) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyName() (value string, err error) {
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

// SetOverwriteExistingDebugFile sets the value of OverwriteExistingDebugFile for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyOverwriteExistingDebugFile(value bool) (err error) {
	return instance.SetProperty("OverwriteExistingDebugFile", value)
}

// GetOverwriteExistingDebugFile gets the value of OverwriteExistingDebugFile for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyOverwriteExistingDebugFile() (value bool, err error) {
	retValue, err := instance.GetProperty("OverwriteExistingDebugFile")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendAdminAlert sets the value of SendAdminAlert for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertySendAdminAlert(value bool) (err error) {
	return instance.SetProperty("SendAdminAlert", value)
}

// GetSendAdminAlert gets the value of SendAdminAlert for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertySendAdminAlert() (value bool, err error) {
	retValue, err := instance.GetProperty("SendAdminAlert")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteDebugInfo sets the value of WriteDebugInfo for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyWriteDebugInfo(value bool) (err error) {
	return instance.SetProperty("WriteDebugInfo", value)
}

// GetWriteDebugInfo gets the value of WriteDebugInfo for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyWriteDebugInfo() (value bool, err error) {
	retValue, err := instance.GetProperty("WriteDebugInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteToSystemLog sets the value of WriteToSystemLog for the instance
func (instance *Win32_OSRecoveryConfiguration) SetPropertyWriteToSystemLog(value bool) (err error) {
	return instance.SetProperty("WriteToSystemLog", value)
}

// GetWriteToSystemLog gets the value of WriteToSystemLog for the instance
func (instance *Win32_OSRecoveryConfiguration) GetPropertyWriteToSystemLog() (value bool, err error) {
	retValue, err := instance.GetProperty("WriteToSystemLog")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
