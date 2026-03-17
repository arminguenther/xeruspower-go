// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outletgroup

import (
	"github.com/arminguenther/xeruspower-go/v40200/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/idl/event"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40200/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40200/sensors/accumulatingnumericsensor"
	"github.com/arminguenther/xeruspower-go/v40200/sensors/numericsensor"
)

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["activePower"] = object.ToMap(s.ActivePower)
	j0["apparentPower"] = object.ToMap(s.ApparentPower)
	j0["activeEnergy"] = object.ToMap(s.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(s.ApparentEnergy)
	return j0
}

func (s *Sensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	s.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentPower", j0)
	if err != nil {
		return err
	}
	s.ApparentPower, err = object.As[numericsensor.NumericSensor](j0["apparentPower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	s.ActiveEnergy, err = object.As[accumulatingnumericsensor.AccumulatingNumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentEnergy", j0)
	if err != nil {
		return err
	}
	s.ApparentEnergy, err = object.As[accumulatingnumericsensor.AccumulatingNumericSensor](j0["apparentEnergy"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["name"] = s.Name
	s1 := make([]any, 0, len(s.Members))
	for _, e1 := range s.Members {
		s1 = append(s1, object.ToMap(e1))
	}
	j0["members"] = s1
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	s.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("members", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["members"])
	if err != nil {
		return err
	}
	s.Members = make([]outlet.Outlet, 0, len(s1))
	for _, a1 := range s1 {
		var e1 outlet.Outlet
		e1, err = object.As[outlet.Outlet](a1, caller)
		if err != nil {
			return err
		}
		s.Members = append(s.Members, e1)
	}
	return nil
}

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["groupId"] = m.GroupId
	j0["uniqueId"] = m.UniqueId
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("groupId", j0)
	if err != nil {
		return err
	}
	m.GroupId, err = encoding.AsInt32(j0["groupId"])
	if err != nil {
		return err
	}
	err = encoding.In("uniqueId", j0)
	if err != nil {
		return err
	}
	m.UniqueId, err = encoding.AsInt32(j0["uniqueId"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SensorsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSensors", value)
	if err != nil {
		return err
	}
	err = s.oldSensors.Decode(value["oldSensors"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSensors", value)
	if err != nil {
		return err
	}
	err = s.newSensors.Decode(value["newSensors"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = s.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = s.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_PowerControlEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.UserEvent = valobj.For[userevent.UserEvent]()
	err := p.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("state", value)
	if err != nil {
		return err
	}
	p.state, err = encoding.AsEnum[outlet.PowerState](value["state"])
	if err != nil {
		return err
	}
	err = encoding.In("cycle", value)
	if err != nil {
		return err
	}
	p.cycle, err = encoding.Is[bool](value["cycle"])
	if err != nil {
		return err
	}
	return nil
}
