// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package ldapmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
)

func init() {
	object.Register(NewLdapManager)
}

type _LdapManager struct {
	idl.Object
}

// NewLdapManager returns a new LdapManager interface for the object at given RID.
func NewLdapManager(rid string, caller idl.Caller) LdapManager {
	return &_LdapManager{idl.NewObject(rid, caller)}
}

func (l *_LdapManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "auth.LdapManager",
		Major: 3, Submajor: 0, Minor: 1,
	}
}

func (l *_LdapManager) GetLdapServers(ctx context.Context) ([]ServerSettings, error) {
	var ret []ServerSettings
	val, err := l.Caller().Call(ctx, l.RID(), "getLdapServers", nil)
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
	ret = make([]ServerSettings, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ServerSettings
		err = e0.Decode(a0, l.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (l *_LdapManager) SetLdapServers(ctx context.Context, in0 []ServerSettings) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := l.Caller().Call(ctx, l.RID(), "setLdapServers", map[string]any{
		"serverList": s0,
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

func (l *_LdapManager) TestLdapServer(ctx context.Context, in0 string, in1 string, in2 ServerSettings) (int32, string, error) {
	var ret int32
	var out0 string
	val, err := l.Caller().Call(ctx, l.RID(), "testLdapServer", map[string]any{
		"username": in0,
		"password": in1,
		"settings": in2.Encode(),
	})
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
	err = encoding.In("diagMsg", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["diagMsg"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}
