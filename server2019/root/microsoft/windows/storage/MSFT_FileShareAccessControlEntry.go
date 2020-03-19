// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_FileShareAccessControlEntry struct
type MSFT_FileShareAccessControlEntry struct {
	*cim.WmiInstance

	//
	AccessControlType uint16

	//
	AccessRight uint16

	//
	AccountName string
}

func NewMSFT_FileShareAccessControlEntryEx1(instance *cim.WmiInstance) (newInstance *MSFT_FileShareAccessControlEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_FileShareAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_FileShareAccessControlEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_FileShareAccessControlEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileShareAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

// SetAccessControlType sets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessControlType(value uint16) (err error) {
	return instance.SetProperty("AccessControlType", value)
}

// GetAccessControlType gets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessControlType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessControlType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessRight sets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessRight(value uint16) (err error) {
	return instance.SetProperty("AccessRight", value)
}

// GetAccessRight gets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessRight() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessRight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", value)
}

// GetAccountName gets the value of AccountName for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("AccountName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
