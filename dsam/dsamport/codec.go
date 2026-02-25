// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsamport

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["dsamNumber"] = i.DsamNumber
	j0["portNumber"] = i.PortNumber
	j0["connected"] = i.Connected
	j0["devIfType"] = i.DevIfType
	j0["state"] = i.State
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("dsamNumber", j0)
	if err != nil {
		return err
	}
	i.DsamNumber, err = encoding.AsInt32(j0["dsamNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("portNumber", j0)
	if err != nil {
		return err
	}
	i.PortNumber, err = encoding.AsInt32(j0["portNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("connected", j0)
	if err != nil {
		return err
	}
	i.Connected, err = encoding.Is[bool](j0["connected"])
	if err != nil {
		return err
	}
	err = encoding.In("devIfType", j0)
	if err != nil {
		return err
	}
	i.DevIfType, err = encoding.AsEnum[DeviceInterfaceType](j0["devIfType"])
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	i.State, err = encoding.AsEnum[State](j0["state"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["name"] = s.Name
	j0["devIfType"] = s.DevIfType
	j0["baudRate"] = s.BaudRate
	j0["parity"] = s.Parity
	j0["stopBits"] = s.StopBits
	j0["flowCtrl"] = s.FlowCtrl
	j0["breakDurationMs"] = s.BreakDurationMs
	j0["sshDpaPortEnabled"] = s.SshDpaPortEnabled
	j0["sshDpaPort"] = s.SshDpaPort
	j0["allowSharedAccess"] = s.AllowSharedAccess
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	s.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("devIfType", j0)
	if err != nil {
		return err
	}
	s.DevIfType, err = encoding.AsEnum[DeviceInterfaceType](j0["devIfType"])
	if err != nil {
		return err
	}
	err = encoding.In("baudRate", j0)
	if err != nil {
		return err
	}
	s.BaudRate, err = encoding.AsInt32(j0["baudRate"])
	if err != nil {
		return err
	}
	err = encoding.In("parity", j0)
	if err != nil {
		return err
	}
	s.Parity, err = encoding.AsEnum[Parity](j0["parity"])
	if err != nil {
		return err
	}
	err = encoding.In("stopBits", j0)
	if err != nil {
		return err
	}
	s.StopBits, err = encoding.AsInt32(j0["stopBits"])
	if err != nil {
		return err
	}
	err = encoding.In("flowCtrl", j0)
	if err != nil {
		return err
	}
	s.FlowCtrl, err = encoding.AsEnum[FlowControl](j0["flowCtrl"])
	if err != nil {
		return err
	}
	err = encoding.In("breakDurationMs", j0)
	if err != nil {
		return err
	}
	s.BreakDurationMs, err = encoding.AsInt32(j0["breakDurationMs"])
	if err != nil {
		return err
	}
	err = encoding.In("sshDpaPortEnabled", j0)
	if err != nil {
		return err
	}
	s.SshDpaPortEnabled, err = encoding.Is[bool](j0["sshDpaPortEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("sshDpaPort", j0)
	if err != nil {
		return err
	}
	s.SshDpaPort, err = encoding.AsInt32(j0["sshDpaPort"])
	if err != nil {
		return err
	}
	err = encoding.In("allowSharedAccess", j0)
	if err != nil {
		return err
	}
	s.AllowSharedAccess, err = encoding.Is[bool](j0["allowSharedAccess"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_InfoChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	i.UserEvent = valobj.For[userevent.UserEvent]()
	err := i.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldInfo", value)
	if err != nil {
		return err
	}
	err = i.oldInfo.Decode(value["oldInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newInfo", value)
	if err != nil {
		return err
	}
	err = i.newInfo.Decode(value["newInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("portName", value)
	if err != nil {
		return err
	}
	i.portName, err = encoding.Is[string](value["portName"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("dsamNumber", value)
	if err != nil {
		return err
	}
	s.dsamNumber, err = encoding.AsInt32(value["dsamNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("portNumber", value)
	if err != nil {
		return err
	}
	s.portNumber, err = encoding.AsInt32(value["portNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = s.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = s.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}
