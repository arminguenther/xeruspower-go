// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkconfiguration

import (
	"github.com/arminguenther/xeruspower-go/v40410/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() RestoredEvent { return &_RestoredEvent{} })
	valobj.Register(func() SavedEvent { return &_SavedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	userevent.UserEvent
}

func (s *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "bulkcfg.BulkConfiguration_1_0_2.SettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SettingsChangedEvent) isSettingsChangedEvent() {}

type _SavedEvent struct {
	userevent.UserEvent
	isBackup bool
}

func (s *_SavedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "bulkcfg.BulkConfiguration_1_0_2.SavedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SavedEvent) IsBackup() bool {
	return s.isBackup
}

func (s *_SavedEvent) isSavedEvent() {}

type _RestoredEvent struct {
	userevent.UserEvent
	isBackup bool
}

func (r *_RestoredEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "bulkcfg.BulkConfiguration_1_0_2.RestoredEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RestoredEvent) IsBackup() bool {
	return r.isBackup
}

func (r *_RestoredEvent) isRestoredEvent() {}
