// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package keypadmanager

import (
	"github.com/arminguenther/xeruspower-go/v40032/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/idl/event"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40032/smartlock/keypad"
)

func (k *KeypadSettings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["name"] = k.Name
	j0["description"] = k.Description
	return j0
}

func (k *KeypadSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	k.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	k.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	return nil
}

func (k *_KeypadEvent) Decode(value map[string]any, caller idl.Caller) error {
	k.Event = valobj.For[event.Event]()
	err := k.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("keypad", value)
	if err != nil {
		return err
	}
	k.keypad, err = object.As[keypad.Keypad](value["keypad"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("metaData", value)
	if err != nil {
		return err
	}
	err = k.metaData.Decode(value["metaData"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (k *_KeypadAttachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	k.KeypadEvent = valobj.For[KeypadEvent]()
	err := k.KeypadEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (k *_KeypadDetachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	k.KeypadEvent = valobj.For[KeypadEvent]()
	err := k.KeypadEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (k *_KeypadSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	k.UserEvent = valobj.For[userevent.UserEvent]()
	err := k.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("keypad", value)
	if err != nil {
		return err
	}
	k.keypad, err = object.As[keypad.Keypad](value["keypad"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = k.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = k.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("position", value)
	if err != nil {
		return err
	}
	k.position, err = encoding.Is[string](value["position"])
	if err != nil {
		return err
	}
	return nil
}
