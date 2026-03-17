// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package displaycontrol

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
)

func init() {
	object.Register(NewDisplayControl)
}

type _DisplayControl struct {
	idl.Object
}

// NewDisplayControl returns a new DisplayControl interface for the object at given RID.
func NewDisplayControl(rid string, caller idl.Caller) DisplayControl {
	return &_DisplayControl{idl.NewObject(rid, caller)}
}

func (d *_DisplayControl) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "display.DisplayControl",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (d *_DisplayControl) GetAvailableDefaultViews(ctx context.Context) ([]DefaultViewItem, error) {
	var ret []DefaultViewItem
	val, err := d.Caller().Call(ctx, d.RID(), "getAvailableDefaultViews", nil)
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
	ret = make([]DefaultViewItem, 0, len(s0))
	for _, a0 := range s0 {
		var e0 DefaultViewItem
		err = e0.Decode(a0, d.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (d *_DisplayControl) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := d.Caller().Call(ctx, d.RID(), "getSettings", nil)
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

func (d *_DisplayControl) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "setSettings", map[string]any{
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

func (d *_DisplayControl) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
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

func (d *_DisplayControl) ShowSmiley(ctx context.Context, in0 bool, in1 bool, in2 string) error {
	_, err := d.Caller().Call(ctx, d.RID(), "showSmiley", map[string]any{
		"on":    in0,
		"happy": in1,
		"msg":   in2,
	})
	return err
}
