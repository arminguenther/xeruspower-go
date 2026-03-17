// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package authmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
)

func init() {
	object.Register(NewAuthManager)
}

type _AuthManager struct {
	idl.Object
}

// NewAuthManager returns a new AuthManager interface for the object at given RID.
func NewAuthManager(rid string, caller idl.Caller) AuthManager {
	return &_AuthManager{idl.NewObject(rid, caller)}
}

func (a *_AuthManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "auth.AuthManager",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AuthManager) GetPolicy(ctx context.Context) (Policy, error) {
	var ret Policy
	val, err := a.Caller().Call(ctx, a.RID(), "getPolicy", nil)
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
	err = ret.Decode(res["_ret_"], a.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AuthManager) SetPolicy(ctx context.Context, in0 Policy) (int32, error) {
	var ret int32
	val, err := a.Caller().Call(ctx, a.RID(), "setPolicy", map[string]any{
		"p": in0.Encode(),
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
