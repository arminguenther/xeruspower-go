// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package fitness

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
)

func init() {
	object.Register(NewFitness)
}

type _Fitness struct {
	idl.Object
}

// NewFitness returns a new Fitness interface for the object at given RID.
func NewFitness(rid string, caller idl.Caller) Fitness {
	return &_Fitness{idl.NewObject(rid, caller)}
}

func (f *_Fitness) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "fitness.Fitness",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_Fitness) GetDataEntries(ctx context.Context) ([]DataEntry, error) {
	var ret []DataEntry
	val, err := f.Caller().Call(ctx, f.RID(), "getDataEntries", nil)
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
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]DataEntry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 DataEntry
		err = e0.Decode(a0, f.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (f *_Fitness) GetErrorLogIndexRange(ctx context.Context) (int32, int32, error) {
	var out0 int32
	var out1 int32
	val, err := f.Caller().Call(ctx, f.RID(), "getErrorLogIndexRange", nil)
	if err != nil {
		return out0, out1, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("firstIndex", res)
	if err != nil {
		return out0, out1, err
	}
	out0, err = encoding.AsInt32(res["firstIndex"])
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("entryCount", res)
	if err != nil {
		return out0, out1, err
	}
	out1, err = encoding.AsInt32(res["entryCount"])
	if err != nil {
		return out0, out1, err
	}
	return out0, out1, nil
}

func (f *_Fitness) GetErrorLogEntries(ctx context.Context, in0 int32, in1 int32) ([]ErrorLogEntry, error) {
	var ret []ErrorLogEntry
	val, err := f.Caller().Call(ctx, f.RID(), "getErrorLogEntries", map[string]any{
		"startIndex": in0,
		"count":      in1,
	})
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
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]ErrorLogEntry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ErrorLogEntry
		err = e0.Decode(a0, f.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
