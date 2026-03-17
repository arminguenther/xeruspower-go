// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package webcam

import (
	"github.com/arminguenther/xeruspower-go/v40020/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
)

func (f *Format) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["width"] = f.Width
	j0["height"] = f.Height
	j0["pixelFormat"] = f.PixelFormat
	return j0
}

func (f *Format) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("width", j0)
	if err != nil {
		return err
	}
	f.Width, err = encoding.AsInt32(j0["width"])
	if err != nil {
		return err
	}
	err = encoding.In("height", j0)
	if err != nil {
		return err
	}
	f.Height, err = encoding.AsInt32(j0["height"])
	if err != nil {
		return err
	}
	err = encoding.In("pixelFormat", j0)
	if err != nil {
		return err
	}
	f.PixelFormat, err = encoding.AsEnum[PixelFormat](j0["pixelFormat"])
	if err != nil {
		return err
	}
	return nil
}

func (c *Controls) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["brightness"] = c.Brightness
	j0["contrast"] = c.Contrast
	j0["saturation"] = c.Saturation
	j0["gain"] = c.Gain
	j0["gamma"] = c.Gamma
	j0["powerLineFrequency"] = c.PowerLineFrequency
	return j0
}

func (c *Controls) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("brightness", j0)
	if err != nil {
		return err
	}
	c.Brightness, err = encoding.AsInt32(j0["brightness"])
	if err != nil {
		return err
	}
	err = encoding.In("contrast", j0)
	if err != nil {
		return err
	}
	c.Contrast, err = encoding.AsInt32(j0["contrast"])
	if err != nil {
		return err
	}
	err = encoding.In("saturation", j0)
	if err != nil {
		return err
	}
	c.Saturation, err = encoding.AsInt32(j0["saturation"])
	if err != nil {
		return err
	}
	err = encoding.In("gain", j0)
	if err != nil {
		return err
	}
	c.Gain, err = encoding.AsInt32(j0["gain"])
	if err != nil {
		return err
	}
	err = encoding.In("gamma", j0)
	if err != nil {
		return err
	}
	c.Gamma, err = encoding.AsInt32(j0["gamma"])
	if err != nil {
		return err
	}
	err = encoding.In("powerLineFrequency", j0)
	if err != nil {
		return err
	}
	c.PowerLineFrequency, err = encoding.AsEnum[PowerLineFrequency](j0["powerLineFrequency"])
	if err != nil {
		return err
	}
	return nil
}

func (l *Location) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["name"] = l.Name
	j0["x"] = l.X
	j0["y"] = l.Y
	j0["z"] = l.Z
	return j0
}

func (l *Location) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	l.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("x", j0)
	if err != nil {
		return err
	}
	l.X, err = encoding.Is[string](j0["x"])
	if err != nil {
		return err
	}
	err = encoding.In("y", j0)
	if err != nil {
		return err
	}
	l.Y, err = encoding.Is[string](j0["y"])
	if err != nil {
		return err
	}
	err = encoding.In("z", j0)
	if err != nil {
		return err
	}
	l.Z, err = encoding.Is[string](j0["z"])
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageMetaData) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["format"] = i.Format.Encode()
	j0["timestamp"] = i.Timestamp
	j0["location"] = i.Location.Encode()
	return j0
}

func (i *ImageMetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("format", j0)
	if err != nil {
		return err
	}
	err = i.Format.Decode(j0["format"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	i.Timestamp, err = encoding.AsInt64(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("location", j0)
	if err != nil {
		return err
	}
	err = i.Location.Decode(j0["location"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *Image) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["meta"] = i.Meta.Encode()
	j0["data"] = i.Data
	return j0
}

func (i *Image) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("meta", j0)
	if err != nil {
		return err
	}
	err = i.Meta.Decode(j0["meta"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("data", j0)
	if err != nil {
		return err
	}
	i.Data, err = encoding.Is[string](j0["data"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["format"] = s.Format.Encode()
	j0["controls"] = s.Controls.Encode()
	j0["name"] = s.Name
	j0["location"] = s.Location.Encode()
	j0["refreshInterval"] = s.RefreshInterval
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("format", j0)
	if err != nil {
		return err
	}
	err = s.Format.Decode(j0["format"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("controls", j0)
	if err != nil {
		return err
	}
	err = s.Controls.Decode(j0["controls"], caller)
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
	err = encoding.In("location", j0)
	if err != nil {
		return err
	}
	err = s.Location.Decode(j0["location"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("refreshInterval", j0)
	if err != nil {
		return err
	}
	s.RefreshInterval, err = encoding.AsInt32(j0["refreshInterval"])
	if err != nil {
		return err
	}
	return nil
}

func (i *Information) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["id"] = i.Id
	j0["model"] = i.Model
	s1 := make([]any, 0, len(i.SupportedFormats))
	for _, e1 := range i.SupportedFormats {
		s1 = append(s1, e1.Encode())
	}
	j0["supportedFormats"] = s1
	return j0
}

func (i *Information) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	i.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("model", j0)
	if err != nil {
		return err
	}
	i.Model, err = encoding.Is[string](j0["model"])
	if err != nil {
		return err
	}
	err = encoding.In("supportedFormats", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["supportedFormats"])
	if err != nil {
		return err
	}
	i.SupportedFormats = make([]Format, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Format
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.SupportedFormats = append(i.SupportedFormats, e1)
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
