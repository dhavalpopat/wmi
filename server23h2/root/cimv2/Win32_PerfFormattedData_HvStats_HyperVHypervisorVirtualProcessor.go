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

// Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor struct
type Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor struct {
	*Win32_PerfFormattedData

	//
	AddressDomainFlushesPersec uint64

	//
	AddressSpaceEvictionsPersec uint64

	//
	AddressSpaceFlushesPersec uint64

	//
	AddressSpaceSwitchesPersec uint64

	//
	APICEOIAccessesPersec uint64

	//
	APICIPIsSentPersec uint64

	//
	APICMMIOAccessesPersec uint64

	//
	APICSelfIPIsSentPersec uint64

	//
	APICTPRAccessesPersec uint64

	//
	BusLockAcquisitionsPersec uint64

	//
	ControlRegisterAccessesCost uint64

	//
	ControlRegisterAccessesForwardedPersec uint64

	//
	ControlRegisterAccessesForwardingCost uint64

	//
	ControlRegisterAccessesPersec uint64

	//
	CPPCRequestContextSwitchesPersec uint64

	//
	CPUGroupHypercallsPersec uint64

	//
	CPUIDInstructionsCost uint64

	//
	CPUIDInstructionsForwardedPersec uint64

	//
	CPUIDInstructionsForwardingCost uint64

	//
	CPUIDInstructionsPersec uint64

	//
	CPUWaitTimePerDispatch uint64

	//
	DebugRegisterAccessesCost uint64

	//
	DebugRegisterAccessesForwardedPersec uint64

	//
	DebugRegisterAccessesForwardingCost uint64

	//
	DebugRegisterAccessesPersec uint64

	//
	DepositHypercallsPersec uint64

	//
	DeviceDomainHypercallsPersec uint64

	//
	EmulatedInstructionsCost uint64

	//
	EmulatedInstructionsForwardedPersec uint64

	//
	EmulatedInstructionsForwardingCost uint64

	//
	EmulatedInstructionsPersec uint64

	//
	EventLogHypercallsPersec uint64

	//
	ExtendedHypercallInterceptMessagesPersec uint64

	//
	ExtendedHypercallsPersec uint64

	//
	ExternalInterruptsCost uint64

	//
	ExternalInterruptsForwardedPersec uint64

	//
	ExternalInterruptsPersec uint64

	//
	FlushPhysicalAddressListHypercallsPersec uint64

	//
	FlushPhysicalAddressSpaceHypercallsPersec uint64

	//
	GIFInstructionEmulationCost uint64

	//
	GIFInstructionEmulationInterceptsPersec uint64

	//
	GlobalGVARangeFlushesPersec uint64

	//
	GlobalIOTLBFlushCost uint64

	//
	GlobalIOTLBFlushesPersec uint64

	//
	GPASpaceHypercallsPersec uint64

	//
	GuestPageTableMapsPersec uint64

	//
	HardwareInterruptsPersec uint64

	//
	HLTInstructionsCost uint64

	//
	HLTInstructionsForwardedPersec uint64

	//
	HLTInstructionsForwardingCost uint64

	//
	HLTInstructionsPersec uint64

	//
	HypercallsCost uint64

	//
	HypercallsForwardedPersec uint64

	//
	HypercallsForwardingCost uint64

	//
	HypercallsPersec uint64

	//
	InvEptAllContextEmulationInterceptsPersec uint64

	//
	InvEptAllContextInstructionEmulationCost uint64

	//
	InvEptSingleContextEmulationInterceptsPersec uint64

	//
	InvEptSingleContextInstructionEmulationCost uint64

	//
	InvVpidAllContextEmulationInterceptsPersec uint64

	//
	InvVpidAllContextInstructionEmulationCost uint64

	//
	InvVpidSingleAddressEmulationInterceptsPersec uint64

	//
	InvVpidSingleAddressInstructionEmulationCost uint64

	//
	InvVpidSingleContextEmulationInterceptsPersec uint64

	//
	InvVpidSingleContextInstructionEmulationCost uint64

	//
	IOInstructionsCost uint64

	//
	IOInstructionsForwardedPersec uint64

	//
	IOInstructionsForwardingCost uint64

	//
	IOInstructionsPersec uint64

	//
	IOInterceptMessagesPersec uint64

	//
	IOMMUHypercallsPersec uint64

	//
	LargePageTLBFillsPersec uint64

	//
	LocalFlushedGVARangesPersec uint64

	//
	LocalIOTLBFlushCost uint64

	//
	LocalIOTLBFlushesPersec uint64

	//
	LogicalProcessorDispatchesPersec uint64

	//
	LogicalProcessorHypercallsPersec uint64

	//
	LogicalProcessorMigrationsPersec uint64

	//
	LongSpinWaitHypercallsPersec uint64

	//
	MBECNestedPageTableSwitchesPersec uint64

	//
	MemoryInterceptMessagesPersec uint64

	//
	MSRAccessesCost uint64

	//
	MSRAccessesForwardedPersec uint64

	//
	MSRAccessesForwardingCost uint64

	//
	MSRAccessesPersec uint64

	//
	MWAITInstructionsCost uint64

	//
	MWAITInstructionsForwardedPersec uint64

	//
	MWAITInstructionsForwardingCost uint64

	//
	MWAITInstructionsPersec uint64

	//
	NestedPageFaultInterceptsCost uint64

	//
	NestedPageFaultInterceptsPersec uint64

	//
	NestedSLATHardPageFaultsCost uint64

	//
	NestedSLATHardPageFaultsPersec uint64

	//
	NestedSLATSoftPageFaultsCost uint64

	//
	NestedSLATSoftPageFaultsPersec uint64

	//
	NestedTLBPageTableEvictionsPersec uint64

	//
	NestedTLBPageTableReclamationsPersec uint64

	//
	NestedVMEntriesCost uint64

	//
	NestedVMEntriesPersec uint64

	//
	OtherHypercallsPersec uint64

	//
	OtherInterceptsCost uint64

	//
	OtherInterceptsForwardedPersec uint64

	//
	OtherInterceptsForwardingCost uint64

	//
	OtherInterceptsPersec uint64

	//
	OtherMessagesPersec uint64

	//
	OtherReflectedGuestExceptionsPersec uint64

	//
	PageFaultInterceptsCost uint64

	//
	PageFaultInterceptsForwardedPersec uint64

	//
	PageFaultInterceptsForwardingCost uint64

	//
	PageFaultInterceptsPersec uint64

	//
	PageInvalidationsCost uint64

	//
	PageInvalidationsForwardedPersec uint64

	//
	PageInvalidationsForwardingCost uint64

	//
	PageInvalidationsPersec uint64

	//
	PageScansPersec uint64

	//
	PageTableAllocationsPersec uint64

	//
	PageTableEvictionsPersec uint64

	//
	PageTableReclamationsPersec uint64

	//
	PageTableResetsPersec uint64

	//
	PageTableValidationsPersec uint64

	//
	PageTableWriteInterceptsPersec uint64

	//
	PendingInterruptsCost uint64

	//
	PendingInterruptsForwardedPersec uint64

	//
	PendingInterruptsForwardingCost uint64

	//
	PendingInterruptsPersec uint64

	//
	PercentGuestRelativeUtilization uint64

	//
	PercentGuestRunTime uint64

	//
	PercentHypervisorRunTime uint64

	//
	PercentRemoteRunTime uint64

	//
	PercentTotalCoreRunTime uint64

	//
	PercentTotalRunTime uint64

	//
	PercentVTL1RunTime uint64

	//
	PercentVTL2RunTime uint64

	//
	PerformanceMonitoringInterruptsPersec uint64

	//
	PerformanceMonitoringIPTMSRAccessesPersec uint64

	//
	PerformanceMonitoringLBRMSRAccessesPersec uint64

	//
	PerformanceMonitoringvPMUMSRAccessesPersec uint64

	//
	PostedInterruptNotificationsPersec uint64

	//
	PostedInterruptScansPersec uint64

	//
	RDPMCInstructionsCost uint64

	//
	RDPMCInstructionsPersec uint64

	//
	ReflectedGuestPageFaultsPersec uint64

	//
	SchedulingPriority uint64

	//
	SmallPageTLBFillsPersec uint64

	//
	SVMHypercallsPersec uint64

	//
	SyntheticInterruptHypercallsPersec uint64

	//
	SyntheticInterruptsPersec uint64

	//
	TotalInterceptsCost uint64

	//
	TotalInterceptsPersec uint64

	//
	TotalMessagesPersec uint64

	//
	TotalVirtualizationInstructionsEmulatedPersec uint64

	//
	TotalVirtualizationInstructionsEmulationCost uint64

	//
	VirtualInterruptHypercallsPersec uint64

	//
	VirtualInterruptsPersec uint64

	//
	VirtualMMUHypercallsPersec uint64

	//
	VirtualProcessorHypercallsPersec uint64

	//
	VMCLEAREmulationInterceptsPersec uint64

	//
	VMCLEARInstructionEmulationCost uint64

	//
	VMLOADEmulationInterceptsPersec uint64

	//
	VMLOADInstructionEmulationCost uint64

	//
	VMPTRLDEmulationInterceptsPersec uint64

	//
	VMPTRLDInstructionEmulationCost uint64

	//
	VMPTRSTEmulationInterceptsPersec uint64

	//
	VMPTRSTInstructionEmulationCost uint64

	//
	VMREADEmulationInterceptsPersec uint64

	//
	VMREADInstructionEmulationCost uint64

	//
	VMSAVEEmulationInterceptsPersec uint64

	//
	VMSAVEInstructionEmulationCost uint64

	//
	VMWRITEEmulationInterceptsPersec uint64

	//
	VMWRITEInstructionEmulationCost uint64

	//
	VMXOFFEmulationInterceptsPersec uint64

	//
	VMXOFFInstructionEmulationCost uint64

	//
	VMXONEmulationInterceptsPersec uint64

	//
	VMXONInstructionEmulationCost uint64

	//
	VSMHypercallsPersec uint64

	//
	VTL1AverageRunTime uint64

	//
	VTL1DispatchesPersec uint64

	//
	VTL2AverageRunTime uint64

	//
	VTL2DispatchesPersec uint64
}

