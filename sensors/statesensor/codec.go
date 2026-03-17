// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package statesensor

import (
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/idl/event"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
)

func (s *State) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["timestamp"] = s.Timestamp.Unix()
	j0["available"] = s.Available
	j0["value"] = s.Value
	return j0
}

func (s *State) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	s.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("available", j0)
	if err != nil {
		return err
	}
	s.Available, err = encoding.Is[bool](j0["available"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	s.Value, err = encoding.AsInt32(j0["value"])
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
	err = encoding.In("oldState", value)
	if err != nil {
		return err
	}
	err = s.oldState.Decode(value["oldState"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	err = s.newState.Decode(value["newState"], caller)
	if err != nil {
		return err
	}
	return nil
}
