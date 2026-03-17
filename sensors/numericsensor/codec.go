// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package numericsensor

import (
	"github.com/arminguenther/xeruspower-go/v40010/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/idl/event"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/valobj"
)

func (r *Range) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["lower"] = r.Lower
	j0["upper"] = r.Upper
	return j0
}

func (r *Range) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("lower", j0)
	if err != nil {
		return err
	}
	r.Lower, err = encoding.AsFloat64(j0["lower"])
	if err != nil {
		return err
	}
	err = encoding.In("upper", j0)
	if err != nil {
		return err
	}
	r.Upper, err = encoding.AsFloat64(j0["upper"])
	if err != nil {
		return err
	}
	return nil
}

func (t *ThresholdCapabilities) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["hasUpperCritical"] = t.HasUpperCritical
	j0["hasUpperWarning"] = t.HasUpperWarning
	j0["hasLowerWarning"] = t.HasLowerWarning
	j0["hasLowerCritical"] = t.HasLowerCritical
	return j0
}

func (t *ThresholdCapabilities) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("hasUpperCritical", j0)
	if err != nil {
		return err
	}
	t.HasUpperCritical, err = encoding.Is[bool](j0["hasUpperCritical"])
	if err != nil {
		return err
	}
	err = encoding.In("hasUpperWarning", j0)
	if err != nil {
		return err
	}
	t.HasUpperWarning, err = encoding.Is[bool](j0["hasUpperWarning"])
	if err != nil {
		return err
	}
	err = encoding.In("hasLowerWarning", j0)
	if err != nil {
		return err
	}
	t.HasLowerWarning, err = encoding.Is[bool](j0["hasLowerWarning"])
	if err != nil {
		return err
	}
	err = encoding.In("hasLowerCritical", j0)
	if err != nil {
		return err
	}
	t.HasLowerCritical, err = encoding.Is[bool](j0["hasLowerCritical"])
	if err != nil {
		return err
	}
	return nil
}

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["type"] = m.Type.Encode()
	j0["decdigits"] = m.Decdigits
	j0["accuracy"] = m.Accuracy
	j0["resolution"] = m.Resolution
	j0["tolerance"] = m.Tolerance
	j0["noiseThreshold"] = m.NoiseThreshold
	j0["range"] = m.Range.Encode()
	j0["thresholdCaps"] = m.ThresholdCaps.Encode()
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	err = m.Type.Decode(j0["type"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("decdigits", j0)
	if err != nil {
		return err
	}
	m.Decdigits, err = encoding.AsInt32(j0["decdigits"])
	if err != nil {
		return err
	}
	err = encoding.In("accuracy", j0)
	if err != nil {
		return err
	}
	m.Accuracy, err = encoding.AsFloat32(j0["accuracy"])
	if err != nil {
		return err
	}
	err = encoding.In("resolution", j0)
	if err != nil {
		return err
	}
	m.Resolution, err = encoding.AsFloat32(j0["resolution"])
	if err != nil {
		return err
	}
	err = encoding.In("tolerance", j0)
	if err != nil {
		return err
	}
	m.Tolerance, err = encoding.AsFloat32(j0["tolerance"])
	if err != nil {
		return err
	}
	err = encoding.In("noiseThreshold", j0)
	if err != nil {
		return err
	}
	m.NoiseThreshold, err = encoding.AsFloat32(j0["noiseThreshold"])
	if err != nil {
		return err
	}
	err = encoding.In("range", j0)
	if err != nil {
		return err
	}
	err = m.Range.Decode(j0["range"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("thresholdCaps", j0)
	if err != nil {
		return err
	}
	err = m.ThresholdCaps.Decode(j0["thresholdCaps"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (t *Thresholds) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["upperCriticalActive"] = t.UpperCriticalActive
	j0["upperCritical"] = t.UpperCritical
	j0["upperWarningActive"] = t.UpperWarningActive
	j0["upperWarning"] = t.UpperWarning
	j0["lowerWarningActive"] = t.LowerWarningActive
	j0["lowerWarning"] = t.LowerWarning
	j0["lowerCriticalActive"] = t.LowerCriticalActive
	j0["lowerCritical"] = t.LowerCritical
	j0["assertionTimeout"] = t.AssertionTimeout
	j0["deassertionHysteresis"] = t.DeassertionHysteresis
	return j0
}

func (t *Thresholds) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("upperCriticalActive", j0)
	if err != nil {
		return err
	}
	t.UpperCriticalActive, err = encoding.Is[bool](j0["upperCriticalActive"])
	if err != nil {
		return err
	}
	err = encoding.In("upperCritical", j0)
	if err != nil {
		return err
	}
	t.UpperCritical, err = encoding.AsFloat64(j0["upperCritical"])
	if err != nil {
		return err
	}
	err = encoding.In("upperWarningActive", j0)
	if err != nil {
		return err
	}
	t.UpperWarningActive, err = encoding.Is[bool](j0["upperWarningActive"])
	if err != nil {
		return err
	}
	err = encoding.In("upperWarning", j0)
	if err != nil {
		return err
	}
	t.UpperWarning, err = encoding.AsFloat64(j0["upperWarning"])
	if err != nil {
		return err
	}
	err = encoding.In("lowerWarningActive", j0)
	if err != nil {
		return err
	}
	t.LowerWarningActive, err = encoding.Is[bool](j0["lowerWarningActive"])
	if err != nil {
		return err
	}
	err = encoding.In("lowerWarning", j0)
	if err != nil {
		return err
	}
	t.LowerWarning, err = encoding.AsFloat64(j0["lowerWarning"])
	if err != nil {
		return err
	}
	err = encoding.In("lowerCriticalActive", j0)
	if err != nil {
		return err
	}
	t.LowerCriticalActive, err = encoding.Is[bool](j0["lowerCriticalActive"])
	if err != nil {
		return err
	}
	err = encoding.In("lowerCritical", j0)
	if err != nil {
		return err
	}
	t.LowerCritical, err = encoding.AsFloat64(j0["lowerCritical"])
	if err != nil {
		return err
	}
	err = encoding.In("assertionTimeout", j0)
	if err != nil {
		return err
	}
	t.AssertionTimeout, err = encoding.AsInt32(j0["assertionTimeout"])
	if err != nil {
		return err
	}
	err = encoding.In("deassertionHysteresis", j0)
	if err != nil {
		return err
	}
	t.DeassertionHysteresis, err = encoding.AsFloat32(j0["deassertionHysteresis"])
	if err != nil {
		return err
	}
	return nil
}

func (r *Reading) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["timestamp"] = r.Timestamp.Unix()
	j0["available"] = r.Available
	j1 := make(map[string]any, 4)
	j1["aboveUpperCritical"] = r.Status.AboveUpperCritical
	j1["aboveUpperWarning"] = r.Status.AboveUpperWarning
	j1["belowLowerWarning"] = r.Status.BelowLowerWarning
	j1["belowLowerCritical"] = r.Status.BelowLowerCritical
	j0["status"] = j1
	j0["valid"] = r.Valid
	j0["value"] = r.Value
	return j0
}

func (r *Reading) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	r.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("available", j0)
	if err != nil {
		return err
	}
	r.Available, err = encoding.Is[bool](j0["available"])
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	j1, err := encoding.Is[map[string]any](j0["status"])
	if err != nil {
		return err
	}
	err = encoding.In("aboveUpperCritical", j1)
	if err != nil {
		return err
	}
	r.Status.AboveUpperCritical, err = encoding.Is[bool](j1["aboveUpperCritical"])
	if err != nil {
		return err
	}
	err = encoding.In("aboveUpperWarning", j1)
	if err != nil {
		return err
	}
	r.Status.AboveUpperWarning, err = encoding.Is[bool](j1["aboveUpperWarning"])
	if err != nil {
		return err
	}
	err = encoding.In("belowLowerWarning", j1)
	if err != nil {
		return err
	}
	r.Status.BelowLowerWarning, err = encoding.Is[bool](j1["belowLowerWarning"])
	if err != nil {
		return err
	}
	err = encoding.In("belowLowerCritical", j1)
	if err != nil {
		return err
	}
	r.Status.BelowLowerCritical, err = encoding.Is[bool](j1["belowLowerCritical"])
	if err != nil {
		return err
	}
	err = encoding.In("valid", j0)
	if err != nil {
		return err
	}
	r.Valid, err = encoding.Is[bool](j0["valid"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	r.Value, err = encoding.AsFloat64(j0["value"])
	if err != nil {
		return err
	}
	return nil
}

func (m *MinMax) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["minReading"] = m.MinReading
	j0["minReadingTimestamp"] = m.MinReadingTimestamp.Unix()
	j0["maxReading"] = m.MaxReading
	j0["maxReadingTimestamp"] = m.MaxReadingTimestamp.Unix()
	j0["valid"] = m.Valid
	j0["observedSince"] = m.ObservedSince.Unix()
	return j0
}

func (m *MinMax) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("minReading", j0)
	if err != nil {
		return err
	}
	m.MinReading, err = encoding.AsFloat64(j0["minReading"])
	if err != nil {
		return err
	}
	err = encoding.In("minReadingTimestamp", j0)
	if err != nil {
		return err
	}
	m.MinReadingTimestamp, err = encoding.AsTime(j0["minReadingTimestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("maxReading", j0)
	if err != nil {
		return err
	}
	m.MaxReading, err = encoding.AsFloat64(j0["maxReading"])
	if err != nil {
		return err
	}
	err = encoding.In("maxReadingTimestamp", j0)
	if err != nil {
		return err
	}
	m.MaxReadingTimestamp, err = encoding.AsTime(j0["maxReadingTimestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("valid", j0)
	if err != nil {
		return err
	}
	m.Valid, err = encoding.Is[bool](j0["valid"])
	if err != nil {
		return err
	}
	err = encoding.In("observedSince", j0)
	if err != nil {
		return err
	}
	m.ObservedSince, err = encoding.AsTime(j0["observedSince"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_ReadingChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.Event = valobj.For[event.Event]()
	err := r.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newReading", value)
	if err != nil {
		return err
	}
	err = r.newReading.Decode(value["newReading"], caller)
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
	err = encoding.In("oldReading", value)
	if err != nil {
		return err
	}
	err = s.oldReading.Decode(value["oldReading"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newReading", value)
	if err != nil {
		return err
	}
	err = s.newReading.Decode(value["newReading"], caller)
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

func (t *_ThresholdsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	t.UserEvent = valobj.For[userevent.UserEvent]()
	err := t.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldThresholds", value)
	if err != nil {
		return err
	}
	err = t.oldThresholds.Decode(value["oldThresholds"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newThresholds", value)
	if err != nil {
		return err
	}
	err = t.newThresholds.Decode(value["newThresholds"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_MinMaxChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.Event = valobj.For[event.Event]()
	err := m.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newMinMax", value)
	if err != nil {
		return err
	}
	err = m.newMinMax.Decode(value["newMinMax"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_MinMaxResetEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.UserEvent = valobj.For[userevent.UserEvent]()
	err := m.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldMinMax", value)
	if err != nil {
		return err
	}
	err = m.oldMinMax.Decode(value["oldMinMax"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newMinMax", value)
	if err != nil {
		return err
	}
	err = m.newMinMax.Decode(value["newMinMax"], caller)
	if err != nil {
		return err
	}
	return nil
}
