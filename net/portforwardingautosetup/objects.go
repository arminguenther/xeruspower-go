// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portforwardingautosetup

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewPortForwardingAutoSetup)
}

type _PortForwardingAutoSetup struct {
	idl.Object
}

// NewPortForwardingAutoSetup returns a new PortForwardingAutoSetup interface for the object at given RID.
func NewPortForwardingAutoSetup(rid string, caller idl.Caller) PortForwardingAutoSetup {
	return &_PortForwardingAutoSetup{idl.NewObject(rid, caller)}
}

func (p *_PortForwardingAutoSetup) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.PortForwardingAutoSetup",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PortForwardingAutoSetup) Start(ctx context.Context, in0 int32, in1 string, in2 string, in3 string, in4 bool) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "start", map[string]any{
		"numberOfExpansionUnits":   in0,
		"username":                 in1,
		"password":                 in2,
		"newPassword":              in3,
		"disableStrongPasswordReq": in4,
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

func (p *_PortForwardingAutoSetup) GetStatus(ctx context.Context) (Status, error) {
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_PortForwardingAutoSetup) Cancel(ctx context.Context) error {
	_, err := p.Caller().Call(ctx, p.RID(), "cancel", nil)
	return err
}
