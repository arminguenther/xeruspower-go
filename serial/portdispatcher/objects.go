// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portdispatcher

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40211/serial/serialport"
)

func init() {
	object.Register(NewPortDispatcher)
}

type _PortDispatcher struct {
	idl.Object
}

// NewPortDispatcher returns a new PortDispatcher interface for the object at given RID.
func NewPortDispatcher(rid string, caller idl.Caller) PortDispatcher {
	return &_PortDispatcher{idl.NewObject(rid, caller)}
}

func (p *_PortDispatcher) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.PortDispatcher",
		Major: 1, Submajor: 2, Minor: 2,
	}
}

func (p *_PortDispatcher) GetPorts(ctx context.Context) (map[string]serialport.SerialPort, error) {
	var ret map[string]serialport.SerialPort
	val, err := p.Caller().Call(ctx, p.RID(), "getPorts", nil)
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
	ret = make(map[string]serialport.SerialPort, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 serialport.SerialPort
		v0, err = object.As[serialport.SerialPort](b0, p.Caller())
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
