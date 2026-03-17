// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package circuit

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40000/pdumodel/pole"
	"github.com/arminguenther/xeruspower-go/v40000/sensors/numericsensor"
)

func (c *Config) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["position"] = c.Position
	j0["type"] = c.Type
	return j0
}

func (c *Config) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	c.Position, err = encoding.AsInt32(j0["position"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	c.Type, err = encoding.AsEnum[Type](j0["type"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 11)
	j0["current"] = object.ToMap(s.Current)
	j0["activePower"] = object.ToMap(s.ActivePower)
	j0["reactivePower"] = object.ToMap(s.ReactivePower)
	j0["apparentPower"] = object.ToMap(s.ApparentPower)
	j0["powerFactor"] = object.ToMap(s.PowerFactor)
	j0["phaseAngle"] = object.ToMap(s.PhaseAngle)
	j0["displacementPowerFactor"] = object.ToMap(s.DisplacementPowerFactor)
	j0["activeEnergy"] = object.ToMap(s.ActiveEnergy)
	j0["unbalancedCurrent"] = object.ToMap(s.UnbalancedCurrent)
	j0["crestFactor"] = object.ToMap(s.CrestFactor)
	j0["activePowerDemand"] = object.ToMap(s.ActivePowerDemand)
	return j0
}

func (s *Sensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	s.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
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
	err = encoding.In("reactivePower", j0)
	if err != nil {
		return err
	}
	s.ReactivePower, err = object.As[numericsensor.NumericSensor](j0["reactivePower"], caller)
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
	err = encoding.In("powerFactor", j0)
	if err != nil {
		return err
	}
	s.PowerFactor, err = object.As[numericsensor.NumericSensor](j0["powerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("phaseAngle", j0)
	if err != nil {
		return err
	}
	s.PhaseAngle, err = object.As[numericsensor.NumericSensor](j0["phaseAngle"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("displacementPowerFactor", j0)
	if err != nil {
		return err
	}
	s.DisplacementPowerFactor, err = object.As[numericsensor.NumericSensor](j0["displacementPowerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	s.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("unbalancedCurrent", j0)
	if err != nil {
		return err
	}
	s.UnbalancedCurrent, err = object.As[numericsensor.NumericSensor](j0["unbalancedCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("crestFactor", j0)
	if err != nil {
		return err
	}
	s.CrestFactor, err = object.As[numericsensor.NumericSensor](j0["crestFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePowerDemand", j0)
	if err != nil {
		return err
	}
	s.ActivePowerDemand, err = object.As[numericsensor.NumericSensor](j0["activePowerDemand"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *PoleSettings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["line"] = p.Line
	j0["meterChannel"] = p.MeterChannel
	return j0
}

func (p *PoleSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("line", j0)
	if err != nil {
		return err
	}
	p.Line, err = encoding.AsEnum[pole.PowerLine](j0["line"])
	if err != nil {
		return err
	}
	err = encoding.In("meterChannel", j0)
	if err != nil {
		return err
	}
	p.MeterChannel, err = encoding.AsInt32(j0["meterChannel"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["name"] = s.Name
	j0["rating"] = s.Rating
	j0["ctRating"] = s.CtRating
	s1 := make([]any, 0, len(s.PoleSettings))
	for _, e1 := range s.PoleSettings {
		s1 = append(s1, e1.Encode())
	}
	j0["poleSettings"] = s1
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
	err = encoding.In("rating", j0)
	if err != nil {
		return err
	}
	s.Rating, err = encoding.AsInt32(j0["rating"])
	if err != nil {
		return err
	}
	err = encoding.In("ctRating", j0)
	if err != nil {
		return err
	}
	s.CtRating, err = encoding.AsInt32(j0["ctRating"])
	if err != nil {
		return err
	}
	err = encoding.In("poleSettings", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["poleSettings"])
	if err != nil {
		return err
	}
	s.PoleSettings = make([]PoleSettings, 0, len(s1))
	for _, a1 := range s1 {
		var e1 PoleSettings
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		s.PoleSettings = append(s.PoleSettings, e1)
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
