// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outletgroup

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/pdumodel/outlet"
)

func init() {
	valobj.Register(func() PowerControlEvent { return &_PowerControlEvent{} })
	valobj.Register(func() SensorsChangedEvent { return &_SensorsChangedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SensorsChangedEvent struct {
	event.Event
	oldSensors Sensors
	newSensors Sensors
}

func (s *_SensorsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroup_1_1_10.SensorsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SensorsChangedEvent) OldSensors() Sensors {
	return s.oldSensors
}

func (s *_SensorsChangedEvent) NewSensors() Sensors {
	return s.newSensors
}

func (s *_SensorsChangedEvent) isSensorsChangedEvent() {}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroup_1_1_10.SettingsChangedEvent",
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

type _PowerControlEvent struct {
	userevent.UserEvent
	state outlet.PowerState
	cycle bool
}

func (p *_PowerControlEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroup_1_1_10.PowerControlEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PowerControlEvent) State() outlet.PowerState {
	return p.state
}

func (p *_PowerControlEvent) Cycle() bool {
	return p.cycle
}

func (p *_PowerControlEvent) isPowerControlEvent() {}