func NewWin32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessorEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAddressDomainFlushesPersec sets the value of AddressDomainFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAddressDomainFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("AddressDomainFlushesPersec", (value))
}

// GetAddressDomainFlushesPersec gets the value of AddressDomainFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAddressDomainFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AddressDomainFlushesPersec")
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

// SetAddressSpaceEvictionsPersec sets the value of AddressSpaceEvictionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAddressSpaceEvictionsPersec(value uint64) (err error) {
	return instance.SetProperty("AddressSpaceEvictionsPersec", (value))
}

// GetAddressSpaceEvictionsPersec gets the value of AddressSpaceEvictionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAddressSpaceEvictionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AddressSpaceEvictionsPersec")
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

// SetAddressSpaceFlushesPersec sets the value of AddressSpaceFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAddressSpaceFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("AddressSpaceFlushesPersec", (value))
}

// GetAddressSpaceFlushesPersec gets the value of AddressSpaceFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAddressSpaceFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AddressSpaceFlushesPersec")
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

// SetAddressSpaceSwitchesPersec sets the value of AddressSpaceSwitchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAddressSpaceSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("AddressSpaceSwitchesPersec", (value))
}

// GetAddressSpaceSwitchesPersec gets the value of AddressSpaceSwitchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAddressSpaceSwitchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AddressSpaceSwitchesPersec")
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

// SetAPICEOIAccessesPersec sets the value of APICEOIAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAPICEOIAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("APICEOIAccessesPersec", (value))
}

// GetAPICEOIAccessesPersec gets the value of APICEOIAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAPICEOIAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("APICEOIAccessesPersec")
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

// SetAPICIPIsSentPersec sets the value of APICIPIsSentPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAPICIPIsSentPersec(value uint64) (err error) {
	return instance.SetProperty("APICIPIsSentPersec", (value))
}

// GetAPICIPIsSentPersec gets the value of APICIPIsSentPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAPICIPIsSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("APICIPIsSentPersec")
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

// SetAPICMMIOAccessesPersec sets the value of APICMMIOAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAPICMMIOAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("APICMMIOAccessesPersec", (value))
}

// GetAPICMMIOAccessesPersec gets the value of APICMMIOAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAPICMMIOAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("APICMMIOAccessesPersec")
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

// SetAPICSelfIPIsSentPersec sets the value of APICSelfIPIsSentPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAPICSelfIPIsSentPersec(value uint64) (err error) {
	return instance.SetProperty("APICSelfIPIsSentPersec", (value))
}

// GetAPICSelfIPIsSentPersec gets the value of APICSelfIPIsSentPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAPICSelfIPIsSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("APICSelfIPIsSentPersec")
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

// SetAPICTPRAccessesPersec sets the value of APICTPRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyAPICTPRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("APICTPRAccessesPersec", (value))
}

// GetAPICTPRAccessesPersec gets the value of APICTPRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyAPICTPRAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("APICTPRAccessesPersec")
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

// SetBusLockAcquisitionsPersec sets the value of BusLockAcquisitionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyBusLockAcquisitionsPersec(value uint64) (err error) {
	return instance.SetProperty("BusLockAcquisitionsPersec", (value))
}

// GetBusLockAcquisitionsPersec gets the value of BusLockAcquisitionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyBusLockAcquisitionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BusLockAcquisitionsPersec")
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

// SetControlRegisterAccessesCost sets the value of ControlRegisterAccessesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyControlRegisterAccessesCost(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesCost", (value))
}

// GetControlRegisterAccessesCost gets the value of ControlRegisterAccessesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyControlRegisterAccessesCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRegisterAccessesCost")
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

// SetControlRegisterAccessesForwardedPersec sets the value of ControlRegisterAccessesForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyControlRegisterAccessesForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesForwardedPersec", (value))
}

// GetControlRegisterAccessesForwardedPersec gets the value of ControlRegisterAccessesForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyControlRegisterAccessesForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRegisterAccessesForwardedPersec")
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

// SetControlRegisterAccessesForwardingCost sets the value of ControlRegisterAccessesForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyControlRegisterAccessesForwardingCost(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesForwardingCost", (value))
}

// GetControlRegisterAccessesForwardingCost gets the value of ControlRegisterAccessesForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyControlRegisterAccessesForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRegisterAccessesForwardingCost")
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

// SetControlRegisterAccessesPersec sets the value of ControlRegisterAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyControlRegisterAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesPersec", (value))
}

// GetControlRegisterAccessesPersec gets the value of ControlRegisterAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyControlRegisterAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRegisterAccessesPersec")
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

// SetCPPCRequestContextSwitchesPersec sets the value of CPPCRequestContextSwitchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPPCRequestContextSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("CPPCRequestContextSwitchesPersec", (value))
}

// GetCPPCRequestContextSwitchesPersec gets the value of CPPCRequestContextSwitchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPPCRequestContextSwitchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPPCRequestContextSwitchesPersec")
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

// SetCPUGroupHypercallsPersec sets the value of CPUGroupHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPUGroupHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("CPUGroupHypercallsPersec", (value))
}

// GetCPUGroupHypercallsPersec gets the value of CPUGroupHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPUGroupHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUGroupHypercallsPersec")
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

// SetCPUIDInstructionsCost sets the value of CPUIDInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPUIDInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsCost", (value))
}

// GetCPUIDInstructionsCost gets the value of CPUIDInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPUIDInstructionsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUIDInstructionsCost")
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

// SetCPUIDInstructionsForwardedPersec sets the value of CPUIDInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPUIDInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsForwardedPersec", (value))
}

// GetCPUIDInstructionsForwardedPersec gets the value of CPUIDInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPUIDInstructionsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUIDInstructionsForwardedPersec")
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

// SetCPUIDInstructionsForwardingCost sets the value of CPUIDInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPUIDInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsForwardingCost", (value))
}

// GetCPUIDInstructionsForwardingCost gets the value of CPUIDInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPUIDInstructionsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUIDInstructionsForwardingCost")
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

// SetCPUIDInstructionsPersec sets the value of CPUIDInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPUIDInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsPersec", (value))
}

