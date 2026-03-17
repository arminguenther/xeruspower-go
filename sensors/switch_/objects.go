// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package switch_

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40100/sensors/statesensor"
)

func init() {
	object.Register(NewSwitch)
}

type _Switch struct {
	statesensor.StateSensor
}

// NewSwitch returns a new Switch interface for the object at given RID.
func NewSwitch(rid string, caller idl.Caller) Switch {
	return &_Switch{statesensor.NewStateSensor(rid, caller)}
}

func (s *_Switch) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Switch",
		Major: 2, Submajor: 0, Minor: 8,
	}
}

func (s *_Switch) SetState(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setState", map[string]any{
		"newState": in0,
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
