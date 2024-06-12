// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_NodeToHostedService struct
type MSCluster_NodeToHostedService struct {
	*CIM_HostedService
}

func NewMSCluster_NodeToHostedServiceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_NodeToHostedService, err error) {
	tmp, err := NewCIM_HostedServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeToHostedService{
		CIM_HostedService: tmp,
	}
	return
}

func NewMSCluster_NodeToHostedServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_NodeToHostedService, err error) {
	tmp, err := NewCIM_HostedServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeToHostedService{
		CIM_HostedService: tmp,
	}
	return
}
