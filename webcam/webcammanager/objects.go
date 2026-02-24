// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package webcammanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/webcam/webcam"
	"github.com/arminguenther/xeruspower-go/webcam/webcamchannel"
)

func init() {
	object.Register(NewWebcamManager)
}

type _WebcamManager struct {
	idl.Object
}

// NewWebcamManager returns a new WebcamManager interface for the object at given RID.
func NewWebcamManager(rid string, caller idl.Caller) WebcamManager {
	return &_WebcamManager{idl.NewObject(rid, caller)}
}

func (w *_WebcamManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.WebcamManager",
		Major: 2, Submajor: 0, Minor: 2,
	}
}

func (w *_WebcamManager) GetWebcams(ctx context.Context) ([]webcam.Webcam, error) {
	var ret []webcam.Webcam
	val, err := w.Caller().Call(ctx, w.RID(), "getWebcams", nil)
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
	ret = make([]webcam.Webcam, 0, len(s0))
	for _, a0 := range s0 {
		var e0 webcam.Webcam
		e0, err = object.As[webcam.Webcam](a0, w.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (w *_WebcamManager) GetChannel(ctx context.Context, in0 webcam.Webcam, in1 string) (int32, webcamchannel.Channel, error) {
	var ret int32
	var out0 webcamchannel.Channel
	val, err := w.Caller().Call(ctx, w.RID(), "getChannel", map[string]any{
		"webcam":     object.ToMap(in0),
		"clientType": in1,
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
	err = encoding.In("channel", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = object.As[webcamchannel.Channel](res["channel"], w.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (w *_WebcamManager) GetChannels(ctx context.Context) ([]webcamchannel.Channel, error) {
	var ret []webcamchannel.Channel
	val, err := w.Caller().Call(ctx, w.RID(), "getChannels", nil)
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
	ret = make([]webcamchannel.Channel, 0, len(s0))
	for _, a0 := range s0 {
		var e0 webcamchannel.Channel
		e0, err = object.As[webcamchannel.Channel](a0, w.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (w *_WebcamManager) RemoveClientType(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := w.Caller().Call(ctx, w.RID(), "removeClientType", map[string]any{
		"clientType": in0,
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

func (w *_WebcamManager) GetClientTypes(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := w.Caller().Call(ctx, w.RID(), "getClientTypes", nil)
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

func (w *_WebcamManager) GetClientTypePriorities(ctx context.Context) (map[string]Priority, error) {
	var ret map[string]Priority
	val, err := w.Caller().Call(ctx, w.RID(), "getClientTypePriorities", nil)
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
	ret = make(map[string]Priority, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 Priority
		v0, err = encoding.AsEnum[Priority](b0)
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

func (w *_WebcamManager) SetClientTypePriorities(ctx context.Context, in0 map[string]Priority) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0})
	}
	val, err := w.Caller().Call(ctx, w.RID(), "setClientTypePriorities", map[string]any{
		"priorities": s0,
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

func (w *_WebcamManager) GetWebcamPriorities(ctx context.Context) (map[string]Priority, error) {
	var ret map[string]Priority
	val, err := w.Caller().Call(ctx, w.RID(), "getWebcamPriorities", nil)
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
	ret = make(map[string]Priority, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 Priority
		v0, err = encoding.AsEnum[Priority](b0)
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

func (w *_WebcamManager) SetWebcamPriorities(ctx context.Context, in0 map[string]Priority) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for k0, v0 := range in0 {
		s0 = append(s0, map[string]any{"key": k0, "value": v0})
	}
	val, err := w.Caller().Call(ctx, w.RID(), "setWebcamPriorities", map[string]any{
		"priorities": s0,
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
