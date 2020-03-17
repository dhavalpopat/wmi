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

// MDM_Policy_Config01_FileExplorer02 struct
type MDM_Policy_Config01_FileExplorer02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	TurnOffDataExecutionPreventionForExplorer string

	//
	TurnOffHeapTerminationOnCorruption string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_FileExplorer02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) GetPropertyParentID() (value string, err error) {
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

// SetTurnOffDataExecutionPreventionForExplorer sets the value of TurnOffDataExecutionPreventionForExplorer for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) SetPropertyTurnOffDataExecutionPreventionForExplorer(value string) (err error) {
	return instance.SetProperty("TurnOffDataExecutionPreventionForExplorer", value)
}

// GetTurnOffDataExecutionPreventionForExplorer gets the value of TurnOffDataExecutionPreventionForExplorer for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) GetPropertyTurnOffDataExecutionPreventionForExplorer() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOffDataExecutionPreventionForExplorer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurnOffHeapTerminationOnCorruption sets the value of TurnOffHeapTerminationOnCorruption for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) SetPropertyTurnOffHeapTerminationOnCorruption(value string) (err error) {
	return instance.SetProperty("TurnOffHeapTerminationOnCorruption", value)
}

// GetTurnOffHeapTerminationOnCorruption gets the value of TurnOffHeapTerminationOnCorruption for the instance
func (instance *MDM_Policy_Config01_FileExplorer02) GetPropertyTurnOffHeapTerminationOnCorruption() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOffHeapTerminationOnCorruption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
