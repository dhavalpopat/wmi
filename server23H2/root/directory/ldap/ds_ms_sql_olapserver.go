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

// ds_ms_sql_olapserver struct
type ds_ms_sql_olapserver struct {
	*ads_ms_sql_olapserver
}

func Newds_ms_sql_olapserverEx1(instance *cim.WmiInstance) (newInstance *ds_ms_sql_olapserver, err error) {
	tmp, err := Newads_ms_sql_olapserverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ms_sql_olapserver{
		ads_ms_sql_olapserver: tmp,
	}
	return
}

func Newds_ms_sql_olapserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ms_sql_olapserver, err error) {
	tmp, err := Newads_ms_sql_olapserverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ms_sql_olapserver{
		ads_ms_sql_olapserver: tmp,
	}
	return
}
