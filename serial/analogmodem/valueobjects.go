// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package analogmodem

import (
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() CallEndedEvent { return &_CallEndedEvent{} })
	valobj.Register(func() CallReceivedEvent { return &_CallReceivedEvent{} })
	valobj.Register(func() DialInEvent { return &_DialInEvent{} })
}

type _DialInEvent struct {
	event.Event
	number string
}

func (d *_DialInEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.AnalogModem.DialInEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DialInEvent) Number() string {
	return d.number
}

func (d *_DialInEvent) isDialInEvent() {}

type _CallReceivedEvent struct {
	DialInEvent
}

func (c *_CallReceivedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.AnalogModem.CallReceivedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CallReceivedEvent) isCallReceivedEvent() {}

type _CallEndedEvent struct {
	DialInEvent
	disconnectedRemotely bool
}

func (c *_CallEndedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.AnalogModem.CallEndedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CallEndedEvent) DisconnectedRemotely() bool {
	return c.disconnectedRemotely
}

func (c *_CallEndedEvent) isCallEndedEvent() {}
