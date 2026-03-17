// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package timereventmanager

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
)

func (r *Range) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["start"] = r.Start
	j0["end"] = r.End
	j0["step"] = r.Step
	return j0
}

func (r *Range) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("start", j0)
	if err != nil {
		return err
	}
	r.Start, err = encoding.AsInt32(j0["start"])
	if err != nil {
		return err
	}
	err = encoding.In("end", j0)
	if err != nil {
		return err
	}
	r.End, err = encoding.AsInt32(j0["end"])
	if err != nil {
		return err
	}
	err = encoding.In("step", j0)
	if err != nil {
		return err
	}
	r.Step, err = encoding.AsInt32(j0["step"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["enabled"] = s.Enabled
	s1 := make([]any, 0, len(s.Minute))
	for _, e1 := range s.Minute {
		s1 = append(s1, e1.Encode())
	}
	j0["minute"] = s1
	s2 := make([]any, 0, len(s.Hour))
	for _, e2 := range s.Hour {
		s2 = append(s2, e2.Encode())
	}
	j0["hour"] = s2
	s3 := make([]any, 0, len(s.DayOfMonth))
	for _, e3 := range s.DayOfMonth {
		s3 = append(s3, e3.Encode())
	}
	j0["dayOfMonth"] = s3
	s4 := make([]any, 0, len(s.Month))
	for _, e4 := range s.Month {
		s4 = append(s4, e4.Encode())
	}
	j0["month"] = s4
	s5 := make([]any, 0, len(s.DayOfWeek))
	for _, e5 := range s.DayOfWeek {
		s5 = append(s5, e5.Encode())
	}
	j0["dayOfWeek"] = s5
	return j0
}

func (s *Schedule) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	s.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("minute", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["minute"])
	if err != nil {
		return err
	}
	s.Minute = make([]Range, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Range
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		s.Minute = append(s.Minute, e1)
	}
	err = encoding.In("hour", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["hour"])
	if err != nil {
		return err
	}
	s.Hour = make([]Range, 0, len(s2))
	for _, a2 := range s2 {
		var e2 Range
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		s.Hour = append(s.Hour, e2)
	}
	err = encoding.In("dayOfMonth", j0)
	if err != nil {
		return err
	}
	var s3 []any
	s3, err = encoding.Is[[]any](j0["dayOfMonth"])
	if err != nil {
		return err
	}
	s.DayOfMonth = make([]Range, 0, len(s3))
	for _, a3 := range s3 {
		var e3 Range
		err = e3.Decode(a3, caller)
		if err != nil {
			return err
		}
		s.DayOfMonth = append(s.DayOfMonth, e3)
	}
	err = encoding.In("month", j0)
	if err != nil {
		return err
	}
	var s4 []any
	s4, err = encoding.Is[[]any](j0["month"])
	if err != nil {
		return err
	}
	s.Month = make([]Range, 0, len(s4))
	for _, a4 := range s4 {
		var e4 Range
		err = e4.Decode(a4, caller)
		if err != nil {
			return err
		}
		s.Month = append(s.Month, e4)
	}
	err = encoding.In("dayOfWeek", j0)
	if err != nil {
		return err
	}
	var s5 []any
	s5, err = encoding.Is[[]any](j0["dayOfWeek"])
	if err != nil {
		return err
	}
	s.DayOfWeek = make([]Range, 0, len(s5))
	for _, a5 := range s5 {
		var e5 Range
		err = e5.Decode(a5, caller)
		if err != nil {
			return err
		}
		s.DayOfWeek = append(s.DayOfWeek, e5)
	}
	return nil
}

func (t *TimerEvent) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["eventId"] = encoding.NonNilSlice(t.EventId)
	j0["executionTime"] = t.ExecutionTime.Encode()
	return j0
}

func (t *TimerEvent) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("eventId", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["eventId"])
	if err != nil {
		return err
	}
	t.EventId = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		t.EventId = append(t.EventId, e1)
	}
	err = encoding.In("executionTime", j0)
	if err != nil {
		return err
	}
	err = t.ExecutionTime.Decode(j0["executionTime"], caller)
	if err != nil {
		return err
	}
	return nil
}
