// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicemanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40020/peripheral/gatewaysensormanager"
	"github.com/arminguenther/xeruspower-go/v40020/peripheral/peripheraldevicepackage"
	"github.com/arminguenther/xeruspower-go/v40020/peripheral/peripheraldeviceslot"
	"github.com/arminguenther/xeruspower-go/v40020/peripheral/sensorhub"
	"github.com/arminguenther/xeruspower-go/v40020/portsmodel/portfuse"
)

func init() {
	object.Register(NewDeviceManager)
}

type _DeviceManager struct {
	idl.Object
}

// NewDeviceManager returns a new DeviceManager interface for the object at given RID.
func NewDeviceManager(rid string, caller idl.Caller) DeviceManager {
	return &_DeviceManager{idl.NewObject(rid, caller)}
}

func (d *_DeviceManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceManager",
		Major: 5, Submajor: 0, Minor: 3,
	}
}

func (d *_DeviceManager) GetDeviceSlots(ctx context.Context) ([]peripheraldeviceslot.DeviceSlot, error) {
	var ret []peripheraldeviceslot.DeviceSlot
	val, err := d.Caller().Call(ctx, d.RID(), "getDeviceSlots", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]peripheraldeviceslot.DeviceSlot, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldeviceslot.DeviceSlot
		e0, err = object.As[peripheraldeviceslot.DeviceSlot](a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DeviceManager) GetDeviceSlot(ctx context.Context, in0 int32) (peripheraldeviceslot.DeviceSlot, error) {
	var ret peripheraldeviceslot.DeviceSlot
	val, err := d.Caller().Call(ctx, d.RID(), "getDeviceSlot", map[string]any{
		"idx": in0,
	})
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	ret, err = object.As[peripheraldeviceslot.DeviceSlot](res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) GetSensorHubs(ctx context.Context) ([]sensorhub.SensorHub, error) {
	var ret []sensorhub.SensorHub
	val, err := d.Caller().Call(ctx, d.RID(), "getSensorHubs", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]sensorhub.SensorHub, 0, len(s0))
	for _, a0 := range s0 {
		var e0 sensorhub.SensorHub
		e0, err = object.As[sensorhub.SensorHub](a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DeviceManager) GetDiscoveredDevices(ctx context.Context) ([]peripheraldeviceslot.Device, error) {
	var ret []peripheraldeviceslot.Device
	val, err := d.Caller().Call(ctx, d.RID(), "getDiscoveredDevices", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]peripheraldeviceslot.Device, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldeviceslot.Device
		e0, err = valobj.As[peripheraldeviceslot.Device](a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DeviceManager) GetDiscoveredPackageInfos(ctx context.Context) ([]peripheraldevicepackage.PackageInfo, error) {
	var ret []peripheraldevicepackage.PackageInfo
	val, err := d.Caller().Call(ctx, d.RID(), "getDiscoveredPackageInfos", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]peripheraldevicepackage.PackageInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldevicepackage.PackageInfo
		err = e0.Decode(a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DeviceManager) GetSettings(ctx context.Context) (DeviceManagerSettings, error) {
	var ret DeviceManagerSettings
	val, err := d.Caller().Call(ctx, d.RID(), "getSettings", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) SetSettings(ctx context.Context, in0 DeviceManagerSettings) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
	})
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) GetMetaData(ctx context.Context) (DeviceManagerMetaData, error) {
	var ret DeviceManagerMetaData
	val, err := d.Caller().Call(ctx, d.RID(), "getMetaData", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) GetDeviceTypeInfos(ctx context.Context) ([]DeviceManagerDeviceTypeInfo, error) {
	var ret []DeviceManagerDeviceTypeInfo
	val, err := d.Caller().Call(ctx, d.RID(), "getDeviceTypeInfos", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]DeviceManagerDeviceTypeInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 DeviceManagerDeviceTypeInfo
		err = e0.Decode(a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DeviceManager) GetFirmwareUpdateState(ctx context.Context) (DeviceManagerFirmwareUpdateState, error) {
	var ret DeviceManagerFirmwareUpdateState
	val, err := d.Caller().Call(ctx, d.RID(), "getFirmwareUpdateState", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) GetStatistics(ctx context.Context) (DeviceManagerStatistics, error) {
	var ret DeviceManagerStatistics
	val, err := d.Caller().Call(ctx, d.RID(), "getStatistics", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) GetDiscoveredPackages(ctx context.Context) ([]peripheraldevicepackage.Package, error) {
	var ret []peripheraldevicepackage.Package
	val, err := d.Caller().Call(ctx, d.RID(), "getDiscoveredPackages", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]peripheraldevicepackage.Package, 0, len(s0))
	for _, a0 := range s0 {
		var e0 peripheraldevicepackage.Package
		e0, err = object.As[peripheraldevicepackage.Package](a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DeviceManager) GetPortFuse(ctx context.Context) (portfuse.PortFuse, error) {
	var ret portfuse.PortFuse
	val, err := d.Caller().Call(ctx, d.RID(), "getPortFuse", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	ret, err = object.As[portfuse.PortFuse](res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceManager) GetGatewaySensorManager(ctx context.Context) (gatewaysensormanager.GatewaySensorManager, error) {
	var ret gatewaysensormanager.GatewaySensorManager
	val, err := d.Caller().Call(ctx, d.RID(), "getGatewaySensorManager", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	ret, err = object.As[gatewaysensormanager.GatewaySensorManager](res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
