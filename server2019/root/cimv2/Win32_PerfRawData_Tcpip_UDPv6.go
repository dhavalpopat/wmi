// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Tcpip_UDPv6 struct
type Win32_PerfRawData_Tcpip_UDPv6 struct {
	Win32_PerfRawData

	//
	DatagramsNoPortPersec uint32

	//
	DatagramsPersec uint32

	//
	DatagramsReceivedErrors uint32

	//
	DatagramsReceivedPersec uint32

	//
	DatagramsSentPersec uint32
}

// SetDatagramsNoPortPersec sets the value of DatagramsNoPortPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) SetPropertyDatagramsNoPortPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsNoPortPersec", value)
}

// GetDatagramsNoPortPersec gets the value of DatagramsNoPortPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) GetPropertyDatagramsNoPortPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsNoPortPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsPersec sets the value of DatagramsPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) SetPropertyDatagramsPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsPersec", value)
}

// GetDatagramsPersec gets the value of DatagramsPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) GetPropertyDatagramsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedErrors sets the value of DatagramsReceivedErrors for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) SetPropertyDatagramsReceivedErrors(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedErrors", value)
}

// GetDatagramsReceivedErrors gets the value of DatagramsReceivedErrors for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) GetPropertyDatagramsReceivedErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedPersec sets the value of DatagramsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) SetPropertyDatagramsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedPersec", value)
}

// GetDatagramsReceivedPersec gets the value of DatagramsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) GetPropertyDatagramsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsSentPersec sets the value of DatagramsSentPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) SetPropertyDatagramsSentPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsSentPersec", value)
}

// GetDatagramsSentPersec gets the value of DatagramsSentPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_UDPv6) GetPropertyDatagramsSentPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
