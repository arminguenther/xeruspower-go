// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package inlet

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() EnableStateChangedEvent { return &_EnableStateChangedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Inlet_3_0_3.SettingsChangedEvent",
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

type _EnableStateChangedEvent struct {
	userevent.UserEvent
	enabled bool
}

func (e *_EnableStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Inlet_3_0_3.EnableStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EnableStateChangedEvent) Enabled() bool {
	return e.enabled
}

func (e *_EnableStateChangedEvent) isEnableStateChangedEvent() {}
