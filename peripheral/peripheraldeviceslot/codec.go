// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldeviceslot

import (
	"github.com/arminguenther/xeruspower-go/v40200/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/idl/event"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40200/peripheral/poselement"
	"github.com/arminguenther/xeruspower-go/v40200/sensors/sensor"
)

func (d *DeviceID) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["serial"] = d.Serial
	j0["type"] = d.Type.Encode()
	j0["isActuator"] = d.IsActuator
	j0["channel"] = d.Channel
	return j0
}

func (d *DeviceID) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("serial", j0)
	if err != nil {
		return err
	}
	d.Serial, err = encoding.Is[string](j0["serial"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	err = d.Type.Decode(j0["type"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("isActuator", j0)
	if err != nil {
		return err
	}
	d.IsActuator, err = encoding.Is[bool](j0["isActuator"])
	if err != nil {
		return err
	}
	err = encoding.In("channel", j0)
	if err != nil {
		return err
	}
	d.Channel, err = encoding.AsInt32(j0["channel"])
	if err != nil {
		return err
	}
	return nil
}

func (a *Address) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	s1 := make([]any, 0, len(a.Position))
	for _, e1 := range a.Position {
		s1 = append(s1, e1.Encode())
	}
	j0["position"] = s1
	j0["type"] = a.Type.Encode()
	j0["isActuator"] = a.IsActuator
	j0["channel"] = a.Channel
	return j0
}

func (a *Address) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["position"])
	if err != nil {
		return err
	}
	a.Position = make([]poselement.PosElement, 0, len(s1))
	for _, a1 := range s1 {
		var e1 poselement.PosElement
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		a.Position = append(a.Position, e1)
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	err = a.Type.Decode(j0["type"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("isActuator", j0)
	if err != nil {
		return err
	}
	a.IsActuator, err = encoding.Is[bool](j0["isActuator"])
	if err != nil {
		return err
	}
	err = encoding.In("channel", j0)
	if err != nil {
		return err
	}
	a.Channel, err = encoding.AsInt32(j0["channel"])
	if err != nil {
		return err
	}
	return nil
}

func (d *_Device) Decode(value map[string]any, caller idl.Caller) error {
	d.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("deviceID", value)
	if err != nil {
		return err
	}
	err = d.deviceID.Decode(value["deviceID"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("position", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["position"])
	if err != nil {
		return err
	}
	d.position = make([]poselement.PosElement, 0, len(s0))
	for _, a0 := range s0 {
		var e0 poselement.PosElement
		err = e0.Decode(a0, caller)
		if err != nil {
			return err
		}
		d.position = append(d.position, e0)
	}
	err = encoding.In("packageClass", value)
	if err != nil {
		return err
	}
	d.packageClass, err = encoding.Is[string](value["packageClass"])
	if err != nil {
		return err
	}
	err = encoding.In("device", value)
	if err != nil {
		return err
	}
	d.device, err = object.As[sensor.Sensor](value["device"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (l *DeviceSlotLocation) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["x"] = l.X
	j0["y"] = l.Y
	j0["z"] = l.Z
	return j0
}

func (l *DeviceSlotLocation) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("x", j0)
	if err != nil {
		return err
	}
	l.X, err = encoding.Is[string](j0["x"])
	if err != nil {
		return err
	}
	err = encoding.In("y", j0)
	if err != nil {
		return err
	}
	l.Y, err = encoding.Is[string](j0["y"])
	if err != nil {
		return err
	}
	err = encoding.In("z", j0)
	if err != nil {
		return err
	}
	l.Z, err = encoding.Is[string](j0["z"])
	if err != nil {
		return err
	}
	return nil
}

func (s *DeviceSlotSettings) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["name"] = s.Name
	j0["description"] = s.Description
	j0["location"] = s.Location.Encode()
	j0["useDefaultThresholds"] = s.UseDefaultThresholds
	s1 := make([]any, 0, len(s.Properties))
	for k1, v1 := range s.Properties {
		s1 = append(s1, map[string]any{"key": k1, "value": v1})
	}
	j0["properties"] = s1
	return j0
}

func (s *DeviceSlotSettings) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	s.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	err = encoding.In("location", j0)
	if err != nil {
		return err
	}
	err = s.Location.Decode(j0["location"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("useDefaultThresholds", j0)
	if err != nil {
		return err
	}
	s.UseDefaultThresholds, err = encoding.Is[bool](j0["useDefaultThresholds"])
	if err != nil {
		return err
	}
	err = encoding.In("properties", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["properties"])
	s.Properties = make(map[string]string, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 string
		v1, err = encoding.Is[string](b1)
		if err != nil {
			return err
		}
		s.Properties[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	return nil
}

func (d *_DeviceSlotDeviceChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldDevice", value)
	if err != nil {
		return err
	}
	d.oldDevice, err = valobj.As[Device](value["oldDevice"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newDevice", value)
	if err != nil {
		return err
	}
	d.newDevice, err = valobj.As[Device](value["newDevice"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_DeviceSlotSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
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
