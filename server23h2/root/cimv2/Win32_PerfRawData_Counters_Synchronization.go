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

// Win32_PerfRawData_Counters_Synchronization struct
type Win32_PerfRawData_Counters_Synchronization struct {
	*Win32_PerfRawData

	//
	ExecResourceAcquiresAcqExclLitePersec uint32

	//
	ExecResourceAcquiresAcqShrdLitePersec uint32

	//
	ExecResourceAcquiresAcqShrdStarveExclPersec uint32

	//
	ExecResourceAcquiresAcqShrdWaitForExclPersec uint32

	//
	ExecResourceAttemptsAcqExclLitePersec uint32

	//
	ExecResourceAttemptsAcqShrdLitePersec uint32

	//
	ExecResourceAttemptsAcqShrdStarveExclPersec uint32

	//
	ExecResourceAttemptsAcqShrdWaitForExclPersec uint32

	//
	ExecResourceBoostExclOwnerPersec uint32

	//
	ExecResourceBoostSharedOwnersPersec uint32

	//
	ExecResourceContentionAcqExclLitePersec uint32

	//
	ExecResourceContentionAcqShrdLitePersec uint32

	//
	ExecResourceContentionAcqShrdStarveExclPersec uint32

	//
	ExecResourceContentionAcqShrdWaitForExclPersec uint32

	//
	ExecResourcenoWaitsAcqExclLitePersec uint32

	//
	ExecResourcenoWaitsAcqShrdLitePersec uint32

	//
	ExecResourcenoWaitsAcqShrdStarveExclPersec uint32

	//
	ExecResourcenoWaitsAcqShrdWaitForExclPersec uint32

	//
	ExecResourceRecursiveExclAcquiresAcqExclLitePersec uint32

	//
	ExecResourceRecursiveExclAcquiresAcqShrdLitePersec uint32

	//
	ExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec uint32

	//
	ExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec uint32

	//
	ExecResourceRecursiveShAcquiresAcqShrdLitePersec uint32

	//
	ExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec uint32

	//
	ExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec uint32

	//
	ExecResourceSetOwnerPointerExclusivePersec uint32

	//
	ExecResourceSetOwnerPointerSharedExistingOwnerPersec uint32

	//
	ExecResourceSetOwnerPointerSharedNewOwnerPersec uint32

	//
	ExecResourceTotalAcquiresPersec uint32

	//
	ExecResourceTotalContentionsPersec uint32

	//
	ExecResourceTotalConvExclusiveToSharedPersec uint32

	//
	ExecResourceTotalDeletePersec uint32

	//
	ExecResourceTotalExclusiveReleasesPersec uint32

	//
	ExecResourceTotalInitializePersec uint32

	//
	ExecResourceTotalReInitializePersec uint32

	//
	ExecResourceTotalSharedReleasesPersec uint32

	//
	IPISendBroadcastRequestsPersec uint32

	//
	IPISendRoutineRequestsPersec uint32

	//
	IPISendSoftwareInterruptsPersec uint32

	//
	SpinlockAcquiresPersec uint32

	//
	SpinlockContentionsPersec uint32

	//
	SpinlockSpinsPersec uint32
}

func NewWin32_PerfRawData_Counters_SynchronizationEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_Synchronization, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_Synchronization{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_SynchronizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_Synchronization, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_Synchronization{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetExecResourceAcquiresAcqExclLitePersec sets the value of ExecResourceAcquiresAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAcquiresAcqExclLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAcquiresAcqExclLitePersec", (value))
}

// GetExecResourceAcquiresAcqExclLitePersec gets the value of ExecResourceAcquiresAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAcquiresAcqExclLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAcquiresAcqExclLitePersec")
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

// SetExecResourceAcquiresAcqShrdLitePersec sets the value of ExecResourceAcquiresAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAcquiresAcqShrdLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAcquiresAcqShrdLitePersec", (value))
}

