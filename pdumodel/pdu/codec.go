// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package pdu

import (
	"github.com/arminguenther/xeruspower-go/v40040/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/controller"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40040/pdumodel/overcurrentprotector"
	"github.com/arminguenther/xeruspower-go/v40040/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40040/sensors/statesensor"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 12)
	j0["nameplate"] = m.Nameplate.Encode()
	j0["ctrlBoardSerial"] = m.CtrlBoardSerial
	j0["hwRevision"] = m.HwRevision
	j0["fwRevision"] = m.FwRevision
	j0["macAddress"] = m.MacAddress
	j0["hasSwitchableOutlets"] = m.HasSwitchableOutlets
	j0["hasMeteredOutlets"] = m.HasMeteredOutlets
	j0["hasLatchingOutletRelays"] = m.HasLatchingOutletRelays
	j0["isInlineMeter"] = m.IsInlineMeter
	j0["isEnergyPulseSupported"] = m.IsEnergyPulseSupported
	j0["hasDCInlets"] = m.HasDCInlets
	j0["pduOrientation"] = m.PduOrientation
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("nameplate", j0)
	if err != nil {
		return err
	}
	err = m.Nameplate.Decode(j0["nameplate"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("ctrlBoardSerial", j0)
	if err != nil {
		return err
	}
	m.CtrlBoardSerial, err = encoding.Is[string](j0["ctrlBoardSerial"])
	if err != nil {
		return err
	}
	err = encoding.In("hwRevision", j0)
	if err != nil {
		return err
	}
	m.HwRevision, err = encoding.Is[string](j0["hwRevision"])
	if err != nil {
		return err
	}
	err = encoding.In("fwRevision", j0)
	if err != nil {
		return err
	}
	m.FwRevision, err = encoding.Is[string](j0["fwRevision"])
	if err != nil {
		return err
	}
	err = encoding.In("macAddress", j0)
	if err != nil {
		return err
	}
	m.MacAddress, err = encoding.Is[string](j0["macAddress"])
	if err != nil {
		return err
	}
	err = encoding.In("hasSwitchableOutlets", j0)
	if err != nil {
		return err
	}
	m.HasSwitchableOutlets, err = encoding.Is[bool](j0["hasSwitchableOutlets"])
	if err != nil {
		return err
	}
	err = encoding.In("hasMeteredOutlets", j0)
	if err != nil {
		return err
	}
	m.HasMeteredOutlets, err = encoding.Is[bool](j0["hasMeteredOutlets"])
	if err != nil {
		return err
	}
	err = encoding.In("hasLatchingOutletRelays", j0)
	if err != nil {
		return err
	}
	m.HasLatchingOutletRelays, err = encoding.Is[bool](j0["hasLatchingOutletRelays"])
	if err != nil {
		return err
	}
	err = encoding.In("isInlineMeter", j0)
	if err != nil {
		return err
	}
	m.IsInlineMeter, err = encoding.Is[bool](j0["isInlineMeter"])
	if err != nil {
		return err
	}
	err = encoding.In("isEnergyPulseSupported", j0)
	if err != nil {
		return err
	}
	m.IsEnergyPulseSupported, err = encoding.Is[bool](j0["isEnergyPulseSupported"])
	if err != nil {
		return err
	}
	err = encoding.In("hasDCInlets", j0)
	if err != nil {
		return err
	}
	m.HasDCInlets, err = encoding.Is[bool](j0["hasDCInlets"])
	if err != nil {
		return err
	}
	err = encoding.In("pduOrientation", j0)
	if err != nil {
		return err
	}
	m.PduOrientation, err = encoding.AsEnum[Orientation](j0["pduOrientation"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	s1 := make([]any, 0, len(s.PowerSupplyStatus))
	for _, e1 := range s.PowerSupplyStatus {
		s1 = append(s1, object.ToMap(e1))
	}
	j0["powerSupplyStatus"] = s1
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
	err = encoding.In("powerSupplyStatus", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["powerSupplyStatus"])
	if err != nil {
		return err
	}
	s.PowerSupplyStatus = make([]statesensor.StateSensor, 0, len(s1))
	for _, a1 := range s1 {
		var e1 statesensor.StateSensor
		e1, err = object.As[statesensor.StateSensor](a1, caller)
		if err != nil {
			return err
		}
		s.PowerSupplyStatus = append(s.PowerSupplyStatus, e1)
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
	return nil
}

func (s *Statistic) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	s1 := make([]any, 0, len(s.CbStats))
	for _, e1 := range s.CbStats {
		s1 = append(s1, e1.Encode())
	}
	j0["cbStats"] = s1
	s2 := make([]any, 0, len(s.CtrlStats))
	for _, e2 := range s.CtrlStats {
		s2 = append(s2, e2.Encode())
	}
	j0["ctrlStats"] = s2
	s3 := make([]any, 0, len(s.OutletStats))
	for _, e3 := range s.OutletStats {
		s3 = append(s3, e3.Encode())
	}
	j0["outletStats"] = s3
	j0["peripheralStats"] = s.PeripheralStats.Encode()
	return j0
}

func (s *Statistic) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("cbStats", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["cbStats"])
	if err != nil {
		return err
	}
	s.CbStats = make([]overcurrentprotector.CircuitBreakerStatistic, 0, len(s1))
	for _, a1 := range s1 {
		var e1 overcurrentprotector.CircuitBreakerStatistic
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		s.CbStats = append(s.CbStats, e1)
	}
	err = encoding.In("ctrlStats", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["ctrlStats"])
	if err != nil {
		return err
	}
	s.CtrlStats = make([]controller.CtrlStatistic, 0, len(s2))
	for _, a2 := range s2 {
		var e2 controller.CtrlStatistic
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		s.CtrlStats = append(s.CtrlStats, e2)
	}
	err = encoding.In("outletStats", j0)
	if err != nil {
		return err
	}
	var s3 []any
	s3, err = encoding.Is[[]any](j0["outletStats"])
	if err != nil {
		return err
	}
	s.OutletStats = make([]outlet.Statistic, 0, len(s3))
	for _, a3 := range s3 {
		var e3 outlet.Statistic
		err = e3.Decode(a3, caller)
		if err != nil {
			return err
		}
		s.OutletStats = append(s.OutletStats, e3)
	}
	err = encoding.In("peripheralStats", j0)
	if err != nil {
		return err
	}
	err = s.PeripheralStats.Decode(j0["peripheralStats"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 12)
	j0["name"] = s.Name
	j0["startupState"] = s.StartupState
	j0["cycleDelay"] = s.CycleDelay
	j0["inRushGuardDelay"] = s.InRushGuardDelay
	j0["outletPowerStateSequence"] = encoding.NonNilSlice(s.OutletPowerStateSequence)
	j0["powerOnDelay"] = s.PowerOnDelay
	j0["latchingRelays"] = s.LatchingRelays
	j0["energyPulseEnabled"] = s.EnergyPulseEnabled
	j0["energyPulsesPerKWh"] = s.EnergyPulsesPerKWh
	j0["demandUpdateInterval"] = s.DemandUpdateInterval
	j0["demandAveragingIntervals"] = s.DemandAveragingIntervals
	j0["suspendTripCauseOutlets"] = s.SuspendTripCauseOutlets
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
	err = encoding.In("cycleDelay", j0)
	if err != nil {
		return err
	}
	s.CycleDelay, err = encoding.AsInt32(j0["cycleDelay"])
	if err != nil {
		return err
	}
	err = encoding.In("inRushGuardDelay", j0)
	if err != nil {
		return err
	}
	s.InRushGuardDelay, err = encoding.AsInt32(j0["inRushGuardDelay"])
	if err != nil {
		return err
	}
	err = encoding.In("outletPowerStateSequence", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["outletPowerStateSequence"])
	if err != nil {
		return err
	}
	s.OutletPowerStateSequence = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		s.OutletPowerStateSequence = append(s.OutletPowerStateSequence, e1)
	}
	err = encoding.In("powerOnDelay", j0)
	if err != nil {
		return err
	}
	s.PowerOnDelay, err = encoding.AsInt32(j0["powerOnDelay"])
	if err != nil {
		return err
	}
	err = encoding.In("latchingRelays", j0)
	if err != nil {
		return err
	}
	s.LatchingRelays, err = encoding.Is[bool](j0["latchingRelays"])
	if err != nil {
		return err
	}
	err = encoding.In("energyPulseEnabled", j0)
	if err != nil {
		return err
	}
	s.EnergyPulseEnabled, err = encoding.Is[bool](j0["energyPulseEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("energyPulsesPerKWh", j0)
	if err != nil {
		return err
	}
	s.EnergyPulsesPerKWh, err = encoding.AsInt32(j0["energyPulsesPerKWh"])
	if err != nil {
		return err
	}
	err = encoding.In("demandUpdateInterval", j0)
	if err != nil {
		return err
	}
	s.DemandUpdateInterval, err = encoding.AsInt32(j0["demandUpdateInterval"])
	if err != nil {
		return err
	}
	err = encoding.In("demandAveragingIntervals", j0)
	if err != nil {
		return err
	}
	s.DemandAveragingIntervals, err = encoding.AsInt32(j0["demandAveragingIntervals"])
	if err != nil {
		return err
	}
	err = encoding.In("suspendTripCauseOutlets", j0)
	if err != nil {
		return err
	}
	s.SuspendTripCauseOutlets, err = encoding.Is[bool](j0["suspendTripCauseOutlets"])
	if err != nil {
		return err
	}
	return nil
}

func (o *OutletSequenceState) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["sequenceRunning"] = o.SequenceRunning
	j0["nextOutletToSwitch"] = o.NextOutletToSwitch
	j0["timeUntilNextSwitch"] = o.TimeUntilNextSwitch
	j0["outletsRemaining"] = o.OutletsRemaining
	j0["cancelableOutletsRemaining"] = o.CancelableOutletsRemaining
	return j0
}

func (o *OutletSequenceState) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sequenceRunning", j0)
	if err != nil {
		return err
	}
	o.SequenceRunning, err = encoding.Is[bool](j0["sequenceRunning"])
	if err != nil {
		return err
	}
	err = encoding.In("nextOutletToSwitch", j0)
	if err != nil {
		return err
	}
	o.NextOutletToSwitch, err = encoding.AsInt32(j0["nextOutletToSwitch"])
	if err != nil {
		return err
	}
	err = encoding.In("timeUntilNextSwitch", j0)
	if err != nil {
		return err
	}
	o.TimeUntilNextSwitch, err = encoding.AsInt32(j0["timeUntilNextSwitch"])
	if err != nil {
		return err
	}
	err = encoding.In("outletsRemaining", j0)
	if err != nil {
		return err
	}
	o.OutletsRemaining, err = encoding.AsInt32(j0["outletsRemaining"])
	if err != nil {
		return err
	}
	err = encoding.In("cancelableOutletsRemaining", j0)
	if err != nil {
		return err
	}
	o.CancelableOutletsRemaining, err = encoding.AsInt32(j0["cancelableOutletsRemaining"])
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

func (l *_LoadSheddingModeChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.UserEvent = valobj.For[userevent.UserEvent]()
	err := l.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", value)
	if err != nil {
		return err
	}
	l.enabled, err = encoding.Is[bool](value["enabled"])
	if err != nil {
		return err
	}
	return nil
}

func (o *_OutletSequenceStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	o.Event = valobj.For[event.Event]()
	err := o.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	err = o.newState.Decode(value["newState"], caller)
	if err != nil {
		return err
	}
	return nil
}
