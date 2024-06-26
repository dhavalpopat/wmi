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

// Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor struct
type Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor struct {
	*Win32_PerfRawData

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
	ControlRegisterAccessesCost_Base uint64

	//
	ControlRegisterAccessesForwardedPersec uint64

	//
	ControlRegisterAccessesForwardingCost uint64

	//
	ControlRegisterAccessesForwardingCost_Base uint64

	//
	ControlRegisterAccessesPersec uint64

	//
	CPPCRequestContextSwitchesPersec uint64

	//
	CPUGroupHypercallsPersec uint64

	//
	CPUIDInstructionsCost uint64

	//
	CPUIDInstructionsCost_Base uint64

	//
	CPUIDInstructionsForwardedPersec uint64

	//
	CPUIDInstructionsForwardingCost uint64

	//
	CPUIDInstructionsForwardingCost_Base uint64

	//
	CPUIDInstructionsPersec uint64

	//
	CPUWaitTimePerDispatch uint64

	//
	CPUWaitTimePerDispatch_Base uint64

	//
	DebugRegisterAccessesCost uint64

	//
	DebugRegisterAccessesCost_Base uint64

	//
	DebugRegisterAccessesForwardedPersec uint64

	//
	DebugRegisterAccessesForwardingCost uint64

	//
	DebugRegisterAccessesForwardingCost_Base uint64

	//
	DebugRegisterAccessesPersec uint64

	//
	DepositHypercallsPersec uint64

	//
	DeviceDomainHypercallsPersec uint64

	//
	EmulatedInstructionsCost uint64

	//
	EmulatedInstructionsCost_Base uint64

	//
	EmulatedInstructionsForwardedPersec uint64

	//
	EmulatedInstructionsForwardingCost uint64

	//
	EmulatedInstructionsForwardingCost_Base uint64

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
	ExternalInterruptsCost_Base uint64

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
	GIFInstructionEmulationCost_Base uint64

	//
	GIFInstructionEmulationInterceptsPersec uint64

	//
	GlobalGVARangeFlushesPersec uint64

	//
	GlobalIOTLBFlushCost uint64

	//
	GlobalIOTLBFlushCost_Base uint64

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
	HLTInstructionsCost_Base uint64

	//
	HLTInstructionsForwardedPersec uint64

	//
	HLTInstructionsForwardingCost uint64

	//
	HLTInstructionsForwardingCost_Base uint64

	//
	HLTInstructionsPersec uint64

	//
	HypercallsCost uint64

	//
	HypercallsCost_Base uint64

	//
	HypercallsForwardedPersec uint64

	//
	HypercallsForwardingCost uint64

	//
	HypercallsForwardingCost_Base uint64

	//
	HypercallsPersec uint64

	//
	InvEptAllContextEmulationInterceptsPersec uint64

	//
	InvEptAllContextInstructionEmulationCost uint64

	//
	InvEptAllContextInstructionEmulationCost_Base uint64

	//
	InvEptSingleContextEmulationInterceptsPersec uint64

	//
	InvEptSingleContextInstructionEmulationCost uint64

	//
	InvEptSingleContextInstructionEmulationCost_Base uint64

	//
	InvVpidAllContextEmulationInterceptsPersec uint64

	//
	InvVpidAllContextInstructionEmulationCost uint64

	//
	InvVpidAllContextInstructionEmulationCost_Base uint64

	//
	InvVpidSingleAddressEmulationInterceptsPersec uint64

	//
	InvVpidSingleAddressInstructionEmulationCost uint64

	//
	InvVpidSingleAddressInstructionEmulationCost_Base uint64

	//
	InvVpidSingleContextEmulationInterceptsPersec uint64

	//
	InvVpidSingleContextInstructionEmulationCost uint64

	//
	InvVpidSingleContextInstructionEmulationCost_Base uint64

	//
	IOInstructionsCost uint64

	//
	IOInstructionsCost_Base uint64

	//
	IOInstructionsForwardedPersec uint64

	//
	IOInstructionsForwardingCost uint64

	//
	IOInstructionsForwardingCost_Base uint64

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
	LocalIOTLBFlushCost_Base uint64

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
	MSRAccessesCost_Base uint64

	//
	MSRAccessesForwardedPersec uint64

	//
	MSRAccessesForwardingCost uint64

	//
	MSRAccessesForwardingCost_Base uint64

	//
	MSRAccessesPersec uint64

	//
	MWAITInstructionsCost uint64

	//
	MWAITInstructionsCost_Base uint64

	//
	MWAITInstructionsForwardedPersec uint64

	//
	MWAITInstructionsForwardingCost uint64

	//
	MWAITInstructionsForwardingCost_Base uint64

	//
	MWAITInstructionsPersec uint64

	//
	NestedPageFaultInterceptsCost uint64

	//
	NestedPageFaultInterceptsCost_Base uint64

	//
	NestedPageFaultInterceptsPersec uint64

	//
	NestedSLATHardPageFaultsCost uint64

	//
	NestedSLATHardPageFaultsCost_Base uint64

	//
	NestedSLATHardPageFaultsPersec uint64

	//
	NestedSLATSoftPageFaultsCost uint64

	//
	NestedSLATSoftPageFaultsCost_Base uint64

	//
	NestedSLATSoftPageFaultsPersec uint64

	//
	NestedTLBPageTableEvictionsPersec uint64

	//
	NestedTLBPageTableReclamationsPersec uint64

	//
	NestedVMEntriesCost uint64

	//
	NestedVMEntriesCost_Base uint64

	//
	NestedVMEntriesPersec uint64

	//
	OtherHypercallsPersec uint64

	//
	OtherInterceptsCost uint64

	//
	OtherInterceptsCost_Base uint64

	//
	OtherInterceptsForwardedPersec uint64

	//
	OtherInterceptsForwardingCost uint64

	//
	OtherInterceptsForwardingCost_Base uint64

	//
	OtherInterceptsPersec uint64

	//
	OtherMessagesPersec uint64

	//
	OtherReflectedGuestExceptionsPersec uint64

	//
	PageFaultInterceptsCost uint64

	//
	PageFaultInterceptsCost_Base uint64

	//
	PageFaultInterceptsForwardedPersec uint64

	//
	PageFaultInterceptsForwardingCost uint64

	//
	PageFaultInterceptsForwardingCost_Base uint64

	//
	PageFaultInterceptsPersec uint64

	//
	PageInvalidationsCost uint64

	//
	PageInvalidationsCost_Base uint64

	//
	PageInvalidationsForwardedPersec uint64

	//
	PageInvalidationsForwardingCost uint64

	//
	PageInvalidationsForwardingCost_Base uint64

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
	PendingInterruptsCost_Base uint64

	//
	PendingInterruptsForwardedPersec uint64

	//
	PendingInterruptsForwardingCost uint64

	//
	PendingInterruptsForwardingCost_Base uint64

	//
	PendingInterruptsPersec uint64

	//
	PercentGuestRelativeUtilization uint64

	//
	PercentGuestRelativeUtilization_Base uint64

	//
	PercentGuestRunTime uint64

	//
	PercentGuestRunTime_Base uint64

	//
	PercentHypervisorRunTime uint64

	//
	PercentHypervisorRunTime_Base uint64

	//
	PercentRemoteRunTime uint64

	//
	PercentRemoteRunTime_Base uint64

	//
	PercentTotalCoreRunTime uint64

	//
	PercentTotalCoreRunTime_Base uint64

	//
	PercentTotalRunTime uint64

	//
	PercentTotalRunTime_Base uint64

	//
	PercentVTL1RunTime uint64

	//
	PercentVTL1RunTime_Base uint64

	//
	PercentVTL2RunTime uint64

	//
	PercentVTL2RunTime_Base uint64

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
	RDPMCInstructionsCost_Base uint64

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
	TotalInterceptsCost_Base uint64

	//
	TotalInterceptsPersec uint64

	//
	TotalMessagesPersec uint64

	//
	TotalVirtualizationInstructionsEmulatedPersec uint64

	//
	TotalVirtualizationInstructionsEmulationCost uint64

	//
	TotalVirtualizationInstructionsEmulationCost_Base uint64

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
	VMCLEARInstructionEmulationCost_Base uint64

	//
	VMLOADEmulationInterceptsPersec uint64

	//
	VMLOADInstructionEmulationCost uint64

	//
	VMLOADInstructionEmulationCost_Base uint64

	//
	VMPTRLDEmulationInterceptsPersec uint64

	//
	VMPTRLDInstructionEmulationCost uint64

	//
	VMPTRLDInstructionEmulationCost_Base uint64

	//
	VMPTRSTEmulationInterceptsPersec uint64

	//
	VMPTRSTInstructionEmulationCost uint64

	//
	VMPTRSTInstructionEmulationCost_Base uint64

	//
	VMREADEmulationInterceptsPersec uint64

	//
	VMREADInstructionEmulationCost uint64

	//
	VMREADInstructionEmulationCost_Base uint64

	//
	VMSAVEEmulationInterceptsPersec uint64

	//
	VMSAVEInstructionEmulationCost uint64

	//
	VMSAVEInstructionEmulationCost_Base uint64

	//
	VMWRITEEmulationInterceptsPersec uint64

	//
	VMWRITEInstructionEmulationCost uint64

	//
	VMWRITEInstructionEmulationCost_Base uint64

	//
	VMXOFFEmulationInterceptsPersec uint64

	//
	VMXOFFInstructionEmulationCost uint64

	//
	VMXOFFInstructionEmulationCost_Base uint64

	//
	VMXONEmulationInterceptsPersec uint64

	//
	VMXONInstructionEmulationCost uint64

	//
	VMXONInstructionEmulationCost_Base uint64

	//
	VSMHypercallsPersec uint64

	//
	VTL1AverageRunTime uint64

	//
	VTL1AverageRunTime_Base uint64

	//
	VTL1DispatchesPersec uint64

	//
	VTL2AverageRunTime uint64

	//
	VTL2AverageRunTime_Base uint64

	//
	VTL2DispatchesPersec uint64
}

