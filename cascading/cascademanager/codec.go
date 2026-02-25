// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cascademanager

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func (p *PrimaryUnitSettings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["caCertChain"] = p.CaCertChain
	j0["allowOffTimeRangeCerts"] = p.AllowOffTimeRangeCerts
	return j0
}

func (p *PrimaryUnitSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("caCertChain", j0)
	if err != nil {
		return err
	}
	p.CaCertChain, err = encoding.Is[string](j0["caCertChain"])
	if err != nil {
		return err
	}
	err = encoding.In("allowOffTimeRangeCerts", j0)
	if err != nil {
		return err
	}
	p.AllowOffTimeRangeCerts, err = encoding.Is[bool](j0["allowOffTimeRangeCerts"])
	if err != nil {
		return err
	}
	return nil
}

func (l *LinkUnit) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["type"] = l.Type
	j0["host"] = l.Host
	j0["status"] = l.Status
	j0["fwVersion"] = l.FwVersion
	return j0
}

func (l *LinkUnit) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	l.Type, err = encoding.AsEnum[LinkUnitType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("host", j0)
	if err != nil {
		return err
	}
	l.Host, err = encoding.Is[string](j0["host"])
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	l.Status, err = encoding.AsEnum[LinkUnitStatus](j0["status"])
	if err != nil {
		return err
	}
	err = encoding.In("fwVersion", j0)
	if err != nil {
		return err
	}
	l.FwVersion, err = encoding.Is[string](j0["fwVersion"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Status) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["role"] = s.Role
	j0["primaryUnit"] = s.PrimaryUnit
	j0["primaryLinkId"] = s.PrimaryLinkId
	s1 := make([]any, 0, len(s.LinkUnits))
	for k1, v1 := range s.LinkUnits {
		s1 = append(s1, map[string]any{"key": k1, "value": v1.Encode()})
	}
	j0["linkUnits"] = s1
	return j0
}

func (s *Status) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("role", j0)
	if err != nil {
		return err
	}
	s.Role, err = encoding.AsEnum[Role](j0["role"])
	if err != nil {
		return err
	}
	err = encoding.In("primaryUnit", j0)
	if err != nil {
		return err
	}
	s.PrimaryUnit, err = encoding.Is[string](j0["primaryUnit"])
	if err != nil {
		return err
	}
	err = encoding.In("primaryLinkId", j0)
	if err != nil {
		return err
	}
	s.PrimaryLinkId, err = encoding.AsInt32(j0["primaryLinkId"])
	if err != nil {
		return err
	}
	err = encoding.In("linkUnits", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["linkUnits"])
	s.LinkUnits = make(map[int32]LinkUnit, l1)
	for a1, b1 := range i1 {
		var k1 int32
		k1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		var v1 LinkUnit
		err = v1.Decode(b1, caller)
		if err != nil {
			return err
		}
		s.LinkUnits[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	return nil
}

func (l *LinkPortStatus) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["isSupported"] = l.IsSupported
	j0["isLinkDetected"] = l.IsLinkDetected
	j0["isLinkingConfirmationNeeded"] = l.IsLinkingConfirmationNeeded
	j0["connectedNeighborAddr"] = l.ConnectedNeighborAddr
	return j0
}

func (l *LinkPortStatus) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("isSupported", j0)
	if err != nil {
		return err
	}
	l.IsSupported, err = encoding.Is[bool](j0["isSupported"])
	if err != nil {
		return err
	}
	err = encoding.In("isLinkDetected", j0)
	if err != nil {
		return err
	}
	l.IsLinkDetected, err = encoding.Is[bool](j0["isLinkDetected"])
	if err != nil {
		return err
	}
	err = encoding.In("isLinkingConfirmationNeeded", j0)
	if err != nil {
		return err
	}
	l.IsLinkingConfirmationNeeded, err = encoding.Is[bool](j0["isLinkingConfirmationNeeded"])
	if err != nil {
		return err
	}
	err = encoding.In("connectedNeighborAddr", j0)
	if err != nil {
		return err
	}
	l.ConnectedNeighborAddr, err = encoding.Is[string](j0["connectedNeighborAddr"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RoleChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.Event = valobj.For[event.Event]()
	err := r.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldRole", value)
	if err != nil {
		return err
	}
	r.oldRole, err = encoding.AsEnum[Role](value["oldRole"])
	if err != nil {
		return err
	}
	err = encoding.In("newRole", value)
	if err != nil {
		return err
	}
	r.newRole, err = encoding.AsEnum[Role](value["newRole"])
	if err != nil {
		return err
	}
	err = encoding.In("primaryUnit", value)
	if err != nil {
		return err
	}
	r.primaryUnit, err = encoding.Is[string](value["primaryUnit"])
	if err != nil {
		return err
	}
	return nil
}

func (l *_LinkUnitAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.UserEvent = valobj.For[userevent.UserEvent]()
	err := l.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("linkId", value)
	if err != nil {
		return err
	}
	l.linkId, err = encoding.AsInt32(value["linkId"])
	if err != nil {
		return err
	}
	err = encoding.In("type", value)
	if err != nil {
		return err
	}
	l.type_, err = encoding.AsEnum[LinkUnitType](value["type"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	l.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	return nil
}

func (l *_LinkUnitReleasedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.UserEvent = valobj.For[userevent.UserEvent]()
	err := l.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("linkId", value)
	if err != nil {
		return err
	}
	l.linkId, err = encoding.AsInt32(value["linkId"])
	if err != nil {
		return err
	}
	err = encoding.In("type", value)
	if err != nil {
		return err
	}
	l.type_, err = encoding.AsEnum[LinkUnitType](value["type"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	l.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	return nil
}

func (l *_LinkUnitStatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.Event = valobj.For[event.Event]()
	err := l.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("linkId", value)
	if err != nil {
		return err
	}
	l.linkId, err = encoding.AsInt32(value["linkId"])
	if err != nil {
		return err
	}
	err = encoding.In("type", value)
	if err != nil {
		return err
	}
	l.type_, err = encoding.AsEnum[LinkUnitType](value["type"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	l.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	err = encoding.In("oldStatus", value)
	if err != nil {
		return err
	}
	l.oldStatus, err = encoding.AsEnum[LinkUnitStatus](value["oldStatus"])
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	l.newStatus, err = encoding.AsEnum[LinkUnitStatus](value["newStatus"])
	if err != nil {
		return err
	}
	return nil
}

func (l *_LinkPortStatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.Event = valobj.For[event.Event]()
	err := l.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldStatus", value)
	if err != nil {
		return err
	}
	err = l.oldStatus.Decode(value["oldStatus"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	err = l.newStatus.Decode(value["newStatus"], caller)
	if err != nil {
		return err
	}
	return nil
}
