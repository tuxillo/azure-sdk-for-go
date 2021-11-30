//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ArtifactsListPager provides operations for iterating over paged responses.
type ArtifactsListPager struct {
	client    *ArtifactsClient
	current   ArtifactsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ArtifactsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ArtifactsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ArtifactsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArtifactList.NextLink == nil || len(*p.current.ArtifactList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ArtifactsListResponse page.
func (p *ArtifactsListPager) PageResponse() ArtifactsListResponse {
	return p.current
}

// AssignmentOperationsListPager provides operations for iterating over paged responses.
type AssignmentOperationsListPager struct {
	client    *AssignmentOperationsClient
	current   AssignmentOperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AssignmentOperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AssignmentOperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AssignmentOperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AssignmentOperationList.NextLink == nil || len(*p.current.AssignmentOperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AssignmentOperationsListResponse page.
func (p *AssignmentOperationsListPager) PageResponse() AssignmentOperationsListResponse {
	return p.current
}

// AssignmentsListPager provides operations for iterating over paged responses.
type AssignmentsListPager struct {
	client    *AssignmentsClient
	current   AssignmentsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AssignmentsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AssignmentsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AssignmentsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AssignmentList.NextLink == nil || len(*p.current.AssignmentList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AssignmentsListResponse page.
func (p *AssignmentsListPager) PageResponse() AssignmentsListResponse {
	return p.current
}

// BlueprintsListPager provides operations for iterating over paged responses.
type BlueprintsListPager struct {
	client    *BlueprintsClient
	current   BlueprintsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BlueprintsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BlueprintsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BlueprintsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.BlueprintList.NextLink == nil || len(*p.current.BlueprintList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BlueprintsListResponse page.
func (p *BlueprintsListPager) PageResponse() BlueprintsListResponse {
	return p.current
}

// PublishedArtifactsListPager provides operations for iterating over paged responses.
type PublishedArtifactsListPager struct {
	client    *PublishedArtifactsClient
	current   PublishedArtifactsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PublishedArtifactsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PublishedArtifactsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PublishedArtifactsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArtifactList.NextLink == nil || len(*p.current.ArtifactList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PublishedArtifactsListResponse page.
func (p *PublishedArtifactsListPager) PageResponse() PublishedArtifactsListResponse {
	return p.current
}

// PublishedBlueprintsListPager provides operations for iterating over paged responses.
type PublishedBlueprintsListPager struct {
	client    *PublishedBlueprintsClient
	current   PublishedBlueprintsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PublishedBlueprintsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PublishedBlueprintsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PublishedBlueprintsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PublishedBlueprintList.NextLink == nil || len(*p.current.PublishedBlueprintList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PublishedBlueprintsListResponse page.
func (p *PublishedBlueprintsListPager) PageResponse() PublishedBlueprintsListResponse {
	return p.current
}
