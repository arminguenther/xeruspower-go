// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeemanager

import (
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/idl/event"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DeviceAddedEvent { return &_DeviceAddedEvent{} })
	valobj.Register(func() DeviceEvent { return &_DeviceEvent{} })
	valobj.Register(func() DeviceRemovedEvent { return &_DeviceRemovedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
}

type _DeviceEvent struct {
	event.Event
	sourceId int32
}

func (d *_DeviceEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeManager.DeviceEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceEvent) SourceId() int32 {
	return d.sourceId
}

func (d *_DeviceEvent) isDeviceEvent() {}

type _DeviceAddedEvent struct {
	DeviceEvent
}

func (d *_DeviceAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeManager.DeviceAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceAddedEvent) isDeviceAddedEvent() {}

type _DeviceRemovedEvent struct {
	DeviceEvent
}

func (d *_DeviceRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeManager.DeviceRemovedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceRemovedEvent) isDeviceRemovedEvent() {}

type _SettingsChangedEvent struct {
	event.Event
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeManager.SettingsChangedEvent",
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

type _StateChangedEvent struct {
	event.Event
	oldState DongleState
	newState DongleState
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeManager.StateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StateChangedEvent) OldState() DongleState {
	return s.oldState
}

func (s *_StateChangedEvent) NewState() DongleState {
	return s.newState
}

func (s *_StateChangedEvent) isStateChangedEvent() {}
