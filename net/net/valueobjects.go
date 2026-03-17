// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package net

import (
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/idl/event"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() CommonInfoChangedEvent { return &_CommonInfoChangedEvent{} })
	valobj.Register(func() EthInfoChangedEvent { return &_EthInfoChangedEvent{} })
	valobj.Register(func() InterfaceInfoChangedEvent { return &_InterfaceInfoChangedEvent{} })
	valobj.Register(func() LinkStateChangedEvent { return &_LinkStateChangedEvent{} })
	valobj.Register(func() PortForwardingExpansionUnitPresenceStateChangedEvent {
		return &_PortForwardingExpansionUnitPresenceStateChangedEvent{}
	})
	valobj.Register(func() WlanInfoChangedEvent { return &_WlanInfoChangedEvent{} })
}

type _CommonInfoChangedEvent struct {
	event.Event
	commonInfo CommonInfo
}

func (c *_CommonInfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.CommonInfoChangedEvent",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (c *_CommonInfoChangedEvent) CommonInfo() CommonInfo {
	return c.commonInfo
}

func (c *_CommonInfoChangedEvent) isCommonInfoChangedEvent() {}

type _InterfaceInfoChangedEvent struct {
	event.Event
	ifInfo InterfaceInfo
}

func (i *_InterfaceInfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.InterfaceInfoChangedEvent",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (i *_InterfaceInfoChangedEvent) IfInfo() InterfaceInfo {
	return i.ifInfo
}

func (i *_InterfaceInfoChangedEvent) isInterfaceInfoChangedEvent() {}

type _EthInfoChangedEvent struct {
	event.Event
	ifName  string
	ifLabel string
	ethInfo EthInfo
}

func (e *_EthInfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.EthInfoChangedEvent",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (e *_EthInfoChangedEvent) IfName() string {
	return e.ifName
}

func (e *_EthInfoChangedEvent) IfLabel() string {
	return e.ifLabel
}

func (e *_EthInfoChangedEvent) EthInfo() EthInfo {
	return e.ethInfo
}

func (e *_EthInfoChangedEvent) isEthInfoChangedEvent() {}

type _WlanInfoChangedEvent struct {
	event.Event
	ifName   string
	wlanInfo WlanInfo
}

func (w *_WlanInfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.WlanInfoChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (w *_WlanInfoChangedEvent) IfName() string {
	return w.ifName
}

func (w *_WlanInfoChangedEvent) WlanInfo() WlanInfo {
	return w.wlanInfo
}

func (w *_WlanInfoChangedEvent) isWlanInfoChangedEvent() {}

type _LinkStateChangedEvent struct {
	event.Event
	ifName  string
	ifLabel string
	ifType  InterfaceType
	ifState InterfaceOpState
}

func (l *_LinkStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.LinkStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (l *_LinkStateChangedEvent) IfName() string {
	return l.ifName
}

func (l *_LinkStateChangedEvent) IfLabel() string {
	return l.ifLabel
}

func (l *_LinkStateChangedEvent) IfType() InterfaceType {
	return l.ifType
}

func (l *_LinkStateChangedEvent) IfState() InterfaceOpState {
	return l.ifState
}

func (l *_LinkStateChangedEvent) isLinkStateChangedEvent() {}

type _PortForwardingExpansionUnitPresenceStateChangedEvent struct {
	event.Event
	expansionUnitPresent bool
}

func (p *_PortForwardingExpansionUnitPresenceStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.PortForwardingExpansionUnitPresenceStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PortForwardingExpansionUnitPresenceStateChangedEvent) ExpansionUnitPresent() bool {
	return p.expansionUnitPresent
}

func (p *_PortForwardingExpansionUnitPresenceStateChangedEvent) isPortForwardingExpansionUnitPresenceStateChangedEvent() {
}
