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

// Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler struct
type Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler struct {
	*Win32_PerfFormattedData

	//
	DspPerSysAvgQueueLength uint64

	//
	DspPerSysHighAvgQueueLength uint64

	//
	DspPerSysHighAvgsecPerDataRequest uint32

	//
	DspPerSysHighCurrentQueueLength uint64

	//
	DspPerSysIdlePerLowAvgQueueLength uint64

	//
	DspPerSysIdlePerLowAvgsecPerDataRequest uint32

	//
	DspPerSysIdlePerLowCurrentQueueLength uint64

	//
	DspPerSysNormalAvgQueueLength uint64

	//
	DspPerSysNormalAvgsecPerDataRequest uint32

	//
	DspPerSysNormalCurrentQueueLength uint64

	//
	DspPerUsrAvgQueueLength uint64

	//
	DspPerUsrHighAvgQueueLength uint64

	//
	DspPerUsrHighAvgsecPerDataRequest uint32

	//
	DspPerUsrHighCurrentQueueLength uint64

	//
	DspPerUsrIdlePerLowAvgQueueLength uint64

	//
	DspPerUsrIdlePerLowAvgsecPerDataRequest uint32

	//
	DspPerUsrIdlePerLowCurrentQueueLength uint64

	//
	DspPerUsrNormalAvgQueueLength uint64

	//
	DspPerUsrNormalAvgsecPerDataRequest uint32

	//
	DspPerUsrNormalCurrentQueueLength uint64

	//
	QuePerSysAvgQueueLength uint64

	//
	QuePerSysHighAvgQueueLength uint64

	//
	QuePerSysHighAvgsecPerDataRequest uint32

	//
	QuePerSysHighBytesPersec uint64

	//
	QuePerSysHighCurrentQueueLength uint64

	//
	QuePerSysHighDataRequestsPersec uint64

	//
	QuePerSysIdlePerLowAvgQueueLength uint64

	//
	QuePerSysIdlePerLowAvgsecPerDataRequest uint32

	//
	QuePerSysIdlePerLowBytesPersec uint64

	//
	QuePerSysIdlePerLowCurrentQueueLength uint64

	//
	QuePerSysIdlePerLowDataRequestsPersec uint64

	//
	QuePerSysNormalAvgQueueLength uint64

	//
	QuePerSysNormalAvgsecPerDataRequest uint32

	//
	QuePerSysNormalBytesPersec uint64

	//
	QuePerSysNormalCurrentQueueLength uint64

	//
	QuePerSysNormalDataRequestsPersec uint64

	//
	QuePerUsrAvgQueueLength uint64

	//
	QuePerUsrHighAvgQueueLength uint64

	//
	QuePerUsrHighAvgsecPerDataRequest uint32

	//
	QuePerUsrHighBytesPersec uint64

	//
	QuePerUsrHighCurrentQueueLength uint64

	//
	QuePerUsrHighDataRequestsPersec uint64

	//
	QuePerUsrIdlePerLowAvgQueueLength uint64

	//
	QuePerUsrIdlePerLowAvgsecPerDataRequest uint32

	//
	QuePerUsrIdlePerLowBytesPersec uint64

	//
	QuePerUsrIdlePerLowCurrentQueueLength uint64

	//
	QuePerUsrIdlePerLowDataRequestsPersec uint64

	//
	QuePerUsrNormalAvgQueueLength uint64

	//
	QuePerUsrNormalAvgsecPerDataRequest uint32

	//
	QuePerUsrNormalBytesPersec uint64

	//
	QuePerUsrNormalCurrentQueueLength uint64

	//
	QuePerUsrNormalDataRequestsPersec uint64
}

func NewWin32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetDspPerSysAvgQueueLength sets the value of DspPerSysAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysAvgQueueLength", (value))
}

// GetDspPerSysAvgQueueLength gets the value of DspPerSysAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysAvgQueueLength")
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

// SetDspPerSysHighAvgQueueLength sets the value of DspPerSysHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysHighAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysHighAvgQueueLength", (value))
}

// GetDspPerSysHighAvgQueueLength gets the value of DspPerSysHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysHighAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysHighAvgQueueLength")
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

