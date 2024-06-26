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

// Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages struct
type Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages struct {
	*Win32_PerfRawData

	//
	MessagesOutstanding uint64

	//
	MessagesSent uint64

	//
	MessagesSentPersec uint64
}

func NewWin32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessagesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessagesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetMessagesOutstanding sets the value of MessagesOutstanding for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) SetPropertyMessagesOutstanding(value uint64) (err error) {
	return instance.SetProperty("MessagesOutstanding", (value))
}

// GetMessagesOutstanding gets the value of MessagesOutstanding for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) GetPropertyMessagesOutstanding() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessagesOutstanding")
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

// SetMessagesSent sets the value of MessagesSent for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) SetPropertyMessagesSent(value uint64) (err error) {
	return instance.SetProperty("MessagesSent", (value))
}

// GetMessagesSent gets the value of MessagesSent for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) GetPropertyMessagesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessagesSent")
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

// SetMessagesSentPersec sets the value of MessagesSentPersec for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) SetPropertyMessagesSentPersec(value uint64) (err error) {
	return instance.SetProperty("MessagesSentPersec", (value))
}

// GetMessagesSentPersec gets the value of MessagesSentPersec for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterMulticastRequestResponseMessages) GetPropertyMessagesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessagesSentPersec")
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
