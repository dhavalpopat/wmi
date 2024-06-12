// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ads_intersitetransportcontainer struct
type ads_intersitetransportcontainer struct {
	*ds_top
}

func Newads_intersitetransportcontainerEx1(instance *cim.WmiInstance) (newInstance *ads_intersitetransportcontainer, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_intersitetransportcontainer{
		ds_top: tmp,
	}
	return
}

func Newads_intersitetransportcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_intersitetransportcontainer, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_intersitetransportcontainer{
		ds_top: tmp,
	}
	return
}
