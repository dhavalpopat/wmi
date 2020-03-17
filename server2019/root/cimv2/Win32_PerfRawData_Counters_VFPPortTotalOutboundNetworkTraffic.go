// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic struct
type Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic struct {
	Win32_PerfRawData

	//
	TotalOutboundBytes uint64

	//
	TotalOutboundForwardedMulticastPackets uint64

	//
	TotalOutboundForwardedUnicastPackets uint64

	//
	TotalOutboundGFTBytes uint64

	//
	TotalOutboundGFTCopyFINPackets uint64

	//
	TotalOutboundGFTCopyPackets uint64

	//
	TotalOutboundGFTCopyResetPackets uint64

	//
	TotalOutboundGFTExceptionPackets uint64

	//
	TotalOutboundGFTExceptionUFOffloadBlockedPackets uint64

	//
	TotalOutboundGFTExceptionUFOffloadDeferredPackets uint64

	//
	TotalOutboundGFTExceptionUFOffloadedTCPPackets uint64

	//
	TotalOutboundGFTExceptionUFOffloadedUDPPackets uint64

	//
	TotalOutboundGFTExceptionUFOffloadFailedPackets uint64

	//
	TotalOutboundGFTExceptionUFOffloadPendingPackets uint64

	//
	TotalOutboundGFTExceptionUFPackets uint64

	//
	TotalOutboundGFTRetryAwaitingPackets uint64

	//
	TotalOutboundGftTotalPackets uint64

	//
	TotalOutboundHairPinnedPackets uint64

	//
	TotalOutboundInterceptedPackets uint64

	//
	TotalOutboundMissedInterceptedPackets uint64

	//
	TotalOutboundNonIPPackets uint64

	//
	TotalOutboundPackets uint64

	//
	TotalOutboundPendingPackets uint64

	//
	TotalOutboundTCPSYNACKPackets uint64

	//
	TotalOutboundTCPSYNPackets uint64

	//
	TotalOutboundThrottledPackets uint64

	//
	TotalOutboundUnicastForwardedGFTExceptionPackets uint64
}

// SetTotalOutboundBytes sets the value of TotalOutboundBytes for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundBytes(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundBytes", value)
}

// GetTotalOutboundBytes gets the value of TotalOutboundBytes for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundForwardedMulticastPackets sets the value of TotalOutboundForwardedMulticastPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundForwardedMulticastPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundForwardedMulticastPackets", value)
}

// GetTotalOutboundForwardedMulticastPackets gets the value of TotalOutboundForwardedMulticastPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundForwardedMulticastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundForwardedMulticastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundForwardedUnicastPackets sets the value of TotalOutboundForwardedUnicastPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundForwardedUnicastPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundForwardedUnicastPackets", value)
}

// GetTotalOutboundForwardedUnicastPackets gets the value of TotalOutboundForwardedUnicastPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundForwardedUnicastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundForwardedUnicastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTBytes sets the value of TotalOutboundGFTBytes for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTBytes(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTBytes", value)
}

// GetTotalOutboundGFTBytes gets the value of TotalOutboundGFTBytes for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTCopyFINPackets sets the value of TotalOutboundGFTCopyFINPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTCopyFINPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTCopyFINPackets", value)
}

// GetTotalOutboundGFTCopyFINPackets gets the value of TotalOutboundGFTCopyFINPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTCopyFINPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTCopyFINPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTCopyPackets sets the value of TotalOutboundGFTCopyPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTCopyPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTCopyPackets", value)
}

// GetTotalOutboundGFTCopyPackets gets the value of TotalOutboundGFTCopyPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTCopyPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTCopyPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTCopyResetPackets sets the value of TotalOutboundGFTCopyResetPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTCopyResetPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTCopyResetPackets", value)
}

// GetTotalOutboundGFTCopyResetPackets gets the value of TotalOutboundGFTCopyResetPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTCopyResetPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTCopyResetPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionPackets sets the value of TotalOutboundGFTExceptionPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionPackets", value)
}

// GetTotalOutboundGFTExceptionPackets gets the value of TotalOutboundGFTExceptionPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFOffloadBlockedPackets sets the value of TotalOutboundGFTExceptionUFOffloadBlockedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFOffloadBlockedPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFOffloadBlockedPackets", value)
}

// GetTotalOutboundGFTExceptionUFOffloadBlockedPackets gets the value of TotalOutboundGFTExceptionUFOffloadBlockedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFOffloadBlockedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFOffloadBlockedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFOffloadDeferredPackets sets the value of TotalOutboundGFTExceptionUFOffloadDeferredPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFOffloadDeferredPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFOffloadDeferredPackets", value)
}

// GetTotalOutboundGFTExceptionUFOffloadDeferredPackets gets the value of TotalOutboundGFTExceptionUFOffloadDeferredPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFOffloadDeferredPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFOffloadDeferredPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFOffloadedTCPPackets sets the value of TotalOutboundGFTExceptionUFOffloadedTCPPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFOffloadedTCPPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFOffloadedTCPPackets", value)
}

