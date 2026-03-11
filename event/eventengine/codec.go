// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventengine

import (
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
)

func (k *KeyValue) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["key"] = k.Key
	j0["value"] = k.Value
	return j0
}

func (k *KeyValue) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("key", j0)
	if err != nil {
		return err
	}
	k.Key, err = encoding.Is[string](j0["key"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	k.Value, err = encoding.Is[string](j0["value"])
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["type"] = e.Type
	j0["id"] = encoding.NonNilSlice(e.Id)
	j0["asserted"] = e.Asserted
	j0["timeStamp"] = e.TimeStamp.Unix()
	s1 := make([]any, 0, len(e.Context))
	for _, e1 := range e.Context {
		s1 = append(s1, e1.Encode())
	}
	j0["context"] = s1
	return j0
}

func (e *Event) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	e.Type, err = encoding.AsEnum[EventType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["id"])
	if err != nil {
		return err
	}
	e.Id = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		e.Id = append(e.Id, e1)
	}
	err = encoding.In("asserted", j0)
	if err != nil {
		return err
	}
	e.Asserted, err = encoding.Is[bool](j0["asserted"])
	if err != nil {
		return err
	}
	err = encoding.In("timeStamp", j0)
	if err != nil {
		return err
	}
	e.TimeStamp, err = encoding.AsTime(j0["timeStamp"])
	if err != nil {
		return err
	}
	err = encoding.In("context", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["context"])
	if err != nil {
		return err
	}
	e.Context = make([]KeyValue, 0, len(s2))
	for _, a2 := range s2 {
		var e2 KeyValue
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		e.Context = append(e.Context, e2)
	}
	return nil
}

func (e *EngineEventDesc) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["eventDescType"] = e.EventDescType
	j0["eventType"] = e.EventType
	j0["dynNodeContext"] = e.DynNodeContext
	j0["idComp"] = e.IdComp
	j0["name"] = e.Name
	s1 := make([]any, 0, len(e.Entries))
	for _, e1 := range e.Entries {
		s1 = append(s1, e1.Encode())
	}
	j0["entries"] = s1
	return j0
}

func (e *EngineEventDesc) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("eventDescType", j0)
	if err != nil {
		return err
	}
	e.EventDescType, err = encoding.AsEnum[EngineEventDescType](j0["eventDescType"])
	if err != nil {
		return err
	}
	err = encoding.In("eventType", j0)
	if err != nil {
		return err
	}
	e.EventType, err = encoding.AsEnum[EventType](j0["eventType"])
	if err != nil {
		return err
	}
	err = encoding.In("dynNodeContext", j0)
	if err != nil {
		return err
	}
	e.DynNodeContext, err = encoding.Is[string](j0["dynNodeContext"])
	if err != nil {
		return err
	}
	err = encoding.In("idComp", j0)
	if err != nil {
		return err
	}
	e.IdComp, err = encoding.Is[string](j0["idComp"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	e.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("entries", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["entries"])
	if err != nil {
		return err
	}
	e.Entries = make([]EngineEventDesc, 0, len(s1))
	for _, a1 := range s1 {
		var e1 EngineEventDesc
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		e.Entries = append(e.Entries, e1)
	}
	return nil
}

func (a *EngineAction) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["id"] = a.Id
	j0["name"] = a.Name
	j0["isSystem"] = a.IsSystem
	j0["type"] = a.Type
	s1 := make([]any, 0, len(a.Arguments))
	for _, e1 := range a.Arguments {
		s1 = append(s1, e1.Encode())
	}
	j0["arguments"] = s1
	return j0
}

func (a *EngineAction) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	a.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	a.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("isSystem", j0)
	if err != nil {
		return err
	}
	a.IsSystem, err = encoding.Is[bool](j0["isSystem"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	a.Type, err = encoding.Is[string](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("arguments", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["arguments"])
	if err != nil {
		return err
	}
	a.Arguments = make([]KeyValue, 0, len(s1))
	for _, a1 := range s1 {
		var e1 KeyValue
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		a.Arguments = append(a.Arguments, e1)
	}
	return nil
}

func (c *EngineCondition) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["negate"] = c.Negate
	j0["operation"] = c.Operation
	j0["matchType"] = c.MatchType
	j0["eventId"] = encoding.NonNilSlice(c.EventId)
	s1 := make([]any, 0, len(c.Conditions))
	for _, e1 := range c.Conditions {
		s1 = append(s1, e1.Encode())
	}
	j0["conditions"] = s1
	return j0
}

func (c *EngineCondition) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("negate", j0)
	if err != nil {
		return err
	}
	c.Negate, err = encoding.Is[bool](j0["negate"])
	if err != nil {
		return err
	}
	err = encoding.In("operation", j0)
	if err != nil {
		return err
	}
	c.Operation, err = encoding.AsEnum[EngineConditionOp](j0["operation"])
	if err != nil {
		return err
	}
	err = encoding.In("matchType", j0)
	if err != nil {
		return err
	}
	c.MatchType, err = encoding.AsEnum[EngineConditionMatchType](j0["matchType"])
	if err != nil {
		return err
	}
	err = encoding.In("eventId", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["eventId"])
	if err != nil {
		return err
	}
	c.EventId = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		c.EventId = append(c.EventId, e1)
	}
	err = encoding.In("conditions", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["conditions"])
	if err != nil {
		return err
	}
	c.Conditions = make([]EngineCondition, 0, len(s2))
	for _, a2 := range s2 {
		var e2 EngineCondition
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		c.Conditions = append(c.Conditions, e2)
	}
	return nil
}

func (r *EngineRule) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["id"] = r.Id
	j0["name"] = r.Name
	j0["isSystem"] = r.IsSystem
	j0["isEnabled"] = r.IsEnabled
	j0["isAutoRearm"] = r.IsAutoRearm
	j0["hasMatched"] = r.HasMatched
	j0["condition"] = r.Condition.Encode()
	j0["actionIds"] = encoding.NonNilSlice(r.ActionIds)
	s1 := make([]any, 0, len(r.Arguments))
	for _, e1 := range r.Arguments {
		s1 = append(s1, e1.Encode())
	}
	j0["arguments"] = s1
	return j0
}

func (r *EngineRule) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	r.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	r.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("isSystem", j0)
	if err != nil {
		return err
	}
	r.IsSystem, err = encoding.Is[bool](j0["isSystem"])
	if err != nil {
		return err
	}
	err = encoding.In("isEnabled", j0)
	if err != nil {
		return err
	}
	r.IsEnabled, err = encoding.Is[bool](j0["isEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("isAutoRearm", j0)
	if err != nil {
		return err
	}
	r.IsAutoRearm, err = encoding.Is[bool](j0["isAutoRearm"])
	if err != nil {
		return err
	}
	err = encoding.In("hasMatched", j0)
	if err != nil {
		return err
	}
	r.HasMatched, err = encoding.Is[bool](j0["hasMatched"])
	if err != nil {
		return err
	}
	err = encoding.In("condition", j0)
	if err != nil {
		return err
	}
	err = r.Condition.Decode(j0["condition"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("actionIds", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["actionIds"])
	if err != nil {
		return err
	}
	r.ActionIds = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		r.ActionIds = append(r.ActionIds, e1)
	}
	err = encoding.In("arguments", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["arguments"])
	if err != nil {
		return err
	}
	r.Arguments = make([]KeyValue, 0, len(s2))
	for _, a2 := range s2 {
		var e2 KeyValue
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		r.Arguments = append(r.Arguments, e2)
	}
	return nil
}
