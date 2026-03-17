// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powermeter

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/controller"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/pole"
)

func init() {
	object.Register(NewPowerMeter)
}

type _PowerMeter struct {
	idl.Object
}

// NewPowerMeter returns a new PowerMeter interface for the object at given RID.
func NewPowerMeter(rid string, caller idl.Caller) PowerMeter {
	return &_PowerMeter{idl.NewObject(rid, caller)}
}

func (p *_PowerMeter) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerMeter",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (p *_PowerMeter) GetConfig(ctx context.Context) (Config, error) {
	var ret Config
	val, err := p.Caller().Call(ctx, p.RID(), "getConfig", nil)
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

func (p *_PowerMeter) GetSensors(ctx context.Context) (Sensors, error) {
	var ret Sensors
	val, err := p.Caller().Call(ctx, p.RID(), "getSensors", nil)
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

func (p *_PowerMeter) GetPoles(ctx context.Context) ([]pole.Pole, error) {
	var ret []pole.Pole
	val, err := p.Caller().Call(ctx, p.RID(), "getPoles", nil)
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
		err = e0.Decode(a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_PowerMeter) GetControllers(ctx context.Context) ([]controller.Controller, error) {
	var ret []controller.Controller
	val, err := p.Caller().Call(ctx, p.RID(), "getControllers", nil)
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
	ret = make([]controller.Controller, 0, len(s0))
	for _, a0 := range s0 {
		var e0 controller.Controller
		e0, err = object.As[controller.Controller](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_PowerMeter) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := p.Caller().Call(ctx, p.RID(), "getSettings", nil)
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

func (p *_PowerMeter) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setSettings", map[string]any{
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

func (p *_PowerMeter) GetEnergyPulseSettings(ctx context.Context) (EnergyPulseSettings, error) {
	var ret EnergyPulseSettings
	val, err := p.Caller().Call(ctx, p.RID(), "getEnergyPulseSettings", nil)
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

func (p *_PowerMeter) SetEnergyPulseSettings(ctx context.Context, in0 EnergyPulseSettings) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setEnergyPulseSettings", map[string]any{
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
