// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package testdisplay

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
)

func init() {
	object.Register(NewDisplay)
}

type _Display struct {
	idl.Object
}

// NewDisplay returns a new Display interface for the object at given RID.
func NewDisplay(rid string, caller idl.Caller) Display {
	return &_Display{idl.NewObject(rid, caller)}
}

func (d *_Display) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "test.Display",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (d *_Display) GetInfo(ctx context.Context) (DisplayInfo, error) {
	var ret DisplayInfo
	val, err := d.Caller().Call(ctx, d.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], d.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (d *_Display) TestSequence(ctx context.Context, in0 int32) error {
	_, err := d.Caller().Call(ctx, d.RID(), "testSequence", map[string]any{
		"cycleTime_ms": in0,
	})
	return err
}

func (d *_Display) EnterTestMode(ctx context.Context, in0 bool) error {
	_, err := d.Caller().Call(ctx, d.RID(), "enterTestMode", map[string]any{
		"showColorNames": in0,
	})
	return err
}

func (d *_Display) GetTestStatus(ctx context.Context) (DisplayTestStatus, error) {
	var ret DisplayTestStatus
	val, err := d.Caller().Call(ctx, d.RID(), "getTestStatus", nil)
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
	ret, err = encoding.AsEnum[DisplayTestStatus](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
