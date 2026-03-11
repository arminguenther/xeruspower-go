// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package port

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40411/peripheral/poselement"
)

func (d *DeviceTypeWithId) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["id"] = d.Id
	j0["type"] = d.Type
	return j0
}

func (d *DeviceTypeWithId) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	d.Id, err = encoding.AsEnum[DeviceTypeId](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	d.Type, err = encoding.Is[string](j0["type"])
	if err != nil {
		return err
	}
	return nil
}

func (d *DetectionMode) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["type"] = d.Type
	j0["pinnedDeviceType"] = d.PinnedDeviceType
	return j0
}

func (d *DetectionMode) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	d.Type, err = encoding.AsEnum[DetectionType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("pinnedDeviceType", j0)
	if err != nil {
		return err
	}
	d.PinnedDeviceType, err = encoding.Is[string](j0["pinnedDeviceType"])
	if err != nil {
		return err
	}
	return nil
}

func (p *Properties) Encode() map[string]any {
	j0 := make(map[string]any, 12)
	j0["name"] = p.Name
	j0["label"] = p.Label
	j0["mode"] = p.Mode.Encode()
	j0["detectedDeviceType"] = p.DetectedDeviceType
	j0["detectedDeviceName"] = p.DetectedDeviceName
	j0["pinnedDeviceTypeId"] = p.PinnedDeviceTypeId
	j0["detectedDeviceTypeId"] = p.DetectedDeviceTypeId
	s1 := make([]any, 0, len(p.DetectableDeviceTypes))
	for _, e1 := range p.DetectableDeviceTypes {
		s1 = append(s1, e1.Encode())
	}
	j0["detectableDeviceTypes"] = s1
	j0["topoId"] = p.TopoId
	j0["serialId"] = p.SerialId
	j0["portType"] = p.PortType
	s2 := make([]any, 0, len(p.Position))
	for _, e2 := range p.Position {
		s2 = append(s2, e2.Encode())
	}
	j0["position"] = s2
	return j0
}

func (p *Properties) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	p.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	p.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("mode", j0)
	if err != nil {
		return err
	}
	err = p.Mode.Decode(j0["mode"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("detectedDeviceType", j0)
	if err != nil {
		return err
	}
	p.DetectedDeviceType, err = encoding.Is[string](j0["detectedDeviceType"])
	if err != nil {
		return err
	}
	err = encoding.In("detectedDeviceName", j0)
	if err != nil {
		return err
	}
	p.DetectedDeviceName, err = encoding.Is[string](j0["detectedDeviceName"])
	if err != nil {
		return err
	}
	err = encoding.In("pinnedDeviceTypeId", j0)
	if err != nil {
		return err
	}
	p.PinnedDeviceTypeId, err = encoding.AsEnum[DeviceTypeId](j0["pinnedDeviceTypeId"])
	if err != nil {
		return err
	}
	err = encoding.In("detectedDeviceTypeId", j0)
	if err != nil {
		return err
	}
	p.DetectedDeviceTypeId, err = encoding.AsEnum[DeviceTypeId](j0["detectedDeviceTypeId"])
	if err != nil {
		return err
	}
	err = encoding.In("detectableDeviceTypes", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["detectableDeviceTypes"])
	if err != nil {
		return err
	}
	p.DetectableDeviceTypes = make([]DeviceTypeWithId, 0, len(s1))
	for _, a1 := range s1 {
		var e1 DeviceTypeWithId
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		p.DetectableDeviceTypes = append(p.DetectableDeviceTypes, e1)
	}
	err = encoding.In("topoId", j0)
	if err != nil {
		return err
	}
	p.TopoId, err = encoding.Is[string](j0["topoId"])
	if err != nil {
		return err
	}
	err = encoding.In("serialId", j0)
	if err != nil {
		return err
	}
	p.SerialId, err = encoding.Is[string](j0["serialId"])
	if err != nil {
		return err
	}
	err = encoding.In("portType", j0)
	if err != nil {
		return err
	}
	p.PortType, err = encoding.AsEnum[poselement.PortType](j0["portType"])
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["position"])
	if err != nil {
		return err
	}
	p.Position = make([]poselement.PosElement, 0, len(s2))
	for _, a2 := range s2 {
		var e2 poselement.PosElement
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		p.Position = append(p.Position, e2)
	}
	return nil
}

func (p *_PropertiesChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.Event = valobj.For[event.Event]()
	err := p.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldProperties", value)
	if err != nil {
		return err
	}
	err = p.oldProperties.Decode(value["oldProperties"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newProperties", value)
	if err != nil {
		return err
	}
	err = p.newProperties.Decode(value["newProperties"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DeviceChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldDevice", value)
	if err != nil {
		return err
	}
	d.oldDevice, err = object.As[idl.Object](value["oldDevice"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newDevice", value)
	if err != nil {
		return err
	}
	d.newDevice, err = object.As[idl.Object](value["newDevice"], caller)
	if err != nil {
		return err
	}
	return nil
}
