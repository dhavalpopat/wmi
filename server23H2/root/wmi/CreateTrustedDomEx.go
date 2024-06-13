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

// CreateTrustedDomEx struct
type CreateTrustedDomEx struct {
	*MSLSATrace
}

func NewCreateTrustedDomExEx1(instance *cim.WmiInstance) (newInstance *CreateTrustedDomEx, err error) {
	tmp, err := NewMSLSATraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CreateTrustedDomEx{
		MSLSATrace: tmp,
	}
	return
}

func NewCreateTrustedDomExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CreateTrustedDomEx, err error) {
	tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CreateTrustedDomEx{
		MSLSATrace: tmp,
	}
	return
}
