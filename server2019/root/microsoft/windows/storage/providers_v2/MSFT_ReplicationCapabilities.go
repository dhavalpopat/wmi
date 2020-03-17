// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_ReplicationCapabilities struct
type MSFT_ReplicationCapabilities struct {
	MSFT_StorageObject

	// Default value for recovery point
	DefaultRecoveryPointObjective uint32

	// Enumeration indicating what operations will be executed as asynchronous jobs. If an operation is included in both this and SupportedSynchronousActions properties then the underlying implementation is indicating that it may or may not create a job.
	///Note: the following methods are not supported asynchronously, hence the gap between 11 and 19:
	///	 - CreateGroup
	///	 - DeleteGroup
	///	 - AddMembers
	///	 - RemoveMembers
	///	 - AddReplicationEntity
	///	 - AddServiceAccessPoint
	///	 - AddSharedSecret.
	SupportedAsynchronousActions []ReplicationCapabilities_SupportedAsynchronousActions

	// An array of supported features of partition objects for replication.
	SupportedLogVolumeFeatures []ReplicationCapabilities_SupportedLogVolumeFeatures

	// Maximum log size in bytes supported for replication.
	SupportedMaximumLogSize uint64

	// Minimum log size in bytes supported for replication.
	SupportedMinimumLogSize uint64

	// Enumeration indicating the supported object types associated with these replication capabilities.
	SupportedObjectTypes []ReplicationCapabilities_SupportedObjectTypes

	// An array of supported features of partition objects for replication.
	SupportedReplicatedPartitionFeatures []ReplicationCapabilities_SupportedReplicatedPartitionFeatures

	// Enumeration indicating the supported SyncType/Mode/Local-or-Remote combinations.
	SupportedReplicationTypes []ReplicationCapabilities_SupportedReplicationTypes

	// Enumeration indicating what operations will be executed synchronously -- without the creation of a job. If an operation is included in both this property and SupportedAsynchronousActions then the underlying implementation is indicating that it may or may not create a job.
	///Note: the following methods are not supported asynchronously:
	///	 - CreateGroup
	///	 - DeleteGroup
	///	 - AddMembers
	///	 - RemoveMembers
	///	 - AddReplicationEntity
	///	 - AddServiceAccessPoint
	///	 - AddSharedSecret.
	SupportedSynchronousActions []ReplicationCapabilities_SupportedSynchronousActions

	// Indicates if CreateReplicationShip operation is supported
	SupportsCreateReplicationRelationshipMethod bool

	// Indicates if empty Replicaiotn Groups are allowed
	SupportsEmptyReplicationGroup bool

	// Indicates if this is a fully discovered model
	SupportsFullDiscovery bool

	// Indicates if Replication Groups is supported
	SupportsReplicationGroup bool
}

// SetDefaultRecoveryPointObjective sets the value of DefaultRecoveryPointObjective for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertyDefaultRecoveryPointObjective(value uint32) (err error) {
	return instance.SetProperty("DefaultRecoveryPointObjective", value)
}

// GetDefaultRecoveryPointObjective gets the value of DefaultRecoveryPointObjective for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertyDefaultRecoveryPointObjective() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultRecoveryPointObjective")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedAsynchronousActions sets the value of SupportedAsynchronousActions for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedAsynchronousActions(value []ReplicationCapabilities_SupportedAsynchronousActions) (err error) {
	return instance.SetProperty("SupportedAsynchronousActions", value)
}

// GetSupportedAsynchronousActions gets the value of SupportedAsynchronousActions for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedAsynchronousActions() (value []ReplicationCapabilities_SupportedAsynchronousActions, err error) {
	retValue, err := instance.GetProperty("SupportedAsynchronousActions")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationCapabilities_SupportedAsynchronousActions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedLogVolumeFeatures sets the value of SupportedLogVolumeFeatures for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedLogVolumeFeatures(value []ReplicationCapabilities_SupportedLogVolumeFeatures) (err error) {
	return instance.SetProperty("SupportedLogVolumeFeatures", value)
}

// GetSupportedLogVolumeFeatures gets the value of SupportedLogVolumeFeatures for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedLogVolumeFeatures() (value []ReplicationCapabilities_SupportedLogVolumeFeatures, err error) {
	retValue, err := instance.GetProperty("SupportedLogVolumeFeatures")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationCapabilities_SupportedLogVolumeFeatures)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedMaximumLogSize sets the value of SupportedMaximumLogSize for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedMaximumLogSize(value uint64) (err error) {
	return instance.SetProperty("SupportedMaximumLogSize", value)
}

// GetSupportedMaximumLogSize gets the value of SupportedMaximumLogSize for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedMaximumLogSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SupportedMaximumLogSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedMinimumLogSize sets the value of SupportedMinimumLogSize for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedMinimumLogSize(value uint64) (err error) {
	return instance.SetProperty("SupportedMinimumLogSize", value)
}

// GetSupportedMinimumLogSize gets the value of SupportedMinimumLogSize for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedMinimumLogSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SupportedMinimumLogSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedObjectTypes sets the value of SupportedObjectTypes for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedObjectTypes(value []ReplicationCapabilities_SupportedObjectTypes) (err error) {
	return instance.SetProperty("SupportedObjectTypes", value)
}