// GetTotalOutboundGFTExceptionUFOffloadedTCPPackets gets the value of TotalOutboundGFTExceptionUFOffloadedTCPPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFOffloadedTCPPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFOffloadedTCPPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFOffloadedUDPPackets sets the value of TotalOutboundGFTExceptionUFOffloadedUDPPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFOffloadedUDPPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFOffloadedUDPPackets", value)
}

// GetTotalOutboundGFTExceptionUFOffloadedUDPPackets gets the value of TotalOutboundGFTExceptionUFOffloadedUDPPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFOffloadedUDPPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFOffloadedUDPPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFOffloadFailedPackets sets the value of TotalOutboundGFTExceptionUFOffloadFailedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFOffloadFailedPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFOffloadFailedPackets", value)
}

// GetTotalOutboundGFTExceptionUFOffloadFailedPackets gets the value of TotalOutboundGFTExceptionUFOffloadFailedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFOffloadFailedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFOffloadFailedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFOffloadPendingPackets sets the value of TotalOutboundGFTExceptionUFOffloadPendingPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFOffloadPendingPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFOffloadPendingPackets", value)
}

// GetTotalOutboundGFTExceptionUFOffloadPendingPackets gets the value of TotalOutboundGFTExceptionUFOffloadPendingPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFOffloadPendingPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFOffloadPendingPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTExceptionUFPackets sets the value of TotalOutboundGFTExceptionUFPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTExceptionUFPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTExceptionUFPackets", value)
}

// GetTotalOutboundGFTExceptionUFPackets gets the value of TotalOutboundGFTExceptionUFPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTExceptionUFPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTExceptionUFPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGFTRetryAwaitingPackets sets the value of TotalOutboundGFTRetryAwaitingPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGFTRetryAwaitingPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGFTRetryAwaitingPackets", value)
}

// GetTotalOutboundGFTRetryAwaitingPackets gets the value of TotalOutboundGFTRetryAwaitingPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGFTRetryAwaitingPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGFTRetryAwaitingPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundGftTotalPackets sets the value of TotalOutboundGftTotalPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundGftTotalPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundGftTotalPackets", value)
}

// GetTotalOutboundGftTotalPackets gets the value of TotalOutboundGftTotalPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundGftTotalPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundGftTotalPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundHairPinnedPackets sets the value of TotalOutboundHairPinnedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundHairPinnedPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundHairPinnedPackets", value)
}

// GetTotalOutboundHairPinnedPackets gets the value of TotalOutboundHairPinnedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundHairPinnedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundHairPinnedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundInterceptedPackets sets the value of TotalOutboundInterceptedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundInterceptedPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundInterceptedPackets", value)
}

// GetTotalOutboundInterceptedPackets gets the value of TotalOutboundInterceptedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundInterceptedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundInterceptedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundMissedInterceptedPackets sets the value of TotalOutboundMissedInterceptedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundMissedInterceptedPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundMissedInterceptedPackets", value)
}

// GetTotalOutboundMissedInterceptedPackets gets the value of TotalOutboundMissedInterceptedPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundMissedInterceptedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundMissedInterceptedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundNonIPPackets sets the value of TotalOutboundNonIPPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundNonIPPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundNonIPPackets", value)
}

// GetTotalOutboundNonIPPackets gets the value of TotalOutboundNonIPPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundNonIPPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundNonIPPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundPackets sets the value of TotalOutboundPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundPackets", value)
}

// GetTotalOutboundPackets gets the value of TotalOutboundPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundPendingPackets sets the value of TotalOutboundPendingPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundPendingPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundPendingPackets", value)
}

// GetTotalOutboundPendingPackets gets the value of TotalOutboundPendingPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundPendingPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundPendingPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundTCPSYNACKPackets sets the value of TotalOutboundTCPSYNACKPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundTCPSYNACKPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundTCPSYNACKPackets", value)
}

// GetTotalOutboundTCPSYNACKPackets gets the value of TotalOutboundTCPSYNACKPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundTCPSYNACKPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundTCPSYNACKPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundTCPSYNPackets sets the value of TotalOutboundTCPSYNPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundTCPSYNPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundTCPSYNPackets", value)
}

// GetTotalOutboundTCPSYNPackets gets the value of TotalOutboundTCPSYNPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundTCPSYNPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundTCPSYNPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundThrottledPackets sets the value of TotalOutboundThrottledPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundThrottledPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundThrottledPackets", value)
}

// GetTotalOutboundThrottledPackets gets the value of TotalOutboundThrottledPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundThrottledPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundThrottledPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundUnicastForwardedGFTExceptionPackets sets the value of TotalOutboundUnicastForwardedGFTExceptionPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) SetPropertyTotalOutboundUnicastForwardedGFTExceptionPackets(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundUnicastForwardedGFTExceptionPackets", value)
}

// GetTotalOutboundUnicastForwardedGFTExceptionPackets gets the value of TotalOutboundUnicastForwardedGFTExceptionPackets for the instance
func (instance *Win32_PerfRawData_Counters_VFPPortTotalOutboundNetworkTraffic) GetPropertyTotalOutboundUnicastForwardedGFTExceptionPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundUnicastForwardedGFTExceptionPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
