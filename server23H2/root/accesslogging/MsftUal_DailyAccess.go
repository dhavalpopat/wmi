// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.AccessLogging
//////////////////////////////////////////////
package accesslogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsftUal_DailyAccess struct
type MsftUal_DailyAccess struct {
	*cim.WmiInstance

	// The number of accesses of a role, or installed product, on the local server from a unique client device.
	AccessCount uint32

	// The date that a device accessed a role, or installed product, on the local server.
	AccessDate string

	// The IP address of a client device that accompanies the UAL payload from installed roles and products.
	IPAddress string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string

	// A unique GUID for a tenant client of an installed role or product which accompanies the UAL data payload, if applicable.
	TenantIdentifier string

	// The client user name that accompanies the UAL payload from installed roles and products, if applicable.
	UserName string
}

func NewMsftUal_DailyAccessEx1(instance *cim.WmiInstance) (newInstance *MsftUal_DailyAccess, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_DailyAccess{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_DailyAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_DailyAccess, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_DailyAccess{
		WmiInstance: tmp,
	}
	return
}

// SetAccessCount sets the value of AccessCount for the instance
func (instance *MsftUal_DailyAccess) SetPropertyAccessCount(value uint32) (err error) {
	return instance.SetProperty("AccessCount", (value))
}

// GetAccessCount gets the value of AccessCount for the instance
func (instance *MsftUal_DailyAccess) GetPropertyAccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAccessDate sets the value of AccessDate for the instance
func (instance *MsftUal_DailyAccess) SetPropertyAccessDate(value string) (err error) {
	return instance.SetProperty("AccessDate", (value))
}

// GetAccessDate gets the value of AccessDate for the instance
func (instance *MsftUal_DailyAccess) GetPropertyAccessDate() (value string, err error) {
	retValue, err := instance.GetProperty("AccessDate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MsftUal_DailyAccess) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_DailyAccess) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProductName sets the value of ProductName for the instance
func (instance *MsftUal_DailyAccess) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_DailyAccess) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRoleGuid sets the value of RoleGuid for the instance
func (instance *MsftUal_DailyAccess) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", (value))
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_DailyAccess) GetPropertyRoleGuid() (value string, err error) {
	retValue, err := instance.GetProperty("RoleGuid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftUal_DailyAccess) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", (value))
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_DailyAccess) GetPropertyRoleName() (value string, err error) {
	retValue, err := instance.GetProperty("RoleName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTenantIdentifier sets the value of TenantIdentifier for the instance
func (instance *MsftUal_DailyAccess) SetPropertyTenantIdentifier(value string) (err error) {
	return instance.SetProperty("TenantIdentifier", (value))
}

// GetTenantIdentifier gets the value of TenantIdentifier for the instance
func (instance *MsftUal_DailyAccess) GetPropertyTenantIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("TenantIdentifier")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetUserName sets the value of UserName for the instance
func (instance *MsftUal_DailyAccess) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *MsftUal_DailyAccess) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
