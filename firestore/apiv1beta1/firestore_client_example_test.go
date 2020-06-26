// Copyright 2020 Google LLC
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

package firestore_test

import (
	"context"
	"io"

	firestore "github.com/adrianduke/google-cloud-go/apiv1beta1"
	"google.golang.org/api/iterator"
	firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_GetDocument() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.GetDocumentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListDocuments() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.ListDocumentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDocuments(ctx, req)
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

func ExampleClient_CreateDocument() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.CreateDocumentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateDocument() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.UpdateDocumentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteDocument() {
	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.DeleteDocumentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_BeginTransaction() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.BeginTransactionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.BeginTransaction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_Commit() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.CommitRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Commit(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_Rollback() {
	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.RollbackRequest{
		// TODO: Fill request struct fields.
	}
	err = c.Rollback(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_Write() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	stream, err := c.Write(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	go func() {
		reqs := []*firestorepb.WriteRequest{
			// TODO: Create requests.
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				// TODO: Handle error.
			}
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_Listen() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	stream, err := c.Listen(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	go func() {
		reqs := []*firestorepb.ListenRequest{
			// TODO: Create requests.
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				// TODO: Handle error.
			}
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListCollectionIds() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.ListCollectionIdsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCollectionIds(ctx, req)
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
