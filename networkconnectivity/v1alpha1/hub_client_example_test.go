// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package networkconnectivity_test

import (
	"context"

	networkconnectivity "cloud.google.com/go/networkconnectivity/v1alpha1"
	"google.golang.org/api/iterator"
	networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"
)

func ExampleNewHubClient() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleHubClient_ListHubs() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.ListHubsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListHubs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleHubClient_GetHub() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.GetHubRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_CreateHub() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.CreateHubRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_UpdateHub() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.UpdateHubRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_DeleteHub() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.DeleteHubRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleHubClient_ListSpokes() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.ListSpokesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSpokes(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleHubClient_GetSpoke() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.GetSpokeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_CreateSpoke() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.CreateSpokeRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_UpdateSpoke() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.UpdateSpokeRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_DeleteSpoke() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.DeleteSpokeRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
