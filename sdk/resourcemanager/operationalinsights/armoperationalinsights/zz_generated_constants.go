//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

const (
	moduleName    = "armoperationalinsights"
	moduleVersion = "v0.5.0"
)

// BillingType - Configures whether billing will be only on the cluster or each workspace will be billed by its proportional
// use. This does not change the overall billing, only how it will be distributed. Default
// value is 'Cluster'
type BillingType string

const (
	BillingTypeCluster    BillingType = "Cluster"
	BillingTypeWorkspaces BillingType = "Workspaces"
)

// PossibleBillingTypeValues returns the possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{
		BillingTypeCluster,
		BillingTypeWorkspaces,
	}
}

// Capacity - The capacity value
type Capacity int64

const (
	CapacityFiveHundred  Capacity = 500
	CapacityTenHundred   Capacity = 1000
	CapacityTwoThousand  Capacity = 2000
	CapacityFiveThousand Capacity = 5000
)

// PossibleCapacityValues returns the possible values for the Capacity const type.
func PossibleCapacityValues() []Capacity {
	return []Capacity{
		CapacityFiveHundred,
		CapacityTenHundred,
		CapacityTwoThousand,
		CapacityFiveThousand,
	}
}

// CapacityReservationLevel - The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
type CapacityReservationLevel int32

const (
	CapacityReservationLevelOneHundred   CapacityReservationLevel = 100
	CapacityReservationLevelTwoHundred   CapacityReservationLevel = 200
	CapacityReservationLevelThreeHundred CapacityReservationLevel = 300
	CapacityReservationLevelFourHundred  CapacityReservationLevel = 400
	CapacityReservationLevelFiveHundred  CapacityReservationLevel = 500
	CapacityReservationLevelTenHundred   CapacityReservationLevel = 1000
	CapacityReservationLevelTwoThousand  CapacityReservationLevel = 2000
	CapacityReservationLevelFiveThousand CapacityReservationLevel = 5000
)

// PossibleCapacityReservationLevelValues returns the possible values for the CapacityReservationLevel const type.
func PossibleCapacityReservationLevelValues() []CapacityReservationLevel {
	return []CapacityReservationLevel{
		CapacityReservationLevelOneHundred,
		CapacityReservationLevelTwoHundred,
		CapacityReservationLevelThreeHundred,
		CapacityReservationLevelFourHundred,
		CapacityReservationLevelFiveHundred,
		CapacityReservationLevelTenHundred,
		CapacityReservationLevelTwoThousand,
		CapacityReservationLevelFiveThousand,
	}
}

// ClusterEntityStatus - The provisioning state of the cluster.
type ClusterEntityStatus string

const (
	ClusterEntityStatusCanceled            ClusterEntityStatus = "Canceled"
	ClusterEntityStatusCreating            ClusterEntityStatus = "Creating"
	ClusterEntityStatusDeleting            ClusterEntityStatus = "Deleting"
	ClusterEntityStatusFailed              ClusterEntityStatus = "Failed"
	ClusterEntityStatusProvisioningAccount ClusterEntityStatus = "ProvisioningAccount"
	ClusterEntityStatusSucceeded           ClusterEntityStatus = "Succeeded"
	ClusterEntityStatusUpdating            ClusterEntityStatus = "Updating"
)

// PossibleClusterEntityStatusValues returns the possible values for the ClusterEntityStatus const type.
func PossibleClusterEntityStatusValues() []ClusterEntityStatus {
	return []ClusterEntityStatus{
		ClusterEntityStatusCanceled,
		ClusterEntityStatusCreating,
		ClusterEntityStatusDeleting,
		ClusterEntityStatusFailed,
		ClusterEntityStatusProvisioningAccount,
		ClusterEntityStatusSucceeded,
		ClusterEntityStatusUpdating,
	}
}

// ClusterSKUNameEnum - The name of the SKU.
type ClusterSKUNameEnum string

const (
	ClusterSKUNameEnumCapacityReservation ClusterSKUNameEnum = "CapacityReservation"
)

