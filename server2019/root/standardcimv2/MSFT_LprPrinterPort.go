// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_LprPrinterPort struct
type MSFT_LprPrinterPort struct {
	*MSFT_PrinterPort

	//
	HostName string

	//
	PrinterName string
}

func NewMSFT_LprPrinterPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_LprPrinterPort, err error) {
	tmp, err := NewMSFT_PrinterPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_LprPrinterPort{
		MSFT_PrinterPort: tmp,
	}
	return
}

func NewMSFT_LprPrinterPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_LprPrinterPort, err error) {
	tmp, err := NewMSFT_PrinterPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_LprPrinterPort{
		MSFT_PrinterPort: tmp,
	}
	return
}

// SetHostName sets the value of HostName for the instance
func (instance *MSFT_LprPrinterPort) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", value)
}

// GetHostName gets the value of HostName for the instance
func (instance *MSFT_LprPrinterPort) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrinterName sets the value of PrinterName for the instance
func (instance *MSFT_LprPrinterPort) SetPropertyPrinterName(value string) (err error) {
	return instance.SetProperty("PrinterName", value)
}

// GetPrinterName gets the value of PrinterName for the instance
func (instance *MSFT_LprPrinterPort) GetPropertyPrinterName() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
