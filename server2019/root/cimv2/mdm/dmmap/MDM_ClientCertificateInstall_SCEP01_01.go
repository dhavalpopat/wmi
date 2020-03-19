// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_ClientCertificateInstall_SCEP01_01 struct
type MDM_ClientCertificateInstall_SCEP01_01 struct {
	*cim.WmiInstance

	//
	CertThumbprint string

	//
	ErrorCode int32

	//
	InstanceID string

	//
	ParentID string

	//
	RespondentServerUrl string

	//
	Status int32
}

func NewMDM_ClientCertificateInstall_SCEP01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_ClientCertificateInstall_SCEP01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_ClientCertificateInstall_SCEP01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ClientCertificateInstall_SCEP01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_ClientCertificateInstall_SCEP01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_ClientCertificateInstall_SCEP01_01{
		WmiInstance: tmp,
	}
	return
}

// SetCertThumbprint sets the value of CertThumbprint for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) SetPropertyCertThumbprint(value string) (err error) {
	return instance.SetProperty("CertThumbprint", value)
}

// GetCertThumbprint gets the value of CertThumbprint for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) GetPropertyCertThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("CertThumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) SetPropertyErrorCode(value int32) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) GetPropertyErrorCode() (value int32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_ClientCertificateInstall_SCEP01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) GetPropertyParentID() (value string, err error) {
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

// SetRespondentServerUrl sets the value of RespondentServerUrl for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) SetPropertyRespondentServerUrl(value string) (err error) {
	return instance.SetProperty("RespondentServerUrl", value)
}

// GetRespondentServerUrl gets the value of RespondentServerUrl for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) GetPropertyRespondentServerUrl() (value string, err error) {
	retValue, err := instance.GetProperty("RespondentServerUrl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_ClientCertificateInstall_SCEP01_01) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
