// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package luaservice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
)

func init() {
	object.Register(NewManager)
}

type _Manager struct {
	idl.Object
}

// NewManager returns a new Manager interface for the object at given RID.
func NewManager(rid string, caller idl.Caller) Manager {
	return &_Manager{idl.NewObject(rid, caller)}
}

func (m *_Manager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "luaservice.Manager",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (m *_Manager) SetScript(ctx context.Context, in0 string, in1 string, in2 ScriptOptions) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "setScript", map[string]any{
		"name":    in0,
		"script":  in1,
		"options": in2.Encode(),
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

func (m *_Manager) GetScript(ctx context.Context, in0 string) (int32, string, error) {
	var ret int32
	var out0 string
	val, err := m.Caller().Call(ctx, m.RID(), "getScript", map[string]any{
		"name": in0,
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
	err = encoding.In("script", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["script"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (m *_Manager) GetScriptNames(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := m.Caller().Call(ctx, m.RID(), "getScriptNames", nil)
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

func (m *_Manager) DeleteScript(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "deleteScript", map[string]any{
		"name": in0,
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

func (m *_Manager) SetScriptOptions(ctx context.Context, in0 string, in1 ScriptOptions) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "setScriptOptions", map[string]any{
		"name":    in0,
		"options": in1.Encode(),
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

func (m *_Manager) GetScriptOptions(ctx context.Context, in0 string) (int32, ScriptOptions, error) {
	var ret int32
	var out0 ScriptOptions
	val, err := m.Caller().Call(ctx, m.RID(), "getScriptOptions", map[string]any{
		"name": in0,
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
	err = encoding.In("options", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["options"], m.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (m *_Manager) GetEnvironment(ctx context.Context) (Environment, error) {
	var ret Environment
	val, err := m.Caller().Call(ctx, m.RID(), "getEnvironment", nil)
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
	err = ret.Decode(res["_ret_"], m.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (m *_Manager) GetScriptOutput(ctx context.Context, in0 string, in1 int64) (int32, int64, int64, string, bool, error) {
	var ret int32
	var out0 int64
	var out1 int64
	var out2 string
	var out3 bool
	val, err := m.Caller().Call(ctx, m.RID(), "getScriptOutput", map[string]any{
		"name":  in0,
		"iAddr": in1,
	})
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	err = encoding.In("oAddr", res)
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	out0, err = encoding.AsInt64(res["oAddr"])
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	err = encoding.In("nAddr", res)
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	out1, err = encoding.AsInt64(res["nAddr"])
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	err = encoding.In("oString", res)
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	out2, err = encoding.Is[string](res["oString"])
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	err = encoding.In("more", res)
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	out3, err = encoding.Is[bool](res["more"])
	if err != nil {
		return ret, out0, out1, out2, out3, err
	}
	return ret, out0, out1, out2, out3, nil
}

func (m *_Manager) ClearScriptOutput(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "clearScriptOutput", map[string]any{
		"name": in0,
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

func (m *_Manager) StartScript(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "startScript", map[string]any{
		"name": in0,
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

func (m *_Manager) StartScriptWithArgs(ctx context.Context, in0 string, in1 map[string]string) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in1))
	for k0, v0 := range in1 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0})
	}
	val, err := m.Caller().Call(ctx, m.RID(), "startScriptWithArgs", map[string]any{
		"name":      in0,
		"arguments": s0,
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

func (m *_Manager) TerminateScript(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "terminateScript", map[string]any{
		"name": in0,
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

func (m *_Manager) GetScriptState(ctx context.Context, in0 string) (int32, ScriptState, error) {
	var ret int32
	var out0 ScriptState
	val, err := m.Caller().Call(ctx, m.RID(), "getScriptState", map[string]any{
		"name": in0,
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
	err = encoding.In("state", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["state"], m.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (m *_Manager) GetScriptStates(ctx context.Context) (map[string]ScriptState, error) {
	var ret map[string]ScriptState
	val, err := m.Caller().Call(ctx, m.RID(), "getScriptStates", nil)
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
	i0, e0, l0 := encoding.AsMapItems(res["_ret_"])
	ret = make(map[string]ScriptState, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 ScriptState
		err = v0.Decode(b0, m.Caller())
		if err != nil {
			return ret, err
		}
		ret[k0] = v0
	}
	err = e0()
	if err != nil {
		return ret, err
	}
	return ret, nil
}
