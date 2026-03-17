// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicepackage

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() BatteryPoweredDevicePackageVoltageChangedEvent {
		return &_BatteryPoweredDevicePackageVoltageChangedEvent{}
	})
}

type _BatteryPoweredDevicePackageVoltageChangedEvent struct {
	event.Event
	oldVoltage float64
	newVoltage float64
}

func (v *_BatteryPoweredDevicePackageVoltageChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.BatteryPoweredDevicePackage.VoltageChangedEvent",
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
