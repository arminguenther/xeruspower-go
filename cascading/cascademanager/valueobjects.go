// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cascademanager

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/idl/event"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() LinkPortStatusChangedEvent { return &_LinkPortStatusChangedEvent{} })
	valobj.Register(func() LinkUnitAddedEvent { return &_LinkUnitAddedEvent{} })
	valobj.Register(func() LinkUnitReleasedEvent { return &_LinkUnitReleasedEvent{} })
	valobj.Register(func() LinkUnitStatusChangedEvent { return &_LinkUnitStatusChangedEvent{} })
	valobj.Register(func() RoleChangedEvent { return &_RoleChangedEvent{} })
}

type _RoleChangedEvent struct {
	event.Event
	oldRole     Role
	newRole     Role
	primaryUnit string
}

func (r *_RoleChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cascading.CascadeManager_3_0_0.RoleChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RoleChangedEvent) OldRole() Role {
	return r.oldRole
}

func (r *_RoleChangedEvent) NewRole() Role {
	return r.newRole
}

func (r *_RoleChangedEvent) PrimaryUnit() string {
	return r.primaryUnit
}

func (r *_RoleChangedEvent) isRoleChangedEvent() {}

type _LinkUnitAddedEvent struct {
	userevent.UserEvent
	linkId int32
	type_  LinkUnitType
	host   string
}

func (l *_LinkUnitAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cascading.CascadeManager_3_0_0.LinkUnitAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LinkUnitAddedEvent) LinkId() int32 {
	return l.linkId
}

func (l *_LinkUnitAddedEvent) Type() LinkUnitType {
	return l.type_
}

func (l *_LinkUnitAddedEvent) Host() string {
	return l.host
}

func (l *_LinkUnitAddedEvent) isLinkUnitAddedEvent() {}

type _LinkUnitReleasedEvent struct {
	userevent.UserEvent
	linkId int32
	type_  LinkUnitType
	host   string
}

func (l *_LinkUnitReleasedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cascading.CascadeManager_3_0_0.LinkUnitReleasedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LinkUnitReleasedEvent) LinkId() int32 {
	return l.linkId
}

func (l *_LinkUnitReleasedEvent) Type() LinkUnitType {
	return l.type_
}

func (l *_LinkUnitReleasedEvent) Host() string {
	return l.host
}

func (l *_LinkUnitReleasedEvent) isLinkUnitReleasedEvent() {}

type _LinkUnitStatusChangedEvent struct {
	event.Event
	linkId    int32
	type_     LinkUnitType
	host      string
	oldStatus LinkUnitStatus
	newStatus LinkUnitStatus
}

func (l *_LinkUnitStatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cascading.CascadeManager_3_0_0.LinkUnitStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LinkUnitStatusChangedEvent) LinkId() int32 {
	return l.linkId
}

func (l *_LinkUnitStatusChangedEvent) Type() LinkUnitType {
	return l.type_
}

func (l *_LinkUnitStatusChangedEvent) Host() string {
	return l.host
}

func (l *_LinkUnitStatusChangedEvent) OldStatus() LinkUnitStatus {
	return l.oldStatus
}

func (l *_LinkUnitStatusChangedEvent) NewStatus() LinkUnitStatus {
	return l.newStatus
}

func (l *_LinkUnitStatusChangedEvent) isLinkUnitStatusChangedEvent() {}

type _LinkPortStatusChangedEvent struct {
	event.Event
	oldStatus LinkPortStatus
	newStatus LinkPortStatus
}

func (l *_LinkPortStatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cascading.CascadeManager_3_0_0.LinkPortStatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LinkPortStatusChangedEvent) OldStatus() LinkPortStatus {
	return l.oldStatus
}

func (l *_LinkPortStatusChangedEvent) NewStatus() LinkPortStatus {
	return l.newStatus
}

func (l *_LinkPortStatusChangedEvent) isLinkPortStatusChangedEvent() {}
