// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msds_shadowprincipalcontainer struct
type ds_msds_shadowprincipalcontainer struct {
	*ads_msds_shadowprincipalcontainer
}

func Newds_msds_shadowprincipalcontainerEx1(instance *cim.WmiInstance) (newInstance *ds_msds_shadowprincipalcontainer, err error) {
	tmp, err := Newads_msds_shadowprincipalcontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_shadowprincipalcontainer{
		ads_msds_shadowprincipalcontainer: tmp,
	}
	return
}

func Newds_msds_shadowprincipalcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_shadowprincipalcontainer, err error) {
	tmp, err := Newads_msds_shadowprincipalcontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_shadowprincipalcontainer{
		ads_msds_shadowprincipalcontainer: tmp,
	}
	return
}
