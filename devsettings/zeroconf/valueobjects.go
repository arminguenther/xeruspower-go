// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zeroconf

import (
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/idl/event"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	event.Event
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Zeroconf_2_0_0.SettingsChangedEvent",
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
