// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zeroconf

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewZeroconf)
}

type _Zeroconf struct {
	idl.Object
}

// NewZeroconf returns a new Zeroconf interface for the object at given RID.
func NewZeroconf(rid string, caller idl.Caller) Zeroconf {
	return &_Zeroconf{idl.NewObject(rid, caller)}
}

func (z *_Zeroconf) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Zeroconf",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (z *_Zeroconf) GetSettings(ctx context.Context) (Settings, error) {
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

func (z *_Zeroconf) SetSettings(ctx context.Context, in0 Settings) error {
	_, err := z.Caller().Call(ctx, z.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
	})
	return err
}