func NewWin32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessorEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAddressDomainFlushesPersec sets the value of AddressDomainFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAddressDomainFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("AddressDomainFlushesPersec", (value))
}

// GetAddressDomainFlushesPersec gets the value of AddressDomainFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAddressDomainFlushesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAddressSpaceEvictionsPersec(value uint64) (err error) {
	return instance.SetProperty("AddressSpaceEvictionsPersec", (value))
}

// GetAddressSpaceEvictionsPersec gets the value of AddressSpaceEvictionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAddressSpaceEvictionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAddressSpaceFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("AddressSpaceFlushesPersec", (value))
}

// GetAddressSpaceFlushesPersec gets the value of AddressSpaceFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAddressSpaceFlushesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAddressSpaceSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("AddressSpaceSwitchesPersec", (value))
}

// GetAddressSpaceSwitchesPersec gets the value of AddressSpaceSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAddressSpaceSwitchesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAPICEOIAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("APICEOIAccessesPersec", (value))
}

// GetAPICEOIAccessesPersec gets the value of APICEOIAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAPICEOIAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAPICIPIsSentPersec(value uint64) (err error) {
	return instance.SetProperty("APICIPIsSentPersec", (value))
}

// GetAPICIPIsSentPersec gets the value of APICIPIsSentPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAPICIPIsSentPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAPICMMIOAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("APICMMIOAccessesPersec", (value))
}

// GetAPICMMIOAccessesPersec gets the value of APICMMIOAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAPICMMIOAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAPICSelfIPIsSentPersec(value uint64) (err error) {
	return instance.SetProperty("APICSelfIPIsSentPersec", (value))
}

// GetAPICSelfIPIsSentPersec gets the value of APICSelfIPIsSentPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAPICSelfIPIsSentPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyAPICTPRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("APICTPRAccessesPersec", (value))
}

// GetAPICTPRAccessesPersec gets the value of APICTPRAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyAPICTPRAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyBusLockAcquisitionsPersec(value uint64) (err error) {
	return instance.SetProperty("BusLockAcquisitionsPersec", (value))
}

// GetBusLockAcquisitionsPersec gets the value of BusLockAcquisitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyBusLockAcquisitionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyControlRegisterAccessesCost(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesCost", (value))
}

// GetControlRegisterAccessesCost gets the value of ControlRegisterAccessesCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyControlRegisterAccessesCost() (value uint64, err error) {
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

// SetControlRegisterAccessesCost_Base sets the value of ControlRegisterAccessesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyControlRegisterAccessesCost_Base(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesCost_Base", (value))
}

// GetControlRegisterAccessesCost_Base gets the value of ControlRegisterAccessesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyControlRegisterAccessesCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRegisterAccessesCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyControlRegisterAccessesForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesForwardedPersec", (value))
}

// GetControlRegisterAccessesForwardedPersec gets the value of ControlRegisterAccessesForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyControlRegisterAccessesForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyControlRegisterAccessesForwardingCost(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesForwardingCost", (value))
}

// GetControlRegisterAccessesForwardingCost gets the value of ControlRegisterAccessesForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyControlRegisterAccessesForwardingCost() (value uint64, err error) {
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

// SetControlRegisterAccessesForwardingCost_Base sets the value of ControlRegisterAccessesForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyControlRegisterAccessesForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesForwardingCost_Base", (value))
}

// GetControlRegisterAccessesForwardingCost_Base gets the value of ControlRegisterAccessesForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyControlRegisterAccessesForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRegisterAccessesForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyControlRegisterAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("ControlRegisterAccessesPersec", (value))
}

// GetControlRegisterAccessesPersec gets the value of ControlRegisterAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyControlRegisterAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPPCRequestContextSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("CPPCRequestContextSwitchesPersec", (value))
}

// GetCPPCRequestContextSwitchesPersec gets the value of CPPCRequestContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPPCRequestContextSwitchesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUGroupHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("CPUGroupHypercallsPersec", (value))
}

// GetCPUGroupHypercallsPersec gets the value of CPUGroupHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUGroupHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUIDInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsCost", (value))
}

// GetCPUIDInstructionsCost gets the value of CPUIDInstructionsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUIDInstructionsCost() (value uint64, err error) {
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

// SetCPUIDInstructionsCost_Base sets the value of CPUIDInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUIDInstructionsCost_Base(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsCost_Base", (value))
}

// GetCPUIDInstructionsCost_Base gets the value of CPUIDInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUIDInstructionsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUIDInstructionsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUIDInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsForwardedPersec", (value))
}

// GetCPUIDInstructionsForwardedPersec gets the value of CPUIDInstructionsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUIDInstructionsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUIDInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsForwardingCost", (value))
}

// GetCPUIDInstructionsForwardingCost gets the value of CPUIDInstructionsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUIDInstructionsForwardingCost() (value uint64, err error) {
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

// SetCPUIDInstructionsForwardingCost_Base sets the value of CPUIDInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUIDInstructionsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsForwardingCost_Base", (value))
}

// GetCPUIDInstructionsForwardingCost_Base gets the value of CPUIDInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUIDInstructionsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUIDInstructionsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUIDInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("CPUIDInstructionsPersec", (value))
}

// GetCPUIDInstructionsPersec gets the value of CPUIDInstructionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUIDInstructionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUWaitTimePerDispatch(value uint64) (err error) {
	return instance.SetProperty("CPUWaitTimePerDispatch", (value))
}

// GetCPUWaitTimePerDispatch gets the value of CPUWaitTimePerDispatch for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUWaitTimePerDispatch() (value uint64, err error) {
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

// SetCPUWaitTimePerDispatch_Base sets the value of CPUWaitTimePerDispatch_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyCPUWaitTimePerDispatch_Base(value uint64) (err error) {
	return instance.SetProperty("CPUWaitTimePerDispatch_Base", (value))
}

// GetCPUWaitTimePerDispatch_Base gets the value of CPUWaitTimePerDispatch_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyCPUWaitTimePerDispatch_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUWaitTimePerDispatch_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDebugRegisterAccessesCost(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesCost", (value))
}

// GetDebugRegisterAccessesCost gets the value of DebugRegisterAccessesCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDebugRegisterAccessesCost() (value uint64, err error) {
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

// SetDebugRegisterAccessesCost_Base sets the value of DebugRegisterAccessesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDebugRegisterAccessesCost_Base(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesCost_Base", (value))
}

