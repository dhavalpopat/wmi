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

// Win32_PerfRawData_Counters_ReFS struct
type Win32_PerfRawData_Counters_ReFS struct {
	*Win32_PerfRawData

	//
	AllocationofDataClustersonFastTierPersec uint64

	//
	AllocationofDataClustersonSlowTierPersec uint64

	//
	AllocationofMetadataClustersonFastTierPersec uint64

	//
	AllocationofMetadataClustersonSlowTierPersec uint64

	//
	AllocatorCacheCollisions uint32

	//
	AllocatorCacheDeletions uint32

	//
	AllocatorCacheInsertions uint32

	//
	AllocatorCachePurges uint32

	//
	AllocatorCacheSize uint32

	//
	AllocatorRegionsExaminedAverage uint64

	//
	AllocatorRegionsExaminedAverage_Base uint32

	//
	AllocatorRegionsExaminedMax uint32

	//
	Checkpointastreeupdatelatency100ns uint64

	//
	Checkpointastreeupdatelatency100ns_Base uint32

	//
	Checkpointexclusivelockslatency100ns uint64

	//
	Checkpointexclusivelockslatency100ns_Base uint32

	//
	Checkpointlatency100ns uint64

	//
	Checkpointlatency100ns_Base uint32

	//
	CheckpointsPersec uint64

	//
	CompactedContainerFillRatioPercent uint64

	//
	CompactedContainerFillRatioPercent_Base uint32

	//
	CompactionFailureCount uint32

	//
	Compactionreadlatency100ns uint64

	//
	Compactionreadlatency100ns_Base uint32

	//
	Compactionsfailedduetoineligiblecontainer uint32

	//
	Compactionsfailedduetomaxfragmentation uint32

	//
	Compactionwritelatency100ns uint64

	//
	Compactionwritelatency100ns_Base uint32

	//
	ContainerDestagesFromFastTierPersec uint64

	//
	ContainerDestagesFromSlowTierPersec uint64

	//
	ContainerMoveFailureCount uint32

	//
	ContainerMoveRetryCount uint32

	//
	Containermovesfailedduetoineligiblecontainer uint32

	//
	CurrentFastTierDataFillPercentage uint32

	//
	CurrentFastTierMetadataFillPercentage uint32

	//
	CurrentSlowTierDataFillPercentage uint32

	//
	CurrentSlowTierMetadataFillPercentage uint32

	//
	DataCompactionsPersec uint64

	//
	DataInPlaceWriteClustersPersec uint64

	//
	DeleteQueueentries uint32

	//
	Dirtymetadatapages uint64

	//
	Dirtytablelistentries uint32

	//
	FastTierDestagedContainerFillRatioPercent uint64

	//
	FastTierDestagedContainerFillRatioPercent_Base uint32

	//
	Fasttierdestagereadlatency100ns uint64

	//
	Fasttierdestagereadlatency100ns_Base uint32

	//
	Fasttierdestagewritelatency100ns uint64

	//
	Fasttierdestagewritelatency100ns_Base uint32

	//
	Logfillpercentage uint32

	//
	LogwritesPersec uint64

	//
	SlowTierDestagedContainerFillRatioPercent uint64

	//
	SlowTierDestagedContainerFillRatioPercent_Base uint32

	//
	Slowtierdestagereadlatency100ns uint64

	//
	Slowtierdestagereadlatency100ns_Base uint32

	//
	Slowtierdestagewritelatency100ns uint64

	//
	Slowtierdestagewritelatency100ns_Base uint32

	//
	TotalAllocationofClustersPersec uint64

	//
	Treeupdateexclusivelockslatency100ns uint64

	//
	Treeupdateexclusivelockslatency100ns_Base uint32

	//
	Treeupdatelatency100ns uint64

	//
	Treeupdatelatency100ns_Base uint32

	//
	TreeupdatesPersec uint64

	//
	Trimlatency100ns uint64

	//
	Trimlatency100ns_Base uint32
}