// GetCPUIDInstructionsPersec gets the value of CPUIDInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPUIDInstructionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUIDInstructionsPersec")
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

// SetCPUWaitTimePerDispatch sets the value of CPUWaitTimePerDispatch for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyCPUWaitTimePerDispatch(value uint64) (err error) {
	return instance.SetProperty("CPUWaitTimePerDispatch", (value))
}

// GetCPUWaitTimePerDispatch gets the value of CPUWaitTimePerDispatch for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyCPUWaitTimePerDispatch() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUWaitTimePerDispatch")
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

// SetDebugRegisterAccessesCost sets the value of DebugRegisterAccessesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyDebugRegisterAccessesCost(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesCost", (value))
}

// GetDebugRegisterAccessesCost gets the value of DebugRegisterAccessesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyDebugRegisterAccessesCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("DebugRegisterAccessesCost")
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

// SetDebugRegisterAccessesForwardedPersec sets the value of DebugRegisterAccessesForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyDebugRegisterAccessesForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesForwardedPersec", (value))
}

// GetDebugRegisterAccessesForwardedPersec gets the value of DebugRegisterAccessesForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyDebugRegisterAccessesForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DebugRegisterAccessesForwardedPersec")
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

// SetDebugRegisterAccessesForwardingCost sets the value of DebugRegisterAccessesForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyDebugRegisterAccessesForwardingCost(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesForwardingCost", (value))
}

// GetDebugRegisterAccessesForwardingCost gets the value of DebugRegisterAccessesForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyDebugRegisterAccessesForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("DebugRegisterAccessesForwardingCost")
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

// SetDebugRegisterAccessesPersec sets the value of DebugRegisterAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyDebugRegisterAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesPersec", (value))
}

// GetDebugRegisterAccessesPersec gets the value of DebugRegisterAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyDebugRegisterAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DebugRegisterAccessesPersec")
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

// SetDepositHypercallsPersec sets the value of DepositHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyDepositHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("DepositHypercallsPersec", (value))
}

// GetDepositHypercallsPersec gets the value of DepositHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyDepositHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DepositHypercallsPersec")
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

// SetDeviceDomainHypercallsPersec sets the value of DeviceDomainHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyDeviceDomainHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("DeviceDomainHypercallsPersec", (value))
}

// GetDeviceDomainHypercallsPersec gets the value of DeviceDomainHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyDeviceDomainHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceDomainHypercallsPersec")
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

// SetEmulatedInstructionsCost sets the value of EmulatedInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyEmulatedInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsCost", (value))
}

// GetEmulatedInstructionsCost gets the value of EmulatedInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyEmulatedInstructionsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("EmulatedInstructionsCost")
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

// SetEmulatedInstructionsForwardedPersec sets the value of EmulatedInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyEmulatedInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsForwardedPersec", (value))
}

// GetEmulatedInstructionsForwardedPersec gets the value of EmulatedInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyEmulatedInstructionsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EmulatedInstructionsForwardedPersec")
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

// SetEmulatedInstructionsForwardingCost sets the value of EmulatedInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyEmulatedInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsForwardingCost", (value))
}

// GetEmulatedInstructionsForwardingCost gets the value of EmulatedInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyEmulatedInstructionsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("EmulatedInstructionsForwardingCost")
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

// SetEmulatedInstructionsPersec sets the value of EmulatedInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyEmulatedInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsPersec", (value))
}

// GetEmulatedInstructionsPersec gets the value of EmulatedInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyEmulatedInstructionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EmulatedInstructionsPersec")
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

// SetEventLogHypercallsPersec sets the value of EventLogHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyEventLogHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("EventLogHypercallsPersec", (value))
}

// GetEventLogHypercallsPersec gets the value of EventLogHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyEventLogHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventLogHypercallsPersec")
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

// SetExtendedHypercallInterceptMessagesPersec sets the value of ExtendedHypercallInterceptMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyExtendedHypercallInterceptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("ExtendedHypercallInterceptMessagesPersec", (value))
}

// GetExtendedHypercallInterceptMessagesPersec gets the value of ExtendedHypercallInterceptMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyExtendedHypercallInterceptMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtendedHypercallInterceptMessagesPersec")
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

// SetExtendedHypercallsPersec sets the value of ExtendedHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyExtendedHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("ExtendedHypercallsPersec", (value))
}

// GetExtendedHypercallsPersec gets the value of ExtendedHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyExtendedHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtendedHypercallsPersec")
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

// SetExternalInterruptsCost sets the value of ExternalInterruptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyExternalInterruptsCost(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsCost", (value))
}

// GetExternalInterruptsCost gets the value of ExternalInterruptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyExternalInterruptsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExternalInterruptsCost")
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

// SetExternalInterruptsForwardedPersec sets the value of ExternalInterruptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyExternalInterruptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsForwardedPersec", (value))
}

// GetExternalInterruptsForwardedPersec gets the value of ExternalInterruptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyExternalInterruptsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExternalInterruptsForwardedPersec")
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

// SetExternalInterruptsPersec sets the value of ExternalInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyExternalInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsPersec", (value))
}

// GetExternalInterruptsPersec gets the value of ExternalInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyExternalInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExternalInterruptsPersec")
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

// SetFlushPhysicalAddressListHypercallsPersec sets the value of FlushPhysicalAddressListHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyFlushPhysicalAddressListHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("FlushPhysicalAddressListHypercallsPersec", (value))
}

// GetFlushPhysicalAddressListHypercallsPersec gets the value of FlushPhysicalAddressListHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyFlushPhysicalAddressListHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushPhysicalAddressListHypercallsPersec")
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

// SetFlushPhysicalAddressSpaceHypercallsPersec sets the value of FlushPhysicalAddressSpaceHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyFlushPhysicalAddressSpaceHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("FlushPhysicalAddressSpaceHypercallsPersec", (value))
}

// GetFlushPhysicalAddressSpaceHypercallsPersec gets the value of FlushPhysicalAddressSpaceHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyFlushPhysicalAddressSpaceHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushPhysicalAddressSpaceHypercallsPersec")
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

// SetGIFInstructionEmulationCost sets the value of GIFInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGIFInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("GIFInstructionEmulationCost", (value))
}

// GetGIFInstructionEmulationCost gets the value of GIFInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGIFInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("GIFInstructionEmulationCost")
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

// SetGIFInstructionEmulationInterceptsPersec sets the value of GIFInstructionEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGIFInstructionEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("GIFInstructionEmulationInterceptsPersec", (value))
}

// GetGIFInstructionEmulationInterceptsPersec gets the value of GIFInstructionEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGIFInstructionEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GIFInstructionEmulationInterceptsPersec")
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

// SetGlobalGVARangeFlushesPersec sets the value of GlobalGVARangeFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGlobalGVARangeFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("GlobalGVARangeFlushesPersec", (value))
}

// GetGlobalGVARangeFlushesPersec gets the value of GlobalGVARangeFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGlobalGVARangeFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GlobalGVARangeFlushesPersec")
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

// SetGlobalIOTLBFlushCost sets the value of GlobalIOTLBFlushCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGlobalIOTLBFlushCost(value uint64) (err error) {
	return instance.SetProperty("GlobalIOTLBFlushCost", (value))
}

// GetGlobalIOTLBFlushCost gets the value of GlobalIOTLBFlushCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGlobalIOTLBFlushCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("GlobalIOTLBFlushCost")
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

// SetGlobalIOTLBFlushesPersec sets the value of GlobalIOTLBFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGlobalIOTLBFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("GlobalIOTLBFlushesPersec", (value))
}

// GetGlobalIOTLBFlushesPersec gets the value of GlobalIOTLBFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGlobalIOTLBFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GlobalIOTLBFlushesPersec")
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

// SetGPASpaceHypercallsPersec sets the value of GPASpaceHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGPASpaceHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("GPASpaceHypercallsPersec", (value))
}

// GetGPASpaceHypercallsPersec gets the value of GPASpaceHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGPASpaceHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GPASpaceHypercallsPersec")
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

