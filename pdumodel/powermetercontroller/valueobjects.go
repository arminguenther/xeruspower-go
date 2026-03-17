// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powermetercontroller

import (
	"github.com/arminguenther/xeruspower-go/v40010/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40010/pdumodel/panel"
	"github.com/arminguenther/xeruspower-go/v40010/pdumodel/powermeter"
)

func init() {
	valobj.Register(func() PanelCreatedEvent { return &_PanelCreatedEvent{} })
	valobj.Register(func() PanelDeletedEvent { return &_PanelDeletedEvent{} })
	valobj.Register(func() PowerMeterCreatedEvent { return &_PowerMeterCreatedEvent{} })
	valobj.Register(func() PowerMeterDeletedEvent { return &_PowerMeterDeletedEvent{} })
}

type _PowerMeterCreatedEvent struct {
	userevent.UserEvent
	powerMeter powermeter.PowerMeter
	config     powermeter.Config
	settings   powermeter.Settings
}

func (p *_PowerMeterCreatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerMeterController_1_2_8.PowerMeterCreatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PowerMeterCreatedEvent) PowerMeter() powermeter.PowerMeter {
	return p.powerMeter
}

func (p *_PowerMeterCreatedEvent) Config() powermeter.Config {
	return p.config
}

func (p *_PowerMeterCreatedEvent) Settings() powermeter.Settings {
	return p.settings
}

func (p *_PowerMeterCreatedEvent) isPowerMeterCreatedEvent() {}

type _PowerMeterDeletedEvent struct {
	userevent.UserEvent
	config   powermeter.Config
	settings powermeter.Settings
}

func (p *_PowerMeterDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerMeterController_1_2_8.PowerMeterDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PowerMeterDeletedEvent) Config() powermeter.Config {
	return p.config
}

func (p *_PowerMeterDeletedEvent) Settings() powermeter.Settings {
	return p.settings
}

func (p *_PowerMeterDeletedEvent) isPowerMeterDeletedEvent() {}

type _PanelCreatedEvent struct {
	PowerMeterCreatedEvent
	panelSettings panel.Settings
}

func (p *_PanelCreatedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerMeterController_1_2_8.PanelCreatedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PanelCreatedEvent) PanelSettings() panel.Settings {
	return p.panelSettings
}

func (p *_PanelCreatedEvent) isPanelCreatedEvent() {}

type _PanelDeletedEvent struct {
	PowerMeterDeletedEvent
	panelSettings panel.Settings
}

func (p *_PanelDeletedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.PowerMeterController_1_2_8.PanelDeletedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (p *_PanelDeletedEvent) PanelSettings() panel.Settings {
	return p.panelSettings
}

func (p *_PanelDeletedEvent) isPanelDeletedEvent() {}
