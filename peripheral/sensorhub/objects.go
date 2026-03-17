// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorhub

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40220/peripheral/peripheraldevicepackage"
)

func init() {
	object.Register(NewSensorHub)
}

type _SensorHub struct {
	idl.Object
}

// NewSensorHub returns a new SensorHub interface for the object at given RID.
func NewSensorHub(rid string, caller idl.Caller) SensorHub {
	return &_SensorHub{idl.NewObject(rid, caller)}
}

func (s *_SensorHub) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.SensorHub",
		Major: 2, Submajor: 0, Minor: 4,
	}
}

func (s *_SensorHub) GetDeviceInfo(ctx context.Context) (DeviceInfo, error) {
	var ret DeviceInfo
	val, err := s.Caller().Call(ctx, s.RID(), "getDeviceInfo", nil)
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

func (s *_SensorHub) GetPackageInfo(ctx context.Context) (peripheraldevicepackage.PackageInfo, error) {
	var ret peripheraldevicepackage.PackageInfo
	val, err := s.Caller().Call(ctx, s.RID(), "getPackageInfo", nil)
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
