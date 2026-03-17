// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package panel

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
	circuit_ "github.com/arminguenther/xeruspower-go/pdumodel/circuit"
)

func init() {
	valobj.Register(func() CircuitCreatedEvent { return &_CircuitCreatedEvent{} })
	valobj.Register(func() CircuitDeletedEvent { return &_CircuitDeletedEvent{} })
	valobj.Register(func() SettingsChangedEvent { return &_SettingsChangedEvent{} })
}

type _SettingsChangedEvent struct {
	userevent.UserEvent
	oldSettings Settings
	newSettings Settings
}

func (p *_SettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Panel_2_0_1.PanelSettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_SettingsChangedEvent) OldSettings() Settings {
	return p.oldSettings
}

func (p *_SettingsChangedEvent) NewSettings() Settings {
	return p.newSettings
}

func (p *_SettingsChangedEvent) isSettingsChangedEvent() {}

type _CircuitCreatedEvent struct {
	userevent.UserEvent
	circuit  circuit_.Circuit
	config   circuit_.Config
	settings circuit_.Settings
}

func (c *_CircuitCreatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Panel_2_0_1.CircuitCreatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CircuitCreatedEvent) Circuit() circuit_.Circuit {
	return c.circuit
}

func (c *_CircuitCreatedEvent) Config() circuit_.Config {
	return c.config
}

func (c *_CircuitCreatedEvent) Settings() circuit_.Settings {
	return c.settings
}

func (c *_CircuitCreatedEvent) isCircuitCreatedEvent() {}

type _CircuitDeletedEvent struct {
	userevent.UserEvent
	config   circuit_.Config
	settings circuit_.Settings
}

func (c *_CircuitDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Panel_2_0_1.CircuitDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CircuitDeletedEvent) Config() circuit_.Config {
	return c.config
}

func (c *_CircuitDeletedEvent) Settings() circuit_.Settings {
	return c.settings
}

func (c *_CircuitDeletedEvent) isCircuitDeletedEvent() {}
