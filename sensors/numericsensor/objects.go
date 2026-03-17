// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package numericsensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40000/sensors/sensor"
)

func init() {
	object.Register(NewNumericSensor)
}

type _NumericSensor struct {
	sensor.Sensor
}

// NewNumericSensor returns a new NumericSensor interface for the object at given RID.
func NewNumericSensor(rid string, caller idl.Caller) NumericSensor {
	return &_NumericSensor{sensor.NewSensor(rid, caller)}
}

func (n *_NumericSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor",
		Major: 4, Submajor: 0, Minor: 5,
	}
}

func (n *_NumericSensor) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := n.Caller().Call(ctx, n.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_NumericSensor) GetDefaultThresholds(ctx context.Context) (Thresholds, error) {
	var ret Thresholds
	val, err := n.Caller().Call(ctx, n.RID(), "getDefaultThresholds", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_NumericSensor) GetThresholds(ctx context.Context) (Thresholds, error) {
	var ret Thresholds
	val, err := n.Caller().Call(ctx, n.RID(), "getThresholds", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_NumericSensor) SetThresholds(ctx context.Context, in0 Thresholds) (int32, error) {
	var ret int32
	val, err := n.Caller().Call(ctx, n.RID(), "setThresholds", map[string]any{
		"thresh": in0.Encode(),
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

func (n *_NumericSensor) GetReading(ctx context.Context) (Reading, error) {
	var ret Reading
	val, err := n.Caller().Call(ctx, n.RID(), "getReading", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_NumericSensor) GetMinMax(ctx context.Context) (MinMax, error) {
	var ret MinMax
	val, err := n.Caller().Call(ctx, n.RID(), "getMinMax", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_NumericSensor) ResetMinMax(ctx context.Context) error {
	_, err := n.Caller().Call(ctx, n.RID(), "resetMinMax", nil)
	return err
}