// PossibleClusterSKUNameEnumValues returns the possible values for the ClusterSKUNameEnum const type.
func PossibleClusterSKUNameEnumValues() []ClusterSKUNameEnum {
	return []ClusterSKUNameEnum{
		ClusterSKUNameEnumCapacityReservation,
	}
}

// ColumnDataTypeHintEnum - Column data type logical hint.
type ColumnDataTypeHintEnum string

const (
	// ColumnDataTypeHintEnumArmPath - An Azure Resource Model (ARM) path: /subscriptions/{...}/resourceGroups/{...}/providers/Microsoft.{...}/{...}/{...}/{...}...
	ColumnDataTypeHintEnumArmPath ColumnDataTypeHintEnum = "armPath"
	// ColumnDataTypeHintEnumGUID - A standard 128-bit GUID following the standard shape, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ColumnDataTypeHintEnumGUID ColumnDataTypeHintEnum = "guid"
	// ColumnDataTypeHintEnumIP - A standard V4/V6 ip address following the standard shape, x.x.x.x/y:y:y:y:y:y:y:y
	ColumnDataTypeHintEnumIP ColumnDataTypeHintEnum = "ip"
	// ColumnDataTypeHintEnumURI - A string that matches the pattern of a URI, for example, scheme://username:password@host:1234/this/is/a/path?k1=v1&k2=v2#fragment
	ColumnDataTypeHintEnumURI ColumnDataTypeHintEnum = "uri"
)

// PossibleColumnDataTypeHintEnumValues returns the possible values for the ColumnDataTypeHintEnum const type.
func PossibleColumnDataTypeHintEnumValues() []ColumnDataTypeHintEnum {
	return []ColumnDataTypeHintEnum{
		ColumnDataTypeHintEnumArmPath,
		ColumnDataTypeHintEnumGUID,
		ColumnDataTypeHintEnumIP,
		ColumnDataTypeHintEnumURI,
	}
}

// ColumnTypeEnum - Column data type.
type ColumnTypeEnum string

const (
	ColumnTypeEnumBoolean  ColumnTypeEnum = "boolean"
	ColumnTypeEnumDateTime ColumnTypeEnum = "dateTime"
	ColumnTypeEnumDynamic  ColumnTypeEnum = "dynamic"
	ColumnTypeEnumGUID     ColumnTypeEnum = "guid"
	ColumnTypeEnumInt      ColumnTypeEnum = "int"
	ColumnTypeEnumLong     ColumnTypeEnum = "long"
	ColumnTypeEnumReal     ColumnTypeEnum = "real"
	ColumnTypeEnumString   ColumnTypeEnum = "string"
)

