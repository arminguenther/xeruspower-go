// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package usb

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
)

func init() {
	object.Register(NewUsb)
}

type _Usb struct {
	idl.Object
}

// NewUsb returns a new Usb interface for the object at given RID.
func NewUsb(rid string, caller idl.Caller) Usb {
	return &_Usb{idl.NewObject(rid, caller)}
}

func (u *_Usb) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usb.Usb",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (u *_Usb) GetSettings(ctx context.Context) (Settings, error) {
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

func (u *_Usb) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
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

func (u *_Usb) GetDevices(ctx context.Context) ([]Device, error) {
	var out0 []Device
	val, err := u.Caller().Call(ctx, u.RID(), "getDevices", nil)
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("usbDevices", res)
	if err != nil {
		return out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["usbDevices"])
	if err != nil {
		return out0, err
	}
	out0 = make([]Device, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Device
		err = e0.Decode(a0, u.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e0)
	}
	return out0, nil
}
