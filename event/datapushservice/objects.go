// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datapushservice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewDataPushService)
}

type _DataPushService struct {
	idl.Object
}

// NewDataPushService returns a new DataPushService interface for the object at given RID.
func NewDataPushService(rid string, caller idl.Caller) DataPushService {
	return &_DataPushService{idl.NewObject(rid, caller)}
}

func (d *_DataPushService) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.DataPushService",
		Major: 1, Submajor: 0, Minor: 3,
	}
}

func (d *_DataPushService) AddEntry(ctx context.Context, in0 EntrySettings) (int32, int32, error) {
	var ret int32
	var out0 int32
	val, err := d.Caller().Call(ctx, d.RID(), "addEntry", map[string]any{
		"entrySettings": in0.Encode(),
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
	err = encoding.In("entryId", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.AsInt32(res["entryId"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (d *_DataPushService) ModifyEntry(ctx context.Context, in0 int32, in1 EntrySettings) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "modifyEntry", map[string]any{
		"entryId":       in0,
		"entrySettings": in1.Encode(),
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

func (d *_DataPushService) DeleteEntry(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "deleteEntry", map[string]any{
		"entryId": in0,
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

func (d *_DataPushService) GetEntry(ctx context.Context, in0 int32) (int32, EntrySettings, error) {
	var ret int32
	var out0 EntrySettings
	val, err := d.Caller().Call(ctx, d.RID(), "getEntry", map[string]any{
		"entryId": in0,
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
	err = encoding.In("entrySettings", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["entrySettings"], d.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (d *_DataPushService) ListEntries(ctx context.Context) (map[int32]EntrySettings, error) {
	var ret map[int32]EntrySettings
	val, err := d.Caller().Call(ctx, d.RID(), "listEntries", nil)
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
	ret = make(map[int32]EntrySettings, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return ret, err
		}
		var v0 EntrySettings
		err = v0.Decode(b0, d.Caller())
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

func (d *_DataPushService) PushData(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := d.Caller().Call(ctx, d.RID(), "pushData", map[string]any{
		"entryId": in0,
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

func (d *_DataPushService) CancelDataPush(ctx context.Context, in0 int32) error {
	_, err := d.Caller().Call(ctx, d.RID(), "cancelDataPush", map[string]any{
		"entryId": in0,
	})
	return err
}

func (d *_DataPushService) GetEntryStatus(ctx context.Context, in0 int32) (int32, EntryStatus, error) {
	var ret int32
	var out0 EntryStatus
	val, err := d.Caller().Call(ctx, d.RID(), "getEntryStatus", map[string]any{
		"entryId": in0,
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
	err = encoding.In("entryStatus", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["entryStatus"], d.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}
