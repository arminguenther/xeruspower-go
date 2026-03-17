// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldeviceslot

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/sensors/sensor"
)

func init() {
	valobj.Register(func() Device { return &_Device{} })
	valobj.Register(func() DeviceSlotDeviceChangedEvent { return &_DeviceSlotDeviceChangedEvent{} })
	valobj.Register(func() DeviceSlotSettingsChangedEvent { return &_DeviceSlotSettingsChangedEvent{} })
}

type _Device struct {
	idl.ValueObject
	deviceID     DeviceID
	position     []PosElement
	packageClass string
	device       sensor.Sensor
}

func (d *_Device) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.Device",
		Major: 7, Submajor: 0, Minor: 1,
	}
}

func (d *_Device) DeviceID() DeviceID {
	return d.deviceID
}

func (d *_Device) Position() []PosElement {
	return d.position
}

func (d *_Device) PackageClass() string {
	return d.packageClass
}

func (d *_Device) Device() sensor.Sensor {
	return d.device
}

func (d *_Device) isDevice() {}

type _DeviceSlotDeviceChangedEvent struct {
	event.Event
	oldDevice Device
	newDevice Device
}

func (d *_DeviceSlotDeviceChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceSlot_5_0_1.DeviceChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceSlotDeviceChangedEvent) OldDevice() Device {
	return d.oldDevice
}

func (d *_DeviceSlotDeviceChangedEvent) NewDevice() Device {
	return d.newDevice
}

func (d *_DeviceSlotDeviceChangedEvent) isDeviceSlotDeviceChangedEvent() {}

type _DeviceSlotSettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings DeviceSlotSettings
	newSettings DeviceSlotSettings
}

func (s *_DeviceSlotSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceSlot_5_0_1.SettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_DeviceSlotSettingsChangedEvent) OldSettings() DeviceSlotSettings {
	return s.oldSettings
}

func (s *_DeviceSlotSettingsChangedEvent) NewSettings() DeviceSlotSettings {
	return s.newSettings
}

func (s *_DeviceSlotSettingsChangedEvent) isDeviceSlotSettingsChangedEvent() {}