// SetGuestPageTableMapsPersec sets the value of GuestPageTableMapsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyGuestPageTableMapsPersec(value uint64) (err error) {
	return instance.SetProperty("GuestPageTableMapsPersec", (value))
}

// GetGuestPageTableMapsPersec gets the value of GuestPageTableMapsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyGuestPageTableMapsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GuestPageTableMapsPersec")
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

// SetHardwareInterruptsPersec sets the value of HardwareInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHardwareInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("HardwareInterruptsPersec", (value))
}

// GetHardwareInterruptsPersec gets the value of HardwareInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHardwareInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HardwareInterruptsPersec")
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

// SetHLTInstructionsCost sets the value of HLTInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHLTInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsCost", (value))
}

// GetHLTInstructionsCost gets the value of HLTInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHLTInstructionsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("HLTInstructionsCost")
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

// SetHLTInstructionsForwardedPersec sets the value of HLTInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHLTInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsForwardedPersec", (value))
}

// GetHLTInstructionsForwardedPersec gets the value of HLTInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHLTInstructionsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HLTInstructionsForwardedPersec")
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

// SetHLTInstructionsForwardingCost sets the value of HLTInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHLTInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsForwardingCost", (value))
}

// GetHLTInstructionsForwardingCost gets the value of HLTInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHLTInstructionsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("HLTInstructionsForwardingCost")
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

// SetHLTInstructionsPersec sets the value of HLTInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHLTInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsPersec", (value))
}

// GetHLTInstructionsPersec gets the value of HLTInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHLTInstructionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HLTInstructionsPersec")
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

// SetHypercallsCost sets the value of HypercallsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHypercallsCost(value uint64) (err error) {
	return instance.SetProperty("HypercallsCost", (value))
}

// GetHypercallsCost gets the value of HypercallsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHypercallsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypercallsCost")
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

// SetHypercallsForwardedPersec sets the value of HypercallsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHypercallsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("HypercallsForwardedPersec", (value))
}

// GetHypercallsForwardedPersec gets the value of HypercallsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHypercallsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypercallsForwardedPersec")
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

// SetHypercallsForwardingCost sets the value of HypercallsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHypercallsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("HypercallsForwardingCost", (value))
}

// GetHypercallsForwardingCost gets the value of HypercallsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHypercallsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypercallsForwardingCost")
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

// SetHypercallsPersec sets the value of HypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("HypercallsPersec", (value))
}

// GetHypercallsPersec gets the value of HypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypercallsPersec")
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

// SetInvEptAllContextEmulationInterceptsPersec sets the value of InvEptAllContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvEptAllContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvEptAllContextEmulationInterceptsPersec", (value))
}

// GetInvEptAllContextEmulationInterceptsPersec gets the value of InvEptAllContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvEptAllContextEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvEptAllContextEmulationInterceptsPersec")
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

// SetInvEptAllContextInstructionEmulationCost sets the value of InvEptAllContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvEptAllContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvEptAllContextInstructionEmulationCost", (value))
}

// GetInvEptAllContextInstructionEmulationCost gets the value of InvEptAllContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvEptAllContextInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvEptAllContextInstructionEmulationCost")
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

// SetInvEptSingleContextEmulationInterceptsPersec sets the value of InvEptSingleContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvEptSingleContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvEptSingleContextEmulationInterceptsPersec", (value))
}

// GetInvEptSingleContextEmulationInterceptsPersec gets the value of InvEptSingleContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvEptSingleContextEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvEptSingleContextEmulationInterceptsPersec")
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

// SetInvEptSingleContextInstructionEmulationCost sets the value of InvEptSingleContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvEptSingleContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvEptSingleContextInstructionEmulationCost", (value))
}

// GetInvEptSingleContextInstructionEmulationCost gets the value of InvEptSingleContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvEptSingleContextInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvEptSingleContextInstructionEmulationCost")
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

// SetInvVpidAllContextEmulationInterceptsPersec sets the value of InvVpidAllContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvVpidAllContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvVpidAllContextEmulationInterceptsPersec", (value))
}

// GetInvVpidAllContextEmulationInterceptsPersec gets the value of InvVpidAllContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvVpidAllContextEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidAllContextEmulationInterceptsPersec")
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

// SetInvVpidAllContextInstructionEmulationCost sets the value of InvVpidAllContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvVpidAllContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvVpidAllContextInstructionEmulationCost", (value))
}

// GetInvVpidAllContextInstructionEmulationCost gets the value of InvVpidAllContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvVpidAllContextInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidAllContextInstructionEmulationCost")
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

// SetInvVpidSingleAddressEmulationInterceptsPersec sets the value of InvVpidSingleAddressEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvVpidSingleAddressEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleAddressEmulationInterceptsPersec", (value))
}

// GetInvVpidSingleAddressEmulationInterceptsPersec gets the value of InvVpidSingleAddressEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvVpidSingleAddressEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidSingleAddressEmulationInterceptsPersec")
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

// SetInvVpidSingleAddressInstructionEmulationCost sets the value of InvVpidSingleAddressInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvVpidSingleAddressInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleAddressInstructionEmulationCost", (value))
}

// GetInvVpidSingleAddressInstructionEmulationCost gets the value of InvVpidSingleAddressInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvVpidSingleAddressInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidSingleAddressInstructionEmulationCost")
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

// SetInvVpidSingleContextEmulationInterceptsPersec sets the value of InvVpidSingleContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvVpidSingleContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleContextEmulationInterceptsPersec", (value))
}

// GetInvVpidSingleContextEmulationInterceptsPersec gets the value of InvVpidSingleContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvVpidSingleContextEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidSingleContextEmulationInterceptsPersec")
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

// SetInvVpidSingleContextInstructionEmulationCost sets the value of InvVpidSingleContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyInvVpidSingleContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleContextInstructionEmulationCost", (value))
}

// GetInvVpidSingleContextInstructionEmulationCost gets the value of InvVpidSingleContextInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyInvVpidSingleContextInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidSingleContextInstructionEmulationCost")
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

// SetIOInstructionsCost sets the value of IOInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyIOInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsCost", (value))
}

// GetIOInstructionsCost gets the value of IOInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyIOInstructionsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInstructionsCost")
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

// SetIOInstructionsForwardedPersec sets the value of IOInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyIOInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsForwardedPersec", (value))
}

// GetIOInstructionsForwardedPersec gets the value of IOInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyIOInstructionsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInstructionsForwardedPersec")
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

// SetIOInstructionsForwardingCost sets the value of IOInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyIOInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsForwardingCost", (value))
}

// GetIOInstructionsForwardingCost gets the value of IOInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyIOInstructionsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInstructionsForwardingCost")
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

// SetIOInstructionsPersec sets the value of IOInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyIOInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsPersec", (value))
}

// GetIOInstructionsPersec gets the value of IOInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyIOInstructionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInstructionsPersec")
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

// SetIOInterceptMessagesPersec sets the value of IOInterceptMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyIOInterceptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("IOInterceptMessagesPersec", (value))
}

// GetIOInterceptMessagesPersec gets the value of IOInterceptMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyIOInterceptMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInterceptMessagesPersec")
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

// SetIOMMUHypercallsPersec sets the value of IOMMUHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyIOMMUHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("IOMMUHypercallsPersec", (value))
}

// GetIOMMUHypercallsPersec gets the value of IOMMUHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyIOMMUHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOMMUHypercallsPersec")
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

// SetLargePageTLBFillsPersec sets the value of LargePageTLBFillsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLargePageTLBFillsPersec(value uint64) (err error) {
	return instance.SetProperty("LargePageTLBFillsPersec", (value))
}

// GetLargePageTLBFillsPersec gets the value of LargePageTLBFillsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLargePageTLBFillsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LargePageTLBFillsPersec")
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

// SetLocalFlushedGVARangesPersec sets the value of LocalFlushedGVARangesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLocalFlushedGVARangesPersec(value uint64) (err error) {
	return instance.SetProperty("LocalFlushedGVARangesPersec", (value))
}

// GetLocalFlushedGVARangesPersec gets the value of LocalFlushedGVARangesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLocalFlushedGVARangesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalFlushedGVARangesPersec")
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

