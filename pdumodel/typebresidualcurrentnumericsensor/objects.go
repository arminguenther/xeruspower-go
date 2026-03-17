// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package typebresidualcurrentnumericsensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/sensors/numericsensor"
)

func init() {
	object.Register(NewTypeBResidualCurrentNumericSensor)
}

type _TypeBResidualCurrentNumericSensor struct {
	numericsensor.NumericSensor
}

// NewTypeBResidualCurrentNumericSensor returns a new TypeBResidualCurrentNumericSensor interface for the object at given RID.
func NewTypeBResidualCurrentNumericSensor(rid string, caller idl.Caller) TypeBResidualCurrentNumericSensor {
	return &_TypeBResidualCurrentNumericSensor{numericsensor.NewNumericSensor(rid, caller)}
}

func (t *_TypeBResidualCurrentNumericSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.TypeBResidualCurrentNumericSensor",
		Major: 1, Submajor: 0, Minor: 5,
	}
}

func (t *_TypeBResidualCurrentNumericSensor) Degauss(ctx context.Context) error {
	_, err := t.Caller().Call(ctx, t.RID(), "degauss", nil)
	return err
}
