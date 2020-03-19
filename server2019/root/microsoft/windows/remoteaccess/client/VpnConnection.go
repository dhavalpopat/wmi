// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// VpnConnection struct
type VpnConnection struct {
	*VpnCommonConfig

	//
	AllUserConnection bool

	//
	AuthenticationMethod []string

	//
	EapConfigXmlStream string

	//
	EncryptionLevel string

	//
	IPSecCustomPolicy []VpnConnectionIPsecConfiguration

	//
	L2tpIPsecAuth string

	//
	MachineCertificateEKUFilter []string

	//
	MachineCertificateIssuerFilter []uint8

	//
	NapState string

	//
	TunnelType string

	//
	UseWinlogonCredential bool

	//
	VpnConfigurationXml string
}

func NewVpnConnectionEx1(instance *cim.WmiInstance) (newInstance *VpnConnection, err error) {
	tmp, err := NewVpnCommonConfigEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnConnection{
		VpnCommonConfig: tmp,
	}
	return
}

func NewVpnConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnConnection, err error) {
	tmp, err := NewVpnCommonConfigEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnConnection{
		VpnCommonConfig: tmp,
	}
	return
}

// SetAllUserConnection sets the value of AllUserConnection for the instance
func (instance *VpnConnection) SetPropertyAllUserConnection(value bool) (err error) {
	return instance.SetProperty("AllUserConnection", value)
}

// GetAllUserConnection gets the value of AllUserConnection for the instance
func (instance *VpnConnection) GetPropertyAllUserConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllUserConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *VpnConnection) SetPropertyAuthenticationMethod(value []string) (err error) {
	return instance.SetProperty("AuthenticationMethod", value)
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *VpnConnection) GetPropertyAuthenticationMethod() (value []string, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEapConfigXmlStream sets the value of EapConfigXmlStream for the instance
func (instance *VpnConnection) SetPropertyEapConfigXmlStream(value string) (err error) {
	return instance.SetProperty("EapConfigXmlStream", value)
}

// GetEapConfigXmlStream gets the value of EapConfigXmlStream for the instance
func (instance *VpnConnection) GetPropertyEapConfigXmlStream() (value string, err error) {
	retValue, err := instance.GetProperty("EapConfigXmlStream")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionLevel sets the value of EncryptionLevel for the instance
func (instance *VpnConnection) SetPropertyEncryptionLevel(value string) (err error) {
	return instance.SetProperty("EncryptionLevel", value)
}

// GetEncryptionLevel gets the value of EncryptionLevel for the instance
func (instance *VpnConnection) GetPropertyEncryptionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPSecCustomPolicy sets the value of IPSecCustomPolicy for the instance
func (instance *VpnConnection) SetPropertyIPSecCustomPolicy(value []VpnConnectionIPsecConfiguration) (err error) {
	return instance.SetProperty("IPSecCustomPolicy", value)
}

// GetIPSecCustomPolicy gets the value of IPSecCustomPolicy for the instance
func (instance *VpnConnection) GetPropertyIPSecCustomPolicy() (value []VpnConnectionIPsecConfiguration, err error) {
	retValue, err := instance.GetProperty("IPSecCustomPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.([]VpnConnectionIPsecConfiguration)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetL2tpIPsecAuth sets the value of L2tpIPsecAuth for the instance
func (instance *VpnConnection) SetPropertyL2tpIPsecAuth(value string) (err error) {
	return instance.SetProperty("L2tpIPsecAuth", value)
}

// GetL2tpIPsecAuth gets the value of L2tpIPsecAuth for the instance
func (instance *VpnConnection) GetPropertyL2tpIPsecAuth() (value string, err error) {
	retValue, err := instance.GetProperty("L2tpIPsecAuth")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineCertificateEKUFilter sets the value of MachineCertificateEKUFilter for the instance
func (instance *VpnConnection) SetPropertyMachineCertificateEKUFilter(value []string) (err error) {
	return instance.SetProperty("MachineCertificateEKUFilter", value)
}

// GetMachineCertificateEKUFilter gets the value of MachineCertificateEKUFilter for the instance
func (instance *VpnConnection) GetPropertyMachineCertificateEKUFilter() (value []string, err error) {
	retValue, err := instance.GetProperty("MachineCertificateEKUFilter")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineCertificateIssuerFilter sets the value of MachineCertificateIssuerFilter for the instance
func (instance *VpnConnection) SetPropertyMachineCertificateIssuerFilter(value []uint8) (err error) {
	return instance.SetProperty("MachineCertificateIssuerFilter", value)
}

// GetMachineCertificateIssuerFilter gets the value of MachineCertificateIssuerFilter for the instance
func (instance *VpnConnection) GetPropertyMachineCertificateIssuerFilter() (value []uint8, err error) {
	retValue, err := instance.GetProperty("MachineCertificateIssuerFilter")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNapState sets the value of NapState for the instance
func (instance *VpnConnection) SetPropertyNapState(value string) (err error) {
	return instance.SetProperty("NapState", value)
}

// GetNapState gets the value of NapState for the instance
func (instance *VpnConnection) GetPropertyNapState() (value string, err error) {
	retValue, err := instance.GetProperty("NapState")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTunnelType sets the value of TunnelType for the instance
func (instance *VpnConnection) SetPropertyTunnelType(value string) (err error) {
	return instance.SetProperty("TunnelType", value)
}

// GetTunnelType gets the value of TunnelType for the instance
func (instance *VpnConnection) GetPropertyTunnelType() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseWinlogonCredential sets the value of UseWinlogonCredential for the instance
func (instance *VpnConnection) SetPropertyUseWinlogonCredential(value bool) (err error) {
	return instance.SetProperty("UseWinlogonCredential", value)
}

// GetUseWinlogonCredential gets the value of UseWinlogonCredential for the instance
func (instance *VpnConnection) GetPropertyUseWinlogonCredential() (value bool, err error) {
	retValue, err := instance.GetProperty("UseWinlogonCredential")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnConfigurationXml sets the value of VpnConfigurationXml for the instance
func (instance *VpnConnection) SetPropertyVpnConfigurationXml(value string) (err error) {
	return instance.SetProperty("VpnConfigurationXml", value)
}

// GetVpnConfigurationXml gets the value of VpnConfigurationXml for the instance
func (instance *VpnConnection) GetPropertyVpnConfigurationXml() (value string, err error) {
	retValue, err := instance.GetProperty("VpnConfigurationXml")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
