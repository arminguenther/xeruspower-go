// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package shelfpowersupply

import (
	"github.com/arminguenther/xeruspower-go/v40412/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/idl/event"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40412/sensors/numericsensor"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 13)
	j0["slotLabel"] = m.SlotLabel
	j0["manufacturer"] = m.Manufacturer
	j0["model"] = m.Model
	j0["serialNumber"] = m.SerialNumber
	j0["firmwareVersion"] = m.FirmwareVersion
	j0["productionDate"] = m.ProductionDate
	j0["hardwareRevision"] = m.HardwareRevision
	j0["maxOutputPower"] = m.MaxOutputPower
	j0["hotplugSupported"] = m.HotplugSupported
	j0["powerControlSupported"] = m.PowerControlSupported
	j0["fanSpeedControlSupported"] = m.FanSpeedControlSupported
	j0["ratedLoad"] = m.RatedLoad
	j0["efficiencyPercent"] = m.EfficiencyPercent
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("slotLabel", j0)
	if err != nil {
		return err
	}
	m.SlotLabel, err = encoding.Is[string](j0["slotLabel"])
	if err != nil {
		return err
	}
	err = encoding.In("manufacturer", j0)
	if err != nil {
		return err
	}
	m.Manufacturer, err = encoding.Is[string](j0["manufacturer"])
	if err != nil {
		return err
	}
	err = encoding.In("model", j0)
	if err != nil {
		return err
	}
	m.Model, err = encoding.Is[string](j0["model"])
	if err != nil {
		return err
	}
	err = encoding.In("serialNumber", j0)
	if err != nil {
		return err
	}
	m.SerialNumber, err = encoding.Is[string](j0["serialNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("firmwareVersion", j0)
	if err != nil {
		return err
	}
	m.FirmwareVersion, err = encoding.Is[string](j0["firmwareVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("productionDate", j0)
	if err != nil {
		return err
	}
	m.ProductionDate, err = encoding.Is[string](j0["productionDate"])
	if err != nil {
		return err
	}
	err = encoding.In("hardwareRevision", j0)
	if err != nil {
		return err
	}
	m.HardwareRevision, err = encoding.Is[string](j0["hardwareRevision"])
	if err != nil {
		return err
	}
	err = encoding.In("maxOutputPower", j0)
	if err != nil {
		return err
	}
	m.MaxOutputPower, err = encoding.AsInt32(j0["maxOutputPower"])
	if err != nil {
		return err
	}
	err = encoding.In("hotplugSupported", j0)
	if err != nil {
		return err
	}
	m.HotplugSupported, err = encoding.Is[bool](j0["hotplugSupported"])
	if err != nil {
		return err
	}
	err = encoding.In("powerControlSupported", j0)
	if err != nil {
		return err
	}
	m.PowerControlSupported, err = encoding.Is[bool](j0["powerControlSupported"])
	if err != nil {
		return err
	}
	err = encoding.In("fanSpeedControlSupported", j0)
	if err != nil {
		return err
	}
	m.FanSpeedControlSupported, err = encoding.Is[bool](j0["fanSpeedControlSupported"])
	if err != nil {
		return err
	}
	err = encoding.In("ratedLoad", j0)
	if err != nil {
		return err
	}
	m.RatedLoad, err = encoding.AsInt32(j0["ratedLoad"])
	if err != nil {
		return err
	}
	err = encoding.In("efficiencyPercent", j0)
	if err != nil {
		return err
	}
	m.EfficiencyPercent, err = encoding.AsInt32(j0["efficiencyPercent"])
	if err != nil {
		return err
	}
	return nil
}

func (s *State) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["available"] = s.Available
	j0["isPresent"] = s.IsPresent
	j0["powerState"] = s.PowerState
	j0["fanSpeedOverridePercent"] = s.FanSpeedOverridePercent
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
	err = encoding.In("isPresent", j0)
	if err != nil {
		return err
	}
	s.IsPresent, err = encoding.Is[bool](j0["isPresent"])
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
	err = encoding.In("fanSpeedOverridePercent", j0)
	if err != nil {
		return err
	}
	s.FanSpeedOverridePercent, err = encoding.AsInt32(j0["fanSpeedOverridePercent"])
	if err != nil {
		return err
	}
	return nil
}

func (i *InputSensors) Encode() map[string]any {
	j0 := make(map[string]any, 16)
	j0["voltage"] = object.ToMap(i.Voltage)
	j0["current"] = object.ToMap(i.Current)
	j0["peakCurrent"] = object.ToMap(i.PeakCurrent)
	j0["unbalancedCurrent"] = object.ToMap(i.UnbalancedCurrent)
	j0["activePower"] = object.ToMap(i.ActivePower)
	j0["reactivePower"] = object.ToMap(i.ReactivePower)
	j0["apparentPower"] = object.ToMap(i.ApparentPower)
	j0["powerFactor"] = object.ToMap(i.PowerFactor)
	j0["displacementPowerFactor"] = object.ToMap(i.DisplacementPowerFactor)
	j0["activeEnergy"] = object.ToMap(i.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(i.ApparentEnergy)
	j0["phaseAngle"] = object.ToMap(i.PhaseAngle)
	j0["lineFrequency"] = object.ToMap(i.LineFrequency)
	j0["crestFactor"] = object.ToMap(i.CrestFactor)
	j0["voltageThd"] = object.ToMap(i.VoltageThd)
	j0["currentThd"] = object.ToMap(i.CurrentThd)
	return j0
}

func (i *InputSensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	i.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	i.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("peakCurrent", j0)
	if err != nil {
		return err
	}
	i.PeakCurrent, err = object.As[numericsensor.NumericSensor](j0["peakCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("unbalancedCurrent", j0)
	if err != nil {
		return err
	}
	i.UnbalancedCurrent, err = object.As[numericsensor.NumericSensor](j0["unbalancedCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	i.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("reactivePower", j0)
	if err != nil {
		return err
	}
	i.ReactivePower, err = object.As[numericsensor.NumericSensor](j0["reactivePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentPower", j0)
	if err != nil {
		return err
	}
	i.ApparentPower, err = object.As[numericsensor.NumericSensor](j0["apparentPower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("powerFactor", j0)
	if err != nil {
		return err
	}
	i.PowerFactor, err = object.As[numericsensor.NumericSensor](j0["powerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("displacementPowerFactor", j0)
	if err != nil {
		return err
	}
	i.DisplacementPowerFactor, err = object.As[numericsensor.NumericSensor](j0["displacementPowerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	i.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentEnergy", j0)
	if err != nil {
		return err
	}
	i.ApparentEnergy, err = object.As[numericsensor.NumericSensor](j0["apparentEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("phaseAngle", j0)
	if err != nil {
		return err
	}
	i.PhaseAngle, err = object.As[numericsensor.NumericSensor](j0["phaseAngle"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("lineFrequency", j0)
	if err != nil {
		return err
	}
	i.LineFrequency, err = object.As[numericsensor.NumericSensor](j0["lineFrequency"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("crestFactor", j0)
	if err != nil {
		return err
	}
	i.CrestFactor, err = object.As[numericsensor.NumericSensor](j0["crestFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("voltageThd", j0)
	if err != nil {
		return err
	}
	i.VoltageThd, err = object.As[numericsensor.NumericSensor](j0["voltageThd"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("currentThd", j0)
	if err != nil {
		return err
	}
	i.CurrentThd, err = object.As[numericsensor.NumericSensor](j0["currentThd"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (o *OutputSensors) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["voltage"] = object.ToMap(o.Voltage)
	j0["current"] = object.ToMap(o.Current)
	j0["activePower"] = object.ToMap(o.ActivePower)
	j0["activeEnergy"] = object.ToMap(o.ActiveEnergy)
	return j0
}

func (o *OutputSensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	o.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	o.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	o.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	o.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *PowerSupplySensors) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	s1 := make([]any, 0, len(p.Fans))
	for k1, v1 := range p.Fans {
		s1 = append(s1, map[string]any{"key": k1, "value": object.ToMap(v1)})
	}
	j0["fans"] = s1
	s2 := make([]any, 0, len(p.Temperatures))
	for k2, v2 := range p.Temperatures {
		s2 = append(s2, map[string]any{"key": k2, "value": object.ToMap(v2)})
	}
	j0["temperatures"] = s2
	return j0
}

func (p *PowerSupplySensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("fans", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["fans"])
	p.Fans = make(map[string]numericsensor.NumericSensor, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 numericsensor.NumericSensor
		v1, err = object.As[numericsensor.NumericSensor](b1, caller)
		if err != nil {
			return err
		}
		p.Fans[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	err = encoding.In("temperatures", j0)
	if err != nil {
		return err
	}
	i2, e2, l2 := encoding.AsMapItems(j0["temperatures"])
	p.Temperatures = make(map[string]numericsensor.NumericSensor, l2)
	for a2, b2 := range i2 {
		var k2 string
		k2, err = encoding.Is[string](a2)
		if err != nil {
			return err
		}
		var v2 numericsensor.NumericSensor
		v2, err = object.As[numericsensor.NumericSensor](b2, caller)
		if err != nil {
			return err
		}
		p.Temperatures[k2] = v2
	}
	err = e2()
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
	return nil
}

func (f *_FanSpeedControlEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.UserEvent = valobj.For[userevent.UserEvent]()
	err := f.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("overrideEnabled", value)
	if err != nil {
		return err
	}
	f.overrideEnabled, err = encoding.Is[bool](value["overrideEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("targetPercent", value)
	if err != nil {
		return err
	}
	f.targetPercent, err = encoding.AsInt32(value["targetPercent"])
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

func (m *_MetaDataChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.Event = valobj.For[event.Event]()
	err := m.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldMetaData", value)
	if err != nil {
		return err
	}
	err = m.oldMetaData.Decode(value["oldMetaData"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newMetaData", value)
	if err != nil {
		return err
	}
	err = m.newMetaData.Decode(value["newMetaData"], caller)
	if err != nil {
		return err
	}
	return nil
}
