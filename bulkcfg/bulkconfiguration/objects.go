// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkconfiguration

import (
	"context"
	"time"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
)

func init() {
	object.Register(NewBulkConfiguration)
}

type _BulkConfiguration struct {
	idl.Object
}

// NewBulkConfiguration returns a new BulkConfiguration interface for the object at given RID.
func NewBulkConfiguration(rid string, caller idl.Caller) BulkConfiguration {
	return &_BulkConfiguration{idl.NewObject(rid, caller)}
}

func (b *_BulkConfiguration) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "bulkcfg.BulkConfiguration",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (b *_BulkConfiguration) GetStatus(ctx context.Context) (Status, time.Time, error) {
	var out0 Status
	var out1 time.Time
	val, err := b.Caller().Call(ctx, b.RID(), "getStatus", nil)
	if err != nil {
		return out0, out1, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("status", res)
	if err != nil {
		return out0, out1, err
	}
	out0, err = encoding.AsEnum[Status](res["status"])
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("timeStamp", res)
	if err != nil {
		return out0, out1, err
	}
	out1, err = encoding.AsTime(res["timeStamp"])
	if err != nil {
		return out0, out1, err
	}
	return out0, out1, nil
}

func (b *_BulkConfiguration) GetFilters(ctx context.Context) ([]Filter, error) {
	var ret []Filter
	val, err := b.Caller().Call(ctx, b.RID(), "getFilters", nil)
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
	ret = make([]Filter, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Filter
		err = e0.Decode(a0, b.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (b *_BulkConfiguration) GetFilterProfiles(ctx context.Context) ([]FilterProfile, error) {
	var ret []FilterProfile
	val, err := b.Caller().Call(ctx, b.RID(), "getFilterProfiles", nil)
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
	ret = make([]FilterProfile, 0, len(s0))
	for _, a0 := range s0 {
		var e0 FilterProfile
		err = e0.Decode(a0, b.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (b *_BulkConfiguration) AddFilterProfile(ctx context.Context, in0 FilterProfile) (int32, error) {
	var ret int32
	val, err := b.Caller().Call(ctx, b.RID(), "addFilterProfile", map[string]any{
		"profile": in0.Encode(),
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

func (b *_BulkConfiguration) ModifyFilterProfile(ctx context.Context, in0 FilterProfile) (int32, error) {
	var ret int32
	val, err := b.Caller().Call(ctx, b.RID(), "modifyFilterProfile", map[string]any{
		"profile": in0.Encode(),
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

func (b *_BulkConfiguration) DeleteFilterProfile(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := b.Caller().Call(ctx, b.RID(), "deleteFilterProfile", map[string]any{
		"profileName": in0,
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

func (b *_BulkConfiguration) GetDefaultFilterProfileName(ctx context.Context) (string, error) {
	var ret string
	val, err := b.Caller().Call(ctx, b.RID(), "getDefaultFilterProfileName", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (b *_BulkConfiguration) SelectDefaultFilterProfile(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := b.Caller().Call(ctx, b.RID(), "selectDefaultFilterProfile", map[string]any{
		"profileName": in0,
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

func (b *_BulkConfiguration) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := b.Caller().Call(ctx, b.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], b.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (b *_BulkConfiguration) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := b.Caller().Call(ctx, b.RID(), "setSettings", map[string]any{
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
