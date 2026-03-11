// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sx

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func (c *ClientInfo) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["portId"] = c.PortId
	j0["portName"] = c.PortName
	j0["userName"] = c.UserName
	j0["userIp"] = c.UserIp
	return j0
}

func (c *ClientInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("portId", j0)
	if err != nil {
		return err
	}
	c.PortId, err = encoding.Is[string](j0["portId"])
	if err != nil {
		return err
	}
	err = encoding.In("portName", j0)
	if err != nil {
		return err
	}
	c.PortName, err = encoding.Is[string](j0["portName"])
	if err != nil {
		return err
	}
	err = encoding.In("userName", j0)
	if err != nil {
		return err
	}
	c.UserName, err = encoding.Is[string](j0["userName"])
	if err != nil {
		return err
	}
	err = encoding.In("userIp", j0)
	if err != nil {
		return err
	}
	c.UserIp, err = encoding.Is[string](j0["userIp"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_ClientConnectionStatusEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.UserEvent = valobj.For[userevent.UserEvent]()
	err := c.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("clientInfo", value)
	if err != nil {
		return err
	}
	err = c.clientInfo.Decode(value["clientInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("status", value)
	if err != nil {
		return err
	}
	c.status, err = encoding.Is[bool](value["status"])
	if err != nil {
		return err
	}
	err = encoding.In("clientCnt", value)
	if err != nil {
		return err
	}
	c.clientCnt, err = encoding.AsInt32(value["clientCnt"])
	if err != nil {
		return err
	}
	return nil
}
