// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package tacplusmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
)

func init() {
	object.Register(NewTacPlusManager)
}

type _TacPlusManager struct {
	idl.Object
}

// NewTacPlusManager returns a new TacPlusManager interface for the object at given RID.
func NewTacPlusManager(rid string, caller idl.Caller) TacPlusManager {
	return &_TacPlusManager{idl.NewObject(rid, caller)}
}

func (t *_TacPlusManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "auth.TacPlusManager",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (t *_TacPlusManager) GetTacPlusServers(ctx context.Context) ([]ServerSettings, error) {
	var ret []ServerSettings
	val, err := t.Caller().Call(ctx, t.RID(), "getTacPlusServers", nil)
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
	ret = make([]ServerSettings, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ServerSettings
		err = e0.Decode(a0, t.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (t *_TacPlusManager) SetTacPlusServers(ctx context.Context, in0 []ServerSettings) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := t.Caller().Call(ctx, t.RID(), "setTacPlusServers", map[string]any{
		"serverList": s0,
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

func (t *_TacPlusManager) TestTacPlusServer(ctx context.Context, in0 string, in1 string, in2 ServerSettings) (int32, error) {
	var ret int32
	val, err := t.Caller().Call(ctx, t.RID(), "testTacPlusServer", map[string]any{
		"username": in0,
		"password": in1,
		"settings": in2.Encode(),
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
