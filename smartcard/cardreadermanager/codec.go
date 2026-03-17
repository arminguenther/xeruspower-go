// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cardreadermanager

import (
	"github.com/arminguenther/xeruspower-go/v40100/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/idl/event"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40100/smartcard/cardreader"
)

func (c *CardReaderSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["name"] = c.Name
	j0["description"] = c.Description
	j0["cardFormat"] = c.CardFormat
	return j0
}

func (c *CardReaderSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	c.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	c.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	err = encoding.In("cardFormat", j0)
	if err != nil {
		return err
	}
	c.CardFormat, err = encoding.Is[string](j0["cardFormat"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_CardReaderEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("cardReader", value)
	if err != nil {
		return err
	}
	c.cardReader, err = object.As[cardreader.CardReader](value["cardReader"], caller)
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

func (c *_CardReaderAttachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.CardReaderEvent = valobj.For[CardReaderEvent]()
	err := c.CardReaderEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CardReaderDetachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.CardReaderEvent = valobj.For[CardReaderEvent]()
	err := c.CardReaderEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CardReaderSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.UserEvent = valobj.For[userevent.UserEvent]()
	err := c.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("cardReader", value)
	if err != nil {
		return err
	}
	c.cardReader, err = object.As[cardreader.CardReader](value["cardReader"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = c.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = c.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("position", value)
	if err != nil {
		return err
	}
	c.position, err = encoding.Is[string](value["position"])
	if err != nil {
		return err
	}
	return nil
}
