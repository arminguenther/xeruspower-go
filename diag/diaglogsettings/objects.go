// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package diaglogsettings

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
)

func init() {
	object.Register(NewDiagLogSettings)
}

type _DiagLogSettings struct {
	idl.Object
}

// NewDiagLogSettings returns a new DiagLogSettings interface for the object at given RID.
func NewDiagLogSettings(rid string, caller idl.Caller) DiagLogSettings {
	return &_DiagLogSettings{idl.NewObject(rid, caller)}
}

func (d *_DiagLogSettings) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "diag.DiagLogSettings",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DiagLogSettings) ResetLogLevelsForAllCtxNames(ctx context.Context) error {
	_, err := d.Caller().Call(ctx, d.RID(), "resetLogLevelsForAllCtxNames", nil)
	return err
}

func (d *_DiagLogSettings) GetLogLevelsForAllCtxNames(ctx context.Context) ([]LogLevelEntry, error) {
	var ret []LogLevelEntry
	val, err := d.Caller().Call(ctx, d.RID(), "getLogLevelsForAllCtxNames", nil)
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
	ret = make([]LogLevelEntry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 LogLevelEntry
		err = e0.Decode(a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DiagLogSettings) SetLogLevelForAllCtxNames(ctx context.Context, in0 LogLevel) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setLogLevelForAllCtxNames", map[string]any{
		"logLevel": in0,
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

func (d *_DiagLogSettings) GetLogLevelByCtxName(ctx context.Context, in0 string) (int32, LogLevel, error) {
	var ret int32
	var out0 LogLevel
	val, err := d.Caller().Call(ctx, d.RID(), "getLogLevelByCtxName", map[string]any{
		"ctxName": in0,
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
	err = encoding.In("logLevel", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.AsEnum[LogLevel](res["logLevel"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (d *_DiagLogSettings) SetLogLevelByCtxName(ctx context.Context, in0 string, in1 LogLevel) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setLogLevelByCtxName", map[string]any{
		"ctxName":  in0,
		"logLevel": in1,
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
