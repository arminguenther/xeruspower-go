// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package rawconfiguration

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() RawConfigDownloadedEvent { return &_RawConfigDownloadedEvent{} })
	valobj.Register(func() RawConfigUpdatedEvent { return &_RawConfigUpdatedEvent{} })
}

type _RawConfigDownloadedEvent struct {
	userevent.UserEvent
}

func (r *_RawConfigDownloadedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "rawcfg.RawConfiguration_1_0_1.RawConfigDownloadedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RawConfigDownloadedEvent) isRawConfigDownloadedEvent() {}

type _RawConfigUpdatedEvent struct {
	userevent.UserEvent
}

func (r *_RawConfigUpdatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "rawcfg.RawConfiguration_1_0_1.RawConfigUpdatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RawConfigUpdatedEvent) isRawConfigUpdatedEvent() {}
