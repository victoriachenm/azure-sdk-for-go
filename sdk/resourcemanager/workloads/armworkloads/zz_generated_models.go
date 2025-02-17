//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloads

import "time"

// DB2ProviderInstanceProperties - Gets or sets the DB2 provider properties.
type DB2ProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`

	// Gets or sets the db2 database name.
	DbName *string `json:"dbName,omitempty"`

	// Gets or sets the db2 database password.
	DbPassword *string `json:"dbPassword,omitempty"`

	// Gets or sets the key vault URI to secret with the database password.
	DbPasswordURI *string `json:"dbPasswordUri,omitempty"`

	// Gets or sets the db2 database sql port.
	DbPort *string `json:"dbPort,omitempty"`

	// Gets or sets the db2 database user name.
	DbUsername *string `json:"dbUsername,omitempty"`

	// Gets or sets the target virtual machine name.
	Hostname *string `json:"hostname,omitempty"`

	// Gets or sets the SAP System Identifier
	SapSid *string `json:"sapSid,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type DB2ProviderInstanceProperties.
func (d *DB2ProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: d.ProviderType,
	}
}

// Error - Standard error object.
type Error struct {
	// READ-ONLY; Server-defined set of error codes.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Array of details about specific errors that led to this reported error.
	Details []*Error `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError *ErrorInnerError `json:"innerError,omitempty" azure:"ro"`

	// READ-ONLY; Human-readable representation of the error.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Target of the error.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorInnerError - Object containing more specific information than the current object about the error.
type ErrorInnerError struct {
	// Standard error object.
	InnerError *Error `json:"innerError,omitempty"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// HanaDbProviderInstanceProperties - Gets or sets the provider properties.
type HanaDbProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`

	// Gets or sets the hana database name.
	DbName *string `json:"dbName,omitempty"`

	// Gets or sets the database password.
	DbPassword *string `json:"dbPassword,omitempty"`

	// Gets or sets the key vault URI to secret with the database password.
	DbPasswordURI *string `json:"dbPasswordUri,omitempty"`

	// Gets or sets the blob URI to SSL certificate for the DB.
	DbSSLCertificateURI *string `json:"dbSslCertificateUri,omitempty"`

	// Gets or sets the database user name.
	DbUsername *string `json:"dbUsername,omitempty"`

	// Gets or sets the target virtual machine size.
	Hostname *string `json:"hostname,omitempty"`

	// Gets or sets the database instance number.
	InstanceNumber *string `json:"instanceNumber,omitempty"`

	// Gets or sets the database sql port.
	SQLPort *string `json:"sqlPort,omitempty"`

	// Gets or sets the hostname(s) in the SSL certificate.
	SSLHostNameInCertificate *string `json:"sslHostNameInCertificate,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type HanaDbProviderInstanceProperties.
func (h *HanaDbProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: h.ProviderType,
	}
}

// ManagedRGConfiguration - Managed resource group configuration
type ManagedRGConfiguration struct {
	// Managed resource group name
	Name *string `json:"name,omitempty"`
}

// Monitor - SAP monitor info on Azure (ARM properties and SAP monitor properties)
type Monitor struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// SAP monitor properties
	Properties *MonitorProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MonitorListResult - The response from the List SAP monitors operation.
type MonitorListResult struct {
	// The URL to get the next set of SAP monitors.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of SAP monitors.
	Value []*Monitor `json:"value,omitempty"`
}