// GetDebugRegisterAccessesCost_Base gets the value of DebugRegisterAccessesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDebugRegisterAccessesCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("DebugRegisterAccessesCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDebugRegisterAccessesForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesForwardedPersec", (value))
}

// GetDebugRegisterAccessesForwardedPersec gets the value of DebugRegisterAccessesForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDebugRegisterAccessesForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDebugRegisterAccessesForwardingCost(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesForwardingCost", (value))
}

// GetDebugRegisterAccessesForwardingCost gets the value of DebugRegisterAccessesForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDebugRegisterAccessesForwardingCost() (value uint64, err error) {
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

// SetDebugRegisterAccessesForwardingCost_Base sets the value of DebugRegisterAccessesForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDebugRegisterAccessesForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesForwardingCost_Base", (value))
}

// GetDebugRegisterAccessesForwardingCost_Base gets the value of DebugRegisterAccessesForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDebugRegisterAccessesForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("DebugRegisterAccessesForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDebugRegisterAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("DebugRegisterAccessesPersec", (value))
}

// GetDebugRegisterAccessesPersec gets the value of DebugRegisterAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDebugRegisterAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDepositHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("DepositHypercallsPersec", (value))
}

// GetDepositHypercallsPersec gets the value of DepositHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDepositHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyDeviceDomainHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("DeviceDomainHypercallsPersec", (value))
}

// GetDeviceDomainHypercallsPersec gets the value of DeviceDomainHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyDeviceDomainHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEmulatedInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsCost", (value))
}

// GetEmulatedInstructionsCost gets the value of EmulatedInstructionsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEmulatedInstructionsCost() (value uint64, err error) {
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

// SetEmulatedInstructionsCost_Base sets the value of EmulatedInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEmulatedInstructionsCost_Base(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsCost_Base", (value))
}

// GetEmulatedInstructionsCost_Base gets the value of EmulatedInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEmulatedInstructionsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("EmulatedInstructionsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEmulatedInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsForwardedPersec", (value))
}

// GetEmulatedInstructionsForwardedPersec gets the value of EmulatedInstructionsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEmulatedInstructionsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEmulatedInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsForwardingCost", (value))
}

// GetEmulatedInstructionsForwardingCost gets the value of EmulatedInstructionsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEmulatedInstructionsForwardingCost() (value uint64, err error) {
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

// SetEmulatedInstructionsForwardingCost_Base sets the value of EmulatedInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEmulatedInstructionsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsForwardingCost_Base", (value))
}

// GetEmulatedInstructionsForwardingCost_Base gets the value of EmulatedInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEmulatedInstructionsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("EmulatedInstructionsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEmulatedInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("EmulatedInstructionsPersec", (value))
}

// GetEmulatedInstructionsPersec gets the value of EmulatedInstructionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEmulatedInstructionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyEventLogHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("EventLogHypercallsPersec", (value))
}

// GetEventLogHypercallsPersec gets the value of EventLogHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyEventLogHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyExtendedHypercallInterceptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("ExtendedHypercallInterceptMessagesPersec", (value))
}

// GetExtendedHypercallInterceptMessagesPersec gets the value of ExtendedHypercallInterceptMessagesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyExtendedHypercallInterceptMessagesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyExtendedHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("ExtendedHypercallsPersec", (value))
}

// GetExtendedHypercallsPersec gets the value of ExtendedHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyExtendedHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyExternalInterruptsCost(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsCost", (value))
}

// GetExternalInterruptsCost gets the value of ExternalInterruptsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyExternalInterruptsCost() (value uint64, err error) {
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

// SetExternalInterruptsCost_Base sets the value of ExternalInterruptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyExternalInterruptsCost_Base(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsCost_Base", (value))
}

// GetExternalInterruptsCost_Base gets the value of ExternalInterruptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyExternalInterruptsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExternalInterruptsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyExternalInterruptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsForwardedPersec", (value))
}

// GetExternalInterruptsForwardedPersec gets the value of ExternalInterruptsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyExternalInterruptsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyExternalInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("ExternalInterruptsPersec", (value))
}

// GetExternalInterruptsPersec gets the value of ExternalInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyExternalInterruptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyFlushPhysicalAddressListHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("FlushPhysicalAddressListHypercallsPersec", (value))
}

// GetFlushPhysicalAddressListHypercallsPersec gets the value of FlushPhysicalAddressListHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyFlushPhysicalAddressListHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyFlushPhysicalAddressSpaceHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("FlushPhysicalAddressSpaceHypercallsPersec", (value))
}

// GetFlushPhysicalAddressSpaceHypercallsPersec gets the value of FlushPhysicalAddressSpaceHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyFlushPhysicalAddressSpaceHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGIFInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("GIFInstructionEmulationCost", (value))
}

// GetGIFInstructionEmulationCost gets the value of GIFInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGIFInstructionEmulationCost() (value uint64, err error) {
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

// SetGIFInstructionEmulationCost_Base sets the value of GIFInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGIFInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("GIFInstructionEmulationCost_Base", (value))
}

// GetGIFInstructionEmulationCost_Base gets the value of GIFInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGIFInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("GIFInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGIFInstructionEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("GIFInstructionEmulationInterceptsPersec", (value))
}

// GetGIFInstructionEmulationInterceptsPersec gets the value of GIFInstructionEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGIFInstructionEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGlobalGVARangeFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("GlobalGVARangeFlushesPersec", (value))
}

// GetGlobalGVARangeFlushesPersec gets the value of GlobalGVARangeFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGlobalGVARangeFlushesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGlobalIOTLBFlushCost(value uint64) (err error) {
	return instance.SetProperty("GlobalIOTLBFlushCost", (value))
}

// GetGlobalIOTLBFlushCost gets the value of GlobalIOTLBFlushCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGlobalIOTLBFlushCost() (value uint64, err error) {
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

// SetGlobalIOTLBFlushCost_Base sets the value of GlobalIOTLBFlushCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGlobalIOTLBFlushCost_Base(value uint64) (err error) {
	return instance.SetProperty("GlobalIOTLBFlushCost_Base", (value))
}

// GetGlobalIOTLBFlushCost_Base gets the value of GlobalIOTLBFlushCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGlobalIOTLBFlushCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("GlobalIOTLBFlushCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGlobalIOTLBFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("GlobalIOTLBFlushesPersec", (value))
}

// GetGlobalIOTLBFlushesPersec gets the value of GlobalIOTLBFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGlobalIOTLBFlushesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGPASpaceHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("GPASpaceHypercallsPersec", (value))
}

// GetGPASpaceHypercallsPersec gets the value of GPASpaceHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGPASpaceHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyGuestPageTableMapsPersec(value uint64) (err error) {
	return instance.SetProperty("GuestPageTableMapsPersec", (value))
}

// GetGuestPageTableMapsPersec gets the value of GuestPageTableMapsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyGuestPageTableMapsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHardwareInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("HardwareInterruptsPersec", (value))
}

// GetHardwareInterruptsPersec gets the value of HardwareInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHardwareInterruptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHLTInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsCost", (value))
}

// GetHLTInstructionsCost gets the value of HLTInstructionsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHLTInstructionsCost() (value uint64, err error) {
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

// SetHLTInstructionsCost_Base sets the value of HLTInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHLTInstructionsCost_Base(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsCost_Base", (value))
}

// GetHLTInstructionsCost_Base gets the value of HLTInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHLTInstructionsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("HLTInstructionsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHLTInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsForwardedPersec", (value))
}

// GetHLTInstructionsForwardedPersec gets the value of HLTInstructionsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHLTInstructionsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHLTInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsForwardingCost", (value))
}

// GetHLTInstructionsForwardingCost gets the value of HLTInstructionsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHLTInstructionsForwardingCost() (value uint64, err error) {
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

// SetHLTInstructionsForwardingCost_Base sets the value of HLTInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHLTInstructionsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsForwardingCost_Base", (value))
}

// GetHLTInstructionsForwardingCost_Base gets the value of HLTInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHLTInstructionsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("HLTInstructionsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHLTInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("HLTInstructionsPersec", (value))
}

// GetHLTInstructionsPersec gets the value of HLTInstructionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHLTInstructionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHypercallsCost(value uint64) (err error) {
	return instance.SetProperty("HypercallsCost", (value))
}

