// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serialport

import (
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["hasModemSupport"] = m.HasModemSupport
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("hasModemSupport", j0)
	if err != nil {
		return err
	}
	m.HasModemSupport, err = encoding.Is[bool](j0["hasModemSupport"])
	if err != nil {
		return err
	}
	return nil
}

func (s *State) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["state"] = s.State
	j0["deviceName"] = s.DeviceName
	return j0
}

func (s *State) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	s.State, err = encoding.AsEnum[PortState](j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceName", j0)
	if err != nil {
		return err
	}
	s.DeviceName, err = encoding.Is[string](j0["deviceName"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["consoleBaudRate"] = s.ConsoleBaudRate
	j0["modemBaudRate"] = s.ModemBaudRate
	j0["detectType"] = s.DetectType
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("consoleBaudRate", j0)
	if err != nil {
		return err
	}
	s.ConsoleBaudRate, err = encoding.AsEnum[BaudRate](j0["consoleBaudRate"])
	if err != nil {
		return err
	}
	err = encoding.In("modemBaudRate", j0)
	if err != nil {
		return err
	}
	s.ModemBaudRate, err = encoding.AsEnum[BaudRate](j0["modemBaudRate"])
	if err != nil {
		return err
	}
	err = encoding.In("detectType", j0)
	if err != nil {
		return err
	}
	s.DetectType, err = encoding.AsEnum[DetectionType](j0["detectType"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModemEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.Event = valobj.For[event.Event]()
	err := m.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("modem", value)
	if err != nil {
		return err
	}
	m.modem, err = object.As[idl.Object](value["modem"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModemAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.ModemEvent = valobj.For[ModemEvent]()
	err := m.ModemEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModemRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.ModemEvent = valobj.For[ModemEvent]()
	err := m.ModemEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
