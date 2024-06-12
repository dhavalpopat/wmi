// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSFC0B587C_1879_46D8_AB82_13F61D22CD8F
//////////////////////////////////////////////
package nsfc0b587c_1879_46d8_ab82_13f61d22cd8f

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __EventDroppedEvent struct
type __EventDroppedEvent struct {
	*__SystemEvent

	//
	Event __Event

	//
	IntendedConsumer __EventConsumer
}

func New__EventDroppedEventEx1(instance *cim.WmiInstance) (newInstance *__EventDroppedEvent, err error) {
	tmp, err := New__SystemEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventDroppedEvent{
		__SystemEvent: tmp,
	}
	return
}

func New__EventDroppedEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventDroppedEvent, err error) {
	tmp, err := New__SystemEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventDroppedEvent{
		__SystemEvent: tmp,
	}
	return
}

// SetEvent sets the value of Event for the instance
func (instance *__EventDroppedEvent) SetPropertyEvent(value __Event) (err error) {
	return instance.SetProperty("Event", (value))
}

// GetEvent gets the value of Event for the instance
func (instance *__EventDroppedEvent) GetPropertyEvent() (value __Event, err error) {
	retValue, err := instance.GetProperty("Event")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(__Event)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " __Event is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = __Event(valuetmp)

	return
}

// SetIntendedConsumer sets the value of IntendedConsumer for the instance
func (instance *__EventDroppedEvent) SetPropertyIntendedConsumer(value __EventConsumer) (err error) {
	return instance.SetProperty("IntendedConsumer", (value))
}

// GetIntendedConsumer gets the value of IntendedConsumer for the instance
func (instance *__EventDroppedEvent) GetPropertyIntendedConsumer() (value __EventConsumer, err error) {
	retValue, err := instance.GetProperty("IntendedConsumer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(__EventConsumer)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " __EventConsumer is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = __EventConsumer(valuetmp)

	return
}
