// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VHDSetInformation struct
type Msvm_VHDSetInformation struct {
	cim.WmiInstance

	//
	AllPaths []string

	//
	Path string

	//
	SnapshotIdList []string
}

// SetAllPaths sets the value of AllPaths for the instance
func (instance *Msvm_VHDSetInformation) SetPropertyAllPaths(value []string) (err error) {
	return instance.SetProperty("AllPaths", value)
}

// GetAllPaths gets the value of AllPaths for the instance
func (instance *Msvm_VHDSetInformation) GetPropertyAllPaths() (value []string, err error) {
	retValue, err := instance.GetProperty("AllPaths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *Msvm_VHDSetInformation) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *Msvm_VHDSetInformation) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSnapshotIdList sets the value of SnapshotIdList for the instance
func (instance *Msvm_VHDSetInformation) SetPropertySnapshotIdList(value []string) (err error) {
	return instance.SetProperty("SnapshotIdList", value)
}

// GetSnapshotIdList gets the value of SnapshotIdList for the instance
func (instance *Msvm_VHDSetInformation) GetPropertySnapshotIdList() (value []string, err error) {
	retValue, err := instance.GetProperty("SnapshotIdList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
