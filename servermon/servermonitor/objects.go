// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package servermonitor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
)

func init() {
	object.Register(NewServerMonitor)
}

type _ServerMonitor struct {
	idl.Object
}

// NewServerMonitor returns a new ServerMonitor interface for the object at given RID.
func NewServerMonitor(rid string, caller idl.Caller) ServerMonitor {
	return &_ServerMonitor{idl.NewObject(rid, caller)}
}

func (s *_ServerMonitor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "servermon.ServerMonitor",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (s *_ServerMonitor) AddServer(ctx context.Context, in0 ServerSettings) (int32, int32, error) {
	var ret int32
	var out0 int32
	val, err := s.Caller().Call(ctx, s.RID(), "addServer", map[string]any{
		"settings": in0.Encode(),
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
	err = encoding.In("id", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.AsInt32(res["id"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (s *_ServerMonitor) ModifyServer(ctx context.Context, in0 int32, in1 ServerSettings) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "modifyServer", map[string]any{
		"id":       in0,
		"settings": in1.Encode(),
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

func (s *_ServerMonitor) DeleteServer(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "deleteServer", map[string]any{
		"id": in0,
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

func (s *_ServerMonitor) GetServer(ctx context.Context, in0 int32) (int32, Server, error) {
	var ret int32
	var out0 Server
	val, err := s.Caller().Call(ctx, s.RID(), "getServer", map[string]any{
		"id": in0,
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
	err = encoding.In("server", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["server"], s.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (s *_ServerMonitor) ListServers(ctx context.Context) (map[int32]Server, error) {
	var ret map[int32]Server
	val, err := s.Caller().Call(ctx, s.RID(), "listServers", nil)
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
	i0, e0, l0 := encoding.AsMapItems(res["_ret_"])
	ret = make(map[int32]Server, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 Server
		err = v0.Decode(b0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret[k0] = v0
	}
	err = e0()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ServerMonitor) PowerControl(ctx context.Context, in0 int32, in1 bool) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "powerControl", map[string]any{
		"id": in0,
		"on": in1,
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
