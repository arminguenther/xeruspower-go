// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gatewaysensormanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/valobj"
)

func init() {
	object.Register(NewGatewaySensorManager)
}

type _GatewaySensorManager struct {
	idl.Object
}

// NewGatewaySensorManager returns a new GatewaySensorManager interface for the object at given RID.
func NewGatewaySensorManager(rid string, caller idl.Caller) GatewaySensorManager {
	return &_GatewaySensorManager{idl.NewObject(rid, caller)}
}

func (g *_GatewaySensorManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (g *_GatewaySensorManager) GetConfiguration(ctx context.Context) (map[string]ConfigurationPackage, error) {
	var ret map[string]ConfigurationPackage
	val, err := g.Caller().Call(ctx, g.RID(), "getConfiguration", nil)
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
	ret = make(map[string]ConfigurationPackage, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 ConfigurationPackage
		err = v0.Decode(b0, g.Caller())
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

func (g *_GatewaySensorManager) SetConfiguration(ctx context.Context, in0 map[string]ConfigurationPackage) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0.Encode()})
	}
	val, err := g.Caller().Call(ctx, g.RID(), "setConfiguration", map[string]any{
		"cfg": s0,
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

func (g *_GatewaySensorManager) GetFeedback(ctx context.Context) ([]Feedback, error) {
	var ret []Feedback
	val, err := g.Caller().Call(ctx, g.RID(), "getFeedback", nil)
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
	ret = make([]Feedback, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Feedback
		e0, err = valobj.As[Feedback](a0, g.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
