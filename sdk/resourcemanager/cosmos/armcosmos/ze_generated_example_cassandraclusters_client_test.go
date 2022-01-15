//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterListBySubscription.json
func ExampleCassandraClustersClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	res, err := client.ListBySubscription(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CassandraClustersClientListBySubscriptionResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterListByResourceGroup.json
func ExampleCassandraClustersClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	res, err := client.ListByResourceGroup(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CassandraClustersClientListByResourceGroupResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterGet.json
func ExampleCassandraClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CassandraClustersClientGetResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterDelete.json
func ExampleCassandraClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterCreate.json
func ExampleCassandraClustersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armcosmos.ClusterResource{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.ClusterResourceProperties{
				AuthenticationMethod: armcosmos.AuthenticationMethod("Cassandra").ToPtr(),
				CassandraVersion:     to.StringPtr("<cassandra-version>"),
				ClientCertificates: []*armcosmos.Certificate{
					{
						Pem: to.StringPtr("<pem>"),
					}},
				ClusterNameOverride:         to.StringPtr("<cluster-name-override>"),
				DelegatedManagementSubnetID: to.StringPtr("<delegated-management-subnet-id>"),
				ExternalGossipCertificates: []*armcosmos.Certificate{
					{
						Pem: to.StringPtr("<pem>"),
					}},
				ExternalSeedNodes: []*armcosmos.SeedNode{
					{
						IPAddress: to.StringPtr("<ipaddress>"),
					},
					{
						IPAddress: to.StringPtr("<ipaddress>"),
					},
					{
						IPAddress: to.StringPtr("<ipaddress>"),
					}},
				HoursBetweenBackups:           to.Int32Ptr(24),
				InitialCassandraAdminPassword: to.StringPtr("<initial-cassandra-admin-password>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CassandraClustersClientCreateUpdateResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterPatch.json
func ExampleCassandraClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armcosmos.ClusterResource{
			Tags: map[string]*string{
				"owner": to.StringPtr("mike"),
			},
			Properties: &armcosmos.ClusterResourceProperties{
				AuthenticationMethod: armcosmos.AuthenticationMethod("None").ToPtr(),
				ExternalGossipCertificates: []*armcosmos.Certificate{
					{
						Pem: to.StringPtr("<pem>"),
					}},
				ExternalSeedNodes: []*armcosmos.SeedNode{
					{
						IPAddress: to.StringPtr("<ipaddress>"),
					},
					{
						IPAddress: to.StringPtr("<ipaddress>"),
					},
					{
						IPAddress: to.StringPtr("<ipaddress>"),
					}},
				HoursBetweenBackups: to.Int32Ptr(12),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CassandraClustersClientUpdateResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraCommand.json
func ExampleCassandraClustersClient_BeginInvokeCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginInvokeCommand(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armcosmos.CommandPostBody{
			Command: to.StringPtr("<command>"),
			Host:    to.StringPtr("<host>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterDeallocate.json
func ExampleCassandraClustersClient_BeginDeallocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeallocate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterStart.json
func ExampleCassandraClustersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraStatus.json
func ExampleCassandraClustersClient_Status() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCassandraClustersClient("<subscription-id>", cred, nil)
	res, err := client.Status(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CassandraClustersClientStatusResult)
}