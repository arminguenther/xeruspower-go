// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventengine

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
)

func init() {
	object.Register(NewEngine)
}

type _Engine struct {
	idl.Object
}

// NewEngine returns a new Engine interface for the object at given RID.
func NewEngine(rid string, caller idl.Caller) Engine {
	return &_Engine{idl.NewObject(rid, caller)}
}

func (e *_Engine) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.Engine",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (e *_Engine) ListEventDescs(ctx context.Context, in0 []string) (int32, []EngineEventDesc, error) {
	var ret int32
	var out0 []EngineEventDesc
	val, err := e.Caller().Call(ctx, e.RID(), "listEventDescs", map[string]any{
		"eventIdPrefix": encoding.NonNilSlice(in0),
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
	err = encoding.In("eventDescs", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["eventDescs"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]EngineEventDesc, 0, len(s0))
	for _, a0 := range s0 {
		var e0 EngineEventDesc
		err = e0.Decode(a0, e.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (e *_Engine) ListActionTypes(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := e.Caller().Call(ctx, e.RID(), "listActionTypes", nil)
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
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (e *_Engine) AddAction(ctx context.Context, in0 EngineAction) (int32, string, error) {
	var ret int32
	var out0 string
	val, err := e.Caller().Call(ctx, e.RID(), "addAction", map[string]any{
		"action": in0.Encode(),
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
	err = encoding.In("actionId", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["actionId"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (e *_Engine) ModifyAction(ctx context.Context, in0 EngineAction) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "modifyAction", map[string]any{
		"action": in0.Encode(),
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

func (e *_Engine) DeleteAction(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "deleteAction", map[string]any{
		"actionId": in0,
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

func (e *_Engine) ListActions(ctx context.Context) ([]EngineAction, error) {
	var ret []EngineAction
	val, err := e.Caller().Call(ctx, e.RID(), "listActions", nil)
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
	ret = make([]EngineAction, 0, len(s0))
	for _, a0 := range s0 {
		var e0 EngineAction
		err = e0.Decode(a0, e.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (e *_Engine) TriggerAction(ctx context.Context, in0 string, in1 []KeyValue) (int32, string, error) {
	var ret int32
	var out0 string
	s0 := make([]any, 0, len(in1))
	for _, e0 := range in1 {
		s0 = append(s0, e0.Encode())
	}
	val, err := e.Caller().Call(ctx, e.RID(), "triggerAction", map[string]any{
		"actionId": in0,
		"context":  s0,
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
	err = encoding.In("errMsg", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["errMsg"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (e *_Engine) TestAction(ctx context.Context, in0 EngineAction, in1 []KeyValue) (int32, string, error) {
	var ret int32
	var out0 string
	s0 := make([]any, 0, len(in1))
	for _, e0 := range in1 {
		s0 = append(s0, e0.Encode())
	}
	val, err := e.Caller().Call(ctx, e.RID(), "testAction", map[string]any{
		"action":  in0.Encode(),
		"context": s0,
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
	err = encoding.In("errMsg", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["errMsg"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (e *_Engine) AddRule(ctx context.Context, in0 EngineRule) (int32, string, error) {
	var ret int32
	var out0 string
	val, err := e.Caller().Call(ctx, e.RID(), "addRule", map[string]any{
		"rule": in0.Encode(),
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
	err = encoding.In("ruleId", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["ruleId"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (e *_Engine) ModifyRule(ctx context.Context, in0 EngineRule) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "modifyRule", map[string]any{
		"rule": in0.Encode(),
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

func (e *_Engine) EnableRule(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "enableRule", map[string]any{
		"ruleId": in0,
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

func (e *_Engine) DisableRule(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "disableRule", map[string]any{
		"ruleId": in0,
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

func (e *_Engine) DeleteRule(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "deleteRule", map[string]any{
		"ruleId": in0,
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

func (e *_Engine) ListRules(ctx context.Context) ([]EngineRule, error) {
	var ret []EngineRule
	val, err := e.Caller().Call(ctx, e.RID(), "listRules", nil)
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
	ret = make([]EngineRule, 0, len(s0))
	for _, a0 := range s0 {
		var e0 EngineRule
		err = e0.Decode(a0, e.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (e *_Engine) DeliverEvent(ctx context.Context, in0 Event) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "deliverEvent", map[string]any{
		"event": in0.Encode(),
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

func (e *_Engine) RearmRule(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := e.Caller().Call(ctx, e.RID(), "rearmRule", map[string]any{
		"ruleId": in0,
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
