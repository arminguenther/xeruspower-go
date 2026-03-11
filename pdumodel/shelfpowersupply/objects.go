// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package shelfpowersupply

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40510/pdumodel/controller"
	"github.com/arminguenther/xeruspower-go/v40510/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/v40510/pdumodel/pole"
)

func init() {
	object.Register(NewShelfPowerSupply)
}

type _ShelfPowerSupply struct {
	edevice.EDevice
}

// NewShelfPowerSupply returns a new ShelfPowerSupply interface for the object at given RID.
func NewShelfPowerSupply(rid string, caller idl.Caller) ShelfPowerSupply {
	return &_ShelfPowerSupply{edevice.NewEDevice(rid, caller)}
}

func (s *_ShelfPowerSupply) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.ShelfPowerSupply",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ShelfPowerSupply) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := s.Caller().Call(ctx, s.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) GetState(ctx context.Context) (State, error) {
	var ret State
	val, err := s.Caller().Call(ctx, s.RID(), "getState", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) GetInputSensors(ctx context.Context) (InputSensors, error) {
	var ret InputSensors
	val, err := s.Caller().Call(ctx, s.RID(), "getInputSensors", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) GetOutputSensors(ctx context.Context) (OutputSensors, error) {
	var ret OutputSensors
	val, err := s.Caller().Call(ctx, s.RID(), "getOutputSensors", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) GetPowerSupplySensors(ctx context.Context) (PowerSupplySensors, error) {
	var ret PowerSupplySensors
	val, err := s.Caller().Call(ctx, s.RID(), "getPowerSupplySensors", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) SetPowerState(ctx context.Context, in0 PowerState) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setPowerState", map[string]any{
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

func (s *_ShelfPowerSupply) SetFanSpeedOverride(ctx context.Context, in0 bool, in1 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setFanSpeedOverride", map[string]any{
		"enableOverride":  in0,
		"fanSpeedPercent": in1,
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

func (s *_ShelfPowerSupply) GetInputPoles(ctx context.Context) ([]pole.Pole, error) {
	var ret []pole.Pole
	val, err := s.Caller().Call(ctx, s.RID(), "getInputPoles", nil)
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
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) GetOutputPoles(ctx context.Context) ([]pole.Pole, error) {
	var ret []pole.Pole
	val, err := s.Caller().Call(ctx, s.RID(), "getOutputPoles", nil)
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
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_ShelfPowerSupply) GetController(ctx context.Context) (controller.Controller, error) {
	var ret controller.Controller
	val, err := s.Caller().Call(ctx, s.RID(), "getController", nil)
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
	ret, err = object.As[controller.Controller](res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
