// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package production

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
)

func init() {
	object.Register(NewProduction)
}

type _Production struct {
	idl.Object
}

// NewProduction returns a new Production interface for the object at given RID.
func NewProduction(rid string, caller idl.Caller) Production {
	return &_Production{idl.NewObject(rid, caller)}
}

func (p *_Production) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "production.Production",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_Production) EnterFactoryConfigMode(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "enterFactoryConfigMode", map[string]any{
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

func (p *_Production) LeaveFactoryConfigMode(ctx context.Context) error {
	_, err := p.Caller().Call(ctx, p.RID(), "leaveFactoryConfigMode", nil)
	return err
}

func (p *_Production) IsFactoryConfigModeEnabled(ctx context.Context) (bool, error) {
	var ret bool
	val, err := p.Caller().Call(ctx, p.RID(), "isFactoryConfigModeEnabled", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
