// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package redfish

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
)

func init() {
	object.Register(NewRedfish)
}

type _Redfish struct {
	idl.Object
}

// NewRedfish returns a new Redfish interface for the object at given RID.
func NewRedfish(rid string, caller idl.Caller) Redfish {
	return &_Redfish{idl.NewObject(rid, caller)}
}

func (r *_Redfish) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Redfish",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_Redfish) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := r.Caller().Call(ctx, r.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], r.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (r *_Redfish) SetSettings(ctx context.Context, in0 Settings) error {
	_, err := r.Caller().Call(ctx, r.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
	})
	return err
}
