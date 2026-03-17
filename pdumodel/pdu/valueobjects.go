// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package pdu

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() LoadSheddingModeChangedEvent { return &_LoadSheddingModeChangedEvent{} })
	valobj.Register(func() OutletSequenceStateChangedEvent { return &_OutletSequenceStateChangedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu_6_1_1.SettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SettingsChangedEvent) OldSettings() Settings {
	return s.oldSettings
}

func (s *_SettingsChangedEvent) NewSettings() Settings {
	return s.newSettings
}

func (s *_SettingsChangedEvent) isSettingsChangedEvent() {}

type _LoadSheddingModeChangedEvent struct {
	userevent.UserEvent
	enabled bool
}

func (l *_LoadSheddingModeChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu_6_1_1.LoadSheddingModeChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LoadSheddingModeChangedEvent) Enabled() bool {
	return l.enabled
}

func (l *_LoadSheddingModeChangedEvent) isLoadSheddingModeChangedEvent() {}

type _OutletSequenceStateChangedEvent struct {
	event.Event
	newState OutletSequenceState
}

func (o *_OutletSequenceStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu_6_1_1.OutletSequenceStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (o *_OutletSequenceStateChangedEvent) NewState() OutletSequenceState {
	return o.newState
}

func (o *_OutletSequenceStateChangedEvent) isOutletSequenceStateChangedEvent() {}
