// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SoftError_EmptyCache struct
type MSFT_SoftError_EmptyCache struct {
	*MSFT_SoftError
}

func NewMSFT_SoftError_EmptyCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_SoftError_EmptyCache, err error) {
	tmp, err := NewMSFT_SoftErrorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_SoftError_EmptyCache{
		MSFT_SoftError: tmp,
	}
	return
}

func NewMSFT_SoftError_EmptyCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SoftError_EmptyCache, err error) {
	tmp, err := NewMSFT_SoftErrorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SoftError_EmptyCache{
		MSFT_SoftError: tmp,
	}
	return
}
