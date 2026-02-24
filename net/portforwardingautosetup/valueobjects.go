// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package portforwardingautosetup

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() StatusChangedEvent { return &_StatusChangedEvent{} })
}

type _StatusChangedEvent struct {
	event.Event
	newStatus Status
}

func (p *_StatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.PortForwardingAutoSetupStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_StatusChangedEvent) NewStatus() Status {
	return p.newStatus
}

func (p *_StatusChangedEvent) isStatusChangedEvent() {}