// PossibleColumnTypeEnumValues returns the possible values for the ColumnTypeEnum const type.
func PossibleColumnTypeEnumValues() []ColumnTypeEnum {
	return []ColumnTypeEnum{
		ColumnTypeEnumBoolean,
		ColumnTypeEnumDateTime,
		ColumnTypeEnumDynamic,
		ColumnTypeEnumGUID,
		ColumnTypeEnumInt,
		ColumnTypeEnumLong,
		ColumnTypeEnumReal,
		ColumnTypeEnumString,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DataIngestionStatus - The status of data ingestion for this workspace.
type DataIngestionStatus string

const (
	// DataIngestionStatusApproachingQuota - 80% of daily cap quota reached.
	DataIngestionStatusApproachingQuota DataIngestionStatus = "ApproachingQuota"
	// DataIngestionStatusForceOff - Ingestion stopped following service setting change.
	DataIngestionStatusForceOff DataIngestionStatus = "ForceOff"
	// DataIngestionStatusForceOn - Ingestion started following service setting change.
	DataIngestionStatusForceOn DataIngestionStatus = "ForceOn"
	// DataIngestionStatusOverQuota - Reached daily cap quota, ingestion stopped.
	DataIngestionStatusOverQuota DataIngestionStatus = "OverQuota"
	// DataIngestionStatusRespectQuota - Ingestion enabled following daily cap quota reset, or subscription enablement.
	DataIngestionStatusRespectQuota DataIngestionStatus = "RespectQuota"
	// DataIngestionStatusSubscriptionSuspended - Ingestion stopped following suspended subscription.
	DataIngestionStatusSubscriptionSuspended DataIngestionStatus = "SubscriptionSuspended"
)

// PossibleDataIngestionStatusValues returns the possible values for the DataIngestionStatus const type.
func PossibleDataIngestionStatusValues() []DataIngestionStatus {
	return []DataIngestionStatus{
		DataIngestionStatusApproachingQuota,
		DataIngestionStatusForceOff,
		DataIngestionStatusForceOn,
		DataIngestionStatusOverQuota,
		DataIngestionStatusRespectQuota,
		DataIngestionStatusSubscriptionSuspended,
	}
}

// DataSourceKind - The kind of the DataSource.
type DataSourceKind string

const (
	DataSourceKindApplicationInsights                                  DataSourceKind = "ApplicationInsights"
	DataSourceKindAzureActivityLog                                     DataSourceKind = "AzureActivityLog"
	DataSourceKindAzureAuditLog                                        DataSourceKind = "AzureAuditLog"
	DataSourceKindChangeTrackingContentLocation                        DataSourceKind = "ChangeTrackingContentLocation"
	DataSourceKindChangeTrackingCustomPath                             DataSourceKind = "ChangeTrackingCustomPath"
	DataSourceKindChangeTrackingDataTypeConfiguration                  DataSourceKind = "ChangeTrackingDataTypeConfiguration"
	DataSourceKindChangeTrackingDefaultRegistry                        DataSourceKind = "ChangeTrackingDefaultRegistry"
	DataSourceKindChangeTrackingLinuxPath                              DataSourceKind = "ChangeTrackingLinuxPath"
	DataSourceKindChangeTrackingPath                                   DataSourceKind = "ChangeTrackingPath"
	DataSourceKindChangeTrackingRegistry                               DataSourceKind = "ChangeTrackingRegistry"
	DataSourceKindChangeTrackingServices                               DataSourceKind = "ChangeTrackingServices"
	DataSourceKindCustomLog                                            DataSourceKind = "CustomLog"
	DataSourceKindCustomLogCollection                                  DataSourceKind = "CustomLogCollection"
	DataSourceKindDNSAnalytics                                         DataSourceKind = "DnsAnalytics"
	DataSourceKindGenericDataSource                                    DataSourceKind = "GenericDataSource"
	DataSourceKindIISLogs                                              DataSourceKind = "IISLogs"
	DataSourceKindImportComputerGroup                                  DataSourceKind = "ImportComputerGroup"
	DataSourceKindItsm                                                 DataSourceKind = "Itsm"
	DataSourceKindLinuxChangeTrackingPath                              DataSourceKind = "LinuxChangeTrackingPath"
	DataSourceKindLinuxPerformanceCollection                           DataSourceKind = "LinuxPerformanceCollection"
	DataSourceKindLinuxPerformanceObject                               DataSourceKind = "LinuxPerformanceObject"
	DataSourceKindLinuxSyslog                                          DataSourceKind = "LinuxSyslog"
	DataSourceKindLinuxSyslogCollection                                DataSourceKind = "LinuxSyslogCollection"
	DataSourceKindNetworkMonitoring                                    DataSourceKind = "NetworkMonitoring"
	DataSourceKindOffice365                                            DataSourceKind = "Office365"
	DataSourceKindSQLDataClassification                                DataSourceKind = "SqlDataClassification"
	DataSourceKindSecurityCenterSecurityWindowsBaselineConfiguration   DataSourceKind = "SecurityCenterSecurityWindowsBaselineConfiguration"
	DataSourceKindSecurityEventCollectionConfiguration                 DataSourceKind = "SecurityEventCollectionConfiguration"
	DataSourceKindSecurityInsightsSecurityEventCollectionConfiguration DataSourceKind = "SecurityInsightsSecurityEventCollectionConfiguration"
	DataSourceKindSecurityWindowsBaselineConfiguration                 DataSourceKind = "SecurityWindowsBaselineConfiguration"
	DataSourceKindWindowsEvent                                         DataSourceKind = "WindowsEvent"
	DataSourceKindWindowsPerformanceCounter                            DataSourceKind = "WindowsPerformanceCounter"
	DataSourceKindWindowsTelemetry                                     DataSourceKind = "WindowsTelemetry"
)

// PossibleDataSourceKindValues returns the possible values for the DataSourceKind const type.
func PossibleDataSourceKindValues() []DataSourceKind {
	return []DataSourceKind{
		DataSourceKindApplicationInsights,
		DataSourceKindAzureActivityLog,
		DataSourceKindAzureAuditLog,
		DataSourceKindChangeTrackingContentLocation,
		DataSourceKindChangeTrackingCustomPath,
		DataSourceKindChangeTrackingDataTypeConfiguration,
		DataSourceKindChangeTrackingDefaultRegistry,
		DataSourceKindChangeTrackingLinuxPath,
		DataSourceKindChangeTrackingPath,
		DataSourceKindChangeTrackingRegistry,
		DataSourceKindChangeTrackingServices,
		DataSourceKindCustomLog,
		DataSourceKindCustomLogCollection,
		DataSourceKindDNSAnalytics,
		DataSourceKindGenericDataSource,
		DataSourceKindIISLogs,
		DataSourceKindImportComputerGroup,
		DataSourceKindItsm,
		DataSourceKindLinuxChangeTrackingPath,
		DataSourceKindLinuxPerformanceCollection,
		DataSourceKindLinuxPerformanceObject,
		DataSourceKindLinuxSyslog,
		DataSourceKindLinuxSyslogCollection,
		DataSourceKindNetworkMonitoring,
		DataSourceKindOffice365,
		DataSourceKindSQLDataClassification,
		DataSourceKindSecurityCenterSecurityWindowsBaselineConfiguration,
		DataSourceKindSecurityEventCollectionConfiguration,
		DataSourceKindSecurityInsightsSecurityEventCollectionConfiguration,
		DataSourceKindSecurityWindowsBaselineConfiguration,
		DataSourceKindWindowsEvent,
		DataSourceKindWindowsPerformanceCounter,
		DataSourceKindWindowsTelemetry,
	}
}

// DataSourceType - Linked storage accounts type.
type DataSourceType string

const (
	DataSourceTypeCustomLogs  DataSourceType = "CustomLogs"
	DataSourceTypeAzureWatson DataSourceType = "AzureWatson"
	DataSourceTypeQuery       DataSourceType = "Query"
	DataSourceTypeAlerts      DataSourceType = "Alerts"
)

// PossibleDataSourceTypeValues returns the possible values for the DataSourceType const type.
func PossibleDataSourceTypeValues() []DataSourceType {
	return []DataSourceType{
		DataSourceTypeCustomLogs,
		DataSourceTypeAzureWatson,
		DataSourceTypeQuery,
		DataSourceTypeAlerts,
	}
}

// IdentityType - The type of identity that creates/modifies resources
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "application"
	IdentityTypeKey             IdentityType = "key"
	IdentityTypeManagedIdentity IdentityType = "managedIdentity"
	IdentityTypeNone            IdentityType = "None"
	IdentityTypeSystemAssigned  IdentityType = "SystemAssigned"
	IdentityTypeUser            IdentityType = "user"
	IdentityTypeUserAssigned    IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeUser,
		IdentityTypeUserAssigned,
	}
}

