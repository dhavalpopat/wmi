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

// Win32_PerfRawData_ESENT_DatabaseDatabases struct
type Win32_PerfRawData_ESENT_DatabaseDatabases struct {
	*Win32_PerfRawData

	//
	DatabaseCachePercentHitUnique uint32

	//
	DatabaseCachePercentHitUnique_Base uint32

	//
	DatabaseCacheRequestsPersecUnique uint32

	//
	DatabaseCacheSizeMB uint64

	//
	IODatabaseReadsAverageLatency uint64

	//
	IODatabaseReadsAverageLatency_Base uint32

	//
	IODatabaseReadsPersec uint32

	//
	IODatabaseWritesAverageLatency uint64

	//
	IODatabaseWritesAverageLatency_Base uint32

	//
	IODatabaseWritesPersec uint32
}

func NewWin32_PerfRawData_ESENT_DatabaseDatabasesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_ESENT_DatabaseDatabases, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ESENT_DatabaseDatabases{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_ESENT_DatabaseDatabasesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_ESENT_DatabaseDatabases, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ESENT_DatabaseDatabases{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetDatabaseCachePercentHitUnique sets the value of DatabaseCachePercentHitUnique for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyDatabaseCachePercentHitUnique(value uint32) (err error) {
	return instance.SetProperty("DatabaseCachePercentHitUnique", (value))
}

// GetDatabaseCachePercentHitUnique gets the value of DatabaseCachePercentHitUnique for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyDatabaseCachePercentHitUnique() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCachePercentHitUnique")
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

// SetDatabaseCachePercentHitUnique_Base sets the value of DatabaseCachePercentHitUnique_Base for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyDatabaseCachePercentHitUnique_Base(value uint32) (err error) {
	return instance.SetProperty("DatabaseCachePercentHitUnique_Base", (value))
}

// GetDatabaseCachePercentHitUnique_Base gets the value of DatabaseCachePercentHitUnique_Base for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyDatabaseCachePercentHitUnique_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCachePercentHitUnique_Base")
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

// SetDatabaseCacheRequestsPersecUnique sets the value of DatabaseCacheRequestsPersecUnique for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyDatabaseCacheRequestsPersecUnique(value uint32) (err error) {
	return instance.SetProperty("DatabaseCacheRequestsPersecUnique", (value))
}

// GetDatabaseCacheRequestsPersecUnique gets the value of DatabaseCacheRequestsPersecUnique for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyDatabaseCacheRequestsPersecUnique() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheRequestsPersecUnique")
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

// SetDatabaseCacheSizeMB sets the value of DatabaseCacheSizeMB for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyDatabaseCacheSizeMB(value uint64) (err error) {
	return instance.SetProperty("DatabaseCacheSizeMB", (value))
}

// GetDatabaseCacheSizeMB gets the value of DatabaseCacheSizeMB for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyDatabaseCacheSizeMB() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheSizeMB")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIODatabaseReadsAverageLatency sets the value of IODatabaseReadsAverageLatency for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyIODatabaseReadsAverageLatency(value uint64) (err error) {
	return instance.SetProperty("IODatabaseReadsAverageLatency", (value))
}

// GetIODatabaseReadsAverageLatency gets the value of IODatabaseReadsAverageLatency for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyIODatabaseReadsAverageLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("IODatabaseReadsAverageLatency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIODatabaseReadsAverageLatency_Base sets the value of IODatabaseReadsAverageLatency_Base for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyIODatabaseReadsAverageLatency_Base(value uint32) (err error) {
	return instance.SetProperty("IODatabaseReadsAverageLatency_Base", (value))
}

// GetIODatabaseReadsAverageLatency_Base gets the value of IODatabaseReadsAverageLatency_Base for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyIODatabaseReadsAverageLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("IODatabaseReadsAverageLatency_Base")
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

// SetIODatabaseReadsPersec sets the value of IODatabaseReadsPersec for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyIODatabaseReadsPersec(value uint32) (err error) {
	return instance.SetProperty("IODatabaseReadsPersec", (value))
}

// GetIODatabaseReadsPersec gets the value of IODatabaseReadsPersec for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyIODatabaseReadsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IODatabaseReadsPersec")
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

// SetIODatabaseWritesAverageLatency sets the value of IODatabaseWritesAverageLatency for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyIODatabaseWritesAverageLatency(value uint64) (err error) {
	return instance.SetProperty("IODatabaseWritesAverageLatency", (value))
}

// GetIODatabaseWritesAverageLatency gets the value of IODatabaseWritesAverageLatency for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyIODatabaseWritesAverageLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("IODatabaseWritesAverageLatency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIODatabaseWritesAverageLatency_Base sets the value of IODatabaseWritesAverageLatency_Base for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyIODatabaseWritesAverageLatency_Base(value uint32) (err error) {
	return instance.SetProperty("IODatabaseWritesAverageLatency_Base", (value))
}

// GetIODatabaseWritesAverageLatency_Base gets the value of IODatabaseWritesAverageLatency_Base for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyIODatabaseWritesAverageLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("IODatabaseWritesAverageLatency_Base")
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

// SetIODatabaseWritesPersec sets the value of IODatabaseWritesPersec for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) SetPropertyIODatabaseWritesPersec(value uint32) (err error) {
	return instance.SetProperty("IODatabaseWritesPersec", (value))
}

// GetIODatabaseWritesPersec gets the value of IODatabaseWritesPersec for the instance
func (instance *Win32_PerfRawData_ESENT_DatabaseDatabases) GetPropertyIODatabaseWritesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IODatabaseWritesPersec")
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