// SetDspPerSysHighAvgsecPerDataRequest sets the value of DspPerSysHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysHighAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("DspPerSysHighAvgsecPerDataRequest", (value))
}

// GetDspPerSysHighAvgsecPerDataRequest gets the value of DspPerSysHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysHighAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("DspPerSysHighAvgsecPerDataRequest")
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

// SetDspPerSysHighCurrentQueueLength sets the value of DspPerSysHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysHighCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysHighCurrentQueueLength", (value))
}

// GetDspPerSysHighCurrentQueueLength gets the value of DspPerSysHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysHighCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysHighCurrentQueueLength")
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

// SetDspPerSysIdlePerLowAvgQueueLength sets the value of DspPerSysIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysIdlePerLowAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysIdlePerLowAvgQueueLength", (value))
}

// GetDspPerSysIdlePerLowAvgQueueLength gets the value of DspPerSysIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysIdlePerLowAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysIdlePerLowAvgQueueLength")
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

// SetDspPerSysIdlePerLowAvgsecPerDataRequest sets the value of DspPerSysIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysIdlePerLowAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("DspPerSysIdlePerLowAvgsecPerDataRequest", (value))
}

// GetDspPerSysIdlePerLowAvgsecPerDataRequest gets the value of DspPerSysIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysIdlePerLowAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("DspPerSysIdlePerLowAvgsecPerDataRequest")
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

// SetDspPerSysIdlePerLowCurrentQueueLength sets the value of DspPerSysIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysIdlePerLowCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysIdlePerLowCurrentQueueLength", (value))
}

// GetDspPerSysIdlePerLowCurrentQueueLength gets the value of DspPerSysIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysIdlePerLowCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysIdlePerLowCurrentQueueLength")
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

// SetDspPerSysNormalAvgQueueLength sets the value of DspPerSysNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysNormalAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysNormalAvgQueueLength", (value))
}

// GetDspPerSysNormalAvgQueueLength gets the value of DspPerSysNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysNormalAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysNormalAvgQueueLength")
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

// SetDspPerSysNormalAvgsecPerDataRequest sets the value of DspPerSysNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysNormalAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("DspPerSysNormalAvgsecPerDataRequest", (value))
}

// GetDspPerSysNormalAvgsecPerDataRequest gets the value of DspPerSysNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysNormalAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("DspPerSysNormalAvgsecPerDataRequest")
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

// SetDspPerSysNormalCurrentQueueLength sets the value of DspPerSysNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerSysNormalCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerSysNormalCurrentQueueLength", (value))
}

// GetDspPerSysNormalCurrentQueueLength gets the value of DspPerSysNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerSysNormalCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerSysNormalCurrentQueueLength")
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

// SetDspPerUsrAvgQueueLength sets the value of DspPerUsrAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrAvgQueueLength", (value))
}

// GetDspPerUsrAvgQueueLength gets the value of DspPerUsrAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrAvgQueueLength")
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

// SetDspPerUsrHighAvgQueueLength sets the value of DspPerUsrHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrHighAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrHighAvgQueueLength", (value))
}

// GetDspPerUsrHighAvgQueueLength gets the value of DspPerUsrHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrHighAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrHighAvgQueueLength")
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

// SetDspPerUsrHighAvgsecPerDataRequest sets the value of DspPerUsrHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrHighAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("DspPerUsrHighAvgsecPerDataRequest", (value))
}

// GetDspPerUsrHighAvgsecPerDataRequest gets the value of DspPerUsrHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrHighAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("DspPerUsrHighAvgsecPerDataRequest")
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

// SetDspPerUsrHighCurrentQueueLength sets the value of DspPerUsrHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrHighCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrHighCurrentQueueLength", (value))
}

// GetDspPerUsrHighCurrentQueueLength gets the value of DspPerUsrHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrHighCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrHighCurrentQueueLength")
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

// SetDspPerUsrIdlePerLowAvgQueueLength sets the value of DspPerUsrIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrIdlePerLowAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrIdlePerLowAvgQueueLength", (value))
}

// GetDspPerUsrIdlePerLowAvgQueueLength gets the value of DspPerUsrIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrIdlePerLowAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrIdlePerLowAvgQueueLength")
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