// GetExecResourceAcquiresAcqShrdLitePersec gets the value of ExecResourceAcquiresAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAcquiresAcqShrdLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAcquiresAcqShrdLitePersec")
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

// SetExecResourceAcquiresAcqShrdStarveExclPersec sets the value of ExecResourceAcquiresAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAcquiresAcqShrdStarveExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAcquiresAcqShrdStarveExclPersec", (value))
}

// GetExecResourceAcquiresAcqShrdStarveExclPersec gets the value of ExecResourceAcquiresAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAcquiresAcqShrdStarveExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAcquiresAcqShrdStarveExclPersec")
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

// SetExecResourceAcquiresAcqShrdWaitForExclPersec sets the value of ExecResourceAcquiresAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAcquiresAcqShrdWaitForExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAcquiresAcqShrdWaitForExclPersec", (value))
}

// GetExecResourceAcquiresAcqShrdWaitForExclPersec gets the value of ExecResourceAcquiresAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAcquiresAcqShrdWaitForExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAcquiresAcqShrdWaitForExclPersec")
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

// SetExecResourceAttemptsAcqExclLitePersec sets the value of ExecResourceAttemptsAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAttemptsAcqExclLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAttemptsAcqExclLitePersec", (value))
}

// GetExecResourceAttemptsAcqExclLitePersec gets the value of ExecResourceAttemptsAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAttemptsAcqExclLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAttemptsAcqExclLitePersec")
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

// SetExecResourceAttemptsAcqShrdLitePersec sets the value of ExecResourceAttemptsAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAttemptsAcqShrdLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAttemptsAcqShrdLitePersec", (value))
}

// GetExecResourceAttemptsAcqShrdLitePersec gets the value of ExecResourceAttemptsAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAttemptsAcqShrdLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAttemptsAcqShrdLitePersec")
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

// SetExecResourceAttemptsAcqShrdStarveExclPersec sets the value of ExecResourceAttemptsAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAttemptsAcqShrdStarveExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAttemptsAcqShrdStarveExclPersec", (value))
}

// GetExecResourceAttemptsAcqShrdStarveExclPersec gets the value of ExecResourceAttemptsAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAttemptsAcqShrdStarveExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAttemptsAcqShrdStarveExclPersec")
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

// SetExecResourceAttemptsAcqShrdWaitForExclPersec sets the value of ExecResourceAttemptsAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceAttemptsAcqShrdWaitForExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceAttemptsAcqShrdWaitForExclPersec", (value))
}

// GetExecResourceAttemptsAcqShrdWaitForExclPersec gets the value of ExecResourceAttemptsAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceAttemptsAcqShrdWaitForExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceAttemptsAcqShrdWaitForExclPersec")
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

// SetExecResourceBoostExclOwnerPersec sets the value of ExecResourceBoostExclOwnerPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceBoostExclOwnerPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceBoostExclOwnerPersec", (value))
}

// GetExecResourceBoostExclOwnerPersec gets the value of ExecResourceBoostExclOwnerPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceBoostExclOwnerPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceBoostExclOwnerPersec")
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

// SetExecResourceBoostSharedOwnersPersec sets the value of ExecResourceBoostSharedOwnersPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceBoostSharedOwnersPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceBoostSharedOwnersPersec", (value))
}

// GetExecResourceBoostSharedOwnersPersec gets the value of ExecResourceBoostSharedOwnersPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceBoostSharedOwnersPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceBoostSharedOwnersPersec")
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

// SetExecResourceContentionAcqExclLitePersec sets the value of ExecResourceContentionAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceContentionAcqExclLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceContentionAcqExclLitePersec", (value))
}

// GetExecResourceContentionAcqExclLitePersec gets the value of ExecResourceContentionAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceContentionAcqExclLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceContentionAcqExclLitePersec")
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

// SetExecResourceContentionAcqShrdLitePersec sets the value of ExecResourceContentionAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceContentionAcqShrdLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceContentionAcqShrdLitePersec", (value))
}

