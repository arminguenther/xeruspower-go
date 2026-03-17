// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicepackage

import (
	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/idl/event"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() BatteryPoweredDevicePackageVoltageChangedEvent {
		return &_BatteryPoweredDevicePackageVoltageChangedEvent{}
	})
	valobj.Register(func() DoorHandleControllerPackageDoorForcedOpenEvent {
		return &_DoorHandleControllerPackageDoorForcedOpenEvent{}
	})
	valobj.Register(func() DoorHandleControllerPackageMechanicallyUnlockedEvent {
		return &_DoorHandleControllerPackageMechanicallyUnlockedEvent{}
	})
}

type _DoorHandleControllerPackageMechanicallyUnlockedEvent struct {
	event.Event
	packageInfo    PackageInfo
	channel        int32
	doorStateName  string
	doorHandleName string
	doorLockName   string
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DoorHandleControllerPackage_3_0_2.MechanicallyUnlockedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) PackageInfo() PackageInfo {
	return m.packageInfo
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) Channel() int32 {
	return m.channel
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) DoorStateName() string {
	return m.doorStateName
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) DoorHandleName() string {
	return m.doorHandleName
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) DoorLockName() string {
	return m.doorLockName
}

func (m *_DoorHandleControllerPackageMechanicallyUnlockedEvent) isDoorHandleControllerPackageMechanicallyUnlockedEvent() {
}

type _DoorHandleControllerPackageDoorForcedOpenEvent struct {
	event.Event
	packageInfo    PackageInfo
	channel        int32
	doorStateName  string
	doorHandleName string
	doorLockName   string
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DoorHandleControllerPackage_3_0_2.DoorForcedOpenEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) PackageInfo() PackageInfo {
	return d.packageInfo
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) Channel() int32 {
	return d.channel
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) DoorStateName() string {
	return d.doorStateName
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) DoorHandleName() string {
	return d.doorHandleName
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) DoorLockName() string {
	return d.doorLockName
}

func (d *_DoorHandleControllerPackageDoorForcedOpenEvent) isDoorHandleControllerPackageDoorForcedOpenEvent() {
}

type _BatteryPoweredDevicePackageVoltageChangedEvent struct {
	event.Event
	oldVoltage float64
	newVoltage float64
}

func (v *_BatteryPoweredDevicePackageVoltageChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.BatteryPoweredDevicePackage_1_0_1.VoltageChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (v *_BatteryPoweredDevicePackageVoltageChangedEvent) OldVoltage() float64 {
	return v.oldVoltage
}

func (v *_BatteryPoweredDevicePackageVoltageChangedEvent) NewVoltage() float64 {
	return v.newVoltage
}

func (v *_BatteryPoweredDevicePackageVoltageChangedEvent) isBatteryPoweredDevicePackageVoltageChangedEvent() {
}
