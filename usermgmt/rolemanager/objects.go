// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package rolemanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/usermgmt/role"
)

func init() {
	object.Register(NewRoleManager)
}

type _RoleManager struct {
	idl.Object
}

// NewRoleManager returns a new RoleManager interface for the object at given RID.
func NewRoleManager(rid string, caller idl.Caller) RoleManager {
	return &_RoleManager{idl.NewObject(rid, caller)}
}

func (r *_RoleManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.RoleManager",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RoleManager) CreateRoleFull(ctx context.Context, in0 string, in1 role.Info) (int32, error) {
	var ret int32
	val, err := r.Caller().Call(ctx, r.RID(), "createRoleFull", map[string]any{
		"name": in0,
		"info": in1.Encode(),
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

func (r *_RoleManager) DeleteRole(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := r.Caller().Call(ctx, r.RID(), "deleteRole", map[string]any{
		"name": in0,
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

func (r *_RoleManager) GetAllRoleNames(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := r.Caller().Call(ctx, r.RID(), "getAllRoleNames", nil)
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

func (r *_RoleManager) GetAllRoles(ctx context.Context) ([]RoleAccount, error) {
	var ret []RoleAccount
	val, err := r.Caller().Call(ctx, r.RID(), "getAllRoles", nil)
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
	ret = make([]RoleAccount, 0, len(s0))
	for _, a0 := range s0 {
		var e0 RoleAccount
		err = e0.Decode(a0, r.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (r *_RoleManager) GetAllPrivileges(ctx context.Context) ([]PrivilegeDesc, error) {
	var ret []PrivilegeDesc
	val, err := r.Caller().Call(ctx, r.RID(), "getAllPrivileges", nil)
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
	ret = make([]PrivilegeDesc, 0, len(s0))
	for _, a0 := range s0 {
		var e0 PrivilegeDesc
		err = e0.Decode(a0, r.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (r *_RoleManager) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := r.Caller().Call(ctx, r.RID(), "getInfo", nil)
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
