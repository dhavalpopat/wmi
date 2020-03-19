// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CAURun_Report_Chunk struct
type MSFT_CAURun_Report_Chunk struct {
	*MSFT_CAURun_Report_ID

	//
	Data string

	//
	SequenceNumber uint32
}

func NewMSFT_CAURun_Report_ChunkEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAURun_Report_Chunk, err error) {
	tmp, err := NewMSFT_CAURun_Report_IDEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAURun_Report_Chunk{
		MSFT_CAURun_Report_ID: tmp,
	}
	return
}

func NewMSFT_CAURun_Report_ChunkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAURun_Report_Chunk, err error) {
	tmp, err := NewMSFT_CAURun_Report_IDEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAURun_Report_Chunk{
		MSFT_CAURun_Report_ID: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSFT_CAURun_Report_Chunk) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_CAURun_Report_Chunk) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequenceNumber sets the value of SequenceNumber for the instance
func (instance *MSFT_CAURun_Report_Chunk) SetPropertySequenceNumber(value uint32) (err error) {
	return instance.SetProperty("SequenceNumber", value)
}

// GetSequenceNumber gets the value of SequenceNumber for the instance
func (instance *MSFT_CAURun_Report_Chunk) GetPropertySequenceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SequenceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
