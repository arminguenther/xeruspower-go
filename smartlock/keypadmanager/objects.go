// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package keypadmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40020/smartlock/keypad"
)

func init() {
	object.Register(NewKeypadManager)
}

type _KeypadManager struct {
	idl.Object
}

// NewKeypadManager returns a new KeypadManager interface for the object at given RID.
func NewKeypadManager(rid string, caller idl.Caller) KeypadManager {
	return &_KeypadManager{idl.NewObject(rid, caller)}
}

func (k *_KeypadManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.KeypadManager",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (k *_KeypadManager) GetKeypads(ctx context.Context) ([]keypad.Keypad, error) {
	var ret []keypad.Keypad
	val, err := k.Caller().Call(ctx, k.RID(), "getKeypads", nil)
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
	ret = make([]keypad.Keypad, 0, len(s0))
	for _, a0 := range s0 {
		var e0 keypad.Keypad
		e0, err = object.As[keypad.Keypad](a0, k.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (k *_KeypadManager) GetKeypadById(ctx context.Context, in0 string) (keypad.Keypad, error) {
	var ret keypad.Keypad
	val, err := k.Caller().Call(ctx, k.RID(), "getKeypadById", map[string]any{
		"keypadId": in0,
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
	ret, err = object.As[keypad.Keypad](res["_ret_"], k.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (k *_KeypadManager) SetKeypadSettings(ctx context.Context, in0 string, in1 KeypadSettings) (int32, error) {
	var ret int32
	val, err := k.Caller().Call(ctx, k.RID(), "setKeypadSettings", map[string]any{
		"position": in0,
		"setting":  in1.Encode(),
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

func (k *_KeypadManager) GetAllKeypadSettings(ctx context.Context) (map[string]KeypadSettings, error) {
	var ret map[string]KeypadSettings
	val, err := k.Caller().Call(ctx, k.RID(), "getAllKeypadSettings", nil)
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
	ret = make(map[string]KeypadSettings, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 KeypadSettings
		err = v0.Decode(b0, k.Caller())
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
