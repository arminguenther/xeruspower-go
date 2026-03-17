// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package unit

import (
	"github.com/arminguenther/xeruspower-go/v40211/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() IdentificationStartedEvent { return &_IdentificationStartedEvent{} })
}

type _IdentificationStartedEvent struct {
	userevent.UserEvent
	duration int32
}

func (i *_IdentificationStartedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Unit_2_0_1.IdentificationStartedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_IdentificationStartedEvent) Duration() int32 {
	return i.duration
}

func (i *_IdentificationStartedEvent) isIdentificationStartedEvent() {}
