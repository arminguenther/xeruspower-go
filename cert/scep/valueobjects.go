// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package scep

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/idl/event"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() StatusChangedEvent { return &_StatusChangedEvent{} })
}

type _StatusChangedEvent struct {
	event.Event
	newStatus Status
}

func (s *_StatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cert.Scep.ScepStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StatusChangedEvent) NewStatus() Status {
	return s.newStatus
}

func (s *_StatusChangedEvent) isStatusChangedEvent() {}