// LinkedServiceEntityStatus - The provisioning state of the linked service.
type LinkedServiceEntityStatus string

const (
	LinkedServiceEntityStatusDeleting            LinkedServiceEntityStatus = "Deleting"
	LinkedServiceEntityStatusProvisioningAccount LinkedServiceEntityStatus = "ProvisioningAccount"
	LinkedServiceEntityStatusSucceeded           LinkedServiceEntityStatus = "Succeeded"
	LinkedServiceEntityStatusUpdating            LinkedServiceEntityStatus = "Updating"
)

// PossibleLinkedServiceEntityStatusValues returns the possible values for the LinkedServiceEntityStatus const type.
func PossibleLinkedServiceEntityStatusValues() []LinkedServiceEntityStatus {
	return []LinkedServiceEntityStatus{
		LinkedServiceEntityStatusDeleting,
		LinkedServiceEntityStatusProvisioningAccount,
		LinkedServiceEntityStatusSucceeded,
		LinkedServiceEntityStatusUpdating,
	}
}

// ProvisioningStateEnum - Table's current provisioning state. If set to 'updating', indicates a resource lock due to ongoing
// operation, forbidding any update to the table until the ongoing operation is concluded.
type ProvisioningStateEnum string

const (
	// ProvisioningStateEnumInProgress - Table schema is stable and without changes, table data is being updated.
	ProvisioningStateEnumInProgress ProvisioningStateEnum = "InProgress"
	// ProvisioningStateEnumSucceeded - Table state is stable and without changes, table is unlocked and open for new updates.
	ProvisioningStateEnumSucceeded ProvisioningStateEnum = "Succeeded"
	// ProvisioningStateEnumUpdating - Table schema is still being built and updated, table is currently locked for any changes
	// till the procedure is done.
	ProvisioningStateEnumUpdating ProvisioningStateEnum = "Updating"
)