// GetHypercallsCost gets the value of HypercallsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHypercallsCost() (value uint64, err error) {
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

// SetHypercallsCost_Base sets the value of HypercallsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHypercallsCost_Base(value uint64) (err error) {
	return instance.SetProperty("HypercallsCost_Base", (value))
}

// GetHypercallsCost_Base gets the value of HypercallsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHypercallsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypercallsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHypercallsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("HypercallsForwardedPersec", (value))
}

// GetHypercallsForwardedPersec gets the value of HypercallsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHypercallsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHypercallsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("HypercallsForwardingCost", (value))
}

// GetHypercallsForwardingCost gets the value of HypercallsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHypercallsForwardingCost() (value uint64, err error) {
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

// SetHypercallsForwardingCost_Base sets the value of HypercallsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHypercallsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("HypercallsForwardingCost_Base", (value))
}

// GetHypercallsForwardingCost_Base gets the value of HypercallsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHypercallsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypercallsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("HypercallsPersec", (value))
}

// GetHypercallsPersec gets the value of HypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvEptAllContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvEptAllContextEmulationInterceptsPersec", (value))
}

// GetInvEptAllContextEmulationInterceptsPersec gets the value of InvEptAllContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvEptAllContextEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvEptAllContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvEptAllContextInstructionEmulationCost", (value))
}

// GetInvEptAllContextInstructionEmulationCost gets the value of InvEptAllContextInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvEptAllContextInstructionEmulationCost() (value uint64, err error) {
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

// SetInvEptAllContextInstructionEmulationCost_Base sets the value of InvEptAllContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvEptAllContextInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("InvEptAllContextInstructionEmulationCost_Base", (value))
}

// GetInvEptAllContextInstructionEmulationCost_Base gets the value of InvEptAllContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvEptAllContextInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvEptAllContextInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvEptSingleContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvEptSingleContextEmulationInterceptsPersec", (value))
}

// GetInvEptSingleContextEmulationInterceptsPersec gets the value of InvEptSingleContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvEptSingleContextEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvEptSingleContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvEptSingleContextInstructionEmulationCost", (value))
}

// GetInvEptSingleContextInstructionEmulationCost gets the value of InvEptSingleContextInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvEptSingleContextInstructionEmulationCost() (value uint64, err error) {
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

// SetInvEptSingleContextInstructionEmulationCost_Base sets the value of InvEptSingleContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvEptSingleContextInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("InvEptSingleContextInstructionEmulationCost_Base", (value))
}

// GetInvEptSingleContextInstructionEmulationCost_Base gets the value of InvEptSingleContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvEptSingleContextInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvEptSingleContextInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidAllContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvVpidAllContextEmulationInterceptsPersec", (value))
}

// GetInvVpidAllContextEmulationInterceptsPersec gets the value of InvVpidAllContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidAllContextEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidAllContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvVpidAllContextInstructionEmulationCost", (value))
}

// GetInvVpidAllContextInstructionEmulationCost gets the value of InvVpidAllContextInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidAllContextInstructionEmulationCost() (value uint64, err error) {
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

// SetInvVpidAllContextInstructionEmulationCost_Base sets the value of InvVpidAllContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidAllContextInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("InvVpidAllContextInstructionEmulationCost_Base", (value))
}

// GetInvVpidAllContextInstructionEmulationCost_Base gets the value of InvVpidAllContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidAllContextInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidAllContextInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidSingleAddressEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleAddressEmulationInterceptsPersec", (value))
}

// GetInvVpidSingleAddressEmulationInterceptsPersec gets the value of InvVpidSingleAddressEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidSingleAddressEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidSingleAddressInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleAddressInstructionEmulationCost", (value))
}

// GetInvVpidSingleAddressInstructionEmulationCost gets the value of InvVpidSingleAddressInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidSingleAddressInstructionEmulationCost() (value uint64, err error) {
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

// SetInvVpidSingleAddressInstructionEmulationCost_Base sets the value of InvVpidSingleAddressInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidSingleAddressInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleAddressInstructionEmulationCost_Base", (value))
}

// GetInvVpidSingleAddressInstructionEmulationCost_Base gets the value of InvVpidSingleAddressInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidSingleAddressInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidSingleAddressInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidSingleContextEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleContextEmulationInterceptsPersec", (value))
}

// GetInvVpidSingleContextEmulationInterceptsPersec gets the value of InvVpidSingleContextEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidSingleContextEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidSingleContextInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleContextInstructionEmulationCost", (value))
}

// GetInvVpidSingleContextInstructionEmulationCost gets the value of InvVpidSingleContextInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidSingleContextInstructionEmulationCost() (value uint64, err error) {
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

// SetInvVpidSingleContextInstructionEmulationCost_Base sets the value of InvVpidSingleContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyInvVpidSingleContextInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("InvVpidSingleContextInstructionEmulationCost_Base", (value))
}

// GetInvVpidSingleContextInstructionEmulationCost_Base gets the value of InvVpidSingleContextInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyInvVpidSingleContextInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvVpidSingleContextInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsCost", (value))
}

// GetIOInstructionsCost gets the value of IOInstructionsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInstructionsCost() (value uint64, err error) {
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

// SetIOInstructionsCost_Base sets the value of IOInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInstructionsCost_Base(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsCost_Base", (value))
}

// GetIOInstructionsCost_Base gets the value of IOInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInstructionsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInstructionsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsForwardedPersec", (value))
}

// GetIOInstructionsForwardedPersec gets the value of IOInstructionsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInstructionsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsForwardingCost", (value))
}

// GetIOInstructionsForwardingCost gets the value of IOInstructionsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInstructionsForwardingCost() (value uint64, err error) {
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

// SetIOInstructionsForwardingCost_Base sets the value of IOInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInstructionsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsForwardingCost_Base", (value))
}

// GetIOInstructionsForwardingCost_Base gets the value of IOInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInstructionsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOInstructionsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("IOInstructionsPersec", (value))
}

// GetIOInstructionsPersec gets the value of IOInstructionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInstructionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOInterceptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("IOInterceptMessagesPersec", (value))
}

// GetIOInterceptMessagesPersec gets the value of IOInterceptMessagesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOInterceptMessagesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyIOMMUHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("IOMMUHypercallsPersec", (value))
}

// GetIOMMUHypercallsPersec gets the value of IOMMUHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyIOMMUHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLargePageTLBFillsPersec(value uint64) (err error) {
	return instance.SetProperty("LargePageTLBFillsPersec", (value))
}

// GetLargePageTLBFillsPersec gets the value of LargePageTLBFillsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLargePageTLBFillsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLocalFlushedGVARangesPersec(value uint64) (err error) {
	return instance.SetProperty("LocalFlushedGVARangesPersec", (value))
}

// GetLocalFlushedGVARangesPersec gets the value of LocalFlushedGVARangesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLocalFlushedGVARangesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLocalIOTLBFlushCost(value uint64) (err error) {
	return instance.SetProperty("LocalIOTLBFlushCost", (value))
}

// GetLocalIOTLBFlushCost gets the value of LocalIOTLBFlushCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLocalIOTLBFlushCost() (value uint64, err error) {
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

// SetLocalIOTLBFlushCost_Base sets the value of LocalIOTLBFlushCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLocalIOTLBFlushCost_Base(value uint64) (err error) {
	return instance.SetProperty("LocalIOTLBFlushCost_Base", (value))
}

// GetLocalIOTLBFlushCost_Base gets the value of LocalIOTLBFlushCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLocalIOTLBFlushCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalIOTLBFlushCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLocalIOTLBFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("LocalIOTLBFlushesPersec", (value))
}

// GetLocalIOTLBFlushesPersec gets the value of LocalIOTLBFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLocalIOTLBFlushesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLogicalProcessorDispatchesPersec(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessorDispatchesPersec", (value))
}

// GetLogicalProcessorDispatchesPersec gets the value of LogicalProcessorDispatchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLogicalProcessorDispatchesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLogicalProcessorHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessorHypercallsPersec", (value))
}

