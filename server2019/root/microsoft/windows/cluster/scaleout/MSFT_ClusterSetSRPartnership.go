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

// MSFT_ClusterSetSRPartnership struct
type MSFT_ClusterSetSRPartnership struct {
	cim.WmiInstance

	//
	DestinationClusterId uint64

	//
	DestinationClusterName string

	//
	Id uint64

	//
	Name string

	//
	PartnershipId string

	//
	SourceClusterId uint64

	//
	SourceClusterName string

	//
	State uint32
}

// SetDestinationClusterId sets the value of DestinationClusterId for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertyDestinationClusterId(value uint64) (err error) {
	return instance.SetProperty("DestinationClusterId", value)
}

// GetDestinationClusterId gets the value of DestinationClusterId for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertyDestinationClusterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("DestinationClusterId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationClusterName sets the value of DestinationClusterName for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertyDestinationClusterName(value string) (err error) {
	return instance.SetProperty("DestinationClusterName", value)
}

// GetDestinationClusterName gets the value of DestinationClusterName for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertyDestinationClusterName() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationClusterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertyId() (value uint64, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertyName() (value string, err error) {
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

// SetPartnershipId sets the value of PartnershipId for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertyPartnershipId(value string) (err error) {
	return instance.SetProperty("PartnershipId", value)
}

// GetPartnershipId gets the value of PartnershipId for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertyPartnershipId() (value string, err error) {
	retValue, err := instance.GetProperty("PartnershipId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceClusterId sets the value of SourceClusterId for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertySourceClusterId(value uint64) (err error) {
	return instance.SetProperty("SourceClusterId", value)
}

// GetSourceClusterId gets the value of SourceClusterId for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertySourceClusterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("SourceClusterId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceClusterName sets the value of SourceClusterName for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertySourceClusterName(value string) (err error) {
	return instance.SetProperty("SourceClusterName", value)
}

// GetSourceClusterName gets the value of SourceClusterName for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertySourceClusterName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceClusterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_ClusterSetSRPartnership) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_ClusterSetSRPartnership) GetPropertyState() (value uint32, err error) {
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

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetSRPartnership) RemoveSRPartnership( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveSRPartnership", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetSRPartnership) SetSRPartnershipProperties( /* IN */ Name string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSRPartnershipProperties", Name, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
