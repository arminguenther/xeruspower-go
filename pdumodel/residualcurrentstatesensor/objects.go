// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package residualcurrentstatesensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40300/sensors/statesensor"
)

func init() {
	object.Register(NewResidualCurrentStateSensor)
}

type _ResidualCurrentStateSensor struct {
	statesensor.StateSensor
}

// NewResidualCurrentStateSensor returns a new ResidualCurrentStateSensor interface for the object at given RID.
func NewResidualCurrentStateSensor(rid string, caller idl.Caller) ResidualCurrentStateSensor {
	return &_ResidualCurrentStateSensor{statesensor.NewStateSensor(rid, caller)}
}

func (r *_ResidualCurrentStateSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.ResidualCurrentStateSensor",
		Major: 2, Submajor: 0, Minor: 7,
	}
}

func (r *_ResidualCurrentStateSensor) StartSelfTest(ctx context.Context) (int32, error) {
	var ret int32
	val, err := r.Caller().Call(ctx, r.RID(), "startSelfTest", nil)
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
