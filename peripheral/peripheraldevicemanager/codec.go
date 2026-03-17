// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicemanager

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/peripheral/peripheraldevicepackage"
	"github.com/arminguenther/xeruspower-go/peripheral/peripheraldeviceslot"
	"github.com/arminguenther/xeruspower-go/sensors/numericsensor"
)

func (s *DeviceManagerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["zCoordMode"] = s.ZCoordMode
	j0["autoManageNewDevices"] = s.AutoManageNewDevices
	j0["deviceAltitude"] = s.DeviceAltitude
	j0["presenceDetectionTimeout"] = s.PresenceDetectionTimeout
	s1 := make([]any, 0, len(s.DefaultThresholdsMap))
	for k1, v1 := range s.DefaultThresholdsMap {
		s1 = append(s1, map[string]any{"key": k1, "value": v1.Encode()})
	}
	j0["defaultThresholdsMap"] = s1
	j0["maxActivePoweredDryContacts"] = s.MaxActivePoweredDryContacts
	j0["muteOtherAccessControlUnit"] = s.MuteOtherAccessControlUnit
	return j0
}

func (s *DeviceManagerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("zCoordMode", j0)
	if err != nil {
		return err
	}
	s.ZCoordMode, err = encoding.AsEnum[DeviceManagerZCoordMode](j0["zCoordMode"])
	if err != nil {
		return err
	}
	err = encoding.In("autoManageNewDevices", j0)
	if err != nil {
		return err
	}
	s.AutoManageNewDevices, err = encoding.Is[bool](j0["autoManageNewDevices"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceAltitude", j0)
	if err != nil {
		return err
	}
	s.DeviceAltitude, err = encoding.AsFloat32(j0["deviceAltitude"])
	if err != nil {
		return err
	}
	err = encoding.In("presenceDetectionTimeout", j0)
	if err != nil {
		return err
	}
	s.PresenceDetectionTimeout, err = encoding.AsInt32(j0["presenceDetectionTimeout"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultThresholdsMap", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["defaultThresholdsMap"])
	s.DefaultThresholdsMap = make(map[string]numericsensor.Thresholds, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 numericsensor.Thresholds
		err = v1.Decode(b1, caller)
		if err != nil {
			return err
		}
		s.DefaultThresholdsMap[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	err = encoding.In("maxActivePoweredDryContacts", j0)
	if err != nil {
		return err
	}
	s.MaxActivePoweredDryContacts, err = encoding.AsInt32(j0["maxActivePoweredDryContacts"])
	if err != nil {
		return err
	}
	err = encoding.In("muteOtherAccessControlUnit", j0)
	if err != nil {
		return err
	}
	s.MuteOtherAccessControlUnit, err = encoding.Is[bool](j0["muteOtherAccessControlUnit"])
	if err != nil {
		return err
	}
	return nil
}

func (m *DeviceManagerMetaData) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["oneWirePortCount"] = m.OneWirePortCount
	j0["onboardDeviceCount"] = m.OnboardDeviceCount
	return j0
}

func (m *DeviceManagerMetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("oneWirePortCount", j0)
	if err != nil {
		return err
	}
	m.OneWirePortCount, err = encoding.AsInt32(j0["oneWirePortCount"])
	if err != nil {
		return err
	}
	err = encoding.In("onboardDeviceCount", j0)
	if err != nil {
		return err
	}
	m.OnboardDeviceCount, err = encoding.AsInt32(j0["onboardDeviceCount"])
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceManagerDeviceTypeInfo) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["type"] = d.Type.Encode()
	j0["isActuator"] = d.IsActuator
	j0["identifier"] = d.Identifier
	j0["name"] = d.Name
	j0["defaultRange"] = d.DefaultRange.Encode()
	j0["defaultDecDigits"] = d.DefaultDecDigits
	return j0
}

func (d *DeviceManagerDeviceTypeInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
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
	err = encoding.In("identifier", j0)
	if err != nil {
		return err
	}
	d.Identifier, err = encoding.Is[string](j0["identifier"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	d.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultRange", j0)
	if err != nil {
		return err
	}
	err = d.DefaultRange.Decode(j0["defaultRange"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("defaultDecDigits", j0)
	if err != nil {
		return err
	}
	d.DefaultDecDigits, err = encoding.AsInt32(j0["defaultDecDigits"])
	if err != nil {
		return err
	}
	return nil
}

func (f *DeviceManagerFirmwareUpdateState) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["active"] = f.Active
	j0["remaining"] = f.Remaining
	return j0
}

func (f *DeviceManagerFirmwareUpdateState) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("active", j0)
	if err != nil {
		return err
	}
	f.Active, err = encoding.Is[bool](j0["active"])
	if err != nil {
		return err
	}
	err = encoding.In("remaining", j0)
	if err != nil {
		return err
	}
	f.Remaining, err = encoding.AsInt32(j0["remaining"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_DeviceManagerSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
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

func (d *_DeviceManagerDeviceEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("devices", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["devices"])
	if err != nil {
		return err
	}
	d.devices = make([]peripheraldeviceslot.Device, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldeviceslot.Device
		e0, err = valobj.As[peripheraldeviceslot.Device](a0, caller)
		if err != nil {
			return err
		}
		d.devices = append(d.devices, e0)
	}
	err = encoding.In("allDevices", value)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](value["allDevices"])
	if err != nil {
		return err
	}
	d.allDevices = make([]peripheraldeviceslot.Device, 0, len(s1))
	for _, a1 := range s1 {
		var e1 peripheraldeviceslot.Device
		e1, err = valobj.As[peripheraldeviceslot.Device](a1, caller)
		if err != nil {
			return err
		}
		d.allDevices = append(d.allDevices, e1)
	}
	return nil
}

func (d *_DeviceManagerDeviceAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.DeviceManagerDeviceEvent = valobj.For[DeviceManagerDeviceEvent]()
	err := d.DeviceManagerDeviceEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DeviceManagerDeviceRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.DeviceManagerDeviceEvent = valobj.For[DeviceManagerDeviceEvent]()
	err := d.DeviceManagerDeviceEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (u *_DeviceManagerUnknownDeviceAttachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	u.Event = valobj.For[event.Event]()
	err := u.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("romCode", value)
	if err != nil {
		return err
	}
	u.romCode, err = encoding.Is[string](value["romCode"])
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
	u.position = make([]peripheraldeviceslot.PosElement, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldeviceslot.PosElement
		err = e0.Decode(a0, caller)
		if err != nil {
			return err
		}
		u.position = append(u.position, e0)
	}
	return nil
}

func (d *_DeviceManagerDeviceFirmwareUpdateStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldVersion", value)
	if err != nil {
		return err
	}
	d.oldVersion, err = encoding.Is[string](value["oldVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("newVersion", value)
	if err != nil {
		return err
	}
	d.newVersion, err = encoding.Is[string](value["newVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("serial", value)
	if err != nil {
		return err
	}
	d.serial, err = encoding.Is[string](value["serial"])
	if err != nil {
		return err
	}
	err = encoding.In("state", value)
	if err != nil {
		return err
	}
	d.state, err = encoding.AsEnum[DeviceManagerDeviceFirmwareUpdateState](value["state"])
	if err != nil {
		return err
	}
	return nil
}

func (f *_DeviceManagerFirmwareUpdateStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.Event = valobj.For[event.Event]()
	err := f.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	err = f.newState.Decode(value["newState"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_DeviceManagerPackageEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.Event = valobj.For[event.Event]()
	err := p.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("packageInfos", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["packageInfos"])
	if err != nil {
		return err
	}
	p.packageInfos = make([]peripheraldevicepackage.PackageInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldevicepackage.PackageInfo
		err = e0.Decode(a0, caller)
		if err != nil {
			return err
		}
		p.packageInfos = append(p.packageInfos, e0)
	}
	err = encoding.In("allPackages", value)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](value["allPackages"])
	if err != nil {
		return err
	}
	p.allPackages = make([]peripheraldevicepackage.PackageInfo, 0, len(s1))
	for _, a1 := range s1 {
		var e1 peripheraldevicepackage.PackageInfo
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		p.allPackages = append(p.allPackages, e1)
	}
	return nil
}

func (p *_DeviceManagerPackageAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.DeviceManagerPackageEvent = valobj.For[DeviceManagerPackageEvent]()
	err := p.DeviceManagerPackageEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_DeviceManagerPackageRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.DeviceManagerPackageEvent = valobj.For[DeviceManagerPackageEvent]()
	err := p.DeviceManagerPackageEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *DeviceManagerStatistics) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["cSumErrCnt"] = s.CSumErrCnt
	j0["fuseTripCnt"] = s.FuseTripCnt
	return j0
}

func (s *DeviceManagerStatistics) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("cSumErrCnt", j0)
	if err != nil {
		return err
	}
	s.CSumErrCnt, err = encoding.AsInt32(j0["cSumErrCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("fuseTripCnt", j0)
	if err != nil {
		return err
	}
	s.FuseTripCnt, err = encoding.AsInt32(j0["fuseTripCnt"])
	if err != nil {
		return err
	}
	return nil
}
