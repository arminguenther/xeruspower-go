// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zeroconf

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
)

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["mdnsEnabled"] = s.MdnsEnabled
	j0["llmnrEnabled"] = s.LlmnrEnabled
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("mdnsEnabled", j0)
	if err != nil {
		return err
	}
	s.MdnsEnabled, err = encoding.Is[bool](j0["mdnsEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("llmnrEnabled", j0)
	if err != nil {
		return err
	}
	s.LlmnrEnabled, err = encoding.Is[bool](j0["llmnrEnabled"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
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
