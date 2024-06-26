// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle struct
type Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle struct {
	*Win32_PerfFormattedData

	//
	HardConnectsPerSecond uint32

	//
	HardDisconnectsPerSecond uint32

	//
	NumberOfActiveConnectionPoolGroups uint32

	//
	NumberOfActiveConnectionPools uint32

	//
	NumberOfActiveConnections uint32

	//
	NumberOfFreeConnections uint32

	//
	NumberOfInactiveConnectionPoolGroups uint32

	//
	NumberOfInactiveConnectionPools uint32

	//
	NumberOfNonPooledConnections uint32

	//
	NumberOfPooledConnections uint32

	//
	NumberOfReclaimedConnections uint32

	//
	NumberOfStasisConnections uint32

	//
	SoftConnectsPerSecond uint32

	//
	SoftDisconnectsPerSecond uint32
}

func NewWin32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracleEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetHardConnectsPerSecond sets the value of HardConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyHardConnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("HardConnectsPerSecond", (value))
}

// GetHardConnectsPerSecond gets the value of HardConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyHardConnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardConnectsPerSecond")
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

// SetHardDisconnectsPerSecond sets the value of HardDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyHardDisconnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("HardDisconnectsPerSecond", (value))
}

// GetHardDisconnectsPerSecond gets the value of HardDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyHardDisconnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardDisconnectsPerSecond")
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

// SetNumberOfActiveConnectionPoolGroups sets the value of NumberOfActiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfActiveConnectionPoolGroups(value uint32) (err error) {
	return instance.SetProperty("NumberOfActiveConnectionPoolGroups", (value))
}

// GetNumberOfActiveConnectionPoolGroups gets the value of NumberOfActiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfActiveConnectionPoolGroups() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfActiveConnectionPoolGroups")
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

// SetNumberOfActiveConnectionPools sets the value of NumberOfActiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfActiveConnectionPools(value uint32) (err error) {
	return instance.SetProperty("NumberOfActiveConnectionPools", (value))
}

// GetNumberOfActiveConnectionPools gets the value of NumberOfActiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfActiveConnectionPools() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfActiveConnectionPools")
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

// SetNumberOfActiveConnections sets the value of NumberOfActiveConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfActiveConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfActiveConnections", (value))
}

// GetNumberOfActiveConnections gets the value of NumberOfActiveConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfActiveConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfActiveConnections")
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

// SetNumberOfFreeConnections sets the value of NumberOfFreeConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfFreeConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfFreeConnections", (value))
}

// GetNumberOfFreeConnections gets the value of NumberOfFreeConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfFreeConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfFreeConnections")
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

// SetNumberOfInactiveConnectionPoolGroups sets the value of NumberOfInactiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfInactiveConnectionPoolGroups(value uint32) (err error) {
	return instance.SetProperty("NumberOfInactiveConnectionPoolGroups", (value))
}

// GetNumberOfInactiveConnectionPoolGroups gets the value of NumberOfInactiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfInactiveConnectionPoolGroups() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInactiveConnectionPoolGroups")
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

// SetNumberOfInactiveConnectionPools sets the value of NumberOfInactiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfInactiveConnectionPools(value uint32) (err error) {
	return instance.SetProperty("NumberOfInactiveConnectionPools", (value))
}

// GetNumberOfInactiveConnectionPools gets the value of NumberOfInactiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfInactiveConnectionPools() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInactiveConnectionPools")
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

// SetNumberOfNonPooledConnections sets the value of NumberOfNonPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfNonPooledConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfNonPooledConnections", (value))
}

// GetNumberOfNonPooledConnections gets the value of NumberOfNonPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfNonPooledConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfNonPooledConnections")
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

// SetNumberOfPooledConnections sets the value of NumberOfPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfPooledConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfPooledConnections", (value))
}

// GetNumberOfPooledConnections gets the value of NumberOfPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfPooledConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPooledConnections")
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

// SetNumberOfReclaimedConnections sets the value of NumberOfReclaimedConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfReclaimedConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfReclaimedConnections", (value))
}

// GetNumberOfReclaimedConnections gets the value of NumberOfReclaimedConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfReclaimedConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfReclaimedConnections")
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

// SetNumberOfStasisConnections sets the value of NumberOfStasisConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfStasisConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfStasisConnections", (value))
}

// GetNumberOfStasisConnections gets the value of NumberOfStasisConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfStasisConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfStasisConnections")
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

// SetSoftConnectsPerSecond sets the value of SoftConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertySoftConnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("SoftConnectsPerSecond", (value))
}

// GetSoftConnectsPerSecond gets the value of SoftConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertySoftConnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("SoftConnectsPerSecond")
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

// SetSoftDisconnectsPerSecond sets the value of SoftDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertySoftDisconnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("SoftDisconnectsPerSecond", (value))
}

// GetSoftDisconnectsPerSecond gets the value of SoftDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertySoftDisconnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("SoftDisconnectsPerSecond")
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
