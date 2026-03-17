// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outletgroupmanager

import (
	"github.com/arminguenther/xeruspower-go/v40220/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/outletgroup"
)

func init() {
	valobj.Register(func() GroupCreatedEvent { return &_GroupCreatedEvent{} })
	valobj.Register(func() GroupDeletedEvent { return &_GroupDeletedEvent{} })
}

type _GroupCreatedEvent struct {
	userevent.UserEvent
	id       int32
	uniqueId int32
	group    outletgroup.OutletGroup
	settings outletgroup.Settings
}

func (g *_GroupCreatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroupManager_1_1_9.GroupCreatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (g *_GroupCreatedEvent) Id() int32 {
	return g.id
}

func (g *_GroupCreatedEvent) UniqueId() int32 {
	return g.uniqueId
}

func (g *_GroupCreatedEvent) Group() outletgroup.OutletGroup {
	return g.group
}

func (g *_GroupCreatedEvent) Settings() outletgroup.Settings {
	return g.settings
}

func (g *_GroupCreatedEvent) isGroupCreatedEvent() {}

type _GroupDeletedEvent struct {
	userevent.UserEvent
	id       int32
	uniqueId int32
	settings outletgroup.Settings
}

func (g *_GroupDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OutletGroupManager_1_1_9.GroupDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (g *_GroupDeletedEvent) Id() int32 {
	return g.id
}

func (g *_GroupDeletedEvent) UniqueId() int32 {
	return g.uniqueId
}

func (g *_GroupDeletedEvent) Settings() outletgroup.Settings {
	return g.settings
}

func (g *_GroupDeletedEvent) isGroupDeletedEvent() {}
