// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package numericsensor

import (
	"github.com/arminguenther/xeruspower-go/v40410/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/idl/event"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DefaultThresholdsChangedEvent { return &_DefaultThresholdsChangedEvent{} })
	valobj.Register(func() MetaDataChangedEvent { return &_MetaDataChangedEvent{} })
	valobj.Register(func() MinMaxChangedEvent { return &_MinMaxChangedEvent{} })
	valobj.Register(func() MinMaxResetEvent { return &_MinMaxResetEvent{} })
	valobj.Register(func() ReadingChangedEvent { return &_ReadingChangedEvent{} })
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
	valobj.Register(func() ThresholdsChangedEvent { return &_ThresholdsChangedEvent{} })
}

type _ReadingChangedEvent struct {
	event.Event
	newReading Reading
}

func (r *_ReadingChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.ReadingChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_ReadingChangedEvent) NewReading() Reading {
	return r.newReading
}

func (r *_ReadingChangedEvent) isReadingChangedEvent() {}

type _StateChangedEvent struct {
	event.Event
	oldReading Reading
	newReading Reading
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.StateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StateChangedEvent) OldReading() Reading {
	return s.oldReading
}

func (s *_StateChangedEvent) NewReading() Reading {
	return s.newReading
}

func (s *_StateChangedEvent) isStateChangedEvent() {}

type _MetaDataChangedEvent struct {
	event.Event
	oldMetaData MetaData
	newMetaData MetaData
}

func (m *_MetaDataChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.MetaDataChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MetaDataChangedEvent) OldMetaData() MetaData {
	return m.oldMetaData
}

func (m *_MetaDataChangedEvent) NewMetaData() MetaData {
	return m.newMetaData
}

func (m *_MetaDataChangedEvent) isMetaDataChangedEvent() {}

type _DefaultThresholdsChangedEvent struct {
	userevent.UserEvent
	oldDefaultThresholds Thresholds
	newDefaultThresholds Thresholds
}

func (d *_DefaultThresholdsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.DefaultThresholdsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DefaultThresholdsChangedEvent) OldDefaultThresholds() Thresholds {
	return d.oldDefaultThresholds
}

func (d *_DefaultThresholdsChangedEvent) NewDefaultThresholds() Thresholds {
	return d.newDefaultThresholds
}

func (d *_DefaultThresholdsChangedEvent) isDefaultThresholdsChangedEvent() {}

type _ThresholdsChangedEvent struct {
	userevent.UserEvent
	oldThresholds Thresholds
	newThresholds Thresholds
}

func (t *_ThresholdsChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.ThresholdsChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_ThresholdsChangedEvent) OldThresholds() Thresholds {
	return t.oldThresholds
}

func (t *_ThresholdsChangedEvent) NewThresholds() Thresholds {
	return t.newThresholds
}

func (t *_ThresholdsChangedEvent) isThresholdsChangedEvent() {}

type _MinMaxChangedEvent struct {
	event.Event
	newMinMax MinMax
}

func (m *_MinMaxChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.MinMaxChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MinMaxChangedEvent) NewMinMax() MinMax {
	return m.newMinMax
}

func (m *_MinMaxChangedEvent) isMinMaxChangedEvent() {}

type _MinMaxResetEvent struct {
	userevent.UserEvent
	oldMinMax MinMax
	newMinMax MinMax
}

func (m *_MinMaxResetEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sensors.NumericSensor_4_0_8.MinMaxResetEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_MinMaxResetEvent) OldMinMax() MinMax {
	return m.oldMinMax
}

func (m *_MinMaxResetEvent) NewMinMax() MinMax {
	return m.newMinMax
}

func (m *_MinMaxResetEvent) isMinMaxResetEvent() {}
