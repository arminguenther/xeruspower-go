// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstriplogger

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
)

func init() {
	object.Register(NewAssetStripLogger)
}

type _AssetStripLogger struct {
	idl.Object
}

// NewAssetStripLogger returns a new AssetStripLogger interface for the object at given RID.
func NewAssetStripLogger(rid string, caller idl.Caller) AssetStripLogger {
	return &_AssetStripLogger{idl.NewObject(rid, caller)}
}

func (a *_AssetStripLogger) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStripLogger",
		Major: 1, Submajor: 0, Minor: 7,
	}
}

func (a *_AssetStripLogger) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := a.Caller().Call(ctx, a.RID(), "getInfo", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], a.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AssetStripLogger) GetRecords(ctx context.Context, in0 int32, in1 int32) (int32, []Record, error) {
	var ret int32
	var out0 []Record
	val, err := a.Caller().Call(ctx, a.RID(), "getRecords", map[string]any{
		"id":    in0,
		"count": in1,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("records", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["records"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]Record, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Record
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}
