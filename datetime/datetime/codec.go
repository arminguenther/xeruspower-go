// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datetime

import (
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

func (z *ZoneInfo) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["id"] = z.Id
	j0["name"] = z.Name
	j0["hasDSTInfo"] = z.HasDSTInfo
	return j0
}

func (z *ZoneInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	z.Id, err = encoding.AsInt32(j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	z.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("hasDSTInfo", j0)
	if err != nil {
		return err
	}
	z.HasDSTInfo, err = encoding.Is[bool](j0["hasDSTInfo"])
	if err != nil {
		return err
	}
	return nil
}

func (z *ZoneCfg) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["id"] = z.Id
	j0["name"] = z.Name
	j0["enableAutoDST"] = z.EnableAutoDST
	return j0
}

func (z *ZoneCfg) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	z.Id, err = encoding.AsInt32(j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	z.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("enableAutoDST", j0)
	if err != nil {
		return err
	}
	z.EnableAutoDST, err = encoding.Is[bool](j0["enableAutoDST"])
	if err != nil {
		return err
	}
	return nil
}

func (n *NtpCfg) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["server1"] = n.Server1
	j0["server2"] = n.Server2
	return j0
}

func (n *NtpCfg) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("server1", j0)
	if err != nil {
		return err
	}
	n.Server1, err = encoding.Is[string](j0["server1"])
	if err != nil {
		return err
	}
	err = encoding.In("server2", j0)
	if err != nil {
		return err
	}
	n.Server2, err = encoding.Is[string](j0["server2"])
	if err != nil {
		return err
	}
	return nil
}

func (c *Cfg) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["zoneCfg"] = c.ZoneCfg.Encode()
	j0["protocol"] = c.Protocol
	j0["deviceTime"] = c.DeviceTime.Unix()
	j0["ntpCfg"] = c.NtpCfg.Encode()
	return j0
}

func (c *Cfg) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("zoneCfg", j0)
	if err != nil {
		return err
	}
	err = c.ZoneCfg.Decode(j0["zoneCfg"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("protocol", j0)
	if err != nil {
		return err
	}
	c.Protocol, err = encoding.AsEnum[Protocol](j0["protocol"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceTime", j0)
	if err != nil {
		return err
	}
	c.DeviceTime, err = encoding.AsTime(j0["deviceTime"])
	if err != nil {
		return err
	}
	err = encoding.In("ntpCfg", j0)
	if err != nil {
		return err
	}
	err = c.NtpCfg.Decode(j0["ntpCfg"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_ConfigurationChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_ClockChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldTime", value)
	if err != nil {
		return err
	}
	c.oldTime, err = encoding.AsTime(value["oldTime"])
	if err != nil {
		return err
	}
	err = encoding.In("newTime", value)
	if err != nil {
		return err
	}
	c.newTime, err = encoding.AsTime(value["newTime"])
	if err != nil {
		return err
	}
	return nil
}
