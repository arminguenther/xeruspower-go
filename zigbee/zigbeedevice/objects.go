// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeedevice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
)

func init() {
	object.Register(NewZigbeeDevice)
}

type _ZigbeeDevice struct {
	idl.Object
}

// NewZigbeeDevice returns a new ZigbeeDevice interface for the object at given RID.
func NewZigbeeDevice(rid string, caller idl.Caller) ZigbeeDevice {
	return &_ZigbeeDevice{idl.NewObject(rid, caller)}
}

func (z *_ZigbeeDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (z *_ZigbeeDevice) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := z.Caller().Call(ctx, z.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], z.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (z *_ZigbeeDevice) GetClusterValues(ctx context.Context) ([]ClusterValue, error) {
	var ret []ClusterValue
	val, err := z.Caller().Call(ctx, z.RID(), "getClusterValues", nil)
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
	ret = make([]ClusterValue, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ClusterValue
		err = e0.Decode(a0, z.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
