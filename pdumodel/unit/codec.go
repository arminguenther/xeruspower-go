// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package unit

import (
	"github.com/arminguenther/xeruspower-go/v40032/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["hasOrientationSensor"] = m.HasOrientationSensor
	j0["supportedDisplayOrientations"] = encoding.NonNilSlice(m.SupportedDisplayOrientations)
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("hasOrientationSensor", j0)
	if err != nil {
		return err
	}
	m.HasOrientationSensor, err = encoding.Is[bool](j0["hasOrientationSensor"])
	if err != nil {
		return err
	}
	err = encoding.In("supportedDisplayOrientations", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["supportedDisplayOrientations"])
	if err != nil {
		return err
	}
	m.SupportedDisplayOrientations = make([]Orientation, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Orientation
		e1, err = encoding.AsEnum[Orientation](a1)
		if err != nil {
			return err
		}
		m.SupportedDisplayOrientations = append(m.SupportedDisplayOrientations, e1)
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["buzzerMuted"] = s.BuzzerMuted
	j0["autoDisplayOrientation"] = s.AutoDisplayOrientation
	j0["displayOrientation"] = s.DisplayOrientation
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("buzzerMuted", j0)
	if err != nil {
		return err
	}
	s.BuzzerMuted, err = encoding.Is[bool](j0["buzzerMuted"])
	if err != nil {
		return err
	}
	err = encoding.In("autoDisplayOrientation", j0)
	if err != nil {
		return err
	}
	s.AutoDisplayOrientation, err = encoding.Is[bool](j0["autoDisplayOrientation"])
	if err != nil {
		return err
	}
	err = encoding.In("displayOrientation", j0)
	if err != nil {
		return err
	}
	s.DisplayOrientation, err = encoding.AsEnum[Orientation](j0["displayOrientation"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_IdentificationStartedEvent) Decode(value map[string]any, caller idl.Caller) error {
	i.UserEvent = valobj.For[userevent.UserEvent]()
	err := i.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("duration", value)
	if err != nil {
		return err
	}
	i.duration, err = encoding.AsInt32(value["duration"])
	if err != nil {
		return err
	}
	return nil
}
