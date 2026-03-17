// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powerqualitysensor

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/sensors/statesensor"
)

func init() {
	object.Register(NewPowerQualitySensor)
}

type _PowerQualitySensor struct {
	statesensor.StateSensor
}

// NewPowerQualitySensor returns a new PowerQualitySensor interface for the object at given RID.
func NewPowerQualitySensor(rid string, caller idl.Caller) PowerQualitySensor {
	return &_PowerQualitySensor{statesensor.NewStateSensor(rid, caller)}
}

func (p *_PowerQualitySensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerQualitySensor",
		Major: 2, Submajor: 0, Minor: 6,
	}
}