// GetLogicalProcessorHypercallsPersec gets the value of LogicalProcessorHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLogicalProcessorHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLogicalProcessorMigrationsPersec(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessorMigrationsPersec", (value))
}

// GetLogicalProcessorMigrationsPersec gets the value of LogicalProcessorMigrationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLogicalProcessorMigrationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyLongSpinWaitHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("LongSpinWaitHypercallsPersec", (value))
}

// GetLongSpinWaitHypercallsPersec gets the value of LongSpinWaitHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyLongSpinWaitHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMBECNestedPageTableSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("MBECNestedPageTableSwitchesPersec", (value))
}

// GetMBECNestedPageTableSwitchesPersec gets the value of MBECNestedPageTableSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMBECNestedPageTableSwitchesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMemoryInterceptMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("MemoryInterceptMessagesPersec", (value))
}

// GetMemoryInterceptMessagesPersec gets the value of MemoryInterceptMessagesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMemoryInterceptMessagesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMSRAccessesCost(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesCost", (value))
}

// GetMSRAccessesCost gets the value of MSRAccessesCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMSRAccessesCost() (value uint64, err error) {
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

// SetMSRAccessesCost_Base sets the value of MSRAccessesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMSRAccessesCost_Base(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesCost_Base", (value))
}

// GetMSRAccessesCost_Base gets the value of MSRAccessesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMSRAccessesCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSRAccessesCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMSRAccessesForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesForwardedPersec", (value))
}

// GetMSRAccessesForwardedPersec gets the value of MSRAccessesForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMSRAccessesForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMSRAccessesForwardingCost(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesForwardingCost", (value))
}

// GetMSRAccessesForwardingCost gets the value of MSRAccessesForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMSRAccessesForwardingCost() (value uint64, err error) {
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

// SetMSRAccessesForwardingCost_Base sets the value of MSRAccessesForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMSRAccessesForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesForwardingCost_Base", (value))
}

// GetMSRAccessesForwardingCost_Base gets the value of MSRAccessesForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMSRAccessesForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSRAccessesForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("MSRAccessesPersec", (value))
}

// GetMSRAccessesPersec gets the value of MSRAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMSRAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMWAITInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsCost", (value))
}

// GetMWAITInstructionsCost gets the value of MWAITInstructionsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMWAITInstructionsCost() (value uint64, err error) {
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

// SetMWAITInstructionsCost_Base sets the value of MWAITInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMWAITInstructionsCost_Base(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsCost_Base", (value))
}

// GetMWAITInstructionsCost_Base gets the value of MWAITInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMWAITInstructionsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("MWAITInstructionsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMWAITInstructionsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsForwardedPersec", (value))
}

// GetMWAITInstructionsForwardedPersec gets the value of MWAITInstructionsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMWAITInstructionsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMWAITInstructionsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsForwardingCost", (value))
}

// GetMWAITInstructionsForwardingCost gets the value of MWAITInstructionsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMWAITInstructionsForwardingCost() (value uint64, err error) {
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

// SetMWAITInstructionsForwardingCost_Base sets the value of MWAITInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMWAITInstructionsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsForwardingCost_Base", (value))
}

// GetMWAITInstructionsForwardingCost_Base gets the value of MWAITInstructionsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMWAITInstructionsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("MWAITInstructionsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyMWAITInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("MWAITInstructionsPersec", (value))
}

// GetMWAITInstructionsPersec gets the value of MWAITInstructionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyMWAITInstructionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedPageFaultInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("NestedPageFaultInterceptsCost", (value))
}

// GetNestedPageFaultInterceptsCost gets the value of NestedPageFaultInterceptsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedPageFaultInterceptsCost() (value uint64, err error) {
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

// SetNestedPageFaultInterceptsCost_Base sets the value of NestedPageFaultInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedPageFaultInterceptsCost_Base(value uint64) (err error) {
	return instance.SetProperty("NestedPageFaultInterceptsCost_Base", (value))
}

// GetNestedPageFaultInterceptsCost_Base gets the value of NestedPageFaultInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedPageFaultInterceptsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedPageFaultInterceptsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedPageFaultInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedPageFaultInterceptsPersec", (value))
}

// GetNestedPageFaultInterceptsPersec gets the value of NestedPageFaultInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedPageFaultInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedSLATHardPageFaultsCost(value uint64) (err error) {
	return instance.SetProperty("NestedSLATHardPageFaultsCost", (value))
}

// GetNestedSLATHardPageFaultsCost gets the value of NestedSLATHardPageFaultsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedSLATHardPageFaultsCost() (value uint64, err error) {
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

// SetNestedSLATHardPageFaultsCost_Base sets the value of NestedSLATHardPageFaultsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedSLATHardPageFaultsCost_Base(value uint64) (err error) {
	return instance.SetProperty("NestedSLATHardPageFaultsCost_Base", (value))
}

// GetNestedSLATHardPageFaultsCost_Base gets the value of NestedSLATHardPageFaultsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedSLATHardPageFaultsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedSLATHardPageFaultsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedSLATHardPageFaultsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedSLATHardPageFaultsPersec", (value))
}

// GetNestedSLATHardPageFaultsPersec gets the value of NestedSLATHardPageFaultsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedSLATHardPageFaultsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedSLATSoftPageFaultsCost(value uint64) (err error) {
	return instance.SetProperty("NestedSLATSoftPageFaultsCost", (value))
}

// GetNestedSLATSoftPageFaultsCost gets the value of NestedSLATSoftPageFaultsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedSLATSoftPageFaultsCost() (value uint64, err error) {
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

// SetNestedSLATSoftPageFaultsCost_Base sets the value of NestedSLATSoftPageFaultsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedSLATSoftPageFaultsCost_Base(value uint64) (err error) {
	return instance.SetProperty("NestedSLATSoftPageFaultsCost_Base", (value))
}

// GetNestedSLATSoftPageFaultsCost_Base gets the value of NestedSLATSoftPageFaultsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedSLATSoftPageFaultsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedSLATSoftPageFaultsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedSLATSoftPageFaultsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedSLATSoftPageFaultsPersec", (value))
}

// GetNestedSLATSoftPageFaultsPersec gets the value of NestedSLATSoftPageFaultsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedSLATSoftPageFaultsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedTLBPageTableEvictionsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedTLBPageTableEvictionsPersec", (value))
}

// GetNestedTLBPageTableEvictionsPersec gets the value of NestedTLBPageTableEvictionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedTLBPageTableEvictionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedTLBPageTableReclamationsPersec(value uint64) (err error) {
	return instance.SetProperty("NestedTLBPageTableReclamationsPersec", (value))
}

// GetNestedTLBPageTableReclamationsPersec gets the value of NestedTLBPageTableReclamationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedTLBPageTableReclamationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedVMEntriesCost(value uint64) (err error) {
	return instance.SetProperty("NestedVMEntriesCost", (value))
}

// GetNestedVMEntriesCost gets the value of NestedVMEntriesCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedVMEntriesCost() (value uint64, err error) {
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

// SetNestedVMEntriesCost_Base sets the value of NestedVMEntriesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedVMEntriesCost_Base(value uint64) (err error) {
	return instance.SetProperty("NestedVMEntriesCost_Base", (value))
}

// GetNestedVMEntriesCost_Base gets the value of NestedVMEntriesCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedVMEntriesCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("NestedVMEntriesCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyNestedVMEntriesPersec(value uint64) (err error) {
	return instance.SetProperty("NestedVMEntriesPersec", (value))
}

// GetNestedVMEntriesPersec gets the value of NestedVMEntriesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyNestedVMEntriesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("OtherHypercallsPersec", (value))
}

// GetOtherHypercallsPersec gets the value of OtherHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsCost", (value))
}

// GetOtherInterceptsCost gets the value of OtherInterceptsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherInterceptsCost() (value uint64, err error) {
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

// SetOtherInterceptsCost_Base sets the value of OtherInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherInterceptsCost_Base(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsCost_Base", (value))
}

// GetOtherInterceptsCost_Base gets the value of OtherInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherInterceptsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherInterceptsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherInterceptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsForwardedPersec", (value))
}

// GetOtherInterceptsForwardedPersec gets the value of OtherInterceptsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherInterceptsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherInterceptsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsForwardingCost", (value))
}

