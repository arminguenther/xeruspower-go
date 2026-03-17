// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package snmp

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
)

func init() {
	object.Register(NewSnmp)
}

type _Snmp struct {
	idl.Object
}

// NewSnmp returns a new Snmp interface for the object at given RID.
func NewSnmp(rid string, caller idl.Caller) Snmp {
	return &_Snmp{idl.NewObject(rid, caller)}
}

func (s *_Snmp) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Snmp",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (s *_Snmp) GetConfiguration(ctx context.Context) (Configuration, error) {
	var ret Configuration
	val, err := s.Caller().Call(ctx, s.RID(), "getConfiguration", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Snmp) SetConfiguration(ctx context.Context, in0 Configuration) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setConfiguration", map[string]any{
		"cfg": in0.Encode(),
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

func (s *_Snmp) GetV3EngineId(ctx context.Context) (string, error) {
	var ret string
	val, err := s.Caller().Call(ctx, s.RID(), "getV3EngineId", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
