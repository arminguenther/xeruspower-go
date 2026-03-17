// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outlet

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/controller"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/inlet"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/overcurrentprotector"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/pole"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/waveform"
)

func init() {
	object.Register(NewOutlet)
}

type _Outlet struct {
	edevice.EDevice
}

// NewOutlet returns a new Outlet interface for the object at given RID.
func NewOutlet(rid string, caller idl.Caller) Outlet {
	return &_Outlet{edevice.NewEDevice(rid, caller)}
}

func (o *_Outlet) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Outlet",
		Major: 3, Submajor: 0, Minor: 3,
	}
}

func (o *_Outlet) GetMetaData(ctx context.Context) (MetaData, error) {
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

func (o *_Outlet) GetSensors(ctx context.Context) (Sensors, error) {
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

func (o *_Outlet) GetState(ctx context.Context) (State, error) {
	var ret State
	val, err := o.Caller().Call(ctx, o.RID(), "getState", nil)
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

func (o *_Outlet) SetPowerState(ctx context.Context, in0 PowerState) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "setPowerState", map[string]any{
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

func (o *_Outlet) CyclePowerState(ctx context.Context) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "cyclePowerState", nil)
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

func (o *_Outlet) GetSettings(ctx context.Context) (Settings, error) {
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

func (o *_Outlet) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
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

func (o *_Outlet) GetIOP(ctx context.Context) (inlet.Inlet, overcurrentprotector.OverCurrentProtector, []pole.Pole, error) {
	var out0 inlet.Inlet
	var out1 overcurrentprotector.OverCurrentProtector
	var out2 []pole.Pole
	val, err := o.Caller().Call(ctx, o.RID(), "getIOP", nil)
	if err != nil {
		return out0, out1, out2, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, out1, out2, err
	}
	err = encoding.In("i", res)
	if err != nil {
		return out0, out1, out2, err
	}
	out0, err = object.As[inlet.Inlet](res["i"], o.Caller())
	if err != nil {
		return out0, out1, out2, err
	}
	err = encoding.In("o", res)
	if err != nil {
		return out0, out1, out2, err
	}
	out1, err = object.As[overcurrentprotector.OverCurrentProtector](res["o"], o.Caller())
	if err != nil {
		return out0, out1, out2, err
	}
	err = encoding.In("p", res)
	if err != nil {
		return out0, out1, out2, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["p"])
	if err != nil {
		return out0, out1, out2, err
	}
	out2 = make([]pole.Pole, 0, len(s0))
	for _, a0 := range s0 {
		var e0 pole.Pole
		err = e0.Decode(a0, o.Caller())
		if err != nil {
			return out0, out1, out2, err
		}
		out2 = append(out2, e0)
	}
	return out0, out1, out2, nil
}

func (o *_Outlet) GetController(ctx context.Context) (controller.Controller, error) {
	var ret controller.Controller
	val, err := o.Caller().Call(ctx, o.RID(), "getController", nil)
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
	ret, err = object.As[controller.Controller](res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_Outlet) GetInrushWaveform(ctx context.Context) (waveform.Waveform, error) {
	var ret waveform.Waveform
	val, err := o.Caller().Call(ctx, o.RID(), "getInrushWaveform", nil)
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

func (o *_Outlet) SetServiceModeEnabled(ctx context.Context, in0 bool) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "setServiceModeEnabled", map[string]any{
		"enabled": in0,
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

func (o *_Outlet) Unstick(ctx context.Context) (int32, error) {
	var ret int32
	val, err := o.Caller().Call(ctx, o.RID(), "unstick", nil)
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