func NewWin32_PerfRawData_Counters_ReFSEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_ReFS, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_ReFS{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_ReFSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_ReFS, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_ReFS{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAllocationofDataClustersonFastTierPersec sets the value of AllocationofDataClustersonFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocationofDataClustersonFastTierPersec(value uint64) (err error) {
	return instance.SetProperty("AllocationofDataClustersonFastTierPersec", (value))
}

// GetAllocationofDataClustersonFastTierPersec gets the value of AllocationofDataClustersonFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocationofDataClustersonFastTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationofDataClustersonFastTierPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAllocationofDataClustersonSlowTierPersec sets the value of AllocationofDataClustersonSlowTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocationofDataClustersonSlowTierPersec(value uint64) (err error) {
	return instance.SetProperty("AllocationofDataClustersonSlowTierPersec", (value))
}

// GetAllocationofDataClustersonSlowTierPersec gets the value of AllocationofDataClustersonSlowTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocationofDataClustersonSlowTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationofDataClustersonSlowTierPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAllocationofMetadataClustersonFastTierPersec sets the value of AllocationofMetadataClustersonFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocationofMetadataClustersonFastTierPersec(value uint64) (err error) {
	return instance.SetProperty("AllocationofMetadataClustersonFastTierPersec", (value))
}

// GetAllocationofMetadataClustersonFastTierPersec gets the value of AllocationofMetadataClustersonFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocationofMetadataClustersonFastTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationofMetadataClustersonFastTierPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAllocationofMetadataClustersonSlowTierPersec sets the value of AllocationofMetadataClustersonSlowTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocationofMetadataClustersonSlowTierPersec(value uint64) (err error) {
	return instance.SetProperty("AllocationofMetadataClustersonSlowTierPersec", (value))
}

// GetAllocationofMetadataClustersonSlowTierPersec gets the value of AllocationofMetadataClustersonSlowTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocationofMetadataClustersonSlowTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationofMetadataClustersonSlowTierPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAllocatorCacheCollisions sets the value of AllocatorCacheCollisions for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorCacheCollisions(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheCollisions", (value))
}

// GetAllocatorCacheCollisions gets the value of AllocatorCacheCollisions for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorCacheCollisions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheCollisions")
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

// SetAllocatorCacheDeletions sets the value of AllocatorCacheDeletions for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorCacheDeletions(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheDeletions", (value))
}

// GetAllocatorCacheDeletions gets the value of AllocatorCacheDeletions for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorCacheDeletions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheDeletions")
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

// SetAllocatorCacheInsertions sets the value of AllocatorCacheInsertions for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorCacheInsertions(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheInsertions", (value))
}

// GetAllocatorCacheInsertions gets the value of AllocatorCacheInsertions for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorCacheInsertions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheInsertions")
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

// SetAllocatorCachePurges sets the value of AllocatorCachePurges for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorCachePurges(value uint32) (err error) {
	return instance.SetProperty("AllocatorCachePurges", (value))
}

// GetAllocatorCachePurges gets the value of AllocatorCachePurges for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorCachePurges() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCachePurges")
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

// SetAllocatorCacheSize sets the value of AllocatorCacheSize for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorCacheSize(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheSize", (value))
}

// GetAllocatorCacheSize gets the value of AllocatorCacheSize for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheSize")
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

// SetAllocatorRegionsExaminedAverage sets the value of AllocatorRegionsExaminedAverage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorRegionsExaminedAverage(value uint64) (err error) {
	return instance.SetProperty("AllocatorRegionsExaminedAverage", (value))
}

// GetAllocatorRegionsExaminedAverage gets the value of AllocatorRegionsExaminedAverage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorRegionsExaminedAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocatorRegionsExaminedAverage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAllocatorRegionsExaminedAverage_Base sets the value of AllocatorRegionsExaminedAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorRegionsExaminedAverage_Base(value uint32) (err error) {
	return instance.SetProperty("AllocatorRegionsExaminedAverage_Base", (value))
}

// GetAllocatorRegionsExaminedAverage_Base gets the value of AllocatorRegionsExaminedAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorRegionsExaminedAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorRegionsExaminedAverage_Base")
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

