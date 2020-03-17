// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_PerfProc_Thread struct
type Win32_PerfRawData_PerfProc_Thread struct {
	Win32_PerfRawData

	//
	ContextSwitchesPersec uint32

	//
	ElapsedTime uint64

	//
	IDProcess uint32

	//
	IDThread uint32

	//
	PercentPrivilegedTime uint64

	//
	PercentProcessorTime uint64

	//
	PercentUserTime uint64

	//
	PriorityBase uint32

	//
	PriorityCurrent uint32

	//
	StartAddress uint32

	//
	ThreadState uint32

	//
	ThreadWaitReason uint32
}

// SetContextSwitchesPersec sets the value of ContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyContextSwitchesPersec(value uint32) (err error) {
	return instance.SetProperty("ContextSwitchesPersec", value)
}

// GetContextSwitchesPersec gets the value of ContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyContextSwitchesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContextSwitchesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyElapsedTime(value uint64) (err error) {
	return instance.SetProperty("ElapsedTime", value)
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyElapsedTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIDProcess sets the value of IDProcess for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyIDProcess(value uint32) (err error) {
	return instance.SetProperty("IDProcess", value)
}

// GetIDProcess gets the value of IDProcess for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyIDProcess() (value uint32, err error) {
	retValue, err := instance.GetProperty("IDProcess")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIDThread sets the value of IDThread for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyIDThread(value uint32) (err error) {
	return instance.SetProperty("IDThread", value)
}

// GetIDThread gets the value of IDThread for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyIDThread() (value uint32, err error) {
	retValue, err := instance.GetProperty("IDThread")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentPrivilegedTime sets the value of PercentPrivilegedTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyPercentPrivilegedTime(value uint64) (err error) {
	return instance.SetProperty("PercentPrivilegedTime", value)
}

// GetPercentPrivilegedTime gets the value of PercentPrivilegedTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyPercentPrivilegedTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentPrivilegedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentProcessorTime sets the value of PercentProcessorTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyPercentProcessorTime(value uint64) (err error) {
	return instance.SetProperty("PercentProcessorTime", value)
}

// GetPercentProcessorTime gets the value of PercentProcessorTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyPercentProcessorTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentProcessorTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentUserTime sets the value of PercentUserTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyPercentUserTime(value uint64) (err error) {
	return instance.SetProperty("PercentUserTime", value)
}

// GetPercentUserTime gets the value of PercentUserTime for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyPercentUserTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentUserTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriorityBase sets the value of PriorityBase for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyPriorityBase(value uint32) (err error) {
	return instance.SetProperty("PriorityBase", value)
}

// GetPriorityBase gets the value of PriorityBase for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyPriorityBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityBase")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriorityCurrent sets the value of PriorityCurrent for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyPriorityCurrent(value uint32) (err error) {
	return instance.SetProperty("PriorityCurrent", value)
}

// GetPriorityCurrent gets the value of PriorityCurrent for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyPriorityCurrent() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityCurrent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartAddress sets the value of StartAddress for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyStartAddress(value uint32) (err error) {
	return instance.SetProperty("StartAddress", value)
}

// GetStartAddress gets the value of StartAddress for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyStartAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadState sets the value of ThreadState for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyThreadState(value uint32) (err error) {
	return instance.SetProperty("ThreadState", value)
}

// GetThreadState gets the value of ThreadState for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyThreadState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadWaitReason sets the value of ThreadWaitReason for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) SetPropertyThreadWaitReason(value uint32) (err error) {
	return instance.SetProperty("ThreadWaitReason", value)
}

// GetThreadWaitReason gets the value of ThreadWaitReason for the instance
func (instance *Win32_PerfRawData_PerfProc_Thread) GetPropertyThreadWaitReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadWaitReason")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
