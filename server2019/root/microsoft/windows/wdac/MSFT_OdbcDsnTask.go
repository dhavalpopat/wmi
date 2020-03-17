// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_OdbcDsnTask struct
type MSFT_OdbcDsnTask struct {
	cim.WmiInstance
}

//

// <param name="DriverName" type="string "></param>
// <param name="DsnType" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_OdbcDsn []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDsnTask) Get( /* IN */ Name string,
	/* IN */ DriverName string,
	/* IN */ Platform string,
	/* IN */ DsnType string,
	/* OUT */ cmdletOutput []MSFT_OdbcDsn) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, DriverName, Platform, DsnType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DriverName" type="string "></param>
// <param name="DsnType" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>
// <param name="SetPropertyValue" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_OdbcDsn []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDsnTask) Add( /* IN */ Name string,
	/* IN */ DriverName string,
	/* IN */ SetPropertyValue []string,
	/* IN */ PassThru bool,
	/* IN */ Platform string,
	/* IN */ DsnType string,
	/* OUT */ cmdletOutput []MSFT_OdbcDsn) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", Name, DriverName, SetPropertyValue, PassThru, Platform, DsnType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_OdbcDsn []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="RemovePropertyValue" type="string []"></param>
// <param name="SetPropertyValue" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_OdbcDsn []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDsnTask) SetByInputObject( /* IN */ PassThru bool,
	/* IN */ SetPropertyValue []string,
	/* IN */ RemovePropertyValue []string,
	/* IN */ InputObject []MSFT_OdbcDsn,
	/* OUT */ cmdletOutput []MSFT_OdbcDsn) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByInputObject", PassThru, SetPropertyValue, RemovePropertyValue, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DriverName" type="string "></param>
// <param name="DsnType" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>
// <param name="RemovePropertyValue" type="string []"></param>
// <param name="SetPropertyValue" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_OdbcDsn []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDsnTask) SetByName( /* IN */ PassThru bool,
	/* IN */ SetPropertyValue []string,
	/* IN */ RemovePropertyValue []string,
	/* IN */ Name string,
	/* IN */ DriverName string,
	/* IN */ Platform string,
	/* IN */ DsnType string,
	/* OUT */ cmdletOutput []MSFT_OdbcDsn) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByName", PassThru, SetPropertyValue, RemovePropertyValue, Name, DriverName, Platform, DsnType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_OdbcDsn []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_OdbcDsn []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDsnTask) RemoveByDsnObject( /* IN */ PassThru bool,
	/* IN */ InputObject []MSFT_OdbcDsn,
	/* OUT */ cmdletOutput []MSFT_OdbcDsn) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveByDsnObject", PassThru, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DriverName" type="string "></param>
// <param name="DsnType" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_OdbcDsn []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDsnTask) RemoveByName( /* IN */ PassThru bool,
	/* IN */ Name string,
	/* IN */ DriverName string,
	/* IN */ Platform string,
	/* IN */ DsnType string,
	/* OUT */ cmdletOutput []MSFT_OdbcDsn) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveByName", PassThru, Name, DriverName, Platform, DsnType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
