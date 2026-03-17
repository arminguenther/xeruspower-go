// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alarmmanager

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/idl/event"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() AlarmAcknowledgedEvent { return &_AlarmAcknowledgedEvent{} })
	valobj.Register(func() AlarmAddedEvent { return &_AlarmAddedEvent{} })
	valobj.Register(func() AlarmUpdatedEvent { return &_AlarmUpdatedEvent{} })
}

type _AlarmAddedEvent struct {
	event.Event
	alarm Alarm
}

func (a *_AlarmAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.AlarmManager.AlarmAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AlarmAddedEvent) Alarm() Alarm {
	return a.alarm
}

func (a *_AlarmAddedEvent) isAlarmAddedEvent() {}

type _AlarmUpdatedEvent struct {
	event.Event
	alarm Alarm
}

func (a *_AlarmUpdatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.AlarmManager.AlarmUpdatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AlarmUpdatedEvent) Alarm() Alarm {
	return a.alarm
}

func (a *_AlarmUpdatedEvent) isAlarmUpdatedEvent() {}

type _AlarmAcknowledgedEvent struct {
	event.Event
	alarmId string
}

func (a *_AlarmAcknowledgedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.AlarmManager.AlarmAcknowledgedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AlarmAcknowledgedEvent) AlarmId() string {
	return a.alarmId
}

func (a *_AlarmAcknowledgedEvent) isAlarmAcknowledgedEvent() {}
