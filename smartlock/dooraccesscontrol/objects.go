// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dooraccesscontrol

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
)

func init() {
	object.Register(NewDoorAccessControl)
}

type _DoorAccessControl struct {
	idl.Object
}

// NewDoorAccessControl returns a new DoorAccessControl interface for the object at given RID.
func NewDoorAccessControl(rid string, caller idl.Caller) DoorAccessControl {
	return &_DoorAccessControl{idl.NewObject(rid, caller)}
}

func (d *_DoorAccessControl) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.DoorAccessControl",
		Major: 1, Submajor: 2, Minor: 4,
	}
}

func (d *_DoorAccessControl) GetDoorAccessRules(ctx context.Context) (map[int32]DoorAccessRule, error) {
	var ret map[int32]DoorAccessRule
	val, err := d.Caller().Call(ctx, d.RID(), "getDoorAccessRules", nil)
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
	ret = make(map[int32]DoorAccessRule, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 DoorAccessRule
		err = v0.Decode(b0, d.Caller())
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

func (d *_DoorAccessControl) SetAllDoorAccessRules(ctx context.Context, in0 map[int32]DoorAccessRule) (int32, []int32, error) {
	var ret int32
	var out0 []int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0.Encode()})
	}
	val, err := d.Caller().Call(ctx, d.RID(), "setAllDoorAccessRules", map[string]any{
		"rules": s0,
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
	err = encoding.In("invalidRuleIds", res)
	if err != nil {
		return ret, out0, err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](res["invalidRuleIds"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e1)
	}
	return ret, out0, nil
}

func (d *_DoorAccessControl) AddDoorAccessRule(ctx context.Context, in0 DoorAccessRule) (int32, int32, error) {
	var ret int32
	var out0 int32
	val, err := d.Caller().Call(ctx, d.RID(), "addDoorAccessRule", map[string]any{
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
	out0, err = encoding.AsInt32(res["ruleId"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (d *_DoorAccessControl) ModifyDoorAccessRule(ctx context.Context, in0 int32, in1 DoorAccessRule) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "modifyDoorAccessRule", map[string]any{
		"id":           in0,
		"modifiedRule": in1.Encode(),
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

func (d *_DoorAccessControl) DeleteDoorAccessRule(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "deleteDoorAccessRule", map[string]any{
		"id": in0,
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
