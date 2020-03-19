// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __FilterToConsumerBinding struct
type __FilterToConsumerBinding struct {
	*__IndicationRelated

	//
	Consumer __EventConsumer

	//
	CreatorSID []uint8

	//
	DeliverSynchronously bool

	//
	DeliveryQoS uint32

	//
	Filter __EventFilter

	//
	MaintainSecurityContext bool

	//
	SlowDownProviders bool
}

func New__FilterToConsumerBindingEx1(instance *cim.WmiInstance) (newInstance *__FilterToConsumerBinding, err error) {
	tmp, err := New__IndicationRelatedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__FilterToConsumerBinding{
		__IndicationRelated: tmp,
	}
	return
}

func New__FilterToConsumerBindingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__FilterToConsumerBinding, err error) {
	tmp, err := New__IndicationRelatedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__FilterToConsumerBinding{
		__IndicationRelated: tmp,
	}
	return
}

// SetConsumer sets the value of Consumer for the instance
func (instance *__FilterToConsumerBinding) SetPropertyConsumer(value __EventConsumer) (err error) {
	return instance.SetProperty("Consumer", value)
}

// GetConsumer gets the value of Consumer for the instance
func (instance *__FilterToConsumerBinding) GetPropertyConsumer() (value __EventConsumer, err error) {
	retValue, err := instance.GetProperty("Consumer")
	if err != nil {
		return
	}
	value, ok := retValue.(__EventConsumer)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreatorSID sets the value of CreatorSID for the instance
func (instance *__FilterToConsumerBinding) SetPropertyCreatorSID(value []uint8) (err error) {
	return instance.SetProperty("CreatorSID", value)
}

// GetCreatorSID gets the value of CreatorSID for the instance
func (instance *__FilterToConsumerBinding) GetPropertyCreatorSID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CreatorSID")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeliverSynchronously sets the value of DeliverSynchronously for the instance
func (instance *__FilterToConsumerBinding) SetPropertyDeliverSynchronously(value bool) (err error) {
	return instance.SetProperty("DeliverSynchronously", value)
}

// GetDeliverSynchronously gets the value of DeliverSynchronously for the instance
func (instance *__FilterToConsumerBinding) GetPropertyDeliverSynchronously() (value bool, err error) {
	retValue, err := instance.GetProperty("DeliverSynchronously")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeliveryQoS sets the value of DeliveryQoS for the instance
func (instance *__FilterToConsumerBinding) SetPropertyDeliveryQoS(value uint32) (err error) {
	return instance.SetProperty("DeliveryQoS", value)
}

// GetDeliveryQoS gets the value of DeliveryQoS for the instance
func (instance *__FilterToConsumerBinding) GetPropertyDeliveryQoS() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeliveryQoS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFilter sets the value of Filter for the instance
func (instance *__FilterToConsumerBinding) SetPropertyFilter(value __EventFilter) (err error) {
	return instance.SetProperty("Filter", value)
}

// GetFilter gets the value of Filter for the instance
func (instance *__FilterToConsumerBinding) GetPropertyFilter() (value __EventFilter, err error) {
	retValue, err := instance.GetProperty("Filter")
	if err != nil {
		return
	}
	value, ok := retValue.(__EventFilter)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaintainSecurityContext sets the value of MaintainSecurityContext for the instance
func (instance *__FilterToConsumerBinding) SetPropertyMaintainSecurityContext(value bool) (err error) {
	return instance.SetProperty("MaintainSecurityContext", value)
}

// GetMaintainSecurityContext gets the value of MaintainSecurityContext for the instance
func (instance *__FilterToConsumerBinding) GetPropertyMaintainSecurityContext() (value bool, err error) {
	retValue, err := instance.GetProperty("MaintainSecurityContext")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSlowDownProviders sets the value of SlowDownProviders for the instance
func (instance *__FilterToConsumerBinding) SetPropertySlowDownProviders(value bool) (err error) {
	return instance.SetProperty("SlowDownProviders", value)
}

// GetSlowDownProviders gets the value of SlowDownProviders for the instance
func (instance *__FilterToConsumerBinding) GetPropertySlowDownProviders() (value bool, err error) {
	retValue, err := instance.GetProperty("SlowDownProviders")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
