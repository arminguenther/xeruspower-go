// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsamdevice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/dsam/dsamport"
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewDsamDevice)
}

type _DsamDevice struct {
	idl.Object
}

// NewDsamDevice returns a new DsamDevice interface for the object at given RID.
func NewDsamDevice(rid string, caller idl.Caller) DsamDevice {
	return &_DsamDevice{idl.NewObject(rid, caller)}
}

func (d *_DsamDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DsamDevice) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := d.Caller().Call(ctx, d.RID(), "getInfo", nil)
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

func (d *_DsamDevice) GetPorts(ctx context.Context) (map[int32]dsamport.DsamPort, error) {
	var ret map[int32]dsamport.DsamPort
	val, err := d.Caller().Call(ctx, d.RID(), "getPorts", nil)
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
	ret = make(map[int32]dsamport.DsamPort, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 dsamport.DsamPort
		v0, err = object.As[dsamport.DsamPort](b0, d.Caller())
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

func (d *_DsamDevice) StartFirmwareUpdate(ctx context.Context) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "startFirmwareUpdate", nil)
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