// SetAllocatorRegionsExaminedMax sets the value of AllocatorRegionsExaminedMax for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyAllocatorRegionsExaminedMax(value uint32) (err error) {
	return instance.SetProperty("AllocatorRegionsExaminedMax", (value))
}

// GetAllocatorRegionsExaminedMax gets the value of AllocatorRegionsExaminedMax for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyAllocatorRegionsExaminedMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorRegionsExaminedMax")
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

// SetCheckpointastreeupdatelatency100ns sets the value of Checkpointastreeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointastreeupdatelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Checkpointastreeupdatelatency100ns", (value))
}

// GetCheckpointastreeupdatelatency100ns gets the value of Checkpointastreeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointastreeupdatelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Checkpointastreeupdatelatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCheckpointastreeupdatelatency100ns_Base sets the value of Checkpointastreeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointastreeupdatelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Checkpointastreeupdatelatency100ns_Base", (value))
}

// GetCheckpointastreeupdatelatency100ns_Base gets the value of Checkpointastreeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointastreeupdatelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Checkpointastreeupdatelatency100ns_Base")
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

// SetCheckpointexclusivelockslatency100ns sets the value of Checkpointexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointexclusivelockslatency100ns(value uint64) (err error) {
	return instance.SetProperty("Checkpointexclusivelockslatency100ns", (value))
}

// GetCheckpointexclusivelockslatency100ns gets the value of Checkpointexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointexclusivelockslatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Checkpointexclusivelockslatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCheckpointexclusivelockslatency100ns_Base sets the value of Checkpointexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointexclusivelockslatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Checkpointexclusivelockslatency100ns_Base", (value))
}

// GetCheckpointexclusivelockslatency100ns_Base gets the value of Checkpointexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointexclusivelockslatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Checkpointexclusivelockslatency100ns_Base")
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

// SetCheckpointlatency100ns sets the value of Checkpointlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointlatency100ns(value uint64) (err error) {
	return instance.SetProperty("Checkpointlatency100ns", (value))
}

// GetCheckpointlatency100ns gets the value of Checkpointlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointlatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Checkpointlatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCheckpointlatency100ns_Base sets the value of Checkpointlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointlatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Checkpointlatency100ns_Base", (value))
}

// GetCheckpointlatency100ns_Base gets the value of Checkpointlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointlatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Checkpointlatency100ns_Base")
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

// SetCheckpointsPersec sets the value of CheckpointsPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCheckpointsPersec(value uint64) (err error) {
	return instance.SetProperty("CheckpointsPersec", (value))
}

// GetCheckpointsPersec gets the value of CheckpointsPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCheckpointsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CheckpointsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCompactedContainerFillRatioPercent sets the value of CompactedContainerFillRatioPercent for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactedContainerFillRatioPercent(value uint64) (err error) {
	return instance.SetProperty("CompactedContainerFillRatioPercent", (value))
}

// GetCompactedContainerFillRatioPercent gets the value of CompactedContainerFillRatioPercent for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactedContainerFillRatioPercent() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompactedContainerFillRatioPercent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCompactedContainerFillRatioPercent_Base sets the value of CompactedContainerFillRatioPercent_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactedContainerFillRatioPercent_Base(value uint32) (err error) {
	return instance.SetProperty("CompactedContainerFillRatioPercent_Base", (value))
}

// GetCompactedContainerFillRatioPercent_Base gets the value of CompactedContainerFillRatioPercent_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactedContainerFillRatioPercent_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompactedContainerFillRatioPercent_Base")
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

// SetCompactionFailureCount sets the value of CompactionFailureCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionFailureCount(value uint32) (err error) {
	return instance.SetProperty("CompactionFailureCount", (value))
}

// GetCompactionFailureCount gets the value of CompactionFailureCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompactionFailureCount")
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

// SetCompactionreadlatency100ns sets the value of Compactionreadlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionreadlatency100ns(value uint64) (err error) {
	return instance.SetProperty("Compactionreadlatency100ns", (value))
}

