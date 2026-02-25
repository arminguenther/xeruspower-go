// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstrip

import (
	"github.com/arminguenther/xeruspower-go/assetmgrmodel/assetstripconfig"
	"github.com/arminguenther/xeruspower-go/event/userevent"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() BladeOverflowChangedEvent { return &_BladeOverflowChangedEvent{} })
	valobj.Register(func() CompositionChangedEvent { return &_CompositionChangedEvent{} })
	valobj.Register(func() FirmwareUpdateStateChangedEvent { return &_FirmwareUpdateStateChangedEvent{} })
	valobj.Register(func() OrientationChangedEvent { return &_OrientationChangedEvent{} })
	valobj.Register(func() RackUnitChangedEvent { return &_RackUnitChangedEvent{} })
	valobj.Register(func() StateChangedEvent { return &_StateChangedEvent{} })
	valobj.Register(func() StripInfoChangedEvent { return &_StripInfoChangedEvent{} })
	valobj.Register(func() TagAddedEvent { return &_TagAddedEvent{} })
	valobj.Register(func() TagEvent { return &_TagEvent{} })
	valobj.Register(func() TagRemovedEvent { return &_TagRemovedEvent{} })
}

type _StripInfoChangedEvent struct {
	event.Event
	oldInfo StripInfo
	newInfo StripInfo
}

func (s *_StripInfoChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.StripInfoChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StripInfoChangedEvent) OldInfo() StripInfo {
	return s.oldInfo
}

func (s *_StripInfoChangedEvent) NewInfo() StripInfo {
	return s.newInfo
}

func (s *_StripInfoChangedEvent) isStripInfoChangedEvent() {}

type _StateChangedEvent struct {
	event.Event
	oldState   State
	newState   State
	deviceInfo DeviceInfo
}

func (s *_StateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.StateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StateChangedEvent) OldState() State {
	return s.oldState
}

func (s *_StateChangedEvent) NewState() State {
	return s.newState
}

func (s *_StateChangedEvent) DeviceInfo() DeviceInfo {
	return s.deviceInfo
}

func (s *_StateChangedEvent) isStateChangedEvent() {}

type _RackUnitChangedEvent struct {
	userevent.UserEvent
	rackUnitNumber int32
	rackUnit       RackUnitInfo
}

func (r *_RackUnitChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.RackUnitChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RackUnitChangedEvent) RackUnitNumber() int32 {
	return r.rackUnitNumber
}

func (r *_RackUnitChangedEvent) RackUnit() RackUnitInfo {
	return r.rackUnit
}

func (r *_RackUnitChangedEvent) isRackUnitChangedEvent() {}

type _TagEvent struct {
	event.Event
	tags    []TagChangeInfo
	allTags []TagInfo
}

func (t *_TagEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.TagEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_TagEvent) Tags() []TagChangeInfo {
	return t.tags
}

func (t *_TagEvent) AllTags() []TagInfo {
	return t.allTags
}

func (t *_TagEvent) isTagEvent() {}

type _TagAddedEvent struct {
	TagEvent
}

func (t *_TagAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.TagAddedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_TagAddedEvent) isTagAddedEvent() {}

type _TagRemovedEvent struct {
	TagEvent
}

func (t *_TagRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.TagRemovedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_TagRemovedEvent) isTagRemovedEvent() {}

type _FirmwareUpdateStateChangedEvent struct {
	event.Event
	state FirmwareUpdateState
}

func (f *_FirmwareUpdateStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.FirmwareUpdateStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FirmwareUpdateStateChangedEvent) State() FirmwareUpdateState {
	return f.state
}

func (f *_FirmwareUpdateStateChangedEvent) isFirmwareUpdateStateChangedEvent() {}

type _BladeOverflowChangedEvent struct {
	event.Event
	overflow bool
}

func (b *_BladeOverflowChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.BladeOverflowChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (b *_BladeOverflowChangedEvent) Overflow() bool {
	return b.overflow
}

func (b *_BladeOverflowChangedEvent) isBladeOverflowChangedEvent() {}

type _OrientationChangedEvent struct {
	event.Event
	oldOrientation assetstripconfig.Orientation
	newOrientation assetstripconfig.Orientation
}

func (o *_OrientationChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.OrientationChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (o *_OrientationChangedEvent) OldOrientation() assetstripconfig.Orientation {
	return o.oldOrientation
}

func (o *_OrientationChangedEvent) NewOrientation() assetstripconfig.Orientation {
	return o.newOrientation
}

func (o *_OrientationChangedEvent) isOrientationChangedEvent() {}

type _CompositionChangedEvent struct {
	event.Event
	oldComponentCount int32
	newComponentCount int32
}

func (c *_CompositionChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "assetmgrmodel.AssetStrip_2_0_6.CompositionChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CompositionChangedEvent) OldComponentCount() int32 {
	return c.oldComponentCount
}

func (c *_CompositionChangedEvent) NewComponentCount() int32 {
	return c.newComponentCount
}

func (c *_CompositionChangedEvent) isCompositionChangedEvent() {}
