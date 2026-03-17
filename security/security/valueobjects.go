// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package security

import (
	"github.com/arminguenther/xeruspower-go/v40100/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() FrontPanelPrivilegesChanged { return &_FrontPanelPrivilegesChanged{} })
	valobj.Register(func() PasswordSettingsChanged { return &_PasswordSettingsChanged{} })
}

type _PasswordSettingsChanged struct {
	userevent.UserEvent
	oldSettings PasswordSettings
	newSettings PasswordSettings
}

func (p *_PasswordSettingsChanged) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "security.PasswordSettingsChanged",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PasswordSettingsChanged) OldSettings() PasswordSettings {
	return p.oldSettings
}

func (p *_PasswordSettingsChanged) NewSettings() PasswordSettings {
	return p.newSettings
}

func (p *_PasswordSettingsChanged) isPasswordSettingsChanged() {}

type _FrontPanelPrivilegesChanged struct {
	userevent.UserEvent
	oldPrivileges []string
	newPrivileges []string
}

func (f *_FrontPanelPrivilegesChanged) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "security.FrontPanelPrivilegesChanged",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FrontPanelPrivilegesChanged) OldPrivileges() []string {
	return f.oldPrivileges
}

func (f *_FrontPanelPrivilegesChanged) NewPrivileges() []string {
	return f.newPrivileges
}

func (f *_FrontPanelPrivilegesChanged) isFrontPanelPrivilegesChanged() {}
