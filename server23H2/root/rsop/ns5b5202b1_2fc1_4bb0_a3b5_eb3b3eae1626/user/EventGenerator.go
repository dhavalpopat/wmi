// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS5B5202B1_2FC1_4BB0_A3B5_EB3B3EAE1626.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __EventGenerator struct
type __EventGenerator struct {
	*__IndicationRelated
}

func New__EventGeneratorEx1(instance *cim.WmiInstance) (newInstance *__EventGenerator, err error) {
	tmp, err := New__IndicationRelatedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventGenerator{
		__IndicationRelated: tmp,
	}
	return
}

func New__EventGeneratorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventGenerator, err error) {
	tmp, err := New__IndicationRelatedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventGenerator{
		__IndicationRelated: tmp,
	}
	return
}
