// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS7344EDEA_49A9_449C_A6C2_37ECE9C9272F
//////////////////////////////////////////////
package ns7344edea_49a9_449c_a6c2_37ece9c9272f

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __IntervalTimerInstruction struct
type __IntervalTimerInstruction struct {
	*__TimerInstruction

	//
	IntervalBetweenEvents uint32
}

func New__IntervalTimerInstructionEx1(instance *cim.WmiInstance) (newInstance *__IntervalTimerInstruction, err error) {
	tmp, err := New__TimerInstructionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__IntervalTimerInstruction{
		__TimerInstruction: tmp,
	}
	return
}

func New__IntervalTimerInstructionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__IntervalTimerInstruction, err error) {
	tmp, err := New__TimerInstructionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__IntervalTimerInstruction{
		__TimerInstruction: tmp,
	}
	return
}

// SetIntervalBetweenEvents sets the value of IntervalBetweenEvents for the instance
func (instance *__IntervalTimerInstruction) SetPropertyIntervalBetweenEvents(value uint32) (err error) {
	return instance.SetProperty("IntervalBetweenEvents", (value))
}

// GetIntervalBetweenEvents gets the value of IntervalBetweenEvents for the instance
func (instance *__IntervalTimerInstruction) GetPropertyIntervalBetweenEvents() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntervalBetweenEvents")
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
