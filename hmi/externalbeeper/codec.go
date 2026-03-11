// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package externalbeeper

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
)

func (s *_StateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldState", value)
	if err != nil {
		return err
	}
	s.oldState, err = encoding.AsEnum[State](value["oldState"])
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	s.newState, err = encoding.AsEnum[State](value["newState"])
	if err != nil {
		return err
	}
	return nil
}
