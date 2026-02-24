// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package powerbus

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/pdumodel/pole"
)

func init() {
	object.Register(NewPowerBus)
}

type _PowerBus struct {
	edevice.EDevice
}

// NewPowerBus returns a new PowerBus interface for the object at given RID.
func NewPowerBus(rid string, caller idl.Caller) PowerBus {
	return &_PowerBus{edevice.NewEDevice(rid, caller)}
}

func (p *_PowerBus) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerBus",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PowerBus) GetSensors(ctx context.Context) (Sensors, error) {
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

func (p *_PowerBus) GetPoles(ctx context.Context) ([]pole.Pole, error) {
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
