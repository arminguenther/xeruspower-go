// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serialport

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ModemAddedEvent { return &_ModemAddedEvent{} })
	valobj.Register(func() ModemEvent { return &_ModemEvent{} })
	valobj.Register(func() ModemRemovedEvent { return &_ModemRemovedEvent{} })
}

type _ModemEvent struct {
	event.Event
	modem idl.Object
}

func (m *_ModemEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.SerialPort_3_0_1.ModemEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModemEvent) Modem() idl.Object {
	return m.modem
}

func (m *_ModemEvent) isModemEvent() {}

type _ModemAddedEvent struct {
	ModemEvent
}

func (m *_ModemAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.SerialPort_3_0_1.ModemAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModemAddedEvent) isModemAddedEvent() {}

type _ModemRemovedEvent struct {
	ModemEvent
}

func (m *_ModemRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.SerialPort_3_0_1.ModemRemovedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModemRemovedEvent) isModemRemovedEvent() {}
