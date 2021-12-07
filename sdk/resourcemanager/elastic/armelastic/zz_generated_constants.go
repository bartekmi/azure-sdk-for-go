//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelastic

const (
	module  = "armelastic"
	version = "v0.1.0"
)

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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// ElasticDeploymentStatus - Flag specifying if the Elastic deployment status is healthy or not.
type ElasticDeploymentStatus string

const (
	ElasticDeploymentStatusHealthy   ElasticDeploymentStatus = "Healthy"
	ElasticDeploymentStatusUnhealthy ElasticDeploymentStatus = "Unhealthy"
)

// PossibleElasticDeploymentStatusValues returns the possible values for the ElasticDeploymentStatus const type.
func PossibleElasticDeploymentStatusValues() []ElasticDeploymentStatus {
	return []ElasticDeploymentStatus{
		ElasticDeploymentStatusHealthy,
		ElasticDeploymentStatusUnhealthy,
	}
}

// ToPtr returns a *ElasticDeploymentStatus pointing to the current value.
func (c ElasticDeploymentStatus) ToPtr() *ElasticDeploymentStatus {
	return &c
}

type LiftrResourceCategories string

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = "MonitorLogs"
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = "Unknown"
)

// PossibleLiftrResourceCategoriesValues returns the possible values for the LiftrResourceCategories const type.
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return []LiftrResourceCategories{
		LiftrResourceCategoriesMonitorLogs,
		LiftrResourceCategoriesUnknown,
	}
}

// ToPtr returns a *LiftrResourceCategories pointing to the current value.
func (c LiftrResourceCategories) ToPtr() *LiftrResourceCategories {
	return &c
}

// ManagedIdentityTypes - Managed Identity types.
type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = "SystemAssigned"
)

// PossibleManagedIdentityTypesValues returns the possible values for the ManagedIdentityTypes const type.
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return []ManagedIdentityTypes{
		ManagedIdentityTypesSystemAssigned,
	}
}

// ToPtr returns a *ManagedIdentityTypes pointing to the current value.
func (c ManagedIdentityTypes) ToPtr() *ManagedIdentityTypes {
	return &c
}

// MonitoringStatus - Flag specifying if the resource monitoring is enabled or disabled.
type MonitoringStatus string

const (
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
)

// PossibleMonitoringStatusValues returns the possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{
		MonitoringStatusDisabled,
		MonitoringStatusEnabled,
	}
}

// ToPtr returns a *MonitoringStatus pointing to the current value.
func (c MonitoringStatus) ToPtr() *MonitoringStatus {
	return &c
}

// OperationName - Operation to be performed on the given vm resource id.
type OperationName string

const (
	OperationNameAdd    OperationName = "Add"
	OperationNameDelete OperationName = "Delete"
)

// PossibleOperationNameValues returns the possible values for the OperationName const type.
func PossibleOperationNameValues() []OperationName {
	return []OperationName{
		OperationNameAdd,
		OperationNameDelete,
	}
}

// ToPtr returns a *OperationName pointing to the current value.
func (c OperationName) ToPtr() *OperationName {
	return &c
}

// ProvisioningState - Provisioning state of Elastic resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// SendingLogs - Flag indicating the status of the resource for sending logs operation to Elastic.
type SendingLogs string

const (
	SendingLogsFalse SendingLogs = "False"
	SendingLogsTrue  SendingLogs = "True"
)

// PossibleSendingLogsValues returns the possible values for the SendingLogs const type.
func PossibleSendingLogsValues() []SendingLogs {
	return []SendingLogs{
		SendingLogsFalse,
		SendingLogsTrue,
	}
}

// ToPtr returns a *SendingLogs pointing to the current value.
func (c SendingLogs) ToPtr() *SendingLogs {
	return &c
}

// TagAction - Valid actions for a filtering tag. Exclusion takes priority over inclusion.
type TagAction string

const (
	TagActionExclude TagAction = "Exclude"
	TagActionInclude TagAction = "Include"
)

// PossibleTagActionValues returns the possible values for the TagAction const type.
func PossibleTagActionValues() []TagAction {
	return []TagAction{
		TagActionExclude,
		TagActionInclude,
	}
}

// ToPtr returns a *TagAction pointing to the current value.
func (c TagAction) ToPtr() *TagAction {
	return &c
}