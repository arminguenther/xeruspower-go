// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package port

import (
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

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
	j0 := make(map[string]any, 5)
	j0["name"] = p.Name
	j0["label"] = p.Label
	j0["mode"] = p.Mode.Encode()
	j0["detectedDeviceType"] = p.DetectedDeviceType
	j0["detectedDeviceName"] = p.DetectedDeviceName
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
