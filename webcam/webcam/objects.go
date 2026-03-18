// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package webcam

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
)

func init() {
	object.Register(NewWebcam)
}

type _Webcam struct {
	idl.Object
}

// NewWebcam returns a new Webcam interface for the object at given RID.
func NewWebcam(rid string, caller idl.Caller) Webcam {
	return &_Webcam{idl.NewObject(rid, caller)}
}

func (w *_Webcam) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.Webcam",
		Major: 2, Submajor: 0, Minor: 1,
	}
}

func (w *_Webcam) GetInformation(ctx context.Context) (Information, error) {
	var ret Information
	val, err := w.Caller().Call(ctx, w.RID(), "getInformation", nil)
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
	err = ret.Decode(res["_ret_"], w.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (w *_Webcam) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := w.Caller().Call(ctx, w.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], w.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (w *_Webcam) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := w.Caller().Call(ctx, w.RID(), "setSettings", map[string]any{
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

func (w *_Webcam) SetControls(ctx context.Context, in0 Controls) (int32, error) {
	var ret int32
	val, err := w.Caller().Call(ctx, w.RID(), "setControls", map[string]any{
		"controls": in0.Encode(),
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

func (w *_Webcam) GetControlDefaults(ctx context.Context) (Controls, error) {
	var ret Controls
	val, err := w.Caller().Call(ctx, w.RID(), "getControlDefaults", nil)
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
	err = ret.Decode(res["_ret_"], w.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
