// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package modbus

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
)

func init() {
	object.Register(NewModbus)
}

type _Modbus struct {
	idl.Object
}

// NewModbus returns a new Modbus interface for the object at given RID.
func NewModbus(rid string, caller idl.Caller) Modbus {
	return &_Modbus{idl.NewObject(rid, caller)}
}

func (m *_Modbus) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Modbus",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (m *_Modbus) GetCapabilities(ctx context.Context) (Capabilities, error) {
	var ret Capabilities
	val, err := m.Caller().Call(ctx, m.RID(), "getCapabilities", nil)
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
	err = ret.Decode(res["_ret_"], m.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (m *_Modbus) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := m.Caller().Call(ctx, m.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], m.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (m *_Modbus) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := m.Caller().Call(ctx, m.RID(), "setSettings", map[string]any{
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
