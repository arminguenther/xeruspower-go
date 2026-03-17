// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package transferswitch

import (
	"github.com/arminguenther/xeruspower-go/v40040/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.TransferSwitch_5_0_2.SettingsChangedEvent",
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
