// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package switch_

import (
	"github.com/arminguenther/xeruspower-go/v40413/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() Event { return &_Event{} })
}

type _Event struct {
	userevent.UserEvent
	targetState int32
}

func (s *_Event) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Switch_2_0_9.SwitchEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_Event) TargetState() int32 {
	return s.targetState
}

func (s *_Event) isEvent() {}
