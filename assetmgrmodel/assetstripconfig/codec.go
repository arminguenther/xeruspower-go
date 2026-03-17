// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstripconfig

import (
	"github.com/arminguenther/xeruspower-go/v40200/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/valobj"
)

func (l *LEDColor) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["r"] = l.R
	j0["g"] = l.G
	j0["b"] = l.B
	return j0
}

func (l *LEDColor) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("r", j0)
	if err != nil {
		return err
	}
	l.R, err = encoding.AsInt32(j0["r"])
	if err != nil {
		return err
	}
	err = encoding.In("g", j0)
	if err != nil {
		return err
	}
	l.G, err = encoding.AsInt32(j0["g"])
	if err != nil {
		return err
	}
	err = encoding.In("b", j0)
	if err != nil {
		return err
	}
	l.B, err = encoding.AsInt32(j0["b"])
	if err != nil {
		return err
	}
	return nil
}

func (s *StripSettings) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["rackUnitCount"] = s.RackUnitCount
	j0["name"] = s.Name
	j0["scanMode"] = s.ScanMode
	j0["defaultColorConnected"] = s.DefaultColorConnected.Encode()
	j0["defaultColorDisconnected"] = s.DefaultColorDisconnected.Encode()
	j0["numberingMode"] = s.NumberingMode
	j0["numberingOffset"] = s.NumberingOffset
	j0["orientation"] = s.Orientation
	return j0
}

func (s *StripSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitCount", j0)
	if err != nil {
		return err
	}
	s.RackUnitCount, err = encoding.AsInt32(j0["rackUnitCount"])
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
	err = encoding.In("scanMode", j0)
	if err != nil {
		return err
	}
	s.ScanMode, err = encoding.AsEnum[ScanMode](j0["scanMode"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultColorConnected", j0)
	if err != nil {
		return err
	}
	err = s.DefaultColorConnected.Decode(j0["defaultColorConnected"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("defaultColorDisconnected", j0)
	if err != nil {
		return err
	}
	err = s.DefaultColorDisconnected.Decode(j0["defaultColorDisconnected"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("numberingMode", j0)
	if err != nil {
		return err
	}
	s.NumberingMode, err = encoding.AsEnum[NumberingMode](j0["numberingMode"])
	if err != nil {
		return err
	}
	err = encoding.In("numberingOffset", j0)
	if err != nil {
		return err
	}
	s.NumberingOffset, err = encoding.AsInt32(j0["numberingOffset"])
	if err != nil {
		return err
	}
	err = encoding.In("orientation", j0)
	if err != nil {
		return err
	}
	s.Orientation, err = encoding.AsEnum[Orientation](j0["orientation"])
	if err != nil {
		return err
	}
	return nil
}

func (r *RackUnitSettings) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["opmode"] = r.Opmode
	j0["mode"] = r.Mode
	j0["color"] = r.Color.Encode()
	j0["name"] = r.Name
	return j0
}

func (r *RackUnitSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("opmode", j0)
	if err != nil {
		return err
	}
	r.Opmode, err = encoding.AsEnum[LEDOperationMode](j0["opmode"])
	if err != nil {
		return err
	}
	err = encoding.In("mode", j0)
	if err != nil {
		return err
	}
	r.Mode, err = encoding.AsEnum[LEDMode](j0["mode"])
	if err != nil {
		return err
	}
	err = encoding.In("color", j0)
	if err != nil {
		return err
	}
	err = r.Color.Decode(j0["color"], caller)
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
	return nil
}

func (s *_StripSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
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

func (r *_RackUnitSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitNumber", value)
	if err != nil {
		return err
	}
	r.rackUnitNumber, err = encoding.AsInt32(value["rackUnitNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = r.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = r.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}
