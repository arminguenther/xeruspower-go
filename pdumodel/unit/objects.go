// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package unit

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
)

func init() {
	object.Register(NewUnit)
}

type _Unit struct {
	idl.Object
}

// NewUnit returns a new Unit interface for the object at given RID.
func NewUnit(rid string, caller idl.Caller) Unit {
	return &_Unit{idl.NewObject(rid, caller)}
}

func (u *_Unit) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Unit",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (u *_Unit) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := u.Caller().Call(ctx, u.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], u.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (u *_Unit) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := u.Caller().Call(ctx, u.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], u.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (u *_Unit) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "setSettings", map[string]any{
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

func (u *_Unit) Identify(ctx context.Context, in0 int32) error {
	_, err := u.Caller().Call(ctx, u.RID(), "identify", map[string]any{
		"seconds": in0,
	})
	return err
}

func (u *_Unit) MuteBuzzer(ctx context.Context, in0 bool) error {
	_, err := u.Caller().Call(ctx, u.RID(), "muteBuzzer", map[string]any{
		"mute": in0,
	})
	return err
}

func (u *_Unit) GetDisplayOrientation(ctx context.Context) (Orientation, error) {
	var ret Orientation
	val, err := u.Caller().Call(ctx, u.RID(), "getDisplayOrientation", nil)
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
	ret, err = encoding.AsEnum[Orientation](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
