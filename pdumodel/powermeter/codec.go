// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powermeter

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40000/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40000/sensors/statesensor"
)

func (c *Config) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["powerMeterId"] = c.PowerMeterId
	j0["type"] = c.Type
	return j0
}

func (c *Config) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("powerMeterId", j0)
	if err != nil {
		return err
	}
	c.PowerMeterId, err = encoding.AsInt32(j0["powerMeterId"])
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
	j0 := make(map[string]any, 14)
	j0["voltage"] = object.ToMap(s.Voltage)
	j0["lineFrequency"] = object.ToMap(s.LineFrequency)
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
	j0["powerQuality"] = object.ToMap(s.PowerQuality)
	return j0
}

func (s *Sensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	s.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("lineFrequency", j0)
	if err != nil {
		return err
	}
	s.LineFrequency, err = object.As[numericsensor.NumericSensor](j0["lineFrequency"], caller)
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
	err = encoding.In("powerQuality", j0)
	if err != nil {
		return err
	}
	s.PowerQuality, err = object.As[statesensor.StateSensor](j0["powerQuality"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["name"] = s.Name
	j0["currentRating"] = s.CurrentRating
	j0["phaseCtRating"] = s.PhaseCtRating
	j0["neutralCtRating"] = s.NeutralCtRating
	j0["earthCtRating"] = s.EarthCtRating
	j0["modbusUnitAddress"] = s.ModbusUnitAddress
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
	err = encoding.In("currentRating", j0)
	if err != nil {
		return err
	}
	s.CurrentRating, err = encoding.AsInt32(j0["currentRating"])
	if err != nil {
		return err
	}
	err = encoding.In("phaseCtRating", j0)
	if err != nil {
		return err
	}
	s.PhaseCtRating, err = encoding.AsInt32(j0["phaseCtRating"])
	if err != nil {
		return err
	}
	err = encoding.In("neutralCtRating", j0)
	if err != nil {
		return err
	}
	s.NeutralCtRating, err = encoding.AsInt32(j0["neutralCtRating"])
	if err != nil {
		return err
	}
	err = encoding.In("earthCtRating", j0)
	if err != nil {
		return err
	}
	s.EarthCtRating, err = encoding.AsInt32(j0["earthCtRating"])
	if err != nil {
		return err
	}
	err = encoding.In("modbusUnitAddress", j0)
	if err != nil {
		return err
	}
	s.ModbusUnitAddress, err = encoding.AsInt32(j0["modbusUnitAddress"])
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

func (e *EnergyPulseSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["pulseEnabled"] = e.PulseEnabled
	j0["poles"] = encoding.NonNilSlice(e.Poles)
	j0["pulsesPerKWh"] = e.PulsesPerKWh
	return j0
}

func (e *EnergyPulseSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("pulseEnabled", j0)
	if err != nil {
		return err
	}
	e.PulseEnabled, err = encoding.Is[bool](j0["pulseEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("poles", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["poles"])
	if err != nil {
		return err
	}
	e.Poles = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		e.Poles = append(e.Poles, e1)
	}
	err = encoding.In("pulsesPerKWh", j0)
	if err != nil {
		return err
	}
	e.PulsesPerKWh, err = encoding.AsInt32(j0["pulsesPerKWh"])
	if err != nil {
		return err
	}
	return nil
}

func (e *_EnergyPulseSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
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
