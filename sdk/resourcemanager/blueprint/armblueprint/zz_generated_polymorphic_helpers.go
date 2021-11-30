//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

import "encoding/json"

func unmarshalArtifactClassification(rawMsg json.RawMessage) (ArtifactClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ArtifactClassification
	switch m["kind"] {
	case string(ArtifactKindPolicyAssignment):
		b = &PolicyAssignmentArtifact{}
	case string(ArtifactKindRoleAssignment):
		b = &RoleAssignmentArtifact{}
	case string(ArtifactKindTemplate):
		b = &TemplateArtifact{}
	default:
		b = &Artifact{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalArtifactClassificationArray(rawMsg json.RawMessage) ([]ArtifactClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ArtifactClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalArtifactClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalArtifactClassificationMap(rawMsg json.RawMessage) (map[string]ArtifactClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ArtifactClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalArtifactClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
