// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorlogger

import (
	"github.com/arminguenther/xeruspower-go/v40220/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/idl/event"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() LoggerInfoChangedEvent { return &_LoggerInfoChangedEvent{} })
	valobj.Register(func() LoggerLoggedSensorsChangedEvent { return &_LoggerLoggedSensorsChangedEvent{} })
	valobj.Register(func() LoggerSettingsChangedEvent { return &_LoggerSettingsChangedEvent{} })
}

type _LoggerInfoChangedEvent struct {
	event.Event
	oldInfo LoggerInfo
	newInfo LoggerInfo
}

func (i *_LoggerInfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Logger_3_1_3.InfoChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_LoggerInfoChangedEvent) OldInfo() LoggerInfo {
	return i.oldInfo
}

func (i *_LoggerInfoChangedEvent) NewInfo() LoggerInfo {
	return i.newInfo
}

func (i *_LoggerInfoChangedEvent) isLoggerInfoChangedEvent() {}

type _LoggerSettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings LoggerSettings
	newSettings LoggerSettings
}

func (s *_LoggerSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Logger_3_1_3.SettingsChangedEvent",
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
		Name:  "sensors.Logger_3_1_3.LoggedSensorsChangedEvent",
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
