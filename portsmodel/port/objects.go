// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package port

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40412/portsmodel/portfuse"
)

func init() {
	object.Register(NewPort)
}

type _Port struct {
	idl.Object
}

// NewPort returns a new Port interface for the object at given RID.
func NewPort(rid string, caller idl.Caller) Port {
	return &_Port{idl.NewObject(rid, caller)}
}

func (p *_Port) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "portsmodel.Port",
		Major: 2, Submajor: 0, Minor: 4,
	}
}

func (p *_Port) GetProperties(ctx context.Context) (Properties, error) {
	var ret Properties
	val, err := p.Caller().Call(ctx, p.RID(), "getProperties", nil)
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

func (p *_Port) SetName(ctx context.Context, in0 string) error {
	_, err := p.Caller().Call(ctx, p.RID(), "setName", map[string]any{
		"name": in0,
	})
	return err
}

func (p *_Port) SetDetectionMode(ctx context.Context, in0 DetectionMode) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setDetectionMode", map[string]any{
		"mode": in0.Encode(),
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

func (p *_Port) GetDetectableDevices(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := p.Caller().Call(ctx, p.RID(), "getDetectableDevices", nil)
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

func (p *_Port) GetDevice(ctx context.Context) (idl.Object, error) {
	var ret idl.Object
	val, err := p.Caller().Call(ctx, p.RID(), "getDevice", nil)
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
	ret, err = object.As[idl.Object](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Port) GetDeviceConfig(ctx context.Context, in0 string) (idl.Object, error) {
	var ret idl.Object
	val, err := p.Caller().Call(ctx, p.RID(), "getDeviceConfig", map[string]any{
		"deviceType": in0,
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
	ret, err = object.As[idl.Object](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Port) GetFuse(ctx context.Context) (portfuse.PortFuse, error) {
	var ret portfuse.PortFuse
	val, err := p.Caller().Call(ctx, p.RID(), "getFuse", nil)
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
	ret, err = object.As[portfuse.PortFuse](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
