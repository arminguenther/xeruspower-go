// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outlet

import (
	"github.com/arminguenther/xeruspower-go/v40410/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/idl/event"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40410/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40410/sensors/statesensor"
)

func (o *Statistic) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["relayCycleCnt"] = o.RelayCycleCnt
	j0["relayFailCnt"] = o.RelayFailCnt
	return j0
}

func (o *Statistic) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("relayCycleCnt", j0)
	if err != nil {
		return err
	}
	o.RelayCycleCnt, err = encoding.AsInt32(j0["relayCycleCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("relayFailCnt", j0)
	if err != nil {
		return err
	}
	o.RelayFailCnt, err = encoding.AsInt32(j0["relayFailCnt"])
	if err != nil {
		return err
	}
	return nil
}

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["label"] = m.Label
	j0["receptacleType"] = m.ReceptacleType
	j0["namePlate"] = m.NamePlate.Encode()
	j0["rating"] = m.Rating.Encode()
	j0["isSwitchable"] = m.IsSwitchable
	j0["isLatching"] = m.IsLatching
	j0["maxRelayCycleCnt"] = m.MaxRelayCycleCnt
	j0["hasWaveformSupport"] = m.HasWaveformSupport
	j0["hasServiceModeSupport"] = m.HasServiceModeSupport
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	m.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("receptacleType", j0)
	if err != nil {
		return err
	}
	m.ReceptacleType, err = encoding.Is[string](j0["receptacleType"])
	if err != nil {
		return err
	}
	err = encoding.In("namePlate", j0)
	if err != nil {
		return err
	}
	err = m.NamePlate.Decode(j0["namePlate"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("rating", j0)
	if err != nil {
		return err
	}
	err = m.Rating.Decode(j0["rating"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("isSwitchable", j0)
	if err != nil {
		return err
	}
	m.IsSwitchable, err = encoding.Is[bool](j0["isSwitchable"])
	if err != nil {
		return err
	}
	err = encoding.In("isLatching", j0)
	if err != nil {
		return err
	}
	m.IsLatching, err = encoding.Is[bool](j0["isLatching"])
	if err != nil {
		return err
	}
	err = encoding.In("maxRelayCycleCnt", j0)
	if err != nil {
		return err
	}
	m.MaxRelayCycleCnt, err = encoding.AsInt32(j0["maxRelayCycleCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("hasWaveformSupport", j0)
	if err != nil {
		return err
	}
	m.HasWaveformSupport, err = encoding.Is[bool](j0["hasWaveformSupport"])
	if err != nil {
		return err
	}
	err = encoding.In("hasServiceModeSupport", j0)
	if err != nil {
		return err
	}
	m.HasServiceModeSupport, err = encoding.Is[bool](j0["hasServiceModeSupport"])
	if err != nil {
		return err
	}
	return nil
}

func (l *LedState) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["red"] = l.Red
	j0["green"] = l.Green
	j0["blinking"] = l.Blinking
	return j0
}

func (l *LedState) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("red", j0)
	if err != nil {
		return err
	}
	l.Red, err = encoding.Is[bool](j0["red"])
	if err != nil {
		return err
	}
	err = encoding.In("green", j0)
	if err != nil {
		return err
	}
	l.Green, err = encoding.Is[bool](j0["green"])
	if err != nil {
		return err
	}
	err = encoding.In("blinking", j0)
	if err != nil {
		return err
	}
	l.Blinking, err = encoding.Is[bool](j0["blinking"])
	if err != nil {
		return err
	}
	return nil
}

func (s *State) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["available"] = s.Available
	j0["powerState"] = s.PowerState
	j0["switchOnInProgress"] = s.SwitchOnInProgress
	j0["cycleInProgress"] = s.CycleInProgress
	j0["isLoadShed"] = s.IsLoadShed
	j0["isSuspended"] = s.IsSuspended
	j0["hasInrushWaveform"] = s.HasInrushWaveform
	j0["inServiceMode"] = s.InServiceMode
	j0["ledState"] = s.LedState.Encode()
	j0["lastPowerStateChange"] = s.LastPowerStateChange.Unix()
	return j0
}

func (s *State) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("available", j0)
	if err != nil {
		return err
	}
	s.Available, err = encoding.Is[bool](j0["available"])
	if err != nil {
		return err
	}
	err = encoding.In("powerState", j0)
	if err != nil {
		return err
	}
	s.PowerState, err = encoding.AsEnum[PowerState](j0["powerState"])
	if err != nil {
		return err
	}
	err = encoding.In("switchOnInProgress", j0)
	if err != nil {
		return err
	}
	s.SwitchOnInProgress, err = encoding.Is[bool](j0["switchOnInProgress"])
	if err != nil {
		return err
	}
	err = encoding.In("cycleInProgress", j0)
	if err != nil {
		return err
	}
	s.CycleInProgress, err = encoding.Is[bool](j0["cycleInProgress"])
	if err != nil {
		return err
	}
	err = encoding.In("isLoadShed", j0)
	if err != nil {
		return err
	}
	s.IsLoadShed, err = encoding.Is[bool](j0["isLoadShed"])
	if err != nil {
		return err
	}
	err = encoding.In("isSuspended", j0)
	if err != nil {
		return err
	}
	s.IsSuspended, err = encoding.Is[bool](j0["isSuspended"])
	if err != nil {
		return err
	}
	err = encoding.In("hasInrushWaveform", j0)
	if err != nil {
		return err
	}
	s.HasInrushWaveform, err = encoding.Is[bool](j0["hasInrushWaveform"])
	if err != nil {
		return err
	}
	err = encoding.In("inServiceMode", j0)
	if err != nil {
		return err
	}
	s.InServiceMode, err = encoding.Is[bool](j0["inServiceMode"])
	if err != nil {
		return err
	}
	err = encoding.In("ledState", j0)
	if err != nil {
		return err
	}
	err = s.LedState.Decode(j0["ledState"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("lastPowerStateChange", j0)
	if err != nil {
		return err
	}
	s.LastPowerStateChange, err = encoding.AsTime(j0["lastPowerStateChange"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["name"] = s.Name
	j0["startupState"] = s.StartupState
	j0["usePduCycleDelay"] = s.UsePduCycleDelay
	j0["cycleDelay"] = s.CycleDelay
	j0["nonCritical"] = s.NonCritical
	j0["sequenceDelay"] = s.SequenceDelay
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
	err = encoding.In("startupState", j0)
	if err != nil {
		return err
	}
	s.StartupState, err = encoding.AsEnum[StartupState](j0["startupState"])
	if err != nil {
		return err
	}
	err = encoding.In("usePduCycleDelay", j0)
	if err != nil {
		return err
	}
	s.UsePduCycleDelay, err = encoding.Is[bool](j0["usePduCycleDelay"])
	if err != nil {
		return err
	}
	err = encoding.In("cycleDelay", j0)
	if err != nil {
		return err
	}
	s.CycleDelay, err = encoding.AsInt32(j0["cycleDelay"])
	if err != nil {
		return err
	}
	err = encoding.In("nonCritical", j0)
	if err != nil {
		return err
	}
	s.NonCritical, err = encoding.Is[bool](j0["nonCritical"])
	if err != nil {
		return err
	}
	err = encoding.In("sequenceDelay", j0)
	if err != nil {
		return err
	}
	s.SequenceDelay, err = encoding.AsInt32(j0["sequenceDelay"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 20)
	j0["voltage"] = object.ToMap(s.Voltage)
	j0["current"] = object.ToMap(s.Current)
	j0["peakCurrent"] = object.ToMap(s.PeakCurrent)
	j0["maximumCurrent"] = object.ToMap(s.MaximumCurrent)
	j0["unbalancedCurrent"] = object.ToMap(s.UnbalancedCurrent)
	j0["unbalancedVoltage"] = object.ToMap(s.UnbalancedVoltage)
	j0["activePower"] = object.ToMap(s.ActivePower)
	j0["reactivePower"] = object.ToMap(s.ReactivePower)
	j0["apparentPower"] = object.ToMap(s.ApparentPower)
	j0["powerFactor"] = object.ToMap(s.PowerFactor)
	j0["displacementPowerFactor"] = object.ToMap(s.DisplacementPowerFactor)
	j0["activeEnergy"] = object.ToMap(s.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(s.ApparentEnergy)
	j0["phaseAngle"] = object.ToMap(s.PhaseAngle)
	j0["lineFrequency"] = object.ToMap(s.LineFrequency)
	j0["crestFactor"] = object.ToMap(s.CrestFactor)
	j0["voltageThd"] = object.ToMap(s.VoltageThd)
	j0["currentThd"] = object.ToMap(s.CurrentThd)
	j0["inrushCurrent"] = object.ToMap(s.InrushCurrent)
	j0["outletState"] = object.ToMap(s.OutletState)
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
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	s.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("peakCurrent", j0)
	if err != nil {
		return err
	}
	s.PeakCurrent, err = object.As[numericsensor.NumericSensor](j0["peakCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("maximumCurrent", j0)
	if err != nil {
		return err
	}
	s.MaximumCurrent, err = object.As[numericsensor.NumericSensor](j0["maximumCurrent"], caller)
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
	err = encoding.In("unbalancedVoltage", j0)
	if err != nil {
		return err
	}
	s.UnbalancedVoltage, err = object.As[numericsensor.NumericSensor](j0["unbalancedVoltage"], caller)
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
	err = encoding.In("apparentEnergy", j0)
	if err != nil {
		return err
	}
	s.ApparentEnergy, err = object.As[numericsensor.NumericSensor](j0["apparentEnergy"], caller)
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
	err = encoding.In("lineFrequency", j0)
	if err != nil {
		return err
	}
	s.LineFrequency, err = object.As[numericsensor.NumericSensor](j0["lineFrequency"], caller)
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
	err = encoding.In("voltageThd", j0)
	if err != nil {
		return err
	}
	s.VoltageThd, err = object.As[numericsensor.NumericSensor](j0["voltageThd"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("currentThd", j0)
	if err != nil {
		return err
	}
	s.CurrentThd, err = object.As[numericsensor.NumericSensor](j0["currentThd"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("inrushCurrent", j0)
	if err != nil {
		return err
	}
	s.InrushCurrent, err = object.As[numericsensor.NumericSensor](j0["inrushCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("outletState", j0)
	if err != nil {
		return err
	}
	s.OutletState, err = object.As[statesensor.StateSensor](j0["outletState"], caller)
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
	p.state, err = encoding.AsEnum[PowerState](value["state"])
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

func (s *_StateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldState", value)
	if err != nil {
		return err
	}
	err = s.oldState.Decode(value["oldState"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	err = s.newState.Decode(value["newState"], caller)
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

func (s *_ServiceModeChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", value)
	if err != nil {
		return err
	}
	s.enabled, err = encoding.Is[bool](value["enabled"])
	if err != nil {
		return err
	}
	return nil
}
