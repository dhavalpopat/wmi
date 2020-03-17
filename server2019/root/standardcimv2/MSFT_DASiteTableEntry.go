// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_DASiteTableEntry struct
type MSFT_DASiteTableEntry struct {
	MSFT_NetSettingData

	//
	ADSite string

	//
	EntryPointIPAddress string

	//
	EntryPointName string

	//
	EntryPointRange []string

	//
	GslbIP string

	//
	IPHttpsProfile string

	//
	PolicyStore string

	//
	State uint32

	//
	TeredoServerIP string
}

// SetADSite sets the value of ADSite for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyADSite(value string) (err error) {
	return instance.SetProperty("ADSite", value)
}

// GetADSite gets the value of ADSite for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyADSite() (value string, err error) {
	retValue, err := instance.GetProperty("ADSite")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntryPointIPAddress sets the value of EntryPointIPAddress for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyEntryPointIPAddress(value string) (err error) {
	return instance.SetProperty("EntryPointIPAddress", value)
}

// GetEntryPointIPAddress gets the value of EntryPointIPAddress for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyEntryPointIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("EntryPointIPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntryPointName sets the value of EntryPointName for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyEntryPointName(value string) (err error) {
	return instance.SetProperty("EntryPointName", value)
}

// GetEntryPointName gets the value of EntryPointName for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyEntryPointName() (value string, err error) {
	retValue, err := instance.GetProperty("EntryPointName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntryPointRange sets the value of EntryPointRange for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyEntryPointRange(value []string) (err error) {
	return instance.SetProperty("EntryPointRange", value)
}

// GetEntryPointRange gets the value of EntryPointRange for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyEntryPointRange() (value []string, err error) {
	retValue, err := instance.GetProperty("EntryPointRange")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGslbIP sets the value of GslbIP for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyGslbIP(value string) (err error) {
	return instance.SetProperty("GslbIP", value)
}

// GetGslbIP gets the value of GslbIP for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyGslbIP() (value string, err error) {
	retValue, err := instance.GetProperty("GslbIP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPHttpsProfile sets the value of IPHttpsProfile for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyIPHttpsProfile(value string) (err error) {
	return instance.SetProperty("IPHttpsProfile", value)
}

// GetIPHttpsProfile gets the value of IPHttpsProfile for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyIPHttpsProfile() (value string, err error) {
	retValue, err := instance.GetProperty("IPHttpsProfile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", value)
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTeredoServerIP sets the value of TeredoServerIP for the instance
func (instance *MSFT_DASiteTableEntry) SetPropertyTeredoServerIP(value string) (err error) {
	return instance.SetProperty("TeredoServerIP", value)
}

// GetTeredoServerIP gets the value of TeredoServerIP for the instance
func (instance *MSFT_DASiteTableEntry) GetPropertyTeredoServerIP() (value string, err error) {
	retValue, err := instance.GetProperty("TeredoServerIP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="EntryPointName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DASiteTableEntry) Enable( /* IN */ EntryPointName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable", EntryPointName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DASiteTableEntry) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewName" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="OutputObject" type="MSFT_DASiteTableEntry "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DASiteTableEntry) Rename( /* IN */ NewName string,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_DASiteTableEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Rename", NewName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GslbIP" type="bool "></param>
// <param name="IPHttpsProfile" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="TeredoServerIP" type="bool "></param>

// <param name="OutputObject" type="MSFT_DASiteTableEntry "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DASiteTableEntry) Reset( /* IN */ TeredoServerIP bool,
	/* IN */ IPHttpsProfile bool,
	/* IN */ GslbIP bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_DASiteTableEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", TeredoServerIP, IPHttpsProfile, GslbIP, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
