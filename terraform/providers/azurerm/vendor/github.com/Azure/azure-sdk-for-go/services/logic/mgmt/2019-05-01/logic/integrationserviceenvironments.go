package logic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IntegrationServiceEnvironmentsClient is the REST API for Azure Logic Apps.
type IntegrationServiceEnvironmentsClient struct {
	BaseClient
}

// NewIntegrationServiceEnvironmentsClient creates an instance of the IntegrationServiceEnvironmentsClient client.
func NewIntegrationServiceEnvironmentsClient(subscriptionID string) IntegrationServiceEnvironmentsClient {
	return NewIntegrationServiceEnvironmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationServiceEnvironmentsClientWithBaseURI creates an instance of the IntegrationServiceEnvironmentsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewIntegrationServiceEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationServiceEnvironmentsClient {
	return IntegrationServiceEnvironmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration service environment.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
// integrationServiceEnvironment - the integration service environment.
func (client IntegrationServiceEnvironmentsClient) CreateOrUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment) (result IntegrationServiceEnvironmentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationServiceEnvironmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", pathParameters),
		autorest.WithJSON(integrationServiceEnvironment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) CreateOrUpdateSender(req *http.Request) (future IntegrationServiceEnvironmentsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationServiceEnvironment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration service environment.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
func (client IntegrationServiceEnvironmentsClient) Delete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroup, integrationServiceEnvironmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationServiceEnvironmentsClient) DeletePreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration service environment.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
func (client IntegrationServiceEnvironmentsClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result IntegrationServiceEnvironment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroup, integrationServiceEnvironmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationServiceEnvironmentsClient) GetPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) GetResponder(resp *http.Response) (result IntegrationServiceEnvironment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets a list of integration service environments by resource group.
// Parameters:
// resourceGroup - the resource group.
// top - the number of items to be included in the result.
func (client IntegrationServiceEnvironmentsClient) ListByResourceGroup(ctx context.Context, resourceGroup string, top *int32) (result IntegrationServiceEnvironmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.iselr.Response.Response != nil {
				sc = result.iselr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroup, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.iselr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.iselr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.iselr.hasNextLink() && result.iselr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client IntegrationServiceEnvironmentsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroup string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroup":  autorest.Encode("path", resourceGroup),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) ListByResourceGroupResponder(resp *http.Response) (result IntegrationServiceEnvironmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client IntegrationServiceEnvironmentsClient) listByResourceGroupNextResults(ctx context.Context, lastResults IntegrationServiceEnvironmentListResult) (result IntegrationServiceEnvironmentListResult, err error) {
	req, err := lastResults.integrationServiceEnvironmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationServiceEnvironmentsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroup string, top *int32) (result IntegrationServiceEnvironmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroup, top)
	return
}

// ListBySubscription gets a list of integration service environments by subscription.
// Parameters:
// top - the number of items to be included in the result.
func (client IntegrationServiceEnvironmentsClient) ListBySubscription(ctx context.Context, top *int32) (result IntegrationServiceEnvironmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.iselr.Response.Response != nil {
				sc = result.iselr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.iselr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.iselr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.iselr.hasNextLink() && result.iselr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client IntegrationServiceEnvironmentsClient) ListBySubscriptionPreparer(ctx context.Context, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationServiceEnvironments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) ListBySubscriptionResponder(resp *http.Response) (result IntegrationServiceEnvironmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client IntegrationServiceEnvironmentsClient) listBySubscriptionNextResults(ctx context.Context, lastResults IntegrationServiceEnvironmentListResult) (result IntegrationServiceEnvironmentListResult, err error) {
	req, err := lastResults.integrationServiceEnvironmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationServiceEnvironmentsClient) ListBySubscriptionComplete(ctx context.Context, top *int32) (result IntegrationServiceEnvironmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, top)
	return
}

// Restart restarts an integration service environment.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
func (client IntegrationServiceEnvironmentsClient) Restart(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.Restart")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RestartPreparer(ctx, resourceGroup, integrationServiceEnvironmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Restart", nil, "Failure preparing request")
		return
	}

	resp, err := client.RestartSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Restart", resp, "Failure sending request")
		return
	}

	result, err = client.RestartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Restart", resp, "Failure responding to request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client IntegrationServiceEnvironmentsClient) RestartPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) RestartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates an integration service environment.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
// integrationServiceEnvironment - the integration service environment.
func (client IntegrationServiceEnvironmentsClient) Update(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment) (result IntegrationServiceEnvironmentsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client IntegrationServiceEnvironmentsClient) UpdatePreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", pathParameters),
		autorest.WithJSON(integrationServiceEnvironment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentsClient) UpdateSender(req *http.Request) (future IntegrationServiceEnvironmentsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentsClient) UpdateResponder(resp *http.Response) (result IntegrationServiceEnvironment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}