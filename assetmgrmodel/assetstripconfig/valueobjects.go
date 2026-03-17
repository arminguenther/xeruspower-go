// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstripconfig

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() RackUnitSettingsChangedEvent { return &_RackUnitSettingsChangedEvent{} })
	valobj.Register(func() StripSettingsChangedEvent { return &_StripSettingsChangedEvent{} })
}

type _StripSettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings StripSettings
	newSettings StripSettings
}

func (s *_StripSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStripConfig_1_0_1.StripSettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StripSettingsChangedEvent) OldSettings() StripSettings {
	return s.oldSettings
}

func (s *_StripSettingsChangedEvent) NewSettings() StripSettings {
	return s.newSettings
}

func (s *_StripSettingsChangedEvent) isStripSettingsChangedEvent() {}

type _RackUnitSettingsChangedEvent struct {
	userevent.UserEvent
	rackUnitNumber int32
	oldSettings    RackUnitSettings
	newSettings    RackUnitSettings
}

func (r *_RackUnitSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStripConfig_1_0_1.RackUnitSettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RackUnitSettingsChangedEvent) RackUnitNumber() int32 {
	return r.rackUnitNumber
}

func (r *_RackUnitSettingsChangedEvent) OldSettings() RackUnitSettings {
	return r.oldSettings
}

func (r *_RackUnitSettingsChangedEvent) NewSettings() RackUnitSettings {
	return r.newSettings
}

func (r *_RackUnitSettingsChangedEvent) isRackUnitSettingsChangedEvent() {}
