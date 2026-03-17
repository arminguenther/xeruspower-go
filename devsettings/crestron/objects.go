// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package crestron

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/object"
)

func init() {
	object.Register(NewCrestron)
}

type _Crestron struct {
	idl.Object
}

// NewCrestron returns a new Crestron interface for the object at given RID.
func NewCrestron(rid string, caller idl.Caller) Crestron {
	return &_Crestron{idl.NewObject(rid, caller)}
}

func (c *_Crestron) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Crestron",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_Crestron) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := c.Caller().Call(ctx, c.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Crestron) SetSettings(ctx context.Context, in0 Settings) error {
	_, err := c.Caller().Call(ctx, c.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
	})
	return err
}
