// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package voltagemonitoringsensor

import (
	"github.com/arminguenther/xeruspower-go/v40200/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	event_ "github.com/arminguenther/xeruspower-go/v40200/idl/event"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DipSwellThresholdsChangedEvent { return &_DipSwellThresholdsChangedEvent{} })
	valobj.Register(func() EventListClearedEvent { return &_EventListClearedEvent{} })
	valobj.Register(func() EventOccurredEvent { return &_EventOccurredEvent{} })
}

type __Event = event_.Event

type _EventOccurredEvent struct {
	__Event
	event Event
}

func (e *_EventOccurredEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.VoltageMonitoringSensor_1_0_3.EventOccurredEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EventOccurredEvent) Event() Event {
	return e.event
}

func (e *_EventOccurredEvent) isEventOccurredEvent() {}

type _EventListClearedEvent struct {
	userevent.UserEvent
}

func (e *_EventListClearedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.VoltageMonitoringSensor_1_0_3.EventListClearedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EventListClearedEvent) isEventListClearedEvent() {}

type _DipSwellThresholdsChangedEvent struct {
	userevent.UserEvent
	oldThresholds DipSwellThresholds
	newThresholds DipSwellThresholds
}

func (d *_DipSwellThresholdsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.VoltageMonitoringSensor_1_0_3.DipSwellThresholdsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DipSwellThresholdsChangedEvent) OldThresholds() DipSwellThresholds {
	return d.oldThresholds
}

func (d *_DipSwellThresholdsChangedEvent) NewThresholds() DipSwellThresholds {
	return d.newThresholds
}

func (d *_DipSwellThresholdsChangedEvent) isDipSwellThresholdsChangedEvent() {}