// SetLocalIOTLBFlushCost sets the value of LocalIOTLBFlushCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLocalIOTLBFlushCost(value uint64) (err error) {
	return instance.SetProperty("LocalIOTLBFlushCost", (value))
}

// GetLocalIOTLBFlushCost gets the value of LocalIOTLBFlushCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLocalIOTLBFlushCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalIOTLBFlushCost")
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

// SetLocalIOTLBFlushesPersec sets the value of LocalIOTLBFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLocalIOTLBFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("LocalIOTLBFlushesPersec", (value))
}

// GetLocalIOTLBFlushesPersec gets the value of LocalIOTLBFlushesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLocalIOTLBFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalIOTLBFlushesPersec")
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

// SetLogicalProcessorDispatchesPersec sets the value of LogicalProcessorDispatchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLogicalProcessorDispatchesPersec(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessorDispatchesPersec", (value))
}

// GetLogicalProcessorDispatchesPersec gets the value of LogicalProcessorDispatchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLogicalProcessorDispatchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalProcessorDispatchesPersec")
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

// SetLogicalProcessorHypercallsPersec sets the value of LogicalProcessorHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLogicalProcessorHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessorHypercallsPersec", (value))
}

// GetLogicalProcessorHypercallsPersec gets the value of LogicalProcessorHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLogicalProcessorHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalProcessorHypercallsPersec")
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

// SetLogicalProcessorMigrationsPersec sets the value of LogicalProcessorMigrationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLogicalProcessorMigrationsPersec(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessorMigrationsPersec", (value))
}

// GetLogicalProcessorMigrationsPersec gets the value of LogicalProcessorMigrationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLogicalProcessorMigrationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalProcessorMigrationsPersec")
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

// SetLongSpinWaitHypercallsPersec sets the value of LongSpinWaitHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyLongSpinWaitHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("LongSpinWaitHypercallsPersec", (value))
}

// GetLongSpinWaitHypercallsPersec gets the value of LongSpinWaitHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyLongSpinWaitHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LongSpinWaitHypercallsPersec")
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

// SetMBECNestedPageTableSwitchesPersec sets the value of MBECNestedPageTableSwitchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMBECNestedPageTableSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("MBECNestedPageTableSwitchesPersec", (value))
}

// GetMBECNestedPageTableSwitchesPersec gets the value of MBECNestedPageTableSwitchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMBECNestedPageTableSwitchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MBECNestedPageTableSwitchesPersec")
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

// SetMemoryInterceptMessagesPersec sets the value of MemoryInterceptMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMemoryInterceptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("MemoryInterceptMessagesPersec", (value))
}

// GetMemoryInterceptMessagesPersec gets the value of MemoryInterceptMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMemoryInterceptMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryInterceptMessagesPersec")
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

// SetMSRAccessesCost sets the value of MSRAccessesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMSRAccessesCost(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesCost", (value))
}

// GetMSRAccessesCost gets the value of MSRAccessesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMSRAccessesCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSRAccessesCost")
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

// SetMSRAccessesForwardedPersec sets the value of MSRAccessesForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMSRAccessesForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesForwardedPersec", (value))
}

// GetMSRAccessesForwardedPersec gets the value of MSRAccessesForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMSRAccessesForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSRAccessesForwardedPersec")
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

// SetMSRAccessesForwardingCost sets the value of MSRAccessesForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMSRAccessesForwardingCost(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesForwardingCost", (value))
}

// GetMSRAccessesForwardingCost gets the value of MSRAccessesForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMSRAccessesForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSRAccessesForwardingCost")
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

// SetMSRAccessesPersec sets the value of MSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesPersec", (value))
}

// GetMSRAccessesPersec gets the value of MSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMSRAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSRAccessesPersec")
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

// SetMWAITInstructionsCost sets the value of MWAITInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMWAITInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsCost", (value))
}

// GetMWAITInstructionsCost gets the value of MWAITInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMWAITInstructionsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("MWAITInstructionsCost")
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

// SetMWAITInstructionsForwardedPersec sets the value of MWAITInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMWAITInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsForwardedPersec", (value))
}

// GetMWAITInstructionsForwardedPersec gets the value of MWAITInstructionsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMWAITInstructionsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MWAITInstructionsForwardedPersec")
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

// SetMWAITInstructionsForwardingCost sets the value of MWAITInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMWAITInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsForwardingCost", (value))
}

// GetMWAITInstructionsForwardingCost gets the value of MWAITInstructionsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMWAITInstructionsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("MWAITInstructionsForwardingCost")
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

// SetMWAITInstructionsPersec sets the value of MWAITInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyMWAITInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsPersec", (value))
}

// GetMWAITInstructionsPersec gets the value of MWAITInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyMWAITInstructionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MWAITInstructionsPersec")
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

// SetNestedPageFaultInterceptsCost sets the value of NestedPageFaultInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedPageFaultInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("NestedPageFaultInterceptsCost", (value))
}

// GetNestedPageFaultInterceptsCost gets the value of NestedPageFaultInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedPageFaultInterceptsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedPageFaultInterceptsCost")
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

// SetNestedPageFaultInterceptsPersec sets the value of NestedPageFaultInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedPageFaultInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedPageFaultInterceptsPersec", (value))
}

// GetNestedPageFaultInterceptsPersec gets the value of NestedPageFaultInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedPageFaultInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedPageFaultInterceptsPersec")
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

// SetNestedSLATHardPageFaultsCost sets the value of NestedSLATHardPageFaultsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedSLATHardPageFaultsCost(value uint64) (err error) {
	return instance.SetProperty("NestedSLATHardPageFaultsCost", (value))
}

// GetNestedSLATHardPageFaultsCost gets the value of NestedSLATHardPageFaultsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedSLATHardPageFaultsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedSLATHardPageFaultsCost")
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

// SetNestedSLATHardPageFaultsPersec sets the value of NestedSLATHardPageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedSLATHardPageFaultsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedSLATHardPageFaultsPersec", (value))
}

// GetNestedSLATHardPageFaultsPersec gets the value of NestedSLATHardPageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedSLATHardPageFaultsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedSLATHardPageFaultsPersec")
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

// SetNestedSLATSoftPageFaultsCost sets the value of NestedSLATSoftPageFaultsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedSLATSoftPageFaultsCost(value uint64) (err error) {
	return instance.SetProperty("NestedSLATSoftPageFaultsCost", (value))
}

// GetNestedSLATSoftPageFaultsCost gets the value of NestedSLATSoftPageFaultsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedSLATSoftPageFaultsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedSLATSoftPageFaultsCost")
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

// SetNestedSLATSoftPageFaultsPersec sets the value of NestedSLATSoftPageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedSLATSoftPageFaultsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedSLATSoftPageFaultsPersec", (value))
}

// GetNestedSLATSoftPageFaultsPersec gets the value of NestedSLATSoftPageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedSLATSoftPageFaultsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedSLATSoftPageFaultsPersec")
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

// SetNestedTLBPageTableEvictionsPersec sets the value of NestedTLBPageTableEvictionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedTLBPageTableEvictionsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedTLBPageTableEvictionsPersec", (value))
}

// GetNestedTLBPageTableEvictionsPersec gets the value of NestedTLBPageTableEvictionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedTLBPageTableEvictionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedTLBPageTableEvictionsPersec")
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

// SetNestedTLBPageTableReclamationsPersec sets the value of NestedTLBPageTableReclamationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedTLBPageTableReclamationsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedTLBPageTableReclamationsPersec", (value))
}

// GetNestedTLBPageTableReclamationsPersec gets the value of NestedTLBPageTableReclamationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedTLBPageTableReclamationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedTLBPageTableReclamationsPersec")
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

// SetNestedVMEntriesCost sets the value of NestedVMEntriesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedVMEntriesCost(value uint64) (err error) {
	return instance.SetProperty("NestedVMEntriesCost", (value))
}

// GetNestedVMEntriesCost gets the value of NestedVMEntriesCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedVMEntriesCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedVMEntriesCost")
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

