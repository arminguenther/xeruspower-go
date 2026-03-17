// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datetime

import (
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/idl/event"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ConfigurationChangedEvent { return &_ConfigurationChangedEvent{} })
}

type _ConfigurationChangedEvent struct {
	event.Event
}

func (c *_ConfigurationChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "datetime.DateTime_3_0_2.ConfigurationChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ConfigurationChangedEvent) isConfigurationChangedEvent() {}
