// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cfg

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
)

func init() {
	object.Register(NewCfg)
}

type _Cfg struct {
	idl.Object
}

// NewCfg returns a new Cfg interface for the object at given RID.
func NewCfg(rid string, caller idl.Caller) Cfg {
	return &_Cfg{idl.NewObject(rid, caller)}
}

func (c *_Cfg) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cfg.Cfg",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (c *_Cfg) GetValues(ctx context.Context, in0 []string) (int32, []string, error) {
	var ret int32
	var out0 []string
	val, err := c.Caller().Call(ctx, c.RID(), "getValues", map[string]any{
		"keys": encoding.NonNilSlice(in0),
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
	err = encoding.In("values", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["values"])
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

func (c *_Cfg) SetValues(ctx context.Context, in0 []KeyValue) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := c.Caller().Call(ctx, c.RID(), "setValues", map[string]any{
		"keyvaluepairs": s0,
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
