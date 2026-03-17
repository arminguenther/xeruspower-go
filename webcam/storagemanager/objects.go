// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package storagemanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40200/webcam/webcam"
)

func init() {
	object.Register(NewStorageManager)
}

type _StorageManager struct {
	idl.Object
}

// NewStorageManager returns a new StorageManager interface for the object at given RID.
func NewStorageManager(rid string, caller idl.Caller) StorageManager {
	return &_StorageManager{idl.NewObject(rid, caller)}
}

func (s *_StorageManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.StorageManager",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (s *_StorageManager) GetSupportedStorageTypes(ctx context.Context) ([]StorageType, error) {
	var ret []StorageType
	val, err := s.Caller().Call(ctx, s.RID(), "getSupportedStorageTypes", nil)
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
	ret = make([]StorageType, 0, len(s0))
	for _, a0 := range s0 {
		var e0 StorageType
		e0, err = encoding.AsEnum[StorageType](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_StorageManager) GetInformation(ctx context.Context) (StorageInformation, error) {
	var ret StorageInformation
	val, err := s.Caller().Call(ctx, s.RID(), "getInformation", nil)
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

func (s *_StorageManager) GetSettings(ctx context.Context) (StorageSettings, error) {
	var ret StorageSettings
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

func (s *_StorageManager) SetSettings(ctx context.Context, in0 StorageSettings) (int32, error) {
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

func (s *_StorageManager) AddImage(ctx context.Context, in0 webcam.Webcam, in1 webcam.Image) (int32, int64, error) {
	var ret int32
	var out0 int64
	val, err := s.Caller().Call(ctx, s.RID(), "addImage", map[string]any{
		"webcam": object.ToMap(in0),
		"image":  in1.Encode(),
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
	err = encoding.In("index", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.AsInt64(res["index"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (s *_StorageManager) RemoveImages(ctx context.Context, in0 webcam.Webcam, in1 int64, in2 int32, in3 Direction) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "removeImages", map[string]any{
		"webcam":    object.ToMap(in0),
		"start":     in1,
		"count":     in2,
		"direction": in3,
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

func (s *_StorageManager) GetMetaData(ctx context.Context, in0 webcam.Webcam, in1 int64, in2 int32, in3 Direction) (int32, []ImageStorageMetaData, error) {
	var ret int32
	var out0 []ImageStorageMetaData
	val, err := s.Caller().Call(ctx, s.RID(), "getMetaData", map[string]any{
		"webcam":    object.ToMap(in0),
		"start":     in1,
		"count":     in2,
		"direction": in3,
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
	err = encoding.In("meta", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["meta"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]ImageStorageMetaData, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ImageStorageMetaData
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (s *_StorageManager) GetImages(ctx context.Context, in0 webcam.Webcam, in1 int64, in2 int32, in3 Direction) (int32, []StorageImage, error) {
	var ret int32
	var out0 []StorageImage
	val, err := s.Caller().Call(ctx, s.RID(), "getImages", map[string]any{
		"webcam":    object.ToMap(in0),
		"start":     in1,
		"count":     in2,
		"direction": in3,
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
	err = encoding.In("image", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["image"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]StorageImage, 0, len(s0))
	for _, a0 := range s0 {
		var e0 StorageImage
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (s *_StorageManager) GetActivities(ctx context.Context) ([]Activity, error) {
	var ret []Activity
	val, err := s.Caller().Call(ctx, s.RID(), "getActivities", nil)
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
	ret = make([]Activity, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Activity
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_StorageManager) StartActivity(ctx context.Context, in0 webcam.Webcam, in1 int32, in2 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "startActivity", map[string]any{
		"webcam":   object.ToMap(in0),
		"count":    in1,
		"interval": in2,
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

func (s *_StorageManager) StartActivityWithFolder(ctx context.Context, in0 webcam.Webcam, in1 int32, in2 int32, in3 string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "startActivityWithFolder", map[string]any{
		"webcam":   object.ToMap(in0),
		"count":    in1,
		"interval": in2,
		"folder":   in3,
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

func (s *_StorageManager) StopActivity(ctx context.Context, in0 webcam.Webcam) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "stopActivity", map[string]any{
		"webcam": object.ToMap(in0),
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
