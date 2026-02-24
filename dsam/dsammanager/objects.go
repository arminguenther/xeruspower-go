// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package dsammanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/dsam/dsamdevice"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewDsamManager)
}

type _DsamManager struct {
	idl.Object
}

// NewDsamManager returns a new DsamManager interface for the object at given RID.
func NewDsamManager(rid string, caller idl.Caller) DsamManager {
	return &_DsamManager{idl.NewObject(rid, caller)}
}

func (d *_DsamManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamManager",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DsamManager) GetDsamDevices(ctx context.Context) (map[int32]dsamdevice.DsamDevice, error) {
	var ret map[int32]dsamdevice.DsamDevice
	val, err := d.Caller().Call(ctx, d.RID(), "getDsamDevices", nil)
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
	i0, e0, l0 := encoding.AsMapItems(res["_ret_"])
	ret = make(map[int32]dsamdevice.DsamDevice, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 dsamdevice.DsamDevice
		v0, err = object.As[dsamdevice.DsamDevice](b0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret[k0] = v0
	}
	err = e0()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DsamManager) GetFirmwareUpdateFileVersion(ctx context.Context) (dsamdevice.FirmwareVersion, error) {
	var ret dsamdevice.FirmwareVersion
	val, err := d.Caller().Call(ctx, d.RID(), "getFirmwareUpdateFileVersion", nil)
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
