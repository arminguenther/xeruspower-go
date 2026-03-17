// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package keypad

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
)

func init() {
	object.Register(NewKeypad)
}

type _Keypad struct {
	idl.Object
}

// NewKeypad returns a new Keypad interface for the object at given RID.
func NewKeypad(rid string, caller idl.Caller) Keypad {
	return &_Keypad{idl.NewObject(rid, caller)}
}

func (k *_Keypad) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.Keypad",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (k *_Keypad) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := k.Caller().Call(ctx, k.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], k.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (k *_Keypad) GetPIN(ctx context.Context) (int32, string, error) {
	var ret int32
	var out0 string
	val, err := k.Caller().Call(ctx, k.RID(), "getPIN", nil)
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("pin", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["pin"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}
