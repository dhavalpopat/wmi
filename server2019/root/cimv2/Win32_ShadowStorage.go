// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ShadowStorage struct
type Win32_ShadowStorage struct {
	cim.WmiInstance

	//
	AllocatedSpace uint64

	//
	DiffVolume Win32_Volume

	//
	MaxSpace uint64

	//
	UsedSpace uint64

	//
	Volume Win32_Volume
}

// SetAllocatedSpace sets the value of AllocatedSpace for the instance
func (instance *Win32_ShadowStorage) SetPropertyAllocatedSpace(value uint64) (err error) {
	return instance.SetProperty("AllocatedSpace", value)
}

// GetAllocatedSpace gets the value of AllocatedSpace for the instance
func (instance *Win32_ShadowStorage) GetPropertyAllocatedSpace() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocatedSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiffVolume sets the value of DiffVolume for the instance
func (instance *Win32_ShadowStorage) SetPropertyDiffVolume(value Win32_Volume) (err error) {
	return instance.SetProperty("DiffVolume", value)
}

// GetDiffVolume gets the value of DiffVolume for the instance
func (instance *Win32_ShadowStorage) GetPropertyDiffVolume() (value Win32_Volume, err error) {
	retValue, err := instance.GetProperty("DiffVolume")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSpace sets the value of MaxSpace for the instance
func (instance *Win32_ShadowStorage) SetPropertyMaxSpace(value uint64) (err error) {
	return instance.SetProperty("MaxSpace", value)
}

// GetMaxSpace gets the value of MaxSpace for the instance
func (instance *Win32_ShadowStorage) GetPropertyMaxSpace() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsedSpace sets the value of UsedSpace for the instance
func (instance *Win32_ShadowStorage) SetPropertyUsedSpace(value uint64) (err error) {
	return instance.SetProperty("UsedSpace", value)
}

// GetUsedSpace gets the value of UsedSpace for the instance
func (instance *Win32_ShadowStorage) GetPropertyUsedSpace() (value uint64, err error) {
	retValue, err := instance.GetProperty("UsedSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolume sets the value of Volume for the instance
func (instance *Win32_ShadowStorage) SetPropertyVolume(value Win32_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *Win32_ShadowStorage) GetPropertyVolume() (value Win32_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="DiffVolume" type="string "></param>
// <param name="MaxSpace" type="uint64 "></param>
// <param name="Volume" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_ShadowStorage) Create( /* IN */ Volume string,
	/* IN */ DiffVolume string,
	/* IN */ MaxSpace uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Create", Volume, DiffVolume, MaxSpace)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
