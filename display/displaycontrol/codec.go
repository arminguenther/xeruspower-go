// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package displaycontrol

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["fwAppVersion"] = i.FwAppVersion
	j0["fwBootVersion"] = i.FwBootVersion
	j0["orientation"] = i.Orientation
	j0["width"] = i.Width
	j0["height"] = i.Height
	j0["versionAvailable"] = i.VersionAvailable
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("fwAppVersion", j0)
	if err != nil {
		return err
	}
	i.FwAppVersion, err = encoding.AsInt32(j0["fwAppVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("fwBootVersion", j0)
	if err != nil {
		return err
	}
	i.FwBootVersion, err = encoding.AsInt32(j0["fwBootVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("orientation", j0)
	if err != nil {
		return err
	}
	i.Orientation, err = encoding.AsInt32(j0["orientation"])
	if err != nil {
		return err
	}
	err = encoding.In("width", j0)
	if err != nil {
		return err
	}
	i.Width, err = encoding.AsInt32(j0["width"])
	if err != nil {
		return err
	}
	err = encoding.In("height", j0)
	if err != nil {
		return err
	}
	i.Height, err = encoding.AsInt32(j0["height"])
	if err != nil {
		return err
	}
	err = encoding.In("versionAvailable", j0)
	if err != nil {
		return err
	}
	i.VersionAvailable, err = encoding.Is[bool](j0["versionAvailable"])
	if err != nil {
		return err
	}
	return nil
}

func (d *DefaultViewItem) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["id"] = d.Id
	j0["description"] = d.Description
	return j0
}

func (d *DefaultViewItem) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	d.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	d.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["defaultView"] = s.DefaultView
	j0["backlightTimeout"] = s.BacklightTimeout
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("defaultView", j0)
	if err != nil {
		return err
	}
	s.DefaultView, err = encoding.Is[string](j0["defaultView"])
	if err != nil {
		return err
	}
	err = encoding.In("backlightTimeout", j0)
	if err != nil {
		return err
	}
	s.BacklightTimeout, err = encoding.AsInt32(j0["backlightTimeout"])
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
