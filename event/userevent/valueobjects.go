// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package userevent

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/idl/event"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() UserEvent { return &_UserEvent{} })
}

type _UserEvent struct {
	event.Event
	actUserName string
	actIpAddr   string
}

func (u *_UserEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.UserEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (u *_UserEvent) ActUserName() string {
	return u.actUserName
}

func (u *_UserEvent) ActIpAddr() string {
	return u.actIpAddr
}

func (u *_UserEvent) isUserEvent() {}
