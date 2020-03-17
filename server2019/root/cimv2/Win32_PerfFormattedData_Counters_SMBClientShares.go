// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_SMBClientShares struct
type Win32_PerfFormattedData_Counters_SMBClientShares struct {
	Win32_PerfFormattedData

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
	CompressedRequestsPersec uint32

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

// SetAvgBytesPerRead sets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgBytesPerRead(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerRead", value)
}

// GetAvgBytesPerRead gets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgBytesPerRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgBytesPerWrite sets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgBytesPerWrite(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerWrite", value)
}

// GetAvgBytesPerWrite gets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgBytesPerWrite() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerWrite")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDataBytesPerRequest sets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgDataBytesPerRequest(value uint64) (err error) {
	return instance.SetProperty("AvgDataBytesPerRequest", value)
}

// GetAvgDataBytesPerRequest gets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgDataBytesPerRequest() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDataBytesPerRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDataQueueLength sets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgDataQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgDataQueueLength", value)
}

// GetAvgDataQueueLength gets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgDataQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDataQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgReadQueueLength sets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgReadQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgReadQueueLength", value)
}

// GetAvgReadQueueLength gets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgReadQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgReadQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerDataRequest sets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerDataRequest", value)
}

// GetAvgsecPerDataRequest gets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerDataRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerRead sets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgsecPerRead(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRead", value)
}

// GetAvgsecPerRead gets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgsecPerRead() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerWrite sets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgsecPerWrite(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerWrite", value)
}

// GetAvgsecPerWrite gets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgsecPerWrite() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerWrite")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgWriteQueueLength sets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyAvgWriteQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgWriteQueueLength", value)
}

// GetAvgWriteQueueLength gets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyAvgWriteQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgWriteQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressedBytesSentPersec sets the value of CompressedBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCompressedBytesSentPersec(value uint32) (err error) {
	return instance.SetProperty("CompressedBytesSentPersec", value)
}

// GetCompressedBytesSentPersec gets the value of CompressedBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCompressedBytesSentPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompressedBytesSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressedRequestsPersec sets the value of CompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCompressedRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("CompressedRequestsPersec", value)
}

// GetCompressedRequestsPersec gets the value of CompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCompressedRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompressedRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressedResponsesPersec sets the value of CompressedResponsesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCompressedResponsesPersec(value uint32) (err error) {
	return instance.SetProperty("CompressedResponsesPersec", value)
}

// GetCompressedResponsesPersec gets the value of CompressedResponsesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCompressedResponsesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompressedResponsesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreditStallsPersec sets the value of CreditStallsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCreditStallsPersec(value uint32) (err error) {
	return instance.SetProperty("CreditStallsPersec", value)
}

// GetCreditStallsPersec gets the value of CreditStallsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCreditStallsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CreditStallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentDataQueueLength sets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyCurrentDataQueueLength(value uint32) (err error) {
	return instance.SetProperty("CurrentDataQueueLength", value)
}

// GetCurrentDataQueueLength gets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyCurrentDataQueueLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentDataQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataBytesPersec sets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyDataBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DataBytesPersec", value)
}

// GetDataBytesPersec gets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyDataBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataRequestsPersec sets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyDataRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("DataRequestsPersec", value)
}

// GetDataRequestsPersec gets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyDataRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetadataRequestsPersec sets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyMetadataRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("MetadataRequestsPersec", value)
}

// GetMetadataRequestsPersec gets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyMetadataRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", value)
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBytestransmittedviaSMBDirectPersec sets the value of ReadBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadBytestransmittedviaSMBDirectPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytestransmittedviaSMBDirectPersec", value)
}

// GetReadBytestransmittedviaSMBDirectPersec gets the value of ReadBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadBytestransmittedviaSMBDirectPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytestransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadRequestsPersec sets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequestsPersec", value)
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadRequeststransmittedviaSMBDirectPersec sets the value of ReadRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyReadRequeststransmittedviaSMBDirectPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequeststransmittedviaSMBDirectPersec", value)
}

// GetReadRequeststransmittedviaSMBDirectPersec gets the value of ReadRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyReadRequeststransmittedviaSMBDirectPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequeststransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurboIOReadsPersec sets the value of TurboIOReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyTurboIOReadsPersec(value uint32) (err error) {
	return instance.SetProperty("TurboIOReadsPersec", value)
}

// GetTurboIOReadsPersec gets the value of TurboIOReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyTurboIOReadsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TurboIOReadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurboIOWritesPersec sets the value of TurboIOWritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyTurboIOWritesPersec(value uint32) (err error) {
	return instance.SetProperty("TurboIOWritesPersec", value)
}

// GetTurboIOWritesPersec gets the value of TurboIOWritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyTurboIOWritesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TurboIOWritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", value)
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytestransmittedviaSMBDirectPersec sets the value of WriteBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteBytestransmittedviaSMBDirectPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytestransmittedviaSMBDirectPersec", value)
}

// GetWriteBytestransmittedviaSMBDirectPersec gets the value of WriteBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteBytestransmittedviaSMBDirectPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytestransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteRequestsPersec sets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequestsPersec", value)
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteRequeststransmittedviaSMBDirectPersec sets the value of WriteRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) SetPropertyWriteRequeststransmittedviaSMBDirectPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequeststransmittedviaSMBDirectPersec", value)
}

// GetWriteRequeststransmittedviaSMBDirectPersec gets the value of WriteRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBClientShares) GetPropertyWriteRequeststransmittedviaSMBDirectPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequeststransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
