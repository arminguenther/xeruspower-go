// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powermetercontroller

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/panel"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/powermeter"
)

func init() {
	object.Register(NewPowerMeterController)
}

type _PowerMeterController struct {
	idl.Object
}

// NewPowerMeterController returns a new PowerMeterController interface for the object at given RID.
func NewPowerMeterController(rid string, caller idl.Caller) PowerMeterController {
	return &_PowerMeterController{idl.NewObject(rid, caller)}
}

func (p *_PowerMeterController) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerMeterController",
		Major: 1, Submajor: 2, Minor: 9,
	}
}

func (p *_PowerMeterController) ScanMeterBoard(ctx context.Context, in0 int32) (int32, ScanResult, error) {
	var ret int32
	var out0 ScanResult
	val, err := p.Caller().Call(ctx, p.RID(), "scanMeterBoard", map[string]any{
		"powerMeterId": in0,
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
	err = encoding.In("result", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["result"], p.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (p *_PowerMeterController) GetPowerMeters(ctx context.Context) (map[int32]powermeter.PowerMeter, error) {
	var ret map[int32]powermeter.PowerMeter
	val, err := p.Caller().Call(ctx, p.RID(), "getPowerMeters", nil)
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
	ret = make(map[int32]powermeter.PowerMeter, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 powermeter.PowerMeter
		v0, err = object.As[powermeter.PowerMeter](b0, p.Caller())
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

func (p *_PowerMeterController) CreatePowerMeter(ctx context.Context, in0 powermeter.Config, in1 powermeter.Settings) (int32, powermeter.PowerMeter, error) {
	var ret int32
	var out0 powermeter.PowerMeter
	val, err := p.Caller().Call(ctx, p.RID(), "createPowerMeter", map[string]any{
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
	err = encoding.In("powerMeter", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = object.As[powermeter.PowerMeter](res["powerMeter"], p.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (p *_PowerMeterController) CreatePanel(ctx context.Context, in0 powermeter.Config, in1 powermeter.Settings, in2 panel.Settings) (int32, panel.Panel, error) {
	var ret int32
	var out0 panel.Panel
	val, err := p.Caller().Call(ctx, p.RID(), "createPanel", map[string]any{
		"config":             in0.Encode(),
		"powerMeterSettings": in1.Encode(),
		"panelSettings":      in2.Encode(),
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
	err = encoding.In("panel", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = object.As[panel.Panel](res["panel"], p.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (p *_PowerMeterController) DeletePowerMeter(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "deletePowerMeter", map[string]any{
		"powerMeterId": in0,
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
