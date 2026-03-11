// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package pole

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40413/pdumodel/residualcurrentstatesensor"
	"github.com/arminguenther/xeruspower-go/v40413/sensors/numericsensor"
)

func (p *Pole) Encode() map[string]any {
	j0 := make(map[string]any, 22)
	j0["label"] = p.Label
	j0["line"] = p.Line
	j0["nodeId"] = p.NodeId
	j0["voltage"] = object.ToMap(p.Voltage)
	j0["voltageLN"] = object.ToMap(p.VoltageLN)
	j0["current"] = object.ToMap(p.Current)
	j0["peakCurrent"] = object.ToMap(p.PeakCurrent)
	j0["activePower"] = object.ToMap(p.ActivePower)
	j0["reactivePower"] = object.ToMap(p.ReactivePower)
	j0["apparentPower"] = object.ToMap(p.ApparentPower)
	j0["powerFactor"] = object.ToMap(p.PowerFactor)
	j0["phaseAngle"] = object.ToMap(p.PhaseAngle)
	j0["displacementPowerFactor"] = object.ToMap(p.DisplacementPowerFactor)
	j0["activeEnergy"] = object.ToMap(p.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(p.ApparentEnergy)
	j0["residualCurrent"] = object.ToMap(p.ResidualCurrent)
	j0["residualACCurrent"] = object.ToMap(p.ResidualACCurrent)
	j0["residualDCCurrent"] = object.ToMap(p.ResidualDCCurrent)
	j0["crestFactor"] = object.ToMap(p.CrestFactor)
	j0["voltageThd"] = object.ToMap(p.VoltageThd)
	j0["currentThd"] = object.ToMap(p.CurrentThd)
	j0["residualCurrentStatus"] = object.ToMap(p.ResidualCurrentStatus)
	return j0
}

func (p *Pole) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	p.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("line", j0)
	if err != nil {
		return err
	}
	p.Line, err = encoding.AsEnum[PowerLine](j0["line"])
	if err != nil {
		return err
	}
	err = encoding.In("nodeId", j0)
	if err != nil {
		return err
	}
	p.NodeId, err = encoding.AsInt32(j0["nodeId"])
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	p.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("voltageLN", j0)
	if err != nil {
		return err
	}
	p.VoltageLN, err = object.As[numericsensor.NumericSensor](j0["voltageLN"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	p.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("peakCurrent", j0)
	if err != nil {
		return err
	}
	p.PeakCurrent, err = object.As[numericsensor.NumericSensor](j0["peakCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	p.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("reactivePower", j0)
	if err != nil {
		return err
	}
	p.ReactivePower, err = object.As[numericsensor.NumericSensor](j0["reactivePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentPower", j0)
	if err != nil {
		return err
	}
	p.ApparentPower, err = object.As[numericsensor.NumericSensor](j0["apparentPower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("powerFactor", j0)
	if err != nil {
		return err
	}
	p.PowerFactor, err = object.As[numericsensor.NumericSensor](j0["powerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("phaseAngle", j0)
	if err != nil {
		return err
	}
	p.PhaseAngle, err = object.As[numericsensor.NumericSensor](j0["phaseAngle"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("displacementPowerFactor", j0)
	if err != nil {
		return err
	}
	p.DisplacementPowerFactor, err = object.As[numericsensor.NumericSensor](j0["displacementPowerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	p.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentEnergy", j0)
	if err != nil {
		return err
	}
	p.ApparentEnergy, err = object.As[numericsensor.NumericSensor](j0["apparentEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualCurrent", j0)
	if err != nil {
		return err
	}
	p.ResidualCurrent, err = object.As[numericsensor.NumericSensor](j0["residualCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualACCurrent", j0)
	if err != nil {
		return err
	}
	p.ResidualACCurrent, err = object.As[numericsensor.NumericSensor](j0["residualACCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualDCCurrent", j0)
	if err != nil {
		return err
	}
	p.ResidualDCCurrent, err = object.As[numericsensor.NumericSensor](j0["residualDCCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("crestFactor", j0)
	if err != nil {
		return err
	}
	p.CrestFactor, err = object.As[numericsensor.NumericSensor](j0["crestFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("voltageThd", j0)
	if err != nil {
		return err
	}
	p.VoltageThd, err = object.As[numericsensor.NumericSensor](j0["voltageThd"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("currentThd", j0)
	if err != nil {
		return err
	}
	p.CurrentThd, err = object.As[numericsensor.NumericSensor](j0["currentThd"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("residualCurrentStatus", j0)
	if err != nil {
		return err
	}
	p.ResidualCurrentStatus, err = object.As[residualcurrentstatesensor.ResidualCurrentStateSensor](j0["residualCurrentStatus"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *MeteredLinePair) Encode() map[string]any {
	j0 := make(map[string]any, 18)
	j0["leftLine"] = m.LeftLine
	j0["rightLine"] = m.RightLine
	j0["leftNodeId"] = m.LeftNodeId
	j0["rightNodeId"] = m.RightNodeId
	j0["voltage"] = object.ToMap(m.Voltage)
	j0["current"] = object.ToMap(m.Current)
	j0["peakCurrent"] = object.ToMap(m.PeakCurrent)
	j0["activePower"] = object.ToMap(m.ActivePower)
	j0["reactivePower"] = object.ToMap(m.ReactivePower)
	j0["apparentPower"] = object.ToMap(m.ApparentPower)
	j0["powerFactor"] = object.ToMap(m.PowerFactor)
	j0["phaseAngle"] = object.ToMap(m.PhaseAngle)
	j0["displacementPowerFactor"] = object.ToMap(m.DisplacementPowerFactor)
	j0["activeEnergy"] = object.ToMap(m.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(m.ApparentEnergy)
	j0["crestFactor"] = object.ToMap(m.CrestFactor)
	j0["voltageThd"] = object.ToMap(m.VoltageThd)
	j0["currentThd"] = object.ToMap(m.CurrentThd)
	return j0
}

func (m *MeteredLinePair) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("leftLine", j0)
	if err != nil {
		return err
	}
	m.LeftLine, err = encoding.AsEnum[PowerLine](j0["leftLine"])
	if err != nil {
		return err
	}
	err = encoding.In("rightLine", j0)
	if err != nil {
		return err
	}
	m.RightLine, err = encoding.AsEnum[PowerLine](j0["rightLine"])
	if err != nil {
		return err
	}
	err = encoding.In("leftNodeId", j0)
	if err != nil {
		return err
	}
	m.LeftNodeId, err = encoding.AsInt32(j0["leftNodeId"])
	if err != nil {
		return err
	}
	err = encoding.In("rightNodeId", j0)
	if err != nil {
		return err
	}
	m.RightNodeId, err = encoding.AsInt32(j0["rightNodeId"])
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	m.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	m.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("peakCurrent", j0)
	if err != nil {
		return err
	}
	m.PeakCurrent, err = object.As[numericsensor.NumericSensor](j0["peakCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	m.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("reactivePower", j0)
	if err != nil {
		return err
	}
	m.ReactivePower, err = object.As[numericsensor.NumericSensor](j0["reactivePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentPower", j0)
	if err != nil {
		return err
	}
	m.ApparentPower, err = object.As[numericsensor.NumericSensor](j0["apparentPower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("powerFactor", j0)
	if err != nil {
		return err
	}
	m.PowerFactor, err = object.As[numericsensor.NumericSensor](j0["powerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("phaseAngle", j0)
	if err != nil {
		return err
	}
	m.PhaseAngle, err = object.As[numericsensor.NumericSensor](j0["phaseAngle"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("displacementPowerFactor", j0)
	if err != nil {
		return err
	}
	m.DisplacementPowerFactor, err = object.As[numericsensor.NumericSensor](j0["displacementPowerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	m.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentEnergy", j0)
	if err != nil {
		return err
	}
	m.ApparentEnergy, err = object.As[numericsensor.NumericSensor](j0["apparentEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("crestFactor", j0)
	if err != nil {
		return err
	}
	m.CrestFactor, err = object.As[numericsensor.NumericSensor](j0["crestFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("voltageThd", j0)
	if err != nil {
		return err
	}
	m.VoltageThd, err = object.As[numericsensor.NumericSensor](j0["voltageThd"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("currentThd", j0)
	if err != nil {
		return err
	}
	m.CurrentThd, err = object.As[numericsensor.NumericSensor](j0["currentThd"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *DoublePole) Encode() map[string]any {
	j0 := make(map[string]any, 13)
	j0["label"] = d.Label
	j0["line"] = d.Line
	j0["inNodeId"] = d.InNodeId
	j0["outNodeId"] = d.OutNodeId
	j0["voltage"] = object.ToMap(d.Voltage)
	j0["voltageLN"] = object.ToMap(d.VoltageLN)
	j0["current"] = object.ToMap(d.Current)
	j0["peakCurrent"] = object.ToMap(d.PeakCurrent)
	j0["activePower"] = object.ToMap(d.ActivePower)
	j0["apparentPower"] = object.ToMap(d.ApparentPower)
	j0["powerFactor"] = object.ToMap(d.PowerFactor)
	j0["activeEnergy"] = object.ToMap(d.ActiveEnergy)
	j0["apparentEnergy"] = object.ToMap(d.ApparentEnergy)
	return j0
}

func (d *DoublePole) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	d.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("line", j0)
	if err != nil {
		return err
	}
	d.Line, err = encoding.AsEnum[PowerLine](j0["line"])
	if err != nil {
		return err
	}
	err = encoding.In("inNodeId", j0)
	if err != nil {
		return err
	}
	d.InNodeId, err = encoding.AsInt32(j0["inNodeId"])
	if err != nil {
		return err
	}
	err = encoding.In("outNodeId", j0)
	if err != nil {
		return err
	}
	d.OutNodeId, err = encoding.AsInt32(j0["outNodeId"])
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	d.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("voltageLN", j0)
	if err != nil {
		return err
	}
	d.VoltageLN, err = object.As[numericsensor.NumericSensor](j0["voltageLN"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	d.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("peakCurrent", j0)
	if err != nil {
		return err
	}
	d.PeakCurrent, err = object.As[numericsensor.NumericSensor](j0["peakCurrent"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	d.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentPower", j0)
	if err != nil {
		return err
	}
	d.ApparentPower, err = object.As[numericsensor.NumericSensor](j0["apparentPower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("powerFactor", j0)
	if err != nil {
		return err
	}
	d.PowerFactor, err = object.As[numericsensor.NumericSensor](j0["powerFactor"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	d.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("apparentEnergy", j0)
	if err != nil {
		return err
	}
	d.ApparentEnergy, err = object.As[numericsensor.NumericSensor](j0["apparentEnergy"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (t *ThrowPole) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["label"] = t.Label
	j0["line"] = t.Line
	j0["inNodeIds"] = encoding.NonNilSlice(t.InNodeIds)
	j0["outNodeId"] = t.OutNodeId
	return j0
}

func (t *ThrowPole) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	t.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("line", j0)
	if err != nil {
		return err
	}
	t.Line, err = encoding.AsEnum[PowerLine](j0["line"])
	if err != nil {
		return err
	}
	err = encoding.In("inNodeIds", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["inNodeIds"])
	if err != nil {
		return err
	}
	t.InNodeIds = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		t.InNodeIds = append(t.InNodeIds, e1)
	}
	err = encoding.In("outNodeId", j0)
	if err != nil {
		return err
	}
	t.OutNodeId, err = encoding.AsInt32(j0["outNodeId"])
	if err != nil {
		return err
	}
	return nil
}
