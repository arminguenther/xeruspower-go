// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package debuglog

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/logging/log"
)

func init() {
	object.Register(NewDebugLog)
}

type _DebugLog struct {
	idl.Object
}

// NewDebugLog returns a new DebugLog interface for the object at given RID.
func NewDebugLog(rid string, caller idl.Caller) DebugLog {
	return &_DebugLog{idl.NewObject(rid, caller)}
}

func (d *_DebugLog) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "logging.DebugLog",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (d *_DebugLog) Clear(ctx context.Context) error {
	_, err := d.Caller().Call(ctx, d.RID(), "clear", nil)
	return err
}

func (d *_DebugLog) GetInfo(ctx context.Context) (log.Info, error) {
	var ret log.Info
	val, err := d.Caller().Call(ctx, d.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DebugLog) GetChunk(ctx context.Context, in0 int32, in1 int32, in2 log.RangeDirection) (log.Chunk, error) {
	var ret log.Chunk
	val, err := d.Caller().Call(ctx, d.RID(), "getChunk", map[string]any{
		"refId":     in0,
		"count":     in1,
		"direction": in2,
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
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
