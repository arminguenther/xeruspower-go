// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portfuse

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
)

func init() {
	object.Register(NewPortFuse)
}

type _PortFuse struct {
	idl.Object
}

// NewPortFuse returns a new PortFuse interface for the object at given RID.
func NewPortFuse(rid string, caller idl.Caller) PortFuse {
	return &_PortFuse{idl.NewObject(rid, caller)}
}

func (p *_PortFuse) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "portsmodel.PortFuse",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (p *_PortFuse) GetStatus(ctx context.Context) (Status, error) {
	var ret Status
	val, err := p.Caller().Call(ctx, p.RID(), "getStatus", nil)
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
	ret, err = encoding.AsEnum[Status](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_PortFuse) Reset(ctx context.Context) error {
	_, err := p.Caller().Call(ctx, p.RID(), "reset", nil)
	return err
}

func (p *_PortFuse) GetTripCount(ctx context.Context) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "getTripCount", nil)
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