// PossibleProvisioningStateEnumValues returns the possible values for the ProvisioningStateEnum const type.
func PossibleProvisioningStateEnumValues() []ProvisioningStateEnum {
	return []ProvisioningStateEnum{
		ProvisioningStateEnumInProgress,
		ProvisioningStateEnumSucceeded,
		ProvisioningStateEnumUpdating,
	}
}

// PublicNetworkAccessType - The network access type for operating on the Log Analytics Workspace. By default it is Enabled
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeDisabled - Disables public connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	// PublicNetworkAccessTypeEnabled - Enables connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
)

// PossiblePublicNetworkAccessTypeValues returns the possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{
		PublicNetworkAccessTypeDisabled,
		PublicNetworkAccessTypeEnabled,
	}
}

// PurgeState - Status of the operation represented by the requested Id.
type PurgeState string

const (
	PurgeStateCompleted PurgeState = "completed"
	PurgeStatePending   PurgeState = "pending"
)

// PossiblePurgeStateValues returns the possible values for the PurgeState const type.
func PossiblePurgeStateValues() []PurgeState {
	return []PurgeState{
		PurgeStateCompleted,
		PurgeStatePending,
	}
}

// SKUNameEnum - The name of the Service Tier.
type SKUNameEnum string

const (
	SKUNameEnumCapacityReservation SKUNameEnum = "CapacityReservation"
	SKUNameEnumFree                SKUNameEnum = "Free"
	SKUNameEnumPerGB2018           SKUNameEnum = "PerGB2018"
	SKUNameEnumPerNode             SKUNameEnum = "PerNode"
	SKUNameEnumPremium             SKUNameEnum = "Premium"
	SKUNameEnumStandalone          SKUNameEnum = "Standalone"
	SKUNameEnumStandard            SKUNameEnum = "Standard"
)

// PossibleSKUNameEnumValues returns the possible values for the SKUNameEnum const type.
func PossibleSKUNameEnumValues() []SKUNameEnum {
	return []SKUNameEnum{
		SKUNameEnumCapacityReservation,
		SKUNameEnumFree,
		SKUNameEnumPerGB2018,
		SKUNameEnumPerNode,
		SKUNameEnumPremium,
		SKUNameEnumStandalone,
		SKUNameEnumStandard,
	}
}

// SearchSortEnum - The sort order of the search.
type SearchSortEnum string

