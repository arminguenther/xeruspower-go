// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package controller

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() MetaDataChangedEvent { return &_MetaDataChangedEvent{} })
	valobj.Register(func() StatusChangedEvent { return &_StatusChangedEvent{} })
}

type _StatusChangedEvent struct {
	event.Event
	oldStatus Status
	newStatus Status
}

func (s *_StatusChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Controller_5_0_1.StatusChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StatusChangedEvent) OldStatus() Status {
	return s.oldStatus
}

func (s *_StatusChangedEvent) NewStatus() Status {
	return s.newStatus
}

func (s *_StatusChangedEvent) isStatusChangedEvent() {}

type _MetaDataChangedEvent struct {
	event.Event
	oldMetaData MetaData
	newMetaData MetaData
}

func (m *_MetaDataChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Controller_5_0_1.MetaDataChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MetaDataChangedEvent) OldMetaData() MetaData {
	return m.oldMetaData
}

func (m *_MetaDataChangedEvent) NewMetaData() MetaData {
	return m.newMetaData
}

func (m *_MetaDataChangedEvent) isMetaDataChangedEvent() {}
