// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package memorymapcontroller

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/pdumodel/controller"
)

func init() {
	object.Register(NewMemoryMapController)
}

type _MemoryMapController struct {
	controller.Controller
}

// NewMemoryMapController returns a new MemoryMapController interface for the object at given RID.
func NewMemoryMapController(rid string, caller idl.Caller) MemoryMapController {
	return &_MemoryMapController{controller.NewController(rid, caller)}
}

func (m *_MemoryMapController) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.MemoryMapController",
		Major: 5, Submajor: 0, Minor: 0,
	}
}

func (m *_MemoryMapController) ReadMemory(ctx context.Context, in0 int32, in1 int32) (int32, []byte, error) {
	var ret int32
	var out0 []byte
	val, err := m.Caller().Call(ctx, m.RID(), "readMemory", map[string]any{
		"address": in0,
		"size":    in1,
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
	err = encoding.In("memory", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["memory"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]byte, 0, len(s0))
	for _, a0 := range s0 {
		var e0 byte
		e0, err = encoding.AsByte(a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (m *_MemoryMapController) WriteMemory(ctx context.Context, in0 int32, in1 []byte) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in1))
	for _, e0 := range in1 {
		s0 = append(s0, e0)
	}
	val, err := m.Caller().Call(ctx, m.RID(), "writeMemory", map[string]any{
		"address": in0,
		"memory":  s0,
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