const (
	SearchSortEnumAsc  SearchSortEnum = "asc"
	SearchSortEnumDesc SearchSortEnum = "desc"
)

// PossibleSearchSortEnumValues returns the possible values for the SearchSortEnum const type.
func PossibleSearchSortEnumValues() []SearchSortEnum {
	return []SearchSortEnum{
		SearchSortEnumAsc,
		SearchSortEnumDesc,
	}
}

// SourceEnum - Table's creator.
type SourceEnum string

const (
	// SourceEnumCustomer - Tables created by the owner of the Workspace, and only found in this Workspace.
	SourceEnumCustomer SourceEnum = "customer"
	// SourceEnumMicrosoft - Tables provisioned by the system, as collected via Diagnostic Settings, the Agents, or any other
	// standard data collection means.
	SourceEnumMicrosoft SourceEnum = "microsoft"
)

// PossibleSourceEnumValues returns the possible values for the SourceEnum const type.
func PossibleSourceEnumValues() []SourceEnum {
	return []SourceEnum{
		SourceEnumCustomer,
		SourceEnumMicrosoft,
	}
}

// StorageInsightState - The state of the storage insight connection to the workspace
type StorageInsightState string

const (
	StorageInsightStateERROR StorageInsightState = "ERROR"
	StorageInsightStateOK    StorageInsightState = "OK"
)

// PossibleStorageInsightStateValues returns the possible values for the StorageInsightState const type.
func PossibleStorageInsightStateValues() []StorageInsightState {
	return []StorageInsightState{
		StorageInsightStateERROR,
		StorageInsightStateOK,
	}
}

// TablePlanEnum - Instruct the system how to handle and charge the logs ingested to this table.
type TablePlanEnum string

const (
	// TablePlanEnumAnalytics - Logs that allow monitoring and analytics.
	TablePlanEnumAnalytics TablePlanEnum = "Analytics"
	// TablePlanEnumBasic - Logs that are adjusted to support high volume low value verbose logs.
	TablePlanEnumBasic TablePlanEnum = "Basic"
)

// PossibleTablePlanEnumValues returns the possible values for the TablePlanEnum const type.
func PossibleTablePlanEnumValues() []TablePlanEnum {
	return []TablePlanEnum{
		TablePlanEnumAnalytics,
		TablePlanEnumBasic,
	}
}

// TableSubTypeEnum - The subtype describes what APIs can be used to interact with the table, and what features are available
// against it.
type TableSubTypeEnum string

const (
	// TableSubTypeEnumAny - The default subtype with which built-in tables are created.
	TableSubTypeEnumAny TableSubTypeEnum = "Any"
	// TableSubTypeEnumClassic - Indicates a table created through the Data Collector API or with the custom logs feature of the
	// MMA agent, or any table against which Custom Fields were created.
	TableSubTypeEnumClassic TableSubTypeEnum = "Classic"
	// TableSubTypeEnumDataCollectionRuleBased - A table eligible to have data sent into it via any of the means supported by
	// Data Collection Rules: the Data Collection Endpoint API, ingestion-time transformations, or any other mechanism provided
	// by Data Collection Rules
	TableSubTypeEnumDataCollectionRuleBased TableSubTypeEnum = "DataCollectionRuleBased"
)

// PossibleTableSubTypeEnumValues returns the possible values for the TableSubTypeEnum const type.
func PossibleTableSubTypeEnumValues() []TableSubTypeEnum {
	return []TableSubTypeEnum{
		TableSubTypeEnumAny,
		TableSubTypeEnumClassic,
		TableSubTypeEnumDataCollectionRuleBased,
	}
}

// TableTypeEnum - Table's creator.
type TableTypeEnum string

