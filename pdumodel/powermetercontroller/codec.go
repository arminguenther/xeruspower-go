// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powermetercontroller

import (
	"github.com/arminguenther/xeruspower-go/v40220/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/powermeter"
)

func (s *ScanResult) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["present"] = s.Present
	j0["meterCount"] = s.MeterCount
	return j0
}

func (s *ScanResult) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("present", j0)
	if err != nil {
		return err
	}
	s.Present, err = encoding.Is[bool](j0["present"])
	if err != nil {
		return err
	}
	err = encoding.In("meterCount", j0)
	if err != nil {
		return err
	}
	s.MeterCount, err = encoding.AsInt32(j0["meterCount"])
	if err != nil {
		return err
	}
	return nil
}

func (p *_PowerMeterCreatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.UserEvent = valobj.For[userevent.UserEvent]()
	err := p.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("powerMeter", value)
	if err != nil {
		return err
	}
	p.powerMeter, err = object.As[powermeter.PowerMeter](value["powerMeter"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("config", value)
	if err != nil {
		return err
	}
	err = p.config.Decode(value["config"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = p.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_PowerMeterDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.UserEvent = valobj.For[userevent.UserEvent]()
	err := p.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("config", value)
	if err != nil {
		return err
	}
	err = p.config.Decode(value["config"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = p.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_PanelCreatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.PowerMeterCreatedEvent = valobj.For[PowerMeterCreatedEvent]()
	err := p.PowerMeterCreatedEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("panelSettings", value)
	if err != nil {
		return err
	}
	err = p.panelSettings.Decode(value["panelSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_PanelDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.PowerMeterDeletedEvent = valobj.For[PowerMeterDeletedEvent]()
	err := p.PowerMeterDeletedEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("panelSettings", value)
	if err != nil {
		return err
	}
	err = p.panelSettings.Decode(value["panelSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}