// SetDspPerUsrIdlePerLowAvgsecPerDataRequest sets the value of DspPerUsrIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrIdlePerLowAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("DspPerUsrIdlePerLowAvgsecPerDataRequest", (value))
}

// GetDspPerUsrIdlePerLowAvgsecPerDataRequest gets the value of DspPerUsrIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrIdlePerLowAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("DspPerUsrIdlePerLowAvgsecPerDataRequest")
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

// SetDspPerUsrIdlePerLowCurrentQueueLength sets the value of DspPerUsrIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrIdlePerLowCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrIdlePerLowCurrentQueueLength", (value))
}

// GetDspPerUsrIdlePerLowCurrentQueueLength gets the value of DspPerUsrIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrIdlePerLowCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrIdlePerLowCurrentQueueLength")
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

// SetDspPerUsrNormalAvgQueueLength sets the value of DspPerUsrNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrNormalAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrNormalAvgQueueLength", (value))
}

// GetDspPerUsrNormalAvgQueueLength gets the value of DspPerUsrNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrNormalAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrNormalAvgQueueLength")
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

// SetDspPerUsrNormalAvgsecPerDataRequest sets the value of DspPerUsrNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrNormalAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("DspPerUsrNormalAvgsecPerDataRequest", (value))
}

// GetDspPerUsrNormalAvgsecPerDataRequest gets the value of DspPerUsrNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrNormalAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("DspPerUsrNormalAvgsecPerDataRequest")
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

// SetDspPerUsrNormalCurrentQueueLength sets the value of DspPerUsrNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyDspPerUsrNormalCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("DspPerUsrNormalCurrentQueueLength", (value))
}

// GetDspPerUsrNormalCurrentQueueLength gets the value of DspPerUsrNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyDspPerUsrNormalCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("DspPerUsrNormalCurrentQueueLength")
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

// SetQuePerSysAvgQueueLength sets the value of QuePerSysAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysAvgQueueLength", (value))
}

// GetQuePerSysAvgQueueLength gets the value of QuePerSysAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysAvgQueueLength")
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

// SetQuePerSysHighAvgQueueLength sets the value of QuePerSysHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysHighAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysHighAvgQueueLength", (value))
}

// GetQuePerSysHighAvgQueueLength gets the value of QuePerSysHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysHighAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysHighAvgQueueLength")
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

// SetQuePerSysHighAvgsecPerDataRequest sets the value of QuePerSysHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysHighAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("QuePerSysHighAvgsecPerDataRequest", (value))
}

// GetQuePerSysHighAvgsecPerDataRequest gets the value of QuePerSysHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysHighAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuePerSysHighAvgsecPerDataRequest")
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

// SetQuePerSysHighBytesPersec sets the value of QuePerSysHighBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysHighBytesPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerSysHighBytesPersec", (value))
}

// GetQuePerSysHighBytesPersec gets the value of QuePerSysHighBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysHighBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysHighBytesPersec")
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

// SetQuePerSysHighCurrentQueueLength sets the value of QuePerSysHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysHighCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysHighCurrentQueueLength", (value))
}

// GetQuePerSysHighCurrentQueueLength gets the value of QuePerSysHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysHighCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysHighCurrentQueueLength")
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

// SetQuePerSysHighDataRequestsPersec sets the value of QuePerSysHighDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysHighDataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerSysHighDataRequestsPersec", (value))
}

// GetQuePerSysHighDataRequestsPersec gets the value of QuePerSysHighDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysHighDataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysHighDataRequestsPersec")
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

// SetQuePerSysIdlePerLowAvgQueueLength sets the value of QuePerSysIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysIdlePerLowAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysIdlePerLowAvgQueueLength", (value))
}

// GetQuePerSysIdlePerLowAvgQueueLength gets the value of QuePerSysIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysIdlePerLowAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysIdlePerLowAvgQueueLength")
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

// SetQuePerSysIdlePerLowAvgsecPerDataRequest sets the value of QuePerSysIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysIdlePerLowAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("QuePerSysIdlePerLowAvgsecPerDataRequest", (value))
}

