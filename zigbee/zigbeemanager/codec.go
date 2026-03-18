// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeemanager

import (
	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/idl/event"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/valobj"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["firmwareVersion"] = m.FirmwareVersion
	j0["stackVersion"] = m.StackVersion
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("firmwareVersion", j0)
	if err != nil {
		return err
	}
	m.FirmwareVersion, err = encoding.Is[string](j0["firmwareVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("stackVersion", j0)
	if err != nil {
		return err
	}
	m.StackVersion, err = encoding.Is[string](j0["stackVersion"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["channel"] = s.Channel
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("channel", j0)
	if err != nil {
		return err
	}
	s.Channel, err = encoding.AsInt32(j0["channel"])
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceRegistration) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["sourceId"] = d.SourceId
	j0["key"] = d.Key
	j0["slot"] = d.Slot
	return j0
}

func (d *DeviceRegistration) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sourceId", j0)
	if err != nil {
		return err
	}
	d.SourceId, err = encoding.AsInt32(j0["sourceId"])
	if err != nil {
		return err
	}
	err = encoding.In("key", j0)
	if err != nil {
		return err
	}
	d.Key, err = encoding.Is[string](j0["key"])
	if err != nil {
		return err
	}
	err = encoding.In("slot", j0)
	if err != nil {
		return err
	}
	d.Slot, err = encoding.AsInt32(j0["slot"])
	if err != nil {
		return err
	}
	return nil
}

func (d *_DeviceEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("sourceId", value)
	if err != nil {
		return err
	}
	d.sourceId, err = encoding.AsInt32(value["sourceId"])
	if err != nil {
		return err
	}
	return nil
}

func (d *_DeviceAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.DeviceEvent = valobj.For[DeviceEvent]()
	err := d.DeviceEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DeviceRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.DeviceEvent = valobj.For[DeviceEvent]()
	err := d.DeviceEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
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

func (s *_StateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldState", value)
	if err != nil {
		return err
	}
	s.oldState, err = encoding.AsEnum[DongleState](value["oldState"])
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	s.newState, err = encoding.AsEnum[DongleState](value["newState"])
	if err != nil {
		return err
	}
	return nil
}
