// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
)

func init() {
	object.Register(NewSensor)
}

type _Sensor struct {
	idl.Object
}

// NewSensor returns a new Sensor interface for the object at given RID.
func NewSensor(rid string, caller idl.Caller) Sensor {
	return &_Sensor{idl.NewObject(rid, caller)}
}

func (s *_Sensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Sensor",
		Major: 4, Submajor: 0, Minor: 7,
	}
}

func (s *_Sensor) GetTypeSpec(ctx context.Context) (TypeSpec, error) {
	var ret TypeSpec
	val, err := s.Caller().Call(ctx, s.RID(), "getTypeSpec", nil)
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

func (s *_Sensor) IsTypeChangeAllowed(ctx context.Context) (bool, error) {
	var ret bool
	val, err := s.Caller().Call(ctx, s.RID(), "isTypeChangeAllowed", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Sensor) SetType(ctx context.Context, in0 int32, in1 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setType", map[string]any{
		"type": in0,
		"unit": in1,
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
