// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetOffloadGlobalSetting struct
type MSFT_NetOffloadGlobalSetting struct {
	MSFT_NetSettingData

	//
	Chimney uint8

	//
	NetworkDirect uint8

	//
	NetworkDirectAcrossIPSubnets uint8

	//
	PacketCoalescingFilter uint8

	//
	ReceiveSegmentCoalescing uint8

	//
	ReceiveSideScaling uint8

	//
	TaskOffload uint8
}

// SetChimney sets the value of Chimney for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyChimney(value uint8) (err error) {
	return instance.SetProperty("Chimney", value)
}

// GetChimney gets the value of Chimney for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyChimney() (value uint8, err error) {
	retValue, err := instance.GetProperty("Chimney")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkDirect sets the value of NetworkDirect for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyNetworkDirect(value uint8) (err error) {
	return instance.SetProperty("NetworkDirect", value)
}

// GetNetworkDirect gets the value of NetworkDirect for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyNetworkDirect() (value uint8, err error) {
	retValue, err := instance.GetProperty("NetworkDirect")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkDirectAcrossIPSubnets sets the value of NetworkDirectAcrossIPSubnets for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyNetworkDirectAcrossIPSubnets(value uint8) (err error) {
	return instance.SetProperty("NetworkDirectAcrossIPSubnets", value)
}

// GetNetworkDirectAcrossIPSubnets gets the value of NetworkDirectAcrossIPSubnets for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyNetworkDirectAcrossIPSubnets() (value uint8, err error) {
	retValue, err := instance.GetProperty("NetworkDirectAcrossIPSubnets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketCoalescingFilter sets the value of PacketCoalescingFilter for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyPacketCoalescingFilter(value uint8) (err error) {
	return instance.SetProperty("PacketCoalescingFilter", value)
}

// GetPacketCoalescingFilter gets the value of PacketCoalescingFilter for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyPacketCoalescingFilter() (value uint8, err error) {
	retValue, err := instance.GetProperty("PacketCoalescingFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveSegmentCoalescing sets the value of ReceiveSegmentCoalescing for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyReceiveSegmentCoalescing(value uint8) (err error) {
	return instance.SetProperty("ReceiveSegmentCoalescing", value)
}

// GetReceiveSegmentCoalescing gets the value of ReceiveSegmentCoalescing for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyReceiveSegmentCoalescing() (value uint8, err error) {
	retValue, err := instance.GetProperty("ReceiveSegmentCoalescing")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveSideScaling sets the value of ReceiveSideScaling for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyReceiveSideScaling(value uint8) (err error) {
	return instance.SetProperty("ReceiveSideScaling", value)
}

// GetReceiveSideScaling gets the value of ReceiveSideScaling for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyReceiveSideScaling() (value uint8, err error) {
	retValue, err := instance.GetProperty("ReceiveSideScaling")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskOffload sets the value of TaskOffload for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyTaskOffload(value uint8) (err error) {
	return instance.SetProperty("TaskOffload", value)
}

// GetTaskOffload gets the value of TaskOffload for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyTaskOffload() (value uint8, err error) {
	retValue, err := instance.GetProperty("TaskOffload")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
