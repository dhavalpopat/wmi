// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_BIOSElement struct
type Msvm_BIOSElement struct {
	CIM_BIOSElement

	//
	BaseBoardSerialNumber string

	//
	BIOSGUID string

	//
	BIOSNumLock bool

	//
	BIOSSerialNumber string

	//
	BootOrder []uint16

	//
	BootPciExpress bool

	//
	BootPciExpressInstanceFilter string

	//
	ChassisAssetTag string

	//
	ChassisSerialNumber string

	//
	EnableHibernation bool
}

// SetBaseBoardSerialNumber sets the value of BaseBoardSerialNumber for the instance
func (instance *Msvm_BIOSElement) SetPropertyBaseBoardSerialNumber(value string) (err error) {
	return instance.SetProperty("BaseBoardSerialNumber", value)
}

// GetBaseBoardSerialNumber gets the value of BaseBoardSerialNumber for the instance
func (instance *Msvm_BIOSElement) GetPropertyBaseBoardSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BaseBoardSerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBIOSGUID sets the value of BIOSGUID for the instance
func (instance *Msvm_BIOSElement) SetPropertyBIOSGUID(value string) (err error) {
	return instance.SetProperty("BIOSGUID", value)
}

// GetBIOSGUID gets the value of BIOSGUID for the instance
func (instance *Msvm_BIOSElement) GetPropertyBIOSGUID() (value string, err error) {
	retValue, err := instance.GetProperty("BIOSGUID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBIOSNumLock sets the value of BIOSNumLock for the instance
func (instance *Msvm_BIOSElement) SetPropertyBIOSNumLock(value bool) (err error) {
	return instance.SetProperty("BIOSNumLock", value)
}

// GetBIOSNumLock gets the value of BIOSNumLock for the instance
func (instance *Msvm_BIOSElement) GetPropertyBIOSNumLock() (value bool, err error) {
	retValue, err := instance.GetProperty("BIOSNumLock")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBIOSSerialNumber sets the value of BIOSSerialNumber for the instance
func (instance *Msvm_BIOSElement) SetPropertyBIOSSerialNumber(value string) (err error) {
	return instance.SetProperty("BIOSSerialNumber", value)
}

// GetBIOSSerialNumber gets the value of BIOSSerialNumber for the instance
func (instance *Msvm_BIOSElement) GetPropertyBIOSSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BIOSSerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBootOrder sets the value of BootOrder for the instance
func (instance *Msvm_BIOSElement) SetPropertyBootOrder(value []uint16) (err error) {
	return instance.SetProperty("BootOrder", value)
}

// GetBootOrder gets the value of BootOrder for the instance
func (instance *Msvm_BIOSElement) GetPropertyBootOrder() (value []uint16, err error) {
	retValue, err := instance.GetProperty("BootOrder")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBootPciExpress sets the value of BootPciExpress for the instance
func (instance *Msvm_BIOSElement) SetPropertyBootPciExpress(value bool) (err error) {
	return instance.SetProperty("BootPciExpress", value)
}

// GetBootPciExpress gets the value of BootPciExpress for the instance
func (instance *Msvm_BIOSElement) GetPropertyBootPciExpress() (value bool, err error) {
	retValue, err := instance.GetProperty("BootPciExpress")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBootPciExpressInstanceFilter sets the value of BootPciExpressInstanceFilter for the instance
func (instance *Msvm_BIOSElement) SetPropertyBootPciExpressInstanceFilter(value string) (err error) {
	return instance.SetProperty("BootPciExpressInstanceFilter", value)
}

// GetBootPciExpressInstanceFilter gets the value of BootPciExpressInstanceFilter for the instance
func (instance *Msvm_BIOSElement) GetPropertyBootPciExpressInstanceFilter() (value string, err error) {
	retValue, err := instance.GetProperty("BootPciExpressInstanceFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChassisAssetTag sets the value of ChassisAssetTag for the instance
func (instance *Msvm_BIOSElement) SetPropertyChassisAssetTag(value string) (err error) {
	return instance.SetProperty("ChassisAssetTag", value)
}

// GetChassisAssetTag gets the value of ChassisAssetTag for the instance
func (instance *Msvm_BIOSElement) GetPropertyChassisAssetTag() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisAssetTag")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChassisSerialNumber sets the value of ChassisSerialNumber for the instance
func (instance *Msvm_BIOSElement) SetPropertyChassisSerialNumber(value string) (err error) {
	return instance.SetProperty("ChassisSerialNumber", value)
}

// GetChassisSerialNumber gets the value of ChassisSerialNumber for the instance
func (instance *Msvm_BIOSElement) GetPropertyChassisSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisSerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableHibernation sets the value of EnableHibernation for the instance
func (instance *Msvm_BIOSElement) SetPropertyEnableHibernation(value bool) (err error) {
	return instance.SetProperty("EnableHibernation", value)
}

// GetEnableHibernation gets the value of EnableHibernation for the instance
func (instance *Msvm_BIOSElement) GetPropertyEnableHibernation() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableHibernation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
