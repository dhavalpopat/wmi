// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetImPlatTeam struct
type MSFT_NetImPlatTeam struct {
	*CIM_ManagedElement

	//
	Name string
}

func NewMSFT_NetImPlatTeamEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetImPlatTeam, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetImPlatTeam{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetImPlatTeamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetImPlatTeam, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetImPlatTeam{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_NetImPlatTeam) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetImPlatTeam) GetPropertyName() (value string, err error) {
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
