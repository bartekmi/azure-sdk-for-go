//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccountBackupsClient contains the methods for the AccountBackups group.
// Don't use this type directly, use NewAccountBackupsClient() instead.
type AccountBackupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountBackupsClient creates a new instance of AccountBackupsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccountBackupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountBackupsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AccountBackupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Delete the specified Backup for a Netapp Account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// backupName - The name of the backup
// options - AccountBackupsClientBeginDeleteOptions contains the optional parameters for the AccountBackupsClient.BeginDelete
// method.
func (client *AccountBackupsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, backupName string, options *AccountBackupsClientBeginDeleteOptions) (*armruntime.Poller[AccountBackupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, backupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[AccountBackupsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[AccountBackupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the specified Backup for a Netapp Account
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AccountBackupsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, backupName string, options *AccountBackupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, backupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccountBackupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupName string, options *AccountBackupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/accountBackups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified backup for a Netapp Account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// backupName - The name of the backup
// options - AccountBackupsClientGetOptions contains the optional parameters for the AccountBackupsClient.Get method.
func (client *AccountBackupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, backupName string, options *AccountBackupsClientGetOptions) (AccountBackupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, backupName, options)
	if err != nil {
		return AccountBackupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountBackupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountBackupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountBackupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupName string, options *AccountBackupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/accountBackups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountBackupsClient) getHandleResponse(resp *http.Response) (AccountBackupsClientGetResponse, error) {
	result := AccountBackupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Backup); err != nil {
		return AccountBackupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Backups for a Netapp Account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// options - AccountBackupsClientListOptions contains the optional parameters for the AccountBackupsClient.List method.
func (client *AccountBackupsClient) NewListPager(resourceGroupName string, accountName string, options *AccountBackupsClientListOptions) *runtime.Pager[AccountBackupsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AccountBackupsClientListResponse]{
		More: func(page AccountBackupsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AccountBackupsClientListResponse) (AccountBackupsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return AccountBackupsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccountBackupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccountBackupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccountBackupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountBackupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/accountBackups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccountBackupsClient) listHandleResponse(resp *http.Response) (AccountBackupsClientListResponse, error) {
	result := AccountBackupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupsList); err != nil {
		return AccountBackupsClientListResponse{}, err
	}
	return result, nil
}
