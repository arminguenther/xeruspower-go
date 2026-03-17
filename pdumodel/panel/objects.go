// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package panel

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/circuit"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/powermeter"
)

func init() {
	object.Register(NewPanel)
}

type _Panel struct {
	powermeter.PowerMeter
}

// NewPanel returns a new Panel interface for the object at given RID.
func NewPanel(rid string, caller idl.Caller) Panel {
	return &_Panel{powermeter.NewPowerMeter(rid, caller)}
}

func (p *_Panel) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Panel",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (p *_Panel) GetPanelSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := p.Caller().Call(ctx, p.RID(), "getPanelSettings", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Panel) SetPanelSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setPanelSettings", map[string]any{
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

func (p *_Panel) GetCircuits(ctx context.Context) (map[int32]circuit.Circuit, error) {
	var ret map[int32]circuit.Circuit
	val, err := p.Caller().Call(ctx, p.RID(), "getCircuits", nil)
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
	ret = make(map[int32]circuit.Circuit, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 circuit.Circuit
		v0, err = object.As[circuit.Circuit](b0, p.Caller())
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

func (p *_Panel) CreateCircuit(ctx context.Context, in0 circuit.Config, in1 circuit.Settings) (int32, circuit.Circuit, error) {
	var ret int32
	var out0 circuit.Circuit
	val, err := p.Caller().Call(ctx, p.RID(), "createCircuit", map[string]any{
		"config":   in0.Encode(),
		"settings": in1.Encode(),
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
	err = encoding.In("circuit", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = object.As[circuit.Circuit](res["circuit"], p.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (p *_Panel) DeleteCircuit(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "deleteCircuit", map[string]any{
		"position": in0,
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
