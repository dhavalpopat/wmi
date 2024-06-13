// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TP_V2_TimerSet struct
type TP_V2_TimerSet struct {
	*ThreadPoolTrace_V2

	//
	Absolute uint32

	//
	DueTime uint64

	//
	Period uint32

	//
	SubQueue uint32

	//
	Timer uint32

	//
	WindowLength uint32
}

func NewTP_V2_TimerSetEx1(instance *cim.WmiInstance) (newInstance *TP_V2_TimerSet, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_TimerSet{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_TimerSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_TimerSet, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_TimerSet{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetAbsolute sets the value of Absolute for the instance
func (instance *TP_V2_TimerSet) SetPropertyAbsolute(value uint32) (err error) {
	return instance.SetProperty("Absolute", (value))
}

// GetAbsolute gets the value of Absolute for the instance
func (instance *TP_V2_TimerSet) GetPropertyAbsolute() (value uint32, err error) {
	retValue, err := instance.GetProperty("Absolute")
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

// SetDueTime sets the value of DueTime for the instance
func (instance *TP_V2_TimerSet) SetPropertyDueTime(value uint64) (err error) {
	return instance.SetProperty("DueTime", (value))
}

// GetDueTime gets the value of DueTime for the instance
func (instance *TP_V2_TimerSet) GetPropertyDueTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("DueTime")
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

// SetPeriod sets the value of Period for the instance
func (instance *TP_V2_TimerSet) SetPropertyPeriod(value uint32) (err error) {
	return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *TP_V2_TimerSet) GetPropertyPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("Period")
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

// SetSubQueue sets the value of SubQueue for the instance
func (instance *TP_V2_TimerSet) SetPropertySubQueue(value uint32) (err error) {
	return instance.SetProperty("SubQueue", (value))
}

// GetSubQueue gets the value of SubQueue for the instance
func (instance *TP_V2_TimerSet) GetPropertySubQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SubQueue")
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

// SetTimer sets the value of Timer for the instance
func (instance *TP_V2_TimerSet) SetPropertyTimer(value uint32) (err error) {
	return instance.SetProperty("Timer", (value))
}

// GetTimer gets the value of Timer for the instance
func (instance *TP_V2_TimerSet) GetPropertyTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timer")
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

// SetWindowLength sets the value of WindowLength for the instance
func (instance *TP_V2_TimerSet) SetPropertyWindowLength(value uint32) (err error) {
	return instance.SetProperty("WindowLength", (value))
}

// GetWindowLength gets the value of WindowLength for the instance
func (instance *TP_V2_TimerSet) GetPropertyWindowLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("WindowLength")
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
