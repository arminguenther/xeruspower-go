// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alertedsensormanager

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40411/sensors/sensor"
)

func init() {
	valobj.Register(func() AlertedSensorsChangedEvent { return &_AlertedSensorsChangedEvent{} })
	valobj.Register(func() MonitoredSensorsChangedEvent { return &_MonitoredSensorsChangedEvent{} })
}

type _MonitoredSensorsChangedEvent struct {
	event.Event
	counts SensorCounts
}

func (m *_MonitoredSensorsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.AlertedSensorManager_1_0_5.MonitoredSensorsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MonitoredSensorsChangedEvent) Counts() SensorCounts {
	return m.counts
}

func (m *_MonitoredSensorsChangedEvent) isMonitoredSensorsChangedEvent() {}

type _AlertedSensorsChangedEvent struct {
	event.Event
	counts         SensorCounts
	changedSensors []SensorData
	removedSensors []sensor.Sensor
}

func (a *_AlertedSensorsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.AlertedSensorManager_1_0_5.AlertedSensorsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AlertedSensorsChangedEvent) Counts() SensorCounts {
	return a.counts
}

func (a *_AlertedSensorsChangedEvent) ChangedSensors() []SensorData {
	return a.changedSensors
}

func (a *_AlertedSensorsChangedEvent) RemovedSensors() []sensor.Sensor {
	return a.removedSensors
}

func (a *_AlertedSensorsChangedEvent) isAlertedSensorsChangedEvent() {}