// SetNestedVMEntriesPersec sets the value of NestedVMEntriesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyNestedVMEntriesPersec(value uint64) (err error) {
	return instance.SetProperty("NestedVMEntriesPersec", (value))
}

// GetNestedVMEntriesPersec gets the value of NestedVMEntriesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyNestedVMEntriesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedVMEntriesPersec")
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

// SetOtherHypercallsPersec sets the value of OtherHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("OtherHypercallsPersec", (value))
}

// GetOtherHypercallsPersec gets the value of OtherHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherHypercallsPersec")
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

// SetOtherInterceptsCost sets the value of OtherInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsCost", (value))
}

// GetOtherInterceptsCost gets the value of OtherInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherInterceptsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherInterceptsCost")
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

// SetOtherInterceptsForwardedPersec sets the value of OtherInterceptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherInterceptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsForwardedPersec", (value))
}

// GetOtherInterceptsForwardedPersec gets the value of OtherInterceptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherInterceptsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherInterceptsForwardedPersec")
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

// SetOtherInterceptsForwardingCost sets the value of OtherInterceptsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherInterceptsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsForwardingCost", (value))
}

// GetOtherInterceptsForwardingCost gets the value of OtherInterceptsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherInterceptsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherInterceptsForwardingCost")
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

// SetOtherInterceptsPersec sets the value of OtherInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsPersec", (value))
}

// GetOtherInterceptsPersec gets the value of OtherInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherInterceptsPersec")
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

// SetOtherMessagesPersec sets the value of OtherMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("OtherMessagesPersec", (value))
}

// GetOtherMessagesPersec gets the value of OtherMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherMessagesPersec")
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

// SetOtherReflectedGuestExceptionsPersec sets the value of OtherReflectedGuestExceptionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyOtherReflectedGuestExceptionsPersec(value uint64) (err error) {
	return instance.SetProperty("OtherReflectedGuestExceptionsPersec", (value))
}

// GetOtherReflectedGuestExceptionsPersec gets the value of OtherReflectedGuestExceptionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyOtherReflectedGuestExceptionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherReflectedGuestExceptionsPersec")
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

// SetPageFaultInterceptsCost sets the value of PageFaultInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageFaultInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsCost", (value))
}

// GetPageFaultInterceptsCost gets the value of PageFaultInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageFaultInterceptsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFaultInterceptsCost")
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

// SetPageFaultInterceptsForwardedPersec sets the value of PageFaultInterceptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageFaultInterceptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsForwardedPersec", (value))
}

// GetPageFaultInterceptsForwardedPersec gets the value of PageFaultInterceptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageFaultInterceptsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFaultInterceptsForwardedPersec")
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

// SetPageFaultInterceptsForwardingCost sets the value of PageFaultInterceptsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageFaultInterceptsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsForwardingCost", (value))
}

// GetPageFaultInterceptsForwardingCost gets the value of PageFaultInterceptsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageFaultInterceptsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFaultInterceptsForwardingCost")
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

// SetPageFaultInterceptsPersec sets the value of PageFaultInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageFaultInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsPersec", (value))
}

// GetPageFaultInterceptsPersec gets the value of PageFaultInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageFaultInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFaultInterceptsPersec")
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

// SetPageInvalidationsCost sets the value of PageInvalidationsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageInvalidationsCost(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsCost", (value))
}

// GetPageInvalidationsCost gets the value of PageInvalidationsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageInvalidationsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageInvalidationsCost")
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

// SetPageInvalidationsForwardedPersec sets the value of PageInvalidationsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageInvalidationsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsForwardedPersec", (value))
}

// GetPageInvalidationsForwardedPersec gets the value of PageInvalidationsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageInvalidationsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageInvalidationsForwardedPersec")
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

// SetPageInvalidationsForwardingCost sets the value of PageInvalidationsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageInvalidationsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsForwardingCost", (value))
}

// GetPageInvalidationsForwardingCost gets the value of PageInvalidationsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageInvalidationsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageInvalidationsForwardingCost")
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

// SetPageInvalidationsPersec sets the value of PageInvalidationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageInvalidationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsPersec", (value))
}

// GetPageInvalidationsPersec gets the value of PageInvalidationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageInvalidationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageInvalidationsPersec")
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

// SetPageScansPersec sets the value of PageScansPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageScansPersec(value uint64) (err error) {
	return instance.SetProperty("PageScansPersec", (value))
}

// GetPageScansPersec gets the value of PageScansPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageScansPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageScansPersec")
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

// SetPageTableAllocationsPersec sets the value of PageTableAllocationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageTableAllocationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableAllocationsPersec", (value))
}

// GetPageTableAllocationsPersec gets the value of PageTableAllocationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageTableAllocationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageTableAllocationsPersec")
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

// SetPageTableEvictionsPersec sets the value of PageTableEvictionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageTableEvictionsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableEvictionsPersec", (value))
}

// GetPageTableEvictionsPersec gets the value of PageTableEvictionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageTableEvictionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageTableEvictionsPersec")
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

// SetPageTableReclamationsPersec sets the value of PageTableReclamationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageTableReclamationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableReclamationsPersec", (value))
}

// GetPageTableReclamationsPersec gets the value of PageTableReclamationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageTableReclamationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageTableReclamationsPersec")
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

// SetPageTableResetsPersec sets the value of PageTableResetsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageTableResetsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableResetsPersec", (value))
}

// GetPageTableResetsPersec gets the value of PageTableResetsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageTableResetsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageTableResetsPersec")
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

// SetPageTableValidationsPersec sets the value of PageTableValidationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageTableValidationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableValidationsPersec", (value))
}

// GetPageTableValidationsPersec gets the value of PageTableValidationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageTableValidationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageTableValidationsPersec")
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

// SetPageTableWriteInterceptsPersec sets the value of PageTableWriteInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPageTableWriteInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableWriteInterceptsPersec", (value))
}

// GetPageTableWriteInterceptsPersec gets the value of PageTableWriteInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPageTableWriteInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageTableWriteInterceptsPersec")
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

// SetPendingInterruptsCost sets the value of PendingInterruptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPendingInterruptsCost(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsCost", (value))
}

// GetPendingInterruptsCost gets the value of PendingInterruptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPendingInterruptsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingInterruptsCost")
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

// SetPendingInterruptsForwardedPersec sets the value of PendingInterruptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPendingInterruptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsForwardedPersec", (value))
}

// GetPendingInterruptsForwardedPersec gets the value of PendingInterruptsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPendingInterruptsForwardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingInterruptsForwardedPersec")
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

// SetPendingInterruptsForwardingCost sets the value of PendingInterruptsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPendingInterruptsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsForwardingCost", (value))
}

// GetPendingInterruptsForwardingCost gets the value of PendingInterruptsForwardingCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPendingInterruptsForwardingCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingInterruptsForwardingCost")
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

// SetPendingInterruptsPersec sets the value of PendingInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPendingInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsPersec", (value))
}

// GetPendingInterruptsPersec gets the value of PendingInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPendingInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingInterruptsPersec")
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

// SetPercentGuestRelativeUtilization sets the value of PercentGuestRelativeUtilization for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentGuestRelativeUtilization(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRelativeUtilization", (value))
}

// GetPercentGuestRelativeUtilization gets the value of PercentGuestRelativeUtilization for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentGuestRelativeUtilization() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentGuestRelativeUtilization")
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

// SetPercentGuestRunTime sets the value of PercentGuestRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentGuestRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRunTime", (value))
}

// GetPercentGuestRunTime gets the value of PercentGuestRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentGuestRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentGuestRunTime")
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

// SetPercentHypervisorRunTime sets the value of PercentHypervisorRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentHypervisorRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentHypervisorRunTime", (value))
}

// GetPercentHypervisorRunTime gets the value of PercentHypervisorRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentHypervisorRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentHypervisorRunTime")
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

// SetPercentRemoteRunTime sets the value of PercentRemoteRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentRemoteRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentRemoteRunTime", (value))
}

// GetPercentRemoteRunTime gets the value of PercentRemoteRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentRemoteRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentRemoteRunTime")
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

// SetPercentTotalCoreRunTime sets the value of PercentTotalCoreRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentTotalCoreRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentTotalCoreRunTime", (value))
}

