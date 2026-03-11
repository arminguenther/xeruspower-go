// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeemanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40411/zigbee/zigbeedevice"
)

func init() {
	object.Register(NewZigbeeManager)
}

type _ZigbeeManager struct {
	idl.Object
}

// NewZigbeeManager returns a new ZigbeeManager interface for the object at given RID.
func NewZigbeeManager(rid string, caller idl.Caller) ZigbeeManager {
	return &_ZigbeeManager{idl.NewObject(rid, caller)}
}

func (z *_ZigbeeManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeManager",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (z *_ZigbeeManager) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := z.Caller().Call(ctx, z.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], z.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (z *_ZigbeeManager) GetDongleState(ctx context.Context) (DongleState, error) {
	var ret DongleState
	val, err := z.Caller().Call(ctx, z.RID(), "getDongleState", nil)
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
	ret, err = encoding.AsEnum[DongleState](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (z *_ZigbeeManager) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := z.Caller().Call(ctx, z.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], z.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (z *_ZigbeeManager) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := z.Caller().Call(ctx, z.RID(), "setSettings", map[string]any{
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

func (z *_ZigbeeManager) AddDevices(ctx context.Context, in0 []DeviceRegistration) ([]int32, error) {
	var ret []int32
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := z.Caller().Call(ctx, z.RID(), "addDevices", map[string]any{
		"sensors": s0,
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
	var s1 []any
	s1, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e1)
	}
	return ret, nil
}

func (z *_ZigbeeManager) RemoveDevice(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := z.Caller().Call(ctx, z.RID(), "removeDevice", map[string]any{
		"sourceId": in0,
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

func (z *_ZigbeeManager) GetRegisteredDevices(ctx context.Context) ([]zigbeedevice.ZigbeeDevice, error) {
	var ret []zigbeedevice.ZigbeeDevice
	val, err := z.Caller().Call(ctx, z.RID(), "getRegisteredDevices", nil)
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
	ret = make([]zigbeedevice.ZigbeeDevice, 0, len(s0))
	for _, a0 := range s0 {
		var e0 zigbeedevice.ZigbeeDevice
		e0, err = object.As[zigbeedevice.ZigbeeDevice](a0, z.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