// GetSupportedObjectTypes gets the value of SupportedObjectTypes for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedObjectTypes() (value []ReplicationCapabilities_SupportedObjectTypes, err error) {
	retValue, err := instance.GetProperty("SupportedObjectTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationCapabilities_SupportedObjectTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedReplicatedPartitionFeatures sets the value of SupportedReplicatedPartitionFeatures for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedReplicatedPartitionFeatures(value []ReplicationCapabilities_SupportedReplicatedPartitionFeatures) (err error) {
	return instance.SetProperty("SupportedReplicatedPartitionFeatures", value)
}

// GetSupportedReplicatedPartitionFeatures gets the value of SupportedReplicatedPartitionFeatures for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedReplicatedPartitionFeatures() (value []ReplicationCapabilities_SupportedReplicatedPartitionFeatures, err error) {
	retValue, err := instance.GetProperty("SupportedReplicatedPartitionFeatures")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationCapabilities_SupportedReplicatedPartitionFeatures)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedReplicationTypes sets the value of SupportedReplicationTypes for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedReplicationTypes(value []ReplicationCapabilities_SupportedReplicationTypes) (err error) {
	return instance.SetProperty("SupportedReplicationTypes", value)
}

// GetSupportedReplicationTypes gets the value of SupportedReplicationTypes for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedReplicationTypes() (value []ReplicationCapabilities_SupportedReplicationTypes, err error) {
	retValue, err := instance.GetProperty("SupportedReplicationTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationCapabilities_SupportedReplicationTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedSynchronousActions sets the value of SupportedSynchronousActions for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportedSynchronousActions(value []ReplicationCapabilities_SupportedSynchronousActions) (err error) {
	return instance.SetProperty("SupportedSynchronousActions", value)
}

// GetSupportedSynchronousActions gets the value of SupportedSynchronousActions for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportedSynchronousActions() (value []ReplicationCapabilities_SupportedSynchronousActions, err error) {
	retValue, err := instance.GetProperty("SupportedSynchronousActions")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationCapabilities_SupportedSynchronousActions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsCreateReplicationRelationshipMethod sets the value of SupportsCreateReplicationRelationshipMethod for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportsCreateReplicationRelationshipMethod(value bool) (err error) {
	return instance.SetProperty("SupportsCreateReplicationRelationshipMethod", value)
}

// GetSupportsCreateReplicationRelationshipMethod gets the value of SupportsCreateReplicationRelationshipMethod for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportsCreateReplicationRelationshipMethod() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCreateReplicationRelationshipMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsEmptyReplicationGroup sets the value of SupportsEmptyReplicationGroup for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportsEmptyReplicationGroup(value bool) (err error) {
	return instance.SetProperty("SupportsEmptyReplicationGroup", value)
}

// GetSupportsEmptyReplicationGroup gets the value of SupportsEmptyReplicationGroup for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportsEmptyReplicationGroup() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsEmptyReplicationGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsFullDiscovery sets the value of SupportsFullDiscovery for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportsFullDiscovery(value bool) (err error) {
	return instance.SetProperty("SupportsFullDiscovery", value)
}

// GetSupportsFullDiscovery gets the value of SupportsFullDiscovery for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportsFullDiscovery() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsFullDiscovery")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsReplicationGroup sets the value of SupportsReplicationGroup for the instance
func (instance *MSFT_ReplicationCapabilities) SetPropertySupportsReplicationGroup(value bool) (err error) {
	return instance.SetProperty("SupportsReplicationGroup", value)
}

// GetSupportsReplicationGroup gets the value of SupportsReplicationGroup for the instance
func (instance *MSFT_ReplicationCapabilities) GetPropertySupportsReplicationGroup() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsReplicationGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedOperations" type="uint16 []"></param>
func (instance *MSFT_ReplicationCapabilities) GetSupportedOperations( /* IN */ ReplicationType uint16,
	/* OUT */ SupportedOperations []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedOperations", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedGroupOperations" type="uint16 []"></param>
func (instance *MSFT_ReplicationCapabilities) GetSupportedGroupOperations( /* IN */ ReplicationType uint16,
	/* OUT */ SupportedGroupOperations []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedGroupOperations", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="Features" type="uint16 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationCapabilities) GetSupportedFeatures( /* IN */ ReplicationType uint16,
	/* OUT */ Features []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedFeatures", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="GroupFeatures" type="uint16 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationCapabilities) GetSupportedGroupFeatures( /* IN */ ReplicationType uint16,
	/* OUT */ GroupFeatures []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedGroupFeatures", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedCopyStates" type="uint16 []"></param>
func (instance *MSFT_ReplicationCapabilities) GetSupportedCopyStates( /* IN */ ReplicationType uint16,
	/* OUT */ SupportedCopyStates []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedCopyStates", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedCopyStates" type="uint16 []"></param>
func (instance *MSFT_ReplicationCapabilities) GetSupportedGroupCopyStates( /* IN */ ReplicationType uint16,
	/* OUT */ SupportedCopyStates []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedGroupCopyStates", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationType" type="uint16 "></param>

// <param name="DefaultRecoveryPoint" type="uint32 "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="RecoveryPointIndicator" type="uint16 "></param>
// <param name="RecoveryPointValues" type="uint32 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationCapabilities) GetRecoveryPointData( /* IN */ ReplicationType uint16,
	/* OUT */ DefaultRecoveryPoint uint32,
	/* OUT */ RecoveryPointValues []uint32,
	/* OUT */ RecoveryPointIndicator uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetRecoveryPointData", ReplicationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
