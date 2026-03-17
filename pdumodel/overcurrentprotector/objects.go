// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package overcurrentprotector

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/inlet"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/pole"
)

func init() {
	object.Register(NewOverCurrentProtector)
}

type _OverCurrentProtector struct {
	edevice.EDevice
}

// NewOverCurrentProtector returns a new OverCurrentProtector interface for the object at given RID.
func NewOverCurrentProtector(rid string, caller idl.Caller) OverCurrentProtector {
	return &_OverCurrentProtector{edevice.NewEDevice(rid, caller)}
}

func (o *_OverCurrentProtector) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OverCurrentProtector",
		Major: 4, Submajor: 0, Minor: 1,
	}
}

func (o *_OverCurrentProtector) GetMetaData(ctx context.Context) (MetaData, error) {
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

func (o *_OverCurrentProtector) GetSensors(ctx context.Context) (Sensors, error) {
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

func (o *_OverCurrentProtector) GetPoles(ctx context.Context) ([]pole.DoublePole, error) {
	var ret []pole.DoublePole
	val, err := o.Caller().Call(ctx, o.RID(), "getPoles", nil)
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
	ret = make([]pole.DoublePole, 0, len(s0))
	for _, a0 := range s0 {
		var e0 pole.DoublePole
		err = e0.Decode(a0, o.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (o *_OverCurrentProtector) GetInlet(ctx context.Context) (inlet.Inlet, error) {
	var ret inlet.Inlet
	val, err := o.Caller().Call(ctx, o.RID(), "getInlet", nil)
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
	ret, err = object.As[inlet.Inlet](res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OverCurrentProtector) GetOCP(ctx context.Context) (OverCurrentProtector, error) {
	var ret OverCurrentProtector
	val, err := o.Caller().Call(ctx, o.RID(), "getOCP", nil)
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
	ret, err = object.As[OverCurrentProtector](res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OverCurrentProtector) GetSettings(ctx context.Context) (Settings, error) {
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

func (o *_OverCurrentProtector) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
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
