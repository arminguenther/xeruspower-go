// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package pdu

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/idl/event"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() LoadSheddingModeChangedEvent { return &_LoadSheddingModeChangedEvent{} })
	valobj.Register(func() OutletSequenceStateChangedEvent { return &_OutletSequenceStateChangedEvent{} })
	valobj.Register(func() PortAppearedEvent { return &_PortAppearedEvent{} })
	valobj.Register(func() PortDisappearedEvent { return &_PortDisappearedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu_6_5_5.SettingsChangedEvent",
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
		Name:  "pdumodel.Pdu_6_5_5.LoadSheddingModeChangedEvent",
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
		Name:  "pdumodel.Pdu_6_5_5.OutletSequenceStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (o *_OutletSequenceStateChangedEvent) NewState() OutletSequenceState {
	return o.newState
}

func (o *_OutletSequenceStateChangedEvent) isOutletSequenceStateChangedEvent() {}

type _PortAppearedEvent struct {
	event.Event
	port PortWithProperties
}

func (p *_PortAppearedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu_6_5_5.PortAppearedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PortAppearedEvent) Port() PortWithProperties {
	return p.port
}

func (p *_PortAppearedEvent) isPortAppearedEvent() {}

type _PortDisappearedEvent struct {
	event.Event
	portId string
}

func (p *_PortDisappearedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu_6_5_5.PortDisappearedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PortDisappearedEvent) PortId() string {
	return p.portId
}

func (p *_PortDisappearedEvent) isPortDisappearedEvent() {}
