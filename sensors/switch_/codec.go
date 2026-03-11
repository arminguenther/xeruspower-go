// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package switch_

import (
	"github.com/arminguenther/xeruspower-go/v40413/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func (s *_Event) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("targetState", value)
	if err != nil {
		return err
	}
	s.targetState, err = encoding.AsInt32(value["targetState"])
	if err != nil {
		return err
	}
	return nil
}
