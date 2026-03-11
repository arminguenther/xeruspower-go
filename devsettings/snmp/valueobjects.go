// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package snmp

import (
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/idl/event"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ConfigurationChangedEvent { return &_ConfigurationChangedEvent{} })
}

type _ConfigurationChangedEvent struct {
	event.Event
	userName  string
	ipAddr    string
	oldConfig Configuration
	newConfig Configuration
}

func (c *_ConfigurationChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Snmp_1_0_2.ConfigurationChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ConfigurationChangedEvent) UserName() string {
	return c.userName
}

func (c *_ConfigurationChangedEvent) IpAddr() string {
	return c.ipAddr
}

func (c *_ConfigurationChangedEvent) OldConfig() Configuration {
	return c.oldConfig
}

func (c *_ConfigurationChangedEvent) NewConfig() Configuration {
	return c.newConfig
}

func (c *_ConfigurationChangedEvent) isConfigurationChangedEvent() {}
