//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConfigurationProfilesClient contains the methods for the ConfigurationProfiles group.
// Don't use this type directly, use NewConfigurationProfilesClient() instead.
type ConfigurationProfilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationProfilesClient creates a new instance of ConfigurationProfilesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfigurationProfilesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ConfigurationProfilesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
// configurationProfileName - Name of the configuration profile.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// parameters - Parameters supplied to create or update configuration profile.
// options - ConfigurationProfilesClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationProfilesClient.CreateOrUpdate
// method.
func (client *ConfigurationProfilesClient) CreateOrUpdate(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfile, options *ConfigurationProfilesClientCreateOrUpdateOptions) (ConfigurationProfilesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, configurationProfileName, resourceGroupName, parameters, options)
	if err != nil {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationProfilesClient) createOrUpdateCreateRequest(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfile, options *ConfigurationProfilesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationProfilesClientCreateOrUpdateResponse, error) {
	result := ConfigurationProfilesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfile); err != nil {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// configurationProfileName - Name of the configuration profile
// options - ConfigurationProfilesClientDeleteOptions contains the optional parameters for the ConfigurationProfilesClient.Delete
// method.
func (client *ConfigurationProfilesClient) Delete(ctx context.Context, resourceGroupName string, configurationProfileName string, options *ConfigurationProfilesClientDeleteOptions) (ConfigurationProfilesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configurationProfileName, options)
	if err != nil {
		return ConfigurationProfilesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfilesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConfigurationProfilesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConfigurationProfilesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileName string, options *ConfigurationProfilesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get information about a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
// configurationProfileName - The configuration profile name.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ConfigurationProfilesClientGetOptions contains the optional parameters for the ConfigurationProfilesClient.Get
// method.
func (client *ConfigurationProfilesClient) Get(ctx context.Context, configurationProfileName string, resourceGroupName string, options *ConfigurationProfilesClientGetOptions) (ConfigurationProfilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, configurationProfileName, resourceGroupName, options)
	if err != nil {
		return ConfigurationProfilesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationProfilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationProfilesClient) getCreateRequest(ctx context.Context, configurationProfileName string, resourceGroupName string, options *ConfigurationProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationProfilesClient) getHandleResponse(resp *http.Response) (ConfigurationProfilesClientGetResponse, error) {
	result := ConfigurationProfilesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfile); err != nil {
		return ConfigurationProfilesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieve a list of configuration profile within a given resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ConfigurationProfilesClientListByResourceGroupOptions contains the optional parameters for the ConfigurationProfilesClient.ListByResourceGroup
// method.
func (client *ConfigurationProfilesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ConfigurationProfilesClientListByResourceGroupOptions) (ConfigurationProfilesClientListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ConfigurationProfilesClientListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfilesClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationProfilesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationProfilesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationProfilesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationProfilesClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationProfilesClientListByResourceGroupResponse, error) {
	result := ConfigurationProfilesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileList); err != nil {
		return ConfigurationProfilesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieve a list of configuration profile within a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - ConfigurationProfilesClientListBySubscriptionOptions contains the optional parameters for the ConfigurationProfilesClient.ListBySubscription
// method.
func (client *ConfigurationProfilesClient) ListBySubscription(ctx context.Context, options *ConfigurationProfilesClientListBySubscriptionOptions) (ConfigurationProfilesClientListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ConfigurationProfilesClientListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfilesClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationProfilesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationProfilesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationProfilesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Automanage/configurationProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConfigurationProfilesClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationProfilesClientListBySubscriptionResponse, error) {
	result := ConfigurationProfilesClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileList); err != nil {
		return ConfigurationProfilesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
// configurationProfileName - Name of the configuration profile.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// parameters - Parameters supplied to update configuration profile.
// options - ConfigurationProfilesClientUpdateOptions contains the optional parameters for the ConfigurationProfilesClient.Update
// method.
func (client *ConfigurationProfilesClient) Update(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfileUpdate, options *ConfigurationProfilesClientUpdateOptions) (ConfigurationProfilesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, configurationProfileName, resourceGroupName, parameters, options)
	if err != nil {
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationProfilesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ConfigurationProfilesClient) updateCreateRequest(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfileUpdate, options *ConfigurationProfilesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ConfigurationProfilesClient) updateHandleResponse(resp *http.Response) (ConfigurationProfilesClientUpdateResponse, error) {
	result := ConfigurationProfilesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfile); err != nil {
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	return result, nil
}