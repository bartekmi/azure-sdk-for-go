//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// DeviceCapacityCheckClient contains the methods for the DeviceCapacityCheck group.
// Don't use this type directly, use NewDeviceCapacityCheckClient() instead.
type DeviceCapacityCheckClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDeviceCapacityCheckClient creates a new instance of DeviceCapacityCheckClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDeviceCapacityCheckClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeviceCapacityCheckClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DeviceCapacityCheckClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCheckResourceCreationFeasibility - Posts the device capacity request info to check feasibility.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - The resource group name.
// deviceName - The device name.
// deviceCapacityRequestInfo - The device capacity request info.
// options - DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions contains the optional parameters for the
// DeviceCapacityCheckClient.BeginCheckResourceCreationFeasibility method.
func (client *DeviceCapacityCheckClient) BeginCheckResourceCreationFeasibility(ctx context.Context, resourceGroupName string, deviceName string, deviceCapacityRequestInfo DeviceCapacityRequestInfo, options *DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions) (*runtime.Poller[DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.checkResourceCreationFeasibility(ctx, resourceGroupName, deviceName, deviceCapacityRequestInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse](options.ResumeToken, client.pl, nil)
	}
}

// CheckResourceCreationFeasibility - Posts the device capacity request info to check feasibility.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
func (client *DeviceCapacityCheckClient) checkResourceCreationFeasibility(ctx context.Context, resourceGroupName string, deviceName string, deviceCapacityRequestInfo DeviceCapacityRequestInfo, options *DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions) (*http.Response, error) {
	req, err := client.checkResourceCreationFeasibilityCreateRequest(ctx, resourceGroupName, deviceName, deviceCapacityRequestInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// checkResourceCreationFeasibilityCreateRequest creates the CheckResourceCreationFeasibility request.
func (client *DeviceCapacityCheckClient) checkResourceCreationFeasibilityCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, deviceCapacityRequestInfo DeviceCapacityRequestInfo, options *DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/deviceCapacityCheck"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	if options != nil && options.CapacityName != nil {
		reqQP.Set("capacityName", *options.CapacityName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, deviceCapacityRequestInfo)
}
