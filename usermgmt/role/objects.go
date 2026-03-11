// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package role

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewRole)
}

type _Role struct {
	idl.Object
}

// NewRole returns a new Role interface for the object at given RID.
func NewRole(rid string, caller idl.Caller) Role {
	return &_Role{idl.NewObject(rid, caller)}
}

func (r *_Role) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.Role",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_Role) GetInfo(ctx context.Context) (Info, error) {
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

func (r *_Role) UpdateFull(ctx context.Context, in0 Info) (int32, error) {
	var ret int32
	val, err := r.Caller().Call(ctx, r.RID(), "updateFull", map[string]any{
		"info": in0.Encode(),
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
