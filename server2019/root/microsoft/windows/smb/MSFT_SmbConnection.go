// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbConnection struct
type MSFT_SmbConnection struct {
	cim.WmiInstance

	//
	ContinuouslyAvailable bool

	//
	Credential string

	//
	Dialect string

	//
	Encrypted bool

	//
	NumOpens uint64

	//
	Redirected bool

	//
	ServerName string

	//
	ShareName string

	//
	Signed bool

	//
	SmbInstance SmbConnection_SmbInstance

	//
	UserName string
}

// SetContinuouslyAvailable sets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_SmbConnection) SetPropertyContinuouslyAvailable(value bool) (err error) {
	return instance.SetProperty("ContinuouslyAvailable", value)
}

// GetContinuouslyAvailable gets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_SmbConnection) GetPropertyContinuouslyAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("ContinuouslyAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCredential sets the value of Credential for the instance
func (instance *MSFT_SmbConnection) SetPropertyCredential(value string) (err error) {
	return instance.SetProperty("Credential", value)
}

// GetCredential gets the value of Credential for the instance
func (instance *MSFT_SmbConnection) GetPropertyCredential() (value string, err error) {
	retValue, err := instance.GetProperty("Credential")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDialect sets the value of Dialect for the instance
func (instance *MSFT_SmbConnection) SetPropertyDialect(value string) (err error) {
	return instance.SetProperty("Dialect", value)
}

// GetDialect gets the value of Dialect for the instance
func (instance *MSFT_SmbConnection) GetPropertyDialect() (value string, err error) {
	retValue, err := instance.GetProperty("Dialect")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncrypted sets the value of Encrypted for the instance
func (instance *MSFT_SmbConnection) SetPropertyEncrypted(value bool) (err error) {
	return instance.SetProperty("Encrypted", value)
}

// GetEncrypted gets the value of Encrypted for the instance
func (instance *MSFT_SmbConnection) GetPropertyEncrypted() (value bool, err error) {
	retValue, err := instance.GetProperty("Encrypted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumOpens sets the value of NumOpens for the instance
func (instance *MSFT_SmbConnection) SetPropertyNumOpens(value uint64) (err error) {
	return instance.SetProperty("NumOpens", value)
}

// GetNumOpens gets the value of NumOpens for the instance
func (instance *MSFT_SmbConnection) GetPropertyNumOpens() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumOpens")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRedirected sets the value of Redirected for the instance
func (instance *MSFT_SmbConnection) SetPropertyRedirected(value bool) (err error) {
	return instance.SetProperty("Redirected", value)
}

// GetRedirected gets the value of Redirected for the instance
func (instance *MSFT_SmbConnection) GetPropertyRedirected() (value bool, err error) {
	retValue, err := instance.GetProperty("Redirected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerName sets the value of ServerName for the instance
func (instance *MSFT_SmbConnection) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *MSFT_SmbConnection) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareName sets the value of ShareName for the instance
func (instance *MSFT_SmbConnection) SetPropertyShareName(value string) (err error) {
	return instance.SetProperty("ShareName", value)
}

// GetShareName gets the value of ShareName for the instance
func (instance *MSFT_SmbConnection) GetPropertyShareName() (value string, err error) {
	retValue, err := instance.GetProperty("ShareName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSigned sets the value of Signed for the instance
func (instance *MSFT_SmbConnection) SetPropertySigned(value bool) (err error) {
	return instance.SetProperty("Signed", value)
}

// GetSigned gets the value of Signed for the instance
func (instance *MSFT_SmbConnection) GetPropertySigned() (value bool, err error) {
	retValue, err := instance.GetProperty("Signed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmbInstance sets the value of SmbInstance for the instance
func (instance *MSFT_SmbConnection) SetPropertySmbInstance(value SmbConnection_SmbInstance) (err error) {
	return instance.SetProperty("SmbInstance", value)
}

// GetSmbInstance gets the value of SmbInstance for the instance
func (instance *MSFT_SmbConnection) GetPropertySmbInstance() (value SmbConnection_SmbInstance, err error) {
	retValue, err := instance.GetProperty("SmbInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbConnection_SmbInstance)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *MSFT_SmbConnection) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *MSFT_SmbConnection) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
