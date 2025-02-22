//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

// DicomServicesClientCreateOrUpdateResponse contains the response from method DicomServicesClient.CreateOrUpdate.
type DicomServicesClientCreateOrUpdateResponse struct {
	DicomService
}

// DicomServicesClientDeleteResponse contains the response from method DicomServicesClient.Delete.
type DicomServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DicomServicesClientGetResponse contains the response from method DicomServicesClient.Get.
type DicomServicesClientGetResponse struct {
	DicomService
}

// DicomServicesClientListByWorkspaceResponse contains the response from method DicomServicesClient.ListByWorkspace.
type DicomServicesClientListByWorkspaceResponse struct {
	DicomServiceCollection
}

// DicomServicesClientUpdateResponse contains the response from method DicomServicesClient.Update.
type DicomServicesClientUpdateResponse struct {
	DicomService
}

// FhirDestinationsClientListByIotConnectorResponse contains the response from method FhirDestinationsClient.ListByIotConnector.
type FhirDestinationsClientListByIotConnectorResponse struct {
	IotFhirDestinationCollection
}

// FhirServicesClientCreateOrUpdateResponse contains the response from method FhirServicesClient.CreateOrUpdate.
type FhirServicesClientCreateOrUpdateResponse struct {
	FhirService
}

// FhirServicesClientDeleteResponse contains the response from method FhirServicesClient.Delete.
type FhirServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// FhirServicesClientGetResponse contains the response from method FhirServicesClient.Get.
type FhirServicesClientGetResponse struct {
	FhirService
}

// FhirServicesClientListByWorkspaceResponse contains the response from method FhirServicesClient.ListByWorkspace.
type FhirServicesClientListByWorkspaceResponse struct {
	FhirServiceCollection
}

// FhirServicesClientUpdateResponse contains the response from method FhirServicesClient.Update.
type FhirServicesClientUpdateResponse struct {
	FhirService
}

// IotConnectorFhirDestinationClientCreateOrUpdateResponse contains the response from method IotConnectorFhirDestinationClient.CreateOrUpdate.
type IotConnectorFhirDestinationClientCreateOrUpdateResponse struct {
	IotFhirDestination
}

// IotConnectorFhirDestinationClientDeleteResponse contains the response from method IotConnectorFhirDestinationClient.Delete.
type IotConnectorFhirDestinationClientDeleteResponse struct {
	// placeholder for future response values
}

// IotConnectorFhirDestinationClientGetResponse contains the response from method IotConnectorFhirDestinationClient.Get.
type IotConnectorFhirDestinationClientGetResponse struct {
	IotFhirDestination
}

// IotConnectorsClientCreateOrUpdateResponse contains the response from method IotConnectorsClient.CreateOrUpdate.
type IotConnectorsClientCreateOrUpdateResponse struct {
	IotConnector
}

// IotConnectorsClientDeleteResponse contains the response from method IotConnectorsClient.Delete.
type IotConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// IotConnectorsClientGetResponse contains the response from method IotConnectorsClient.Get.
type IotConnectorsClientGetResponse struct {
	IotConnector
}

// IotConnectorsClientListByWorkspaceResponse contains the response from method IotConnectorsClient.ListByWorkspace.
type IotConnectorsClientListByWorkspaceResponse struct {
	IotConnectorCollection
}

// IotConnectorsClientUpdateResponse contains the response from method IotConnectorsClient.Update.
type IotConnectorsClientUpdateResponse struct {
	IotConnector
}

// OperationResultsClientGetResponse contains the response from method OperationResultsClient.Get.
type OperationResultsClientGetResponse struct {
	OperationResultsDescription
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ListOperations
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnectionDescription
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnectionDescription
}

// PrivateEndpointConnectionsClientListByServiceResponse contains the response from method PrivateEndpointConnectionsClient.ListByService.
type PrivateEndpointConnectionsClientListByServiceResponse struct {
	PrivateEndpointConnectionListResultDescription
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResourceDescription
}

// PrivateLinkResourcesClientListByServiceResponse contains the response from method PrivateLinkResourcesClient.ListByService.
type PrivateLinkResourcesClientListByServiceResponse struct {
	PrivateLinkResourceListResultDescription
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	ServicesNameAvailabilityInfo
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServicesDescription
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServicesDescription
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResponse struct {
	ServicesDescriptionListResult
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	ServicesDescriptionListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	ServicesDescription
}

// WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.CreateOrUpdate.
type WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnectionDescription
}

// WorkspacePrivateEndpointConnectionsClientDeleteResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.Delete.
type WorkspacePrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacePrivateEndpointConnectionsClientGetResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.Get.
type WorkspacePrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnectionDescription
}

// WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.ListByWorkspace.
type WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse struct {
	PrivateEndpointConnectionListResultDescription
}

// WorkspacePrivateLinkResourcesClientGetResponse contains the response from method WorkspacePrivateLinkResourcesClient.Get.
type WorkspacePrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResourceDescription
}

// WorkspacePrivateLinkResourcesClientListByWorkspaceResponse contains the response from method WorkspacePrivateLinkResourcesClient.ListByWorkspace.
type WorkspacePrivateLinkResourcesClientListByWorkspaceResponse struct {
	PrivateLinkResourceListResultDescription
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.CreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.Delete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.ListByResourceGroup.
type WorkspacesClientListByResourceGroupResponse struct {
	WorkspaceList
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.ListBySubscription.
type WorkspacesClientListBySubscriptionResponse struct {
	WorkspaceList
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.Update.
type WorkspacesClientUpdateResponse struct {
	Workspace
}
