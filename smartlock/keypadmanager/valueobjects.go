// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package keypadmanager

import (
	"github.com/arminguenther/xeruspower-go/v40040/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
	keypad_ "github.com/arminguenther/xeruspower-go/v40040/smartlock/keypad"
)

func init() {
	valobj.Register(func() KeypadAttachedEvent { return &_KeypadAttachedEvent{} })
	valobj.Register(func() KeypadDetachedEvent { return &_KeypadDetachedEvent{} })
	valobj.Register(func() KeypadEvent { return &_KeypadEvent{} })
	valobj.Register(func() KeypadSettingsChangedEvent { return &_KeypadSettingsChangedEvent{} })
}

type _KeypadEvent struct {
	event.Event
	keypad   keypad_.Keypad
	metaData keypad_.MetaData
}

func (k *_KeypadEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.KeypadManager_1_0_1.KeypadEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (k *_KeypadEvent) Keypad() keypad_.Keypad {
	return k.keypad
}

func (k *_KeypadEvent) MetaData() keypad_.MetaData {
	return k.metaData
}

func (k *_KeypadEvent) isKeypadEvent() {}

type _KeypadAttachedEvent struct {
	KeypadEvent
}

func (k *_KeypadAttachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.KeypadManager_1_0_1.KeypadAttachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (k *_KeypadAttachedEvent) isKeypadAttachedEvent() {}

type _KeypadDetachedEvent struct {
	KeypadEvent
}

func (k *_KeypadDetachedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.KeypadManager_1_0_1.KeypadDetachedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (k *_KeypadDetachedEvent) isKeypadDetachedEvent() {}

type _KeypadSettingsChangedEvent struct {
	userevent.UserEvent
	keypad      keypad_.Keypad
	oldSettings KeypadSettings
	newSettings KeypadSettings
	position    string
}

func (k *_KeypadSettingsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartlock.KeypadManager_1_0_1.KeypadSettingsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (k *_KeypadSettingsChangedEvent) Keypad() keypad_.Keypad {
	return k.keypad
}

func (k *_KeypadSettingsChangedEvent) OldSettings() KeypadSettings {
	return k.oldSettings
}

func (k *_KeypadSettingsChangedEvent) NewSettings() KeypadSettings {
	return k.newSettings
}

func (k *_KeypadSettingsChangedEvent) Position() string {
	return k.position
}

func (k *_KeypadSettingsChangedEvent) isKeypadSettingsChangedEvent() {}
