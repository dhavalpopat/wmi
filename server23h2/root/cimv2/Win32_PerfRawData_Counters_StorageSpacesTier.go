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

// Win32_PerfRawData_Counters_StorageSpacesTier struct
type Win32_PerfRawData_Counters_StorageSpacesTier struct {
	*Win32_PerfRawData

	//
	TierReadBytesAverage uint64

	//
	TierReadBytesAverage_Base uint32

	//
	TierReadBytesPersec uint64

	//
	TierReadLatency uint32

	//
	TierReadLatency_Base uint32

	//
	TierReadsAverage uint64

	//
	TierReadsPersec uint64

	//
	TierTransferBytesAverage uint64

	//
	TierTransferBytesAverage_Base uint32

	//
	TierTransferBytesPersec uint64

	//
	TierTransferLatency uint32

	//
	TierTransferLatency_Base uint32

	//
	TierTransfersAverage uint64

	//
	TierTransfersCurrent uint32

	//
	TierTransfersPersec uint64

	//
	TierWriteBytesAverage uint64

	//
	TierWriteBytesAverage_Base uint32

	//
	TierWriteBytesPersec uint64

	//
	TierWriteLatency uint32

	//
	TierWriteLatency_Base uint32

	//
	TierWritesAverage uint64

	//
	TierWritesPersec uint64
}

func NewWin32_PerfRawData_Counters_StorageSpacesTierEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_StorageSpacesTier, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesTier{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_StorageSpacesTierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_StorageSpacesTier, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesTier{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetTierReadBytesAverage sets the value of TierReadBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TierReadBytesAverage", (value))
}

// GetTierReadBytesAverage gets the value of TierReadBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadBytesAverage")
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

// SetTierReadBytesAverage_Base sets the value of TierReadBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadBytesAverage_Base(value uint32) (err error) {
	return instance.SetProperty("TierReadBytesAverage_Base", (value))
}

// GetTierReadBytesAverage_Base gets the value of TierReadBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadBytesAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierReadBytesAverage_Base")
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

// SetTierReadBytesPersec sets the value of TierReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TierReadBytesPersec", (value))
}

// GetTierReadBytesPersec gets the value of TierReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadBytesPersec")
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

// SetTierReadLatency sets the value of TierReadLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadLatency(value uint32) (err error) {
	return instance.SetProperty("TierReadLatency", (value))
}

// GetTierReadLatency gets the value of TierReadLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierReadLatency")
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

// SetTierReadLatency_Base sets the value of TierReadLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadLatency_Base(value uint32) (err error) {
	return instance.SetProperty("TierReadLatency_Base", (value))
}

// GetTierReadLatency_Base gets the value of TierReadLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierReadLatency_Base")
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

// SetTierReadsAverage sets the value of TierReadsAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadsAverage(value uint64) (err error) {
	return instance.SetProperty("TierReadsAverage", (value))
}

// GetTierReadsAverage gets the value of TierReadsAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadsAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadsAverage")
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

// SetTierReadsPersec sets the value of TierReadsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierReadsPersec(value uint64) (err error) {
	return instance.SetProperty("TierReadsPersec", (value))
}

// GetTierReadsPersec gets the value of TierReadsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadsPersec")
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

// SetTierTransferBytesAverage sets the value of TierTransferBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransferBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TierTransferBytesAverage", (value))
}

// GetTierTransferBytesAverage gets the value of TierTransferBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransferBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransferBytesAverage")
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

// SetTierTransferBytesAverage_Base sets the value of TierTransferBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransferBytesAverage_Base(value uint32) (err error) {
	return instance.SetProperty("TierTransferBytesAverage_Base", (value))
}

// GetTierTransferBytesAverage_Base gets the value of TierTransferBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransferBytesAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierTransferBytesAverage_Base")
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

// SetTierTransferBytesPersec sets the value of TierTransferBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransferBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TierTransferBytesPersec", (value))
}

// GetTierTransferBytesPersec gets the value of TierTransferBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransferBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransferBytesPersec")
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

// SetTierTransferLatency sets the value of TierTransferLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransferLatency(value uint32) (err error) {
	return instance.SetProperty("TierTransferLatency", (value))
}

// GetTierTransferLatency gets the value of TierTransferLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransferLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierTransferLatency")
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

// SetTierTransferLatency_Base sets the value of TierTransferLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransferLatency_Base(value uint32) (err error) {
	return instance.SetProperty("TierTransferLatency_Base", (value))
}

// GetTierTransferLatency_Base gets the value of TierTransferLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransferLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierTransferLatency_Base")
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

// SetTierTransfersAverage sets the value of TierTransfersAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransfersAverage(value uint64) (err error) {
	return instance.SetProperty("TierTransfersAverage", (value))
}

// GetTierTransfersAverage gets the value of TierTransfersAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransfersAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransfersAverage")
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

// SetTierTransfersCurrent sets the value of TierTransfersCurrent for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransfersCurrent(value uint32) (err error) {
	return instance.SetProperty("TierTransfersCurrent", (value))
}

// GetTierTransfersCurrent gets the value of TierTransfersCurrent for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransfersCurrent() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierTransfersCurrent")
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

// SetTierTransfersPersec sets the value of TierTransfersPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierTransfersPersec(value uint64) (err error) {
	return instance.SetProperty("TierTransfersPersec", (value))
}

// GetTierTransfersPersec gets the value of TierTransfersPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierTransfersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransfersPersec")
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

// SetTierWriteBytesAverage sets the value of TierWriteBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWriteBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TierWriteBytesAverage", (value))
}

// GetTierWriteBytesAverage gets the value of TierWriteBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWriteBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWriteBytesAverage")
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

// SetTierWriteBytesAverage_Base sets the value of TierWriteBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWriteBytesAverage_Base(value uint32) (err error) {
	return instance.SetProperty("TierWriteBytesAverage_Base", (value))
}

// GetTierWriteBytesAverage_Base gets the value of TierWriteBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWriteBytesAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierWriteBytesAverage_Base")
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

// SetTierWriteBytesPersec sets the value of TierWriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TierWriteBytesPersec", (value))
}

// GetTierWriteBytesPersec gets the value of TierWriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWriteBytesPersec")
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

// SetTierWriteLatency sets the value of TierWriteLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWriteLatency(value uint32) (err error) {
	return instance.SetProperty("TierWriteLatency", (value))
}

// GetTierWriteLatency gets the value of TierWriteLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWriteLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierWriteLatency")
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

// SetTierWriteLatency_Base sets the value of TierWriteLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWriteLatency_Base(value uint32) (err error) {
	return instance.SetProperty("TierWriteLatency_Base", (value))
}

// GetTierWriteLatency_Base gets the value of TierWriteLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWriteLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierWriteLatency_Base")
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

// SetTierWritesAverage sets the value of TierWritesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWritesAverage(value uint64) (err error) {
	return instance.SetProperty("TierWritesAverage", (value))
}

// GetTierWritesAverage gets the value of TierWritesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWritesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWritesAverage")
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

// SetTierWritesPersec sets the value of TierWritesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) SetPropertyTierWritesPersec(value uint64) (err error) {
	return instance.SetProperty("TierWritesPersec", (value))
}

// GetTierWritesPersec gets the value of TierWritesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesTier) GetPropertyTierWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWritesPersec")
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
