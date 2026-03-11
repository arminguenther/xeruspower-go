// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package diagnostics

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewDiagnostics)
}

type _Diagnostics struct {
	idl.Object
}

// NewDiagnostics returns a new Diagnostics interface for the object at given RID.
func NewDiagnostics(rid string, caller idl.Caller) Diagnostics {
	return &_Diagnostics{idl.NewObject(rid, caller)}
}

func (d *_Diagnostics) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.Diagnostics",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (d *_Diagnostics) Ping(ctx context.Context, in0 string, in1 int32, in2 int32) (int32, []string, error) {
	var ret int32
	var out0 []string
	val, err := d.Caller().Call(ctx, d.RID(), "ping", map[string]any{
		"hostName": in0,
		"count":    in1,
		"size":     in2,
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
	err = encoding.In("results", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["results"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (d *_Diagnostics) TraceRoute(ctx context.Context, in0 string, in1 int32, in2 bool) (int32, []string, error) {
	var ret int32
	var out0 []string
	val, err := d.Caller().Call(ctx, d.RID(), "traceRoute", map[string]any{
		"hostName": in0,
		"timeout":  in1,
		"useIcmp":  in2,
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
	err = encoding.In("results", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["results"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (d *_Diagnostics) ListTcpConnections(ctx context.Context) (int32, []string, error) {
	var ret int32
	var out0 []string
	val, err := d.Caller().Call(ctx, d.RID(), "listTcpConnections", nil)
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
	err = encoding.In("results", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["results"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (d *_Diagnostics) ListTcpUdpListenSockets(ctx context.Context) (int32, []string, error) {
	var ret int32
	var out0 []string
	val, err := d.Caller().Call(ctx, d.RID(), "listTcpUdpListenSockets", nil)
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
	err = encoding.In("results", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["results"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (d *_Diagnostics) ResolveHostName(ctx context.Context, in0 string) ([]string, error) {
	var out0 []string
	val, err := d.Caller().Call(ctx, d.RID(), "resolveHostName", map[string]any{
		"hostName": in0,
	})
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("results", res)
	if err != nil {
		return out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["results"])
	if err != nil {
		return out0, err
	}
	out0 = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e0)
	}
	return out0, nil
}

func (d *_Diagnostics) FlushRouteCache(ctx context.Context, in0 string) error {
	_, err := d.Caller().Call(ctx, d.RID(), "flushRouteCache", map[string]any{
		"ifName": in0,
	})
	return err
}
