//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ComputeClient contains the methods for the Compute group.
// Don't use this type directly, use NewComputeClient() instead.
type ComputeClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewComputeClient creates a new instance of ComputeClient with the specified values.
func NewComputeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ComputeClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ComputeClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent
// is to create a new compute, do a GET first to verify that it does not
// exist yet.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeBeginCreateOrUpdateOptions) (ComputeCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return ComputeCreateOrUpdatePollerResponse{}, err
	}
	result := ComputeCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ComputeClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ComputeCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ComputeCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to
// create a new compute, do a GET first to verify that it does not
// exist yet.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ComputeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ComputeClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes specified Machine Learning compute.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeBeginDeleteOptions) (ComputeDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction, options)
	if err != nil {
		return ComputeDeletePollerResponse{}, err
	}
	result := ComputeDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ComputeClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ComputeDeletePollerResponse{}, err
	}
	result.Poller = &ComputeDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes specified Machine Learning compute.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ComputeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	reqQP.Set("underlyingResourceAction", string(underlyingResourceAction))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ComputeClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets compute definition by its name. Any secrets (storage keys, service credentials, etc) are not returned - use 'keys' nested resource to get
// them.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeGetOptions) (ComputeGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComputeGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComputeGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ComputeClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComputeClient) getHandleResponse(resp *http.Response) (ComputeGetResponse, error) {
	result := ComputeGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputeResource); err != nil {
		return ComputeGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ComputeClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets computes in specified workspace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) List(resourceGroupName string, workspaceName string, options *ComputeListOptions) *ComputeListPager {
	return &ComputeListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp ComputeListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PaginatedComputeResourcesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ComputeClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ComputeListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ComputeClient) listHandleResponse(resp *http.Response) (ComputeListResponse, error) {
	result := ComputeListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaginatedComputeResourcesList); err != nil {
		return ComputeListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ComputeClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListKeys - Gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeListKeysOptions) (ComputeListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComputeListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComputeListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ComputeClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ComputeClient) listKeysHandleResponse(resp *http.Response) (ComputeListKeysResponse, error) {
	result := ComputeListKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ComputeListKeysResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *ComputeClient) listKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListNodes - Get the details (e.g IP address, port etc) of all the compute nodes in the compute.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) ListNodes(resourceGroupName string, workspaceName string, computeName string, options *ComputeListNodesOptions) *ComputeListNodesPager {
	return &ComputeListNodesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listNodesCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
		},
		advancer: func(ctx context.Context, resp ComputeListNodesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AmlComputeNodesInformation.NextLink)
		},
	}
}

// listNodesCreateRequest creates the ListNodes request.
func (client *ComputeClient) listNodesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeListNodesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listNodes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listNodesHandleResponse handles the ListNodes response.
func (client *ComputeClient) listNodesHandleResponse(resp *http.Response) (ComputeListNodesResponse, error) {
	result := ComputeListNodesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AmlComputeNodesInformation); err != nil {
		return ComputeListNodesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listNodesHandleError handles the ListNodes error response.
func (client *ComputeClient) listNodesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRestart - Posts a restart action to a compute instance
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) BeginRestart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginRestartOptions) (ComputeRestartPollerResponse, error) {
	resp, err := client.restart(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeRestartPollerResponse{}, err
	}
	result := ComputeRestartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ComputeClient.Restart", "", resp, client.pl, client.restartHandleError)
	if err != nil {
		return ComputeRestartPollerResponse{}, err
	}
	result.Poller = &ComputeRestartPoller{
		pt: pt,
	}
	return result, nil
}

// Restart - Posts a restart action to a compute instance
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) restart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.restartHandleError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ComputeClient) restartCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// restartHandleError handles the Restart error response.
func (client *ComputeClient) restartHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginStart - Posts a start action to a compute instance
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) BeginStart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginStartOptions) (ComputeStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeStartPollerResponse{}, err
	}
	result := ComputeStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ComputeClient.Start", "", resp, client.pl, client.startHandleError)
	if err != nil {
		return ComputeStartPollerResponse{}, err
	}
	result.Poller = &ComputeStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Posts a start action to a compute instance
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) start(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.startHandleError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ComputeClient) startCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startHandleError handles the Start error response.
func (client *ComputeClient) startHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginStop - Posts a stop action to a compute instance
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) BeginStop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginStopOptions) (ComputeStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeStopPollerResponse{}, err
	}
	result := ComputeStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ComputeClient.Stop", "", resp, client.pl, client.stopHandleError)
	if err != nil {
		return ComputeStopPollerResponse{}, err
	}
	result.Poller = &ComputeStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Posts a stop action to a compute instance
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) stop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ComputeClient) stopCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopHandleError handles the Stop error response.
func (client *ComputeClient) stopHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeBeginUpdateOptions) (ComputeUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return ComputeUpdatePollerResponse{}, err
	}
	result := ComputeUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ComputeClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ComputeUpdatePollerResponse{}, err
	}
	result.Poller = &ComputeUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ComputeClient) update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ComputeClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *ComputeClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}