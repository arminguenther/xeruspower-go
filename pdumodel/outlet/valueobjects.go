// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outlet

import (
	"github.com/arminguenther/xeruspower-go/v40200/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/idl/event"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() PowerControlEvent { return &_PowerControlEvent{} })
	valobj.Register(func() ServiceModeChangedEvent { return &_ServiceModeChangedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
}

type _PowerControlEvent struct {
	userevent.UserEvent
	state PowerState
	cycle bool
}

func (p *_PowerControlEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Outlet_3_0_3.PowerControlEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PowerControlEvent) State() PowerState {
	return p.state
}

func (p *_PowerControlEvent) Cycle() bool {
	return p.cycle
}

func (p *_PowerControlEvent) isPowerControlEvent() {}

type _StateChangedEvent struct {
	event.Event
	oldState State
	newState State
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Outlet_3_0_3.StateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StateChangedEvent) OldState() State {
	return s.oldState
}

func (s *_StateChangedEvent) NewState() State {
	return s.newState
}

func (s *_StateChangedEvent) isStateChangedEvent() {}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Outlet_3_0_3.SettingsChangedEvent",
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

type _ServiceModeChangedEvent struct {
	userevent.UserEvent
	enabled bool
}

func (s *_ServiceModeChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Outlet_3_0_3.ServiceModeChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServiceModeChangedEvent) Enabled() bool {
	return s.enabled
}

func (s *_ServiceModeChangedEvent) isServiceModeChangedEvent() {}
