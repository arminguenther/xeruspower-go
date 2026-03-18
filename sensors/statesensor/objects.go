// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package statesensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/sensors/sensor"
)

func init() {
	object.Register(NewStateSensor)
}

type _StateSensor struct {
	sensor.Sensor
}

// NewStateSensor returns a new StateSensor interface for the object at given RID.
func NewStateSensor(rid string, caller idl.Caller) StateSensor {
	return &_StateSensor{sensor.NewSensor(rid, caller)}
}

func (s *_StateSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.StateSensor",
		Major: 4, Submajor: 0, Minor: 5,
	}
}

func (s *_StateSensor) GetState(ctx context.Context) (State, error) {
	var ret State
	val, err := s.Caller().Call(ctx, s.RID(), "getState", nil)
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
