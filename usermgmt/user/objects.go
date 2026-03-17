// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package user

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40010/usermgmt/role"
)

func init() {
	object.Register(NewUser)
}

type _User struct {
	idl.Object
}

// NewUser returns a new User interface for the object at given RID.
func NewUser(rid string, caller idl.Caller) User {
	return &_User{idl.NewObject(rid, caller)}
}

func (u *_User) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.User",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (u *_User) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := u.Caller().Call(ctx, u.RID(), "getInfo", nil)
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

func (u *_User) SetAccountPassword(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "setAccountPassword", map[string]any{
		"password": in0,
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

func (u *_User) UpdateAccountFull(ctx context.Context, in0 string, in1 Info) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "updateAccountFull", map[string]any{
		"password": in0,
		"info":     in1.Encode(),
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

func (u *_User) GetInfoAndPrivileges(ctx context.Context) (Info, []role.Privilege, error) {
	var out0 Info
	var out1 []role.Privilege
	val, err := u.Caller().Call(ctx, u.RID(), "getInfoAndPrivileges", nil)
	if err != nil {
		return out0, out1, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("info", res)
	if err != nil {
		return out0, out1, err
	}
	err = out0.Decode(res["info"], u.Caller())
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("privileges", res)
	if err != nil {
		return out0, out1, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["privileges"])
	if err != nil {
		return out0, out1, err
	}
	out1 = make([]role.Privilege, 0, len(s0))
	for _, a0 := range s0 {
		var e0 role.Privilege
		err = e0.Decode(a0, u.Caller())
		if err != nil {
			return out0, out1, err
		}
		out1 = append(out1, e0)
	}
	return out0, out1, nil
}

func (u *_User) SetPreferences(ctx context.Context, in0 Preferences) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "setPreferences", map[string]any{
		"prefs": in0.Encode(),
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

func (u *_User) GetCapabilities(ctx context.Context) (Capabilities, error) {
	var ret Capabilities
	val, err := u.Caller().Call(ctx, u.RID(), "getCapabilities", nil)
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
