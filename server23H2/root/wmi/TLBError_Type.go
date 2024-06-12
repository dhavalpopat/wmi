// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source TLBError_Type
//////////////////////////////////////////////
package wmi

// TLBError_Type
type TLBError_Type int

const (
	// MCA_WARNING_CACHE enum
	TLBError_Type_MCA_WARNING_CACHE TLBError_Type = 1
	// MCA_ERROR_CACHE enum
	TLBError_Type_MCA_ERROR_CACHE TLBError_Type = 2
	// MCA_WARNING_TLB enum
	TLBError_Type_MCA_WARNING_TLB TLBError_Type = 3
	// MCA_ERROR_TLB enum
	TLBError_Type_MCA_ERROR_TLB TLBError_Type = 4
	// MCA_WARNING_CPU_BUS enum
	TLBError_Type_MCA_WARNING_CPU_BUS TLBError_Type = 5
	// MCA_ERROR_CPU_BUS enum
	TLBError_Type_MCA_ERROR_CPU_BUS TLBError_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	TLBError_Type_MCA_WARNING_REGISTER_FILE TLBError_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	TLBError_Type_MCA_ERROR_REGISTER_FILE TLBError_Type = 8
	// MCA_WARNING_MAS enum
	TLBError_Type_MCA_WARNING_MAS TLBError_Type = 9
	// MCA_ERROR_MAS enum
	TLBError_Type_MCA_ERROR_MAS TLBError_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	TLBError_Type_MCA_WARNING_MEM_UNKNOWN TLBError_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	TLBError_Type_MCA_ERROR_MEM_UNKNOWN TLBError_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	TLBError_Type_MCA_WARNING_MEM_1_2 TLBError_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	TLBError_Type_MCA_ERROR_MEM_1_2 TLBError_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	TLBError_Type_MCA_WARNING_MEM_1_2_5 TLBError_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	TLBError_Type_MCA_ERROR_MEM_1_2_5 TLBError_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	TLBError_Type_MCA_WARNING_MEM_1_2_5_4 TLBError_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	TLBError_Type_MCA_ERROR_MEM_1_2_5_4 TLBError_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	TLBError_Type_MCA_WARNING_SYSTEM_EVENT TLBError_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	TLBError_Type_MCA_ERROR_SYSTEM_EVENT TLBError_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	TLBError_Type_MCA_WARNING_PCI_BUS_PARITY TLBError_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	TLBError_Type_MCA_ERROR_PCI_BUS_PARITY TLBError_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	TLBError_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO TLBError_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	TLBError_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO TLBError_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	TLBError_Type_MCA_WARNING_PCI_BUS_SERR TLBError_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	TLBError_Type_MCA_ERROR_PCI_BUS_SERR TLBError_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	TLBError_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO TLBError_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	TLBError_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO TLBError_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	TLBError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT TLBError_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	TLBError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT TLBError_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	TLBError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO TLBError_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	TLBError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO TLBError_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	TLBError_Type_MCA_WARNING_PCI_BUS_TIMEOUT TLBError_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	TLBError_Type_MCA_ERROR_PCI_BUS_TIMEOUT TLBError_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	TLBError_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO TLBError_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	TLBError_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO TLBError_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	TLBError_Type_MCA_WARNING_PCI_BUS_UNKNOWN TLBError_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	TLBError_Type_MCA_ERROR_PCI_BUS_UNKNOWN TLBError_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	TLBError_Type_MCA_WARNING_PCI_DEVICE TLBError_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	TLBError_Type_MCA_ERROR_PCI_DEVICE TLBError_Type = 40
	// MCA_WARNING_SMBIOS enum
	TLBError_Type_MCA_WARNING_SMBIOS TLBError_Type = 41
	// MCA_ERROR_SMBIOS enum
	TLBError_Type_MCA_ERROR_SMBIOS TLBError_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	TLBError_Type_MCA_WARNING_PLATFORM_SPECIFIC TLBError_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	TLBError_Type_MCA_ERROR_PLATFORM_SPECIFIC TLBError_Type = 44
	// MCA_WARNING_UNKNOWN enum
	TLBError_Type_MCA_WARNING_UNKNOWN TLBError_Type = 45
	// MCA_ERROR_UNKNOWN enum
	TLBError_Type_MCA_ERROR_UNKNOWN TLBError_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	TLBError_Type_MCA_WARNING_UNKNOWN_NO_CPU TLBError_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	TLBError_Type_MCA_ERROR_UNKNOWN_NO_CPU TLBError_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	TLBError_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED TLBError_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	TLBError_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED TLBError_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	TLBError_Type_MCA_WARNING_CPU_THERMAL_THROTTLED TLBError_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	TLBError_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED TLBError_Type = 52
	// MCA_WARNING_CPU enum
	TLBError_Type_MCA_WARNING_CPU TLBError_Type = 53
	// MCA_ERROR_CPU enum
	TLBError_Type_MCA_ERROR_CPU TLBError_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	TLBError_Type_MCA_MEMORYHIERARCHY_ERROR TLBError_Type = 55
	// MCA_TLB_ERROR enum
	TLBError_Type_MCA_TLB_ERROR TLBError_Type = 56
	// MCA_BUS_ERROR enum
	TLBError_Type_MCA_BUS_ERROR TLBError_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	TLBError_Type_MCA_BUS_TIMEOUT_ERROR TLBError_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	TLBError_Type_MCA_INTERNALTIMER_ERROR TLBError_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	TLBError_Type_MCA_MICROCODE_ROM_PARITY_ERROR TLBError_Type = 60
	// MCA_EXTERNAL_ERROR enum
	TLBError_Type_MCA_EXTERNAL_ERROR TLBError_Type = 61
	// MCA_FRC_ERROR enum
	TLBError_Type_MCA_FRC_ERROR TLBError_Type = 62
)
