package network

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// SubnetsClient is the client for the Subnets methods of the Network service.
type SubnetsClient struct {
	ManagementClient
}

// NewSubnetsClient creates an instance of the SubnetsClient client.
func NewSubnetsClient(subscriptionID string) SubnetsClient {
	return NewSubnetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubnetsClientWithBaseURI creates an instance of the SubnetsClient client.
func NewSubnetsClientWithBaseURI(baseURI string, subscriptionID string) SubnetsClient {
	return SubnetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put Subnet operation creates/updates a subnet in
// thespecified virtual network
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. subnetName is the name of the subnet.
// subnetParameters is parameters supplied to the create/update Subnet
// operation
func (client SubnetsClient) CreateOrUpdate(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (result Subnet, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkName, subnetName, subnetParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubnetsClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subnetName":         url.QueryEscape(subnetName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets/{subnetName}"),
		autorest.WithJSON(subnetParameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SubnetsClient) CreateOrUpdateResponder(resp *http.Response) (result Subnet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete subnet operation deletes the specified subnet.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. subnetName is the name of the subnet.
func (client SubnetsClient) Delete(resourceGroupName string, virtualNetworkName string, subnetName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubnetsClient) DeletePreparer(resourceGroupName string, virtualNetworkName string, subnetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subnetName":         url.QueryEscape(subnetName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets/{subnetName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubnetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get subnet operation retreives information about the specified
// subnet.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. subnetName is the name of the subnet.
func (client SubnetsClient) Get(resourceGroupName string, virtualNetworkName string, subnetName string) (result Subnet, ae error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubnetsClient) GetPreparer(resourceGroupName string, virtualNetworkName string, subnetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subnetName":         url.QueryEscape(subnetName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets/{subnetName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubnetsClient) GetResponder(resp *http.Response) (result Subnet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List subnets opertion retrieves all the subnets in a virtual
// network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network.
func (client SubnetsClient) List(resourceGroupName string, virtualNetworkName string) (result SubnetListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, virtualNetworkName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SubnetsClient) ListPreparer(resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SubnetsClient) ListResponder(resp *http.Response) (result SubnetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client SubnetsClient) ListNextResults(lastResults SubnetListResult) (result SubnetListResult, ae error) {
	req, err := lastResults.SubnetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "List", "Failure responding to next results request request")
	}

	return
}
