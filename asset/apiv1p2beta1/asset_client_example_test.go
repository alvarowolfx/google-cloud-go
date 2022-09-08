// Copyright 2022 Google LLC
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

package asset_test

import (
	"context"

	asset "cloud.google.com/go/asset/apiv1p2beta1"
	assetpb "cloud.google.com/go/asset/apiv1p2beta1/assetpb"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRESTClient() {
	ctx := context.Background()
	c, err := asset.NewRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateFeed() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assetpb.CreateFeedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/asset/apiv1p2beta1/assetpb#CreateFeedRequest.
	}
	resp, err := c.CreateFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetFeed() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assetpb.GetFeedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/asset/apiv1p2beta1/assetpb#GetFeedRequest.
	}
	resp, err := c.GetFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListFeeds() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assetpb.ListFeedsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/asset/apiv1p2beta1/assetpb#ListFeedsRequest.
	}
	resp, err := c.ListFeeds(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateFeed() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assetpb.UpdateFeedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/asset/apiv1p2beta1/assetpb#UpdateFeedRequest.
	}
	resp, err := c.UpdateFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteFeed() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assetpb.DeleteFeedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/asset/apiv1p2beta1/assetpb#DeleteFeedRequest.
	}
	err = c.DeleteFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
