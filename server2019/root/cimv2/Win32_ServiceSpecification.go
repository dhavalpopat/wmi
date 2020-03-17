// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ServiceSpecification struct
type Win32_ServiceSpecification struct {
	CIM_Check

	//
	Dependencies string

	//
	DisplayName string

	//
	ErrorControl int32

	//
	ID string

	//
	LoadOrderGroup string

	//
	Password string

	//
	ServiceType int32

	//
	StartName string

	//
	StartType int32
}

// SetDependencies sets the value of Dependencies for the instance
func (instance *Win32_ServiceSpecification) SetPropertyDependencies(value string) (err error) {
	return instance.SetProperty("Dependencies", value)
}

// GetDependencies gets the value of Dependencies for the instance
func (instance *Win32_ServiceSpecification) GetPropertyDependencies() (value string, err error) {
	retValue, err := instance.GetProperty("Dependencies")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *Win32_ServiceSpecification) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *Win32_ServiceSpecification) GetPropertyDisplayName() (value string, err error) {
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

// SetErrorControl sets the value of ErrorControl for the instance
func (instance *Win32_ServiceSpecification) SetPropertyErrorControl(value int32) (err error) {
	return instance.SetProperty("ErrorControl", value)
}

// GetErrorControl gets the value of ErrorControl for the instance
func (instance *Win32_ServiceSpecification) GetPropertyErrorControl() (value int32, err error) {
	retValue, err := instance.GetProperty("ErrorControl")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *Win32_ServiceSpecification) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *Win32_ServiceSpecification) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadOrderGroup sets the value of LoadOrderGroup for the instance
func (instance *Win32_ServiceSpecification) SetPropertyLoadOrderGroup(value string) (err error) {
	return instance.SetProperty("LoadOrderGroup", value)
}

// GetLoadOrderGroup gets the value of LoadOrderGroup for the instance
func (instance *Win32_ServiceSpecification) GetPropertyLoadOrderGroup() (value string, err error) {
	retValue, err := instance.GetProperty("LoadOrderGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassword sets the value of Password for the instance
func (instance *Win32_ServiceSpecification) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", value)
}

// GetPassword gets the value of Password for the instance
func (instance *Win32_ServiceSpecification) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServiceType sets the value of ServiceType for the instance
func (instance *Win32_ServiceSpecification) SetPropertyServiceType(value int32) (err error) {
	return instance.SetProperty("ServiceType", value)
}

// GetServiceType gets the value of ServiceType for the instance
func (instance *Win32_ServiceSpecification) GetPropertyServiceType() (value int32, err error) {
	retValue, err := instance.GetProperty("ServiceType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartName sets the value of StartName for the instance
func (instance *Win32_ServiceSpecification) SetPropertyStartName(value string) (err error) {
	return instance.SetProperty("StartName", value)
}

// GetStartName gets the value of StartName for the instance
func (instance *Win32_ServiceSpecification) GetPropertyStartName() (value string, err error) {
	retValue, err := instance.GetProperty("StartName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartType sets the value of StartType for the instance
func (instance *Win32_ServiceSpecification) SetPropertyStartType(value int32) (err error) {
	return instance.SetProperty("StartType", value)
}

// GetStartType gets the value of StartType for the instance
func (instance *Win32_ServiceSpecification) GetPropertyStartType() (value int32, err error) {
	retValue, err := instance.GetProperty("StartType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
