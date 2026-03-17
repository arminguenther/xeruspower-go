// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensor

import (
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/idl/event"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() TypeSpecChangedEvent { return &_TypeSpecChangedEvent{} })
}

type _TypeSpecChangedEvent struct {
	event.Event
	oldTypeSpec TypeSpec
	newTypeSpec TypeSpec
}

func (t *_TypeSpecChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.Sensor_4_0_7.TypeSpecChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_TypeSpecChangedEvent) OldTypeSpec() TypeSpec {
	return t.oldTypeSpec
}

func (t *_TypeSpecChangedEvent) NewTypeSpec() TypeSpec {
	return t.newTypeSpec
}

func (t *_TypeSpecChangedEvent) isTypeSpecChangedEvent() {}
