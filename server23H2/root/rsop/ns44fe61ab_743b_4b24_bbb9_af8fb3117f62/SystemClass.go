// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS44FE61AB_743B_4B24_BBB9_AF8FB3117F62
//////////////////////////////////////////////
package ns44fe61ab_743b_4b24_bbb9_af8fb3117f62

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __SystemClass struct
type __SystemClass struct {
	*cim.WmiInstance
}

func New__SystemClassEx1(instance *cim.WmiInstance) (newInstance *__SystemClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &__SystemClass{
		WmiInstance: tmp,
	}
	return
}

func New__SystemClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__SystemClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__SystemClass{
		WmiInstance: tmp,
	}
	return
}
