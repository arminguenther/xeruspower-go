// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package shelfpowersupply

import (
	"github.com/arminguenther/xeruspower-go/v40412/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/idl/event"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() FanSpeedControlEvent { return &_FanSpeedControlEvent{} })
	valobj.Register(func() MetaDataChangedEvent { return &_MetaDataChangedEvent{} })
	valobj.Register(func() PowerControlEvent { return &_PowerControlEvent{} })
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
}

type _PowerControlEvent struct {
	userevent.UserEvent
	state PowerState
}

func (p *_PowerControlEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.ShelfPowerSupply.PowerControlEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PowerControlEvent) State() PowerState {
	return p.state
}

func (p *_PowerControlEvent) isPowerControlEvent() {}

type _FanSpeedControlEvent struct {
	userevent.UserEvent
	overrideEnabled bool
	targetPercent   int32
}

func (f *_FanSpeedControlEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.ShelfPowerSupply.FanSpeedControlEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FanSpeedControlEvent) OverrideEnabled() bool {
	return f.overrideEnabled
}

func (f *_FanSpeedControlEvent) TargetPercent() int32 {
	return f.targetPercent
}

func (f *_FanSpeedControlEvent) isFanSpeedControlEvent() {}

type _StateChangedEvent struct {
	event.Event
	oldState State
	newState State
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.ShelfPowerSupply.StateChangedEvent",
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

type _MetaDataChangedEvent struct {
	event.Event
	oldMetaData MetaData
	newMetaData MetaData
}

func (m *_MetaDataChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.ShelfPowerSupply.MetaDataChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MetaDataChangedEvent) OldMetaData() MetaData {
	return m.oldMetaData
}

func (m *_MetaDataChangedEvent) NewMetaData() MetaData {
	return m.newMetaData
}

func (m *_MetaDataChangedEvent) isMetaDataChangedEvent() {}
