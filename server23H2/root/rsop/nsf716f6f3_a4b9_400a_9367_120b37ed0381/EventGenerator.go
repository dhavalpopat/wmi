// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSF716F6F3_A4B9_400A_9367_120B37ED0381
//////////////////////////////////////////////
package nsf716f6f3_a4b9_400a_9367_120b37ed0381

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
