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

// Win32_PerfFormattedData_Counters_SMBClientShares struct
type Win32_PerfFormattedData_Counters_SMBClientShares struct {
	*Win32_PerfFormattedData

	//
	AttemptedCompressedRequestsPersec uint32

	//
	AvgBytesPerRead uint64

	//
	AvgBytesPerWrite uint64

	//
	AvgDataBytesPerRequest uint64

	//
	AvgDataQueueLength uint64

	//
	AvgReadQueueLength uint64

	//
	AvgsecPerDataRequest uint32

	//
	AvgsecPerRead uint32

	//
	AvgsecPerWrite uint32

	//
	AvgWriteQueueLength uint64

	//
	CompressedBytesSentPersec uint32

	//
	CompressedResponsesPersec uint32

	//
	CreditStallsPersec uint32

	//
	CurrentDataQueueLength uint32

	//
	DataBytesPersec uint64

	//
	DataRequestsPersec uint32

	//
	MetadataRequestsPersec uint32

	//
	ReadBytesPersec uint64

	//
	ReadBytestransmittedviaSMBDirectPersec uint64

	//
	ReadRequestsPersec uint32

	//
	ReadRequeststransmittedviaSMBDirectPersec uint32

	//
	SuccessfulCompressedRequestsPersec uint32

	//
	TurboIOReadsPersec uint32

	//
	TurboIOWritesPersec uint32

	//
	WriteBytesPersec uint64

	//
	WriteBytestransmittedviaSMBDirectPersec uint64

	//
	WriteRequestsPersec uint32

	//
	WriteRequeststransmittedviaSMBDirectPersec uint32
}

func NewWin32_PerfFormattedData_Counters_SMBClientSharesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_SMBClientShares, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_SMBClientShares{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_SMBClientSharesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_SMBClientShares, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_SMBClientShares{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAttemptedCompressedRequestsPersec sets the value of AttemptedCompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAttemptedCompressedRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("AttemptedCompressedRequestsPersec", (value))
}

// GetAttemptedCompressedRequestsPersec gets the value of AttemptedCompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAttemptedCompressedRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("AttemptedCompressedRequestsPersec")
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

// SetAvgBytesPerRead sets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgBytesPerRead(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerRead", (value))
}

// GetAvgBytesPerRead gets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgBytesPerRead() (value uint64, err error) {
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

// SetAvgBytesPerWrite sets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgBytesPerWrite(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerWrite", (value))
}

// GetAvgBytesPerWrite gets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgBytesPerWrite() (value uint64, err error) {
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

// SetAvgDataBytesPerRequest sets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgDataBytesPerRequest(value uint64) (err error) {
	return instance.SetProperty("AvgDataBytesPerRequest", (value))
}

// GetAvgDataBytesPerRequest gets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgDataBytesPerRequest() (value uint64, err error) {
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

// SetAvgDataQueueLength sets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgDataQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgDataQueueLength", (value))
}

// GetAvgDataQueueLength gets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgDataQueueLength() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgReadQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgReadQueueLength", (value))
}

// GetAvgReadQueueLength gets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgReadQueueLength() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerDataRequest", (value))
}

// GetAvgsecPerDataRequest gets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgsecPerDataRequest() (value uint32, err error) {
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

// SetAvgsecPerRead sets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgsecPerRead(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRead", (value))
}

// GetAvgsecPerRead gets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgsecPerRead() (value uint32, err error) {
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

// SetAvgsecPerWrite sets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgsecPerWrite(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerWrite", (value))
}

// GetAvgsecPerWrite gets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgsecPerWrite() (value uint32, err error) {
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

// SetAvgWriteQueueLength sets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgWriteQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgWriteQueueLength", (value))
}

// GetAvgWriteQueueLength gets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgWriteQueueLength() (value uint64, err error) {
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

// SetCompressedBytesSentPersec sets the value of CompressedBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCompressedBytesSentPersec(value uint32) (err error) {
	return instance.SetProperty("CompressedBytesSentPersec", (value))
}

// GetCompressedBytesSentPersec gets the value of CompressedBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCompressedBytesSentPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompressedBytesSentPersec")
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

