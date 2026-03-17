// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gsmmodem

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
)

func init() {
	object.Register(NewGsmModem)
}

type _GsmModem struct {
	idl.Object
}

// NewGsmModem returns a new GsmModem interface for the object at given RID.
func NewGsmModem(rid string, caller idl.Caller) GsmModem {
	return &_GsmModem{idl.NewObject(rid, caller)}
}

func (g *_GsmModem) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "serial.GsmModem",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (g *_GsmModem) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := g.Caller().Call(ctx, g.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], g.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (g *_GsmModem) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := g.Caller().Call(ctx, g.RID(), "setSettings", map[string]any{
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

func (g *_GsmModem) SendSms(ctx context.Context, in0 string, in1 string) (int32, error) {
	var ret int32
	val, err := g.Caller().Call(ctx, g.RID(), "sendSms", map[string]any{
		"recipient": in0,
		"text":      in1,
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

func (g *_GsmModem) SendTestSms(ctx context.Context, in0 string, in1 Settings) (int32, error) {
	var ret int32
	val, err := g.Caller().Call(ctx, g.RID(), "sendTestSms", map[string]any{
		"recipient":    in0,
		"testSettings": in1.Encode(),
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

func (g *_GsmModem) GetInformation(ctx context.Context) (int32, Information, error) {
	var ret int32
	var out0 Information
	val, err := g.Caller().Call(ctx, g.RID(), "getInformation", nil)
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
	err = encoding.In("info", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["info"], g.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (g *_GsmModem) GetInformationWithPin(ctx context.Context, in0 string) (int32, Information, error) {
	var ret int32
	var out0 Information
	val, err := g.Caller().Call(ctx, g.RID(), "getInformationWithPin", map[string]any{
		"pin": in0,
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
	err = encoding.In("info", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["info"], g.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (g *_GsmModem) GetSimSecurityStatus(ctx context.Context) (int32, SimSecurityStatus, error) {
	var ret int32
	var out0 SimSecurityStatus
	val, err := g.Caller().Call(ctx, g.RID(), "getSimSecurityStatus", nil)
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
	err = encoding.In("simStatus", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.AsEnum[SimSecurityStatus](res["simStatus"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (g *_GsmModem) UnlockSimCard(ctx context.Context, in0 string, in1 string) (int32, error) {
	var ret int32
	val, err := g.Caller().Call(ctx, g.RID(), "unlockSimCard", map[string]any{
		"puk":    in0,
		"newPin": in1,
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
