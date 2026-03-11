// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package net

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewNet)
}

type _Net struct {
	idl.Object
}

// NewNet returns a new Net interface for the object at given RID.
func NewNet(rid string, caller idl.Caller) Net {
	return &_Net{idl.NewObject(rid, caller)}
}

func (n *_Net) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.Net",
		Major: 8, Submajor: 0, Minor: 0,
	}
}

func (n *_Net) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := n.Caller().Call(ctx, n.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_Net) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := n.Caller().Call(ctx, n.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (n *_Net) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := n.Caller().Call(ctx, n.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
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

func (n *_Net) GetPortForwardingProtocolMappings(ctx context.Context) ([]PortForwardingProtocolMapping, error) {
	var ret []PortForwardingProtocolMapping
	val, err := n.Caller().Call(ctx, n.RID(), "getPortForwardingProtocolMappings", nil)
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
	ret = make([]PortForwardingProtocolMapping, 0, len(s0))
	for _, a0 := range s0 {
		var e0 PortForwardingProtocolMapping
		err = e0.Decode(a0, n.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
