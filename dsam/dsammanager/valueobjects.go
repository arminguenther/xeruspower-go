// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsammanager

import (
	"github.com/arminguenther/xeruspower-go/v40413/dsam/dsamdevice"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DsamAttachedEvent { return &_DsamAttachedEvent{} })
	valobj.Register(func() DsamControllerChangedEvent { return &_DsamControllerChangedEvent{} })
	valobj.Register(func() DsamDetachedEvent { return &_DsamDetachedEvent{} })
}

type _DsamAttachedEvent struct {
	event.Event
	info   dsamdevice.Info
	device dsamdevice.DsamDevice
}

func (d *_DsamAttachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamManager.DsamAttachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DsamAttachedEvent) Info() dsamdevice.Info {
	return d.info
}

func (d *_DsamAttachedEvent) Device() dsamdevice.DsamDevice {
	return d.device
}

func (d *_DsamAttachedEvent) isDsamAttachedEvent() {}

type _DsamDetachedEvent struct {
	event.Event
	info dsamdevice.Info
}

func (d *_DsamDetachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamManager.DsamDetachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DsamDetachedEvent) Info() dsamdevice.Info {
	return d.info
}

func (d *_DsamDetachedEvent) isDsamDetachedEvent() {}

type _DsamControllerChangedEvent struct {
	event.Event
	dsamNumber  int32
	reset       bool
	resetReason string
}

func (d *_DsamControllerChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "dsam.DsamManager.DsamControllerChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DsamControllerChangedEvent) DsamNumber() int32 {
	return d.dsamNumber
}

func (d *_DsamControllerChangedEvent) Reset() bool {
	return d.reset
}

func (d *_DsamControllerChangedEvent) ResetReason() string {
	return d.resetReason
}

func (d *_DsamControllerChangedEvent) isDsamControllerChangedEvent() {}
