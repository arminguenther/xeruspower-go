// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package internalbeeper

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func (m *_MuteChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.UserEvent = valobj.For[userevent.UserEvent]()
	err := m.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("muted", value)
	if err != nil {
		return err
	}
	m.muted, err = encoding.Is[bool](value["muted"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_StateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("state", value)
	if err != nil {
		return err
	}
	s.state, err = encoding.AsEnum[State](value["state"])
	if err != nil {
		return err
	}
	err = encoding.In("reason", value)
	if err != nil {
		return err
	}
	s.reason, err = encoding.Is[string](value["reason"])
	if err != nil {
		return err
	}
	err = encoding.In("mutedTemporarily", value)
	if err != nil {
		return err
	}
	s.mutedTemporarily, err = encoding.Is[bool](value["mutedTemporarily"])
	if err != nil {
		return err
	}
	return nil
}
