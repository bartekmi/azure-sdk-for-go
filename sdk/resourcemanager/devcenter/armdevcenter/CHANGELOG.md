# Release History

## 0.3.0 (2022-10-27)
### Breaking Changes

- Type of `OperationStatus.Error` has been changed from `*OperationStatusError` to `*ErrorDetail`
- Struct `OperationStatusError` has been removed

### Features Added

- New const `CatalogSyncStateFailed`
- New const `CatalogSyncStateSucceeded`
- New const `CatalogSyncStateInProgress`
- New const `CatalogSyncStateCanceled`
- New type alias `CatalogSyncState`
- New function `PossibleCatalogSyncStateValues() []CatalogSyncState`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `OperationStatusResult`
- New anonymous field `Schedule` in struct `SchedulesClientUpdateResponse`
- New field `Operations` in struct `OperationStatus`
- New field `ResourceID` in struct `OperationStatus`
- New field `SyncState` in struct `CatalogProperties`


## 0.2.0 (2022-09-29)
### Features Added

- New function `NewProjectAllowedEnvironmentTypesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ProjectAllowedEnvironmentTypesClient, error)`
- New function `*ProjectAllowedEnvironmentTypesClient.NewListPager(string, string, *ProjectAllowedEnvironmentTypesClientListOptions) *runtime.Pager[ProjectAllowedEnvironmentTypesClientListResponse]`
- New function `*ProjectAllowedEnvironmentTypesClient.Get(context.Context, string, string, string, *ProjectAllowedEnvironmentTypesClientGetOptions) (ProjectAllowedEnvironmentTypesClientGetResponse, error)`
- New struct `AllowedEnvironmentType`
- New struct `AllowedEnvironmentTypeListResult`
- New struct `AllowedEnvironmentTypeProperties`
- New struct `ProjectAllowedEnvironmentTypesClient`
- New struct `ProjectAllowedEnvironmentTypesClientGetOptions`
- New struct `ProjectAllowedEnvironmentTypesClientGetResponse`
- New struct `ProjectAllowedEnvironmentTypesClientListOptions`
- New struct `ProjectAllowedEnvironmentTypesClientListResponse`


## 0.1.0 (2022-07-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).