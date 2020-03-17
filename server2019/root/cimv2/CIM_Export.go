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

// CIM_Export struct
type CIM_Export struct {
	cim.WmiInstance

	//
	Directory CIM_Directory

	//
	ExportedDirectoryName string

	//
	LocalFS CIM_LocalFileSystem
}

// SetDirectory sets the value of Directory for the instance
func (instance *CIM_Export) SetPropertyDirectory(value CIM_Directory) (err error) {
	return instance.SetProperty("Directory", value)
}

// GetDirectory gets the value of Directory for the instance
func (instance *CIM_Export) GetPropertyDirectory() (value CIM_Directory, err error) {
	retValue, err := instance.GetProperty("Directory")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Directory)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedDirectoryName sets the value of ExportedDirectoryName for the instance
func (instance *CIM_Export) SetPropertyExportedDirectoryName(value string) (err error) {
	return instance.SetProperty("ExportedDirectoryName", value)
}

// GetExportedDirectoryName gets the value of ExportedDirectoryName for the instance
func (instance *CIM_Export) GetPropertyExportedDirectoryName() (value string, err error) {
	retValue, err := instance.GetProperty("ExportedDirectoryName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalFS sets the value of LocalFS for the instance
func (instance *CIM_Export) SetPropertyLocalFS(value CIM_LocalFileSystem) (err error) {
	return instance.SetProperty("LocalFS", value)
}

// GetLocalFS gets the value of LocalFS for the instance
func (instance *CIM_Export) GetPropertyLocalFS() (value CIM_LocalFileSystem, err error) {
	retValue, err := instance.GetProperty("LocalFS")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_LocalFileSystem)
	if !ok {
		// TODO: Set an error
	}
	return
}
