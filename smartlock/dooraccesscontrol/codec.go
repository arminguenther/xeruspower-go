// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dooraccesscontrol

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/idl/event"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40000/peripheral/peripheraldeviceslot"
)

func (a *AbsoluteTimeCondition) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = a.Enabled
	j0["validFrom"] = a.ValidFrom.Unix()
	j0["validTill"] = a.ValidTill.Unix()
	return j0
}

func (a *AbsoluteTimeCondition) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	a.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("validFrom", j0)
	if err != nil {
		return err
	}
	a.ValidFrom, err = encoding.AsTime(j0["validFrom"])
	if err != nil {
		return err
	}
	err = encoding.In("validTill", j0)
	if err != nil {
		return err
	}
	a.ValidTill, err = encoding.AsTime(j0["validTill"])
	if err != nil {
		return err
	}
	return nil
}

func (p *PeriodicTimeCondition) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["enabled"] = p.Enabled
	j0["daysOfWeek"] = encoding.NonNilSlice(p.DaysOfWeek)
	j0["fromHourOfDay"] = p.FromHourOfDay
	j0["tillHourOfDay"] = p.TillHourOfDay
	j0["fromMinuteOfHour"] = p.FromMinuteOfHour
	j0["tillMinuteOfHour"] = p.TillMinuteOfHour
	return j0
}

func (p *PeriodicTimeCondition) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	p.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("daysOfWeek", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["daysOfWeek"])
	if err != nil {
		return err
	}
	p.DaysOfWeek = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		p.DaysOfWeek = append(p.DaysOfWeek, e1)
	}
	err = encoding.In("fromHourOfDay", j0)
	if err != nil {
		return err
	}
	p.FromHourOfDay, err = encoding.AsInt32(j0["fromHourOfDay"])
	if err != nil {
		return err
	}
	err = encoding.In("tillHourOfDay", j0)
	if err != nil {
		return err
	}
	p.TillHourOfDay, err = encoding.AsInt32(j0["tillHourOfDay"])
	if err != nil {
		return err
	}
	err = encoding.In("fromMinuteOfHour", j0)
	if err != nil {
		return err
	}
	p.FromMinuteOfHour, err = encoding.AsInt32(j0["fromMinuteOfHour"])
	if err != nil {
		return err
	}
	err = encoding.In("tillMinuteOfHour", j0)
	if err != nil {
		return err
	}
	p.TillMinuteOfHour, err = encoding.AsInt32(j0["tillMinuteOfHour"])
	if err != nil {
		return err
	}
	return nil
}

func (c *CardReaderInfo) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["linkId"] = c.LinkId
	j0["position"] = c.Position
	return j0
}

func (c *CardReaderInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("linkId", j0)
	if err != nil {
		return err
	}
	c.LinkId, err = encoding.AsInt32(j0["linkId"])
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	c.Position, err = encoding.Is[string](j0["position"])
	if err != nil {
		return err
	}
	return nil
}

func (k *KeypadInfo) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["linkId"] = k.LinkId
	j0["position"] = k.Position
	return j0
}

func (k *KeypadInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("linkId", j0)
	if err != nil {
		return err
	}
	k.LinkId, err = encoding.AsInt32(j0["linkId"])
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	k.Position, err = encoding.Is[string](j0["position"])
	if err != nil {
		return err
	}
	return nil
}

func (c *CardCondition) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = c.Enabled
	j0["cardUid"] = c.CardUid
	j0["cardReader"] = c.CardReader.Encode()
	return j0
}

