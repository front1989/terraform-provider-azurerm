package datafactory

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateEndPointConnectionsClient is the the Azure Data Factory V2 management API provides a RESTful set of web
// services that interact with Azure Data Factory V2 services.
type PrivateEndPointConnectionsClient struct {
	BaseClient
}

// NewPrivateEndPointConnectionsClient creates an instance of the PrivateEndPointConnectionsClient client.
func NewPrivateEndPointConnectionsClient(subscriptionID string) PrivateEndPointConnectionsClient {
	return NewPrivateEndPointConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndPointConnectionsClientWithBaseURI creates an instance of the PrivateEndPointConnectionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateEndPointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndPointConnectionsClient {
	return PrivateEndPointConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByFactory lists Private endpoint connections
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
func (client PrivateEndPointConnectionsClient) ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result PrivateEndpointConnectionListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndPointConnectionsClient.ListByFactory")
		defer func() {
			sc := -1
			if result.peclr.Response.Response != nil {
				sc = result.peclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PrivateEndPointConnectionsClient", "ListByFactory", err.Error())
	}

	result.fn = client.listByFactoryNextResults
	req, err := client.ListByFactoryPreparer(ctx, resourceGroupName, factoryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PrivateEndPointConnectionsClient", "ListByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.peclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PrivateEndPointConnectionsClient", "ListByFactory", resp, "Failure sending request")
		return
	}

	result.peclr, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PrivateEndPointConnectionsClient", "ListByFactory", resp, "Failure responding to request")
		return
	}
	if result.peclr.hasNextLink() && result.peclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByFactoryPreparer prepares the ListByFactory request.
func (client PrivateEndPointConnectionsClient) ListByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/privateEndPointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByFactorySender sends the ListByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndPointConnectionsClient) ListByFactorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByFactoryResponder handles the response to the ListByFactory request. The method always
// closes the http.Response Body.
func (client PrivateEndPointConnectionsClient) ListByFactoryResponder(resp *http.Response) (result PrivateEndpointConnectionListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFactoryNextResults retrieves the next set of results, if any.
func (client PrivateEndPointConnectionsClient) listByFactoryNextResults(ctx context.Context, lastResults PrivateEndpointConnectionListResponse) (result PrivateEndpointConnectionListResponse, err error) {
	req, err := lastResults.privateEndpointConnectionListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.PrivateEndPointConnectionsClient", "listByFactoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.PrivateEndPointConnectionsClient", "listByFactoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PrivateEndPointConnectionsClient", "listByFactoryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByFactoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndPointConnectionsClient) ListByFactoryComplete(ctx context.Context, resourceGroupName string, factoryName string) (result PrivateEndpointConnectionListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndPointConnectionsClient.ListByFactory")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByFactory(ctx, resourceGroupName, factoryName)
	return
}