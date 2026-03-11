// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstrip

import (
	"github.com/arminguenther/xeruspower-go/v40413/assetmgrmodel/assetstripconfig"
	"github.com/arminguenther/xeruspower-go/v40413/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func (d *DeviceInfo) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["deviceId"] = d.DeviceId
	j0["hardwareId"] = d.HardwareId
	j0["protocolVersion"] = d.ProtocolVersion
	j0["bootVersion"] = d.BootVersion
	j0["appVersion"] = d.AppVersion
	j0["orientationSensAvailable"] = d.OrientationSensAvailable
	j0["isCascadable"] = d.IsCascadable
	j0["rackUnitCountConfigurable"] = d.RackUnitCountConfigurable
	return j0
}

func (d *DeviceInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("deviceId", j0)
	if err != nil {
		return err
	}
	d.DeviceId, err = encoding.AsInt32(j0["deviceId"])
	if err != nil {
		return err
	}
	err = encoding.In("hardwareId", j0)
	if err != nil {
		return err
	}
	d.HardwareId, err = encoding.AsInt32(j0["hardwareId"])
	if err != nil {
		return err
	}
	err = encoding.In("protocolVersion", j0)
	if err != nil {
		return err
	}
	d.ProtocolVersion, err = encoding.AsInt32(j0["protocolVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("bootVersion", j0)
	if err != nil {
		return err
	}
	d.BootVersion, err = encoding.AsInt32(j0["bootVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("appVersion", j0)
	if err != nil {
		return err
	}
	d.AppVersion, err = encoding.AsInt32(j0["appVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("orientationSensAvailable", j0)
	if err != nil {
		return err
	}
	d.OrientationSensAvailable, err = encoding.Is[bool](j0["orientationSensAvailable"])
	if err != nil {
		return err
	}
	err = encoding.In("isCascadable", j0)
	if err != nil {
		return err
	}
	d.IsCascadable, err = encoding.Is[bool](j0["isCascadable"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitCountConfigurable", j0)
	if err != nil {
		return err
	}
	d.RackUnitCountConfigurable, err = encoding.Is[bool](j0["rackUnitCountConfigurable"])
	if err != nil {
		return err
	}
	return nil
}

func (s *StripInfo) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["maxMainTagCount"] = s.MaxMainTagCount
	j0["maxBladeTagCount"] = s.MaxBladeTagCount
	j0["mainTagCount"] = s.MainTagCount
	j0["bladeTagCount"] = s.BladeTagCount
	j0["bladeOverflow"] = s.BladeOverflow
	j0["rackUnitCount"] = s.RackUnitCount
	j0["componentCount"] = s.ComponentCount
	j0["cascadeState"] = s.CascadeState
	return j0
}

func (s *StripInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("maxMainTagCount", j0)
	if err != nil {
		return err
	}
	s.MaxMainTagCount, err = encoding.AsInt32(j0["maxMainTagCount"])
	if err != nil {
		return err
	}
	err = encoding.In("maxBladeTagCount", j0)
	if err != nil {
		return err
	}
	s.MaxBladeTagCount, err = encoding.AsInt32(j0["maxBladeTagCount"])
	if err != nil {
		return err
	}
	err = encoding.In("mainTagCount", j0)
	if err != nil {
		return err
	}
	s.MainTagCount, err = encoding.AsInt32(j0["mainTagCount"])
	if err != nil {
		return err
	}
	err = encoding.In("bladeTagCount", j0)
	if err != nil {
		return err
	}
	s.BladeTagCount, err = encoding.AsInt32(j0["bladeTagCount"])
	if err != nil {
		return err
	}
	err = encoding.In("bladeOverflow", j0)
	if err != nil {
		return err
	}
	s.BladeOverflow, err = encoding.Is[bool](j0["bladeOverflow"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitCount", j0)
	if err != nil {
		return err
	}
	s.RackUnitCount, err = encoding.AsInt32(j0["rackUnitCount"])
	if err != nil {
		return err
	}
	err = encoding.In("componentCount", j0)
	if err != nil {
		return err
	}
	s.ComponentCount, err = encoding.AsInt32(j0["componentCount"])
	if err != nil {
		return err
	}
	err = encoding.In("cascadeState", j0)
	if err != nil {
		return err
	}
	s.CascadeState, err = encoding.AsEnum[CascadeState](j0["cascadeState"])
	if err != nil {
		return err
	}
	return nil
}

func (t *TagInfo) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["rackUnitNumber"] = t.RackUnitNumber
	j0["slotNumber"] = t.SlotNumber
	j0["familyDesc"] = t.FamilyDesc
	j0["rawId"] = t.RawId
	j0["programmable"] = t.Programmable
	return j0
}

func (t *TagInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitNumber", j0)
	if err != nil {
		return err
	}
	t.RackUnitNumber, err = encoding.AsInt32(j0["rackUnitNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("slotNumber", j0)
	if err != nil {
		return err
	}
	t.SlotNumber, err = encoding.AsInt32(j0["slotNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("familyDesc", j0)
	if err != nil {
		return err
	}
	t.FamilyDesc, err = encoding.Is[string](j0["familyDesc"])
	if err != nil {
		return err
	}
	err = encoding.In("rawId", j0)
	if err != nil {
		return err
	}
	t.RawId, err = encoding.Is[string](j0["rawId"])
	if err != nil {
		return err
	}
	err = encoding.In("programmable", j0)
	if err != nil {
		return err
	}
	t.Programmable, err = encoding.Is[bool](j0["programmable"])
	if err != nil {
		return err
	}
	return nil
}

func (r *RackUnitInfo) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["rackUnitNumber"] = r.RackUnitNumber
	j0["rackUnitPosition"] = r.RackUnitPosition
	j0["type"] = r.Type
	j0["size"] = r.Size
	j0["settings"] = r.Settings.Encode()
	j0["assetStripCascadePosition"] = r.AssetStripCascadePosition
	j0["rackUnitRelativePosition"] = r.RackUnitRelativePosition
	j0["assetStripNumberOfRackUnits"] = r.AssetStripNumberOfRackUnits
	return j0
}

func (r *RackUnitInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitNumber", j0)
	if err != nil {
		return err
	}
	r.RackUnitNumber, err = encoding.AsInt32(j0["rackUnitNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitPosition", j0)
	if err != nil {
		return err
	}
	r.RackUnitPosition, err = encoding.AsInt32(j0["rackUnitPosition"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	r.Type, err = encoding.AsEnum[TagType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("size", j0)
	if err != nil {
		return err
	}
	r.Size, err = encoding.AsInt32(j0["size"])
	if err != nil {
		return err
	}
	err = encoding.In("settings", j0)
	if err != nil {
		return err
	}
	err = r.Settings.Decode(j0["settings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("assetStripCascadePosition", j0)
	if err != nil {
		return err
	}
	r.AssetStripCascadePosition, err = encoding.AsInt32(j0["assetStripCascadePosition"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitRelativePosition", j0)
	if err != nil {
		return err
	}
	r.RackUnitRelativePosition, err = encoding.AsInt32(j0["rackUnitRelativePosition"])
	if err != nil {
		return err
	}
	err = encoding.In("assetStripNumberOfRackUnits", j0)
	if err != nil {
		return err
	}
	r.AssetStripNumberOfRackUnits, err = encoding.AsInt32(j0["assetStripNumberOfRackUnits"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_StripInfoChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldInfo", value)
	if err != nil {
		return err
	}
	err = s.oldInfo.Decode(value["oldInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newInfo", value)
	if err != nil {
		return err
	}
	err = s.newInfo.Decode(value["newInfo"], caller)
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
	s.oldState, err = encoding.AsEnum[State](value["oldState"])
	if err != nil {
		return err
	}
	err = encoding.In("newState", value)
	if err != nil {
		return err
	}
	s.newState, err = encoding.AsEnum[State](value["newState"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceInfo", value)
	if err != nil {
		return err
	}
	err = s.deviceInfo.Decode(value["deviceInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *_RackUnitChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitNumber", value)
	if err != nil {
		return err
	}
	r.rackUnitNumber, err = encoding.AsInt32(value["rackUnitNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnit", value)
	if err != nil {
		return err
	}
	err = r.rackUnit.Decode(value["rackUnit"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (t *TagChangeInfo) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["tag"] = t.Tag.Encode()
	j0["info"] = t.Info.Encode()
	j0["parentBladeTagId"] = t.ParentBladeTagId
	j0["slotPosition"] = t.SlotPosition
	return j0
}

func (t *TagChangeInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("tag", j0)
	if err != nil {
		return err
	}
	err = t.Tag.Decode(j0["tag"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("info", j0)
	if err != nil {
		return err
	}
	err = t.Info.Decode(j0["info"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("parentBladeTagId", j0)
	if err != nil {
		return err
	}
	t.ParentBladeTagId, err = encoding.Is[string](j0["parentBladeTagId"])
	if err != nil {
		return err
	}
	err = encoding.In("slotPosition", j0)
	if err != nil {
		return err
	}
	t.SlotPosition, err = encoding.AsInt32(j0["slotPosition"])
	if err != nil {
		return err
	}
	return nil
}

func (t *_TagEvent) Decode(value map[string]any, caller idl.Caller) error {
	t.Event = valobj.For[event.Event]()
	err := t.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("tags", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["tags"])
	if err != nil {
		return err
	}
	t.tags = make([]TagChangeInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 TagChangeInfo
		err = e0.Decode(a0, caller)
		if err != nil {
			return err
		}
		t.tags = append(t.tags, e0)
	}
	err = encoding.In("allTags", value)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](value["allTags"])
	if err != nil {
		return err
	}
	t.allTags = make([]TagInfo, 0, len(s1))
	for _, a1 := range s1 {
		var e1 TagInfo
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		t.allTags = append(t.allTags, e1)
	}
	return nil
}

func (t *_TagAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	t.TagEvent = valobj.For[TagEvent]()
	err := t.TagEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (t *_TagRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	t.TagEvent = valobj.For[TagEvent]()
	err := t.TagEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (f *_FirmwareUpdateStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.Event = valobj.For[event.Event]()
	err := f.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("state", value)
	if err != nil {
		return err
	}
	f.state, err = encoding.AsEnum[FirmwareUpdateState](value["state"])
	if err != nil {
		return err
	}
	return nil
}

func (b *_BladeOverflowChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	b.Event = valobj.For[event.Event]()
	err := b.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("overflow", value)
	if err != nil {
		return err
	}
	b.overflow, err = encoding.Is[bool](value["overflow"])
	if err != nil {
		return err
	}
	return nil
}

func (o *_OrientationChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	o.Event = valobj.For[event.Event]()
	err := o.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldOrientation", value)
	if err != nil {
		return err
	}
	o.oldOrientation, err = encoding.AsEnum[assetstripconfig.Orientation](value["oldOrientation"])
	if err != nil {
		return err
	}
	err = encoding.In("newOrientation", value)
	if err != nil {
		return err
	}
	o.newOrientation, err = encoding.AsEnum[assetstripconfig.Orientation](value["newOrientation"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_CompositionChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldComponentCount", value)
	if err != nil {
		return err
	}
	c.oldComponentCount, err = encoding.AsInt32(value["oldComponentCount"])
	if err != nil {
		return err
	}
	err = encoding.In("newComponentCount", value)
	if err != nil {
		return err
	}
	c.newComponentCount, err = encoding.AsInt32(value["newComponentCount"])
	if err != nil {
		return err
	}
	return nil
}
