// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package testunit

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/test/testdisplay"
)

func init() {
	object.Register(NewUnit)
}

type _Unit struct {
	idl.Object
}

// NewUnit returns a new Unit interface for the object at given RID.
func NewUnit(rid string, caller idl.Caller) Unit {
	return &_Unit{idl.NewObject(rid, caller)}
}

func (u *_Unit) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "test.Unit",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (u *_Unit) GetDisplays(ctx context.Context) ([]testdisplay.Display, error) {
	var ret []testdisplay.Display
	val, err := u.Caller().Call(ctx, u.RID(), "getDisplays", nil)
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
	ret = make([]testdisplay.Display, 0, len(s0))
	for _, a0 := range s0 {
		var e0 testdisplay.Display
		e0, err = object.As[testdisplay.Display](a0, u.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (u *_Unit) GetButtonStates(ctx context.Context) ([]bool, error) {
	var ret []bool
	val, err := u.Caller().Call(ctx, u.RID(), "getButtonStates", nil)
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
	ret = make([]bool, 0, len(s0))
	for _, a0 := range s0 {
		var e0 bool
		e0, err = encoding.Is[bool](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (u *_Unit) SetBuzzer(ctx context.Context, in0 bool) error {
	_, err := u.Caller().Call(ctx, u.RID(), "setBuzzer", map[string]any{
		"isOn": in0,
	})
	return err
}

func (u *_Unit) ResetAllSubControllers(ctx context.Context) error {
	_, err := u.Caller().Call(ctx, u.RID(), "resetAllSubControllers", nil)
	return err
}

func (u *_Unit) TriggerSubControllerWatchdog(ctx context.Context, in0 int32) error {
	_, err := u.Caller().Call(ctx, u.RID(), "triggerSubControllerWatchdog", map[string]any{
		"rs485Addr": in0,
	})
	return err
}
