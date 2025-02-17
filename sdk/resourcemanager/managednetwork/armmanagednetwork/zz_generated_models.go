//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

// ConnectivityCollection - The collection of Connectivity related groups and policies within the Managed Network
type ConnectivityCollection struct {
	// READ-ONLY; The collection of connectivity related Managed Network Groups within the Managed Network
	Groups []*Group `json:"groups,omitempty" azure:"ro"`

	// READ-ONLY; The collection of Managed Network Peering Policies within the Managed Network
	Peerings []*PeeringPolicy `json:"peerings,omitempty" azure:"ro"`
}

// ErrorResponse - The error response that indicates why an operation has failed.
type ErrorResponse struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// Group - The Managed Network Group resource
type Group struct {
	// Responsibility role under which this Managed Network Group will be created
	Kind *Kind `json:"kind,omitempty"`

	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Gets or sets the properties of a network group
	Properties *GroupProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GroupListResult - Result of the request to list Managed Network Groups. It contains a list of groups and a URL link to
// get the next set of results.
type GroupListResult struct {
	// Gets the URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets a page of ManagedNetworkGroup
	Value []*Group `json:"value,omitempty"`
}

// GroupProperties - Properties of a Managed Network Group
type GroupProperties struct {
	// The collection of management groups covered by the Managed Network
	ManagementGroups []*ResourceID `json:"managementGroups,omitempty"`

	// The collection of subnets covered by the Managed Network
	Subnets []*ResourceID `json:"subnets,omitempty"`

	// The collection of subscriptions covered by the Managed Network
	Subscriptions []*ResourceID `json:"subscriptions,omitempty"`

	// The collection of virtual nets covered by the Managed Network
	VirtualNetworks []*ResourceID `json:"virtualNetworks,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// GroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupsClient.BeginCreateOrUpdate method.
type GroupsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GroupsClientBeginDeleteOptions contains the optional parameters for the GroupsClient.BeginDelete method.
type GroupsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GroupsClientGetOptions contains the optional parameters for the GroupsClient.Get method.
type GroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// GroupsClientListByManagedNetworkOptions contains the optional parameters for the GroupsClient.ListByManagedNetwork method.
type GroupsClientListByManagedNetworkOptions struct {
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// May be used to limit the number of results in a page for list queries.
	Top *int32
}

// HubAndSpokePeeringPolicyProperties - Properties of a Hub and Spoke Peering Policy
type HubAndSpokePeeringPolicyProperties struct {
	// REQUIRED; Gets or sets the connectivity type of a network structure policy
	Type *Type `json:"type,omitempty"`

	// Gets or sets the hub virtual network ID
	Hub *ResourceID `json:"hub,omitempty"`

	// Gets or sets the mesh group IDs
	Mesh []*ResourceID `json:"mesh,omitempty"`

	// Gets or sets the spokes group IDs
	Spokes []*ResourceID `json:"spokes,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ListResult - Result of the request to list Managed Network. It contains a list of Managed Networks and a URL link to get
// the next set of results.
type ListResult struct {
	// Gets the URL to get the next page of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets a page of ManagedNetworks
	Value []*ManagedNetwork `json:"value,omitempty"`
}

// ManagedNetwork - The Managed Network resource
type ManagedNetwork struct {
	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The MNC properties
	Properties *Properties `json:"properties,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ManagedNetworksClientBeginDeleteOptions contains the optional parameters for the ManagedNetworksClient.BeginDelete method.
type ManagedNetworksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedNetworksClientBeginUpdateOptions contains the optional parameters for the ManagedNetworksClient.BeginUpdate method.
type ManagedNetworksClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedNetworksClientCreateOrUpdateOptions contains the optional parameters for the ManagedNetworksClient.CreateOrUpdate
// method.
type ManagedNetworksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagedNetworksClientGetOptions contains the optional parameters for the ManagedNetworksClient.Get method.
type ManagedNetworksClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagedNetworksClientListByResourceGroupOptions contains the optional parameters for the ManagedNetworksClient.ListByResourceGroup
// method.
type ManagedNetworksClientListByResourceGroupOptions struct {
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// May be used to limit the number of results in a page for list queries.
	Top *int32
}

// ManagedNetworksClientListBySubscriptionOptions contains the optional parameters for the ManagedNetworksClient.ListBySubscription
// method.
type ManagedNetworksClientListBySubscriptionOptions struct {
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// May be used to limit the number of results in a page for list queries.
	Top *int32
}

// MeshPeeringPolicyProperties - Properties of a Mesh Peering Policy
type MeshPeeringPolicyProperties struct {
	// REQUIRED; Gets or sets the connectivity type of a network structure policy
	Type *Type `json:"type,omitempty"`

	// Gets or sets the hub virtual network ID
	Hub *ResourceID `json:"hub,omitempty"`

	// Gets or sets the mesh group IDs
	Mesh []*ResourceID `json:"mesh,omitempty"`

	// Gets or sets the spokes group IDs
	Spokes []*ResourceID `json:"spokes,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// Operation - REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.ManagedNetwork
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Managed Network operations. It contains a list of operations and a
// URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Resource Provider operations supported by the Managed Network resource provider.
	Value []*Operation `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PeeringPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the PeeringPoliciesClient.BeginCreateOrUpdate
// method.
type PeeringPoliciesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PeeringPoliciesClientBeginDeleteOptions contains the optional parameters for the PeeringPoliciesClient.BeginDelete method.
type PeeringPoliciesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PeeringPoliciesClientGetOptions contains the optional parameters for the PeeringPoliciesClient.Get method.
type PeeringPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PeeringPoliciesClientListByManagedNetworkOptions contains the optional parameters for the PeeringPoliciesClient.ListByManagedNetwork
// method.
type PeeringPoliciesClientListByManagedNetworkOptions struct {
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// May be used to limit the number of results in a page for list queries.
	Top *int32
}

// PeeringPolicy - The Managed Network Peering Policy resource
type PeeringPolicy struct {
	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Gets or sets the properties of a Managed Network Policy
	Properties *PeeringPolicyProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PeeringPolicyListResult - Result of the request to list Managed Network Peering Policies. It contains a list of policies
// and a URL link to get the next set of results.
type PeeringPolicyListResult struct {
	// Gets the URL to get the next page of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets a page of Peering Policies
	Value []*PeeringPolicy `json:"value,omitempty"`
}

// PeeringPolicyProperties - Properties of a Managed Network Peering Policy
type PeeringPolicyProperties struct {
	// REQUIRED; Gets or sets the connectivity type of a network structure policy
	Type *Type `json:"type,omitempty"`

	// Gets or sets the hub virtual network ID
	Hub *ResourceID `json:"hub,omitempty"`

	// Gets or sets the mesh group IDs
	Mesh []*ResourceID `json:"mesh,omitempty"`

	// Gets or sets the spokes group IDs
	Spokes []*ResourceID `json:"spokes,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// Properties of Managed Network
type Properties struct {
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only
	// property that is reflective of all ScopeAssignments for this Managed
	// Network
	Scope *Scope `json:"scope,omitempty"`

	// READ-ONLY; The collection of groups and policies concerned with connectivity
	Connectivity *ConnectivityCollection `json:"connectivity,omitempty" azure:"ro"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location
// and tags
type ProxyResource struct {
	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - The general resource model definition
type Resource struct {
	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceID - Generic pointer to a resource
type ResourceID struct {
	// Resource Id
	ID *string `json:"id,omitempty"`
}

// ResourceProperties - Base for resource properties.
type ResourceProperties struct {
	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// Scope of a Managed Network
type Scope struct {
	// The collection of management groups covered by the Managed Network
	ManagementGroups []*ResourceID `json:"managementGroups,omitempty"`

	// The collection of subnets covered by the Managed Network
	Subnets []*ResourceID `json:"subnets,omitempty"`

	// The collection of subscriptions covered by the Managed Network
	Subscriptions []*ResourceID `json:"subscriptions,omitempty"`

	// The collection of virtual nets covered by the Managed Network
	VirtualNetworks []*ResourceID `json:"virtualNetworks,omitempty"`
}

// ScopeAssignment - The Managed Network resource
type ScopeAssignment struct {
	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The Scope Assignment properties
	Properties *ScopeAssignmentProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ScopeAssignmentListResult - Result of the request to list ScopeAssignment. It contains a list of groups and a URL link
// to get the next set of results.
type ScopeAssignmentListResult struct {
	// Gets the URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets a page of ScopeAssignment
	Value []*ScopeAssignment `json:"value,omitempty"`
}

// ScopeAssignmentProperties - Properties of Managed Network
type ScopeAssignmentProperties struct {
	// The managed network ID with scope will be assigned to.
	AssignedManagedNetwork *string `json:"assignedManagedNetwork,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ScopeAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ScopeAssignmentsClient.CreateOrUpdate
// method.
type ScopeAssignmentsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ScopeAssignmentsClientDeleteOptions contains the optional parameters for the ScopeAssignmentsClient.Delete method.
type ScopeAssignmentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ScopeAssignmentsClientGetOptions contains the optional parameters for the ScopeAssignmentsClient.Get method.
type ScopeAssignmentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ScopeAssignmentsClientListOptions contains the optional parameters for the ScopeAssignmentsClient.List method.
type ScopeAssignmentsClientListOptions struct {
	// placeholder for future optional parameters
}

// TrackedResource - The resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Update Tags of Managed Network
type Update struct {
	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}
