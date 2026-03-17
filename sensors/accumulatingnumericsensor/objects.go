// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package accumulatingnumericsensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40020/sensors/numericsensor"
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
		Major: 2, Submajor: 0, Minor: 6,
	}
}

func (a *_AccumulatingNumericSensor) ResetValue(ctx context.Context) error {
	_, err := a.Caller().Call(ctx, a.RID(), "resetValue", nil)
	return err
}
