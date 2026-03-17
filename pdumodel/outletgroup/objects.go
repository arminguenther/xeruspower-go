// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outletgroup

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/pdumodel/outlet"
)

func init() {
	object.Register(NewOutletGroup)
}

type _OutletGroup struct {
	idl.Object
}

// NewOutletGroup returns a new OutletGroup interface for the object at given RID.
func NewOutletGroup(rid string, caller idl.Caller) OutletGroup {
	return &_OutletGroup{idl.NewObject(rid, caller)}
}

func (o *_OutletGroup) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroup",
		Major: 1, Submajor: 1, Minor: 9,
	}
}

func (o *_OutletGroup) GetSensors(ctx context.Context) (Sensors, error) {
	var ret Sensors
	val, err := o.Caller().Call(ctx, o.RID(), "getSensors", nil)
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
	err = ret.Decode(res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OutletGroup) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := o.Caller().Call(ctx, o.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OutletGroup) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := o.Caller().Call(ctx, o.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OutletGroup) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
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

func (o *_OutletGroup) SetAllOutletPowerStates(ctx context.Context, in0 outlet.PowerState) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "setAllOutletPowerStates", map[string]any{
		"pstate": in0,
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

func (o *_OutletGroup) CycleAllOutletPowerStates(ctx context.Context) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "cycleAllOutletPowerStates", nil)
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