// GetCompactionreadlatency100ns gets the value of Compactionreadlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionreadlatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Compactionreadlatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCompactionreadlatency100ns_Base sets the value of Compactionreadlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionreadlatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Compactionreadlatency100ns_Base", (value))
}

// GetCompactionreadlatency100ns_Base gets the value of Compactionreadlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionreadlatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Compactionreadlatency100ns_Base")
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

// SetCompactionsfailedduetoineligiblecontainer sets the value of Compactionsfailedduetoineligiblecontainer for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionsfailedduetoineligiblecontainer(value uint32) (err error) {
	return instance.SetProperty("Compactionsfailedduetoineligiblecontainer", (value))
}

// GetCompactionsfailedduetoineligiblecontainer gets the value of Compactionsfailedduetoineligiblecontainer for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionsfailedduetoineligiblecontainer() (value uint32, err error) {
	retValue, err := instance.GetProperty("Compactionsfailedduetoineligiblecontainer")
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

// SetCompactionsfailedduetomaxfragmentation sets the value of Compactionsfailedduetomaxfragmentation for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionsfailedduetomaxfragmentation(value uint32) (err error) {
	return instance.SetProperty("Compactionsfailedduetomaxfragmentation", (value))
}

// GetCompactionsfailedduetomaxfragmentation gets the value of Compactionsfailedduetomaxfragmentation for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionsfailedduetomaxfragmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("Compactionsfailedduetomaxfragmentation")
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

// SetCompactionwritelatency100ns sets the value of Compactionwritelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionwritelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Compactionwritelatency100ns", (value))
}

// GetCompactionwritelatency100ns gets the value of Compactionwritelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionwritelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Compactionwritelatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCompactionwritelatency100ns_Base sets the value of Compactionwritelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCompactionwritelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Compactionwritelatency100ns_Base", (value))
}

// GetCompactionwritelatency100ns_Base gets the value of Compactionwritelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCompactionwritelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Compactionwritelatency100ns_Base")
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

// SetContainerDestagesFromFastTierPersec sets the value of ContainerDestagesFromFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyContainerDestagesFromFastTierPersec(value uint64) (err error) {
	return instance.SetProperty("ContainerDestagesFromFastTierPersec", (value))
}

// GetContainerDestagesFromFastTierPersec gets the value of ContainerDestagesFromFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyContainerDestagesFromFastTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ContainerDestagesFromFastTierPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetContainerDestagesFromSlowTierPersec sets the value of ContainerDestagesFromSlowTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyContainerDestagesFromSlowTierPersec(value uint64) (err error) {
	return instance.SetProperty("ContainerDestagesFromSlowTierPersec", (value))
}

// GetContainerDestagesFromSlowTierPersec gets the value of ContainerDestagesFromSlowTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyContainerDestagesFromSlowTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ContainerDestagesFromSlowTierPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetContainerMoveFailureCount sets the value of ContainerMoveFailureCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyContainerMoveFailureCount(value uint32) (err error) {
	return instance.SetProperty("ContainerMoveFailureCount", (value))
}

// GetContainerMoveFailureCount gets the value of ContainerMoveFailureCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyContainerMoveFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContainerMoveFailureCount")
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

// SetContainerMoveRetryCount sets the value of ContainerMoveRetryCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyContainerMoveRetryCount(value uint32) (err error) {
	return instance.SetProperty("ContainerMoveRetryCount", (value))
}

// GetContainerMoveRetryCount gets the value of ContainerMoveRetryCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyContainerMoveRetryCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContainerMoveRetryCount")
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

// SetContainermovesfailedduetoineligiblecontainer sets the value of Containermovesfailedduetoineligiblecontainer for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyContainermovesfailedduetoineligiblecontainer(value uint32) (err error) {
	return instance.SetProperty("Containermovesfailedduetoineligiblecontainer", (value))
}

