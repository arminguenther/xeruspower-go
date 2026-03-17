// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cardreadermanager

import (
	"github.com/arminguenther/xeruspower-go/v40020/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/idl/event"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40020/smartcard/cardreader"
)

func init() {
	valobj.Register(func() CardReaderAttachedEvent { return &_CardReaderAttachedEvent{} })
	valobj.Register(func() CardReaderDetachedEvent { return &_CardReaderDetachedEvent{} })
	valobj.Register(func() CardReaderEvent { return &_CardReaderEvent{} })
	valobj.Register(func() CardReaderSettingsChangedEvent { return &_CardReaderSettingsChangedEvent{} })
}

type _CardReaderEvent struct {
	event.Event
	cardReader cardreader.CardReader
	metaData   cardreader.MetaData
}

func (c *_CardReaderEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReaderManager_1_0_5.CardReaderEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CardReaderEvent) CardReader() cardreader.CardReader {
	return c.cardReader
}

func (c *_CardReaderEvent) MetaData() cardreader.MetaData {
	return c.metaData
}

func (c *_CardReaderEvent) isCardReaderEvent() {}

type _CardReaderAttachedEvent struct {
	CardReaderEvent
}

func (c *_CardReaderAttachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReaderManager_1_0_5.CardReaderAttachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CardReaderAttachedEvent) isCardReaderAttachedEvent() {}

type _CardReaderDetachedEvent struct {
	CardReaderEvent
}

func (c *_CardReaderDetachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReaderManager_1_0_5.CardReaderDetachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CardReaderDetachedEvent) isCardReaderDetachedEvent() {}

type _CardReaderSettingsChangedEvent struct {
	userevent.UserEvent
	cardReader  cardreader.CardReader
	oldSettings CardReaderSettings
	newSettings CardReaderSettings
	position    string
}

func (c *_CardReaderSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReaderManager_1_0_5.CardReaderSettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CardReaderSettingsChangedEvent) CardReader() cardreader.CardReader {
	return c.cardReader
}

func (c *_CardReaderSettingsChangedEvent) OldSettings() CardReaderSettings {
	return c.oldSettings
}

func (c *_CardReaderSettingsChangedEvent) NewSettings() CardReaderSettings {
	return c.newSettings
}

func (c *_CardReaderSettingsChangedEvent) Position() string {
	return c.position
}

func (c *_CardReaderSettingsChangedEvent) isCardReaderSettingsChangedEvent() {}