// GetOtherInterceptsForwardingCost gets the value of OtherInterceptsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherInterceptsForwardingCost() (value uint64, err error) {
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

// SetOtherInterceptsForwardingCost_Base sets the value of OtherInterceptsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherInterceptsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsForwardingCost_Base", (value))
}

// GetOtherInterceptsForwardingCost_Base gets the value of OtherInterceptsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherInterceptsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherInterceptsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("OtherInterceptsPersec", (value))
}

// GetOtherInterceptsPersec gets the value of OtherInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("OtherMessagesPersec", (value))
}

// GetOtherMessagesPersec gets the value of OtherMessagesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherMessagesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyOtherReflectedGuestExceptionsPersec(value uint64) (err error) {
	return instance.SetProperty("OtherReflectedGuestExceptionsPersec", (value))
}

// GetOtherReflectedGuestExceptionsPersec gets the value of OtherReflectedGuestExceptionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyOtherReflectedGuestExceptionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageFaultInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsCost", (value))
}

// GetPageFaultInterceptsCost gets the value of PageFaultInterceptsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageFaultInterceptsCost() (value uint64, err error) {
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

// SetPageFaultInterceptsCost_Base sets the value of PageFaultInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageFaultInterceptsCost_Base(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsCost_Base", (value))
}

// GetPageFaultInterceptsCost_Base gets the value of PageFaultInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageFaultInterceptsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFaultInterceptsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageFaultInterceptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsForwardedPersec", (value))
}

// GetPageFaultInterceptsForwardedPersec gets the value of PageFaultInterceptsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageFaultInterceptsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageFaultInterceptsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsForwardingCost", (value))
}

// GetPageFaultInterceptsForwardingCost gets the value of PageFaultInterceptsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageFaultInterceptsForwardingCost() (value uint64, err error) {
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

// SetPageFaultInterceptsForwardingCost_Base sets the value of PageFaultInterceptsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageFaultInterceptsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsForwardingCost_Base", (value))
}

// GetPageFaultInterceptsForwardingCost_Base gets the value of PageFaultInterceptsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageFaultInterceptsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFaultInterceptsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageFaultInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("PageFaultInterceptsPersec", (value))
}

// GetPageFaultInterceptsPersec gets the value of PageFaultInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageFaultInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageInvalidationsCost(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsCost", (value))
}

// GetPageInvalidationsCost gets the value of PageInvalidationsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageInvalidationsCost() (value uint64, err error) {
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

// SetPageInvalidationsCost_Base sets the value of PageInvalidationsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageInvalidationsCost_Base(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsCost_Base", (value))
}

// GetPageInvalidationsCost_Base gets the value of PageInvalidationsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageInvalidationsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageInvalidationsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageInvalidationsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsForwardedPersec", (value))
}

// GetPageInvalidationsForwardedPersec gets the value of PageInvalidationsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageInvalidationsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageInvalidationsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsForwardingCost", (value))
}

// GetPageInvalidationsForwardingCost gets the value of PageInvalidationsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageInvalidationsForwardingCost() (value uint64, err error) {
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

// SetPageInvalidationsForwardingCost_Base sets the value of PageInvalidationsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageInvalidationsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsForwardingCost_Base", (value))
}

// GetPageInvalidationsForwardingCost_Base gets the value of PageInvalidationsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageInvalidationsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageInvalidationsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageInvalidationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageInvalidationsPersec", (value))
}

// GetPageInvalidationsPersec gets the value of PageInvalidationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageInvalidationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageScansPersec(value uint64) (err error) {
	return instance.SetProperty("PageScansPersec", (value))
}

// GetPageScansPersec gets the value of PageScansPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageScansPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageTableAllocationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableAllocationsPersec", (value))
}

// GetPageTableAllocationsPersec gets the value of PageTableAllocationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageTableAllocationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageTableEvictionsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableEvictionsPersec", (value))
}

// GetPageTableEvictionsPersec gets the value of PageTableEvictionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageTableEvictionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageTableReclamationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableReclamationsPersec", (value))
}

// GetPageTableReclamationsPersec gets the value of PageTableReclamationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageTableReclamationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageTableResetsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableResetsPersec", (value))
}

// GetPageTableResetsPersec gets the value of PageTableResetsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageTableResetsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageTableValidationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableValidationsPersec", (value))
}

// GetPageTableValidationsPersec gets the value of PageTableValidationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageTableValidationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPageTableWriteInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("PageTableWriteInterceptsPersec", (value))
}

// GetPageTableWriteInterceptsPersec gets the value of PageTableWriteInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPageTableWriteInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPendingInterruptsCost(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsCost", (value))
}

// GetPendingInterruptsCost gets the value of PendingInterruptsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPendingInterruptsCost() (value uint64, err error) {
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

// SetPendingInterruptsCost_Base sets the value of PendingInterruptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPendingInterruptsCost_Base(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsCost_Base", (value))
}

// GetPendingInterruptsCost_Base gets the value of PendingInterruptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPendingInterruptsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingInterruptsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPendingInterruptsForwardedPersec(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsForwardedPersec", (value))
}

// GetPendingInterruptsForwardedPersec gets the value of PendingInterruptsForwardedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPendingInterruptsForwardedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPendingInterruptsForwardingCost(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsForwardingCost", (value))
}

// GetPendingInterruptsForwardingCost gets the value of PendingInterruptsForwardingCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPendingInterruptsForwardingCost() (value uint64, err error) {
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

// SetPendingInterruptsForwardingCost_Base sets the value of PendingInterruptsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPendingInterruptsForwardingCost_Base(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsForwardingCost_Base", (value))
}

// GetPendingInterruptsForwardingCost_Base gets the value of PendingInterruptsForwardingCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPendingInterruptsForwardingCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingInterruptsForwardingCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPendingInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("PendingInterruptsPersec", (value))
}

// GetPendingInterruptsPersec gets the value of PendingInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPendingInterruptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentGuestRelativeUtilization(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRelativeUtilization", (value))
}

// GetPercentGuestRelativeUtilization gets the value of PercentGuestRelativeUtilization for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentGuestRelativeUtilization() (value uint64, err error) {
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

// SetPercentGuestRelativeUtilization_Base sets the value of PercentGuestRelativeUtilization_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentGuestRelativeUtilization_Base(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRelativeUtilization_Base", (value))
}

// GetPercentGuestRelativeUtilization_Base gets the value of PercentGuestRelativeUtilization_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentGuestRelativeUtilization_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentGuestRelativeUtilization_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentGuestRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRunTime", (value))
}

// GetPercentGuestRunTime gets the value of PercentGuestRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentGuestRunTime() (value uint64, err error) {
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

// SetPercentGuestRunTime_Base sets the value of PercentGuestRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentGuestRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRunTime_Base", (value))
}

// GetPercentGuestRunTime_Base gets the value of PercentGuestRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentGuestRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentGuestRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentHypervisorRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentHypervisorRunTime", (value))
}

// GetPercentHypervisorRunTime gets the value of PercentHypervisorRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentHypervisorRunTime() (value uint64, err error) {
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

// SetPercentHypervisorRunTime_Base sets the value of PercentHypervisorRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentHypervisorRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentHypervisorRunTime_Base", (value))
}

// GetPercentHypervisorRunTime_Base gets the value of PercentHypervisorRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentHypervisorRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentHypervisorRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentRemoteRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentRemoteRunTime", (value))
}

// GetPercentRemoteRunTime gets the value of PercentRemoteRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentRemoteRunTime() (value uint64, err error) {
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

// SetPercentRemoteRunTime_Base sets the value of PercentRemoteRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentRemoteRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentRemoteRunTime_Base", (value))
}

// GetPercentRemoteRunTime_Base gets the value of PercentRemoteRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentRemoteRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentRemoteRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentTotalCoreRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentTotalCoreRunTime", (value))
}

// GetPercentTotalCoreRunTime gets the value of PercentTotalCoreRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentTotalCoreRunTime() (value uint64, err error) {
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

// SetPercentTotalCoreRunTime_Base sets the value of PercentTotalCoreRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentTotalCoreRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentTotalCoreRunTime_Base", (value))
}

