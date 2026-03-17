// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datapushservice

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/idl/event"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
)

func (e *EntrySettings) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["url"] = e.Url
	j0["allowOffTimeRangeCerts"] = e.AllowOffTimeRangeCerts
	j0["caCertChain"] = e.CaCertChain
	j0["useAuth"] = e.UseAuth
	j0["username"] = e.Username
	j0["password"] = e.Password
	j0["type"] = e.Type
	j0["items"] = encoding.NonNilSlice(e.Items)
	return j0
}

func (e *EntrySettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("url", j0)
	if err != nil {
		return err
	}
	e.Url, err = encoding.Is[string](j0["url"])
	if err != nil {
		return err
	}
	err = encoding.In("allowOffTimeRangeCerts", j0)
	if err != nil {
		return err
	}
	e.AllowOffTimeRangeCerts, err = encoding.Is[bool](j0["allowOffTimeRangeCerts"])
	if err != nil {
		return err
	}
	err = encoding.In("caCertChain", j0)
	if err != nil {
		return err
	}
	e.CaCertChain, err = encoding.Is[string](j0["caCertChain"])
	if err != nil {
		return err
	}
	err = encoding.In("useAuth", j0)
	if err != nil {
		return err
	}
	e.UseAuth, err = encoding.Is[bool](j0["useAuth"])
	if err != nil {
		return err
	}
	err = encoding.In("username", j0)
	if err != nil {
		return err
	}
	e.Username, err = encoding.Is[string](j0["username"])
	if err != nil {
		return err
	}
	err = encoding.In("password", j0)
	if err != nil {
		return err
	}
	e.Password, err = encoding.Is[string](j0["password"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	e.Type, err = encoding.AsEnum[EntryType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("items", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["items"])
	if err != nil {
		return err
	}
	e.Items = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		e.Items = append(e.Items, e1)
	}
	return nil
}

func (e *EntryStatus) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["busy"] = e.Busy
	j0["rescheduled"] = e.Rescheduled
	j0["lastAttemptTime"] = e.LastAttemptTime.Unix()
	j0["lastSuccessTime"] = e.LastSuccessTime.Unix()
	return j0
}

func (e *EntryStatus) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("busy", j0)
	if err != nil {
		return err
	}
	e.Busy, err = encoding.Is[bool](j0["busy"])
	if err != nil {
		return err
	}
	err = encoding.In("rescheduled", j0)
	if err != nil {
		return err
	}
	e.Rescheduled, err = encoding.Is[bool](j0["rescheduled"])
	if err != nil {
		return err
	}
	err = encoding.In("lastAttemptTime", j0)
	if err != nil {
		return err
	}
	e.LastAttemptTime, err = encoding.AsTime(j0["lastAttemptTime"])
	if err != nil {
		return err
	}
	err = encoding.In("lastSuccessTime", j0)
	if err != nil {
		return err
	}
	e.LastSuccessTime, err = encoding.AsTime(j0["lastSuccessTime"])
	if err != nil {
		return err
	}
	return nil
}

func (e *_EntryAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("entryId", value)
	if err != nil {
		return err
	}
	e.entryId, err = encoding.AsInt32(value["entryId"])
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = e.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (e *_EntryModifiedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("entryId", value)
	if err != nil {
		return err
	}
	e.entryId, err = encoding.AsInt32(value["entryId"])
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = e.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = e.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (e *_EntryDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("entryId", value)
	if err != nil {
		return err
	}
	e.entryId, err = encoding.AsInt32(value["entryId"])
	if err != nil {
		return err
	}
	return nil
}

func (e *_EntryStatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.Event = valobj.For[event.Event]()
	err := e.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("entryId", value)
	if err != nil {
		return err
	}
	e.entryId, err = encoding.AsInt32(value["entryId"])
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	err = e.newStatus.Decode(value["newStatus"], caller)
	if err != nil {
		return err
	}
	return nil
}
