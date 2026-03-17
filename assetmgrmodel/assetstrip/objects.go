// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstrip

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
)

func init() {
	object.Register(NewAssetStrip)
}

type _AssetStrip struct {
	idl.Object
}

// NewAssetStrip returns a new AssetStrip interface for the object at given RID.
func NewAssetStrip(rid string, caller idl.Caller) AssetStrip {
	return &_AssetStrip{idl.NewObject(rid, caller)}
}

func (a *_AssetStrip) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip",
		Major: 2, Submajor: 0, Minor: 6,
	}
}

func (a *_AssetStrip) GetState(ctx context.Context) (State, error) {
	var ret State
	val, err := a.Caller().Call(ctx, a.RID(), "getState", nil)
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
	ret, err = encoding.AsEnum[State](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AssetStrip) GetDeviceInfo(ctx context.Context) (DeviceInfo, error) {
	var ret DeviceInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getDeviceInfo", nil)
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
	err = ret.Decode(res["_ret_"], a.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AssetStrip) GetStripInfo(ctx context.Context) (StripInfo, error) {
	var ret StripInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getStripInfo", nil)
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
	err = ret.Decode(res["_ret_"], a.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AssetStrip) GetRackUnitInfo(ctx context.Context, in0 int32) (int32, RackUnitInfo, error) {
	var ret int32
	var out0 RackUnitInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getRackUnitInfo", map[string]any{
		"rackUnitNumber": in0,
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
	err = encoding.In("info", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["info"], a.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (a *_AssetStrip) GetAllRackUnitInfos(ctx context.Context) ([]RackUnitInfo, error) {
	var ret []RackUnitInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getAllRackUnitInfos", nil)
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
	ret = make([]RackUnitInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 RackUnitInfo
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (a *_AssetStrip) GetTag(ctx context.Context, in0 int32, in1 int32) (int32, TagInfo, error) {
	var ret int32
	var out0 TagInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getTag", map[string]any{
		"rackUnitNumber": in0,
		"slotNumber":     in1,
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
	err = encoding.In("tagInfo", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["tagInfo"], a.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (a *_AssetStrip) GetAllTags(ctx context.Context) ([]TagInfo, error) {
	var ret []TagInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getAllTags", nil)
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
	ret = make([]TagInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 TagInfo
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (a *_AssetStrip) GetMainTags(ctx context.Context) ([]TagInfo, error) {
	var ret []TagInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getMainTags", nil)
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
	ret = make([]TagInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 TagInfo
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (a *_AssetStrip) GetExtensionTags(ctx context.Context, in0 int32) (int32, []TagInfo, error) {
	var ret int32
	var out0 []TagInfo
	val, err := a.Caller().Call(ctx, a.RID(), "getExtensionTags", map[string]any{
		"rackUnitNumber": in0,
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
	err = encoding.In("tags", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["tags"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]TagInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 TagInfo
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (a *_AssetStrip) TriggerPowercycle(ctx context.Context, in0 bool) error {
	_, err := a.Caller().Call(ctx, a.RID(), "triggerPowercycle", map[string]any{
		"hard": in0,
	})
	return err
}

func (a *_AssetStrip) ProgramTagIDs(ctx context.Context, in0 []TagInfo) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := a.Caller().Call(ctx, a.RID(), "programTagIDs", map[string]any{
		"tagInfos": s0,
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

func (a *_AssetStrip) GetFirmwareUpdateState(ctx context.Context) (FirmwareUpdateState, error) {
	var ret FirmwareUpdateState
	val, err := a.Caller().Call(ctx, a.RID(), "getFirmwareUpdateState", nil)
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
	ret, err = encoding.AsEnum[FirmwareUpdateState](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