// GetExecResourceContentionAcqShrdLitePersec gets the value of ExecResourceContentionAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceContentionAcqShrdLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceContentionAcqShrdLitePersec")
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

// SetExecResourceContentionAcqShrdStarveExclPersec sets the value of ExecResourceContentionAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceContentionAcqShrdStarveExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceContentionAcqShrdStarveExclPersec", (value))
}

// GetExecResourceContentionAcqShrdStarveExclPersec gets the value of ExecResourceContentionAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceContentionAcqShrdStarveExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceContentionAcqShrdStarveExclPersec")
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

// SetExecResourceContentionAcqShrdWaitForExclPersec sets the value of ExecResourceContentionAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceContentionAcqShrdWaitForExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceContentionAcqShrdWaitForExclPersec", (value))
}

// GetExecResourceContentionAcqShrdWaitForExclPersec gets the value of ExecResourceContentionAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceContentionAcqShrdWaitForExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceContentionAcqShrdWaitForExclPersec")
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

// SetExecResourcenoWaitsAcqExclLitePersec sets the value of ExecResourcenoWaitsAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourcenoWaitsAcqExclLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourcenoWaitsAcqExclLitePersec", (value))
}

// GetExecResourcenoWaitsAcqExclLitePersec gets the value of ExecResourcenoWaitsAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourcenoWaitsAcqExclLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourcenoWaitsAcqExclLitePersec")
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

// SetExecResourcenoWaitsAcqShrdLitePersec sets the value of ExecResourcenoWaitsAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourcenoWaitsAcqShrdLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourcenoWaitsAcqShrdLitePersec", (value))
}

// GetExecResourcenoWaitsAcqShrdLitePersec gets the value of ExecResourcenoWaitsAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourcenoWaitsAcqShrdLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourcenoWaitsAcqShrdLitePersec")
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

// SetExecResourcenoWaitsAcqShrdStarveExclPersec sets the value of ExecResourcenoWaitsAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourcenoWaitsAcqShrdStarveExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourcenoWaitsAcqShrdStarveExclPersec", (value))
}

// GetExecResourcenoWaitsAcqShrdStarveExclPersec gets the value of ExecResourcenoWaitsAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourcenoWaitsAcqShrdStarveExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourcenoWaitsAcqShrdStarveExclPersec")
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

// SetExecResourcenoWaitsAcqShrdWaitForExclPersec sets the value of ExecResourcenoWaitsAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourcenoWaitsAcqShrdWaitForExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourcenoWaitsAcqShrdWaitForExclPersec", (value))
}

// GetExecResourcenoWaitsAcqShrdWaitForExclPersec gets the value of ExecResourcenoWaitsAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourcenoWaitsAcqShrdWaitForExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourcenoWaitsAcqShrdWaitForExclPersec")
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

// SetExecResourceRecursiveExclAcquiresAcqExclLitePersec sets the value of ExecResourceRecursiveExclAcquiresAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveExclAcquiresAcqExclLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveExclAcquiresAcqExclLitePersec", (value))
}

// GetExecResourceRecursiveExclAcquiresAcqExclLitePersec gets the value of ExecResourceRecursiveExclAcquiresAcqExclLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveExclAcquiresAcqExclLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveExclAcquiresAcqExclLitePersec")
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

// SetExecResourceRecursiveExclAcquiresAcqShrdLitePersec sets the value of ExecResourceRecursiveExclAcquiresAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveExclAcquiresAcqShrdLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveExclAcquiresAcqShrdLitePersec", (value))
}

// GetExecResourceRecursiveExclAcquiresAcqShrdLitePersec gets the value of ExecResourceRecursiveExclAcquiresAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveExclAcquiresAcqShrdLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveExclAcquiresAcqShrdLitePersec")
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

// SetExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec sets the value of ExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec", (value))
}

// GetExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec gets the value of ExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveExclAcquiresAcqShrdStarveExclPersec")
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

// SetExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec sets the value of ExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec", (value))
}

// GetExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec gets the value of ExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveExclAcquiresAcqShrdWaitForExclPersec")
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

// SetExecResourceRecursiveShAcquiresAcqShrdLitePersec sets the value of ExecResourceRecursiveShAcquiresAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveShAcquiresAcqShrdLitePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveShAcquiresAcqShrdLitePersec", (value))
}

// GetExecResourceRecursiveShAcquiresAcqShrdLitePersec gets the value of ExecResourceRecursiveShAcquiresAcqShrdLitePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveShAcquiresAcqShrdLitePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveShAcquiresAcqShrdLitePersec")
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

// SetExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec sets the value of ExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec", (value))
}

// GetExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec gets the value of ExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveShAcquiresAcqShrdStarveExclPersec")
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

// SetExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec sets the value of ExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec", (value))
}

// GetExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec gets the value of ExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceRecursiveShAcquiresAcqShrdWaitForExclPersec")
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

// SetExecResourceSetOwnerPointerExclusivePersec sets the value of ExecResourceSetOwnerPointerExclusivePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceSetOwnerPointerExclusivePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceSetOwnerPointerExclusivePersec", (value))
}

// GetExecResourceSetOwnerPointerExclusivePersec gets the value of ExecResourceSetOwnerPointerExclusivePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceSetOwnerPointerExclusivePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceSetOwnerPointerExclusivePersec")
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

// SetExecResourceSetOwnerPointerSharedExistingOwnerPersec sets the value of ExecResourceSetOwnerPointerSharedExistingOwnerPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceSetOwnerPointerSharedExistingOwnerPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceSetOwnerPointerSharedExistingOwnerPersec", (value))
}

// GetExecResourceSetOwnerPointerSharedExistingOwnerPersec gets the value of ExecResourceSetOwnerPointerSharedExistingOwnerPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceSetOwnerPointerSharedExistingOwnerPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceSetOwnerPointerSharedExistingOwnerPersec")
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

// SetExecResourceSetOwnerPointerSharedNewOwnerPersec sets the value of ExecResourceSetOwnerPointerSharedNewOwnerPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceSetOwnerPointerSharedNewOwnerPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceSetOwnerPointerSharedNewOwnerPersec", (value))
}

// GetExecResourceSetOwnerPointerSharedNewOwnerPersec gets the value of ExecResourceSetOwnerPointerSharedNewOwnerPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceSetOwnerPointerSharedNewOwnerPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceSetOwnerPointerSharedNewOwnerPersec")
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

// SetExecResourceTotalAcquiresPersec sets the value of ExecResourceTotalAcquiresPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalAcquiresPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalAcquiresPersec", (value))
}

// GetExecResourceTotalAcquiresPersec gets the value of ExecResourceTotalAcquiresPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalAcquiresPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalAcquiresPersec")
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

// SetExecResourceTotalContentionsPersec sets the value of ExecResourceTotalContentionsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalContentionsPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalContentionsPersec", (value))
}

// GetExecResourceTotalContentionsPersec gets the value of ExecResourceTotalContentionsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalContentionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalContentionsPersec")
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

// SetExecResourceTotalConvExclusiveToSharedPersec sets the value of ExecResourceTotalConvExclusiveToSharedPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalConvExclusiveToSharedPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalConvExclusiveToSharedPersec", (value))
}

// GetExecResourceTotalConvExclusiveToSharedPersec gets the value of ExecResourceTotalConvExclusiveToSharedPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalConvExclusiveToSharedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalConvExclusiveToSharedPersec")
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

// SetExecResourceTotalDeletePersec sets the value of ExecResourceTotalDeletePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalDeletePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalDeletePersec", (value))
}

