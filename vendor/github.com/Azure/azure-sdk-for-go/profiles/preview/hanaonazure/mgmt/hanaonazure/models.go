// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package hanaonazure

import original "github.com/Azure/azure-sdk-for-go/services/hanaonazure/mgmt/2017-11-03-preview/hanaonazure"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type HanaInstancesClient = original.HanaInstancesClient

func NewHanaInstancesClient(subscriptionID string) HanaInstancesClient {
	return original.NewHanaInstancesClient(subscriptionID)
}
func NewHanaInstancesClientWithBaseURI(baseURI string, subscriptionID string) HanaInstancesClient {
	return original.NewHanaInstancesClientWithBaseURI(baseURI, subscriptionID)
}

type HanaHardwareTypeNamesEnum = original.HanaHardwareTypeNamesEnum

const (
	CiscoUCS HanaHardwareTypeNamesEnum = original.CiscoUCS
)

func PossibleHanaHardwareTypeNamesEnumValues() []HanaHardwareTypeNamesEnum {
	return original.PossibleHanaHardwareTypeNamesEnumValues()
}

type HanaInstanceSizeNamesEnum = original.HanaInstanceSizeNamesEnum

const (
	S144  HanaInstanceSizeNamesEnum = original.S144
	S144m HanaInstanceSizeNamesEnum = original.S144m
	S192  HanaInstanceSizeNamesEnum = original.S192
	S192m HanaInstanceSizeNamesEnum = original.S192m
	S72   HanaInstanceSizeNamesEnum = original.S72
	S72m  HanaInstanceSizeNamesEnum = original.S72m
)

func PossibleHanaInstanceSizeNamesEnumValues() []HanaInstanceSizeNamesEnum {
	return original.PossibleHanaInstanceSizeNamesEnumValues()
}

type Disk = original.Disk
type Display = original.Display
type ErrorResponse = original.ErrorResponse
type HanaInstance = original.HanaInstance
type HanaInstanceProperties = original.HanaInstanceProperties
type HanaInstancesListResult = original.HanaInstancesListResult
type HanaInstancesListResultIterator = original.HanaInstancesListResultIterator
type HanaInstancesListResultPage = original.HanaInstancesListResultPage
type HardwareProfile = original.HardwareProfile
type IPAddress = original.IPAddress
type NetworkProfile = original.NetworkProfile
type Operation = original.Operation
type OperationList = original.OperationList
type OSProfile = original.OSProfile
type Resource = original.Resource
type StorageProfile = original.StorageProfile
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