// GetPercentTotalCoreRunTime gets the value of PercentTotalCoreRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentTotalCoreRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTotalCoreRunTime")
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

// SetPercentTotalRunTime sets the value of PercentTotalRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentTotalRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentTotalRunTime", (value))
}

// GetPercentTotalRunTime gets the value of PercentTotalRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentTotalRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTotalRunTime")
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

// SetPercentVTL1RunTime sets the value of PercentVTL1RunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentVTL1RunTime(value uint64) (err error) {
	return instance.SetProperty("PercentVTL1RunTime", (value))
}

// GetPercentVTL1RunTime gets the value of PercentVTL1RunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentVTL1RunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentVTL1RunTime")
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

// SetPercentVTL2RunTime sets the value of PercentVTL2RunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPercentVTL2RunTime(value uint64) (err error) {
	return instance.SetProperty("PercentVTL2RunTime", (value))
}

// GetPercentVTL2RunTime gets the value of PercentVTL2RunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPercentVTL2RunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentVTL2RunTime")
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

// SetPerformanceMonitoringInterruptsPersec sets the value of PerformanceMonitoringInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPerformanceMonitoringInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringInterruptsPersec", (value))
}

// GetPerformanceMonitoringInterruptsPersec gets the value of PerformanceMonitoringInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPerformanceMonitoringInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PerformanceMonitoringInterruptsPersec")
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

// SetPerformanceMonitoringIPTMSRAccessesPersec sets the value of PerformanceMonitoringIPTMSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPerformanceMonitoringIPTMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringIPTMSRAccessesPersec", (value))
}

// GetPerformanceMonitoringIPTMSRAccessesPersec gets the value of PerformanceMonitoringIPTMSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPerformanceMonitoringIPTMSRAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PerformanceMonitoringIPTMSRAccessesPersec")
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

// SetPerformanceMonitoringLBRMSRAccessesPersec sets the value of PerformanceMonitoringLBRMSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPerformanceMonitoringLBRMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringLBRMSRAccessesPersec", (value))
}

// GetPerformanceMonitoringLBRMSRAccessesPersec gets the value of PerformanceMonitoringLBRMSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPerformanceMonitoringLBRMSRAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PerformanceMonitoringLBRMSRAccessesPersec")
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

// SetPerformanceMonitoringvPMUMSRAccessesPersec sets the value of PerformanceMonitoringvPMUMSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPerformanceMonitoringvPMUMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringvPMUMSRAccessesPersec", (value))
}

// GetPerformanceMonitoringvPMUMSRAccessesPersec gets the value of PerformanceMonitoringvPMUMSRAccessesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPerformanceMonitoringvPMUMSRAccessesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PerformanceMonitoringvPMUMSRAccessesPersec")
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

// SetPostedInterruptNotificationsPersec sets the value of PostedInterruptNotificationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPostedInterruptNotificationsPersec(value uint64) (err error) {
	return instance.SetProperty("PostedInterruptNotificationsPersec", (value))
}

// GetPostedInterruptNotificationsPersec gets the value of PostedInterruptNotificationsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPostedInterruptNotificationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostedInterruptNotificationsPersec")
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

// SetPostedInterruptScansPersec sets the value of PostedInterruptScansPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyPostedInterruptScansPersec(value uint64) (err error) {
	return instance.SetProperty("PostedInterruptScansPersec", (value))
}

// GetPostedInterruptScansPersec gets the value of PostedInterruptScansPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyPostedInterruptScansPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostedInterruptScansPersec")
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

// SetRDPMCInstructionsCost sets the value of RDPMCInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyRDPMCInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("RDPMCInstructionsCost", (value))
}

// GetRDPMCInstructionsCost gets the value of RDPMCInstructionsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyRDPMCInstructionsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("RDPMCInstructionsCost")
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

// SetRDPMCInstructionsPersec sets the value of RDPMCInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyRDPMCInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("RDPMCInstructionsPersec", (value))
}

// GetRDPMCInstructionsPersec gets the value of RDPMCInstructionsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyRDPMCInstructionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RDPMCInstructionsPersec")
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

// SetReflectedGuestPageFaultsPersec sets the value of ReflectedGuestPageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyReflectedGuestPageFaultsPersec(value uint64) (err error) {
	return instance.SetProperty("ReflectedGuestPageFaultsPersec", (value))
}

// GetReflectedGuestPageFaultsPersec gets the value of ReflectedGuestPageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyReflectedGuestPageFaultsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReflectedGuestPageFaultsPersec")
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

// SetSchedulingPriority sets the value of SchedulingPriority for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertySchedulingPriority(value uint64) (err error) {
	return instance.SetProperty("SchedulingPriority", (value))
}

// GetSchedulingPriority gets the value of SchedulingPriority for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertySchedulingPriority() (value uint64, err error) {
	retValue, err := instance.GetProperty("SchedulingPriority")
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

// SetSmallPageTLBFillsPersec sets the value of SmallPageTLBFillsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertySmallPageTLBFillsPersec(value uint64) (err error) {
	return instance.SetProperty("SmallPageTLBFillsPersec", (value))
}

// GetSmallPageTLBFillsPersec gets the value of SmallPageTLBFillsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertySmallPageTLBFillsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SmallPageTLBFillsPersec")
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

// SetSVMHypercallsPersec sets the value of SVMHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertySVMHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("SVMHypercallsPersec", (value))
}

// GetSVMHypercallsPersec gets the value of SVMHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertySVMHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SVMHypercallsPersec")
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

// SetSyntheticInterruptHypercallsPersec sets the value of SyntheticInterruptHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertySyntheticInterruptHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("SyntheticInterruptHypercallsPersec", (value))
}

// GetSyntheticInterruptHypercallsPersec gets the value of SyntheticInterruptHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertySyntheticInterruptHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SyntheticInterruptHypercallsPersec")
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

// SetSyntheticInterruptsPersec sets the value of SyntheticInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertySyntheticInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("SyntheticInterruptsPersec", (value))
}

// GetSyntheticInterruptsPersec gets the value of SyntheticInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertySyntheticInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SyntheticInterruptsPersec")
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

// SetTotalInterceptsCost sets the value of TotalInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyTotalInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("TotalInterceptsCost", (value))
}

// GetTotalInterceptsCost gets the value of TotalInterceptsCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyTotalInterceptsCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInterceptsCost")
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

// SetTotalInterceptsPersec sets the value of TotalInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyTotalInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("TotalInterceptsPersec", (value))
}

// GetTotalInterceptsPersec gets the value of TotalInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyTotalInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInterceptsPersec")
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

// SetTotalMessagesPersec sets the value of TotalMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyTotalMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("TotalMessagesPersec", (value))
}

// GetTotalMessagesPersec gets the value of TotalMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyTotalMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalMessagesPersec")
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

// SetTotalVirtualizationInstructionsEmulatedPersec sets the value of TotalVirtualizationInstructionsEmulatedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyTotalVirtualizationInstructionsEmulatedPersec(value uint64) (err error) {
	return instance.SetProperty("TotalVirtualizationInstructionsEmulatedPersec", (value))
}

// GetTotalVirtualizationInstructionsEmulatedPersec gets the value of TotalVirtualizationInstructionsEmulatedPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyTotalVirtualizationInstructionsEmulatedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalVirtualizationInstructionsEmulatedPersec")
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

// SetTotalVirtualizationInstructionsEmulationCost sets the value of TotalVirtualizationInstructionsEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyTotalVirtualizationInstructionsEmulationCost(value uint64) (err error) {
	return instance.SetProperty("TotalVirtualizationInstructionsEmulationCost", (value))
}

// GetTotalVirtualizationInstructionsEmulationCost gets the value of TotalVirtualizationInstructionsEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyTotalVirtualizationInstructionsEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalVirtualizationInstructionsEmulationCost")
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

// SetVirtualInterruptHypercallsPersec sets the value of VirtualInterruptHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVirtualInterruptHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualInterruptHypercallsPersec", (value))
}

// GetVirtualInterruptHypercallsPersec gets the value of VirtualInterruptHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVirtualInterruptHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualInterruptHypercallsPersec")
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

