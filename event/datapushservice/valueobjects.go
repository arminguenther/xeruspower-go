// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datapushservice

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/idl/event"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() EntryAddedEvent { return &_EntryAddedEvent{} })
	valobj.Register(func() EntryDeletedEvent { return &_EntryDeletedEvent{} })
	valobj.Register(func() EntryModifiedEvent { return &_EntryModifiedEvent{} })
	valobj.Register(func() EntryStatusChangedEvent { return &_EntryStatusChangedEvent{} })
}

type _EntryAddedEvent struct {
	userevent.UserEvent
	entryId  int32
	settings EntrySettings
}

func (e *_EntryAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.DataPushService_1_0_3.EntryAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EntryAddedEvent) EntryId() int32 {
	return e.entryId
}

func (e *_EntryAddedEvent) Settings() EntrySettings {
	return e.settings
}

func (e *_EntryAddedEvent) isEntryAddedEvent() {}

type _EntryModifiedEvent struct {
	userevent.UserEvent
	entryId     int32
	oldSettings EntrySettings
	newSettings EntrySettings
}

func (e *_EntryModifiedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.DataPushService_1_0_3.EntryModifiedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EntryModifiedEvent) EntryId() int32 {
	return e.entryId
}

func (e *_EntryModifiedEvent) OldSettings() EntrySettings {
	return e.oldSettings
}

func (e *_EntryModifiedEvent) NewSettings() EntrySettings {
	return e.newSettings
}

func (e *_EntryModifiedEvent) isEntryModifiedEvent() {}

type _EntryDeletedEvent struct {
	userevent.UserEvent
	entryId int32
}

func (e *_EntryDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.DataPushService_1_0_3.EntryDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EntryDeletedEvent) EntryId() int32 {
	return e.entryId
}

func (e *_EntryDeletedEvent) isEntryDeletedEvent() {}

type _EntryStatusChangedEvent struct {
	event.Event
	entryId   int32
	newStatus EntryStatus
}

func (e *_EntryStatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.DataPushService_1_0_3.EntryStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_EntryStatusChangedEvent) EntryId() int32 {
	return e.entryId
}

func (e *_EntryStatusChangedEvent) NewStatus() EntryStatus {
	return e.newStatus
}

func (e *_EntryStatusChangedEvent) isEntryStatusChangedEvent() {}
