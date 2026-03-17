// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventlog

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ClearedEvent { return &_ClearedEvent{} })
}

type _ClearedEvent struct {
	userevent.UserEvent
}

func (e *_ClearedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "logging.EventLogClearedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_ClearedEvent) isClearedEvent() {}
