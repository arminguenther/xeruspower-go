// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cardreader

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() CardEvent { return &_CardEvent{} })
	valobj.Register(func() CardInsertedEvent { return &_CardInsertedEvent{} })
	valobj.Register(func() CardRemovedEvent { return &_CardRemovedEvent{} })
}

type _CardEvent struct {
	event.Event
	metaData MetaData
}

func (c *_CardEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReader_1_0_4.CardEvent",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (c *_CardEvent) MetaData() MetaData {
	return c.metaData
}

func (c *_CardEvent) isCardEvent() {}

type _CardInsertedEvent struct {
	CardEvent
}

func (c *_CardInsertedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReader_1_0_4.CardInsertedEvent",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (c *_CardInsertedEvent) isCardInsertedEvent() {}

type _CardRemovedEvent struct {
	CardEvent
}

func (c *_CardRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReader_1_0_4.CardRemovedEvent",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (c *_CardRemovedEvent) isCardRemovedEvent() {}
