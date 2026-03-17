// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package voltagemonitoringsensor

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/idl/event"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
)

func (d *DipSwellThresholds) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["dipActive"] = d.DipActive
	j0["dipThreshold"] = d.DipThreshold
	j0["swellActive"] = d.SwellActive
	j0["swellThreshold"] = d.SwellThreshold
	j0["deassertionHysteresis"] = d.DeassertionHysteresis
	return j0
}

func (d *DipSwellThresholds) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("dipActive", j0)
	if err != nil {
		return err
	}
	d.DipActive, err = encoding.Is[bool](j0["dipActive"])
	if err != nil {
		return err
	}
	err = encoding.In("dipThreshold", j0)
	if err != nil {
		return err
	}
	d.DipThreshold, err = encoding.AsFloat64(j0["dipThreshold"])
	if err != nil {
		return err
	}
	err = encoding.In("swellActive", j0)
	if err != nil {
		return err
	}
	d.SwellActive, err = encoding.Is[bool](j0["swellActive"])
	if err != nil {
		return err
	}
	err = encoding.In("swellThreshold", j0)
	if err != nil {
		return err
	}
	d.SwellThreshold, err = encoding.AsFloat64(j0["swellThreshold"])
	if err != nil {
		return err
	}
	err = encoding.In("deassertionHysteresis", j0)
	if err != nil {
		return err
	}
	d.DeassertionHysteresis, err = encoding.AsFloat32(j0["deassertionHysteresis"])
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["type"] = e.Type
	j0["timestamp"] = e.Timestamp.Unix()
	j0["duration"] = e.Duration
	j0["voltage"] = e.Voltage
	j0["waveform"] = e.Waveform.Encode()
	return j0
}

func (e *Event) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	e.Type, err = encoding.AsEnum[EventType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	e.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("duration", j0)
	if err != nil {
		return err
	}
	e.Duration, err = encoding.AsInt32(j0["duration"])
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	e.Voltage, err = encoding.AsFloat64(j0["voltage"])
	if err != nil {
		return err
	}
	err = encoding.In("waveform", j0)
	if err != nil {
		return err
	}
	err = e.Waveform.Decode(j0["waveform"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (e *_EventOccurredEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.__Event = valobj.For[event.Event]()
	err := e.__Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("event", value)
	if err != nil {
		return err
	}
	err = e.event.Decode(value["event"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (e *_EventListClearedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DipSwellThresholdsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.UserEvent = valobj.For[userevent.UserEvent]()
	err := d.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldThresholds", value)
	if err != nil {
		return err
	}
	err = d.oldThresholds.Decode(value["oldThresholds"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newThresholds", value)
	if err != nil {
		return err
	}
	err = d.newThresholds.Decode(value["newThresholds"], caller)
	if err != nil {
		return err
	}
	return nil
}
