// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSE8125520_1F5F_48D5_9C53_BB40C321ED3D
//////////////////////////////////////////////
package nse8125520_1f5f_48d5_9c53_bb40c321ed3d

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __IndicationRelated struct
type __IndicationRelated struct {
	*__SystemClass
}

func New__IndicationRelatedEx1(instance *cim.WmiInstance) (newInstance *__IndicationRelated, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__IndicationRelated{
		__SystemClass: tmp,
	}
	return
}

func New__IndicationRelatedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__IndicationRelated, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__IndicationRelated{
		__SystemClass: tmp,
	}
	return
}