// GetQuePerSysIdlePerLowAvgsecPerDataRequest gets the value of QuePerSysIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysIdlePerLowAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuePerSysIdlePerLowAvgsecPerDataRequest")
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

// SetQuePerSysIdlePerLowBytesPersec sets the value of QuePerSysIdlePerLowBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysIdlePerLowBytesPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerSysIdlePerLowBytesPersec", (value))
}

// GetQuePerSysIdlePerLowBytesPersec gets the value of QuePerSysIdlePerLowBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysIdlePerLowBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysIdlePerLowBytesPersec")
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

// SetQuePerSysIdlePerLowCurrentQueueLength sets the value of QuePerSysIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysIdlePerLowCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysIdlePerLowCurrentQueueLength", (value))
}

// GetQuePerSysIdlePerLowCurrentQueueLength gets the value of QuePerSysIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysIdlePerLowCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysIdlePerLowCurrentQueueLength")
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

// SetQuePerSysIdlePerLowDataRequestsPersec sets the value of QuePerSysIdlePerLowDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysIdlePerLowDataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerSysIdlePerLowDataRequestsPersec", (value))
}

// GetQuePerSysIdlePerLowDataRequestsPersec gets the value of QuePerSysIdlePerLowDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysIdlePerLowDataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysIdlePerLowDataRequestsPersec")
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

// SetQuePerSysNormalAvgQueueLength sets the value of QuePerSysNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysNormalAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysNormalAvgQueueLength", (value))
}

// GetQuePerSysNormalAvgQueueLength gets the value of QuePerSysNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysNormalAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysNormalAvgQueueLength")
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

// SetQuePerSysNormalAvgsecPerDataRequest sets the value of QuePerSysNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysNormalAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("QuePerSysNormalAvgsecPerDataRequest", (value))
}

// GetQuePerSysNormalAvgsecPerDataRequest gets the value of QuePerSysNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysNormalAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuePerSysNormalAvgsecPerDataRequest")
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

// SetQuePerSysNormalBytesPersec sets the value of QuePerSysNormalBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysNormalBytesPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerSysNormalBytesPersec", (value))
}

// GetQuePerSysNormalBytesPersec gets the value of QuePerSysNormalBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysNormalBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysNormalBytesPersec")
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

// SetQuePerSysNormalCurrentQueueLength sets the value of QuePerSysNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysNormalCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerSysNormalCurrentQueueLength", (value))
}

// GetQuePerSysNormalCurrentQueueLength gets the value of QuePerSysNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysNormalCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysNormalCurrentQueueLength")
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

// SetQuePerSysNormalDataRequestsPersec sets the value of QuePerSysNormalDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerSysNormalDataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerSysNormalDataRequestsPersec", (value))
}

// GetQuePerSysNormalDataRequestsPersec gets the value of QuePerSysNormalDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerSysNormalDataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerSysNormalDataRequestsPersec")
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

// SetQuePerUsrAvgQueueLength sets the value of QuePerUsrAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrAvgQueueLength", (value))
}

// GetQuePerUsrAvgQueueLength gets the value of QuePerUsrAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrAvgQueueLength")
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

// SetQuePerUsrHighAvgQueueLength sets the value of QuePerUsrHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrHighAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrHighAvgQueueLength", (value))
}

// GetQuePerUsrHighAvgQueueLength gets the value of QuePerUsrHighAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrHighAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrHighAvgQueueLength")
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

// SetQuePerUsrHighAvgsecPerDataRequest sets the value of QuePerUsrHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrHighAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("QuePerUsrHighAvgsecPerDataRequest", (value))
}

// GetQuePerUsrHighAvgsecPerDataRequest gets the value of QuePerUsrHighAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrHighAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuePerUsrHighAvgsecPerDataRequest")
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

// SetQuePerUsrHighBytesPersec sets the value of QuePerUsrHighBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrHighBytesPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrHighBytesPersec", (value))
}

// GetQuePerUsrHighBytesPersec gets the value of QuePerUsrHighBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrHighBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrHighBytesPersec")
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

// SetQuePerUsrHighCurrentQueueLength sets the value of QuePerUsrHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrHighCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrHighCurrentQueueLength", (value))
}

