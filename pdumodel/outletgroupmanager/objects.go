// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outletgroupmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40510/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40510/pdumodel/outletgroup"
)

func init() {
	object.Register(NewOutletGroupManager)
}

type _OutletGroupManager struct {
	idl.Object
}

// NewOutletGroupManager returns a new OutletGroupManager interface for the object at given RID.
func NewOutletGroupManager(rid string, caller idl.Caller) OutletGroupManager {
	return &_OutletGroupManager{idl.NewObject(rid, caller)}
}

func (o *_OutletGroupManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroupManager",
		Major: 1, Submajor: 1, Minor: 11,
	}
}

func (o *_OutletGroupManager) CreateGroup(ctx context.Context, in0 string, in1 []outlet.Outlet) (int32, outletgroup.OutletGroup, error) {
	var ret int32
	var out0 outletgroup.OutletGroup
	s0 := make([]any, 0, len(in1))
	for _, e0 := range in1 {
		s0 = append(s0, object.ToMap(e0))
	}
	val, err := o.Caller().Call(ctx, o.RID(), "createGroup", map[string]any{
		"name":    in0,
		"members": s0,
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
	err = encoding.In("group", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = object.As[outletgroup.OutletGroup](res["group"], o.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (o *_OutletGroupManager) GetAllGroups(ctx context.Context) (map[int32]outletgroup.OutletGroup, error) {
	var ret map[int32]outletgroup.OutletGroup
	val, err := o.Caller().Call(ctx, o.RID(), "getAllGroups", nil)
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
	ret = make(map[int32]outletgroup.OutletGroup, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 outletgroup.OutletGroup
		v0, err = object.As[outletgroup.OutletGroup](b0, o.Caller())
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

func (o *_OutletGroupManager) GetGroup(ctx context.Context, in0 int32) (int32, outletgroup.OutletGroup, error) {
	var ret int32
	var out0 outletgroup.OutletGroup
	val, err := o.Caller().Call(ctx, o.RID(), "getGroup", map[string]any{
		"id": in0,
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
	err = encoding.In("group", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = object.As[outletgroup.OutletGroup](res["group"], o.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (o *_OutletGroupManager) DeleteGroup(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "deleteGroup", map[string]any{
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
