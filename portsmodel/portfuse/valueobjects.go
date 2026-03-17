// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portfuse

import (
	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/idl/event"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() StatusChangedEvent { return &_StatusChangedEvent{} })
}

type _StatusChangedEvent struct {
	event.Event
	oldStatus Status
	newStatus Status
}

func (s *_StatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "portsmodel.PortFuse_1_0_1.StatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StatusChangedEvent) OldStatus() Status {
	return s.oldStatus
}

func (s *_StatusChangedEvent) NewStatus() Status {
	return s.newStatus
}

func (s *_StatusChangedEvent) isStatusChangedEvent() {}
