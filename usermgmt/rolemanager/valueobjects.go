// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package rolemanager

import (
	"github.com/arminguenther/xeruspower-go/v40413/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40413/usermgmt/role"
)

func init() {
	valobj.Register(func() RoleAdded { return &_RoleAdded{} })
	valobj.Register(func() RoleChanged { return &_RoleChanged{} })
	valobj.Register(func() RoleEvent { return &_RoleEvent{} })
	valobj.Register(func() RoleRemoved { return &_RoleRemoved{} })
}

type _RoleEvent struct {
	userevent.UserEvent
	rolename string
}

func (r *_RoleEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.RoleEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RoleEvent) Rolename() string {
	return r.rolename
}

func (r *_RoleEvent) isRoleEvent() {}

type _RoleAdded struct {
	RoleEvent
}

func (r *_RoleAdded) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.RoleAdded",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RoleAdded) isRoleAdded() {}

type _RoleRemoved struct {
	RoleEvent
}

func (r *_RoleRemoved) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.RoleRemoved",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RoleRemoved) isRoleRemoved() {}

type _RoleChanged struct {
	RoleEvent
	oldSettings role.Info
	newSettings role.Info
}

func (r *_RoleChanged) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.RoleChanged",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RoleChanged) OldSettings() role.Info {
	return r.oldSettings
}

func (r *_RoleChanged) NewSettings() role.Info {
	return r.newSettings
}

func (r *_RoleChanged) isRoleChanged() {}
