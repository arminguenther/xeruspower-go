// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package ade

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
)

func init() {
	object.Register(NewAde)
}

type _Ade struct {
	idl.Object
}

// NewAde returns a new Ade interface for the object at given RID.
func NewAde(rid string, caller idl.Caller) Ade {
	return &_Ade{idl.NewObject(rid, caller)}
}

func (a *_Ade) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Ade",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_Ade) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := a.Caller().Call(ctx, a.RID(), "getMetaData", nil)
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

func (a *_Ade) GetLatestSample(ctx context.Context) ([]Sample, error) {
	var ret []Sample
	val, err := a.Caller().Call(ctx, a.RID(), "getLatestSample", nil)
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
	ret = make([]Sample, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Sample
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (a *_Ade) GetCalibrationData(ctx context.Context) (RegisterMap, error) {
	var ret RegisterMap
	val, err := a.Caller().Call(ctx, a.RID(), "getCalibrationData", nil)
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
	ret = make(map[string]int64, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 int64
		v0, err = encoding.AsInt64(b0)
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

func (a *_Ade) SetCalibrationData(ctx context.Context, in0 RegisterMap) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0})
	}
	val, err := a.Caller().Call(ctx, a.RID(), "setCalibrationData", map[string]any{
		"regs": s0,
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
