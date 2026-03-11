// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package accumulatingnumericsensor

import (
	"context"
	"time"

	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40413/sensors/numericsensor"
)

func init() {
	object.Register(NewAccumulatingNumericSensor)
}

type _AccumulatingNumericSensor struct {
	numericsensor.NumericSensor
}

// NewAccumulatingNumericSensor returns a new AccumulatingNumericSensor interface for the object at given RID.
func NewAccumulatingNumericSensor(rid string, caller idl.Caller) AccumulatingNumericSensor {
	return &_AccumulatingNumericSensor{numericsensor.NewNumericSensor(rid, caller)}
}

func (a *_AccumulatingNumericSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.AccumulatingNumericSensor",
		Major: 2, Submajor: 0, Minor: 9,
	}
}

func (a *_AccumulatingNumericSensor) ResetValue(ctx context.Context) error {
	_, err := a.Caller().Call(ctx, a.RID(), "resetValue", nil)
	return err
}

func (a *_AccumulatingNumericSensor) GetLastResetTime(ctx context.Context) (time.Time, error) {
	var ret time.Time
	val, err := a.Caller().Call(ctx, a.RID(), "getLastResetTime", nil)
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
	ret, err = encoding.AsTime(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
