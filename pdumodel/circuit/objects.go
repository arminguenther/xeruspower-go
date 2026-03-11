// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package circuit

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40412/pdumodel/pole"
)

func init() {
	object.Register(NewCircuit)
}

type _Circuit struct {
	idl.Object
}

// NewCircuit returns a new Circuit interface for the object at given RID.
func NewCircuit(rid string, caller idl.Caller) Circuit {
	return &_Circuit{idl.NewObject(rid, caller)}
}

func (c *_Circuit) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Circuit",
		Major: 2, Submajor: 0, Minor: 3,
	}
}

func (c *_Circuit) GetConfig(ctx context.Context) (Config, error) {
	var ret Config
	val, err := c.Caller().Call(ctx, c.RID(), "getConfig", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Circuit) GetSensors(ctx context.Context) (Sensors, error) {
	var ret Sensors
	val, err := c.Caller().Call(ctx, c.RID(), "getSensors", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Circuit) GetPoles(ctx context.Context) ([]pole.Pole, error) {
	var ret []pole.Pole
	val, err := c.Caller().Call(ctx, c.RID(), "getPoles", nil)
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
	ret = make([]pole.Pole, 0, len(s0))
	for _, a0 := range s0 {
		var e0 pole.Pole
		err = e0.Decode(a0, c.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (c *_Circuit) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := c.Caller().Call(ctx, c.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Circuit) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "setSettings", map[string]any{
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