// MonitorProperties - Describes the properties of a SAP monitor.
type MonitorProperties struct {
	// The monitor resources will be deployed in the monitoring region. The subnet region should be same as the monitoring region.
	AppLocation *string `json:"appLocation,omitempty"`

	// The ARM ID of the Log Analytics Workspace that is used for monitoring.
	LogAnalyticsWorkspaceArmID *string `json:"logAnalyticsWorkspaceArmId,omitempty"`

	// Managed resource group configuration
	ManagedResourceGroupConfiguration *ManagedRGConfiguration `json:"managedResourceGroupConfiguration,omitempty"`

	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string `json:"monitorSubnet,omitempty"`

	// Sets the routing preference of the monitor. By default only RFC1918 traffic is routed to the customer VNET.
	RoutingPreference *RoutingPreference `json:"routingPreference,omitempty"`

	// READ-ONLY; Defines the SAP monitor errors.
	Errors *MonitorPropertiesErrors `json:"errors,omitempty" azure:"ro"`

	// READ-ONLY; The ARM ID of the MSI used for monitoring.
	MsiArmID *string `json:"msiArmId,omitempty" azure:"ro"`

	// READ-ONLY; State of provisioning of the monitor.
	ProvisioningState *WorkloadMonitorProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MonitorPropertiesErrors - Defines the SAP monitor errors.
type MonitorPropertiesErrors struct {
	// READ-ONLY; Server-defined set of error codes.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Array of details about specific errors that led to this reported error.
	Details []*Error `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError *ErrorInnerError `json:"innerError,omitempty" azure:"ro"`

	// READ-ONLY; Human-readable representation of the error.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Target of the error.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MonitorsClientBeginCreateOptions contains the optional parameters for the MonitorsClient.BeginCreate method.
type MonitorsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
type MonitorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
type MonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.ListByResourceGroup method.
type MonitorsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListOptions contains the optional parameters for the MonitorsClient.List method.
type MonitorsClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientUpdateOptions contains the optional parameters for the MonitorsClient.Update method.
type MonitorsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// MsSQLServerProviderInstanceProperties - Gets or sets the SQL server provider properties.
type MsSQLServerProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`

	// Gets or sets the database password.
	DbPassword *string `json:"dbPassword,omitempty"`

	// Gets or sets the key vault URI to secret with the database password.
	DbPasswordURI *string `json:"dbPasswordUri,omitempty"`

	// Gets or sets the database sql port.
	DbPort *string `json:"dbPort,omitempty"`

	// Gets or sets the database user name.
	DbUsername *string `json:"dbUsername,omitempty"`

	// Gets or sets the SQL server host name.
	Hostname *string `json:"hostname,omitempty"`

	// Gets or sets the SAP System Identifier
	SapSid *string `json:"sapSid,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type MsSQLServerProviderInstanceProperties.
func (m *MsSQLServerProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: m.ProviderType,
	}
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationStatusResult - The current status of an async operation.
type OperationStatusResult struct {
	// REQUIRED; Operation status.
	Status *string `json:"status,omitempty"`

	// The end time of the operation.
	EndTime *time.Time `json:"endTime,omitempty"`

	// If present, details of the operation error.
	Error *ErrorDetail `json:"error,omitempty"`

	// Fully qualified ID for the async operation.
	ID *string `json:"id,omitempty"`

	// Name of the async operation.
	Name *string `json:"name,omitempty"`

	// The operations list.
	Operations []*OperationStatusResult `json:"operations,omitempty"`

	// Percent of the operation that is complete.
	PercentComplete *float32 `json:"percentComplete,omitempty"`

	// The start time of the operation.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsContent - Defines the workload operation content.
type OperationsContent struct {
	// Operations content.
	Properties *OperationsDefinition `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// OperationsDefinition - Properties of an Operation.
type OperationsDefinition struct {
	// REQUIRED; Display information of the operation.
	Display *OperationsDefinitionDisplay `json:"display,omitempty"`

	// REQUIRED; Name of the operation.
	Name *string `json:"name,omitempty"`

	// Defines the action type of workload operation.
	ActionType *WorkloadMonitorActionType `json:"actionType,omitempty"`

	// Indicates whether the operation applies to data-plane.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Defines the workload operation origin.
	Origin *OperationProperties `json:"origin,omitempty"`

	// Defines the workload operation properties.
	Properties interface{} `json:"properties,omitempty"`
}

// OperationsDefinitionArrayResponseWithContinuation - Defines the workload operation definition response.
type OperationsDefinitionArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// Defines the workload operation definition response properties.
	Value []*OperationsDefinition `json:"value,omitempty"`
}

// OperationsDefinitionDisplay - Display information of the operation.
type OperationsDefinitionDisplay struct {
	// REQUIRED; Describes the workload operation.
	Description *string `json:"description,omitempty"`

	// REQUIRED; Defines the workload operation.
	Operation *string `json:"operation,omitempty"`

	// REQUIRED; Defines the workload provider.
	Provider *string `json:"provider,omitempty"`

	// REQUIRED; Defines the workload resource.
	Resource *string `json:"resource,omitempty"`
}

// OperationsDisplayDefinition - Defines the workload operation.
type OperationsDisplayDefinition struct {
	// REQUIRED; Describes the workload operation.
	Description *string `json:"description,omitempty"`

	// REQUIRED; Defines the workload operation.
	Operation *string `json:"operation,omitempty"`

	// REQUIRED; Defines the workload provider.
	Provider *string `json:"provider,omitempty"`

	// REQUIRED; Defines the workload resource.
	Resource *string `json:"resource,omitempty"`
}

// PrometheusHaClusterProviderInstanceProperties - Gets or sets the PrometheusHaCluster provider properties.
type PrometheusHaClusterProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`

	// Gets or sets the clusterName.
	ClusterName *string `json:"clusterName,omitempty"`

	// Gets or sets the target machine name.
	Hostname *string `json:"hostname,omitempty"`

	// URL of the Node Exporter endpoint.
	PrometheusURL *string `json:"prometheusUrl,omitempty"`

	// Gets or sets the cluster sid.
	Sid *string `json:"sid,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type PrometheusHaClusterProviderInstanceProperties.
func (p *PrometheusHaClusterProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: p.ProviderType,
	}
}

