// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SomFilter struct
type MSFT_SomFilter struct {
	cim.WmiInstance

	//
	Author string

	//
	ChangeDate string

	//
	CreationDate string

	//
	Description string

	//
	Domain string

	//
	ID string

	//
	Name string

	//
	Rules []MSFT_Rule

	//
	SourceOrganization string
}

// SetAuthor sets the value of Author for the instance
func (instance *MSFT_SomFilter) SetPropertyAuthor(value string) (err error) {
	return instance.SetProperty("Author", value)
}

// GetAuthor gets the value of Author for the instance
func (instance *MSFT_SomFilter) GetPropertyAuthor() (value string, err error) {
	retValue, err := instance.GetProperty("Author")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChangeDate sets the value of ChangeDate for the instance
func (instance *MSFT_SomFilter) SetPropertyChangeDate(value string) (err error) {
	return instance.SetProperty("ChangeDate", value)
}

// GetChangeDate gets the value of ChangeDate for the instance
func (instance *MSFT_SomFilter) GetPropertyChangeDate() (value string, err error) {
	retValue, err := instance.GetProperty("ChangeDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationDate sets the value of CreationDate for the instance
func (instance *MSFT_SomFilter) SetPropertyCreationDate(value string) (err error) {
	return instance.SetProperty("CreationDate", value)
}

// GetCreationDate gets the value of CreationDate for the instance
func (instance *MSFT_SomFilter) GetPropertyCreationDate() (value string, err error) {
	retValue, err := instance.GetProperty("CreationDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_SomFilter) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_SomFilter) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomain sets the value of Domain for the instance
func (instance *MSFT_SomFilter) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *MSFT_SomFilter) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *MSFT_SomFilter) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MSFT_SomFilter) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_SomFilter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_SomFilter) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRules sets the value of Rules for the instance
func (instance *MSFT_SomFilter) SetPropertyRules(value []MSFT_Rule) (err error) {
	return instance.SetProperty("Rules", value)
}

// GetRules gets the value of Rules for the instance
func (instance *MSFT_SomFilter) GetPropertyRules() (value []MSFT_Rule, err error) {
	retValue, err := instance.GetProperty("Rules")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_Rule)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceOrganization sets the value of SourceOrganization for the instance
func (instance *MSFT_SomFilter) SetPropertySourceOrganization(value string) (err error) {
	return instance.SetProperty("SourceOrganization", value)
}

// GetSourceOrganization gets the value of SourceOrganization for the instance
func (instance *MSFT_SomFilter) GetPropertySourceOrganization() (value string, err error) {
	retValue, err := instance.GetProperty("SourceOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SomFilter) Evaluate() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Evaluate")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="filters" type="MSFT_SomFilter []"></param>

// <param name="results" type="uint32 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SomFilter) BatchEvaluate( /* IN */ filters []MSFT_SomFilter,
	/* OUT */ results []uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("BatchEvaluate", filters)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
