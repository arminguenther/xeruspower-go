// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package peripheraldeviceslot

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	object.Register(NewDeviceSlot)
}

type _DeviceSlot struct {
	idl.Object
}

// NewDeviceSlot returns a new DeviceSlot interface for the object at given RID.
func NewDeviceSlot(rid string, caller idl.Caller) DeviceSlot {
	return &_DeviceSlot{idl.NewObject(rid, caller)}
}

func (d *_DeviceSlot) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.DeviceSlot",
		Major: 5, Submajor: 0, Minor: 2,
	}
}

func (d *_DeviceSlot) GetDevice(ctx context.Context) (Device, error) {
	var ret Device
	val, err := d.Caller().Call(ctx, d.RID(), "getDevice", nil)
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
	ret, err = valobj.As[Device](res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DeviceSlot) Assign(ctx context.Context, in0 DeviceID) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "assign", map[string]any{
		"devid": in0.Encode(),
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

func (d *_DeviceSlot) AssignAddress(ctx context.Context, in0 string, in1 Address) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "assignAddress", map[string]any{
		"packageClass": in0,
		"address":      in1.Encode(),
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

func (d *_DeviceSlot) Unassign(ctx context.Context) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "unassign", nil)
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

func (d *_DeviceSlot) GetSettings(ctx context.Context) (DeviceSlotSettings, error) {
	var ret DeviceSlotSettings
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

func (d *_DeviceSlot) SetSettings(ctx context.Context, in0 DeviceSlotSettings) (int32, error) {
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