const (
	// TableTypeEnumCustomLog - Custom log table.
	TableTypeEnumCustomLog TableTypeEnum = "CustomLog"
	// TableTypeEnumMicrosoft - Standard data collected by Azure Monitor.
	TableTypeEnumMicrosoft TableTypeEnum = "Microsoft"
	// TableTypeEnumRestoredLogs - Restored data.
	TableTypeEnumRestoredLogs TableTypeEnum = "RestoredLogs"
	// TableTypeEnumSearchResults - Data collected by a search job.
	TableTypeEnumSearchResults TableTypeEnum = "SearchResults"
)

// PossibleTableTypeEnumValues returns the possible values for the TableTypeEnum const type.
func PossibleTableTypeEnumValues() []TableTypeEnum {
	return []TableTypeEnum{
		TableTypeEnumCustomLog,
		TableTypeEnumMicrosoft,
		TableTypeEnumRestoredLogs,
		TableTypeEnumSearchResults,
	}
}

// Type - The type of the destination resource
type Type string

const (
	TypeEventHub       Type = "EventHub"
	TypeStorageAccount Type = "StorageAccount"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeEventHub,
		TypeStorageAccount,
	}
}

// WorkspaceEntityStatus - The provisioning state of the workspace.
type WorkspaceEntityStatus string

const (
	WorkspaceEntityStatusCanceled            WorkspaceEntityStatus = "Canceled"
	WorkspaceEntityStatusCreating            WorkspaceEntityStatus = "Creating"
	WorkspaceEntityStatusDeleting            WorkspaceEntityStatus = "Deleting"
	WorkspaceEntityStatusFailed              WorkspaceEntityStatus = "Failed"
	WorkspaceEntityStatusProvisioningAccount WorkspaceEntityStatus = "ProvisioningAccount"
	WorkspaceEntityStatusSucceeded           WorkspaceEntityStatus = "Succeeded"
	WorkspaceEntityStatusUpdating            WorkspaceEntityStatus = "Updating"
)

// PossibleWorkspaceEntityStatusValues returns the possible values for the WorkspaceEntityStatus const type.
func PossibleWorkspaceEntityStatusValues() []WorkspaceEntityStatus {
	return []WorkspaceEntityStatus{
		WorkspaceEntityStatusCanceled,
		WorkspaceEntityStatusCreating,
		WorkspaceEntityStatusDeleting,
		WorkspaceEntityStatusFailed,
		WorkspaceEntityStatusProvisioningAccount,
		WorkspaceEntityStatusSucceeded,
		WorkspaceEntityStatusUpdating,
	}
}

// WorkspaceSKUNameEnum - The name of the SKU.
type WorkspaceSKUNameEnum string

const (
	WorkspaceSKUNameEnumCapacityReservation WorkspaceSKUNameEnum = "CapacityReservation"
	WorkspaceSKUNameEnumFree                WorkspaceSKUNameEnum = "Free"
	WorkspaceSKUNameEnumLACluster           WorkspaceSKUNameEnum = "LACluster"
	WorkspaceSKUNameEnumPerGB2018           WorkspaceSKUNameEnum = "PerGB2018"
	WorkspaceSKUNameEnumPerNode             WorkspaceSKUNameEnum = "PerNode"
	WorkspaceSKUNameEnumPremium             WorkspaceSKUNameEnum = "Premium"
	WorkspaceSKUNameEnumStandalone          WorkspaceSKUNameEnum = "Standalone"
	WorkspaceSKUNameEnumStandard            WorkspaceSKUNameEnum = "Standard"
)

// PossibleWorkspaceSKUNameEnumValues returns the possible values for the WorkspaceSKUNameEnum const type.
func PossibleWorkspaceSKUNameEnumValues() []WorkspaceSKUNameEnum {
	return []WorkspaceSKUNameEnum{
		WorkspaceSKUNameEnumCapacityReservation,
		WorkspaceSKUNameEnumFree,
		WorkspaceSKUNameEnumLACluster,
		WorkspaceSKUNameEnumPerGB2018,
		WorkspaceSKUNameEnumPerNode,
		WorkspaceSKUNameEnumPremium,
		WorkspaceSKUNameEnumStandalone,
		WorkspaceSKUNameEnumStandard,
	}
}
