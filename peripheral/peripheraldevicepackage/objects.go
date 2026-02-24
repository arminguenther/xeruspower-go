// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package peripheraldevicepackage

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewBatteryPoweredDevicePackage)
	object.Register(NewDoorHandleControllerPackage)
	object.Register(NewPackage)
}

type _Package struct {
	idl.Object
}

// NewPackage returns a new Package interface for the object at given RID.
func NewPackage(rid string, caller idl.Caller) Package {
	return &_Package{idl.NewObject(rid, caller)}
}

func (p *_Package) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.Package",
		Major: 3, Submajor: 0, Minor: 3,
	}
}

func (p *_Package) GetPackageInfo(ctx context.Context) (PackageInfo, error) {
	var ret PackageInfo
	val, err := p.Caller().Call(ctx, p.RID(), "getPackageInfo", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

type _DoorHandleControllerPackage struct {
	Package
}

// NewDoorHandleControllerPackage returns a new DoorHandleControllerPackage interface for the object at given RID.
func NewDoorHandleControllerPackage(rid string, caller idl.Caller) DoorHandleControllerPackage {
	return &_DoorHandleControllerPackage{NewPackage(rid, caller)}
}

func (d *_DoorHandleControllerPackage) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DoorHandleControllerPackage",
		Major: 3, Submajor: 0, Minor: 3,
	}
}

func (d *_DoorHandleControllerPackage) GetSupportedHandleTypes(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := d.Caller().Call(ctx, d.RID(), "getSupportedHandleTypes", nil)
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
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DoorHandleControllerPackage) GetSupportedExternalDeviceTypes(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := d.Caller().Call(ctx, d.RID(), "getSupportedExternalDeviceTypes", nil)
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
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DoorHandleControllerPackage) GetHandleType(ctx context.Context, in0 int32) (string, error) {
	var ret string
	val, err := d.Caller().Call(ctx, d.RID(), "getHandleType", map[string]any{
		"channel": in0,
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DoorHandleControllerPackage) GetExternalDeviceType(ctx context.Context, in0 int32) (string, error) {
	var ret string
	val, err := d.Caller().Call(ctx, d.RID(), "getExternalDeviceType", map[string]any{
		"channel": in0,
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DoorHandleControllerPackage) SetHandleType(ctx context.Context, in0 int32, in1 string) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setHandleType", map[string]any{
		"channel":    in0,
		"handleType": in1,
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

func (d *_DoorHandleControllerPackage) SetExternalDeviceType(ctx context.Context, in0 int32, in1 string) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setExternalDeviceType", map[string]any{
		"channel": in0,
		"type":    in1,
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

type _BatteryPoweredDevicePackage struct {
	Package
}

// NewBatteryPoweredDevicePackage returns a new BatteryPoweredDevicePackage interface for the object at given RID.
func NewBatteryPoweredDevicePackage(rid string, caller idl.Caller) BatteryPoweredDevicePackage {
	return &_BatteryPoweredDevicePackage{NewPackage(rid, caller)}
}

func (b *_BatteryPoweredDevicePackage) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.BatteryPoweredDevicePackage",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (b *_BatteryPoweredDevicePackage) GetBatteryVoltage(ctx context.Context) (float64, error) {
	var ret float64
	val, err := b.Caller().Call(ctx, b.RID(), "getBatteryVoltage", nil)
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
	ret, err = encoding.AsFloat64(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