// GetQuePerUsrHighCurrentQueueLength gets the value of QuePerUsrHighCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrHighCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrHighCurrentQueueLength")
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

// SetQuePerUsrHighDataRequestsPersec sets the value of QuePerUsrHighDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrHighDataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrHighDataRequestsPersec", (value))
}

// GetQuePerUsrHighDataRequestsPersec gets the value of QuePerUsrHighDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrHighDataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrHighDataRequestsPersec")
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

// SetQuePerUsrIdlePerLowAvgQueueLength sets the value of QuePerUsrIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrIdlePerLowAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrIdlePerLowAvgQueueLength", (value))
}

// GetQuePerUsrIdlePerLowAvgQueueLength gets the value of QuePerUsrIdlePerLowAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrIdlePerLowAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrIdlePerLowAvgQueueLength")
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

// SetQuePerUsrIdlePerLowAvgsecPerDataRequest sets the value of QuePerUsrIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrIdlePerLowAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("QuePerUsrIdlePerLowAvgsecPerDataRequest", (value))
}

// GetQuePerUsrIdlePerLowAvgsecPerDataRequest gets the value of QuePerUsrIdlePerLowAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrIdlePerLowAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuePerUsrIdlePerLowAvgsecPerDataRequest")
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

// SetQuePerUsrIdlePerLowBytesPersec sets the value of QuePerUsrIdlePerLowBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrIdlePerLowBytesPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrIdlePerLowBytesPersec", (value))
}

// GetQuePerUsrIdlePerLowBytesPersec gets the value of QuePerUsrIdlePerLowBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrIdlePerLowBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrIdlePerLowBytesPersec")
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

// SetQuePerUsrIdlePerLowCurrentQueueLength sets the value of QuePerUsrIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrIdlePerLowCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrIdlePerLowCurrentQueueLength", (value))
}

// GetQuePerUsrIdlePerLowCurrentQueueLength gets the value of QuePerUsrIdlePerLowCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrIdlePerLowCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrIdlePerLowCurrentQueueLength")
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

// SetQuePerUsrIdlePerLowDataRequestsPersec sets the value of QuePerUsrIdlePerLowDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrIdlePerLowDataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrIdlePerLowDataRequestsPersec", (value))
}

// GetQuePerUsrIdlePerLowDataRequestsPersec gets the value of QuePerUsrIdlePerLowDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrIdlePerLowDataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrIdlePerLowDataRequestsPersec")
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

// SetQuePerUsrNormalAvgQueueLength sets the value of QuePerUsrNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrNormalAvgQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrNormalAvgQueueLength", (value))
}

// GetQuePerUsrNormalAvgQueueLength gets the value of QuePerUsrNormalAvgQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrNormalAvgQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrNormalAvgQueueLength")
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

// SetQuePerUsrNormalAvgsecPerDataRequest sets the value of QuePerUsrNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrNormalAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("QuePerUsrNormalAvgsecPerDataRequest", (value))
}

// GetQuePerUsrNormalAvgsecPerDataRequest gets the value of QuePerUsrNormalAvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrNormalAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuePerUsrNormalAvgsecPerDataRequest")
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

// SetQuePerUsrNormalBytesPersec sets the value of QuePerUsrNormalBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrNormalBytesPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrNormalBytesPersec", (value))
}

// GetQuePerUsrNormalBytesPersec gets the value of QuePerUsrNormalBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrNormalBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrNormalBytesPersec")
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

// SetQuePerUsrNormalCurrentQueueLength sets the value of QuePerUsrNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrNormalCurrentQueueLength(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrNormalCurrentQueueLength", (value))
}

// GetQuePerUsrNormalCurrentQueueLength gets the value of QuePerUsrNormalCurrentQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrNormalCurrentQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrNormalCurrentQueueLength")
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

// SetQuePerUsrNormalDataRequestsPersec sets the value of QuePerUsrNormalDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) SetPropertyQuePerUsrNormalDataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("QuePerUsrNormalDataRequestsPersec", (value))
}

// GetQuePerUsrNormalDataRequestsPersec gets the value of QuePerUsrNormalDataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskScheduler) GetPropertyQuePerUsrNormalDataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuePerUsrNormalDataRequestsPersec")
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
