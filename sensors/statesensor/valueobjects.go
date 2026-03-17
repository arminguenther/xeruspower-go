// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package statesensor

import (
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/idl/event"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
}

type _StateChangedEvent struct {
	event.Event
	oldState State
	newState State
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.StateSensor_4_0_5.StateChangedEvent",
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
