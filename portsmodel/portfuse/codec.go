// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portfuse

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func (s *_StatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldStatus", value)
	if err != nil {
		return err
	}
	s.oldStatus, err = encoding.AsEnum[Status](value["oldStatus"])
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	s.newStatus, err = encoding.AsEnum[Status](value["newStatus"])
	if err != nil {
		return err
	}
	return nil
}
