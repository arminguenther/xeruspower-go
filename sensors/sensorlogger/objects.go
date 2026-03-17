// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorlogger

import (
	"context"
	"time"

	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40032/peripheral/peripheraldeviceslot"
	"github.com/arminguenther/xeruspower-go/v40032/sensors/sensor"
)

func init() {
	object.Register(NewLogger)
}

type _Logger struct {
	idl.Object
}

// NewLogger returns a new Logger interface for the object at given RID.
func NewLogger(rid string, caller idl.Caller) Logger {
	return &_Logger{idl.NewObject(rid, caller)}
}

func (l *_Logger) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Logger",
		Major: 3, Submajor: 0, Minor: 1,
	}
}

func (l *_Logger) GetInfo(ctx context.Context) (LoggerInfo, error) {
	var ret LoggerInfo
	val, err := l.Caller().Call(ctx, l.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], l.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (l *_Logger) GetSettings(ctx context.Context) (LoggerSettings, error) {
	var ret LoggerSettings
	val, err := l.Caller().Call(ctx, l.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], l.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (l *_Logger) SetSettings(ctx context.Context, in0 LoggerSettings) (int32, error) {
	var ret int32
	val, err := l.Caller().Call(ctx, l.RID(), "setSettings", map[string]any{
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

func (l *_Logger) GetTimeStamps(ctx context.Context, in0 int32, in1 int32) (int32, []time.Time, error) {
	var ret int32
	var out0 []time.Time
	val, err := l.Caller().Call(ctx, l.RID(), "getTimeStamps", map[string]any{
		"recid": in0,
		"count": in1,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("timestamps", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["timestamps"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]time.Time, 0, len(s0))
	for _, a0 := range s0 {
		var e0 time.Time
		e0, err = encoding.AsTime(a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (l *_Logger) GetSensorRecords(ctx context.Context, in0 sensor.Sensor, in1 int32, in2 int32) (int32, []LoggerRecord, error) {
	var ret int32
	var out0 []LoggerRecord
	val, err := l.Caller().Call(ctx, l.RID(), "getSensorRecords", map[string]any{
		"sensor": object.ToMap(in0),
		"recid":  in1,
		"count":  in2,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("recs", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["recs"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]LoggerRecord, 0, len(s0))
	for _, a0 := range s0 {
		var e0 LoggerRecord
		err = e0.Decode(a0, l.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (l *_Logger) GetPeripheralDeviceRecords(ctx context.Context, in0 peripheraldeviceslot.DeviceSlot, in1 int32, in2 int32) (int32, []LoggerRecord, error) {
	var ret int32
	var out0 []LoggerRecord
	val, err := l.Caller().Call(ctx, l.RID(), "getPeripheralDeviceRecords", map[string]any{
		"slot":  object.ToMap(in0),
		"recid": in1,
		"count": in2,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("recs", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["recs"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]LoggerRecord, 0, len(s0))
	for _, a0 := range s0 {
		var e0 LoggerRecord
		err = e0.Decode(a0, l.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (l *_Logger) GetSensorTimedRecords(ctx context.Context, in0 sensor.Sensor, in1 int32, in2 int32) (int32, []LoggerTimedRecord, error) {
	var ret int32
	var out0 []LoggerTimedRecord
	val, err := l.Caller().Call(ctx, l.RID(), "getSensorTimedRecords", map[string]any{
		"sensor": object.ToMap(in0),
		"recid":  in1,
		"count":  in2,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("recs", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["recs"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]LoggerTimedRecord, 0, len(s0))
	for _, a0 := range s0 {
		var e0 LoggerTimedRecord
		err = e0.Decode(a0, l.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (l *_Logger) GetPeripheralDeviceTimedRecords(ctx context.Context, in0 peripheraldeviceslot.DeviceSlot, in1 int32, in2 int32) (int32, []LoggerTimedRecord, error) {
	var ret int32
	var out0 []LoggerTimedRecord
	val, err := l.Caller().Call(ctx, l.RID(), "getPeripheralDeviceTimedRecords", map[string]any{
		"slot":  object.ToMap(in0),
		"recid": in1,
		"count": in2,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("recs", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["recs"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]LoggerTimedRecord, 0, len(s0))
	for _, a0 := range s0 {
		var e0 LoggerTimedRecord
		err = e0.Decode(a0, l.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (l *_Logger) GetLoggedSensors(ctx context.Context) (LoggerSensorSet, error) {
	var ret LoggerSensorSet
	val, err := l.Caller().Call(ctx, l.RID(), "getLoggedSensors", nil)
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
	err = ret.Decode(res["_ret_"], l.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (l *_Logger) SetLoggedSensors(ctx context.Context, in0 LoggerSensorSet) (int32, error) {
	var ret int32
	val, err := l.Caller().Call(ctx, l.RID(), "setLoggedSensors", map[string]any{
		"sensors": in0.Encode(),
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

func (l *_Logger) EnableSensors(ctx context.Context, in0 LoggerSensorSet) (int32, error) {
	var ret int32
	val, err := l.Caller().Call(ctx, l.RID(), "enableSensors", map[string]any{
		"sensors": in0.Encode(),
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

func (l *_Logger) DisableSensors(ctx context.Context, in0 LoggerSensorSet) (int32, error) {
	var ret int32
	val, err := l.Caller().Call(ctx, l.RID(), "disableSensors", map[string]any{
		"sensors": in0.Encode(),
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

func (l *_Logger) IsSensorEnabled(ctx context.Context, in0 sensor.Sensor) (bool, error) {
	var ret bool
	val, err := l.Caller().Call(ctx, l.RID(), "isSensorEnabled", map[string]any{
		"sensor": object.ToMap(in0),
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (l *_Logger) IsSlotEnabled(ctx context.Context, in0 peripheraldeviceslot.DeviceSlot) (bool, error) {
	var ret bool
	val, err := l.Caller().Call(ctx, l.RID(), "isSlotEnabled", map[string]any{
		"slot": object.ToMap(in0),
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (l *_Logger) EnableAllSensors(ctx context.Context) error {
	_, err := l.Caller().Call(ctx, l.RID(), "enableAllSensors", nil)
	return err
}

func (l *_Logger) DisableAllSensors(ctx context.Context) error {
	_, err := l.Caller().Call(ctx, l.RID(), "disableAllSensors", nil)
	return err
}

func (l *_Logger) GetSensorSetTimestamp(ctx context.Context) (time.Time, error) {
	var ret time.Time
	val, err := l.Caller().Call(ctx, l.RID(), "getSensorSetTimestamp", nil)
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
	ret, err = encoding.AsTime(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (l *_Logger) GetLogRow(ctx context.Context, in0 int32) (int32, LoggerLogRow, error) {
	var ret int32
	var out0 LoggerLogRow
	val, err := l.Caller().Call(ctx, l.RID(), "getLogRow", map[string]any{
		"recid": in0,
	})
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("row", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["row"], l.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}
