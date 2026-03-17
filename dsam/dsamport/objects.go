// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsamport

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
)

func init() {
	object.Register(NewDsamPort)
}

type _DsamPort struct {
	idl.Object
}

// NewDsamPort returns a new DsamPort interface for the object at given RID.
func NewDsamPort(rid string, caller idl.Caller) DsamPort {
	return &_DsamPort{idl.NewObject(rid, caller)}
}

func (d *_DsamPort) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamPort",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DsamPort) GetInfo(ctx context.Context) (Info, error) {
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

func (d *_DsamPort) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
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

func (d *_DsamPort) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
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

func (d *_DsamPort) GetTtyUsbNumber(ctx context.Context) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "getTtyUsbNumber", nil)
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

func (d *_DsamPort) SetState(ctx context.Context, in0 State) error {
	_, err := d.Caller().Call(ctx, d.RID(), "setState", map[string]any{
		"state": in0,
	})
	return err
}
