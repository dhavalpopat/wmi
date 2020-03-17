// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSetVM struct
type MSFT_ClusterSetVM struct {
	cim.WmiInstance

	//
	AvailabilitySet uint64

	//
	CheckHeartBeat bool

	//
	DefaultMoveType uint32

	//
	FaultDomain uint64

	//
	Id uint64

	//
	MemberId uint64

	//
	MemberName string

	//
	NodeId uint64

	//
	NodeName string

	//
	OfflineAction uint32

	//
	PlacementCondition string

	//
	StartMemory uint32

	//
	State uint32

	//
	UpdateDomain uint64

	//
	VMId string

	//
	VMName string
}

// SetAvailabilitySet sets the value of AvailabilitySet for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyAvailabilitySet(value uint64) (err error) {
	return instance.SetProperty("AvailabilitySet", value)
}

// GetAvailabilitySet gets the value of AvailabilitySet for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyAvailabilitySet() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailabilitySet")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCheckHeartBeat sets the value of CheckHeartBeat for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyCheckHeartBeat(value bool) (err error) {
	return instance.SetProperty("CheckHeartBeat", value)
}

// GetCheckHeartBeat gets the value of CheckHeartBeat for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyCheckHeartBeat() (value bool, err error) {
	retValue, err := instance.GetProperty("CheckHeartBeat")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultMoveType sets the value of DefaultMoveType for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyDefaultMoveType(value uint32) (err error) {
	return instance.SetProperty("DefaultMoveType", value)
}

// GetDefaultMoveType gets the value of DefaultMoveType for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyDefaultMoveType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultMoveType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomain sets the value of FaultDomain for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyFaultDomain(value uint64) (err error) {
	return instance.SetProperty("FaultDomain", value)
}

// GetFaultDomain gets the value of FaultDomain for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyFaultDomain() (value uint64, err error) {
	retValue, err := instance.GetProperty("FaultDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyId() (value uint64, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemberId sets the value of MemberId for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyMemberId(value uint64) (err error) {
	return instance.SetProperty("MemberId", value)
}

// GetMemberId gets the value of MemberId for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyMemberId() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemberId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemberName sets the value of MemberName for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyMemberName(value string) (err error) {
	return instance.SetProperty("MemberName", value)
}

// GetMemberName gets the value of MemberName for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyMemberName() (value string, err error) {
	retValue, err := instance.GetProperty("MemberName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNodeId sets the value of NodeId for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyNodeId(value uint64) (err error) {
	return instance.SetProperty("NodeId", value)
}

// GetNodeId gets the value of NodeId for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyNodeId() (value uint64, err error) {
	retValue, err := instance.GetProperty("NodeId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNodeName sets the value of NodeName for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyNodeName(value string) (err error) {
	return instance.SetProperty("NodeName", value)
}

// GetNodeName gets the value of NodeName for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("NodeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOfflineAction sets the value of OfflineAction for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyOfflineAction(value uint32) (err error) {
	return instance.SetProperty("OfflineAction", value)
}

// GetOfflineAction gets the value of OfflineAction for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyOfflineAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("OfflineAction")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlacementCondition sets the value of PlacementCondition for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyPlacementCondition(value string) (err error) {
	return instance.SetProperty("PlacementCondition", value)
}

// GetPlacementCondition gets the value of PlacementCondition for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyPlacementCondition() (value string, err error) {
	retValue, err := instance.GetProperty("PlacementCondition")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartMemory sets the value of StartMemory for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyStartMemory(value uint32) (err error) {
	return instance.SetProperty("StartMemory", value)
}

// GetStartMemory gets the value of StartMemory for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyStartMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateDomain sets the value of UpdateDomain for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyUpdateDomain(value uint64) (err error) {
	return instance.SetProperty("UpdateDomain", value)
}

// GetUpdateDomain gets the value of UpdateDomain for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyUpdateDomain() (value uint64, err error) {
	retValue, err := instance.GetProperty("UpdateDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMId sets the value of VMId for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyVMId(value string) (err error) {
	return instance.SetProperty("VMId", value)
}

// GetVMId gets the value of VMId for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyVMId() (value string, err error) {
	retValue, err := instance.GetProperty("VMId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMName sets the value of VMName for the instance
func (instance *MSFT_ClusterSetVM) SetPropertyVMName(value string) (err error) {
	return instance.SetProperty("VMName", value)
}

// GetVMName gets the value of VMName for the instance
func (instance *MSFT_ClusterSetVM) GetPropertyVMName() (value string, err error) {
	retValue, err := instance.GetProperty("VMName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetVM) Online() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Online")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetVM) Offline() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Offline")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MoveType" type="uint32 "></param>
// <param name="Node" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetVM) Move( /* IN */ Node string,
	/* IN */ MoveType uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Move", Node, MoveType)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetVM) DestroyVm( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DestroyVm", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetVM) UnclusterVm( /* IN */ Force bool,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UnclusterVm", Force, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AvailabilitySetName" type="string "></param>
// <param name="PlacementCondition" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetVM) SetVmProperties( /* IN */ AvailabilitySetName string,
	/* IN */ PlacementCondition string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetVmProperties", AvailabilitySetName, PlacementCondition)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
