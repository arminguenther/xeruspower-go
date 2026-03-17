// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sx

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ClientConnectionStatusEvent { return &_ClientConnectionStatusEvent{} })
}

type _ClientConnectionStatusEvent struct {
	userevent.UserEvent
	clientInfo ClientInfo
	status     bool
	clientCnt  int32
}

func (c *_ClientConnectionStatusEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sx.Sx.ClientConnectionStatusEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ClientConnectionStatusEvent) ClientInfo() ClientInfo {
	return c.clientInfo
}

func (c *_ClientConnectionStatusEvent) Status() bool {
	return c.status
}

func (c *_ClientConnectionStatusEvent) ClientCnt() int32 {
	return c.clientCnt
}

func (c *_ClientConnectionStatusEvent) isClientConnectionStatusEvent() {}
