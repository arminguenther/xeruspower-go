// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alertedsensormanager

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/idl/event"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40220/sensors/sensor"
)

func (s *SensorCounts) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["total"] = s.Total
	j0["unavailable"] = s.Unavailable
	j0["critical"] = s.Critical
	j0["warned"] = s.Warned
	return j0
}

func (s *SensorCounts) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("total", j0)
	if err != nil {
		return err
	}
	s.Total, err = encoding.AsInt32(j0["total"])
	if err != nil {
		return err
	}
	err = encoding.In("unavailable", j0)
	if err != nil {
		return err
	}
	s.Unavailable, err = encoding.AsInt32(j0["unavailable"])
	if err != nil {
		return err
	}
	err = encoding.In("critical", j0)
	if err != nil {
		return err
	}
	s.Critical, err = encoding.AsInt32(j0["critical"])
	if err != nil {
		return err
	}
	err = encoding.In("warned", j0)
	if err != nil {
		return err
	}
	s.Warned, err = encoding.AsInt32(j0["warned"])
	if err != nil {
		return err
	}
	return nil
}

func (s *SensorData) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["sensor"] = object.ToMap(s.Sensor)
	j0["parent"] = object.ToMap(s.Parent)
	j0["alertState"] = s.AlertState
	return j0
}

func (s *SensorData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sensor", j0)
	if err != nil {
		return err
	}
	s.Sensor, err = object.As[sensor.Sensor](j0["sensor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("parent", j0)
	if err != nil {
		return err
	}
	s.Parent, err = object.As[idl.Object](j0["parent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("alertState", j0)
	if err != nil {
		return err
	}
	s.AlertState, err = encoding.AsEnum[AlertState](j0["alertState"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_MonitoredSensorsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.Event = valobj.For[event.Event]()
	err := m.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("counts", value)
	if err != nil {
		return err
	}
	err = m.counts.Decode(value["counts"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *_AlertedSensorsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	a.Event = valobj.For[event.Event]()
	err := a.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("counts", value)
	if err != nil {
		return err
	}
	err = a.counts.Decode(value["counts"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("changedSensors", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["changedSensors"])
	if err != nil {
		return err
	}
	a.changedSensors = make([]SensorData, 0, len(s0))
	for _, a0 := range s0 {
		var e0 SensorData
		err = e0.Decode(a0, caller)
		if err != nil {
			return err
		}
		a.changedSensors = append(a.changedSensors, e0)
	}
	err = encoding.In("removedSensors", value)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](value["removedSensors"])
	if err != nil {
		return err
	}
	a.removedSensors = make([]sensor.Sensor, 0, len(s1))
	for _, a1 := range s1 {
		var e1 sensor.Sensor
		e1, err = object.As[sensor.Sensor](a1, caller)
		if err != nil {
			return err
		}
		a.removedSensors = append(a.removedSensors, e1)
	}
	return nil
}
