// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_RDS_Connection_Broker struct
type ServerComponent_RDS_Connection_Broker struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_RDS_Connection_BrokerEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_RDS_Connection_Broker, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_RDS_Connection_Broker{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_RDS_Connection_BrokerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_RDS_Connection_Broker, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_RDS_Connection_Broker{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}