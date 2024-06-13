// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// OpenTrustedDomain_End struct
type OpenTrustedDomain_End struct {
	*OpenTrustedDomain
}

func NewOpenTrustedDomain_EndEx1(instance *cim.WmiInstance) (newInstance *OpenTrustedDomain_End, err error) {
	tmp, err := NewOpenTrustedDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OpenTrustedDomain_End{
		OpenTrustedDomain: tmp,
	}
	return
}

func NewOpenTrustedDomain_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OpenTrustedDomain_End, err error) {
	tmp, err := NewOpenTrustedDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OpenTrustedDomain_End{
		OpenTrustedDomain: tmp,
	}
	return
}
