// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bluetooth

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
)

func init() {
	object.Register(NewBluetooth)
}

type _Bluetooth struct {
	idl.Object
}

// NewBluetooth returns a new Bluetooth interface for the object at given RID.
func NewBluetooth(rid string, caller idl.Caller) Bluetooth {
	return &_Bluetooth{idl.NewObject(rid, caller)}
}

func (b *_Bluetooth) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Bluetooth",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (b *_Bluetooth) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := b.Caller().Call(ctx, b.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], b.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (b *_Bluetooth) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := b.Caller().Call(ctx, b.RID(), "setSettings", map[string]any{
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
