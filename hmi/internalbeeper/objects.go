// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package internalbeeper

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewInternalBeeper)
}

type _InternalBeeper struct {
	idl.Object
}

// NewInternalBeeper returns a new InternalBeeper interface for the object at given RID.
func NewInternalBeeper(rid string, caller idl.Caller) InternalBeeper {
	return &_InternalBeeper{idl.NewObject(rid, caller)}
}

func (i *_InternalBeeper) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "hmi.InternalBeeper",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (i *_InternalBeeper) Mute(ctx context.Context, in0 bool) error {
	_, err := i.Caller().Call(ctx, i.RID(), "mute", map[string]any{
		"muted": in0,
	})
	return err
}

func (i *_InternalBeeper) IsMuted(ctx context.Context) (bool, error) {
	var ret bool
	val, err := i.Caller().Call(ctx, i.RID(), "isMuted", nil)
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

func (i *_InternalBeeper) Activate(ctx context.Context, in0 bool, in1 string, in2 int32) error {
	_, err := i.Caller().Call(ctx, i.RID(), "activate", map[string]any{
		"on":      in0,
		"reason":  in1,
		"timeout": in2,
	})
	return err
}

func (i *_InternalBeeper) GetState(ctx context.Context) (State, string, bool, error) {
	var ret State
	var out0 string
	var out1 bool
	val, err := i.Caller().Call(ctx, i.RID(), "getState", nil)
	if err != nil {
		return ret, out0, out1, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, out1, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, out1, err
	}
	ret, err = encoding.AsEnum[State](res["_ret_"])
	if err != nil {
		return ret, out0, out1, err
	}
	err = encoding.In("reason", res)
	if err != nil {
		return ret, out0, out1, err
	}
	out0, err = encoding.Is[string](res["reason"])
	if err != nil {
		return ret, out0, out1, err
	}
	err = encoding.In("mutedTemporarily", res)
	if err != nil {
		return ret, out0, out1, err
	}
	out1, err = encoding.Is[bool](res["mutedTemporarily"])
	if err != nil {
		return ret, out0, out1, err
	}
	return ret, out0, out1, nil
}

func (i *_InternalBeeper) MuteCurrentActivation(ctx context.Context) error {
	_, err := i.Caller().Call(ctx, i.RID(), "muteCurrentActivation", nil)
	return err
}
