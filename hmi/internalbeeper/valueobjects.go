// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package internalbeeper

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/idl/event"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() MuteChangedEvent { return &_MuteChangedEvent{} })
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
}

type _MuteChangedEvent struct {
	userevent.UserEvent
	muted bool
}

func (m *_MuteChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "hmi.InternalBeeper_2_0_1.MuteChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MuteChangedEvent) Muted() bool {
	return m.muted
}

func (m *_MuteChangedEvent) isMuteChangedEvent() {}

type _StateChangedEvent struct {
	event.Event
	state            State
	reason           string
	mutedTemporarily bool
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "hmi.InternalBeeper_2_0_1.StateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StateChangedEvent) State() State {
	return s.state
}

func (s *_StateChangedEvent) Reason() string {
	return s.reason
}

func (s *_StateChangedEvent) MutedTemporarily() bool {
	return s.mutedTemporarily
}

func (s *_StateChangedEvent) isStateChangedEvent() {}
