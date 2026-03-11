// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serialport

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
)

func init() {
	object.Register(NewSerialPort)
}

type _SerialPort struct {
	idl.Object
}

// NewSerialPort returns a new SerialPort interface for the object at given RID.
func NewSerialPort(rid string, caller idl.Caller) SerialPort {
	return &_SerialPort{idl.NewObject(rid, caller)}
}

func (s *_SerialPort) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.SerialPort",
		Major: 3, Submajor: 0, Minor: 1,
	}
}

func (s *_SerialPort) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := s.Caller().Call(ctx, s.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_SerialPort) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := s.Caller().Call(ctx, s.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_SerialPort) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setSettings", map[string]any{
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

func (s *_SerialPort) GetState(ctx context.Context) (State, error) {
	var ret State
	val, err := s.Caller().Call(ctx, s.RID(), "getState", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_SerialPort) GetModem(ctx context.Context) (idl.Object, error) {
	var ret idl.Object
	val, err := s.Caller().Call(ctx, s.RID(), "getModem", nil)
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
	ret, err = object.As[idl.Object](res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
