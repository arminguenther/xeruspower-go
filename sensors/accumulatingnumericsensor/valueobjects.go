// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package accumulatingnumericsensor

import (
	"github.com/arminguenther/xeruspower-go/v40020/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40020/sensors/numericsensor"
)

func init() {
	valobj.Register(func() ResetEvent { return &_ResetEvent{} })
}

type _ResetEvent struct {
	userevent.UserEvent
	oldReading numericsensor.Reading
	newReading numericsensor.Reading
}

func (r *_ResetEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.AccumulatingNumericSensor_2_0_6.ResetEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_ResetEvent) OldReading() numericsensor.Reading {
	return r.oldReading
}

func (r *_ResetEvent) NewReading() numericsensor.Reading {
	return r.newReading
}

func (r *_ResetEvent) isResetEvent() {}
