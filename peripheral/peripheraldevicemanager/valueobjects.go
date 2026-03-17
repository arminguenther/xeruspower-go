// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicemanager

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/idl/event"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40300/peripheral/peripheraldevicepackage"
	"github.com/arminguenther/xeruspower-go/v40300/peripheral/peripheraldeviceslot"
	"github.com/arminguenther/xeruspower-go/v40300/peripheral/poselement"
)

func init() {
	valobj.Register(func() DeviceManagerDeviceAddedEvent { return &_DeviceManagerDeviceAddedEvent{} })
	valobj.Register(func() DeviceManagerDeviceEvent { return &_DeviceManagerDeviceEvent{} })
	valobj.Register(func() DeviceManagerDeviceFirmwareUpdateStateChangedEvent {
		return &_DeviceManagerDeviceFirmwareUpdateStateChangedEvent{}
	})
	valobj.Register(func() DeviceManagerDeviceRemovedEvent { return &_DeviceManagerDeviceRemovedEvent{} })
	valobj.Register(func() DeviceManagerFirmwareUpdateStateChangedEvent {
		return &_DeviceManagerFirmwareUpdateStateChangedEvent{}
	})
	valobj.Register(func() DeviceManagerPackageAddedEvent { return &_DeviceManagerPackageAddedEvent{} })
	valobj.Register(func() DeviceManagerPackageEvent { return &_DeviceManagerPackageEvent{} })
	valobj.Register(func() DeviceManagerPackageRemovedEvent { return &_DeviceManagerPackageRemovedEvent{} })
	valobj.Register(func() DeviceManagerSettingsChangedEvent { return &_DeviceManagerSettingsChangedEvent{} })
	valobj.Register(func() DeviceManagerUnknownDeviceAttachedEvent { return &_DeviceManagerUnknownDeviceAttachedEvent{} })
}

type _DeviceManagerSettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings DeviceManagerSettings
	newSettings DeviceManagerSettings
}

func (s *_DeviceManagerSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.SettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_DeviceManagerSettingsChangedEvent) OldSettings() DeviceManagerSettings {
	return s.oldSettings
}

func (s *_DeviceManagerSettingsChangedEvent) NewSettings() DeviceManagerSettings {
	return s.newSettings
}

func (s *_DeviceManagerSettingsChangedEvent) isDeviceManagerSettingsChangedEvent() {}

type _DeviceManagerDeviceEvent struct {
	event.Event
	devices    []peripheraldeviceslot.Device
	allDevices []peripheraldeviceslot.Device
}

func (d *_DeviceManagerDeviceEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.DeviceEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceManagerDeviceEvent) Devices() []peripheraldeviceslot.Device {
	return d.devices
}

func (d *_DeviceManagerDeviceEvent) AllDevices() []peripheraldeviceslot.Device {
	return d.allDevices
}

func (d *_DeviceManagerDeviceEvent) isDeviceManagerDeviceEvent() {}

type _DeviceManagerDeviceAddedEvent struct {
	DeviceManagerDeviceEvent
}

func (d *_DeviceManagerDeviceAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.DeviceAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceManagerDeviceAddedEvent) isDeviceManagerDeviceAddedEvent() {}

type _DeviceManagerDeviceRemovedEvent struct {
	DeviceManagerDeviceEvent
}

func (d *_DeviceManagerDeviceRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.DeviceRemovedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceManagerDeviceRemovedEvent) isDeviceManagerDeviceRemovedEvent() {}

type _DeviceManagerUnknownDeviceAttachedEvent struct {
	event.Event
	romCode  string
	position []poselement.PosElement
}

func (u *_DeviceManagerUnknownDeviceAttachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.UnknownDeviceAttachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (u *_DeviceManagerUnknownDeviceAttachedEvent) RomCode() string {
	return u.romCode
}

func (u *_DeviceManagerUnknownDeviceAttachedEvent) Position() []poselement.PosElement {
	return u.position
}

func (u *_DeviceManagerUnknownDeviceAttachedEvent) isDeviceManagerUnknownDeviceAttachedEvent() {}

type _DeviceManagerDeviceFirmwareUpdateStateChangedEvent struct {
	event.Event
	oldVersion string
	newVersion string
	serial     string
	state      DeviceManagerDeviceFirmwareUpdateState
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.DeviceFirmwareUpdateStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) OldVersion() string {
	return d.oldVersion
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) NewVersion() string {
	return d.newVersion
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) Serial() string {
	return d.serial
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) State() DeviceManagerDeviceFirmwareUpdateState {
	return d.state
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) isDeviceManagerDeviceFirmwareUpdateStateChangedEvent() {
}

type _DeviceManagerFirmwareUpdateStateChangedEvent struct {
	event.Event
	newState DeviceManagerFirmwareUpdateState
}

func (f *_DeviceManagerFirmwareUpdateStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.FirmwareUpdateStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_DeviceManagerFirmwareUpdateStateChangedEvent) NewState() DeviceManagerFirmwareUpdateState {
	return f.newState
}

func (f *_DeviceManagerFirmwareUpdateStateChangedEvent) isDeviceManagerFirmwareUpdateStateChangedEvent() {
}

type _DeviceManagerPackageEvent struct {
	event.Event
	packageInfos []peripheraldevicepackage.PackageInfo
	allPackages  []peripheraldevicepackage.PackageInfo
}

func (p *_DeviceManagerPackageEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.PackageEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_DeviceManagerPackageEvent) PackageInfos() []peripheraldevicepackage.PackageInfo {
	return p.packageInfos
}

func (p *_DeviceManagerPackageEvent) AllPackages() []peripheraldevicepackage.PackageInfo {
	return p.allPackages
}

func (p *_DeviceManagerPackageEvent) isDeviceManagerPackageEvent() {}

type _DeviceManagerPackageAddedEvent struct {
	DeviceManagerPackageEvent
}

func (p *_DeviceManagerPackageAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.PackageAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_DeviceManagerPackageAddedEvent) isDeviceManagerPackageAddedEvent() {}

type _DeviceManagerPackageRemovedEvent struct {
	DeviceManagerPackageEvent
}

func (p *_DeviceManagerPackageRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager_5_3_5.PackageRemovedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_DeviceManagerPackageRemovedEvent) isDeviceManagerPackageRemovedEvent() {}
