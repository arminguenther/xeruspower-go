// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkrequest

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
)

func init() {
	object.Register(NewBulkRequest)
}

type _BulkRequest struct {
	idl.Object
}

// NewBulkRequest returns a new BulkRequest interface for the object at given RID.
func NewBulkRequest(rid string, caller idl.Caller) BulkRequest {
	return &_BulkRequest{idl.NewObject(rid, caller)}
}

func (b *_BulkRequest) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "bulkrpc.BulkRequest",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (b *_BulkRequest) PerformRequest(ctx context.Context, in0 []Request) ([]Response, error) {
	var out0 []Response
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := b.Caller().Call(ctx, b.RID(), "performRequest", map[string]any{
		"requests": s0,
	})
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("responses", res)
	if err != nil {
		return out0, err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](res["responses"])
	if err != nil {
		return out0, err
	}
	out0 = make([]Response, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Response
		err = e1.Decode(a1, b.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e1)
	}
	return out0, nil
}

func (b *_BulkRequest) PerformBulk(ctx context.Context, in0 []BulkRequestRequest) ([]BulkRequestResponse, error) {
	var out0 []BulkRequestResponse
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := b.Caller().Call(ctx, b.RID(), "performBulk", map[string]any{
		"requests": s0,
	})
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("responses", res)
	if err != nil {
		return out0, err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](res["responses"])
	if err != nil {
		return out0, err
	}
	out0 = make([]BulkRequestResponse, 0, len(s1))
	for _, a1 := range s1 {
		var e1 BulkRequestResponse
		err = e1.Decode(a1, b.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e1)
	}
	return out0, nil
}

func (b *_BulkRequest) PerformBulkTimeout(ctx context.Context, in0 []BulkRequestRequest, in1 int32) ([]BulkRequestResponse, error) {
	var out0 []BulkRequestResponse
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := b.Caller().Call(ctx, b.RID(), "performBulkTimeout", map[string]any{
		"requests":  s0,
		"timeoutMs": in1,
	})
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("responses", res)
	if err != nil {
		return out0, err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](res["responses"])
	if err != nil {
		return out0, err
	}
	out0 = make([]BulkRequestResponse, 0, len(s1))
	for _, a1 := range s1 {
		var e1 BulkRequestResponse
		err = e1.Decode(a1, b.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e1)
	}
	return out0, nil
}
