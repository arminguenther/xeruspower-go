// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorlogger

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/peripheral/peripheraldeviceslot"
	"github.com/arminguenther/xeruspower-go/sensors/sensor"
)

func (s *LoggerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["isEnabled"] = s.IsEnabled
	j0["samplePeriod"] = s.SamplePeriod
	j0["samplesPerRecord"] = s.SamplesPerRecord
	j0["oldestRecId"] = s.OldestRecId
	j0["newestRecId"] = s.NewestRecId
	j0["logCapacity"] = s.LogCapacity
	return j0
}

func (s *LoggerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("isEnabled", j0)
	if err != nil {
		return err
	}
	s.IsEnabled, err = encoding.Is[bool](j0["isEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("samplePeriod", j0)
	if err != nil {
		return err
	}
	s.SamplePeriod, err = encoding.AsInt32(j0["samplePeriod"])
	if err != nil {
		return err
	}
	err = encoding.In("samplesPerRecord", j0)
	if err != nil {
		return err
	}
	s.SamplesPerRecord, err = encoding.AsInt32(j0["samplesPerRecord"])
	if err != nil {
		return err
	}
	err = encoding.In("oldestRecId", j0)
	if err != nil {
		return err
	}
	s.OldestRecId, err = encoding.AsInt32(j0["oldestRecId"])
	if err != nil {
		return err
	}
	err = encoding.In("newestRecId", j0)
	if err != nil {
		return err
	}
	s.NewestRecId, err = encoding.AsInt32(j0["newestRecId"])
	if err != nil {
		return err
	}
	err = encoding.In("logCapacity", j0)
	if err != nil {
		return err
	}
	s.LogCapacity, err = encoding.AsInt32(j0["logCapacity"])
	if err != nil {
		return err
	}
	return nil
}

func (s *LoggerSensorSet) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	s1 := make([]any, 0, len(s.Sensors))
	for _, e1 := range s.Sensors {
		s1 = append(s1, object.ToMap(e1))
	}
	j0["sensors"] = s1
	s2 := make([]any, 0, len(s.Slots))
	for _, e2 := range s.Slots {
		s2 = append(s2, object.ToMap(e2))
	}
	j0["slots"] = s2
	return j0
}

func (s *LoggerSensorSet) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sensors", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["sensors"])
	if err != nil {
		return err
	}
	s.Sensors = make([]sensor.Sensor, 0, len(s1))
	for _, a1 := range s1 {
		var e1 sensor.Sensor
		e1, err = object.As[sensor.Sensor](a1, caller)
		if err != nil {
			return err
		}
		s.Sensors = append(s.Sensors, e1)
	}
	err = encoding.In("slots", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["slots"])
	if err != nil {
		return err
	}
	s.Slots = make([]peripheraldeviceslot.DeviceSlot, 0, len(s2))
	for _, a2 := range s2 {
		var e2 peripheraldeviceslot.DeviceSlot
		e2, err = object.As[peripheraldeviceslot.DeviceSlot](a2, caller)
		if err != nil {
			return err
		}
		s.Slots = append(s.Slots, e2)
	}
	return nil
}

func (s *_LoggerSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = s.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = s.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (l *_LoggerLoggedSensorsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.UserEvent = valobj.For[userevent.UserEvent]()
	err := l.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSensors", value)
	if err != nil {
		return err
	}
	err = l.oldSensors.Decode(value["oldSensors"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSensors", value)
	if err != nil {
		return err
	}
	err = l.newSensors.Decode(value["newSensors"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *LoggerRecord) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["available"] = r.Available
	j0["takenValidSamples"] = r.TakenValidSamples
	j0["state"] = r.State
	j0["minValue"] = r.MinValue
	j0["avgValue"] = r.AvgValue
	j0["maxValue"] = r.MaxValue
	return j0
}

func (r *LoggerRecord) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("available", j0)
	if err != nil {
		return err
	}
	r.Available, err = encoding.Is[bool](j0["available"])
	if err != nil {
		return err
	}
	err = encoding.In("takenValidSamples", j0)
	if err != nil {
		return err
	}
	r.TakenValidSamples, err = encoding.AsInt32(j0["takenValidSamples"])
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	r.State, err = encoding.AsInt32(j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("minValue", j0)
	if err != nil {
		return err
	}
	r.MinValue, err = encoding.AsFloat64(j0["minValue"])
	if err != nil {
		return err
	}
	err = encoding.In("avgValue", j0)
	if err != nil {
		return err
	}
	r.AvgValue, err = encoding.AsFloat64(j0["avgValue"])
	if err != nil {
		return err
	}
	err = encoding.In("maxValue", j0)
	if err != nil {
		return err
	}
	r.MaxValue, err = encoding.AsFloat64(j0["maxValue"])
	if err != nil {
		return err
	}
	return nil
}

func (t *LoggerTimedRecord) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["timestamp"] = t.Timestamp.Unix()
	j0["record"] = t.Record.Encode()
	return j0
}

func (t *LoggerTimedRecord) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	t.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("record", j0)
	if err != nil {
		return err
	}
	err = t.Record.Decode(j0["record"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (l *LoggerLogRow) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["sensorSetTimestamp"] = l.SensorSetTimestamp.Unix()
	j0["timestamp"] = l.Timestamp.Unix()
	s1 := make([]any, 0, len(l.SensorRecords))
	for _, e1 := range l.SensorRecords {
		s1 = append(s1, e1.Encode())
	}
	j0["sensorRecords"] = s1
	s2 := make([]any, 0, len(l.PeripheralDeviceRecords))
	for _, e2 := range l.PeripheralDeviceRecords {
		s2 = append(s2, e2.Encode())
	}
	j0["peripheralDeviceRecords"] = s2
	return j0
}

func (l *LoggerLogRow) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sensorSetTimestamp", j0)
	if err != nil {
		return err
	}
	l.SensorSetTimestamp, err = encoding.AsTime(j0["sensorSetTimestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	l.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("sensorRecords", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["sensorRecords"])
	if err != nil {
		return err
	}
	l.SensorRecords = make([]LoggerRecord, 0, len(s1))
	for _, a1 := range s1 {
		var e1 LoggerRecord
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		l.SensorRecords = append(l.SensorRecords, e1)
	}
	err = encoding.In("peripheralDeviceRecords", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["peripheralDeviceRecords"])
	if err != nil {
		return err
	}
	l.PeripheralDeviceRecords = make([]LoggerRecord, 0, len(s2))
	for _, a2 := range s2 {
		var e2 LoggerRecord
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		l.PeripheralDeviceRecords = append(l.PeripheralDeviceRecords, e2)
	}
	return nil
}
