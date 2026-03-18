// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package port

import (
	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/idl/event"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DeviceChangedEvent { return &_DeviceChangedEvent{} })
	valobj.Register(func() PropertiesChangedEvent { return &_PropertiesChangedEvent{} })
}

type _PropertiesChangedEvent struct {
	event.Event
	oldProperties Properties
	newProperties Properties
}

func (p *_PropertiesChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "portsmodel.Port_2_0_3.PropertiesChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PropertiesChangedEvent) OldProperties() Properties {
	return p.oldProperties
}

func (p *_PropertiesChangedEvent) NewProperties() Properties {
	return p.newProperties
}

func (p *_PropertiesChangedEvent) isPropertiesChangedEvent() {}

type _DeviceChangedEvent struct {
	event.Event
	oldDevice idl.Object
	newDevice idl.Object
}

func (d *_DeviceChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "portsmodel.Port_2_0_3.DeviceChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceChangedEvent) OldDevice() idl.Object {
	return d.oldDevice
}

func (d *_DeviceChangedEvent) NewDevice() idl.Object {
	return d.newDevice
}

func (d *_DeviceChangedEvent) isDeviceChangedEvent() {}
