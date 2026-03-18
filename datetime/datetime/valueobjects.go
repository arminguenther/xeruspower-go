// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datetime

import (
	"time"

	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/idl/event"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ClockChangedEvent { return &_ClockChangedEvent{} })
	valobj.Register(func() ConfigurationChangedEvent { return &_ConfigurationChangedEvent{} })
}

type _ConfigurationChangedEvent struct {
	event.Event
}

func (c *_ConfigurationChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "datetime.DateTime_3_0_3.ConfigurationChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ConfigurationChangedEvent) isConfigurationChangedEvent() {}

type _ClockChangedEvent struct {
	event.Event
	oldTime time.Time
	newTime time.Time
}

func (c *_ClockChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "datetime.DateTime_3_0_3.ClockChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ClockChangedEvent) OldTime() time.Time {
	return c.oldTime
}

func (c *_ClockChangedEvent) NewTime() time.Time {
	return c.newTime
}

func (c *_ClockChangedEvent) isClockChangedEvent() {}