// SetCompressedResponsesPersec sets the value of CompressedResponsesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCompressedResponsesPersec(value uint32) (err error) {
	return instance.SetProperty("CompressedResponsesPersec", (value))
}

// GetCompressedResponsesPersec gets the value of CompressedResponsesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCompressedResponsesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompressedResponsesPersec")
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

// SetCreditStallsPersec sets the value of CreditStallsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCreditStallsPersec(value uint32) (err error) {
	return instance.SetProperty("CreditStallsPersec", (value))
}

// GetCreditStallsPersec gets the value of CreditStallsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCreditStallsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CreditStallsPersec")
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

// SetCurrentDataQueueLength sets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCurrentDataQueueLength(value uint32) (err error) {
	return instance.SetProperty("CurrentDataQueueLength", (value))
}

// GetCurrentDataQueueLength gets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCurrentDataQueueLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentDataQueueLength")
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

// SetDataBytesPersec sets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyDataBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DataBytesPersec", (value))
}

// GetDataBytesPersec gets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyDataBytesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyDataRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("DataRequestsPersec", (value))
}

// GetDataRequestsPersec gets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyDataRequestsPersec() (value uint32, err error) {
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

// SetMetadataRequestsPersec sets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyMetadataRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("MetadataRequestsPersec", (value))
}

// GetMetadataRequestsPersec gets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyMetadataRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataRequestsPersec")
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

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", (value))
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadBytesPersec() (value uint64, err error) {
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

// SetReadBytestransmittedviaSMBDirectPersec sets the value of ReadBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadBytestransmittedviaSMBDirectPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytestransmittedviaSMBDirectPersec", (value))
}

// GetReadBytestransmittedviaSMBDirectPersec gets the value of ReadBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadBytestransmittedviaSMBDirectPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytestransmittedviaSMBDirectPersec")
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
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequestsPersec", (value))
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadRequestsPersec() (value uint32, err error) {
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

// SetReadRequeststransmittedviaSMBDirectPersec sets the value of ReadRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadRequeststransmittedviaSMBDirectPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequeststransmittedviaSMBDirectPersec", (value))
}

// GetReadRequeststransmittedviaSMBDirectPersec gets the value of ReadRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadRequeststransmittedviaSMBDirectPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequeststransmittedviaSMBDirectPersec")
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

// SetSuccessfulCompressedRequestsPersec sets the value of SuccessfulCompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertySuccessfulCompressedRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("SuccessfulCompressedRequestsPersec", (value))
}

// GetSuccessfulCompressedRequestsPersec gets the value of SuccessfulCompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertySuccessfulCompressedRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulCompressedRequestsPersec")
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

// SetTurboIOReadsPersec sets the value of TurboIOReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyTurboIOReadsPersec(value uint32) (err error) {
	return instance.SetProperty("TurboIOReadsPersec", (value))
}

// GetTurboIOReadsPersec gets the value of TurboIOReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyTurboIOReadsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TurboIOReadsPersec")
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

// SetTurboIOWritesPersec sets the value of TurboIOWritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyTurboIOWritesPersec(value uint32) (err error) {
	return instance.SetProperty("TurboIOWritesPersec", (value))
}

// GetTurboIOWritesPersec gets the value of TurboIOWritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyTurboIOWritesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TurboIOWritesPersec")
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

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", (value))
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteBytesPersec() (value uint64, err error) {
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

// SetWriteBytestransmittedviaSMBDirectPersec sets the value of WriteBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteBytestransmittedviaSMBDirectPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytestransmittedviaSMBDirectPersec", (value))
}

// GetWriteBytestransmittedviaSMBDirectPersec gets the value of WriteBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteBytestransmittedviaSMBDirectPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytestransmittedviaSMBDirectPersec")
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
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequestsPersec", (value))
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteRequestsPersec() (value uint32, err error) {
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

// SetWriteRequeststransmittedviaSMBDirectPersec sets the value of WriteRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteRequeststransmittedviaSMBDirectPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequeststransmittedviaSMBDirectPersec", (value))
}

// GetWriteRequeststransmittedviaSMBDirectPersec gets the value of WriteRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteRequeststransmittedviaSMBDirectPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequeststransmittedviaSMBDirectPersec")
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
