// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package voltagemonitoringsensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/sensors/numericsensor"
)

func init() {
	object.Register(NewVoltageMonitoringSensor)
}

type _VoltageMonitoringSensor struct {
	numericsensor.NumericSensor
}

// NewVoltageMonitoringSensor returns a new VoltageMonitoringSensor interface for the object at given RID.
func NewVoltageMonitoringSensor(rid string, caller idl.Caller) VoltageMonitoringSensor {
	return &_VoltageMonitoringSensor{numericsensor.NewNumericSensor(rid, caller)}
}

func (v *_VoltageMonitoringSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.VoltageMonitoringSensor",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (v *_VoltageMonitoringSensor) GetRecentEvents(ctx context.Context) ([]Event, error) {
	var ret []Event
	val, err := v.Caller().Call(ctx, v.RID(), "getRecentEvents", nil)
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
	ret = make([]Event, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Event
		err = e0.Decode(a0, v.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (v *_VoltageMonitoringSensor) ClearRecentEvents(ctx context.Context) error {
	_, err := v.Caller().Call(ctx, v.RID(), "clearRecentEvents", nil)
	return err
}

func (v *_VoltageMonitoringSensor) GetDipSwellThresholds(ctx context.Context) (DipSwellThresholds, error) {
	var ret DipSwellThresholds
	val, err := v.Caller().Call(ctx, v.RID(), "getDipSwellThresholds", nil)
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
	err = ret.Decode(res["_ret_"], v.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (v *_VoltageMonitoringSensor) SetDipSwellThresholds(ctx context.Context, in0 DipSwellThresholds) (int32, error) {
	var ret int32
	val, err := v.Caller().Call(ctx, v.RID(), "setDipSwellThresholds", map[string]any{
		"thresholds": in0.Encode(),
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
