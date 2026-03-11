// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alarmmanager

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func (a *Alert) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["eventCondition"] = a.EventCondition
	j0["message"] = a.Message
	j0["firstAppearance"] = a.FirstAppearance.Unix()
	j0["lastAppearance"] = a.LastAppearance.Unix()
	j0["numberAlerts"] = a.NumberAlerts
	return j0
}

func (a *Alert) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("eventCondition", j0)
	if err != nil {
		return err
	}
	a.EventCondition, err = encoding.Is[string](j0["eventCondition"])
	if err != nil {
		return err
	}
	err = encoding.In("message", j0)
	if err != nil {
		return err
	}
	a.Message, err = encoding.Is[string](j0["message"])
	if err != nil {
		return err
	}
	err = encoding.In("firstAppearance", j0)
	if err != nil {
		return err
	}
	a.FirstAppearance, err = encoding.AsTime(j0["firstAppearance"])
	if err != nil {
		return err
	}
	err = encoding.In("lastAppearance", j0)
	if err != nil {
		return err
	}
	a.LastAppearance, err = encoding.AsTime(j0["lastAppearance"])
	if err != nil {
		return err
	}
	err = encoding.In("numberAlerts", j0)
	if err != nil {
		return err
	}
	a.NumberAlerts, err = encoding.AsInt32(j0["numberAlerts"])
	if err != nil {
		return err
	}
	return nil
}

func (a *Alarm) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["id"] = a.Id
	j0["name"] = a.Name
	j0["actionId"] = a.ActionId
	s1 := make([]any, 0, len(a.Alerts))
	for _, e1 := range a.Alerts {
		s1 = append(s1, e1.Encode())
	}
	j0["alerts"] = s1
	return j0
}

func (a *Alarm) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("actionId", j0)
	if err != nil {
		return err
	}
	a.ActionId, err = encoding.Is[string](j0["actionId"])
	if err != nil {
		return err
	}
	err = encoding.In("alerts", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["alerts"])
	if err != nil {
		return err
	}
	a.Alerts = make([]Alert, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Alert
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		a.Alerts = append(a.Alerts, e1)
	}
	return nil
}

func (a *_AlarmAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	a.Event = valobj.For[event.Event]()
	err := a.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("alarm", value)
	if err != nil {
		return err
	}
	err = a.alarm.Decode(value["alarm"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *_AlarmUpdatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	a.Event = valobj.For[event.Event]()
	err := a.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("alarm", value)
	if err != nil {
		return err
	}
	err = a.alarm.Decode(value["alarm"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *_AlarmAcknowledgedEvent) Decode(value map[string]any, caller idl.Caller) error {
	a.Event = valobj.For[event.Event]()
	err := a.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("alarmId", value)
	if err != nil {
		return err
	}
	a.alarmId, err = encoding.Is[string](value["alarmId"])
	if err != nil {
		return err
	}
	return nil
}
