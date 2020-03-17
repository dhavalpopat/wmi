// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel struct
type Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel struct {
	Win32_PerfRawData

	//
	Numberofdataavailableeventwasreset uint32

	//
	Numberofdataavailableeventwasresetpersecond uint64

	//
	Numberofdataavailablesignalsreceived uint32

	//
	Numberofdataavailablesignalsreceivedpersecond uint64

	//
	Numberofdataavailablesignalssent uint32

	//
	Numberofdataavailablesignalssentpersecond uint64

	//
	Numberofspaceavailableeventwasreset uint32

	//
	Numberofspaceavailableeventwasresetpersecond uint64

	//
	Numberofspaceavailablesignalsreceived uint32

	//
	Numberofspaceavailablesignalsreceivedpersecond uint64

	//
	Numberofspaceavailablesignalssent uint32

	//
	Numberofspaceavailablesignalssentpersecond uint64
}

// SetNumberofdataavailableeventwasreset sets the value of Numberofdataavailableeventwasreset for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofdataavailableeventwasreset(value uint32) (err error) {
	return instance.SetProperty("Numberofdataavailableeventwasreset", value)
}

// GetNumberofdataavailableeventwasreset gets the value of Numberofdataavailableeventwasreset for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofdataavailableeventwasreset() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofdataavailableeventwasreset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofdataavailableeventwasresetpersecond sets the value of Numberofdataavailableeventwasresetpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofdataavailableeventwasresetpersecond(value uint64) (err error) {
	return instance.SetProperty("Numberofdataavailableeventwasresetpersecond", value)
}

// GetNumberofdataavailableeventwasresetpersecond gets the value of Numberofdataavailableeventwasresetpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofdataavailableeventwasresetpersecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofdataavailableeventwasresetpersecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofdataavailablesignalsreceived sets the value of Numberofdataavailablesignalsreceived for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofdataavailablesignalsreceived(value uint32) (err error) {
	return instance.SetProperty("Numberofdataavailablesignalsreceived", value)
}

// GetNumberofdataavailablesignalsreceived gets the value of Numberofdataavailablesignalsreceived for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofdataavailablesignalsreceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofdataavailablesignalsreceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofdataavailablesignalsreceivedpersecond sets the value of Numberofdataavailablesignalsreceivedpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofdataavailablesignalsreceivedpersecond(value uint64) (err error) {
	return instance.SetProperty("Numberofdataavailablesignalsreceivedpersecond", value)
}

// GetNumberofdataavailablesignalsreceivedpersecond gets the value of Numberofdataavailablesignalsreceivedpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofdataavailablesignalsreceivedpersecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofdataavailablesignalsreceivedpersecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofdataavailablesignalssent sets the value of Numberofdataavailablesignalssent for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofdataavailablesignalssent(value uint32) (err error) {
	return instance.SetProperty("Numberofdataavailablesignalssent", value)
}

// GetNumberofdataavailablesignalssent gets the value of Numberofdataavailablesignalssent for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofdataavailablesignalssent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofdataavailablesignalssent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofdataavailablesignalssentpersecond sets the value of Numberofdataavailablesignalssentpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofdataavailablesignalssentpersecond(value uint64) (err error) {
	return instance.SetProperty("Numberofdataavailablesignalssentpersecond", value)
}

// GetNumberofdataavailablesignalssentpersecond gets the value of Numberofdataavailablesignalssentpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofdataavailablesignalssentpersecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofdataavailablesignalssentpersecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofspaceavailableeventwasreset sets the value of Numberofspaceavailableeventwasreset for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofspaceavailableeventwasreset(value uint32) (err error) {
	return instance.SetProperty("Numberofspaceavailableeventwasreset", value)
}

// GetNumberofspaceavailableeventwasreset gets the value of Numberofspaceavailableeventwasreset for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofspaceavailableeventwasreset() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofspaceavailableeventwasreset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofspaceavailableeventwasresetpersecond sets the value of Numberofspaceavailableeventwasresetpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofspaceavailableeventwasresetpersecond(value uint64) (err error) {
	return instance.SetProperty("Numberofspaceavailableeventwasresetpersecond", value)
}

// GetNumberofspaceavailableeventwasresetpersecond gets the value of Numberofspaceavailableeventwasresetpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofspaceavailableeventwasresetpersecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofspaceavailableeventwasresetpersecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofspaceavailablesignalsreceived sets the value of Numberofspaceavailablesignalsreceived for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofspaceavailablesignalsreceived(value uint32) (err error) {
	return instance.SetProperty("Numberofspaceavailablesignalsreceived", value)
}

// GetNumberofspaceavailablesignalsreceived gets the value of Numberofspaceavailablesignalsreceived for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofspaceavailablesignalsreceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofspaceavailablesignalsreceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofspaceavailablesignalsreceivedpersecond sets the value of Numberofspaceavailablesignalsreceivedpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofspaceavailablesignalsreceivedpersecond(value uint64) (err error) {
	return instance.SetProperty("Numberofspaceavailablesignalsreceivedpersecond", value)
}

// GetNumberofspaceavailablesignalsreceivedpersecond gets the value of Numberofspaceavailablesignalsreceivedpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofspaceavailablesignalsreceivedpersecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofspaceavailablesignalsreceivedpersecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofspaceavailablesignalssent sets the value of Numberofspaceavailablesignalssent for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofspaceavailablesignalssent(value uint32) (err error) {
	return instance.SetProperty("Numberofspaceavailablesignalssent", value)
}

// GetNumberofspaceavailablesignalssent gets the value of Numberofspaceavailablesignalssent for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofspaceavailablesignalssent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofspaceavailablesignalssent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofspaceavailablesignalssentpersecond sets the value of Numberofspaceavailablesignalssentpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) SetPropertyNumberofspaceavailablesignalssentpersecond(value uint64) (err error) {
	return instance.SetProperty("Numberofspaceavailablesignalssentpersecond", value)
}

// GetNumberofspaceavailablesignalssentpersecond gets the value of Numberofspaceavailablesignalssentpersecond for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMTransportChannel) GetPropertyNumberofspaceavailablesignalssentpersecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofspaceavailablesignalssentpersecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
