// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package servermonitor

import (
	"github.com/arminguenther/xeruspower-go/v40200/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/idl/event"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ServerAddedEvent { return &_ServerAddedEvent{} })
	valobj.Register(func() ServerDeletedEvent { return &_ServerDeletedEvent{} })
	valobj.Register(func() ServerPowerControlCompletedEvent { return &_ServerPowerControlCompletedEvent{} })
	valobj.Register(func() ServerPowerControlInitiatedEvent { return &_ServerPowerControlInitiatedEvent{} })
	valobj.Register(func() ServerPowerStateEvent { return &_ServerPowerStateEvent{} })
	valobj.Register(func() ServerReachabilityEvent { return &_ServerReachabilityEvent{} })
	valobj.Register(func() ServerSettingsChangedEvent { return &_ServerSettingsChangedEvent{} })
}

type _ServerPowerStateEvent struct {
	event.Event
	id            int32
	host          string
	oldPowerState ServerPowerState
	newPowerState ServerPowerState
}

func (s *_ServerPowerStateEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerPowerStateEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerPowerStateEvent) Id() int32 {
	return s.id
}

func (s *_ServerPowerStateEvent) Host() string {
	return s.host
}

func (s *_ServerPowerStateEvent) OldPowerState() ServerPowerState {
	return s.oldPowerState
}

func (s *_ServerPowerStateEvent) NewPowerState() ServerPowerState {
	return s.newPowerState
}

func (s *_ServerPowerStateEvent) isServerPowerStateEvent() {}

type _ServerPowerControlInitiatedEvent struct {
	userevent.UserEvent
	id   int32
	host string
	on   bool
}

func (s *_ServerPowerControlInitiatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerPowerControlInitiatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerPowerControlInitiatedEvent) Id() int32 {
	return s.id
}

func (s *_ServerPowerControlInitiatedEvent) Host() string {
	return s.host
}

func (s *_ServerPowerControlInitiatedEvent) On() bool {
	return s.on
}

func (s *_ServerPowerControlInitiatedEvent) isServerPowerControlInitiatedEvent() {}

type _ServerPowerControlCompletedEvent struct {
	event.Event
	id     int32
	host   string
	result ServerPowerControlResult
}

func (s *_ServerPowerControlCompletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerPowerControlCompletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerPowerControlCompletedEvent) Id() int32 {
	return s.id
}

func (s *_ServerPowerControlCompletedEvent) Host() string {
	return s.host
}

func (s *_ServerPowerControlCompletedEvent) Result() ServerPowerControlResult {
	return s.result
}

func (s *_ServerPowerControlCompletedEvent) isServerPowerControlCompletedEvent() {}

type _ServerReachabilityEvent struct {
	event.Event
	id        int32
	host      string
	reachable ServerReachability
}

func (s *_ServerReachabilityEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerReachabilityEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerReachabilityEvent) Id() int32 {
	return s.id
}

func (s *_ServerReachabilityEvent) Host() string {
	return s.host
}

func (s *_ServerReachabilityEvent) Reachable() ServerReachability {
	return s.reachable
}

func (s *_ServerReachabilityEvent) isServerReachabilityEvent() {}

type _ServerAddedEvent struct {
	userevent.UserEvent
	id       int32
	settings ServerSettings
}

func (s *_ServerAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerAddedEvent) Id() int32 {
	return s.id
}

func (s *_ServerAddedEvent) Settings() ServerSettings {
	return s.settings
}

func (s *_ServerAddedEvent) isServerAddedEvent() {}

type _ServerSettingsChangedEvent struct {
	userevent.UserEvent
	id          int32
	oldSettings ServerSettings
	newSettings ServerSettings
}

func (s *_ServerSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerSettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerSettingsChangedEvent) Id() int32 {
	return s.id
}

func (s *_ServerSettingsChangedEvent) OldSettings() ServerSettings {
	return s.oldSettings
}

func (s *_ServerSettingsChangedEvent) NewSettings() ServerSettings {
	return s.newSettings
}

func (s *_ServerSettingsChangedEvent) isServerSettingsChangedEvent() {}

type _ServerDeletedEvent struct {
	userevent.UserEvent
	id int32
}

func (s *_ServerDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor_2_0_1.ServerDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ServerDeletedEvent) Id() int32 {
	return s.id
}

func (s *_ServerDeletedEvent) isServerDeletedEvent() {}
