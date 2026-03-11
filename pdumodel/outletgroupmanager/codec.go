// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outletgroupmanager

import (
	"github.com/arminguenther/xeruspower-go/v40413/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40413/pdumodel/outletgroup"
)

func (g *_GroupCreatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	g.UserEvent = valobj.For[userevent.UserEvent]()
	err := g.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	g.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("uniqueId", value)
	if err != nil {
		return err
	}
	g.uniqueId, err = encoding.AsInt32(value["uniqueId"])
	if err != nil {
		return err
	}
	err = encoding.In("group", value)
	if err != nil {
		return err
	}
	g.group, err = object.As[outletgroup.OutletGroup](value["group"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = g.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (g *_GroupDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	g.UserEvent = valobj.For[userevent.UserEvent]()
	err := g.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	g.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("uniqueId", value)
	if err != nil {
		return err
	}
	g.uniqueId, err = encoding.AsInt32(value["uniqueId"])
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = g.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}
