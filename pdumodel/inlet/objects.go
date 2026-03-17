// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package inlet

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40300/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/v40300/pdumodel/pole"
)

func init() {
	object.Register(NewInlet)
}

type _Inlet struct {
	edevice.EDevice
}

// NewInlet returns a new Inlet interface for the object at given RID.
func NewInlet(rid string, caller idl.Caller) Inlet {
	return &_Inlet{edevice.NewEDevice(rid, caller)}
}

func (i *_Inlet) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Inlet",
		Major: 3, Submajor: 0, Minor: 3,
	}
}

func (i *_Inlet) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := i.Caller().Call(ctx, i.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], i.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (i *_Inlet) GetSensors(ctx context.Context) (Sensors, error) {
	var ret Sensors
	val, err := i.Caller().Call(ctx, i.RID(), "getSensors", nil)
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
	err = ret.Decode(res["_ret_"], i.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (i *_Inlet) GetPoles(ctx context.Context) ([]pole.Pole, error) {
	var ret []pole.Pole
	val, err := i.Caller().Call(ctx, i.RID(), "getPoles", nil)
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
		err = e0.Decode(a0, i.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (i *_Inlet) GetLinePairs(ctx context.Context) ([]pole.MeteredLinePair, error) {
	var ret []pole.MeteredLinePair
	val, err := i.Caller().Call(ctx, i.RID(), "getLinePairs", nil)
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
	ret = make([]pole.MeteredLinePair, 0, len(s0))
	for _, a0 := range s0 {
		var e0 pole.MeteredLinePair
		err = e0.Decode(a0, i.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (i *_Inlet) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := i.Caller().Call(ctx, i.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], i.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (i *_Inlet) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := i.Caller().Call(ctx, i.RID(), "setSettings", map[string]any{
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

func (i *_Inlet) SetEnabled(ctx context.Context, in0 bool) error {
	_, err := i.Caller().Call(ctx, i.RID(), "setEnabled", map[string]any{
		"enabled": in0,
	})
	return err
}

func (i *_Inlet) IsEnabled(ctx context.Context) (bool, error) {
	var ret bool
	val, err := i.Caller().Call(ctx, i.RID(), "isEnabled", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
