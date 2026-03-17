// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package timereventmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
)

func init() {
	object.Register(NewTimerEventManager)
}

type _TimerEventManager struct {
	idl.Object
}

// NewTimerEventManager returns a new TimerEventManager interface for the object at given RID.
func NewTimerEventManager(rid string, caller idl.Caller) TimerEventManager {
	return &_TimerEventManager{idl.NewObject(rid, caller)}
}

func (t *_TimerEventManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.TimerEventManager",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (t *_TimerEventManager) AddTimerEvent(ctx context.Context, in0 Schedule) (int32, []string, error) {
	var ret int32
	var out0 []string
	val, err := t.Caller().Call(ctx, t.RID(), "addTimerEvent", map[string]any{
		"schedule": in0.Encode(),
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
	err = encoding.In("eventId", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["eventId"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (t *_TimerEventManager) ModifyTimerEvent(ctx context.Context, in0 []string, in1 Schedule) (int32, error) {
	var ret int32
	val, err := t.Caller().Call(ctx, t.RID(), "modifyTimerEvent", map[string]any{
		"eventId":  encoding.NonNilSlice(in0),
		"schedule": in1.Encode(),
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
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TimerEventManager) DeleteTimerEvent(ctx context.Context, in0 []string) (int32, error) {
	var ret int32
	val, err := t.Caller().Call(ctx, t.RID(), "deleteTimerEvent", map[string]any{
		"eventId": encoding.NonNilSlice(in0),
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
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TimerEventManager) ListTimerEvents(ctx context.Context) ([]TimerEvent, error) {
	var ret []TimerEvent
	val, err := t.Caller().Call(ctx, t.RID(), "listTimerEvents", nil)
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
	ret = make([]TimerEvent, 0, len(s0))
	for _, a0 := range s0 {
		var e0 TimerEvent
		err = e0.Decode(a0, t.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
