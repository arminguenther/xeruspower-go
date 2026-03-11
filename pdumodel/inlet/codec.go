// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package inlet

import (
	"github.com/arminguenther/xeruspower-go/v40412/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40412/pdumodel/residualcurrentstatesensor"
	"github.com/arminguenther/xeruspower-go/v40412/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40412/sensors/statesensor"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["label"] = m.Label
	j0["plugType"] = m.PlugType
	j0["namePlate"] = m.NamePlate.Encode()
	j0["rating"] = m.Rating.Encode()
	j0["hasWaveformSupport"] = m.HasWaveformSupport
	j0["isDC"] = m.IsDC
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
	err = encoding.In("plugType", j0)
	if err != nil {
		return err
	}
	m.PlugType, err = encoding.Is[string](j0["plugType"])
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
	err = encoding.In("hasWaveformSupport", j0)
	if err != nil {
		return err
	}
	m.HasWaveformSupport, err = encoding.Is[bool](j0["hasWaveformSupport"])
	if err != nil {
		return err
	}
	err = encoding.In("isDC", j0)
	if err != nil {
		return err
	}
	m.IsDC, err = encoding.Is[bool](j0["isDC"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 25)
	j0["voltage"] = object.ToMap(s.Voltage)
	j0["current"] = object.ToMap(s.Current)
	j0["peakCurrent"] = object.ToMap(s.PeakCurrent)
	j0["residualCurrent"] = object.ToMap(s.ResidualCurrent)
	j0["residualACCurrent"] = object.ToMap(s.ResidualACCurrent)
	j0["residualDCCurrent"] = object.ToMap(s.ResidualDCCurrent)
	j0["activePower"] = object.ToMap(s.ActivePower)
	j0["reactivePower"] = object.ToMap(s.ReactivePower)
	j0["apparentPower"] = object.ToMap(s.ApparentPower)
	j0["powerFactor"] = object.ToMap(s.PowerFactor)
	j0["displacementPowerFactor"] = object.ToMap(s.DisplacementPowerFactor)
	j0["activeEnergy"] = object.ToMap(s.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(s.ApparentEnergy)
	j0["unbalancedCurrent"] = object.ToMap(s.UnbalancedCurrent)
	j0["unbalancedLineLineCurrent"] = object.ToMap(s.UnbalancedLineLineCurrent)
	j0["unbalancedVoltage"] = object.ToMap(s.UnbalancedVoltage)
	j0["unbalancedLineLineVoltage"] = object.ToMap(s.UnbalancedLineLineVoltage)
	j0["lineFrequency"] = object.ToMap(s.LineFrequency)
	j0["phaseAngle"] = object.ToMap(s.PhaseAngle)
	j0["crestFactor"] = object.ToMap(s.CrestFactor)
	j0["voltageThd"] = object.ToMap(s.VoltageThd)
	j0["currentThd"] = object.ToMap(s.CurrentThd)
	j0["powerQuality"] = object.ToMap(s.PowerQuality)
	j0["surgeProtectorStatus"] = object.ToMap(s.SurgeProtectorStatus)
	j0["residualCurrentStatus"] = object.ToMap(s.ResidualCurrentStatus)
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
	err = encoding.In("residualCurrent", j0)
	if err != nil {
		return err
	}
	s.ResidualCurrent, err = object.As[numericsensor.NumericSensor](j0["residualCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualACCurrent", j0)
	if err != nil {
		return err
	}
	s.ResidualACCurrent, err = object.As[numericsensor.NumericSensor](j0["residualACCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualDCCurrent", j0)
	if err != nil {
		return err
	}
	s.ResidualDCCurrent, err = object.As[numericsensor.NumericSensor](j0["residualDCCurrent"], caller)
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
	err = encoding.In("unbalancedCurrent", j0)
	if err != nil {
		return err
	}
	s.UnbalancedCurrent, err = object.As[numericsensor.NumericSensor](j0["unbalancedCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("unbalancedLineLineCurrent", j0)
	if err != nil {
		return err
	}
	s.UnbalancedLineLineCurrent, err = object.As[numericsensor.NumericSensor](j0["unbalancedLineLineCurrent"], caller)
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
	err = encoding.In("unbalancedLineLineVoltage", j0)
	if err != nil {
		return err
	}
	s.UnbalancedLineLineVoltage, err = object.As[numericsensor.NumericSensor](j0["unbalancedLineLineVoltage"], caller)
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
	err = encoding.In("phaseAngle", j0)
	if err != nil {
		return err
	}
	s.PhaseAngle, err = object.As[numericsensor.NumericSensor](j0["phaseAngle"], caller)
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
	err = encoding.In("powerQuality", j0)
	if err != nil {
		return err
	}
	s.PowerQuality, err = object.As[statesensor.StateSensor](j0["powerQuality"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("surgeProtectorStatus", j0)
	if err != nil {
		return err
	}
	s.SurgeProtectorStatus, err = object.As[statesensor.StateSensor](j0["surgeProtectorStatus"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualCurrentStatus", j0)
	if err != nil {
		return err
	}
	s.ResidualCurrentStatus, err = object.As[residualcurrentstatesensor.ResidualCurrentStateSensor](j0["residualCurrentStatus"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["name"] = s.Name
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

func (e *_EnableStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", value)
	if err != nil {
		return err
	}
	e.enabled, err = encoding.Is[bool](value["enabled"])
	if err != nil {
		return err
	}
	return nil
}
