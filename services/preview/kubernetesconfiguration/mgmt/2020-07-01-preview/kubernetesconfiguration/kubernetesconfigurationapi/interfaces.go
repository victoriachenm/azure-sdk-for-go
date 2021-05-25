package kubernetesconfigurationapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/kubernetesconfiguration/mgmt/2020-07-01-preview/kubernetesconfiguration"
	"github.com/Azure/go-autorest/autorest"
)

// SourceControlConfigurationsClientAPI contains the set of methods on the SourceControlConfigurationsClient type.
type SourceControlConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, sourceControlConfiguration kubernetesconfiguration.SourceControlConfiguration) (result kubernetesconfiguration.SourceControlConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) (result kubernetesconfiguration.SourceControlConfigurationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) (result kubernetesconfiguration.SourceControlConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result kubernetesconfiguration.SourceControlConfigurationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result kubernetesconfiguration.SourceControlConfigurationListIterator, err error)
}

var _ SourceControlConfigurationsClientAPI = (*kubernetesconfiguration.SourceControlConfigurationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result kubernetesconfiguration.ResourceProviderOperationListPage, err error)
	ListComplete(ctx context.Context) (result kubernetesconfiguration.ResourceProviderOperationListIterator, err error)
}

var _ OperationsClientAPI = (*kubernetesconfiguration.OperationsClient)(nil)

// ExtensionsClientAPI contains the set of methods on the ExtensionsClient type.
type ExtensionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionInstanceName string, extensionInstance kubernetesconfiguration.ExtensionInstance) (result kubernetesconfiguration.ExtensionInstance, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionInstanceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionInstanceName string) (result kubernetesconfiguration.ExtensionInstance, err error)
	List(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result kubernetesconfiguration.ExtensionInstancesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result kubernetesconfiguration.ExtensionInstancesListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionInstanceName string, extensionInstance kubernetesconfiguration.ExtensionInstanceUpdate) (result kubernetesconfiguration.ExtensionInstance, err error)
}

var _ ExtensionsClientAPI = (*kubernetesconfiguration.ExtensionsClient)(nil)