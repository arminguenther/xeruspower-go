// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventlog

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40020/logging/log"
)

func init() {
	object.Register(NewEventLog)
}

type _EventLog struct {
	idl.Object
}

// NewEventLog returns a new EventLog interface for the object at given RID.
func NewEventLog(rid string, caller idl.Caller) EventLog {
	return &_EventLog{idl.NewObject(rid, caller)}
}

func (e *_EventLog) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "logging.EventLog",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (e *_EventLog) Clear(ctx context.Context) error {
	_, err := e.Caller().Call(ctx, e.RID(), "clear", nil)
	return err
}

func (e *_EventLog) GetInfo(ctx context.Context) (log.Info, error) {
	var ret log.Info
	val, err := e.Caller().Call(ctx, e.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], e.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (e *_EventLog) GetChunk(ctx context.Context, in0 int32, in1 int32, in2 log.RangeDirection, in3 []string) (log.Chunk, error) {
	var ret log.Chunk
	val, err := e.Caller().Call(ctx, e.RID(), "getChunk", map[string]any{
		"refId":      in0,
		"count":      in1,
		"direction":  in2,
		"categories": encoding.NonNilSlice(in3),
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
	err = ret.Decode(res["_ret_"], e.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
