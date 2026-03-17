// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package usermanager

import (
	"github.com/arminguenther/xeruspower-go/v40100/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() AccountAdded { return &_AccountAdded{} })
	valobj.Register(func() AccountChanged { return &_AccountChanged{} })
	valobj.Register(func() AccountEvent { return &_AccountEvent{} })
	valobj.Register(func() AccountRemoved { return &_AccountRemoved{} })
	valobj.Register(func() AccountRenamed { return &_AccountRenamed{} })
	valobj.Register(func() PasswordChanged { return &_PasswordChanged{} })
}

type _AccountEvent struct {
	userevent.UserEvent
	username string
}

func (a *_AccountEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.AccountEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AccountEvent) Username() string {
	return a.username
}

func (a *_AccountEvent) isAccountEvent() {}

type _AccountAdded struct {
	AccountEvent
}

func (a *_AccountAdded) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.AccountAdded",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AccountAdded) isAccountAdded() {}

type _AccountRenamed struct {
	AccountEvent
	newUsername string
}

func (a *_AccountRenamed) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.AccountRenamed",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AccountRenamed) NewUsername() string {
	return a.newUsername
}

func (a *_AccountRenamed) isAccountRenamed() {}

type _AccountRemoved struct {
	AccountEvent
}

func (a *_AccountRemoved) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.AccountRemoved",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AccountRemoved) isAccountRemoved() {}

type _PasswordChanged struct {
	AccountEvent
}

func (p *_PasswordChanged) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.PasswordChanged",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PasswordChanged) isPasswordChanged() {}

type _AccountChanged struct {
	AccountEvent
}

func (a *_AccountChanged) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "usermgmt.AccountChanged",
		Major: 3, Submajor: 0, Minor: 0,
	}
}

func (a *_AccountChanged) isAccountChanged() {}
