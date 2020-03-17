// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache struct
type Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache struct {
	Win32_PerfFormattedData

	//
	CacheIOReadBytes uint64

	//
	CacheIOReadBytesPersec uint64

	//
	CacheRead uint64

	//
	CacheReadPerSec uint64

	//
	CacheSizeConfigured uint64

	//
	CacheSizeCurrent uint64

	//
	CacheState uint64

	//
	DiskIOReadBytes uint64

	//
	DiskIOReadBytesPerSec uint64

	//
	DiskIOReads uint64

	//
	DiskIOReadsPerSec uint64

	//
	IOReadBytes uint64

	//
	IOReadBytesPerSec uint64

	//
	IOReads uint64

	//
	IOReadsPerSec uint64

	//
	LRUCacheSizeCurrent uint64

	//
	LRUCacheSizeTarget uint64

	//
	PartialRead uint64

	//
	PartialReadPersec uint64

	//
	PercentCacheValid uint64

	//
	ValidCacheSize uint64
}

// SetCacheIOReadBytes sets the value of CacheIOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheIOReadBytes(value uint64) (err error) {
	return instance.SetProperty("CacheIOReadBytes", value)
}

// GetCacheIOReadBytes gets the value of CacheIOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheIOReadBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheIOReadBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheIOReadBytesPersec sets the value of CacheIOReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheIOReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("CacheIOReadBytesPersec", value)
}

// GetCacheIOReadBytesPersec gets the value of CacheIOReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheIOReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheIOReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheRead sets the value of CacheRead for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheRead(value uint64) (err error) {
	return instance.SetProperty("CacheRead", value)
}

// GetCacheRead gets the value of CacheRead for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheReadPerSec sets the value of CacheReadPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheReadPerSec(value uint64) (err error) {
	return instance.SetProperty("CacheReadPerSec", value)
}

// GetCacheReadPerSec gets the value of CacheReadPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheReadPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheReadPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheSizeConfigured sets the value of CacheSizeConfigured for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheSizeConfigured(value uint64) (err error) {
	return instance.SetProperty("CacheSizeConfigured", value)
}

// GetCacheSizeConfigured gets the value of CacheSizeConfigured for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheSizeConfigured() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheSizeConfigured")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheSizeCurrent sets the value of CacheSizeCurrent for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheSizeCurrent(value uint64) (err error) {
	return instance.SetProperty("CacheSizeCurrent", value)
}

// GetCacheSizeCurrent gets the value of CacheSizeCurrent for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheSizeCurrent() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheSizeCurrent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheState sets the value of CacheState for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyCacheState(value uint64) (err error) {
	return instance.SetProperty("CacheState", value)
}

// GetCacheState gets the value of CacheState for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyCacheState() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskIOReadBytes sets the value of DiskIOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyDiskIOReadBytes(value uint64) (err error) {
	return instance.SetProperty("DiskIOReadBytes", value)
}

// GetDiskIOReadBytes gets the value of DiskIOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyDiskIOReadBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskIOReadBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskIOReadBytesPerSec sets the value of DiskIOReadBytesPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyDiskIOReadBytesPerSec(value uint64) (err error) {
	return instance.SetProperty("DiskIOReadBytesPerSec", value)
}

// GetDiskIOReadBytesPerSec gets the value of DiskIOReadBytesPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyDiskIOReadBytesPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskIOReadBytesPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskIOReads sets the value of DiskIOReads for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyDiskIOReads(value uint64) (err error) {
	return instance.SetProperty("DiskIOReads", value)
}

// GetDiskIOReads gets the value of DiskIOReads for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyDiskIOReads() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskIOReads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskIOReadsPerSec sets the value of DiskIOReadsPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyDiskIOReadsPerSec(value uint64) (err error) {
	return instance.SetProperty("DiskIOReadsPerSec", value)
}

// GetDiskIOReadsPerSec gets the value of DiskIOReadsPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyDiskIOReadsPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskIOReadsPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytes sets the value of IOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyIOReadBytes(value uint64) (err error) {
	return instance.SetProperty("IOReadBytes", value)
}

// GetIOReadBytes gets the value of IOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyIOReadBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytesPerSec sets the value of IOReadBytesPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyIOReadBytesPerSec(value uint64) (err error) {
	return instance.SetProperty("IOReadBytesPerSec", value)
}

// GetIOReadBytesPerSec gets the value of IOReadBytesPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyIOReadBytesPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytesPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReads sets the value of IOReads for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyIOReads(value uint64) (err error) {
	return instance.SetProperty("IOReads", value)
}

// GetIOReads gets the value of IOReads for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyIOReads() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadsPerSec sets the value of IOReadsPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyIOReadsPerSec(value uint64) (err error) {
	return instance.SetProperty("IOReadsPerSec", value)
}

// GetIOReadsPerSec gets the value of IOReadsPerSec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyIOReadsPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadsPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLRUCacheSizeCurrent sets the value of LRUCacheSizeCurrent for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyLRUCacheSizeCurrent(value uint64) (err error) {
	return instance.SetProperty("LRUCacheSizeCurrent", value)
}

// GetLRUCacheSizeCurrent gets the value of LRUCacheSizeCurrent for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyLRUCacheSizeCurrent() (value uint64, err error) {
	retValue, err := instance.GetProperty("LRUCacheSizeCurrent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLRUCacheSizeTarget sets the value of LRUCacheSizeTarget for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyLRUCacheSizeTarget(value uint64) (err error) {
	return instance.SetProperty("LRUCacheSizeTarget", value)
}

// GetLRUCacheSizeTarget gets the value of LRUCacheSizeTarget for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyLRUCacheSizeTarget() (value uint64, err error) {
	retValue, err := instance.GetProperty("LRUCacheSizeTarget")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartialRead sets the value of PartialRead for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyPartialRead(value uint64) (err error) {
	return instance.SetProperty("PartialRead", value)
}

// GetPartialRead gets the value of PartialRead for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyPartialRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("PartialRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartialReadPersec sets the value of PartialReadPersec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyPartialReadPersec(value uint64) (err error) {
	return instance.SetProperty("PartialReadPersec", value)
}

// GetPartialReadPersec gets the value of PartialReadPersec for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyPartialReadPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PartialReadPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentCacheValid sets the value of PercentCacheValid for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyPercentCacheValid(value uint64) (err error) {
	return instance.SetProperty("PercentCacheValid", value)
}

// GetPercentCacheValid gets the value of PercentCacheValid for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyPercentCacheValid() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentCacheValid")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidCacheSize sets the value of ValidCacheSize for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) SetPropertyValidCacheSize(value uint64) (err error) {
	return instance.SetProperty("ValidCacheSize", value)
}

// GetValidCacheSize gets the value of ValidCacheSize for the instance
func (instance *Win32_PerfFormattedData_CsvFsPerfProvider_ClusterCSVFSBlockCache) GetPropertyValidCacheSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ValidCacheSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
