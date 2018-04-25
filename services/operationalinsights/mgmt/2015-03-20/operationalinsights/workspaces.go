package operationalinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// WorkspacesClient is the operational Insights Client
type WorkspacesClient struct {
	BaseClient
}

// NewWorkspacesClient creates an instance of the WorkspacesClient client.
func NewWorkspacesClient(subscriptionID string, purgeID string) WorkspacesClient {
	return NewWorkspacesClientWithBaseURI(DefaultBaseURI, subscriptionID, purgeID)
}

// NewWorkspacesClientWithBaseURI creates an instance of the WorkspacesClient client.
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string, purgeID string) WorkspacesClient {
	return WorkspacesClient{NewWithBaseURI(baseURI, subscriptionID, purgeID)}
}

// GetSchema gets the schema for a given workspace.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - log Analytics workspace name
func (client WorkspacesClient) GetSchema(ctx context.Context, resourceGroupName string, workspaceName string) (result SearchGetSchemaResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspacesClient", "GetSchema", err.Error())
	}

	req, err := client.GetSchemaPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "GetSchema", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSchemaSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "GetSchema", resp, "Failure sending request")
		return
	}

	result, err = client.GetSchemaResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "GetSchema", resp, "Failure responding to request")
	}

	return
}

// GetSchemaPreparer prepares the GetSchema request.
func (client WorkspacesClient) GetSchemaPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-03-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/schema", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSchemaSender sends the GetSchema request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) GetSchemaSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetSchemaResponder handles the response to the GetSchema request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) GetSchemaResponder(resp *http.Response) (result SearchGetSchemaResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSearchResults submit a search for a given workspace. The response will contain an id to track the search. User
// can use the id to poll the search status and get the full search result later if the search takes long time to
// finish.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - log Analytics workspace name
// parameters - the parameters required to execute a search query.
func (client WorkspacesClient) GetSearchResults(ctx context.Context, resourceGroupName string, workspaceName string, parameters SearchParameters) (result WorkspacesGetSearchResultsFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Query", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspacesClient", "GetSearchResults", err.Error())
	}

	req, err := client.GetSearchResultsPreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "GetSearchResults", nil, "Failure preparing request")
		return
	}

	result, err = client.GetSearchResultsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "GetSearchResults", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetSearchResultsPreparer prepares the GetSearchResults request.
func (client WorkspacesClient) GetSearchResultsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters SearchParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-03-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/search", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSearchResultsSender sends the GetSearchResults request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) GetSearchResultsSender(req *http.Request) (future WorkspacesGetSearchResultsFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// GetSearchResultsResponder handles the response to the GetSearchResults request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) GetSearchResultsResponder(resp *http.Response) (result SearchResultsResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListLinkTargets get a list of workspaces which the current user has administrator privileges and are not associated
// with an Azure Subscription. The subscriptionId parameter in the Url is ignored.
func (client WorkspacesClient) ListLinkTargets(ctx context.Context) (result ListLinkTarget, err error) {
	req, err := client.ListLinkTargetsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "ListLinkTargets", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListLinkTargetsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "ListLinkTargets", resp, "Failure sending request")
		return
	}

	result, err = client.ListLinkTargetsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "ListLinkTargets", resp, "Failure responding to request")
	}

	return
}

// ListLinkTargetsPreparer prepares the ListLinkTargets request.
func (client WorkspacesClient) ListLinkTargetsPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/linkTargets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListLinkTargetsSender sends the ListLinkTargets request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) ListLinkTargetsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListLinkTargetsResponder handles the response to the ListLinkTargets request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) ListLinkTargetsResponder(resp *http.Response) (result ListLinkTarget, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Purge purges data in an Log Analytics workspace by a set of user-defined filters.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - log Analytics workspace name
// body - describes the body of a request to purge data in a single table of an Log Analytics Workspace
func (client WorkspacesClient) Purge(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody) (result WorkspacesPurgeFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Table", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.Filters", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspacesClient", "Purge", err.Error())
	}

	req, err := client.PurgePreparer(ctx, resourceGroupName, workspaceName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "Purge", nil, "Failure preparing request")
		return
	}

	result, err = client.PurgeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "Purge", result.Response(), "Failure sending request")
		return
	}

	return
}

// PurgePreparer prepares the Purge request.
func (client WorkspacesClient) PurgePreparer(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-03-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/purge", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) PurgeSender(req *http.Request) (future WorkspacesPurgeFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) PurgeResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateSearchResults gets updated search results for a given search query.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - log Analytics workspace name
// ID - the id of the search that will have results updated. You can get the id from the response of the
// GetResults call.
func (client WorkspacesClient) UpdateSearchResults(ctx context.Context, resourceGroupName string, workspaceName string, ID string) (result SearchResultsResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspacesClient", "UpdateSearchResults", err.Error())
	}

	req, err := client.UpdateSearchResultsPreparer(ctx, resourceGroupName, workspaceName, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "UpdateSearchResults", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSearchResultsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "UpdateSearchResults", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateSearchResultsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesClient", "UpdateSearchResults", resp, "Failure responding to request")
	}

	return
}

// UpdateSearchResultsPreparer prepares the UpdateSearchResults request.
func (client WorkspacesClient) UpdateSearchResultsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-03-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/search/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSearchResultsSender sends the UpdateSearchResults request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) UpdateSearchResultsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateSearchResultsResponder handles the response to the UpdateSearchResults request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) UpdateSearchResultsResponder(resp *http.Response) (result SearchResultsResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
