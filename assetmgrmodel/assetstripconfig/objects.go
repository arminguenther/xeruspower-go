// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstripconfig

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
)

func init() {
	object.Register(NewAssetStripConfig)
}

type _AssetStripConfig struct {
	idl.Object
}

// NewAssetStripConfig returns a new AssetStripConfig interface for the object at given RID.
func NewAssetStripConfig(rid string, caller idl.Caller) AssetStripConfig {
	return &_AssetStripConfig{idl.NewObject(rid, caller)}
}

func (a *_AssetStripConfig) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStripConfig",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (a *_AssetStripConfig) GetStripSettings(ctx context.Context) (StripSettings, error) {
	var ret StripSettings
	val, err := a.Caller().Call(ctx, a.RID(), "getStripSettings", nil)
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

func (a *_AssetStripConfig) SetStripSettings(ctx context.Context, in0 StripSettings) (int32, error) {
	var ret int32
	val, err := a.Caller().Call(ctx, a.RID(), "setStripSettings", map[string]any{
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

func (a *_AssetStripConfig) GetRackUnitSettings(ctx context.Context, in0 int32) (int32, RackUnitSettings, error) {
	var ret int32
	var out0 RackUnitSettings
	val, err := a.Caller().Call(ctx, a.RID(), "getRackUnitSettings", map[string]any{
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
	err = encoding.In("settings", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["settings"], a.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (a *_AssetStripConfig) GetAllRackUnitSettings(ctx context.Context) ([]RackUnitSettings, error) {
	var ret []RackUnitSettings
	val, err := a.Caller().Call(ctx, a.RID(), "getAllRackUnitSettings", nil)
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
	ret = make([]RackUnitSettings, 0, len(s0))
	for _, a0 := range s0 {
		var e0 RackUnitSettings
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (a *_AssetStripConfig) SetRackUnitSettings(ctx context.Context, in0 int32, in1 RackUnitSettings) (int32, error) {
	var ret int32
	val, err := a.Caller().Call(ctx, a.RID(), "setRackUnitSettings", map[string]any{
		"rackUnitNumber": in0,
		"settings":       in1.Encode(),
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

func (a *_AssetStripConfig) SetMultipleRackUnitSettings(ctx context.Context, in0 map[int32]RackUnitSettings) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0.Encode()})
	}
	val, err := a.Caller().Call(ctx, a.RID(), "setMultipleRackUnitSettings", map[string]any{
		"settings": s0,
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
