// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gsmmodem

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/idl/event"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() SimPinUpdatedEvent { return &_SimPinUpdatedEvent{} })
	valobj.Register(func() SimSecurityStatusChangedEvent { return &_SimSecurityStatusChangedEvent{} })
}

type _SimSecurityStatusChangedEvent struct {
	event.Event
	newSimStatus SimSecurityStatus
}

func (s *_SimSecurityStatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.GsmModem_1_0_2.SimSecurityStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SimSecurityStatusChangedEvent) NewSimStatus() SimSecurityStatus {
	return s.newSimStatus
}

func (s *_SimSecurityStatusChangedEvent) isSimSecurityStatusChangedEvent() {}

type _SimPinUpdatedEvent struct {
	event.Event
	newPin string
}

func (s *_SimPinUpdatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.GsmModem_1_0_2.SimPinUpdatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SimPinUpdatedEvent) NewPin() string {
	return s.newPin
}

func (s *_SimPinUpdatedEvent) isSimPinUpdatedEvent() {}
