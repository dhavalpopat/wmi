// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

// MSFT_MpThreatCatalog struct
type MSFT_MpThreatCatalog struct {
	BaseStatus

	//
	CategoryID uint8

	//
	SeverityID uint8

	//
	ThreatID int64

	//
	ThreatName string

	//
	TypeID uint8
}

// SetCategoryID sets the value of CategoryID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyCategoryID(value uint8) (err error) {
	return instance.SetProperty("CategoryID", value)
}

// GetCategoryID gets the value of CategoryID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyCategoryID() (value uint8, err error) {
	retValue, err := instance.GetProperty("CategoryID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSeverityID sets the value of SeverityID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertySeverityID(value uint8) (err error) {
	return instance.SetProperty("SeverityID", value)
}

// GetSeverityID gets the value of SeverityID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertySeverityID() (value uint8, err error) {
	retValue, err := instance.GetProperty("SeverityID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreatID sets the value of ThreatID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyThreatID(value int64) (err error) {
	return instance.SetProperty("ThreatID", value)
}

// GetThreatID gets the value of ThreatID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyThreatID() (value int64, err error) {
	retValue, err := instance.GetProperty("ThreatID")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreatName sets the value of ThreatName for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyThreatName(value string) (err error) {
	return instance.SetProperty("ThreatName", value)
}

// GetThreatName gets the value of ThreatName for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyThreatName() (value string, err error) {
	retValue, err := instance.GetProperty("ThreatName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeID sets the value of TypeID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyTypeID(value uint8) (err error) {
	return instance.SetProperty("TypeID", value)
}

// GetTypeID gets the value of TypeID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyTypeID() (value uint8, err error) {
	retValue, err := instance.GetProperty("TypeID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
