// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_BIOSFeatureBIOSElements struct
type CIM_BIOSFeatureBIOSElements struct {
	*CIM_SoftwareFeatureSoftwareElements
}

func NewCIM_BIOSFeatureBIOSElementsEx1(instance *cim.WmiInstance) (newInstance *CIM_BIOSFeatureBIOSElements, err error) {
	tmp, err := NewCIM_SoftwareFeatureSoftwareElementsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSFeatureBIOSElements{
		CIM_SoftwareFeatureSoftwareElements: tmp,
	}
	return
}

func NewCIM_BIOSFeatureBIOSElementsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BIOSFeatureBIOSElements, err error) {
	tmp, err := NewCIM_SoftwareFeatureSoftwareElementsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSFeatureBIOSElements{
		CIM_SoftwareFeatureSoftwareElements: tmp,
	}
	return
}