func (c *CardCondition) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	c.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("cardUid", j0)
	if err != nil {
		return err
	}
	c.CardUid, err = encoding.Is[string](j0["cardUid"])
	if err != nil {
		return err
	}
	err = encoding.In("cardReader", j0)
	if err != nil {
		return err
	}
	err = c.CardReader.Decode(j0["cardReader"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (k *KeypadCondition) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = k.Enabled
	j0["pin"] = k.Pin
	j0["keypad"] = k.Keypad.Encode()
	return j0
}

func (k *KeypadCondition) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	k.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("pin", j0)
	if err != nil {
		return err
	}
	k.Pin, err = encoding.Is[string](j0["pin"])
	if err != nil {
		return err
	}
	err = encoding.In("keypad", j0)
	if err != nil {
		return err
	}
	err = k.Keypad.Decode(j0["keypad"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *DoorAccessRule) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["name"] = d.Name
	s1 := make([]any, 0, len(d.DoorHandleLocks))
	for _, e1 := range d.DoorHandleLocks {
		s1 = append(s1, object.ToMap(e1))
	}
	j0["doorHandleLocks"] = s1
	j0["cardCondition1"] = d.CardCondition1.Encode()
	j0["cardCondition2"] = d.CardCondition2.Encode()
	j0["keypadCondition1"] = d.KeypadCondition1.Encode()
	j0["keypadCondition2"] = d.KeypadCondition2.Encode()
	j0["conditionsTimeout"] = d.ConditionsTimeout
	j0["absoluteTime"] = d.AbsoluteTime.Encode()
	j0["periodicTime"] = d.PeriodicTime.Encode()
	return j0
}

func (d *DoorAccessRule) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	d.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("doorHandleLocks", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["doorHandleLocks"])
	if err != nil {
		return err
	}
	d.DoorHandleLocks = make([]peripheraldeviceslot.DeviceSlot, 0, len(s1))
	for _, a1 := range s1 {
		var e1 peripheraldeviceslot.DeviceSlot
		e1, err = object.As[peripheraldeviceslot.DeviceSlot](a1, caller)
		if err != nil {
			return err
		}
		d.DoorHandleLocks = append(d.DoorHandleLocks, e1)
	}
	err = encoding.In("cardCondition1", j0)
	if err != nil {
		return err
	}
	err = d.CardCondition1.Decode(j0["cardCondition1"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("cardCondition2", j0)
	if err != nil {
		return err
	}
	err = d.CardCondition2.Decode(j0["cardCondition2"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("keypadCondition1", j0)
	if err != nil {
		return err
	}
	err = d.KeypadCondition1.Decode(j0["keypadCondition1"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("keypadCondition2", j0)
	if err != nil {
		return err
	}
	err = d.KeypadCondition2.Decode(j0["keypadCondition2"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("conditionsTimeout", j0)
	if err != nil {
		return err
	}
	d.ConditionsTimeout, err = encoding.AsInt32(j0["conditionsTimeout"])
	if err != nil {
		return err
	}
	err = encoding.In("absoluteTime", j0)
	if err != nil {
		return err
	}
	err = d.AbsoluteTime.Decode(j0["absoluteTime"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("periodicTime", j0)
	if err != nil {
		return err
	}
	err = d.PeriodicTime.Decode(j0["periodicTime"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DoorAccessGrantedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ruleId", value)
	if err != nil {
		return err
	}
	d.ruleId, err = encoding.AsInt32(value["ruleId"])
	if err != nil {
		return err
	}
	err = encoding.In("rule", value)
	if err != nil {
		return err
	}
	err = d.rule.Decode(value["rule"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DoorAccessDeniedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("reason", value)
	if err != nil {
		return err
	}
	d.reason, err = encoding.AsEnum[DoorAccessDenialReason](value["reason"])
	if err != nil {
		return err
	}
	err = encoding.In("ruleId", value)
	if err != nil {
		return err
	}
	d.ruleId, err = encoding.AsInt32(value["ruleId"])
	if err != nil {
		return err
	}
	err = encoding.In("ruleName", value)
	if err != nil {
		return err
	}
	d.ruleName, err = encoding.Is[string](value["ruleName"])
	if err != nil {
		return err
	}
	return nil
}

func (d *_DoorAccessRuleAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.UserEvent = valobj.For[userevent.UserEvent]()
	err := d.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ruleId", value)
	if err != nil {
		return err
	}
	d.ruleId, err = encoding.AsInt32(value["ruleId"])
	if err != nil {
		return err
	}
	err = encoding.In("rule", value)
	if err != nil {
		return err
	}
	err = d.rule.Decode(value["rule"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DoorAccessRuleChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.UserEvent = valobj.For[userevent.UserEvent]()
	err := d.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ruleId", value)
	if err != nil {
		return err
	}
	d.ruleId, err = encoding.AsInt32(value["ruleId"])
	if err != nil {
		return err
	}
	err = encoding.In("oldRule", value)
	if err != nil {
		return err
	}
	err = d.oldRule.Decode(value["oldRule"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newRule", value)
	if err != nil {
		return err
	}
	err = d.newRule.Decode(value["newRule"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DoorAccessRuleDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.UserEvent = valobj.For[userevent.UserEvent]()
	err := d.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ruleId", value)
	if err != nil {
		return err
	}
	d.ruleId, err = encoding.AsInt32(value["ruleId"])
	if err != nil {
		return err
	}
	err = encoding.In("rule", value)
	if err != nil {
		return err
	}
	err = d.rule.Decode(value["rule"], caller)
	if err != nil {
		return err
	}
	return nil
}
