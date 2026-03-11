// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datetime

import (
	"context"
	"time"

	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
)

func init() {
	object.Register(NewDateTime)
}

type _DateTime struct {
	idl.Object
}

// NewDateTime returns a new DateTime interface for the object at given RID.
func NewDateTime(rid string, caller idl.Caller) DateTime {
	return &_DateTime{idl.NewObject(rid, caller)}
}

func (d *_DateTime) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "datetime.DateTime",
		Major: 3, Submajor: 0, Minor: 3,
	}
}

func (d *_DateTime) GetZoneInfos(ctx context.Context, in0 bool) ([]ZoneInfo, error) {
	var out0 []ZoneInfo
	val, err := d.Caller().Call(ctx, d.RID(), "getZoneInfos", map[string]any{
		"useOlson": in0,
	})
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("zoneInfos", res)
	if err != nil {
		return out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["zoneInfos"])
	if err != nil {
		return out0, err
	}
	out0 = make([]ZoneInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ZoneInfo
		err = e0.Decode(a0, d.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e0)
	}
	return out0, nil
}

func (d *_DateTime) CheckNtpServer(ctx context.Context, in0 string) (bool, error) {
	var ret bool
	val, err := d.Caller().Call(ctx, d.RID(), "checkNtpServer", map[string]any{
		"ntpServer": in0,
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_DateTime) GetActiveNtpServers(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := d.Caller().Call(ctx, d.RID(), "getActiveNtpServers", nil)
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

func (d *_DateTime) GetCfg(ctx context.Context) (Cfg, error) {
	var out0 Cfg
	val, err := d.Caller().Call(ctx, d.RID(), "getCfg", nil)
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("cfg", res)
	if err != nil {
		return out0, err
	}
	err = out0.Decode(res["cfg"], d.Caller())
	if err != nil {
		return out0, err
	}
	return out0, nil
}

func (d *_DateTime) SetCfg(ctx context.Context, in0 Cfg) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setCfg", map[string]any{
		"cfg": in0.Encode(),
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

func (d *_DateTime) GetTime(ctx context.Context, in0 bool) (ZoneInfo, bool, int32, time.Time, error) {
	var out0 ZoneInfo
	var out1 bool
	var out2 int32
	var out3 time.Time
	val, err := d.Caller().Call(ctx, d.RID(), "getTime", map[string]any{
		"useOlson": in0,
	})
	if err != nil {
		return out0, out1, out2, out3, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, out1, out2, out3, err
	}
	err = encoding.In("zone", res)
	if err != nil {
		return out0, out1, out2, out3, err
	}
	err = out0.Decode(res["zone"], d.Caller())
	if err != nil {
		return out0, out1, out2, out3, err
	}
	err = encoding.In("dstEnabled", res)
	if err != nil {
		return out0, out1, out2, out3, err
	}
	out1, err = encoding.Is[bool](res["dstEnabled"])
	if err != nil {
		return out0, out1, out2, out3, err
	}
	err = encoding.In("utcOffset", res)
	if err != nil {
		return out0, out1, out2, out3, err
	}
	out2, err = encoding.AsInt32(res["utcOffset"])
	if err != nil {
		return out0, out1, out2, out3, err
	}
	err = encoding.In("currentTime", res)
	if err != nil {
		return out0, out1, out2, out3, err
	}
	out3, err = encoding.AsTime(res["currentTime"])
	if err != nil {
		return out0, out1, out2, out3, err
	}
	return out0, out1, out2, out3, nil
}
