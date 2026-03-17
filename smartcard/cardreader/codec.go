// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cardreader

import (
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/idl/event"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["id"] = m.Id
	j0["manufacturer"] = m.Manufacturer
	j0["product"] = m.Product
	j0["serialNumber"] = m.SerialNumber
	j0["channel"] = m.Channel
	j0["position"] = m.Position
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	m.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("manufacturer", j0)
	if err != nil {
		return err
	}
	m.Manufacturer, err = encoding.Is[string](j0["manufacturer"])
	if err != nil {
		return err
	}
	err = encoding.In("product", j0)
	if err != nil {
		return err
	}
	m.Product, err = encoding.Is[string](j0["product"])
	if err != nil {
		return err
	}
	err = encoding.In("serialNumber", j0)
	if err != nil {
		return err
	}
	m.SerialNumber, err = encoding.Is[string](j0["serialNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("channel", j0)
	if err != nil {
		return err
	}
	m.Channel, err = encoding.AsInt32(j0["channel"])
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	m.Position, err = encoding.Is[string](j0["position"])
	if err != nil {
		return err
	}
	return nil
}

func (c *CardInformation) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["type"] = c.Type
	j0["uid"] = c.Uid
	return j0
}

func (c *CardInformation) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	c.Type, err = encoding.Is[string](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("uid", j0)
	if err != nil {
		return err
	}
	c.Uid, err = encoding.Is[string](j0["uid"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_CardEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("metaData", value)
	if err != nil {
		return err
	}
	err = c.metaData.Decode(value["metaData"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CardInsertedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.CardEvent = valobj.For[CardEvent]()
	err := c.CardEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CardRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.CardEvent = valobj.For[CardEvent]()
	err := c.CardEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
