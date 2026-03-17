// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package keypad

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/idl/event"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() PINEnteredEvent { return &_PINEnteredEvent{} })
}

type _PINEnteredEvent struct {
	event.Event
	metaData MetaData
}

func (p *_PINEnteredEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.Keypad.PINEnteredEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PINEnteredEvent) MetaData() MetaData {
	return p.metaData
}

func (p *_PINEnteredEvent) isPINEnteredEvent() {}