// GetPercentTotalCoreRunTime_Base gets the value of PercentTotalCoreRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentTotalCoreRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTotalCoreRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentTotalRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentTotalRunTime", (value))
}

// GetPercentTotalRunTime gets the value of PercentTotalRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentTotalRunTime() (value uint64, err error) {
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

// SetPercentTotalRunTime_Base sets the value of PercentTotalRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentTotalRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentTotalRunTime_Base", (value))
}

// GetPercentTotalRunTime_Base gets the value of PercentTotalRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentTotalRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTotalRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentVTL1RunTime(value uint64) (err error) {
	return instance.SetProperty("PercentVTL1RunTime", (value))
}

// GetPercentVTL1RunTime gets the value of PercentVTL1RunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentVTL1RunTime() (value uint64, err error) {
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

// SetPercentVTL1RunTime_Base sets the value of PercentVTL1RunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentVTL1RunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentVTL1RunTime_Base", (value))
}

// GetPercentVTL1RunTime_Base gets the value of PercentVTL1RunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentVTL1RunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentVTL1RunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentVTL2RunTime(value uint64) (err error) {
	return instance.SetProperty("PercentVTL2RunTime", (value))
}

// GetPercentVTL2RunTime gets the value of PercentVTL2RunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentVTL2RunTime() (value uint64, err error) {
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

// SetPercentVTL2RunTime_Base sets the value of PercentVTL2RunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPercentVTL2RunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentVTL2RunTime_Base", (value))
}

// GetPercentVTL2RunTime_Base gets the value of PercentVTL2RunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPercentVTL2RunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentVTL2RunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPerformanceMonitoringInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringInterruptsPersec", (value))
}

// GetPerformanceMonitoringInterruptsPersec gets the value of PerformanceMonitoringInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPerformanceMonitoringInterruptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPerformanceMonitoringIPTMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringIPTMSRAccessesPersec", (value))
}

// GetPerformanceMonitoringIPTMSRAccessesPersec gets the value of PerformanceMonitoringIPTMSRAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPerformanceMonitoringIPTMSRAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPerformanceMonitoringLBRMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringLBRMSRAccessesPersec", (value))
}

// GetPerformanceMonitoringLBRMSRAccessesPersec gets the value of PerformanceMonitoringLBRMSRAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPerformanceMonitoringLBRMSRAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPerformanceMonitoringvPMUMSRAccessesPersec(value uint64) (err error) {
	return instance.SetProperty("PerformanceMonitoringvPMUMSRAccessesPersec", (value))
}

// GetPerformanceMonitoringvPMUMSRAccessesPersec gets the value of PerformanceMonitoringvPMUMSRAccessesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPerformanceMonitoringvPMUMSRAccessesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPostedInterruptNotificationsPersec(value uint64) (err error) {
	return instance.SetProperty("PostedInterruptNotificationsPersec", (value))
}

// GetPostedInterruptNotificationsPersec gets the value of PostedInterruptNotificationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPostedInterruptNotificationsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyPostedInterruptScansPersec(value uint64) (err error) {
	return instance.SetProperty("PostedInterruptScansPersec", (value))
}

// GetPostedInterruptScansPersec gets the value of PostedInterruptScansPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyPostedInterruptScansPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyRDPMCInstructionsCost(value uint64) (err error) {
	return instance.SetProperty("RDPMCInstructionsCost", (value))
}

// GetRDPMCInstructionsCost gets the value of RDPMCInstructionsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyRDPMCInstructionsCost() (value uint64, err error) {
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

// SetRDPMCInstructionsCost_Base sets the value of RDPMCInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyRDPMCInstructionsCost_Base(value uint64) (err error) {
	return instance.SetProperty("RDPMCInstructionsCost_Base", (value))
}

// GetRDPMCInstructionsCost_Base gets the value of RDPMCInstructionsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyRDPMCInstructionsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("RDPMCInstructionsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyRDPMCInstructionsPersec(value uint64) (err error) {
	return instance.SetProperty("RDPMCInstructionsPersec", (value))
}

// GetRDPMCInstructionsPersec gets the value of RDPMCInstructionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyRDPMCInstructionsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyReflectedGuestPageFaultsPersec(value uint64) (err error) {
	return instance.SetProperty("ReflectedGuestPageFaultsPersec", (value))
}

// GetReflectedGuestPageFaultsPersec gets the value of ReflectedGuestPageFaultsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyReflectedGuestPageFaultsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertySchedulingPriority(value uint64) (err error) {
	return instance.SetProperty("SchedulingPriority", (value))
}

// GetSchedulingPriority gets the value of SchedulingPriority for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertySchedulingPriority() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertySmallPageTLBFillsPersec(value uint64) (err error) {
	return instance.SetProperty("SmallPageTLBFillsPersec", (value))
}

// GetSmallPageTLBFillsPersec gets the value of SmallPageTLBFillsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertySmallPageTLBFillsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertySVMHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("SVMHypercallsPersec", (value))
}

// GetSVMHypercallsPersec gets the value of SVMHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertySVMHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertySyntheticInterruptHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("SyntheticInterruptHypercallsPersec", (value))
}

// GetSyntheticInterruptHypercallsPersec gets the value of SyntheticInterruptHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertySyntheticInterruptHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertySyntheticInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("SyntheticInterruptsPersec", (value))
}

// GetSyntheticInterruptsPersec gets the value of SyntheticInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertySyntheticInterruptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalInterceptsCost(value uint64) (err error) {
	return instance.SetProperty("TotalInterceptsCost", (value))
}

// GetTotalInterceptsCost gets the value of TotalInterceptsCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalInterceptsCost() (value uint64, err error) {
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

// SetTotalInterceptsCost_Base sets the value of TotalInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalInterceptsCost_Base(value uint64) (err error) {
	return instance.SetProperty("TotalInterceptsCost_Base", (value))
}

// GetTotalInterceptsCost_Base gets the value of TotalInterceptsCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalInterceptsCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInterceptsCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("TotalInterceptsPersec", (value))
}

// GetTotalInterceptsPersec gets the value of TotalInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("TotalMessagesPersec", (value))
}

// GetTotalMessagesPersec gets the value of TotalMessagesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalMessagesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalVirtualizationInstructionsEmulatedPersec(value uint64) (err error) {
	return instance.SetProperty("TotalVirtualizationInstructionsEmulatedPersec", (value))
}

// GetTotalVirtualizationInstructionsEmulatedPersec gets the value of TotalVirtualizationInstructionsEmulatedPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalVirtualizationInstructionsEmulatedPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalVirtualizationInstructionsEmulationCost(value uint64) (err error) {
	return instance.SetProperty("TotalVirtualizationInstructionsEmulationCost", (value))
}

// GetTotalVirtualizationInstructionsEmulationCost gets the value of TotalVirtualizationInstructionsEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalVirtualizationInstructionsEmulationCost() (value uint64, err error) {
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

// SetTotalVirtualizationInstructionsEmulationCost_Base sets the value of TotalVirtualizationInstructionsEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyTotalVirtualizationInstructionsEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("TotalVirtualizationInstructionsEmulationCost_Base", (value))
}

// GetTotalVirtualizationInstructionsEmulationCost_Base gets the value of TotalVirtualizationInstructionsEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyTotalVirtualizationInstructionsEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalVirtualizationInstructionsEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVirtualInterruptHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualInterruptHypercallsPersec", (value))
}

// GetVirtualInterruptHypercallsPersec gets the value of VirtualInterruptHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVirtualInterruptHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVirtualInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualInterruptsPersec", (value))
}

// GetVirtualInterruptsPersec gets the value of VirtualInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVirtualInterruptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVirtualMMUHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualMMUHypercallsPersec", (value))
}

// GetVirtualMMUHypercallsPersec gets the value of VirtualMMUHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVirtualMMUHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVirtualProcessorHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualProcessorHypercallsPersec", (value))
}

// GetVirtualProcessorHypercallsPersec gets the value of VirtualProcessorHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVirtualProcessorHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMCLEAREmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMCLEAREmulationInterceptsPersec", (value))
}

// GetVMCLEAREmulationInterceptsPersec gets the value of VMCLEAREmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMCLEAREmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMCLEARInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMCLEARInstructionEmulationCost", (value))
}

