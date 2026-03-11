// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package panel

import (
	"github.com/arminguenther/xeruspower-go/v40410/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40410/pdumodel/circuit"
)

func (p *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["meterCount"] = p.MeterCount
	j0["panelSize"] = p.PanelSize
	j0["columns"] = p.Columns
	j0["labelingScheme"] = p.LabelingScheme
	return j0
}

func (p *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("meterCount", j0)
	if err != nil {
		return err
	}
	p.MeterCount, err = encoding.AsInt32(j0["meterCount"])
	if err != nil {
		return err
	}
	err = encoding.In("panelSize", j0)
	if err != nil {
		return err
	}
	p.PanelSize, err = encoding.AsInt32(j0["panelSize"])
	if err != nil {
		return err
	}
	err = encoding.In("columns", j0)
	if err != nil {
		return err
	}
	p.Columns, err = encoding.AsInt32(j0["columns"])
	if err != nil {
		return err
	}
	err = encoding.In("labelingScheme", j0)
	if err != nil {
		return err
	}
	p.LabelingScheme, err = encoding.AsEnum[LabelingScheme](j0["labelingScheme"])
	if err != nil {
		return err
	}
	return nil
}

func (p *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.UserEvent = valobj.For[userevent.UserEvent]()
	err := p.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = p.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = p.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CircuitCreatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.UserEvent = valobj.For[userevent.UserEvent]()
	err := c.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("circuit", value)
	if err != nil {
		return err
	}
	c.circuit, err = object.As[circuit.Circuit](value["circuit"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("config", value)
	if err != nil {
		return err
	}
	err = c.config.Decode(value["config"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = c.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CircuitDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.UserEvent = valobj.For[userevent.UserEvent]()
	err := c.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("config", value)
	if err != nil {
		return err
	}
	err = c.config.Decode(value["config"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = c.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}
