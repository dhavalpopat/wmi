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

// Win32_PerfRawData_Counters_SMBServerSessions struct
type Win32_PerfRawData_Counters_SMBServerSessions struct {
	*Win32_PerfRawData

	//
	AvgBytesPerRead uint64

	//
	AvgBytesPerRead_Base uint32

	//
	AvgBytesPerWrite uint64

	//
	AvgBytesPerWrite_Base uint32

	//
	AvgDataBytesPerRequest uint64

	//
	AvgDataBytesPerRequest_Base uint32

	//
	AvgDataQueueLength uint64

	//
	AvgReadQueueLength uint64

	//
	AvgsecPerDataRequest uint32

	//
	AvgsecPerDataRequest_Base uint32

	//
	AvgsecPerRead uint32

	//
	AvgsecPerRead_Base uint32

	//
	AvgsecPerRequest uint32

	//
	AvgsecPerRequest_Base uint32

	//
	AvgsecPerWrite uint32

	//
	AvgsecPerWrite_Base uint32

	//
	AvgWriteQueueLength uint64

	//
	CurrentDataQueueLength uint64

	//
	CurrentDurableOpenFileCount uint64

	//
	CurrentOpenFileCount uint64

	//
	CurrentPendingRequests uint64

	//
	DataBytesPersec uint64

	//
	DataRequestsPersec uint32

	//
	FilesOpenedPersec uint64

	//
	MetadataRequestsPersec uint64

	//
	PercentPersistentHandles uint64

	//
	PercentPersistentHandles_Base uint64

	//
	PercentResilientHandles uint64

	//
	PercentResilientHandles_Base uint64

	//
	ReadBytesPersec uint64

	//
	ReadRequestsPersec uint32

	//
	ReceivedBytesPersec uint64

	//
	RequestsPersec uint64

	//
	SentBytesPersec uint64

	//
	TotalDurableHandleReopenCount uint64

	//
	TotalFailedDurableHandleReopenCount uint64

	//
	TotalFailedPersistentHandleReopenCount uint64

	//
	TotalFailedResilientHandleReopenCount uint64

	//
	TotalFileOpenCount uint64

	//
	TotalPersistentHandleReopenCount uint64

	//
	TotalResilientHandleReopenCount uint64

	//
	TransferredBytesPersec uint64

	//
	TreeConnectCount uint64

	//
	WriteBytesPersec uint64

	//
	WriteRequestsPersec uint32
}

func NewWin32_PerfRawData_Counters_SMBServerSessionsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_SMBServerSessions, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_SMBServerSessions{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_SMBServerSessionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_SMBServerSessions, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_SMBServerSessions{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAvgBytesPerRead sets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgBytesPerRead(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerRead", (value))
}

// GetAvgBytesPerRead gets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgBytesPerRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerRead")
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

// SetAvgBytesPerRead_Base sets the value of AvgBytesPerRead_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgBytesPerRead_Base(value uint32) (err error) {
	return instance.SetProperty("AvgBytesPerRead_Base", (value))
}

// GetAvgBytesPerRead_Base gets the value of AvgBytesPerRead_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgBytesPerRead_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerRead_Base")
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

// SetAvgBytesPerWrite sets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgBytesPerWrite(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerWrite", (value))
}

// GetAvgBytesPerWrite gets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgBytesPerWrite() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerWrite")
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

// SetAvgBytesPerWrite_Base sets the value of AvgBytesPerWrite_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgBytesPerWrite_Base(value uint32) (err error) {
	return instance.SetProperty("AvgBytesPerWrite_Base", (value))
}

// GetAvgBytesPerWrite_Base gets the value of AvgBytesPerWrite_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgBytesPerWrite_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerWrite_Base")
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

// SetAvgDataBytesPerRequest sets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgDataBytesPerRequest(value uint64) (err error) {
	return instance.SetProperty("AvgDataBytesPerRequest", (value))
}

// GetAvgDataBytesPerRequest gets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgDataBytesPerRequest() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDataBytesPerRequest")
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

// SetAvgDataBytesPerRequest_Base sets the value of AvgDataBytesPerRequest_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgDataBytesPerRequest_Base(value uint32) (err error) {
	return instance.SetProperty("AvgDataBytesPerRequest_Base", (value))
}

// GetAvgDataBytesPerRequest_Base gets the value of AvgDataBytesPerRequest_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgDataBytesPerRequest_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgDataBytesPerRequest_Base")
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

// SetAvgDataQueueLength sets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgDataQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgDataQueueLength", (value))
}

// GetAvgDataQueueLength gets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgDataQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDataQueueLength")
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

// SetAvgReadQueueLength sets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgReadQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgReadQueueLength", (value))
}

