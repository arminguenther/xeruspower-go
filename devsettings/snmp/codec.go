// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package snmp

import (
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/idl/event"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
)

func (c *Configuration) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["v2enable"] = c.V2enable
	j0["v3enable"] = c.V3enable
	j0["readComm"] = c.ReadComm
	j0["writeComm"] = c.WriteComm
	j0["sysContact"] = c.SysContact
	j0["sysName"] = c.SysName
	j0["sysLocation"] = c.SysLocation
	return j0
}

func (c *Configuration) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("v2enable", j0)
	if err != nil {
		return err
	}
	c.V2enable, err = encoding.Is[bool](j0["v2enable"])
	if err != nil {
		return err
	}
	err = encoding.In("v3enable", j0)
	if err != nil {
		return err
	}
	c.V3enable, err = encoding.Is[bool](j0["v3enable"])
	if err != nil {
		return err
	}
	err = encoding.In("readComm", j0)
	if err != nil {
		return err
	}
	c.ReadComm, err = encoding.Is[string](j0["readComm"])
	if err != nil {
		return err
	}
	err = encoding.In("writeComm", j0)
	if err != nil {
		return err
	}
	c.WriteComm, err = encoding.Is[string](j0["writeComm"])
	if err != nil {
		return err
	}
	err = encoding.In("sysContact", j0)
	if err != nil {
		return err
	}
	c.SysContact, err = encoding.Is[string](j0["sysContact"])
	if err != nil {
		return err
	}
	err = encoding.In("sysName", j0)
	if err != nil {
		return err
	}
	c.SysName, err = encoding.Is[string](j0["sysName"])
	if err != nil {
		return err
	}
	err = encoding.In("sysLocation", j0)
	if err != nil {
		return err
	}
	c.SysLocation, err = encoding.Is[string](j0["sysLocation"])
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
	err = encoding.In("userName", value)
	if err != nil {
		return err
	}
	c.userName, err = encoding.Is[string](value["userName"])
	if err != nil {
		return err
	}
	err = encoding.In("ipAddr", value)
	if err != nil {
		return err
	}
	c.ipAddr, err = encoding.Is[string](value["ipAddr"])
	if err != nil {
		return err
	}
	err = encoding.In("oldConfig", value)
	if err != nil {
		return err
	}
	err = c.oldConfig.Decode(value["oldConfig"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newConfig", value)
	if err != nil {
		return err
	}
	err = c.newConfig.Decode(value["newConfig"], caller)
	if err != nil {
		return err
	}
	return nil
}