// PrometheusOSProviderInstanceProperties - Gets or sets the PrometheusOS provider properties.
type PrometheusOSProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`

	// URL of the Node Exporter endpoint
	PrometheusURL *string `json:"prometheusUrl,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type PrometheusOSProviderInstanceProperties.
func (p *PrometheusOSProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: p.ProviderType,
	}
}

// ProviderInstance - A provider instance associated with monitor.
type ProviderInstance struct {
	// Provider Instance properties
	Properties *ProviderInstanceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ProviderInstanceListResult - The response from the List provider instances operation.
type ProviderInstanceListResult struct {
	// The URL to get the next set of provider instances.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of provider instances.
	Value []*ProviderInstance `json:"value,omitempty"`
}

// ProviderInstanceProperties - Describes the properties of a provider instance.
type ProviderInstanceProperties struct {
	// Defines the provider instance errors.
	ProviderSettings ProviderSpecificPropertiesClassification `json:"providerSettings,omitempty"`

	// READ-ONLY; Defines the provider instance errors.
	Errors *ProviderInstancePropertiesErrors `json:"errors,omitempty" azure:"ro"`

	// READ-ONLY; State of provisioning of the provider instance
	ProvisioningState *WorkloadMonitorProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ProviderInstancePropertiesErrors - Defines the provider instance errors.
type ProviderInstancePropertiesErrors struct {
	// READ-ONLY; Server-defined set of error codes.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Array of details about specific errors that led to this reported error.
	Details []*Error `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError *ErrorInnerError `json:"innerError,omitempty" azure:"ro"`

	// READ-ONLY; Human-readable representation of the error.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Target of the error.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ProviderInstancesClientBeginCreateOptions contains the optional parameters for the ProviderInstancesClient.BeginCreate
// method.
type ProviderInstancesClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProviderInstancesClientBeginDeleteOptions contains the optional parameters for the ProviderInstancesClient.BeginDelete
// method.
type ProviderInstancesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProviderInstancesClientGetOptions contains the optional parameters for the ProviderInstancesClient.Get method.
type ProviderInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProviderInstancesClientListOptions contains the optional parameters for the ProviderInstancesClient.List method.
type ProviderInstancesClientListOptions struct {
	// placeholder for future optional parameters
}

// ProviderSpecificPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetProviderSpecificProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DB2ProviderInstanceProperties, *HanaDbProviderInstanceProperties, *MsSQLServerProviderInstanceProperties, *PrometheusHaClusterProviderInstanceProperties,
// - *PrometheusOSProviderInstanceProperties, *ProviderSpecificProperties, *SapNetWeaverProviderInstanceProperties
type ProviderSpecificPropertiesClassification interface {
	// GetProviderSpecificProperties returns the ProviderSpecificProperties content of the underlying type.
	GetProviderSpecificProperties() *ProviderSpecificProperties
}

// ProviderSpecificProperties - Gets or sets the provider specific properties.
type ProviderSpecificProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type ProviderSpecificProperties.
func (p *ProviderSpecificProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return p
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SapNetWeaverProviderInstanceProperties - Gets or sets the provider properties.
type SapNetWeaverProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string `json:"providerType,omitempty"`

	// Gets or sets the SAP Client ID.
	SapClientID *string `json:"sapClientId,omitempty"`

	// Gets or sets the list of HostFile Entries
	SapHostFileEntries []*string `json:"sapHostFileEntries,omitempty"`

	// Gets or sets the target virtual machine IP Address/FQDN.
	SapHostname *string `json:"sapHostname,omitempty"`

	// Gets or sets the instance number of SAP NetWeaver.
	SapInstanceNr *string `json:"sapInstanceNr,omitempty"`

	// Sets the SAP password.
	SapPassword *string `json:"sapPassword,omitempty"`

	// Gets or sets the key vault URI to secret with the SAP password.
	SapPasswordURI *string `json:"sapPasswordUri,omitempty"`

	// Gets or sets the SAP HTTP port number.
	SapPortNumber *string `json:"sapPortNumber,omitempty"`

	// Gets or sets the blob URI to SSL certificate for the SAP system.
	SapSSLCertificateURI *string `json:"sapSslCertificateUri,omitempty"`

	// Gets or sets the SAP System Identifier
	SapSid *string `json:"sapSid,omitempty"`

	// Gets or sets the SAP user name.
	SapUsername *string `json:"sapUsername,omitempty"`
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type SapNetWeaverProviderInstanceProperties.
func (s *SapNetWeaverProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: s.ProviderType,
	}
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// Tags field of the resource.
type Tags struct {
	// Tags field of the resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}
