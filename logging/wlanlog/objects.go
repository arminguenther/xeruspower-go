// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package wlanlog

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40010/logging/log"
)

func init() {
	object.Register(NewWlanLog)
}

type _WlanLog struct {
	idl.Object
}

// NewWlanLog returns a new WlanLog interface for the object at given RID.
func NewWlanLog(rid string, caller idl.Caller) WlanLog {
	return &_WlanLog{idl.NewObject(rid, caller)}
}

func (w *_WlanLog) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "logging.WlanLog",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (w *_WlanLog) Clear(ctx context.Context) error {
	_, err := w.Caller().Call(ctx, w.RID(), "clear", nil)
	return err
}

func (w *_WlanLog) GetInfo(ctx context.Context) (log.Info, error) {
	var ret log.Info
	val, err := w.Caller().Call(ctx, w.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], w.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (w *_WlanLog) GetChunk(ctx context.Context, in0 int32, in1 int32, in2 log.RangeDirection) (log.Chunk, error) {
	var ret log.Chunk
	val, err := w.Caller().Call(ctx, w.RID(), "getChunk", map[string]any{
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
	err = ret.Decode(res["_ret_"], w.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