// SetVirtualInterruptsPersec sets the value of VirtualInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVirtualInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualInterruptsPersec", (value))
}

// GetVirtualInterruptsPersec gets the value of VirtualInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVirtualInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualInterruptsPersec")
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

// SetVirtualMMUHypercallsPersec sets the value of VirtualMMUHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVirtualMMUHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualMMUHypercallsPersec", (value))
}

// GetVirtualMMUHypercallsPersec gets the value of VirtualMMUHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVirtualMMUHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualMMUHypercallsPersec")
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

// SetVirtualProcessorHypercallsPersec sets the value of VirtualProcessorHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVirtualProcessorHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualProcessorHypercallsPersec", (value))
}

// GetVirtualProcessorHypercallsPersec gets the value of VirtualProcessorHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVirtualProcessorHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualProcessorHypercallsPersec")
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

// SetVMCLEAREmulationInterceptsPersec sets the value of VMCLEAREmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMCLEAREmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMCLEAREmulationInterceptsPersec", (value))
}

// GetVMCLEAREmulationInterceptsPersec gets the value of VMCLEAREmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMCLEAREmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMCLEAREmulationInterceptsPersec")
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

// SetVMCLEARInstructionEmulationCost sets the value of VMCLEARInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMCLEARInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMCLEARInstructionEmulationCost", (value))
}

// GetVMCLEARInstructionEmulationCost gets the value of VMCLEARInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMCLEARInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMCLEARInstructionEmulationCost")
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

// SetVMLOADEmulationInterceptsPersec sets the value of VMLOADEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMLOADEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMLOADEmulationInterceptsPersec", (value))
}

// GetVMLOADEmulationInterceptsPersec gets the value of VMLOADEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMLOADEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMLOADEmulationInterceptsPersec")
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

// SetVMLOADInstructionEmulationCost sets the value of VMLOADInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMLOADInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMLOADInstructionEmulationCost", (value))
}

// GetVMLOADInstructionEmulationCost gets the value of VMLOADInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMLOADInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMLOADInstructionEmulationCost")
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

// SetVMPTRLDEmulationInterceptsPersec sets the value of VMPTRLDEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMPTRLDEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMPTRLDEmulationInterceptsPersec", (value))
}

// GetVMPTRLDEmulationInterceptsPersec gets the value of VMPTRLDEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMPTRLDEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMPTRLDEmulationInterceptsPersec")
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

// SetVMPTRLDInstructionEmulationCost sets the value of VMPTRLDInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMPTRLDInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMPTRLDInstructionEmulationCost", (value))
}

// GetVMPTRLDInstructionEmulationCost gets the value of VMPTRLDInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMPTRLDInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMPTRLDInstructionEmulationCost")
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

// SetVMPTRSTEmulationInterceptsPersec sets the value of VMPTRSTEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMPTRSTEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMPTRSTEmulationInterceptsPersec", (value))
}

// GetVMPTRSTEmulationInterceptsPersec gets the value of VMPTRSTEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMPTRSTEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMPTRSTEmulationInterceptsPersec")
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

// SetVMPTRSTInstructionEmulationCost sets the value of VMPTRSTInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMPTRSTInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMPTRSTInstructionEmulationCost", (value))
}

// GetVMPTRSTInstructionEmulationCost gets the value of VMPTRSTInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMPTRSTInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMPTRSTInstructionEmulationCost")
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

// SetVMREADEmulationInterceptsPersec sets the value of VMREADEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMREADEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMREADEmulationInterceptsPersec", (value))
}

// GetVMREADEmulationInterceptsPersec gets the value of VMREADEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMREADEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMREADEmulationInterceptsPersec")
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

// SetVMREADInstructionEmulationCost sets the value of VMREADInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMREADInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMREADInstructionEmulationCost", (value))
}

// GetVMREADInstructionEmulationCost gets the value of VMREADInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMREADInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMREADInstructionEmulationCost")
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

// SetVMSAVEEmulationInterceptsPersec sets the value of VMSAVEEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMSAVEEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMSAVEEmulationInterceptsPersec", (value))
}

// GetVMSAVEEmulationInterceptsPersec gets the value of VMSAVEEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMSAVEEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMSAVEEmulationInterceptsPersec")
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

// SetVMSAVEInstructionEmulationCost sets the value of VMSAVEInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMSAVEInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMSAVEInstructionEmulationCost", (value))
}

// GetVMSAVEInstructionEmulationCost gets the value of VMSAVEInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMSAVEInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMSAVEInstructionEmulationCost")
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

// SetVMWRITEEmulationInterceptsPersec sets the value of VMWRITEEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMWRITEEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMWRITEEmulationInterceptsPersec", (value))
}

// GetVMWRITEEmulationInterceptsPersec gets the value of VMWRITEEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMWRITEEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMWRITEEmulationInterceptsPersec")
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

// SetVMWRITEInstructionEmulationCost sets the value of VMWRITEInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMWRITEInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMWRITEInstructionEmulationCost", (value))
}

// GetVMWRITEInstructionEmulationCost gets the value of VMWRITEInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMWRITEInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMWRITEInstructionEmulationCost")
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

// SetVMXOFFEmulationInterceptsPersec sets the value of VMXOFFEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMXOFFEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMXOFFEmulationInterceptsPersec", (value))
}

// GetVMXOFFEmulationInterceptsPersec gets the value of VMXOFFEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMXOFFEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMXOFFEmulationInterceptsPersec")
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

// SetVMXOFFInstructionEmulationCost sets the value of VMXOFFInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMXOFFInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMXOFFInstructionEmulationCost", (value))
}

// GetVMXOFFInstructionEmulationCost gets the value of VMXOFFInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMXOFFInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMXOFFInstructionEmulationCost")
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

// SetVMXONEmulationInterceptsPersec sets the value of VMXONEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMXONEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMXONEmulationInterceptsPersec", (value))
}

// GetVMXONEmulationInterceptsPersec gets the value of VMXONEmulationInterceptsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMXONEmulationInterceptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMXONEmulationInterceptsPersec")
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

// SetVMXONInstructionEmulationCost sets the value of VMXONInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVMXONInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMXONInstructionEmulationCost", (value))
}

// GetVMXONInstructionEmulationCost gets the value of VMXONInstructionEmulationCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVMXONInstructionEmulationCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMXONInstructionEmulationCost")
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

// SetVSMHypercallsPersec sets the value of VSMHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVSMHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VSMHypercallsPersec", (value))
}

// GetVSMHypercallsPersec gets the value of VSMHypercallsPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVSMHypercallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VSMHypercallsPersec")
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

// SetVTL1AverageRunTime sets the value of VTL1AverageRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVTL1AverageRunTime(value uint64) (err error) {
	return instance.SetProperty("VTL1AverageRunTime", (value))
}

// GetVTL1AverageRunTime gets the value of VTL1AverageRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVTL1AverageRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("VTL1AverageRunTime")
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

// SetVTL1DispatchesPersec sets the value of VTL1DispatchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVTL1DispatchesPersec(value uint64) (err error) {
	return instance.SetProperty("VTL1DispatchesPersec", (value))
}

// GetVTL1DispatchesPersec gets the value of VTL1DispatchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVTL1DispatchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VTL1DispatchesPersec")
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

// SetVTL2AverageRunTime sets the value of VTL2AverageRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVTL2AverageRunTime(value uint64) (err error) {
	return instance.SetProperty("VTL2AverageRunTime", (value))
}

// GetVTL2AverageRunTime gets the value of VTL2AverageRunTime for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVTL2AverageRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("VTL2AverageRunTime")
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

// SetVTL2DispatchesPersec sets the value of VTL2DispatchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) SetPropertyVTL2DispatchesPersec(value uint64) (err error) {
	return instance.SetProperty("VTL2DispatchesPersec", (value))
}

// GetVTL2DispatchesPersec gets the value of VTL2DispatchesPersec for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisorVirtualProcessor) GetPropertyVTL2DispatchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VTL2DispatchesPersec")
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
