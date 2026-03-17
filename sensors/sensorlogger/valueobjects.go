// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorlogger

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() LoggerLoggedSensorsChangedEvent { return &_LoggerLoggedSensorsChangedEvent{} })
	valobj.Register(func() LoggerSettingsChangedEvent { return &_LoggerSettingsChangedEvent{} })
}

type _LoggerSettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings LoggerSettings
	newSettings LoggerSettings
}

func (s *_LoggerSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Logger_2_3_8.SettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_LoggerSettingsChangedEvent) OldSettings() LoggerSettings {
	return s.oldSettings
}

func (s *_LoggerSettingsChangedEvent) NewSettings() LoggerSettings {
	return s.newSettings
}

func (s *_LoggerSettingsChangedEvent) isLoggerSettingsChangedEvent() {}

type _LoggerLoggedSensorsChangedEvent struct {
	userevent.UserEvent
	oldSensors LoggerSensorSet
	newSensors LoggerSensorSet
}

func (l *_LoggerLoggedSensorsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Logger_2_3_8.LoggedSensorsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LoggerLoggedSensorsChangedEvent) OldSensors() LoggerSensorSet {
	return l.oldSensors
}

func (l *_LoggerLoggedSensorsChangedEvent) NewSensors() LoggerSensorSet {
	return l.newSensors
}

func (l *_LoggerLoggedSensorsChangedEvent) isLoggerLoggedSensorsChangedEvent() {}