// GetContainermovesfailedduetoineligiblecontainer gets the value of Containermovesfailedduetoineligiblecontainer for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyContainermovesfailedduetoineligiblecontainer() (value uint32, err error) {
	retValue, err := instance.GetProperty("Containermovesfailedduetoineligiblecontainer")
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

// SetCurrentFastTierDataFillPercentage sets the value of CurrentFastTierDataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCurrentFastTierDataFillPercentage(value uint32) (err error) {
	return instance.SetProperty("CurrentFastTierDataFillPercentage", (value))
}

// GetCurrentFastTierDataFillPercentage gets the value of CurrentFastTierDataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCurrentFastTierDataFillPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentFastTierDataFillPercentage")
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

// SetCurrentFastTierMetadataFillPercentage sets the value of CurrentFastTierMetadataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCurrentFastTierMetadataFillPercentage(value uint32) (err error) {
	return instance.SetProperty("CurrentFastTierMetadataFillPercentage", (value))
}

// GetCurrentFastTierMetadataFillPercentage gets the value of CurrentFastTierMetadataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCurrentFastTierMetadataFillPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentFastTierMetadataFillPercentage")
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

// SetCurrentSlowTierDataFillPercentage sets the value of CurrentSlowTierDataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCurrentSlowTierDataFillPercentage(value uint32) (err error) {
	return instance.SetProperty("CurrentSlowTierDataFillPercentage", (value))
}

// GetCurrentSlowTierDataFillPercentage gets the value of CurrentSlowTierDataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCurrentSlowTierDataFillPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentSlowTierDataFillPercentage")
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

// SetCurrentSlowTierMetadataFillPercentage sets the value of CurrentSlowTierMetadataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyCurrentSlowTierMetadataFillPercentage(value uint32) (err error) {
	return instance.SetProperty("CurrentSlowTierMetadataFillPercentage", (value))
}

// GetCurrentSlowTierMetadataFillPercentage gets the value of CurrentSlowTierMetadataFillPercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyCurrentSlowTierMetadataFillPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentSlowTierMetadataFillPercentage")
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

// SetDataCompactionsPersec sets the value of DataCompactionsPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyDataCompactionsPersec(value uint64) (err error) {
	return instance.SetProperty("DataCompactionsPersec", (value))
}

// GetDataCompactionsPersec gets the value of DataCompactionsPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyDataCompactionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataCompactionsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDataInPlaceWriteClustersPersec sets the value of DataInPlaceWriteClustersPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyDataInPlaceWriteClustersPersec(value uint64) (err error) {
	return instance.SetProperty("DataInPlaceWriteClustersPersec", (value))
}

// GetDataInPlaceWriteClustersPersec gets the value of DataInPlaceWriteClustersPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyDataInPlaceWriteClustersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataInPlaceWriteClustersPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDeleteQueueentries sets the value of DeleteQueueentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyDeleteQueueentries(value uint32) (err error) {
	return instance.SetProperty("DeleteQueueentries", (value))
}

// GetDeleteQueueentries gets the value of DeleteQueueentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyDeleteQueueentries() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeleteQueueentries")
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

// SetDirtymetadatapages sets the value of Dirtymetadatapages for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyDirtymetadatapages(value uint64) (err error) {
	return instance.SetProperty("Dirtymetadatapages", (value))
}

// GetDirtymetadatapages gets the value of Dirtymetadatapages for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyDirtymetadatapages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Dirtymetadatapages")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDirtytablelistentries sets the value of Dirtytablelistentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyDirtytablelistentries(value uint32) (err error) {
	return instance.SetProperty("Dirtytablelistentries", (value))
}

// GetDirtytablelistentries gets the value of Dirtytablelistentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyDirtytablelistentries() (value uint32, err error) {
	retValue, err := instance.GetProperty("Dirtytablelistentries")
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

// SetFastTierDestagedContainerFillRatioPercent sets the value of FastTierDestagedContainerFillRatioPercent for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyFastTierDestagedContainerFillRatioPercent(value uint64) (err error) {
	return instance.SetProperty("FastTierDestagedContainerFillRatioPercent", (value))
}

