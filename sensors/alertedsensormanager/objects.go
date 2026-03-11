// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alertedsensormanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
)

func init() {
	object.Register(NewAlertedSensorManager)
}

type _AlertedSensorManager struct {
	idl.Object
}

// NewAlertedSensorManager returns a new AlertedSensorManager interface for the object at given RID.
func NewAlertedSensorManager(rid string, caller idl.Caller) AlertedSensorManager {
	return &_AlertedSensorManager{idl.NewObject(rid, caller)}
}

func (a *_AlertedSensorManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.AlertedSensorManager",
		Major: 1, Submajor: 0, Minor: 5,
	}
}

func (a *_AlertedSensorManager) GetSensorCounts(ctx context.Context) (SensorCounts, error) {
	var ret SensorCounts
	val, err := a.Caller().Call(ctx, a.RID(), "getSensorCounts", nil)
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
	err = ret.Decode(res["_ret_"], a.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AlertedSensorManager) GetAllSensors(ctx context.Context) ([]SensorData, error) {
	var ret []SensorData
	val, err := a.Caller().Call(ctx, a.RID(), "getAllSensors", nil)
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
	ret = make([]SensorData, 0, len(s0))
	for _, a0 := range s0 {
		var e0 SensorData
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (a *_AlertedSensorManager) GetAlertedSensors(ctx context.Context) ([]SensorData, error) {
	var ret []SensorData
	val, err := a.Caller().Call(ctx, a.RID(), "getAlertedSensors", nil)
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
	ret = make([]SensorData, 0, len(s0))
	for _, a0 := range s0 {
		var e0 SensorData
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
