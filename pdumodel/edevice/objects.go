// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package edevice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40032/pdumodel/pole"
	"github.com/arminguenther/xeruspower-go/v40032/pdumodel/waveform"
)

func init() {
	object.Register(NewEDevice)
}

type _EDevice struct {
	idl.Object
}

// NewEDevice returns a new EDevice interface for the object at given RID.
func NewEDevice(rid string, caller idl.Caller) EDevice {
	return &_EDevice{idl.NewObject(rid, caller)}
}

func (e *_EDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.EDevice",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (e *_EDevice) GetParents(ctx context.Context) ([]EDevice, error) {
	var ret []EDevice
	val, err := e.Caller().Call(ctx, e.RID(), "getParents", nil)
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
	ret = make([]EDevice, 0, len(s0))
	for _, a0 := range s0 {
		var e0 EDevice
		e0, err = object.As[EDevice](a0, e.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (e *_EDevice) GetChildren(ctx context.Context) ([]EDevice, error) {
	var ret []EDevice
	val, err := e.Caller().Call(ctx, e.RID(), "getChildren", nil)
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
	ret = make([]EDevice, 0, len(s0))
	for _, a0 := range s0 {
		var e0 EDevice
		e0, err = object.As[EDevice](a0, e.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (e *_EDevice) GetWaveform(ctx context.Context) (waveform.Waveform, error) {
	var ret waveform.Waveform
	val, err := e.Caller().Call(ctx, e.RID(), "getWaveform", nil)
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
	err = ret.Decode(res["_ret_"], e.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (e *_EDevice) GetPoleWaveform(ctx context.Context, in0 pole.PowerLine) (waveform.Waveform, error) {
	var ret waveform.Waveform
	val, err := e.Caller().Call(ctx, e.RID(), "getPoleWaveform", map[string]any{
		"line": in0,
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
	err = ret.Decode(res["_ret_"], e.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