// GetFastTierDestagedContainerFillRatioPercent gets the value of FastTierDestagedContainerFillRatioPercent for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyFastTierDestagedContainerFillRatioPercent() (value uint64, err error) {
	retValue, err := instance.GetProperty("FastTierDestagedContainerFillRatioPercent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFastTierDestagedContainerFillRatioPercent_Base sets the value of FastTierDestagedContainerFillRatioPercent_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyFastTierDestagedContainerFillRatioPercent_Base(value uint32) (err error) {
	return instance.SetProperty("FastTierDestagedContainerFillRatioPercent_Base", (value))
}

// GetFastTierDestagedContainerFillRatioPercent_Base gets the value of FastTierDestagedContainerFillRatioPercent_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyFastTierDestagedContainerFillRatioPercent_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("FastTierDestagedContainerFillRatioPercent_Base")
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

// SetFasttierdestagereadlatency100ns sets the value of Fasttierdestagereadlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyFasttierdestagereadlatency100ns(value uint64) (err error) {
	return instance.SetProperty("Fasttierdestagereadlatency100ns", (value))
}

// GetFasttierdestagereadlatency100ns gets the value of Fasttierdestagereadlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyFasttierdestagereadlatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Fasttierdestagereadlatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFasttierdestagereadlatency100ns_Base sets the value of Fasttierdestagereadlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyFasttierdestagereadlatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Fasttierdestagereadlatency100ns_Base", (value))
}

// GetFasttierdestagereadlatency100ns_Base gets the value of Fasttierdestagereadlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyFasttierdestagereadlatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Fasttierdestagereadlatency100ns_Base")
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

// SetFasttierdestagewritelatency100ns sets the value of Fasttierdestagewritelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyFasttierdestagewritelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Fasttierdestagewritelatency100ns", (value))
}

// GetFasttierdestagewritelatency100ns gets the value of Fasttierdestagewritelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyFasttierdestagewritelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Fasttierdestagewritelatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFasttierdestagewritelatency100ns_Base sets the value of Fasttierdestagewritelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyFasttierdestagewritelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Fasttierdestagewritelatency100ns_Base", (value))
}

// GetFasttierdestagewritelatency100ns_Base gets the value of Fasttierdestagewritelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyFasttierdestagewritelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Fasttierdestagewritelatency100ns_Base")
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

// SetLogfillpercentage sets the value of Logfillpercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyLogfillpercentage(value uint32) (err error) {
	return instance.SetProperty("Logfillpercentage", (value))
}

// GetLogfillpercentage gets the value of Logfillpercentage for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyLogfillpercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("Logfillpercentage")
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

// SetLogwritesPersec sets the value of LogwritesPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyLogwritesPersec(value uint64) (err error) {
	return instance.SetProperty("LogwritesPersec", (value))
}

// GetLogwritesPersec gets the value of LogwritesPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyLogwritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogwritesPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowTierDestagedContainerFillRatioPercent sets the value of SlowTierDestagedContainerFillRatioPercent for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertySlowTierDestagedContainerFillRatioPercent(value uint64) (err error) {
	return instance.SetProperty("SlowTierDestagedContainerFillRatioPercent", (value))
}

// GetSlowTierDestagedContainerFillRatioPercent gets the value of SlowTierDestagedContainerFillRatioPercent for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertySlowTierDestagedContainerFillRatioPercent() (value uint64, err error) {
	retValue, err := instance.GetProperty("SlowTierDestagedContainerFillRatioPercent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowTierDestagedContainerFillRatioPercent_Base sets the value of SlowTierDestagedContainerFillRatioPercent_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertySlowTierDestagedContainerFillRatioPercent_Base(value uint32) (err error) {
	return instance.SetProperty("SlowTierDestagedContainerFillRatioPercent_Base", (value))
}

// GetSlowTierDestagedContainerFillRatioPercent_Base gets the value of SlowTierDestagedContainerFillRatioPercent_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertySlowTierDestagedContainerFillRatioPercent_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("SlowTierDestagedContainerFillRatioPercent_Base")
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

