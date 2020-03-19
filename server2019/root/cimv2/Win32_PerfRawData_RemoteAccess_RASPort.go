// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfRawData_RemoteAccess_RASPort struct
type Win32_PerfRawData_RemoteAccess_RASPort struct {
	*Win32_PerfRawData

	//
	AlignmentErrors uint32

	//
	BufferOverrunErrors uint32

	//
	BytesReceived uint64

	//
	BytesReceivedPerSec uint32

	//
	BytesTransmitted uint64

	//
	BytesTransmittedPerSec uint32

	//
	CRCErrors uint32

	//
	FramesReceived uint32

	//
	FramesReceivedPerSec uint32

	//
	FramesTransmitted uint32

	//
	FramesTransmittedPerSec uint32

	//
	PercentCompressionIn uint32

	//
	PercentCompressionOut uint32

	//
	SerialOverrunErrors uint32

	//
	TimeoutErrors uint32

	//
	TotalErrors uint32

	//
	TotalErrorsPerSec uint32
}

func NewWin32_PerfRawData_RemoteAccess_RASPortEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_RemoteAccess_RASPort, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_RemoteAccess_RASPort{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_RemoteAccess_RASPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_RemoteAccess_RASPort, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_RemoteAccess_RASPort{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAlignmentErrors sets the value of AlignmentErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyAlignmentErrors(value uint32) (err error) {
	return instance.SetProperty("AlignmentErrors", value)
}

// GetAlignmentErrors gets the value of AlignmentErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyAlignmentErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("AlignmentErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBufferOverrunErrors sets the value of BufferOverrunErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyBufferOverrunErrors(value uint32) (err error) {
	return instance.SetProperty("BufferOverrunErrors", value)
}

// GetBufferOverrunErrors gets the value of BufferOverrunErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyBufferOverrunErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferOverrunErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesReceivedPerSec sets the value of BytesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyBytesReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesReceivedPerSec", value)
}

// GetBytesReceivedPerSec gets the value of BytesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyBytesReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesReceivedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTransmitted sets the value of BytesTransmitted for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyBytesTransmitted(value uint64) (err error) {
	return instance.SetProperty("BytesTransmitted", value)
}

// GetBytesTransmitted gets the value of BytesTransmitted for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyBytesTransmitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTransmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTransmittedPerSec sets the value of BytesTransmittedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyBytesTransmittedPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesTransmittedPerSec", value)
}

// GetBytesTransmittedPerSec gets the value of BytesTransmittedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyBytesTransmittedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesTransmittedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCRCErrors sets the value of CRCErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyCRCErrors(value uint32) (err error) {
	return instance.SetProperty("CRCErrors", value)
}

// GetCRCErrors gets the value of CRCErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyCRCErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("CRCErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesReceived sets the value of FramesReceived for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyFramesReceived(value uint32) (err error) {
	return instance.SetProperty("FramesReceived", value)
}

// GetFramesReceived gets the value of FramesReceived for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyFramesReceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("FramesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesReceivedPerSec sets the value of FramesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyFramesReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("FramesReceivedPerSec", value)
}

// GetFramesReceivedPerSec gets the value of FramesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyFramesReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FramesReceivedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesTransmitted sets the value of FramesTransmitted for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyFramesTransmitted(value uint32) (err error) {
	return instance.SetProperty("FramesTransmitted", value)
}

// GetFramesTransmitted gets the value of FramesTransmitted for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyFramesTransmitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("FramesTransmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesTransmittedPerSec sets the value of FramesTransmittedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyFramesTransmittedPerSec(value uint32) (err error) {
	return instance.SetProperty("FramesTransmittedPerSec", value)
}

// GetFramesTransmittedPerSec gets the value of FramesTransmittedPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyFramesTransmittedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FramesTransmittedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentCompressionIn sets the value of PercentCompressionIn for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyPercentCompressionIn(value uint32) (err error) {
	return instance.SetProperty("PercentCompressionIn", value)
}

// GetPercentCompressionIn gets the value of PercentCompressionIn for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyPercentCompressionIn() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentCompressionIn")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentCompressionOut sets the value of PercentCompressionOut for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyPercentCompressionOut(value uint32) (err error) {
	return instance.SetProperty("PercentCompressionOut", value)
}

// GetPercentCompressionOut gets the value of PercentCompressionOut for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyPercentCompressionOut() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentCompressionOut")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialOverrunErrors sets the value of SerialOverrunErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertySerialOverrunErrors(value uint32) (err error) {
	return instance.SetProperty("SerialOverrunErrors", value)
}

// GetSerialOverrunErrors gets the value of SerialOverrunErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertySerialOverrunErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("SerialOverrunErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeoutErrors sets the value of TimeoutErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyTimeoutErrors(value uint32) (err error) {
	return instance.SetProperty("TimeoutErrors", value)
}

// GetTimeoutErrors gets the value of TimeoutErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyTimeoutErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeoutErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalErrors sets the value of TotalErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyTotalErrors(value uint32) (err error) {
	return instance.SetProperty("TotalErrors", value)
}

// GetTotalErrors gets the value of TotalErrors for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyTotalErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalErrorsPerSec sets the value of TotalErrorsPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) SetPropertyTotalErrorsPerSec(value uint32) (err error) {
	return instance.SetProperty("TotalErrorsPerSec", value)
}

// GetTotalErrorsPerSec gets the value of TotalErrorsPerSec for the instance
func (instance *Win32_PerfRawData_RemoteAccess_RASPort) GetPropertyTotalErrorsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalErrorsPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
