//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

// AddonsClientCreateOrUpdateResponse contains the response from method AddonsClient.CreateOrUpdate.
type AddonsClientCreateOrUpdateResponse struct {
	AddonClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddonsClientCreateOrUpdateResponse.
func (a *AddonsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAddonClassification(data)
	if err != nil {
		return err
	}
	a.AddonClassification = res
	return nil
}

// AddonsClientDeleteResponse contains the response from method AddonsClient.Delete.
type AddonsClientDeleteResponse struct {
	// placeholder for future response values
}

// AddonsClientGetResponse contains the response from method AddonsClient.Get.
type AddonsClientGetResponse struct {
	AddonClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddonsClientGetResponse.
func (a *AddonsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAddonClassification(data)
	if err != nil {
		return err
	}
	a.AddonClassification = res
	return nil
}

// AddonsClientListByRoleResponse contains the response from method AddonsClient.ListByRole.
type AddonsClientListByRoleResponse struct {
	AddonList
}

// AlertsClientGetResponse contains the response from method AlertsClient.Get.
type AlertsClientGetResponse struct {
	Alert
}

// AlertsClientListByDataBoxEdgeDeviceResponse contains the response from method AlertsClient.ListByDataBoxEdgeDevice.
type AlertsClientListByDataBoxEdgeDeviceResponse struct {
	AlertList
}

// AvailableSKUsClientListResponse contains the response from method AvailableSKUsClient.List.
type AvailableSKUsClientListResponse struct {
	SKUList
}

// BandwidthSchedulesClientCreateOrUpdateResponse contains the response from method BandwidthSchedulesClient.CreateOrUpdate.
type BandwidthSchedulesClientCreateOrUpdateResponse struct {
	BandwidthSchedule
}

// BandwidthSchedulesClientDeleteResponse contains the response from method BandwidthSchedulesClient.Delete.
type BandwidthSchedulesClientDeleteResponse struct {
	// placeholder for future response values
}

// BandwidthSchedulesClientGetResponse contains the response from method BandwidthSchedulesClient.Get.
type BandwidthSchedulesClientGetResponse struct {
	BandwidthSchedule
}

// BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse contains the response from method BandwidthSchedulesClient.ListByDataBoxEdgeDevice.
type BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse struct {
	BandwidthSchedulesList
}

// ContainersClientCreateOrUpdateResponse contains the response from method ContainersClient.CreateOrUpdate.
type ContainersClientCreateOrUpdateResponse struct {
	Container
}

// ContainersClientDeleteResponse contains the response from method ContainersClient.Delete.
type ContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainersClientGetResponse contains the response from method ContainersClient.Get.
type ContainersClientGetResponse struct {
	Container
}

// ContainersClientListByStorageAccountResponse contains the response from method ContainersClient.ListByStorageAccount.
type ContainersClientListByStorageAccountResponse struct {
	ContainerList
}

// ContainersClientRefreshResponse contains the response from method ContainersClient.Refresh.
type ContainersClientRefreshResponse struct {
	// placeholder for future response values
}

// DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse contains the response from method DeviceCapacityCheckClient.CheckResourceCreationFeasibility.
type DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse struct {
	// placeholder for future response values
}

// DeviceCapacityInfoClientGetDeviceCapacityInfoResponse contains the response from method DeviceCapacityInfoClient.GetDeviceCapacityInfo.
type DeviceCapacityInfoClientGetDeviceCapacityInfoResponse struct {
	DeviceCapacityInfo
}

// DevicesClientCreateOrUpdateResponse contains the response from method DevicesClient.CreateOrUpdate.
type DevicesClientCreateOrUpdateResponse struct {
	Device
}

// DevicesClientCreateOrUpdateSecuritySettingsResponse contains the response from method DevicesClient.CreateOrUpdateSecuritySettings.
type DevicesClientCreateOrUpdateSecuritySettingsResponse struct {
	// placeholder for future response values
}

// DevicesClientDeleteResponse contains the response from method DevicesClient.Delete.
type DevicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DevicesClientDownloadUpdatesResponse contains the response from method DevicesClient.DownloadUpdates.
type DevicesClientDownloadUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientGenerateCertificateResponse contains the response from method DevicesClient.GenerateCertificate.
type DevicesClientGenerateCertificateResponse struct {
	GenerateCertResponse
}

// DevicesClientGetExtendedInformationResponse contains the response from method DevicesClient.GetExtendedInformation.
type DevicesClientGetExtendedInformationResponse struct {
	DeviceExtendedInfo
}

// DevicesClientGetNetworkSettingsResponse contains the response from method DevicesClient.GetNetworkSettings.
type DevicesClientGetNetworkSettingsResponse struct {
	NetworkSettings
}

// DevicesClientGetResponse contains the response from method DevicesClient.Get.
type DevicesClientGetResponse struct {
	Device
}

// DevicesClientGetUpdateSummaryResponse contains the response from method DevicesClient.GetUpdateSummary.
type DevicesClientGetUpdateSummaryResponse struct {
	UpdateSummary
}

// DevicesClientInstallUpdatesResponse contains the response from method DevicesClient.InstallUpdates.
type DevicesClientInstallUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientListByResourceGroupResponse contains the response from method DevicesClient.ListByResourceGroup.
type DevicesClientListByResourceGroupResponse struct {
	DeviceList
}

// DevicesClientListBySubscriptionResponse contains the response from method DevicesClient.ListBySubscription.
type DevicesClientListBySubscriptionResponse struct {
	DeviceList
}

// DevicesClientScanForUpdatesResponse contains the response from method DevicesClient.ScanForUpdates.
type DevicesClientScanForUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientUpdateExtendedInformationResponse contains the response from method DevicesClient.UpdateExtendedInformation.
type DevicesClientUpdateExtendedInformationResponse struct {
	DeviceExtendedInfo
}

// DevicesClientUpdateResponse contains the response from method DevicesClient.Update.
type DevicesClientUpdateResponse struct {
	Device
}

// DevicesClientUploadCertificateResponse contains the response from method DevicesClient.UploadCertificate.
type DevicesClientUploadCertificateResponse struct {
	UploadCertificateResponse
}

// DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse contains the response from method DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings.
type DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse struct {
	DiagnosticProactiveLogCollectionSettings
}

// DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse contains the response from method DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings.
type DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse struct {
	DiagnosticRemoteSupportSettings
}

// DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse contains the response from method DiagnosticSettingsClient.UpdateDiagnosticProactiveLogCollectionSettings.
type DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse struct {
	DiagnosticProactiveLogCollectionSettings
}

// DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse contains the response from method DiagnosticSettingsClient.UpdateDiagnosticRemoteSupportSettings.
type DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse struct {
	DiagnosticRemoteSupportSettings
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	Job
}

// MonitoringConfigClientCreateOrUpdateResponse contains the response from method MonitoringConfigClient.CreateOrUpdate.
type MonitoringConfigClientCreateOrUpdateResponse struct {
	MonitoringMetricConfiguration
}

// MonitoringConfigClientDeleteResponse contains the response from method MonitoringConfigClient.Delete.
type MonitoringConfigClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitoringConfigClientGetResponse contains the response from method MonitoringConfigClient.Get.
type MonitoringConfigClientGetResponse struct {
	MonitoringMetricConfiguration
}

// MonitoringConfigClientListResponse contains the response from method MonitoringConfigClient.List.
type MonitoringConfigClientListResponse struct {
	MonitoringMetricConfigurationList
}

// NodesClientListByDataBoxEdgeDeviceResponse contains the response from method NodesClient.ListByDataBoxEdgeDevice.
type NodesClientListByDataBoxEdgeDeviceResponse struct {
	NodeList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsList
}

// OperationsStatusClientGetResponse contains the response from method OperationsStatusClient.Get.
type OperationsStatusClientGetResponse struct {
	Job
}

// OrdersClientCreateOrUpdateResponse contains the response from method OrdersClient.CreateOrUpdate.
type OrdersClientCreateOrUpdateResponse struct {
	Order
}

// OrdersClientDeleteResponse contains the response from method OrdersClient.Delete.
type OrdersClientDeleteResponse struct {
	// placeholder for future response values
}

// OrdersClientGetResponse contains the response from method OrdersClient.Get.
type OrdersClientGetResponse struct {
	Order
}

// OrdersClientListByDataBoxEdgeDeviceResponse contains the response from method OrdersClient.ListByDataBoxEdgeDevice.
type OrdersClientListByDataBoxEdgeDeviceResponse struct {
	OrderList
}

// OrdersClientListDCAccessCodeResponse contains the response from method OrdersClient.ListDCAccessCode.
type OrdersClientListDCAccessCodeResponse struct {
	DCAccessCode
}

// RolesClientCreateOrUpdateResponse contains the response from method RolesClient.CreateOrUpdate.
type RolesClientCreateOrUpdateResponse struct {
	RoleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RolesClientCreateOrUpdateResponse.
func (r *RolesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalRoleClassification(data)
	if err != nil {
		return err
	}
	r.RoleClassification = res
	return nil
}

// RolesClientDeleteResponse contains the response from method RolesClient.Delete.
type RolesClientDeleteResponse struct {
	// placeholder for future response values
}

// RolesClientGetResponse contains the response from method RolesClient.Get.
type RolesClientGetResponse struct {
	RoleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RolesClientGetResponse.
func (r *RolesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalRoleClassification(data)
	if err != nil {
		return err
	}
	r.RoleClassification = res
	return nil
}

// RolesClientListByDataBoxEdgeDeviceResponse contains the response from method RolesClient.ListByDataBoxEdgeDevice.
type RolesClientListByDataBoxEdgeDeviceResponse struct {
	RoleList
}

// SharesClientCreateOrUpdateResponse contains the response from method SharesClient.CreateOrUpdate.
type SharesClientCreateOrUpdateResponse struct {
	Share
}

// SharesClientDeleteResponse contains the response from method SharesClient.Delete.
type SharesClientDeleteResponse struct {
	// placeholder for future response values
}

// SharesClientGetResponse contains the response from method SharesClient.Get.
type SharesClientGetResponse struct {
	Share
}

// SharesClientListByDataBoxEdgeDeviceResponse contains the response from method SharesClient.ListByDataBoxEdgeDevice.
type SharesClientListByDataBoxEdgeDeviceResponse struct {
	ShareList
}

// SharesClientRefreshResponse contains the response from method SharesClient.Refresh.
type SharesClientRefreshResponse struct {
	// placeholder for future response values
}

// StorageAccountCredentialsClientCreateOrUpdateResponse contains the response from method StorageAccountCredentialsClient.CreateOrUpdate.
type StorageAccountCredentialsClientCreateOrUpdateResponse struct {
	StorageAccountCredential
}

// StorageAccountCredentialsClientDeleteResponse contains the response from method StorageAccountCredentialsClient.Delete.
type StorageAccountCredentialsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountCredentialsClientGetResponse contains the response from method StorageAccountCredentialsClient.Get.
type StorageAccountCredentialsClientGetResponse struct {
	StorageAccountCredential
}

// StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse contains the response from method StorageAccountCredentialsClient.ListByDataBoxEdgeDevice.
type StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse struct {
	StorageAccountCredentialList
}

// StorageAccountsClientCreateOrUpdateResponse contains the response from method StorageAccountsClient.CreateOrUpdate.
type StorageAccountsClientCreateOrUpdateResponse struct {
	StorageAccount
}

// StorageAccountsClientDeleteResponse contains the response from method StorageAccountsClient.Delete.
type StorageAccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountsClientGetResponse contains the response from method StorageAccountsClient.Get.
type StorageAccountsClientGetResponse struct {
	StorageAccount
}

// StorageAccountsClientListByDataBoxEdgeDeviceResponse contains the response from method StorageAccountsClient.ListByDataBoxEdgeDevice.
type StorageAccountsClientListByDataBoxEdgeDeviceResponse struct {
	StorageAccountList
}

// SupportPackagesClientTriggerSupportPackageResponse contains the response from method SupportPackagesClient.TriggerSupportPackage.
type SupportPackagesClientTriggerSupportPackageResponse struct {
	// placeholder for future response values
}

// TriggersClientCreateOrUpdateResponse contains the response from method TriggersClient.CreateOrUpdate.
type TriggersClientCreateOrUpdateResponse struct {
	TriggerClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientCreateOrUpdateResponse.
func (t *TriggersClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// TriggersClientDeleteResponse contains the response from method TriggersClient.Delete.
type TriggersClientDeleteResponse struct {
	// placeholder for future response values
}

// TriggersClientGetResponse contains the response from method TriggersClient.Get.
type TriggersClientGetResponse struct {
	TriggerClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientGetResponse.
func (t *TriggersClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// TriggersClientListByDataBoxEdgeDeviceResponse contains the response from method TriggersClient.ListByDataBoxEdgeDevice.
type TriggersClientListByDataBoxEdgeDeviceResponse struct {
	TriggerList
}

// UsersClientCreateOrUpdateResponse contains the response from method UsersClient.CreateOrUpdate.
type UsersClientCreateOrUpdateResponse struct {
	User
}

// UsersClientDeleteResponse contains the response from method UsersClient.Delete.
type UsersClientDeleteResponse struct {
	// placeholder for future response values
}

// UsersClientGetResponse contains the response from method UsersClient.Get.
type UsersClientGetResponse struct {
	User
}

// UsersClientListByDataBoxEdgeDeviceResponse contains the response from method UsersClient.ListByDataBoxEdgeDevice.
type UsersClientListByDataBoxEdgeDeviceResponse struct {
	UserList
}
