// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package transferswitch

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/pole"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/waveform"
)

func init() {
	object.Register(NewTransferSwitch)
}

type _TransferSwitch struct {
	edevice.EDevice
}

// NewTransferSwitch returns a new TransferSwitch interface for the object at given RID.
func NewTransferSwitch(rid string, caller idl.Caller) TransferSwitch {
	return &_TransferSwitch{edevice.NewEDevice(rid, caller)}
}

func (t *_TransferSwitch) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.TransferSwitch",
		Major: 5, Submajor: 0, Minor: 2,
	}
}

func (t *_TransferSwitch) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := t.Caller().Call(ctx, t.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], t.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TransferSwitch) GetSensors(ctx context.Context) (Sensors, error) {
	var ret Sensors
	val, err := t.Caller().Call(ctx, t.RID(), "getSensors", nil)
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
	err = ret.Decode(res["_ret_"], t.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TransferSwitch) GetPoles(ctx context.Context) ([]pole.ThrowPole, error) {
	var ret []pole.ThrowPole
	val, err := t.Caller().Call(ctx, t.RID(), "getPoles", nil)
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
	ret = make([]pole.ThrowPole, 0, len(s0))
	for _, a0 := range s0 {
		var e0 pole.ThrowPole
		err = e0.Decode(a0, t.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (t *_TransferSwitch) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := t.Caller().Call(ctx, t.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], t.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TransferSwitch) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := t.Caller().Call(ctx, t.RID(), "setSettings", map[string]any{
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

func (t *_TransferSwitch) GetStatistics(ctx context.Context) (Statistics, error) {
	var ret Statistics
	val, err := t.Caller().Call(ctx, t.RID(), "getStatistics", nil)
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
	err = ret.Decode(res["_ret_"], t.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TransferSwitch) TransferToSource(ctx context.Context, in0 int32, in1 bool) (int32, error) {
	var ret int32
	val, err := t.Caller().Call(ctx, t.RID(), "transferToSource", map[string]any{
		"source":        in0,
		"faultOverride": in1,
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

func (t *_TransferSwitch) GetLastTransferReason(ctx context.Context) (TransferReason, error) {
	var ret TransferReason
	val, err := t.Caller().Call(ctx, t.RID(), "getLastTransferReason", nil)
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
	ret, err = encoding.AsEnum[TransferReason](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TransferSwitch) GetLastTransferWaveform(ctx context.Context) (waveform.Waveform, error) {
	var ret waveform.Waveform
	val, err := t.Caller().Call(ctx, t.RID(), "getLastTransferWaveform", nil)
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
	err = ret.Decode(res["_ret_"], t.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (t *_TransferSwitch) GetTransferLog(ctx context.Context) ([]TransferLogEntry, error) {
	var ret []TransferLogEntry
	val, err := t.Caller().Call(ctx, t.RID(), "getTransferLog", nil)
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
	ret = make([]TransferLogEntry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 TransferLogEntry
		err = e0.Decode(a0, t.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (t *_TransferSwitch) GetParameters(ctx context.Context) (Parameters, error) {
	var ret Parameters
	val, err := t.Caller().Call(ctx, t.RID(), "getParameters", nil)
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
	ret = make(map[string]int32, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 int32
		v0, err = encoding.AsInt32(b0)
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

func (t *_TransferSwitch) SetParameters(ctx context.Context, in0 Parameters) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0})
	}
	val, err := t.Caller().Call(ctx, t.RID(), "setParameters", map[string]any{
		"parameters": s0,
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
