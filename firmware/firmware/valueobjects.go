// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package firmware

import (
	"github.com/arminguenther/xeruspower-go/v40032/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/idl/event"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() SystemShutdownEvent { return &_SystemShutdownEvent{} })
	valobj.Register(func() SystemStartupEvent { return &_SystemStartupEvent{} })
	valobj.Register(func() UpdateCompletedEvent { return &_UpdateCompletedEvent{} })
	valobj.Register(func() UpdateEvent { return &_UpdateEvent{} })
	valobj.Register(func() UpdateFailedEvent { return &_UpdateFailedEvent{} })
	valobj.Register(func() UpdateStartedEvent { return &_UpdateStartedEvent{} })
	valobj.Register(func() ValidationFailedEvent { return &_ValidationFailedEvent{} })
}

type _SystemStartupEvent struct {
	event.Event
}

func (s *_SystemStartupEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.SystemStartupEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SystemStartupEvent) isSystemStartupEvent() {}

type _SystemShutdownEvent struct {
	userevent.UserEvent
}

func (s *_SystemShutdownEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.SystemShutdownEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SystemShutdownEvent) isSystemShutdownEvent() {}

type _ValidationFailedEvent struct {
	userevent.UserEvent
}

func (f *_ValidationFailedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.FirmwareValidationFailedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_ValidationFailedEvent) isValidationFailedEvent() {}

type _UpdateEvent struct {
	userevent.UserEvent
	oldVersion string
	newVersion string
}

func (f *_UpdateEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.FirmwareUpdateEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_UpdateEvent) OldVersion() string {
	return f.oldVersion
}

func (f *_UpdateEvent) NewVersion() string {
	return f.newVersion
}

func (f *_UpdateEvent) isUpdateEvent() {}

type _UpdateStartedEvent struct {
	UpdateEvent
}

func (f *_UpdateStartedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.FirmwareUpdateStartedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_UpdateStartedEvent) isUpdateStartedEvent() {}

type _UpdateCompletedEvent struct {
	UpdateEvent
}

func (f *_UpdateCompletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.FirmwareUpdateCompletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_UpdateCompletedEvent) isUpdateCompletedEvent() {}

type _UpdateFailedEvent struct {
	UpdateEvent
}

func (f *_UpdateFailedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.FirmwareUpdateFailedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_UpdateFailedEvent) isUpdateFailedEvent() {}
