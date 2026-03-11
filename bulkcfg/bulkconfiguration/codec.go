// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkconfiguration

import (
	"github.com/arminguenther/xeruspower-go/v40410/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
)

func (f *Filter) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["name"] = f.Name
	j0["displayName"] = f.DisplayName
	j0["noOverride"] = f.NoOverride
	j0["bulkOnly"] = f.BulkOnly
	j0["ruleSpecs"] = encoding.NonNilSlice(f.RuleSpecs)
	return j0
}

func (f *Filter) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	f.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("displayName", j0)
	if err != nil {
		return err
	}
	f.DisplayName, err = encoding.Is[string](j0["displayName"])
	if err != nil {
		return err
	}
	err = encoding.In("noOverride", j0)
	if err != nil {
		return err
	}
	f.NoOverride, err = encoding.Is[bool](j0["noOverride"])
	if err != nil {
		return err
	}
	err = encoding.In("bulkOnly", j0)
	if err != nil {
		return err
	}
	f.BulkOnly, err = encoding.Is[bool](j0["bulkOnly"])
	if err != nil {
		return err
	}
	err = encoding.In("ruleSpecs", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["ruleSpecs"])
	if err != nil {
		return err
	}
	f.RuleSpecs = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		f.RuleSpecs = append(f.RuleSpecs, e1)
	}
	return nil
}

func (f *FilterProfile) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["name"] = f.Name
	j0["description"] = f.Description
	s1 := make([]any, 0, len(f.FilterNameToTypeMap))
	for k1, v1 := range f.FilterNameToTypeMap {
		s1 = append(s1, map[string]any{"key": k1, "value": v1})
	}
	j0["filterNameToTypeMap"] = s1
	return j0
}

func (f *FilterProfile) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	f.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	f.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	err = encoding.In("filterNameToTypeMap", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["filterNameToTypeMap"])
	f.FilterNameToTypeMap = make(map[string]FilterType, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 FilterType
		v1, err = encoding.AsEnum[FilterType](b1)
		if err != nil {
			return err
		}
		f.FilterNameToTypeMap[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	s1 := make([]any, 0, len(s.FilterProfiles))
	for _, e1 := range s.FilterProfiles {
		s1 = append(s1, e1.Encode())
	}
	j0["filterProfiles"] = s1
	j0["defaultProfileName"] = s.DefaultProfileName
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("filterProfiles", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["filterProfiles"])
	if err != nil {
		return err
	}
	s.FilterProfiles = make([]FilterProfile, 0, len(s1))
	for _, a1 := range s1 {
		var e1 FilterProfile
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		s.FilterProfiles = append(s.FilterProfiles, e1)
	}
	err = encoding.In("defaultProfileName", j0)
	if err != nil {
		return err
	}
	s.DefaultProfileName, err = encoding.Is[string](j0["defaultProfileName"])
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
	return nil
}

func (s *_SavedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("isBackup", value)
	if err != nil {
		return err
	}
	s.isBackup, err = encoding.Is[bool](value["isBackup"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RestoredEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("isBackup", value)
	if err != nil {
		return err
	}
	r.isBackup, err = encoding.Is[bool](value["isBackup"])
	if err != nil {
		return err
	}
	return nil
}
