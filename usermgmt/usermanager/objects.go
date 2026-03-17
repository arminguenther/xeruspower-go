// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package usermanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40220/usermgmt/user"
)

func init() {
	object.Register(NewUserManager)
}

type _UserManager struct {
	idl.Object
}

// NewUserManager returns a new UserManager interface for the object at given RID.
func NewUserManager(rid string, caller idl.Caller) UserManager {
	return &_UserManager{idl.NewObject(rid, caller)}
}

func (u *_UserManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.UserManager",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (u *_UserManager) GetAccountNames(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := u.Caller().Call(ctx, u.RID(), "getAccountNames", nil)
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
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (u *_UserManager) CreateAccount(ctx context.Context, in0 string, in1 string) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "createAccount", map[string]any{
		"username": in0,
		"password": in1,
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

func (u *_UserManager) RenameAccount(ctx context.Context, in0 string, in1 string) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "renameAccount", map[string]any{
		"username":    in0,
		"newUsername": in1,
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

func (u *_UserManager) DeleteAccount(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "deleteAccount", map[string]any{
		"username": in0,
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

func (u *_UserManager) GetAllAccounts(ctx context.Context) ([]Account, error) {
	var ret []Account
	val, err := u.Caller().Call(ctx, u.RID(), "getAllAccounts", nil)
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
	ret = make([]Account, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Account
		err = e0.Decode(a0, u.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (u *_UserManager) CreateAccountFull(ctx context.Context, in0 string, in1 string, in2 user.Info) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "createAccountFull", map[string]any{
		"username": in0,
		"password": in1,
		"info":     in2.Encode(),
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

func (u *_UserManager) GetAccountsByRole(ctx context.Context, in0 string) ([]Account, error) {
	var ret []Account
	val, err := u.Caller().Call(ctx, u.RID(), "getAccountsByRole", map[string]any{
		"roleName": in0,
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
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]Account, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Account
		err = e0.Decode(a0, u.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (u *_UserManager) GetDefaultPreferences(ctx context.Context) (user.Preferences, error) {
	var ret user.Preferences
	val, err := u.Caller().Call(ctx, u.RID(), "getDefaultPreferences", nil)
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

func (u *_UserManager) SetDefaultPreferences(ctx context.Context, in0 user.Preferences) (int32, error) {
	var ret int32
	val, err := u.Caller().Call(ctx, u.RID(), "setDefaultPreferences", map[string]any{
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
