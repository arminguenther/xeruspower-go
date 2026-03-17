// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package hardwarehealth

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
)

func init() {
	object.Register(NewHardwareHealth)
}

type _HardwareHealth struct {
	idl.Object
}

// NewHardwareHealth returns a new HardwareHealth interface for the object at given RID.
func NewHardwareHealth(rid string, caller idl.Caller) HardwareHealth {
	return &_HardwareHealth{idl.NewObject(rid, caller)}
}

func (h *_HardwareHealth) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "fitness.HardwareHealth",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (h *_HardwareHealth) GetFailures(ctx context.Context) ([]Failure, error) {
	var ret []Failure
	val, err := h.Caller().Call(ctx, h.RID(), "getFailures", nil)
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
	ret = make([]Failure, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Failure
		err = e0.Decode(a0, h.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