// SetSlowtierdestagereadlatency100ns sets the value of Slowtierdestagereadlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertySlowtierdestagereadlatency100ns(value uint64) (err error) {
	return instance.SetProperty("Slowtierdestagereadlatency100ns", (value))
}

// GetSlowtierdestagereadlatency100ns gets the value of Slowtierdestagereadlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertySlowtierdestagereadlatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Slowtierdestagereadlatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowtierdestagereadlatency100ns_Base sets the value of Slowtierdestagereadlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertySlowtierdestagereadlatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Slowtierdestagereadlatency100ns_Base", (value))
}

// GetSlowtierdestagereadlatency100ns_Base gets the value of Slowtierdestagereadlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertySlowtierdestagereadlatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Slowtierdestagereadlatency100ns_Base")
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

// SetSlowtierdestagewritelatency100ns sets the value of Slowtierdestagewritelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertySlowtierdestagewritelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Slowtierdestagewritelatency100ns", (value))
}

// GetSlowtierdestagewritelatency100ns gets the value of Slowtierdestagewritelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertySlowtierdestagewritelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Slowtierdestagewritelatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowtierdestagewritelatency100ns_Base sets the value of Slowtierdestagewritelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertySlowtierdestagewritelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Slowtierdestagewritelatency100ns_Base", (value))
}

// GetSlowtierdestagewritelatency100ns_Base gets the value of Slowtierdestagewritelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertySlowtierdestagewritelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Slowtierdestagewritelatency100ns_Base")
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

// SetTotalAllocationofClustersPersec sets the value of TotalAllocationofClustersPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTotalAllocationofClustersPersec(value uint64) (err error) {
	return instance.SetProperty("TotalAllocationofClustersPersec", (value))
}

// GetTotalAllocationofClustersPersec gets the value of TotalAllocationofClustersPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTotalAllocationofClustersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalAllocationofClustersPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTreeupdateexclusivelockslatency100ns sets the value of Treeupdateexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTreeupdateexclusivelockslatency100ns(value uint64) (err error) {
	return instance.SetProperty("Treeupdateexclusivelockslatency100ns", (value))
}

// GetTreeupdateexclusivelockslatency100ns gets the value of Treeupdateexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTreeupdateexclusivelockslatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Treeupdateexclusivelockslatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTreeupdateexclusivelockslatency100ns_Base sets the value of Treeupdateexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTreeupdateexclusivelockslatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Treeupdateexclusivelockslatency100ns_Base", (value))
}

// GetTreeupdateexclusivelockslatency100ns_Base gets the value of Treeupdateexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTreeupdateexclusivelockslatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Treeupdateexclusivelockslatency100ns_Base")
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

// SetTreeupdatelatency100ns sets the value of Treeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTreeupdatelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Treeupdatelatency100ns", (value))
}

// GetTreeupdatelatency100ns gets the value of Treeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTreeupdatelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Treeupdatelatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTreeupdatelatency100ns_Base sets the value of Treeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTreeupdatelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Treeupdatelatency100ns_Base", (value))
}

// GetTreeupdatelatency100ns_Base gets the value of Treeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTreeupdatelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Treeupdatelatency100ns_Base")
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

// SetTreeupdatesPersec sets the value of TreeupdatesPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTreeupdatesPersec(value uint64) (err error) {
	return instance.SetProperty("TreeupdatesPersec", (value))
}

// GetTreeupdatesPersec gets the value of TreeupdatesPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTreeupdatesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TreeupdatesPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTrimlatency100ns sets the value of Trimlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTrimlatency100ns(value uint64) (err error) {
	return instance.SetProperty("Trimlatency100ns", (value))
}

// GetTrimlatency100ns gets the value of Trimlatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTrimlatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Trimlatency100ns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTrimlatency100ns_Base sets the value of Trimlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) SetPropertyTrimlatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Trimlatency100ns_Base", (value))
}

// GetTrimlatency100ns_Base gets the value of Trimlatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFS) GetPropertyTrimlatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Trimlatency100ns_Base")
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
