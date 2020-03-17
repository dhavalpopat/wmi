// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DeviceErrorCounts struct
type CIM_DeviceErrorCounts struct {
	CIM_StatisticalInformation

	//
	CriticalErrorCount uint64

	//
	DeviceCreationClassName string

	//
	DeviceID string

	//
	IndeterminateErrorCount uint64

	//
	MajorErrorCount uint64

	//
	MinorErrorCount uint64

	//
	SystemCreationClassName string

	//
	SystemName string

	//
	WarningCount uint64
}

// SetCriticalErrorCount sets the value of CriticalErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyCriticalErrorCount(value uint64) (err error) {
	return instance.SetProperty("CriticalErrorCount", value)
}

// GetCriticalErrorCount gets the value of CriticalErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyCriticalErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CriticalErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceCreationClassName sets the value of DeviceCreationClassName for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyDeviceCreationClassName(value string) (err error) {
	return instance.SetProperty("DeviceCreationClassName", value)
}

// GetDeviceCreationClassName gets the value of DeviceCreationClassName for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyDeviceCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceID sets the value of DeviceID for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyDeviceID(value string) (err error) {
	return instance.SetProperty("DeviceID", value)
}

// GetDeviceID gets the value of DeviceID for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyDeviceID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIndeterminateErrorCount sets the value of IndeterminateErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyIndeterminateErrorCount(value uint64) (err error) {
	return instance.SetProperty("IndeterminateErrorCount", value)
}

// GetIndeterminateErrorCount gets the value of IndeterminateErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyIndeterminateErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("IndeterminateErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMajorErrorCount sets the value of MajorErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyMajorErrorCount(value uint64) (err error) {
	return instance.SetProperty("MajorErrorCount", value)
}

// GetMajorErrorCount gets the value of MajorErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyMajorErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MajorErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinorErrorCount sets the value of MinorErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyMinorErrorCount(value uint64) (err error) {
	return instance.SetProperty("MinorErrorCount", value)
}

// GetMinorErrorCount gets the value of MinorErrorCount for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyMinorErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinorErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWarningCount sets the value of WarningCount for the instance
func (instance *CIM_DeviceErrorCounts) SetPropertyWarningCount(value uint64) (err error) {
	return instance.SetProperty("WarningCount", value)
}

// GetWarningCount gets the value of WarningCount for the instance
func (instance *CIM_DeviceErrorCounts) GetPropertyWarningCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WarningCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="SelectedCounter" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_DeviceErrorCounts) ResetCounter( /* IN */ SelectedCounter uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResetCounter", SelectedCounter)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