// GetAvgReadQueueLength gets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgReadQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgReadQueueLength")
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

// SetAvgsecPerDataRequest sets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerDataRequest", (value))
}

// GetAvgsecPerDataRequest gets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerDataRequest")
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

// SetAvgsecPerDataRequest_Base sets the value of AvgsecPerDataRequest_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerDataRequest_Base(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerDataRequest_Base", (value))
}

// GetAvgsecPerDataRequest_Base gets the value of AvgsecPerDataRequest_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerDataRequest_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerDataRequest_Base")
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

// SetAvgsecPerRead sets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerRead(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRead", (value))
}

// GetAvgsecPerRead gets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerRead() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRead")
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

// SetAvgsecPerRead_Base sets the value of AvgsecPerRead_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerRead_Base(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRead_Base", (value))
}

// GetAvgsecPerRead_Base gets the value of AvgsecPerRead_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerRead_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRead_Base")
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

// SetAvgsecPerRequest sets the value of AvgsecPerRequest for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerRequest(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRequest", (value))
}

// GetAvgsecPerRequest gets the value of AvgsecPerRequest for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRequest")
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

// SetAvgsecPerRequest_Base sets the value of AvgsecPerRequest_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerRequest_Base(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRequest_Base", (value))
}

// GetAvgsecPerRequest_Base gets the value of AvgsecPerRequest_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerRequest_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRequest_Base")
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

// SetAvgsecPerWrite sets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerWrite(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerWrite", (value))
}

// GetAvgsecPerWrite gets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerWrite() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerWrite")
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

// SetAvgsecPerWrite_Base sets the value of AvgsecPerWrite_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgsecPerWrite_Base(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerWrite_Base", (value))
}

// GetAvgsecPerWrite_Base gets the value of AvgsecPerWrite_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgsecPerWrite_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerWrite_Base")
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

// SetAvgWriteQueueLength sets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyAvgWriteQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgWriteQueueLength", (value))
}

// GetAvgWriteQueueLength gets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyAvgWriteQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgWriteQueueLength")
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

// SetCurrentDataQueueLength sets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyCurrentDataQueueLength(value uint64) (err error) {
	return instance.SetProperty("CurrentDataQueueLength", (value))
}

// GetCurrentDataQueueLength gets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyCurrentDataQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentDataQueueLength")
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

// SetCurrentDurableOpenFileCount sets the value of CurrentDurableOpenFileCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyCurrentDurableOpenFileCount(value uint64) (err error) {
	return instance.SetProperty("CurrentDurableOpenFileCount", (value))
}

// GetCurrentDurableOpenFileCount gets the value of CurrentDurableOpenFileCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyCurrentDurableOpenFileCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentDurableOpenFileCount")
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

// SetCurrentOpenFileCount sets the value of CurrentOpenFileCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyCurrentOpenFileCount(value uint64) (err error) {
	return instance.SetProperty("CurrentOpenFileCount", (value))
}

// GetCurrentOpenFileCount gets the value of CurrentOpenFileCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyCurrentOpenFileCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentOpenFileCount")
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

// SetCurrentPendingRequests sets the value of CurrentPendingRequests for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyCurrentPendingRequests(value uint64) (err error) {
	return instance.SetProperty("CurrentPendingRequests", (value))
}

// GetCurrentPendingRequests gets the value of CurrentPendingRequests for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyCurrentPendingRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentPendingRequests")
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

// SetDataBytesPersec sets the value of DataBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyDataBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DataBytesPersec", (value))
}

// GetDataBytesPersec gets the value of DataBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyDataBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataBytesPersec")
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

// SetDataRequestsPersec sets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyDataRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("DataRequestsPersec", (value))
}

// GetDataRequestsPersec gets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyDataRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataRequestsPersec")
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

// SetFilesOpenedPersec sets the value of FilesOpenedPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyFilesOpenedPersec(value uint64) (err error) {
	return instance.SetProperty("FilesOpenedPersec", (value))
}

// GetFilesOpenedPersec gets the value of FilesOpenedPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyFilesOpenedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FilesOpenedPersec")
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

// SetMetadataRequestsPersec sets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyMetadataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("MetadataRequestsPersec", (value))
}

// GetMetadataRequestsPersec gets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyMetadataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MetadataRequestsPersec")
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

// SetPercentPersistentHandles sets the value of PercentPersistentHandles for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyPercentPersistentHandles(value uint64) (err error) {
	return instance.SetProperty("PercentPersistentHandles", (value))
}

// GetPercentPersistentHandles gets the value of PercentPersistentHandles for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyPercentPersistentHandles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentPersistentHandles")
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

