// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsamport

import (
	"github.com/arminguenther/xeruspower-go/v40211/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() InfoChangedEvent { return &_InfoChangedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _InfoChangedEvent struct {
	userevent.UserEvent
	oldInfo  Info
	newInfo  Info
	portName string
}

func (i *_InfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamPort.InfoChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InfoChangedEvent) OldInfo() Info {
	return i.oldInfo
}

func (i *_InfoChangedEvent) NewInfo() Info {
	return i.newInfo
}

func (i *_InfoChangedEvent) PortName() string {
	return i.portName
}

func (i *_InfoChangedEvent) isInfoChangedEvent() {}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	dsamNumber  int32
	portNumber  int32
	oldSettings Settings
	newSettings Settings
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamPort.SettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SettingsChangedEvent) DsamNumber() int32 {
	return s.dsamNumber
}

func (s *_SettingsChangedEvent) PortNumber() int32 {
	return s.portNumber
}

func (s *_SettingsChangedEvent) OldSettings() Settings {
	return s.oldSettings
}

func (s *_SettingsChangedEvent) NewSettings() Settings {
	return s.newSettings
}

func (s *_SettingsChangedEvent) isSettingsChangedEvent() {}
