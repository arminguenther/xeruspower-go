// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package hardwarehealth

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() FailureStatusChangedEvent { return &_FailureStatusChangedEvent{} })
}

type _FailureStatusChangedEvent struct {
	event.Event
	componentId string
	failureType int32
	isAsserted  bool
}

func (f *_FailureStatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "fitness.HardwareHealth_1_0_3.FailureStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FailureStatusChangedEvent) ComponentId() string {
	return f.componentId
}

func (f *_FailureStatusChangedEvent) FailureType() int32 {
	return f.failureType
}

func (f *_FailureStatusChangedEvent) IsAsserted() bool {
	return f.isAsserted
}

func (f *_FailureStatusChangedEvent) isFailureStatusChangedEvent() {}