// SetPercentPersistentHandles_Base sets the value of PercentPersistentHandles_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyPercentPersistentHandles_Base(value uint64) (err error) {
	return instance.SetProperty("PercentPersistentHandles_Base", (value))
}

// GetPercentPersistentHandles_Base gets the value of PercentPersistentHandles_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyPercentPersistentHandles_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentPersistentHandles_Base")
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

// SetPercentResilientHandles sets the value of PercentResilientHandles for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyPercentResilientHandles(value uint64) (err error) {
	return instance.SetProperty("PercentResilientHandles", (value))
}

// GetPercentResilientHandles gets the value of PercentResilientHandles for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyPercentResilientHandles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentResilientHandles")
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

// SetPercentResilientHandles_Base sets the value of PercentResilientHandles_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyPercentResilientHandles_Base(value uint64) (err error) {
	return instance.SetProperty("PercentResilientHandles_Base", (value))
}

// GetPercentResilientHandles_Base gets the value of PercentResilientHandles_Base for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyPercentResilientHandles_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentResilientHandles_Base")
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

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", (value))
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
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

// SetReadRequestsPersec sets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyReadRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequestsPersec", (value))
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyReadRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequestsPersec")
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

// SetReceivedBytesPersec sets the value of ReceivedBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyReceivedBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceivedBytesPersec", (value))
}

// GetReceivedBytesPersec gets the value of ReceivedBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyReceivedBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedBytesPersec")
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

// SetRequestsPersec sets the value of RequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("RequestsPersec", (value))
}

// GetRequestsPersec gets the value of RequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequestsPersec")
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

// SetSentBytesPersec sets the value of SentBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertySentBytesPersec(value uint64) (err error) {
	return instance.SetProperty("SentBytesPersec", (value))
}

// GetSentBytesPersec gets the value of SentBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertySentBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentBytesPersec")
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

// SetTotalDurableHandleReopenCount sets the value of TotalDurableHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalDurableHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalDurableHandleReopenCount", (value))
}

// GetTotalDurableHandleReopenCount gets the value of TotalDurableHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalDurableHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalDurableHandleReopenCount")
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

// SetTotalFailedDurableHandleReopenCount sets the value of TotalFailedDurableHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalFailedDurableHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFailedDurableHandleReopenCount", (value))
}

// GetTotalFailedDurableHandleReopenCount gets the value of TotalFailedDurableHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalFailedDurableHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFailedDurableHandleReopenCount")
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

// SetTotalFailedPersistentHandleReopenCount sets the value of TotalFailedPersistentHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalFailedPersistentHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFailedPersistentHandleReopenCount", (value))
}

// GetTotalFailedPersistentHandleReopenCount gets the value of TotalFailedPersistentHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalFailedPersistentHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFailedPersistentHandleReopenCount")
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

// SetTotalFailedResilientHandleReopenCount sets the value of TotalFailedResilientHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalFailedResilientHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFailedResilientHandleReopenCount", (value))
}

// GetTotalFailedResilientHandleReopenCount gets the value of TotalFailedResilientHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalFailedResilientHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFailedResilientHandleReopenCount")
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

// SetTotalFileOpenCount sets the value of TotalFileOpenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalFileOpenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFileOpenCount", (value))
}

// GetTotalFileOpenCount gets the value of TotalFileOpenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalFileOpenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFileOpenCount")
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

// SetTotalPersistentHandleReopenCount sets the value of TotalPersistentHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalPersistentHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalPersistentHandleReopenCount", (value))
}

// GetTotalPersistentHandleReopenCount gets the value of TotalPersistentHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalPersistentHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalPersistentHandleReopenCount")
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

// SetTotalResilientHandleReopenCount sets the value of TotalResilientHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTotalResilientHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalResilientHandleReopenCount", (value))
}

// GetTotalResilientHandleReopenCount gets the value of TotalResilientHandleReopenCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTotalResilientHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalResilientHandleReopenCount")
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

// SetTransferredBytesPersec sets the value of TransferredBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTransferredBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TransferredBytesPersec", (value))
}

// GetTransferredBytesPersec gets the value of TransferredBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTransferredBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferredBytesPersec")
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

// SetTreeConnectCount sets the value of TreeConnectCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyTreeConnectCount(value uint64) (err error) {
	return instance.SetProperty("TreeConnectCount", (value))
}

// GetTreeConnectCount gets the value of TreeConnectCount for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyTreeConnectCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TreeConnectCount")
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

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", (value))
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytesPersec")
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

// SetWriteRequestsPersec sets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) SetPropertyWriteRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequestsPersec", (value))
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_SMBServerSessions) GetPropertyWriteRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequestsPersec")
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