// GetExecResourceTotalDeletePersec gets the value of ExecResourceTotalDeletePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalDeletePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalDeletePersec")
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

// SetExecResourceTotalExclusiveReleasesPersec sets the value of ExecResourceTotalExclusiveReleasesPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalExclusiveReleasesPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalExclusiveReleasesPersec", (value))
}

// GetExecResourceTotalExclusiveReleasesPersec gets the value of ExecResourceTotalExclusiveReleasesPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalExclusiveReleasesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalExclusiveReleasesPersec")
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

// SetExecResourceTotalInitializePersec sets the value of ExecResourceTotalInitializePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalInitializePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalInitializePersec", (value))
}

// GetExecResourceTotalInitializePersec gets the value of ExecResourceTotalInitializePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalInitializePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalInitializePersec")
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

// SetExecResourceTotalReInitializePersec sets the value of ExecResourceTotalReInitializePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalReInitializePersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalReInitializePersec", (value))
}

// GetExecResourceTotalReInitializePersec gets the value of ExecResourceTotalReInitializePersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalReInitializePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalReInitializePersec")
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

// SetExecResourceTotalSharedReleasesPersec sets the value of ExecResourceTotalSharedReleasesPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyExecResourceTotalSharedReleasesPersec(value uint32) (err error) {
	return instance.SetProperty("ExecResourceTotalSharedReleasesPersec", (value))
}

// GetExecResourceTotalSharedReleasesPersec gets the value of ExecResourceTotalSharedReleasesPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyExecResourceTotalSharedReleasesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExecResourceTotalSharedReleasesPersec")
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

// SetIPISendBroadcastRequestsPersec sets the value of IPISendBroadcastRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyIPISendBroadcastRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("IPISendBroadcastRequestsPersec", (value))
}

// GetIPISendBroadcastRequestsPersec gets the value of IPISendBroadcastRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyIPISendBroadcastRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPISendBroadcastRequestsPersec")
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

// SetIPISendRoutineRequestsPersec sets the value of IPISendRoutineRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyIPISendRoutineRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("IPISendRoutineRequestsPersec", (value))
}

// GetIPISendRoutineRequestsPersec gets the value of IPISendRoutineRequestsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyIPISendRoutineRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPISendRoutineRequestsPersec")
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

// SetIPISendSoftwareInterruptsPersec sets the value of IPISendSoftwareInterruptsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertyIPISendSoftwareInterruptsPersec(value uint32) (err error) {
	return instance.SetProperty("IPISendSoftwareInterruptsPersec", (value))
}

// GetIPISendSoftwareInterruptsPersec gets the value of IPISendSoftwareInterruptsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertyIPISendSoftwareInterruptsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPISendSoftwareInterruptsPersec")
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

// SetSpinlockAcquiresPersec sets the value of SpinlockAcquiresPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertySpinlockAcquiresPersec(value uint32) (err error) {
	return instance.SetProperty("SpinlockAcquiresPersec", (value))
}

// GetSpinlockAcquiresPersec gets the value of SpinlockAcquiresPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertySpinlockAcquiresPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpinlockAcquiresPersec")
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

// SetSpinlockContentionsPersec sets the value of SpinlockContentionsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertySpinlockContentionsPersec(value uint32) (err error) {
	return instance.SetProperty("SpinlockContentionsPersec", (value))
}

// GetSpinlockContentionsPersec gets the value of SpinlockContentionsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertySpinlockContentionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpinlockContentionsPersec")
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

// SetSpinlockSpinsPersec sets the value of SpinlockSpinsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) SetPropertySpinlockSpinsPersec(value uint32) (err error) {
	return instance.SetProperty("SpinlockSpinsPersec", (value))
}

// GetSpinlockSpinsPersec gets the value of SpinlockSpinsPersec for the instance
func (instance *Win32_PerfRawData_Counters_Synchronization) GetPropertySpinlockSpinsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpinlockSpinsPersec")
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