// GetVMCLEARInstructionEmulationCost gets the value of VMCLEARInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMCLEARInstructionEmulationCost() (value uint64, err error) {
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

// SetVMCLEARInstructionEmulationCost_Base sets the value of VMCLEARInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMCLEARInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMCLEARInstructionEmulationCost_Base", (value))
}

// GetVMCLEARInstructionEmulationCost_Base gets the value of VMCLEARInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMCLEARInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMCLEARInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMLOADEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMLOADEmulationInterceptsPersec", (value))
}

// GetVMLOADEmulationInterceptsPersec gets the value of VMLOADEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMLOADEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMLOADInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMLOADInstructionEmulationCost", (value))
}

// GetVMLOADInstructionEmulationCost gets the value of VMLOADInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMLOADInstructionEmulationCost() (value uint64, err error) {
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

// SetVMLOADInstructionEmulationCost_Base sets the value of VMLOADInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMLOADInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMLOADInstructionEmulationCost_Base", (value))
}

// GetVMLOADInstructionEmulationCost_Base gets the value of VMLOADInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMLOADInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMLOADInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMPTRLDEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMPTRLDEmulationInterceptsPersec", (value))
}

// GetVMPTRLDEmulationInterceptsPersec gets the value of VMPTRLDEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMPTRLDEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMPTRLDInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMPTRLDInstructionEmulationCost", (value))
}

// GetVMPTRLDInstructionEmulationCost gets the value of VMPTRLDInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMPTRLDInstructionEmulationCost() (value uint64, err error) {
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

// SetVMPTRLDInstructionEmulationCost_Base sets the value of VMPTRLDInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMPTRLDInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMPTRLDInstructionEmulationCost_Base", (value))
}

// GetVMPTRLDInstructionEmulationCost_Base gets the value of VMPTRLDInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMPTRLDInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMPTRLDInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMPTRSTEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMPTRSTEmulationInterceptsPersec", (value))
}

// GetVMPTRSTEmulationInterceptsPersec gets the value of VMPTRSTEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMPTRSTEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMPTRSTInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMPTRSTInstructionEmulationCost", (value))
}

// GetVMPTRSTInstructionEmulationCost gets the value of VMPTRSTInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMPTRSTInstructionEmulationCost() (value uint64, err error) {
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

// SetVMPTRSTInstructionEmulationCost_Base sets the value of VMPTRSTInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMPTRSTInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMPTRSTInstructionEmulationCost_Base", (value))
}

// GetVMPTRSTInstructionEmulationCost_Base gets the value of VMPTRSTInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMPTRSTInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMPTRSTInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMREADEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMREADEmulationInterceptsPersec", (value))
}

// GetVMREADEmulationInterceptsPersec gets the value of VMREADEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMREADEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMREADInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMREADInstructionEmulationCost", (value))
}

// GetVMREADInstructionEmulationCost gets the value of VMREADInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMREADInstructionEmulationCost() (value uint64, err error) {
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

// SetVMREADInstructionEmulationCost_Base sets the value of VMREADInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMREADInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMREADInstructionEmulationCost_Base", (value))
}

// GetVMREADInstructionEmulationCost_Base gets the value of VMREADInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMREADInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMREADInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMSAVEEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMSAVEEmulationInterceptsPersec", (value))
}

// GetVMSAVEEmulationInterceptsPersec gets the value of VMSAVEEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMSAVEEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMSAVEInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMSAVEInstructionEmulationCost", (value))
}

// GetVMSAVEInstructionEmulationCost gets the value of VMSAVEInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMSAVEInstructionEmulationCost() (value uint64, err error) {
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

// SetVMSAVEInstructionEmulationCost_Base sets the value of VMSAVEInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMSAVEInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMSAVEInstructionEmulationCost_Base", (value))
}

// GetVMSAVEInstructionEmulationCost_Base gets the value of VMSAVEInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMSAVEInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMSAVEInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMWRITEEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMWRITEEmulationInterceptsPersec", (value))
}

// GetVMWRITEEmulationInterceptsPersec gets the value of VMWRITEEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMWRITEEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMWRITEInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMWRITEInstructionEmulationCost", (value))
}

// GetVMWRITEInstructionEmulationCost gets the value of VMWRITEInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMWRITEInstructionEmulationCost() (value uint64, err error) {
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

// SetVMWRITEInstructionEmulationCost_Base sets the value of VMWRITEInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMWRITEInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMWRITEInstructionEmulationCost_Base", (value))
}

// GetVMWRITEInstructionEmulationCost_Base gets the value of VMWRITEInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMWRITEInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMWRITEInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMXOFFEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMXOFFEmulationInterceptsPersec", (value))
}

// GetVMXOFFEmulationInterceptsPersec gets the value of VMXOFFEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMXOFFEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMXOFFInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMXOFFInstructionEmulationCost", (value))
}

// GetVMXOFFInstructionEmulationCost gets the value of VMXOFFInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMXOFFInstructionEmulationCost() (value uint64, err error) {
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

// SetVMXOFFInstructionEmulationCost_Base sets the value of VMXOFFInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMXOFFInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMXOFFInstructionEmulationCost_Base", (value))
}

// GetVMXOFFInstructionEmulationCost_Base gets the value of VMXOFFInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMXOFFInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMXOFFInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMXONEmulationInterceptsPersec(value uint64) (err error) {
	return instance.SetProperty("VMXONEmulationInterceptsPersec", (value))
}

// GetVMXONEmulationInterceptsPersec gets the value of VMXONEmulationInterceptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMXONEmulationInterceptsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMXONInstructionEmulationCost(value uint64) (err error) {
	return instance.SetProperty("VMXONInstructionEmulationCost", (value))
}

// GetVMXONInstructionEmulationCost gets the value of VMXONInstructionEmulationCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMXONInstructionEmulationCost() (value uint64, err error) {
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

// SetVMXONInstructionEmulationCost_Base sets the value of VMXONInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVMXONInstructionEmulationCost_Base(value uint64) (err error) {
	return instance.SetProperty("VMXONInstructionEmulationCost_Base", (value))
}

// GetVMXONInstructionEmulationCost_Base gets the value of VMXONInstructionEmulationCost_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVMXONInstructionEmulationCost_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VMXONInstructionEmulationCost_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVSMHypercallsPersec(value uint64) (err error) {
	return instance.SetProperty("VSMHypercallsPersec", (value))
}

// GetVSMHypercallsPersec gets the value of VSMHypercallsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVSMHypercallsPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVTL1AverageRunTime(value uint64) (err error) {
	return instance.SetProperty("VTL1AverageRunTime", (value))
}

// GetVTL1AverageRunTime gets the value of VTL1AverageRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVTL1AverageRunTime() (value uint64, err error) {
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

// SetVTL1AverageRunTime_Base sets the value of VTL1AverageRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVTL1AverageRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("VTL1AverageRunTime_Base", (value))
}

// GetVTL1AverageRunTime_Base gets the value of VTL1AverageRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVTL1AverageRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VTL1AverageRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVTL1DispatchesPersec(value uint64) (err error) {
	return instance.SetProperty("VTL1DispatchesPersec", (value))
}

// GetVTL1DispatchesPersec gets the value of VTL1DispatchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVTL1DispatchesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVTL2AverageRunTime(value uint64) (err error) {
	return instance.SetProperty("VTL2AverageRunTime", (value))
}

// GetVTL2AverageRunTime gets the value of VTL2AverageRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVTL2AverageRunTime() (value uint64, err error) {
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

// SetVTL2AverageRunTime_Base sets the value of VTL2AverageRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVTL2AverageRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("VTL2AverageRunTime_Base", (value))
}

// GetVTL2AverageRunTime_Base gets the value of VTL2AverageRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVTL2AverageRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VTL2AverageRunTime_Base")
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) SetPropertyVTL2DispatchesPersec(value uint64) (err error) {
	return instance.SetProperty("VTL2DispatchesPersec", (value))
}

// GetVTL2DispatchesPersec gets the value of VTL2DispatchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorRootVirtualProcessor) GetPropertyVTL2DispatchesPersec() (value uint64, err error) {
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
