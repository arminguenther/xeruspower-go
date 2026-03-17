// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package firmware

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
)

func init() {
	object.Register(NewFirmware)
}

type _Firmware struct {
	idl.Object
}

// NewFirmware returns a new Firmware interface for the object at given RID.
func NewFirmware(rid string, caller idl.Caller) Firmware {
	return &_Firmware{idl.NewObject(rid, caller)}
}

func (f *_Firmware) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.Firmware",
		Major: 2, Submajor: 0, Minor: 2,
	}
}

func (f *_Firmware) Reboot(ctx context.Context) error {
	_, err := f.Caller().Call(ctx, f.RID(), "reboot", nil)
	return err
}

func (f *_Firmware) FactoryReset(ctx context.Context) error {
	_, err := f.Caller().Call(ctx, f.RID(), "factoryReset", nil)
	return err
}

func (f *_Firmware) HardFactoryReset(ctx context.Context) (int32, error) {
	var ret int32
	val, err := f.Caller().Call(ctx, f.RID(), "hardFactoryReset", nil)
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

func (f *_Firmware) ManufacturingReset(ctx context.Context) (int32, error) {
	var ret int32
	val, err := f.Caller().Call(ctx, f.RID(), "manufacturingReset", nil)
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

func (f *_Firmware) GetVersion(ctx context.Context) (string, error) {
	var ret string
	val, err := f.Caller().Call(ctx, f.RID(), "getVersion", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (f *_Firmware) GetUpdateHistory(ctx context.Context) ([]UpdateHistoryEntry, error) {
	var ret []UpdateHistoryEntry
	val, err := f.Caller().Call(ctx, f.RID(), "getUpdateHistory", nil)
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
	ret = make([]UpdateHistoryEntry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 UpdateHistoryEntry
		err = e0.Decode(a0, f.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (f *_Firmware) GetImageStatus(ctx context.Context) (ImageStatus, error) {
	var ret ImageStatus
	val, err := f.Caller().Call(ctx, f.RID(), "getImageStatus", nil)
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
	err = ret.Decode(res["_ret_"], f.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (f *_Firmware) DiscardImage(ctx context.Context) error {
	_, err := f.Caller().Call(ctx, f.RID(), "discardImage", nil)
	return err
}

func (f *_Firmware) GetImageInfo(ctx context.Context) (bool, ImageInfo, error) {
	var ret bool
	var out0 ImageInfo
	val, err := f.Caller().Call(ctx, f.RID(), "getImageInfo", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("info", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["info"], f.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (f *_Firmware) StartUpdate(ctx context.Context, in0 []UpdateFlags) error {
	_, err := f.Caller().Call(ctx, f.RID(), "startUpdate", map[string]any{
		"flags": encoding.NonNilSlice(in0),
	})
	return err
}
