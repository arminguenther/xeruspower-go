// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dooraccesscontrol

import (
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DoorAccessDeniedEvent { return &_DoorAccessDeniedEvent{} })
	valobj.Register(func() DoorAccessGrantedEvent { return &_DoorAccessGrantedEvent{} })
	valobj.Register(func() DoorAccessRuleAddedEvent { return &_DoorAccessRuleAddedEvent{} })
	valobj.Register(func() DoorAccessRuleChangedEvent { return &_DoorAccessRuleChangedEvent{} })
	valobj.Register(func() DoorAccessRuleDeletedEvent { return &_DoorAccessRuleDeletedEvent{} })
}

type _DoorAccessGrantedEvent struct {
	event.Event
	ruleId int32
	rule   DoorAccessRule
}

func (d *_DoorAccessGrantedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.DoorAccessControl_1_2_3.DoorAccessGrantedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DoorAccessGrantedEvent) RuleId() int32 {
	return d.ruleId
}

func (d *_DoorAccessGrantedEvent) Rule() DoorAccessRule {
	return d.rule
}

func (d *_DoorAccessGrantedEvent) isDoorAccessGrantedEvent() {}

type _DoorAccessDeniedEvent struct {
	event.Event
	reason   DoorAccessDenialReason
	ruleId   int32
	ruleName string
}

func (d *_DoorAccessDeniedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.DoorAccessControl_1_2_3.DoorAccessDeniedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DoorAccessDeniedEvent) Reason() DoorAccessDenialReason {
	return d.reason
}

func (d *_DoorAccessDeniedEvent) RuleId() int32 {
	return d.ruleId
}

func (d *_DoorAccessDeniedEvent) RuleName() string {
	return d.ruleName
}

func (d *_DoorAccessDeniedEvent) isDoorAccessDeniedEvent() {}

type _DoorAccessRuleAddedEvent struct {
	userevent.UserEvent
	ruleId int32
	rule   DoorAccessRule
}

func (d *_DoorAccessRuleAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.DoorAccessControl_1_2_3.DoorAccessRuleAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DoorAccessRuleAddedEvent) RuleId() int32 {
	return d.ruleId
}

func (d *_DoorAccessRuleAddedEvent) Rule() DoorAccessRule {
	return d.rule
}

func (d *_DoorAccessRuleAddedEvent) isDoorAccessRuleAddedEvent() {}

type _DoorAccessRuleChangedEvent struct {
	userevent.UserEvent
	ruleId  int32
	oldRule DoorAccessRule
	newRule DoorAccessRule
}

func (d *_DoorAccessRuleChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.DoorAccessControl_1_2_3.DoorAccessRuleChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DoorAccessRuleChangedEvent) RuleId() int32 {
	return d.ruleId
}

func (d *_DoorAccessRuleChangedEvent) OldRule() DoorAccessRule {
	return d.oldRule
}

func (d *_DoorAccessRuleChangedEvent) NewRule() DoorAccessRule {
	return d.newRule
}

func (d *_DoorAccessRuleChangedEvent) isDoorAccessRuleChangedEvent() {}

type _DoorAccessRuleDeletedEvent struct {
	userevent.UserEvent
	ruleId int32
	rule   DoorAccessRule
}

func (d *_DoorAccessRuleDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.DoorAccessControl_1_2_3.DoorAccessRuleDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DoorAccessRuleDeletedEvent) RuleId() int32 {
	return d.ruleId
}

func (d *_DoorAccessRuleDeletedEvent) Rule() DoorAccessRule {
	return d.rule
}

func (d *_DoorAccessRuleDeletedEvent) isDoorAccessRuleDeletedEvent() {}
